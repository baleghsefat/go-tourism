package controllers

import (
	"net/http"
	"strconv"
	UserModel "tourism/app/Models/User"

	"github.com/labstack/echo/v4"
)

var NewUser UserModel.User

func Index(context echo.Context) error {
	users := UserModel.GetAllUsers()

	return context.JSON(http.StatusOK, users)
}

func Show(context echo.Context) error {
	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	userId := int(number)
	user := UserModel.GetUserById(userId)

	return context.JSON(http.StatusOK, user)
}

func Store(context echo.Context) (err error) {
	var newUser UserModel.User
	// if err = context.Validate(newUser); err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }
	// newUser.FirstName = context.QueryParam("first_name")
	// newUser.Mobile = context.QueryParam("mobile")
	context.Bind(&newUser)

	UserModel.CreateUser(&newUser)

	return context.JSON(http.StatusCreated, newUser)
}

func Update(context echo.Context) error {
	var user UserModel.User

	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	userId := int(number)

	context.Bind(&user)

	UserModel.UpdateUser(&user, userId)

	return context.JSON(http.StatusOK, user)
}

func Destroy(context echo.Context) error {
	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	userId := int(number) //TODO move to a function

	UserModel.DeleteUser(userId)

	return context.NoContent(http.StatusNoContent)
}
