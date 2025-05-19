package http

import (
	"fmt"

	"github.com/d5/tengo/v2"
)

type Request struct {
	tengo.ObjectImpl

	method  string
	url     string
	body    []byte
	headers map[string][]string
	cookies map[string]string
}

func (r *Request) TypeName() string {
	return "request"
}

func (r *Request) String() string {
	return fmt.Sprintf("%s %s %s", r.method, r.url, r.body)
}

func (r *Request) IndexGet(index tengo.Object) (tengo.Object, error) {
	k, _ := index.(*tengo.String)
	switch k.Value {
	case "set_cookie":
		return &tengo.UserFunction{
			Name: "set_cookie",
			Value: func(args ...tengo.Object) (tengo.Object, error) {
				name, _ := tengo.ToString(args[0])
				value, _ := tengo.ToString(args[1])

				if r.cookies == nil {
					r.cookies = make(map[string]string)
				}

				r.cookies[name] = value

				return nil, nil
			},
		}, nil
	case "set_header":
		return &tengo.UserFunction{
			Name: "set_header",
			Value: func(args ...tengo.Object) (tengo.Object, error) {
				name, _ := tengo.ToString(args[0])
				value, _ := tengo.ToString(args[1])

				if r.headers == nil {
					r.headers = make(map[string][]string)
				}

				r.headers[name] = append(r.headers[name], value)

				return nil, nil
			},
		}, nil
	case "url":
		return &tengo.UserFunction{
			Name: "url",
			Value: func(_ ...tengo.Object) (tengo.Object, error) {
				return &tengo.String{Value: r.url}, nil
			},
		}, nil
	case "method":
		return &tengo.UserFunction{
			Name: "method",
			Value: func(_ ...tengo.Object) (tengo.Object, error) {
				return &tengo.String{Value: r.method}, nil
			},
		}, nil
	case "body":
		return &tengo.UserFunction{
			Name: "body",
			Value: func(_ ...tengo.Object) (tengo.Object, error) {
				return &tengo.Bytes{Value: r.body}, nil
			},
		}, nil
	}

	return tengo.UndefinedValue, nil
}
