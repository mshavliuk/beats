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

package tls

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
		panic(err)
	}
}

// AssetTls returns asset data.
// This is the base64 encoded zlib format compressed contents of protos/tls.
func AssetTls() string {
	return "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"
}
