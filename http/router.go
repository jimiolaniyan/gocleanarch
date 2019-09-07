package http

import (
	"strings"
)

type Router struct {
	routes map[string]Controller
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]Controller)}
}

func (r *Router) Route(request *ParsedRequest) string {
	println(request.Path)
	parts := strings.Split(request.Path, "/")
	controllerKey := parts[1]
	controller := r.routes[controllerKey]
	if controller == nil {
		return ""
	}
	return controller.Handle(request)
}

func (r *Router) AddPath(path string, controller Controller) {
	r.routes[path] = controller
}
