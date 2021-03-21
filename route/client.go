package route

import "net/http"

type Client struct {
	response *http.ResponseWriter
	request  *http.Request
}

func (client *Client) Request() *http.Request {
	return client.request
}

func (client *Client) Response() http.ResponseWriter {
	return *client.response
}

func (client *Client) Send(message string) {
	(*client.response).Write([]byte(message))
}
