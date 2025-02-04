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

package linux

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "linux", asset.ModuleFieldsPri, AssetLinux); err != nil {
		panic(err)
	}
}

// AssetLinux returns asset data.
// This is the base64 encoded zlib format compressed contents of module/linux.
func AssetLinux() string {
	return "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"
}
