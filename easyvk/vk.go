// Package easyvk provides you simple way
// to work with VK API.
package easyvk

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const apiURL = "https://api.vk.com/method/"

// VK defines a set of functions for
// working with VK API.
type VK struct {
	AccessToken string
	Version     string
	Account     Account
	Fave        Fave
	Photos      Photos
	Status      Status
	Upload      Upload
	Wall        Wall
}

// WithToken helps to initialize your
// VK object with token.
func WithToken(token string) VK {
	vk := VK{}
	vk.AccessToken = token
	vk.Version = "5.63"
	vk.Account = Account{&vk }
	vk.Fave = Fave{&vk }
	vk.Photos = Photos{&vk }
	vk.Status = Status{&vk }
	vk.Upload = Upload{}
	vk.Wall = Wall{&vk }
	return vk
}

// Request provides access to VK API methods.
func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(apiURL + method)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if string(body[2:7]) == "error" {
		var e VKError
		err = json.Unmarshal(body, &e)
		return nil, fmt.Errorf("Code %d: %s", e.Error.Code, e.Error.Message)
	}

	return body, nil
}

type response struct {
	Response int `json:"response"`
}
