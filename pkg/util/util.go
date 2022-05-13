package util

import (
	"time"
)

type ResultWriter interface {
	Write(section string, name string, duration time.Duration, err error)
}
