/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package models

type (
	SenML struct {
		Bt  float64                  `json:"bt"`
		Bn  string                   `json:"bn"`
		Bu  string                   `json:"bu"`
		Ver int                      `json:"ver"`
		E   []map[string]interface{} `json:"e"`
	}

	Channel struct {
		Id      string `json:"id"`
		Device  string `json:"device"`
		Created string `json:"created"`
		Updated string `json:"updated"`

		Ts SenML `json:"ts"`

		Msg map[string]interface{} `json:"msg"`

		Metadata  map[string]interface{} `json:"metadata"`
		Mfprivate map[string]interface{} `json:"mfprivate"`
	}
)
