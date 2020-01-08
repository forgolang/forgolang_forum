package api

import (
	model2 "forgolang_forum/database/model"
	"forgolang_forum/model"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

type TokenControllerTest struct {
	*Suite
}

func (s TokenControllerTest) SetupSuite() {
	SetupSuite(s.Suite)
}

func (s TokenControllerTest) Test_PostTokenWithValidParams() {
	pass := "123456"
	userModel := model2.NewUser(&pass)
	userModel.Username = "akdilsiz"
	userModel.Email = "akdilsiz@tecpor.com"
	userModel.IsActive = true
	user := new(model2.User)

	err := s.API.App.Database.Insert(user, userModel, "id", "inserted_at")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(userModel.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	userPassphrase := new(model2.UserPassphrase)
	userPassphraseModel := model2.NewUserPassphrase(userModel.ID)
	userPassphraseModel.InsertedAt = time.Now().UTC()
	err = s.API.App.Database.Insert(userPassphrase, userPassphraseModel, "passphrase")
	s.Nil(err)

	tokenRequest := model.TokenRequest{Passphrase: userPassphrase.Passphrase}

	resp := s.JSON(Post, "/api/v1/auth/token", tokenRequest)

	s.Equal(resp.Status, fasthttp.StatusCreated)

	data, _ := resp.Success.Data.(map[string]interface{})
	s.Equal(data["user_id"], float64(userModel.ID))
	s.NotNil(data["jwt"])
	s.Equal(data["role"], "superadmin")

	s.API.App.Logger.LogInfo("Successfully post token with valid params")
}

func (s TokenControllerTest) Test_Shoul_404Error_PostTokenWithValidParamsIfNotExists() {
	pass := "123456"
	userModel := model2.NewUser(&pass)
	userModel.Username = "akdilsiz2"
	userModel.Email = "akdilsiz2@tecpor.com"
	userModel.IsActive = true
	user := new(model2.User)

	err := s.API.App.Database.Insert(user, userModel, "id")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(userModel.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	tokenRequest := model.TokenRequest{Passphrase: "userPassphrase.Passphrase"}

	resp := s.JSON(Post, "/api/v1/auth/token", tokenRequest)

	s.Equal(resp.Status, fasthttp.StatusNotFound)

	s.API.App.Logger.LogInfo("Should be 404 error post token with valid params " +
		"if passphrase does not exists")
}

func (s TokenControllerTest) Test_Should_404Error_PostTokenWithValidParamsIfUserNotActive() {
	pass := "123456"
	userModel := model2.NewUser(&pass)
	userModel.Username = "akdilsiz3"
	userModel.Email = "akdilsiz3@tecpor.com"
	userModel.IsActive = false
	user := new(model2.User)

	err := s.API.App.Database.Insert(user, userModel, "id")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(userModel.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	userPassphrase := new(model2.UserPassphrase)
	userPassphraseModel := model2.NewUserPassphrase(userModel.ID)
	err = s.API.App.Database.Insert(userPassphrase, userPassphraseModel, "passphrase")
	s.Nil(err)

	tokenRequest := model.TokenRequest{Passphrase: userPassphrase.Passphrase}

	resp := s.JSON(Post, "/api/v1/auth/token", tokenRequest)

	s.Equal(resp.Status, fasthttp.StatusNotFound)

	s.API.App.Logger.LogInfo("Should be 404 error post token with valid params " +
		"if user not active")
}

func (s TokenControllerTest) Test_Should_404Error_PostTokenWithValidParamsIfPassphraseExpire() {
	pass := "123456"
	userModel := model2.NewUser(&pass)
	userModel.Username = "akdilsiz4"
	userModel.Email = "akdilsiz4@tecpor.com"
	userModel.IsActive = true
	user := new(model2.User)

	err := s.API.App.Database.Insert(user, userModel, "id")
	s.Nil(err)

	roleAssignment := model2.NewUserRoleAssignment(userModel.ID, 1)
	err = s.API.App.Database.Insert(new(model2.UserRoleAssignment), roleAssignment, "id")
	s.Nil(err)

	userPassphrase := new(model2.UserPassphrase)
	userPassphraseModel := model2.NewUserPassphrase(userModel.ID)
	userPassphraseModel.InsertedAt = time.Now().UTC().AddDate(0, -4, 0)
	err = s.API.App.Database.Insert(userPassphrase, userPassphraseModel, "passphrase")
	s.Nil(err)

	tokenRequest := model.TokenRequest{Passphrase: userPassphraseModel.Passphrase}

	resp := s.JSON(Post, "/api/v1/auth/token", tokenRequest)

	s.Equal(resp.Status, fasthttp.StatusNotFound)
	s.API.App.Logger.LogInfo("Should be 404 error post token with valid params " +
		"if passphrase expire")
}

func (s TokenControllerTest) Test_Should_404Err_PostTokenWithValidParamsIfUserRoleNotExists() {
	pass := "123456"
	userModel := model2.NewUser(&pass)
	userModel.Username = "akdilsiz5"
	userModel.Email = "akdilsiz5@tecpor.com"
	userModel.IsActive = true
	user := new(model2.User)

	err := s.API.App.Database.Insert(user, userModel, "id", "inserted_at")
	s.Nil(err)

	userPassphrase := new(model2.UserPassphrase)
	userPassphraseModel := model2.NewUserPassphrase(userModel.ID)
	userPassphraseModel.InsertedAt = time.Now().UTC()
	err = s.API.App.Database.Insert(userPassphrase, userPassphraseModel, "passphrase")
	s.Nil(err)

	tokenRequest := model.TokenRequest{Passphrase: userPassphrase.Passphrase}

	resp := s.JSON(Post, "/api/v1/auth/token", tokenRequest)

	s.Equal(resp.Status, fasthttp.StatusNotFound)

	s.API.App.Logger.LogInfo("Should be 404 error post token with valid params " +
		"if user role assignment not exists")
}

func Test_TokenController(t *testing.T) {
	s := TokenControllerTest{NewSuite()}
	Run(t, s)
}
