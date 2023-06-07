package main

import (
	"github.com/kataras/iris/v12/context"
)

func Hello(ctx *context.Context) {
	ctx.WriteString("hello iris")
}
