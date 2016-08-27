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

	e := iris.Tester(t)
	r := e.Request("GET", "http://localhost:7070/status").Expect().JSON().Array()
	fmt.Println(r)

}

/**
func TestServer(t *testing.T) {
	mfl := new(MainfluxLite)
	expectedBody := `{"status":"OK"}`

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:7070/status", expectedBody), nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	// Config
	cfg := config.Config{
		HttpHost:"localhost",
		HttpPort:7070,

		// Mongo
		MongoHost:"localhost",
		MongoPort:27017,
		MongoDatabase:"mainflux_test",
	}

	//mfl.ServeHTTP(recorder, req)
	mfl.ServeHTTP(cfg)

	switch recorder.Body.String() {
	case expectedBody:
		// body is equal so no need to do anything
	default:
		t.Errorf("Body (%s) did not match expectation (%s).",
			recorder.Body.String(),
			expectedBody)
	}
}
*/
