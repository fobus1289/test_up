package route

type route struct {
	Url        string
	action     interface{}
	actionType string
	middlewares []func(client Client) bool
	func1      func(...interface{}) interface{}
}

func (receiver *route) Middleware(fn func(client Client) bool) *route {
	receiver.middlewares = append(receiver.middlewares, fn)
	return receiver
}
