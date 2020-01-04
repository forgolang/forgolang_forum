// Copyright 2019 Abdulkadir DILSIZ
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
)

// ResponseSuccess rest api success response structure
type ResponseSuccess struct {
	ResponseInterface `json:"-"`
	Data              interface{} `json:"data"`
	TotalCount        int64       `json:"total_count"`
}

// ToJSON response structure to json string
func (r ResponseSuccess) ToJSON() string {
	body, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(body)
}

// ResponseSuccessOne rest api success response structure
type ResponseSuccessOne struct {
	ResponseInterface `json:"-"`
	Data              interface{} `json:"data"`
}

// ToJSON response structure to json string
func (r ResponseSuccessOne) ToJSON() string {
	body, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(body)
}
