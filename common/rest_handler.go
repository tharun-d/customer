package common

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type ctxKey struct{}

var routes = []route{}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Route(method, pattern string, handler http.HandlerFunc) {
	routes = append(routes, route{method, regexp.MustCompile("^" + pattern + "$"), handler})
}

func Param(r *http.Request, index int) string {
	params := r.Context().Value(ctxKey{}).([]string)
	return params[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
