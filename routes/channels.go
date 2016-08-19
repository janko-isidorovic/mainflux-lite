/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/mainflux/mainflux-lite/db"
	"github.com/mainflux/mainflux-lite/models"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
)

/** == Functions == */

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
		if vv, okv := v["v"]; okv {
			field["value"] = vv
		} else if vsv, oksv := v["sv"]; oksv {
			field["value"] = vsv
		} else if vbv, okbv := v["bv"]; okbv {
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
			tb = float64(time.Now().Unix())
		}

		// Set relative time
		var tr int64
		if vt, okvt := v["t"]; okvt {
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

func insertMsg(id string, msg map[string]interface{}) int {
	rc := 0

	// Insert Msg in Influx
	// Check if we can insert one message blob as a single datapoint
	tags := map[string]string{
		"attribute": "msg",
	}
	client.NewPoint(id, tags, msg, time.Now())
	return rc
}

// QueryDB convenience function to query the database
func queryInfluxDb(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "Mainflux",
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

/**
 * CreateChannel ()
 */
func CreateChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)
	if validateJsonSchema(body) != true {
		println("Invalid schema")
	}

	// Init new Mongo session
	// and get the "devices" collection
	// from this new session
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Turn map into a JSON to put it in the Device struct later
	j, err := json.Marshal(&body)
	if err != nil {
		fmt.Println(err)
	}

	// Set up defaults and pick up new values from user-provided JSON
	channel := models.Channel{Id: "Some Id"}
	json.Unmarshal(j, &channel)

	// Creating UUID Version 4
	uuid := uuid.NewV4()
	fmt.Println(uuid.String())

	channel.Id = uuid.String()

	fmt.Println(channel)

	// Insert Device
	erri := Db.C("channels").Insert(channel)
	if erri != nil {
		println("CANNOT INSERT")
		panic(erri)
	}

	ctx.Write("Created Device req.channelId")
}

/**
 * GetChannels()
 */
func GetChannels(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	results := []models.Channel{}
	err := Db.C("channels").Find(nil).All(&results)
	if err != nil {
		println("ERROR!!!")
		log.Print(err)
	}

	res, err := json.Marshal(results)
	if err != nil {
		fmt.Println("error:", err)
	}

	ctx.Write(string(res))
}

/**
 * GetChannel()
 */
func GetChannel(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("id")

	result := models.Channel{}
	err := Db.C("channels").Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Print(err)
	}

	res, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(res)

	ctx.Write(string(res))
}

/**
 * UpdateChannel()
 */
func UpdateChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)
	// Validate JSON schema user provided
	if validateJsonSchema(body) != true {
		println("Invalid schema")
	}

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("id")

	// Check if someone is trying to change "id" key
	// and protect us from this
	if _, ok := body["id"]; ok {
		println("Error: can not change device ID")
	}

	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": body}
	err := Db.C("channels").Update(colQuerier, change)
	if err != nil {
		log.Print(err)
	}

	ctx.Write(`{"status":"updated"}`)
}

/**
 * DeleteChannel()
 */
func DeleteChannel(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("id")

	err := Db.C("channels").Remove(bson.M{"id": id})
	if err != nil {
		log.Print(err)
	}

	ctx.Write(`{"status":"deleted"}`)
}

/**
 * SendChannel()
 */
func SendChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)

	id := ctx.Param("id")

	if m, ok := body["msg"]; ok {
		insertMsg(id, m.(map[string]interface{}))
	}

	if t, ok := body["ts"]; ok {
		insertTs(id, t.(models.SenML))
	}

	ctx.Write(`{"status":"inserted"}`)
}
