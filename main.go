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
	"github.com/mainflux/mainflux-lite/servers"
	"github.com/mainflux/mainflux-lite/clients"
	"github.com/fatih/color"
	"runtime"
)

type MainfluxLite struct {
}

func main() {

	// Parse config
	var cfg config.Config
	cfg.Parse()

	// MongoDb
	db.InitMongo(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)

	// MQTT 
	mqc := new(clients.MqttConn)
	//Sub to everything comming on all channels of all devices
	mqc.MqttSub()

	// Serve HTTP
	go servers.HttpServer(cfg)

	// Print banner
	color.Cyan(banner)
	color.Cyan("Magic happens on port " + strconv.Itoa(cfg.HttpPort))

	/** Keep main() runnig */
	runtime.Goexit()
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
