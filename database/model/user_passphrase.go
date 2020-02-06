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
	"forgolang_forum/model"
	"github.com/streetbyters/agente/utils"
	"time"
)

// UserPassphrase authentication access token struct
type UserPassphrase struct {
	database.DBInterface `json:"-"`
	ID                   int64     `db:"id" json:"id"`
	UserID               int64     `db:"user_id" json:"user_id" foreign:"fk_user_passphrases_user_id"`
	Passphrase           string    `db:"passphrase" json:"passphrase" unique:"user_passphrases_passphrase_unique_index"`
	InsertedAt           time.Time `db:"inserted_at" json:"inserted_at"`
}

// NewUserPassphrase generate authentication access token
func NewUserPassphrase(userID int64) *UserPassphrase {
	return &UserPassphrase{
		UserID:     userID,
		Passphrase: utils.Passkey(),
	}
}

// TableName user_passphrase database table name
func (d UserPassphrase) TableName() string {
	return "user_passphrases"
}

// ToJSON UserPassphrase database model to json string
func (d UserPassphrase) ToJSON() string {
	return database.ToJSON(d)
}

// PassphraseQuery generate user_passphrase query string for database type
func (d UserPassphrase) PassphraseQuery(db *database.Database) string {
	passphraseInvalidation := NewUserPassphraseInvalidation()
	var query string
	switch db.Type {
	case model.Postgres:
		query = "SELECT p.* FROM " + d.TableName() + " AS p " +
			"LEFT OUTER JOIN " + passphraseInvalidation.TableName() + " AS pi ON p.id = pi.passphrase_id " +
			"WHERE pi.passphrase_id IS NULL AND p.passphrase = $1 AND " +
			"p.inserted_at >= (CURRENT_TIMESTAMP - interval '3 month')"
		break
	}

	return query
}
