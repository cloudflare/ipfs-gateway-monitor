package pinata

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

var (
	// According to https://docs.pinata.cloud/api-pinning/pin-file, the key has
	// to be "file".
	fileKey  = "file"
	fileName = "filename"
	url      = "https://api.pinata.cloud/pinning/pinFileToIPFS"
)

type Client struct {
	jwt string
}

type PinFileToIpfsResponse struct {
	IpfsHash  string `json:"IpfsHash"`
	PinSize   int    `json:"PinSize"`
	Timestamp string `json:"Timestamp"`
}

func NewClient(jwt string) *Client {
	return &Client{jwt}
}

// Return the CID, if the pinning is successful.
func (client *Client) PinFileToIpfs(content []byte) (string, error) {
	// Prepare a form that you will submit to the URL.
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile(fileKey, fileName)
	if err != nil {
		return "", err
	}
	_, err = fw.Write(content)
	if err != nil {
		return "", err
	}

	// Don't forget to close the multipart writer. If you don't close it, your
	// request will be missing the terminating boundary.
	w.Close()

	// Upload the content. Don't forget to set the content type, this will
	// contain the boundary.
	bearer := "Bearer " + client.jwt
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", w.FormDataContentType())

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var body PinFileToIpfsResponse
	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		return "", err
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Pinata: The status code is not 200")
	}

	return body.IpfsHash, nil
}
