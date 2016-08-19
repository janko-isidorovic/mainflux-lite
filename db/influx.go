/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package db

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"strconv"
)

type InfluxConn struct {
	C  client.Client
	Bp client.BatchPoints
}

var IfxConn InfluxConn

func InitInflux(host string, port int, db string) error {
	// Make client
	icc, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + host + ":" + strconv.Itoa(port),
		//Username: username,
		//Password: password,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	IfxConn.C = icc

	// Create a new point batch
	icbp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	IfxConn.Bp = icbp

	return err
}
