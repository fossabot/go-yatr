package yatr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type YatrContext struct {
	APIKey string
}

type yatrResponse struct {
	Code     int      `json:"code"`
	Language string   `json:"lang"`
	Text     []string `json:"text"`
	Message  string   `json:"message"`
}

// NewContext creates new Yatr context
func NewContext(key string) *YatrContext {
	return &YatrContext{APIKey:key}
}

// Translate gets fromLang - the language into which to translate, targetLang - the language from which you want to transfer, text - text to translate and returns translated text.
// If the "fromLang" is empty, the API will determine the language automatically.
func (y *YatrContext) Translate(fromLang, targetLang, text string) (string, error) {
	var translateLang string
	if fromLang != "" {
		translateLang = fromLang+"-"+targetLang
	} else {
		translateLang = targetLang
	}
	resp, err := http.PostForm(fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?lang=%v&key=%v", translateLang, y.APIKey), url.Values{"text":{text}})
	if err != nil {
		return "", err
	}
	var result yatrResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if result.Code != 200 {
		return "", errors.New(result.Message)
	}
	return strings.Join(result.Text, "\n"), nil
}
