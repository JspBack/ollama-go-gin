package utils

import (
	"ai-thing/schemas"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func DoRequest(client *http.Client, data schemas.AiRequest) (string, error) {
	r, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req := http.Request{
		Method: http.MethodPost,
		URL:    &url.URL{Scheme: os.Getenv("OLLAMA_SCHEME"), Host: os.Getenv("OLLAMA_URL"), Path: "/api/chat"},
		Body:   io.NopCloser(bytes.NewBuffer(r)),
	}

	resp, err := client.Do(&req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return string(body), errors.New("bad status code: " + resp.Status)
	}

	return string(body), nil
}

func ValidateForms(c *gin.Context, forms ...string) error {
	for _, form := range forms {
		if c.PostForm(form) == "" {
			return errors.New("missing form: " + form)
		}
	}

	return nil
}
