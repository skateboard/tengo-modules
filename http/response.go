package http

import (
	"fmt"

	"github.com/d5/tengo/v2"
)

type Response struct {
	tengo.ObjectImpl

	body       []byte
	statusCode int
	cookies    map[string]tengo.Object
}

func (r *Response) TypeName() string {
	return "response"
}

func (r *Response) String() string {
	return fmt.Sprintf("%d %s", r.statusCode, r.body)
}

func (r *Response) IndexGet(index tengo.Object) (tengo.Object, error) {
	k, _ := index.(*tengo.String)
	switch k.Value {
	case "body":
		return &tengo.UserFunction{
			Name: "body",
			Value: func(_ ...tengo.Object) (tengo.Object, error) {
				return &tengo.Bytes{Value: r.body}, nil
			},
		}, nil
	case "status_code":
		return &tengo.UserFunction{
			Name: "status_code",
			Value: func(args ...tengo.Object) (tengo.Object, error) {
				return &tengo.Int{Value: int64(r.statusCode)}, nil
			},
		}, nil
	case "cookies":
		return &tengo.UserFunction{
			Name: "cookies",
			Value: func(args ...tengo.Object) (tengo.Object, error) {
				return &tengo.Map{Value: r.cookies}, nil
			},
		}, nil
	}

	return tengo.UndefinedValue, nil
}
