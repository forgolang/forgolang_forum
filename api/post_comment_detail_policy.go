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

package api

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
)

// PostCommentDetailPolicy post comment detail authorization
type PostCommentDetailPolicy struct {
	Policy
	*API
}

// Delete post comment authorization
func (p PostCommentDetailPolicy) Create(next phi.HandlerFunc) phi.HandlerFunc {
	pcP := PostCommentPolicy{API: p.API}
	return p.API.Authorization.Apply(next, "PostCommentDetailController", "Create",
		func(ctx *fasthttp.RequestCtx) bool {
			if comment := pcP.GetComment(ctx); comment != nil && comment.UserID == p.GetAuthContext(ctx).ID {
				return true
			}
			return false
		})
}
