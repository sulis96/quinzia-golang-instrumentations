package router

import (
	"fmt"
	"net/http"

	"github.com/sulis96/quinzia-golang-instrumentations/config"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/controller"
)

type (
	router struct {
		controller controller.IController
	}

	IRouter interface {
		RunHTTP(c *config.AppConfig)
	}
)

func NewRouter(c controller.IController) IRouter {
	return &router{
		controller: c,
	}
}

func (r *router) RunHTTP(c *config.AppConfig) {
	r.Router()
	fmt.Println("starting server at http://localhost:" + c.HTTPServerPort)
	http.ListenAndServe(":"+c.HTTPServerPort, nil)
}

func (r *router) Router() {
	http.HandleFunc("/health", r.controller.Health)
	// api member
	http.HandleFunc("/member/create", r.controller.CreateMember)
	http.HandleFunc("/member/list", r.controller.ReadMember)
}
