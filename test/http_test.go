package test

import (
	"context"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/skateboard/tengomodules"
)

func TestHTTP(t *testing.T) {
	script := tengo.NewScript([]byte(
		`
		fmt := import("fmt")
		http := import("http")
		json := import("json")

		req := http.new_request("GET", "https://httpbin.io/headers", undefined)
		fmt.println(req.url())

		req.set_header("User-Agent", "TEST")

		req.set_cookie("test", "test")

		res := http.do_request(req)
		statusCode := res.status_code()
		
		fmt.println(statusCode)

		j := json.decode(res.body())
		fmt.println(j)

		cookies := res.cookies()
		fmt.println(cookies)

`))

	script.SetImports(tengomodules.LoadAllModules(true))

	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}
	_ = compiled

}
