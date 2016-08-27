/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"fmt"
	"testing"
	"github.com/mainflux/mainflux-lite/config"

	"github.com/kataras/iris"
)

func TestServer(t *testing.T) {

	mfl := new(MainfluxLite)

	// Config
	cfg := config.Config{
		HttpHost:"localhost",
		HttpPort:7070,

		// Mongo
		MongoHost:"localhost",
		MongoPort:27017,
		MongoDatabase:"mainflux_test",
	}

	go mfl.ServeHTTP(cfg)

	// prepare test framework
	if ok := <-iris.Available; !ok {
		t.Fatal("Unexpected error: server cannot start, please report this as bug!!")
	}


	e := iris.Tester(t)
	r := e.Request("GET", "/status").Expect().Status(iris.StatusOK).JSON()
	fmt.Println("%v", r)

}

