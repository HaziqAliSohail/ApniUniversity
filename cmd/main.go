package main

import (
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "ApniUniversity"
	"ApniUniversity/config"
	"ApniUniversity/gen/restapi"
	"ApniUniversity/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"unix", "http"}
	defer func(server *restapi.Server) {
		err := server.Shutdown()
		if err != nil {

		}
	}(server)

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
