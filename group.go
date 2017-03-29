package httprouter

import (
	"net/http"
)

type RouteGroup struct {
	router *Router
	path   string
}

func newRouteGroup(r *Router, path string) *RouteGroup {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	//Strip traling / (if present) as all added sub paths must start with a /
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return &RouteGroup{router: r, path: path}
}

func (g *RouteGroup) NewGroup(path string) *RouteGroup {
	return newRouteGroup(g.router, g.subPath(path))
}

func (g *RouteGroup) Handle(method, path string, handle Handle) {
	g.router.Handle(method, g.subPath(path), handle)
}

func (g *RouteGroup) Handler(method, path string, handler http.Handler) {
	g.router.Handler(method, g.subPath(path), handler)
}

func (g *RouteGroup) HandlerFunc(method, path string, handler http.HandlerFunc) {
	g.router.HandlerFunc(method, g.subPath(path), handler)
}

func (g *RouteGroup) GET(path string, handle Handle) {
	g.Handle("GET", path, handle)
}

func (g *RouteGroup) HEAD(path string, handle Handle) {
	g.Handle("HEAD", path, handle)
}

func (g *RouteGroup) OPTIONS(path string, handle Handle) {
	g.Handle("OPTIONS", path, handle)
}

func (g *RouteGroup) POST(path string, handle Handle) {
	g.Handle("POST", path, handle)
}

func (g *RouteGroup) PUT(path string, handle Handle) {
	g.Handle("PUT", path, handle)
}

func (g *RouteGroup) PATCH(path string, handle Handle) {
	g.Handle("PATCH", path, handle)
}

func (g *RouteGroup) DELETE(path string, handle Handle) {
	g.Handle("DELETE", path, handle)
}

func (g *RouteGroup) subPath(path string) string {
	if path[0] != '/' {
		panic("path must start with a '/'")
	}
	return g.path + path
}
