// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "./vendor/github.com/elastic/beats/metricbeat/_meta/fields.common.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJyUkz2O2zAQhXud4mH71QFUBAgWAVKl2gvMkmNzYJKjkCMrun0gxkpiYy3D6ijqvW9+nl5x4mWA05Q0d4CJRR7wtp09V1dkNNE84EsHAG+ajSTXiwgH4egr6EwS6SMyJINiBJ85G2wZufYdLp8NXfN4RabEAxJbEVfZ+qR+itwuP6Wuz3vgpoMeYIHxRwMLZDhy5kLGvt00dn+PtZ6fJG3a52FBq+3Dvmu1Kxi5IJlxKJowB3HhpoaZ1uHHyM7Y93gPUv+atTEj0YKshg/GWLiui5gD5+bjyejaAlEdxbjc7aHY1sK6zgFR8/HyovDPSQr7AVamB1P91hJRdMoeVmSESWpxSeKKVnaafd3dWh3J8VUpJ15mLX4f/GOTriP2S6Yk7p/zLXI13uuOf1Ea4/+11cdp8uqmtP0QPb7GmZaKlijFi1f30ne/AwAA//9Prh2/"
}
