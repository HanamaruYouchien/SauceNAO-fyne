package saucenao

import (
	"bytes"
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

func SearchByURL(apikey string, imageUrl string) ([]byte, error) {
	q := url.Values{}
	q.Add(fieldOutputType, fieldOutputTypeValue)
	q.Add(fieldApikey, apikey)
	q.Add(fieldImageUrl, imageUrl)

	resp, err := http.Get(baseUrl + q.Encode())
	if err != nil {
		return []byte{}, err
	}

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return ret, nil
}

func SearchByFile(apikey string, imageFile []byte) ([]byte, error) {
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

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return ret, nil
}
