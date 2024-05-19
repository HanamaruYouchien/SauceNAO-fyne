package saucenao

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

const (
	baseUrl              = "https://saucenao.com/search.php?"
	fieldApikey          = "api_key"
	fieldOutputType      = "output_type"
	fieldOutputTypeValue = "2"
	fieldImageUrl        = "url"
	fieldImageFile       = "file"
)

func SearchByURL(apikey string, imageUrl string) (*Response, error) {
	raw, err := SearchRawByURL(apikey, imageUrl)
	if err != nil {
		return nil, err
	}

	return ParseResponse(raw)
}

func SearchByFile(apikey string, imageFile []byte) (*Response, error) {
	raw, err := SearchRawByFile(apikey, imageFile)
	if err != nil {
		return nil, err
	}

	return ParseResponse(raw)
}

func SearchRawByURL(apikey string, imageUrl string) ([]byte, error) {
	q := url.Values{}
	q.Add(fieldOutputType, fieldOutputTypeValue)
	q.Add(fieldApikey, apikey)
	q.Add(fieldImageUrl, imageUrl)

	resp, err := http.Get(baseUrl + q.Encode())
	if err != nil {
		return []byte{}, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return respBytes, nil
}

func SearchRawByFile(apikey string, imageFile []byte) ([]byte, error) {
	q := url.Values{}
	q.Add(fieldOutputType, fieldOutputTypeValue)
	q.Add(fieldApikey, apikey)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormFile(fieldImageFile, "image") // filename cannot be empty
	if err != nil {
		return []byte{}, nil
	}
	if _, err := fw.Write(imageFile); err != nil {
		return []byte{}, nil
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, baseUrl+q.Encode(), bytes.NewReader(body.Bytes()))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return respBytes, nil
}

func ParseResponse(raw []byte) (*Response, error) {
	resp := &Response{}
	if err := json.Unmarshal(raw, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
