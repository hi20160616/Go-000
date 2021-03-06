//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hi20160616/Go-000/Week04/homework/internal/biz"
	"github.com/hi20160616/Go-000/Week04/homework/internal/data"
)

func InitGreeterCase() *biz.GreeterCase {
	wire.Build(biz.NewGreeterCase, data.NewGreeterRepo)
	return &biz.GreeterCase{}
}
