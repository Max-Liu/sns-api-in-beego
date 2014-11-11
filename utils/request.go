package helper

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func MakeRequest(query map[string]string, urlStr, method string) *http.Request {
	data := url.Values{}
	for k, v := range query {
		data.Add(k, v)
	}
	r, err := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	return r
}

func MakeUploadRequest(query map[string]string, urlStr, method, fileField, filePath string) *http.Request {
	var b bytes.Buffer

	w := multipart.NewWriter(&b)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range query {
		w.WriteField(k, v)
	}

	fw, err := w.CreateFormFile(fileField, filePath)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = io.Copy(fw, f); err != nil {
		log.Fatal(err)
	}
	w.Close()

	request, _ := http.NewRequest(method, urlStr, &b)
	request.Header.Set("Content-Type", w.FormDataContentType())
	request.Header.Set("Accept", "application/json")
	return request
}
