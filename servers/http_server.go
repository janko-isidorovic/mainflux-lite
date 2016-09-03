/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package servers

import (
	"strconv"

	"github.com/mainflux/mainflux-lite/routes"
	"github.com/mainflux/mainflux-lite/config"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)


func HttpServer(cfg config.Config) {
	// Iris config
	iris.Config.DisableBanner = true

	// set the global middlewares
	iris.Use(logger.New(iris.Logger))

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// register public API
	registerRoutes()

	// start the server
	iris.Listen(cfg.HttpHost + ":" + strconv.Itoa(cfg.HttpPort))
}

func registerRoutes() {
	// STATUS
	iris.Get("/status", routes.GetStatus)

	// DEVICES
	iris.Post("/devices", routes.CreateDevice)
	iris.Get("/devices", routes.GetDevices)

	iris.Get("/devices/:device_id", routes.GetDevice)
	iris.Put("/devices/:device_id", routes.UpdateDevice)

	iris.Delete("/devices/:device_id", routes.DeleteDevice)

	// CHANNELS
	iris.Post("/devices/:device_id/channels", routes.CreateChannel)
	iris.Get("/devices/:device_id/channels", routes.GetChannels)

	iris.Get("/devices/:device_id/channels/:channel_id", routes.GetChannel)
	iris.Put("/devices/:device_id/channels/:channel_id", routes.UpdateChannel)

	iris.Delete("/devices/:device_id/channels/:channel_id", routes.DeleteChannel)
}
