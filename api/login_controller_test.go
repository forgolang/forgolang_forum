package api

import (
	"forgolang_forum/database"
	model2 "forgolang_forum/database/model"
	"forgolang_forum/model"
	"github.com/valyala/fasthttp"
	"testing"
)

type LoginControllerTest struct {
	*Suite
}

func (s LoginControllerTest) SetupSuite() {
	SetupSuite(s.Suite)
}

func (s LoginControllerTest) Test_PostLoginWithValidParams() {
	pass := "123456"
	user := model2.NewUser(&pass)
	user.Username = "akdilsiz-login"
	user.Email = "akdilsiz@tecpor.com"
	userModel := new(model2.User)

	err := s.API.App.Database.Insert(userModel, user, "id")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(user.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	userState := model2.NewUserState(user.ID)
	userState.State = database.Active
	err = s.API.App.Database.Insert(new(model2.UserState), userState, "id")
	s.Nil(err)

	loginRequest := model.LoginRequest{
		ID:       "akdilsiz-login",
		Password: "123456",
	}

	resp := s.JSON(Post, "/api/v1/user/sign_in", loginRequest)

	s.Equal(resp.Status, fasthttp.StatusCreated)
	s.Equal(resp.Success.Data.(map[string]interface{})["user_id"], float64(user.ID))
	s.Equal(len(resp.Success.Data.(map[string]interface{})["passphrase"].(string)), 192)

	s.API.App.Logger.LogInfo("Successfully Post login with valid params")
}

func (s LoginControllerTest) Test_Should_422Error_PostLoginWithInvalidParams() {
	loginRequest := model.LoginRequest{
		ID: "akdilsiz",
	}

	resp := s.JSON(Post, "/api/v1/user/sign_in", loginRequest)

	s.Equal(resp.Status, fasthttp.StatusUnprocessableEntity)

	s.API.App.Logger.LogInfo("Should be 422 error post login with invalid params")
}

func (s LoginControllerTest) Test_Should_404Error_PostLoginWithValidParamsIfUserNotExists() {
	loginRequest := model.LoginRequest{
		ID:       "not_found",
		Password: "123456",
	}

	resp := s.JSON(Post, "/api/v1/user/sign_in", loginRequest)

	s.Equal(resp.Status, fasthttp.StatusNotFound)

	s.API.App.Logger.LogInfo("Should be 404 error post login with valid params" +
		"if user does not exists")
}

func (s LoginControllerTest) Test_Should_401Error_PostLoginWithValidParamsIfPasswordNotMatch() {
	pass := "123456789"
	user := model2.NewUser(&pass)
	user.Username = "akdilsiz2-notmatch"
	user.Email = "akdilsiz2@tecpor.com"
	userModel := new(model2.User)

	err := s.API.App.Database.Insert(userModel, user, "id")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(user.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	userState := model2.NewUserState(user.ID)
	userState.State = database.Active
	err = s.API.App.Database.Insert(new(model2.UserState), userState, "id")
	s.Nil(err)

	loginRequest := model.LoginRequest{
		ID:       "akdilsiz2-notmatch",
		Password: "12345",
	}

	resp := s.JSON(Post, "/api/v1/user/sign_in", loginRequest)

	s.Equal(resp.Status, fasthttp.StatusUnauthorized)

	s.API.App.Logger.LogInfo("Should be 404 error post login with valid params" +
		"if password does not match")
}

func (s LoginControllerTest) TearDownSuite() {
	TearDownSuite(s.Suite)
}

func Test_LoginController(t *testing.T) {
	s := LoginControllerTest{NewSuite()}
	Run(t, s)
}
