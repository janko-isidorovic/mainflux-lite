/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package routes

import (
	"log"
	"time"
	"math"

	"github.com/mainflux/mainflux-lite/db"
	"github.com/mainflux/mainflux-lite/models"

	"github.com/influxdata/influxdb/client/v2"
)

/** ===InfluxDB Manipualtion === */
 /**
 * Insert time-series
 */
func insertTs(id string, ts models.SenML) int {
	rc := 0

	// Insert in SenML in Influx
	// SenML can contain several datapoints
	// and can target different tags

	// Loop here for each attribute
	for k, v := range ts.E {
		tags := map[string]string{
			"attribute": ts.E[k]["n"].(string),
		}

		// Examine if "v" exists, then "sv", then "bv"
		var field map[string]interface{}
		if vv, ok := v["v"]; ok {
			field["value"] = vv
		} else if vsv, ok := v["sv"]; ok {
			field["value"] = vsv
		} else if vbv, ok := v["bv"]; ok {
			field["value"] = vbv
		}

		/**
		 * Handle time
		 *
		 * If either the Base Time or Time value is missing, the missing
		 * attribute is considered to have a value of zero.  The Base Time and
		 * Time values are added together to get the time of measurement.  A
		 * time of zero indicates that the sensor does not know the absolute
		 * time and the measurement was made roughly "now".  A negative value is
		 * used to indicate seconds in the past from roughly "now".  A positive
		 * value is used to indicate the number of seconds, excluding leap
		 * seconds, since the start of the year 1970 in UTC.
		 */
		// Set time base
		var tb float64
		if bt := ts.Bt; bt != 0.0 {
			// If bt is sent and is different than zero
			// N.B. if bt was not sent, `ts.Bt` will still be zero, as this is init value
			tb = bt
		} else {
			// If not that means that sensor does not have RTC
			// and want us to use our NTP - "roughly now"
			//tb = float64(time.Now())
		}

		// Set relative time
		var tr int64
		if vt, ok := v["t"]; ok {
			// If there is relative time, use it
			tr = vt.(int64)
		} else {
			// Otherwise it is considered as zero
			tr = 0
		}

		// Total time
		tt := tb + float64(tr)
		// Break into int and fractional nb
		ts, tsf := math.Modf(tt)
		// Find nanoseconds number from fractional part
		tns := tsf * 1000 * 1000

		// Get time in Unix format, based on s and ns
		t := time.Unix(int64(ts), int64(tns))
		pt, err := client.NewPoint(id, tags, field, t)

		if err != nil {
			log.Fatalln("Error: ", err)
		}

		db.IfxConn.Bp.AddPoint(pt)
	}

	// Write the batch
	db.IfxConn.C.Write(db.IfxConn.Bp)

	return rc
}

// QueryDB convenience function to query the database
func queryInfluxDb(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "mainflux",
	}
	if response, err := db.IfxConn.C.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
