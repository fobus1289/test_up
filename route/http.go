package route

import (
	"net/http"
)

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

func handleFunc(url string, method string, action func(client *Client)) *route {
	rou := new(route)

	go http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
		client := new(Client)
		client.request = request
		client.response = &writer

		if !valid(writer, request, rou, method, *client) {
			return
		}

		action(client)
	})

	return rou
}

func Get(url string, action func(client *Client)) *route {
	return handleFunc(url, "GET", action)
}

func Post(url string, action func(client *Client)) *route {
	return handleFunc(url, "POST", action)
}

func Put(url string, action func(client *Client)) *route {
	return handleFunc(url, "PUT", action)
}

func Patch(url string, action func(client *Client)) *route {
	return handleFunc(url, "PATCH", action)
}

func Delete(url string, action func(client *Client)) *route {
	return handleFunc(url, "DELETE", action)
}
