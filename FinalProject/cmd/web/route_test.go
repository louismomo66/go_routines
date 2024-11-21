package main

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/go-chi/chi"
// )

// var routes = []string{
// 	"/",
// 	"/login",
// 	"/logout",
// 	"/register",
// 	"/activate",
// 	"/members/plans",
// 	"/members/subscribe",
// }

// func Test_Routes(t *testing.T) {
// 	testRoutes := testApp.routes()

// 	chiRoutes, ok := testRoutes.(*chi.Mux)
// 	if !ok {
// 		// If direct type assertion fails, try using chi.NewRouter to compare routes
// 		// router := chi.NewRouter()
// 		routeWalker, ok := testRoutes.(interface {
// 			Routes() []chi.Route
// 		})

// 		if !ok {
// 			t.Fatalf("could not extract routes from handler")
// 		}

// 		routes := routeWalker.Routes()

// 		for _, route := range routes {
// 			t.Logf("Found route: %s", route.Pattern)
// 		}
// 	}

// 	for _, route := range routes {
// 		routeExists(t, chiRoutes, route)
// 	}
// }

// func routeExists(t *testing.T, routes *chi.Mux, route string) {
// 	found := false

// 	chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
// 		if route == foundRoute {
// 			found = true
// 		}
// 		return nil
// 	})

// 	if !found {
// 		t.Errorf("did not find %s in registered routes", route)
// 	}
// }
