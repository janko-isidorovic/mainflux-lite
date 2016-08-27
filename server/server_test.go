/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package server

import (
	"fmt"
	"testing"
	"github.com/mainflux/mainflux-lite/config"
	"github.com/mainflux/mainflux-lite/db"
	"github.com/kataras/iris"
)

func TestServer(t *testing.T) {

	// Config
	var cfg config.Config
	cfg.Parse()

	// MongoDb
	db.InitMongo(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)

	go ServeHTTP(cfg)

	// prepare test framework
	if ok := <-iris.Available; !ok {
		t.Fatal("Unexpected error: server cannot start, please report this as bug!!")
	}


	e := iris.Tester(t)
	r := e.Request("GET", "/status").Expect().Status(iris.StatusOK).JSON()
	fmt.Println("%v", r)

}

