// Copyright 2019 Street Byters Community
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
	"forgolang_forum/model"
	"github.com/valyala/fasthttp"
)

// HomeController base controller
type HomeController struct {
	Controller
	*API
}

// Index api base route
func (c HomeController) Index(ctx *fasthttp.RequestCtx) {
	c.JSONResponse(ctx, model.ResponseSuccess{
		Data: "Forgolang",
	}, fasthttp.StatusOK)
}
