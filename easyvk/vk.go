package easyvk

import (
	"net/url"
	"net/http"
	"io/ioutil"
)

const API_URL = "https://api.vk.com/method/"

type VK struct {
	AccessToken string
	Account     Account
}

func WithToken(token string) VK {
	vk := VK{}
	vk.AccessToken = token
	vk.Account = Account{ &vk }
	return vk
}

func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(API_URL + method)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("access_token", vk.AccessToken)
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

	return body, nil
}
