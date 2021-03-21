package route

import (
	"net/http"
)

func Get(url string) *route {

	rou := new(route)

	http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {

	})

	return rou
}

func valid(writer http.ResponseWriter, request *http.Request, rou *route, method string, client Client) bool {
	if request.Method != method {
		writer.WriteHeader(405)
		_, _ = writer.Write([]byte("Method Not Allowed "))
		return false
	}

	middlewares := rou.middlewares

	if middlewares != nil {
		for i := 0; i < len(middlewares); i++ {
			if middlewares[i] != nil && !middlewares[i](client) {
				writer.WriteHeader(403)
				_, _ = writer.Write([]byte("Forbidden "))
				return false
			}
		}
	}

	return true
}

func Post(url string, action func(client *Client)) *route {
	rou := new(route)

	go http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		client := new(Client)
		client.request = request
		client.response = &writer
		if !valid(writer, request, rou, "POST", *client) {
			return
		}

		action(client)
	})

	return rou
}

func PostNoRet(url string, action func(writer http.ResponseWriter, request *http.Request)) *route {
	rou := new(route)

	go http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		action(writer, request)
	})

	return rou
}

func Put(url string, action func(client interface{})) {

}

func Patch(url string, action interface{}) {

}

func Delete(url string, action interface{}) {

}

func New(response http.ResponseWriter, request *http.Request) *Client {

	return &Client{
		&response,
		request,
	}
}
