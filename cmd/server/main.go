package main

import (
	_ "github.com/impzero/roosh/docs"
	"github.com/impzero/roosh/roosh"
	"github.com/impzero/roosh/web"
)

//	@title			Roosh API
//	@version		1.0
//	@description	Roosh's API
//	@BasePath		/

func main() {
	memory := &roosh.Stack{}
	calc := roosh.StackCalculator{Memory: memory}
	w := web.New(web.Config{
		Addr:       ":8080",
		Stack:      memory,
		Calculator: &calc,
	})

	w.Start()
}
