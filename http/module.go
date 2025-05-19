package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/d5/tengo/v2"
)

var Module = map[string]tengo.Object{
	"request":     &tengo.UserFunction{Name: "request", Value: request},
	"new_request": &tengo.UserFunction{Name: "new_request", Value: newRequest},
	"do_request":  &tengo.UserFunction{Name: "do_request", Value: doRequest},
}

func doRequest(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) == 0 {
		return nil, tengo.ErrWrongNumArguments
	}

	r := args[0].(*Request)

	var body io.Reader

	if r.body != nil {
		body = bytes.NewBuffer(r.body)
	} else {
		body = nil
	}

	req, err := http.NewRequest(r.method, r.url, body)
	if err != nil {
		return nil, err
	}

	req.Header = r.headers

	if r.cookies != nil {
		for cookieName, cookieValue := range r.cookies {
			req.AddCookie(&http.Cookie{
				Name:  cookieName,
				Value: cookieValue,
			})
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		body:       b,
		statusCode: res.StatusCode,
		cookies:    make(map[string]tengo.Object),
	}

	for _, cookie := range res.Cookies() {
		response.cookies[cookie.Name] = &tengo.String{Value: cookie.Value}
	}

	return response, nil
}

func newRequest(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) == 0 {
		return nil, tengo.ErrWrongNumArguments
	}

	method, ok := tengo.ToString(args[0])
	if !ok {
		return nil, errors.New("failed to find method")
	}

	url, ok := tengo.ToString(args[1])
	if !ok {
		return nil, errors.New("failed to find url")
	}

	var body []byte

	rB, ok := tengo.ToByteSlice(args[2])
	if ok {
		body = rB
	} else {
		body = nil
	}

	return &Request{
		url:    url,
		method: method,
		body:   body,
	}, nil
}

func request(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) == 0 {
		return nil, tengo.ErrWrongNumArguments
	}

	method, ok := tengo.ToString(args[0])
	if !ok {
		return nil, errors.New("failed to find method")
	}

	url, ok := tengo.ToString(args[1])
	if !ok {
		return nil, errors.New("failed to find url")
	}

	var body io.Reader

	rB, ok := tengo.ToByteSlice(args[2])
	if ok {
		body = bytes.NewBuffer(rB)
	} else {
		body = nil
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{
		body:       b,
		statusCode: res.StatusCode,
		cookies:    make(map[string]tengo.Object),
	}

	for _, cookie := range res.Cookies() {
		r.cookies[cookie.Name] = &tengo.String{Value: cookie.Value}
	}

	return r, nil
}
