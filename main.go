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
	"github.com/mainflux/mainflux-lite/config"
	"github.com/mainflux/mainflux-lite/db"
	"github.com/mainflux/mainflux-lite/server"
	"github.com/fatih/color"
)

type MainfluxLite struct {
}

func main() {

	// Parse config
	var cfg config.Config
	cfg.Parse()

	// MongoDb
	db.InitMongo(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)

	// Print banner
	color.Cyan(banner)
	color.Cyan("Magic happens on port " + strconv.Itoa(cfg.HttpPort))

	// Serve forever
	server.ServeHTTP(cfg)
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
