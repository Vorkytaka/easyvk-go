package easyvk

import (
	"bytes"
	"mime/multipart"
	"os"
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
)

// An Upload describes a set of methods
// that helps with update to VK servers.
type Upload struct{}

// A PhotoWall describes an info
// about uploaded photo.
type PhotoWall struct {
	Server int `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// PhotoWall upload file (on filePath) to given url.
// Return info about uploaded photo.
func (u *Upload) PhotoWall(url, filePath string) (PhotoWall, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("photo", filePath)
	if err != nil {
		return PhotoWall{}, err
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return PhotoWall{}, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return PhotoWall{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return PhotoWall{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PhotoWall{}, err
	}

	var uploaded PhotoWall
	err = json.Unmarshal(body, &uploaded)
	if err != nil {
		return PhotoWall{}, err
	}

	return uploaded, nil
}
