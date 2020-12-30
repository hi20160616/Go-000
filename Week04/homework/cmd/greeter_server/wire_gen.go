// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/hi20160616/Go-000/Week04/homework/internal/biz"
	"github.com/hi20160616/Go-000/Week04/homework/internal/data"
)

// Injectors from wire.go:

func InitGreeterCase() *biz.GreeterCase {
	greeterRepo := data.NewGreeterRepo()
	greeterCase := biz.NewGreeterCase(greeterRepo)
	return greeterCase
}
