/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
    "strconv"

    "github.com/mainflux/mainflux-lite/routes"
    "github.com/mainflux/mainflux-lite/config"
    "github.com/mainflux/mainflux-lite/db"

    "github.com/iris-contrib/middleware/logger"
    "github.com/kataras/iris"

    "github.com/fatih/color"

)

func main() {

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

    // Parse config
    var cfg config.Config
    cfg.Parse()

    // MongoDb
    db.InitMongo(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)

    color.Cyan(banner)
    color.Cyan("Magic happens on port " + strconv.Itoa(cfg.HttpPort))

    // start the server
    iris.Listen(cfg.HttpHost + ":" + strconv.Itoa(cfg.HttpPort))
}

func registerRoutes() {
    // STATUS
	  iris.Get("/status", routes.GetStatus)

    // DEVICES
	  iris.Post("/devices", routes.CreateDevice)
	  iris.Get("/devices", routes.GetDevices)

    iris.Get("/devices/:id", routes.GetDevice)
    iris.Put("/devices/:id", routes.UpdateDevice)

    iris.Delete("/devices/:id", routes.DeleteDevice)

    // CHANNELS
	  iris.Post("/channels", routes.CreateChannel)
	  iris.Get("/channels", routes.GetChannels)

    iris.Get("/channels/:id", routes.GetChannel)
    iris.Put("/channels/:id", routes.UpdateChannel)

    iris.Delete("/channels/:id", routes.DeleteChannel)
}

var banner = `
_|      _|            _|                _|_|  _|                      
_|_|  _|_|    _|_|_|      _|_|_|      _|      _|  _|    _|  _|    _|  
_|  _|  _|  _|    _|  _|  _|    _|  _|_|_|_|  _|  _|    _|    _|_|    
_|      _|  _|    _|  _|  _|    _|    _|      _|  _|    _|  _|    _|  
_|      _|    _|_|_|  _|  _|    _|    _|      _|    _|_|_|  _|    _|  
                                                                     

                == Industrial IoT System ==
       
                Made with <3 by Mainflux Team
[w] http://mainflux.io
[t] @mainflux

                       ** LITE **

`

