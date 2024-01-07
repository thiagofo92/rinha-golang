package routerbuilder

import (
	"thiagofo92/api-test/controller"

	"github.com/go-chi/chi/v5"
)

type RouterBuild struct {
	router *chi.Mux
}

func NewRouterBuild(r *chi.Mux) *RouterBuild {
	return &RouterBuild{
		router: r,
	}
}

func (r RouterBuild) userRouter() {
	controller := controller.NewUserController()
	r.router.Get("/user", controller.Find)
}

func (r RouterBuild) Build() *chi.Mux {
	r.userRouter()
	return r.router
}
