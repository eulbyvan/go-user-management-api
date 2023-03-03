/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 11:15:05 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package controller

import (
	"net/http"
	"strconv"

	"github.com/eulbyvan/go-user-management/model"
	"github.com/eulbyvan/go-user-management/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func (c *UserController) GetAll(ctx *gin.Context) {
	res := c.usecase.GetAll()

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.GetOne(id)
	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) Add(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Add(&user)
	ctx.JSON(http.StatusCreated, res)
}

func (c *UserController) Edit(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&user)
	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.Remove(id)
	ctx.JSON(http.StatusOK, res)
}

func NewUserController(u usecase.UserUsecase) *UserController {
	controller := UserController {
		usecase: u,
	}

	return &controller
}