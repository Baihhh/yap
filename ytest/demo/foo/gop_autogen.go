package main

import "github.com/goplus/yap"

const _ = true

type foo struct {
	yap.App
}
//line ytest/demo/foo/foo_yap.gox:1
func (this *foo) MainEntry() {
//line ytest/demo/foo/foo_yap.gox:1:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line ytest/demo/foo/foo_yap.gox:2:1
		ctx.Json__1(map[string]string{"id": ctx.Param("id")})
	})
//line ytest/demo/foo/foo_yap.gox:7:1
	this.Run(":8080")
}
func main() {
//line ytest/demo/foo/foo_ytest.gox:3:1
	yap.Gopt_App_Main(new(foo))
}