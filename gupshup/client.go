package gupshup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	MessageUrl  = "https://api.gupshup.io/sm/api/v1/msg"
	ContentType = "application/x-www-form-urlencoded"
)

var (
	HttpClient httpClient
	Http       httpBuilder
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpBuilder interface {
	BuildRequest(method, url string, body io.Reader) (*http.Request, error)
}

type httpRequest struct{}

func (r *httpRequest) BuildRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

type Client struct {
	ApiKey  string
	AppName string
}

type Response struct {
	Status    string `json:"status"`
	MessageId string `json:"messageId"`
}

type TextRequest struct {
	Source      string
	Destination string
	Text        string
}

type ImageRequest struct {
	Source      string
	Destination string
	Image       string
	Preview     string
	Caption     string
}

type DocumentRequest struct {
	Source      string
	Destination string
	Url         string
	Filename    string
}

type AudioRequest struct {
	Source      string
	Destination string
	Url         string
}

type VideoRequest struct {
	Source      string
	Destination string
	Url         string
	Caption     string
}

type StickerRequest struct {
	Source      string
	Destination string
	Url         string
}

type InteractiveMessageRequest struct {
	Source             string
	Destination        string
	InteractiveMessage InteractiveMessage
}

type InteractiveMessage struct {
	Type          string                     `json:"type"`
	Title         string                     `json:"title"`
	Body          string                     `json:"body"`
	MessageId     string                     `json:"msgid"`
	GlobalButtons []InteractiveGlobalButtons `json:"globalButtons"`
	Items         []InteractiveMessageItem   `json:"items"`
}

type InteractiveGlobalButtons struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type InteractiveMessageItem struct {
	Title    string                      `json:"title"`
	Subtitle string                      `json:"subtitle"`
	Options  []InteractiveMessageOptions `json:"options"`
}

type InteractiveMessageOptions struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Postback    string `json:"postbackText"`
}

func init() {
	HttpClient = &http.Client{}
	Http = &httpRequest{}
}

func (g *Client) SendText(request TextRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
			"type": "text",
			"text": "%s"
		}`, request.Text)},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) SendImage(request ImageRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
				"type": "image",
				"originalUrl": "%s",
				"previewUrl": "%s",
				"caption": "%s"
			}`, request.Image, request.Preview, request.Caption)},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) SendDocument(request DocumentRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
				"type": "file",
				"url": "%s",
			    "filename": "%s"
			}`, request.Url, request.Filename)},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) SendAudio(request AudioRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
				"type": "audio",
				"url": "%s"
			}`, request.Url)},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) SendVideo(request VideoRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
				"type": "video",
				"url": "%s",
				"caption": "%s"
			}`, request.Url, request.Caption)},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) SendSticker(request StickerRequest) (Response, error) {
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message": {fmt.Sprintf(`{
				"type": "sticker",
				"url": "%s"
			}`, request.Url)},
	}

	return g.sendRequest(requestBody)
}

func NewInteractiveMessage(body, title, messageId string,
	button []InteractiveGlobalButtons, items []InteractiveMessageItem) InteractiveMessage {
	return InteractiveMessage{
		Type:          "list",
		Title:         title,
		Body:          body,
		MessageId:     messageId,
		GlobalButtons: button,
		Items:         items,
	}
}

func NewInteractiveMessageItem(title, subtitle string, options []InteractiveMessageOptions) InteractiveMessageItem {
	return InteractiveMessageItem{
		Title:    title,
		Subtitle: subtitle,
		Options:  options,
	}
}

func NewInteractiveMessageOption(title, description, postback string) InteractiveMessageOptions {
	return InteractiveMessageOptions{
		Type:        "text",
		Title:       title,
		Description: description,
		Postback:    postback,
	}
}

func NewButton(title string) InteractiveGlobalButtons {
	return InteractiveGlobalButtons{
		Type:  "text",
		Title: title,
	}
}

func (g *Client) SendInteractiveMessage(request InteractiveMessageRequest) (Response, error) {
	message, _ := json.Marshal(request.InteractiveMessage)
	requestBody := url.Values{
		"channel":     {"whatsapp"},
		"source":      {request.Source},
		"destination": {request.Destination},
		"src.name":    {g.AppName},
		"message":     {fmt.Sprintf(`%s`, string(message))},
	}

	return g.sendRequest(requestBody)
}

func (g *Client) sendRequest(requestBody url.Values) (Response, error) {
	var response Response
	httpRequest, err := Http.BuildRequest("POST", MessageUrl, strings.NewReader(requestBody.Encode()))
	if err != nil {
		return response, err
	}
	httpRequest.Header.Set("Content-Type", ContentType)
	httpRequest.Header.Set("apikey", g.ApiKey)

	resp, err := HttpClient.Do(httpRequest)

	if err != nil {
		return response, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}
