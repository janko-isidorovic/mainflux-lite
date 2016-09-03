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
	"github.com/krylovsk/gosenml"
)

type ChannelWriteStatus struct {
	Nb int
	Str string
}

/** == Functions == */

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


	c := models.Channel{}
	json.Unmarshal(ctx.RequestCtx.Request.Body(), &c)

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
	err := Db.C("channels").Insert(c)
	if err != nil {
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
 * WriteChannel()
 * Generic function that updates the channel value.
 * Can be called via various protocols
 */
func WriteChannel(id string, bodyBytes []byte) ChannelWriteStatus {
	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		fmt.Println("Error unmarshaling body")
	}

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Check if someone is trying to change "id" key
	// and protect us from this
	s := ChannelWriteStatus{}
	if _, ok := body["id"]; ok {
		s.Nb = iris.StatusBadRequest
		s.Str = "Invalid request: 'id' is read-only"
		return s
	}
	if _, ok := body["device"]; ok {
		println("Error: can not change device")
		s.Nb = iris.StatusBadRequest
		s.Str = "Invalid request: 'device' is read-only"
		return s
	}
	if _, ok := body["created"]; ok {
		println("Error: can not change device")
		s.Nb = iris.StatusBadRequest
		s.Str = "Invalid request: 'created' is read-only"
		return s
	}

	senmlDecoder := gosenml.NewJSONDecoder()
	m, _ := senmlDecoder.DecodeMessage(bodyBytes)
	for _, e := range m.Entries {
		// BaseName
		e.Name = m.BaseName + e.Name

		// BaseTime
		e.Time = m.BaseTime + e.Time
		if e.Time <= 0 {
			e.Time += time.Now().Unix()
		}

		// BaseUnits
		if e.Units == "" {
			e.Units = m.BaseUnits
		}

		/** Insert entry in DB */
		colQuerier := bson.M{"id": id}
		change := bson.M{"$push": bson.M{"values": e}}
		err := Db.C("channels").Update(colQuerier, change)
		if err != nil {
			log.Print(err)
			s.Nb = iris.StatusNotFound
			s.Str = "Not inserted"
			return s
		}
	}

	// Timestamp
	t := time.Now().UTC().Format(time.RFC3339)
	body["updated"] = t

	/** Then update channel timestamp */
	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": bson.M{"updated": body["updated"]}}
	err := Db.C("channels").Update(colQuerier, change)
	if err != nil {
		log.Print(err)
		s.Nb = iris.StatusNotFound
		s.Str = "Not updated"
		return s
	}

	s.Nb = iris.StatusOK
	s.Str = "Updated"
	return s
}


/**
 * UpdateChannel()
 */
func UpdateChannel(ctx *iris.Context) {
	var body map[string]interface{}
	ctx.ReadJSON(&body)
	// Validate JSON schema user provided
	/*
	if validateJsonSchema("channel", body) != true {
		println("Invalid schema")
		ctx.JSON(iris.StatusBadRequest, iris.Map{"response": "invalid json schema in request"})
		return
	}
	*/

	id := ctx.Param("channel_id")

	status := WriteChannel(id, ctx.RequestCtx.Request.Body())
	ctx.JSON(status.Nb, iris.Map{"response": status.Str})
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


