package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Version and BuildTime are filled in during build by the Makefile
var (
	Version   = "N/A"
	BuildTime = "N/A"
)

var (
	NanosecondsInSecond = 1000000000
)

var (
	measureBinary   = flag.String("measure-binary", "", "The executable file for the measurement.")
	metricAddr      = flag.String("metric-addr", ":8080", "The listening addr for the Prometheus metric.")
	executeInterval = flag.String("execute-interval", "", "The interval between executation with a unit suffix of \"ms\", \"s\", \"m\", or \"h\"")
)

const (
	metricsNamespace = "ipfs"
	metricsSubsystem = "monitor"
)

var (
	errorSignal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "errors",
		Help:      "Become 1 if there is an error for the scenario or 0 otherwise.",
	}, []string{"section", "name"})
	responseTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "response_time_seconds",
		Help:      "The reponse time from the IPFS gateway.",
	}, []string{"section", "name"})
	epochCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: metricsSubsystem,
		Name:      "epoch_number",
		Help:      "The epoch counter since the export started.",
	})
)

type scenario struct {
	section string
	name    string
}

func main() {
	fmt.Printf("#Version=%s, BuildTime=%s\n", Version, BuildTime)
	flag.Parse()

	args := os.Args
	measureArgs := []string{}

	for i, arg := range args {
		if arg == "--" {
			measureArgs = args[i+1:]
		}
	}

	interval, err := time.ParseDuration(*executeInterval)
	if err != nil {
		fmt.Fprintf(os.Stderr, "The execute-interval flag is in an invalid format.")
		os.Exit(1)
	}

	// Trigger the executation with 1 element buffered.
	signal := make(chan struct{})

	go func() {
		// Make sure the first round will be executed immediately.
		signal <- struct{}{}
		for {
			time.Sleep(interval)
			select {
			case signal <- struct{}{}:
			default:
				// If there is no receiver, we will sleep anyway.
			}
		}
	}()

	go func() {
		var errMap map[scenario]bool
		var rspMap map[scenario]float64
		for {
			<-signal
			// Set the metrics only at the beginning of the execution so that
			// when we average the response time, it will be reasonable.
			for sc, hasErr := range errMap {
				if hasErr {
					errorSignal.With(prometheus.Labels{"section": sc.section, "name": sc.name}).Set(1)
				} else {
					errorSignal.With(prometheus.Labels{"section": sc.section, "name": sc.name}).Set(0)
				}
			}
			for sc, rspTime := range rspMap {
				responseTime.With(prometheus.Labels{"section": sc.section, "name": sc.name}).Set(rspTime)
			}
			errMap = make(map[scenario]bool)
			rspMap = make(map[scenario]float64)

			fmt.Printf("#Start executing, %s\n", time.Now())
			// The Prometheus metric indicating the new epoch.
			epochCounter.Inc()
			cmd := exec.Command(*measureBinary, measureArgs...)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Printf("#Execute error: %s, time=%d\n", err, time.Now())
				continue
			}
			err = cmd.Start()
			if err != nil {
				fmt.Printf("#Execute error: %s, time=%d\n", err, time.Now())
				continue
			}

			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Printf("%s\n", scanner.Text())
				// Skip the comment.
				if strings.HasPrefix(scanner.Text(), "#") {
					continue
				}

				fields := strings.SplitN(scanner.Text(), ",", 5)

				section := fields[0]
				name := fields[1]

				respTimeNanosec, err := strconv.ParseInt(fields[2], 10, 64)
				if err != nil {
					fmt.Printf("#Error while executing: cannot parse the response time: %s, time=%d\n", err, time.Now())
					break
				}

				scenarioTime := time.Duration(respTimeNanosec)
				scenarioErr := fields[3]

				sc := scenario{section, name}
				if scenarioErr != "nil" {
					errMap[sc] = true
				} else {
					errMap[sc] = false
				}
				rspMap[sc] = float64(scenarioTime.Nanoseconds()) / float64(NanosecondsInSecond)
			}

			err = cmd.Wait()
			if err != nil {
				fmt.Printf("#Error while executing: %s, time=%d\n", err, time.Now())
				continue
			}
			fmt.Printf("#Finished executing, %s\n", time.Now())
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(*metricAddr, nil)
	if err != nil {
		fmt.Printf("#Listening error on the metric server: %s\n", err)
	}
}
