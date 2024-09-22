package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type stdTelegramAPIClient struct {
	baseUrl string
}

func decodeResponseOnSuccess(body io.Reader, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}

func (c *stdTelegramAPIClient) Get(command string, v interface{}) error {
	url := c.baseUrl + command
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponseOnSuccess(resp.Body, v)
}

func (c *stdTelegramAPIClient) Post(command string, data url.Values, v interface{}) error {
	url := c.baseUrl + command
	resp, err := http.PostForm(url, data)
	if err != nil {
		return err
	}
	//log.Println("REQ == : ", url, data)
	defer resp.Body.Close()
	return decodeResponseOnSuccess(resp.Body, v)
}

func NewStdTelegramAPIClient(token string) (TelegramAPIClient, error) {
	if len(token) == 0 {
		return nil, errors.New("fatal: token cannot be empty")
	}

	return &stdTelegramAPIClient{baseUrl: "https://api.telegram.org/bot" + token + "/"}, nil
}

func (c *stdTelegramAPIClient) PostPhoto(command string, data url.Values, path string, v interface{}) error {
	file, _ := os.Open(path)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("png", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	url := c.baseUrl + command

	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("Post photo error", err)
	}
	log.Println(url)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	log.Println("REQ: ", r.Form)
	_, e := client.Do(r)
	if e != nil {
		log.Println("Post photo eroror 2", err)
	}

	return nil
}
