/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */
package routes

import (
	"fmt"

	"github.com/kataras/iris"
)

/** == Functions == */

/**
 * getStatus()
 */
func GetStatus(ctx *iris.Context) {
	fmt.Println("Status OK")
	ctx.Write(`{"status":"OK"}`)
}
