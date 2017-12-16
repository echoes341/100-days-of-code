package main

import (
	"github.com/julienschmidt/httprouter"
)

func initRoutes(m *httprouter.Router) {
	m.GET("/", home)
}
