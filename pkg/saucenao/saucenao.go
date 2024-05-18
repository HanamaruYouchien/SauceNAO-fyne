package saucenao

import (
	"io"
	"net/http"
	"net/url"
)

const (
	baseUrl              = "https://saucenao.co/search.php?"
	fieldApikey          = "api_key"
	fieldOutputType      = "output_type"
	fieldOutputTypeValue = "2"
	fieldImageUrl        = "url"
)

func SearchByURL(apikey string, imageUrl string) ([]byte, error) {
	q := url.Values{}
	q.Add(fieldOutputType, fieldOutputTypeValue)
	q.Add(fieldApikey, apikey)

	v := url.Values{}
	v.Add(fieldImageUrl, imageUrl)

	resp, err := http.PostForm(baseUrl+q.Encode(), v)
	if err != nil {
		return []byte{}, err
	}

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return ret, nil
}
