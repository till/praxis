package server

import (
	"github.com/convox/praxis/api"
	"github.com/convox/praxis/server/controllers"
)

type Server struct {
	*api.Server
}

func New() *Server {
	api := api.New("rack", "convox.rack")

	api.Route("app.create", "POST", "/apps", controllers.AppCreate)
	api.Route("app.delete", "DELETE", "/apps/{app}", controllers.AppDelete)

	api.Route("build.create", "POST", "/apps/{app}/builds", controllers.BuildCreate)

	api.Route("release.get", "GET", "/apps/{app}/releases/{id}", controllers.ReleaseGet)

	api.Route("object.store", "POST", "/apps/{app}/objects/{key:.*}", controllers.ObjectStore)

	return &Server{Server: api}
}