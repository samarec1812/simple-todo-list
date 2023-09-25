package main

import (
	"github.com/samarec1812/simple-todo-list/internal/app"
	"github.com/samarec1812/simple-todo-list/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
