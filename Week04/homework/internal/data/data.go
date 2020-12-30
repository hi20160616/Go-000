package data

import (
	"log"

	"github.com/hi20160616/Go-000/Week04/homework/internal/biz"
)

const greeterID = 100 // imagin that is got from database

var _ biz.GreeterRepo = new(greeterRepo)

type greeterRepo struct{}

func (g *greeterRepo) SayHi(*biz.Greeter) int32 {
	log.Printf("Hi there is data package!")
	return greeterID
}

func NewGreeterRepo() biz.GreeterRepo {
	return &greeterRepo{}
}
