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
	"time"

	"github.com/mainflux/mainflux-lite/db"
	"github.com/mainflux/mainflux-lite/models"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
)

/** == Functions == */
func formatTs(s map[string]interface{}) models.SenML {

	r := models.SenML{}

	// E
	if _, ok := s["E"]; !ok {
		//ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid request: 'id' is read-only"})
		//return

		// XXX How do we do error here?
		println("E does not exist")
	}

	// Loop here for each attribute and format
	e := s["e"].([]map[string]interface{})
	for k, v := range e {
		// t
		if _, ok := s["bt"]; ok {
			// XXX - must parse time in good format!
			if _, ok := v["t"]; ok {
				r.E[k]["t"] = s["bt"].(int64) + v["t"].(int64)
			} else {
				r.E[k]["t"] = s["bt"].(int64)
			}
		} else {
			r.E[k]["t"] = time.Now().UTC().Format(time.RFC3339)
		}

		// n
		if _, ok := s["n"]; ok {
			if _, ok := v["n"]; ok {
				r.E[k]["n"] = s["bn"].(string) + v["n"].(string)
			} else {
				r.E[k]["n"] = s["bn"].(string)
			}
		}

		// u
		if _, ok := s["bu"]; ok {
			if _, ok := v["u"]; ok {
				r.E[k]["u"] = s["bu"].(string) + v["u"].(string)
			} else {
				r.E[k]["u"] = s["bu"].(string)
			}
		}

		// v
		if _, ok := v["v"]; ok {
			r.E[k]["v"] = v["v"].(float64)
		}

		// sv
		if _, ok := v["sv"]; ok {
			r.E[k]["sv"] = v["sv"].(string)
		}

		// bv
		if _, ok := v["bv"]; ok {
			r.E[k]["bv"] = v["bv"].(float64)
		}
	}

	return r
}

/**
 * CreateChannel ()
 */
func CreateChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)
	/*
	if validateJsonSchema("channel", body) != true {
		println("Invalid schema")
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid json schema in request"})
		return
	}
	*/

	// Init new Mongo session
	// and get the "channels" collection
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
	c := models.Channel{Id: "Some Id"}
	json.Unmarshal(j, &c)

	// Creating UUID Version 4
	uuid := uuid.NewV4()
	fmt.Println(uuid.String())

	c.Id = uuid.String()

	// Insert reference to DeviceId
	did := ctx.Param("device_id")
	c.Device = did

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	c.Created, c.Updated = t, t

	// Insert Channel
	erri := Db.C("channels").Insert(c)
	if erri != nil {
		ctx.JSON(iris.StatusInternalServerError, iris.Map{"response": "cannot create device"})
		return
	}

	ctx.JSON(iris.StatusCreated, iris.Map{"response": "created", "id": c.Id})
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
		log.Print(err)
	}

	ctx.JSON(iris.StatusOK, &results)
}

/**
 * GetChannel()
 */
func GetChannel(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("channel_id")

	result := models.Channel{}
	err := Db.C("channels").Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusNotFound, iris.Map{"response": "not found", "id": id})
		return
	}

	ctx.JSON(iris.StatusOK, &result)
}

/**
 * UpdateChannel()
 */
func UpdateChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)
	// Validate JSON schema user provided
	if validateJsonSchema("channel", body) != true {
		println("Invalid schema")
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid json schema in request"})
		return
	}

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("channel_id")

	// Check if someone is trying to change "id" key
	// and protect us from this
	if _, ok := body["id"]; ok {
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid request: 'id' is read-only"})
		return
	}
	if _, ok := body["device"]; ok {
		println("Error: can not change device")
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid request: 'device' is read-only"})
		return
	}
	if _, ok := body["created"]; ok {
		println("Error: can not change device")
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid request: 'created' is read-only"})
		return
	}

	if _, ok := body["device"]; ok {
		body["ts"] = formatTs(body["ts"].(map[string]interface{}))
	}

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	body["updated"] = t

	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": body}
	err := Db.C("channels").Update(colQuerier, change)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusNotFound, iris.Map{"response": "not updated", "id": id})
		return
	}

	ctx.JSON(iris.StatusOK, iris.Map{"response": "updated", "id": id})
}

/**
 * DeleteChannel()
 */
func DeleteChannel(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	id := ctx.Param("channel_id")

	err := Db.C("channels").Remove(bson.M{"id": id})
		if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusNotFound, iris.Map{"response": "not deleted", "id": id})
		return
	}

	ctx.JSON(iris.StatusOK, iris.Map{"response": "deleted", "id": id})
}


