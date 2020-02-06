// Copyright 2019 StreetByters Community
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
	"forgolang_forum/database"
	"gopkg.in/guregu/null.v3/zero"
	"time"
)

// Tag special classifications
type Tag struct {
	database.DBInterface `json:"-"`
	ID                   int64     `db:"id" json:"id"`
	Name                 string    `db:"name" json:"name" validate:"required,gte=2,lte=32"`
	Count                zero.Int  `json:"count,omitempty"`
	InsertedAt           time.Time `db:"inserted_at" json:"inserted_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

// NewTag generate tag structure
func NewTag() *Tag {
	return &Tag{}
}

// TableName tag database
func (m Tag) TableName() string {
	return "tags"
}

// ToJSON tag structure to json string
func (m Tag) ToJSON() string {
	return database.ToJSON(m)
}

// Timestamps generate timestamp fields
func (m Tag) Timestamps() bool {
	return true
}
