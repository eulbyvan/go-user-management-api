/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 11:15:05 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eulbyvan/go-user-management/model/app_error"
	"github.com/eulbyvan/go-user-management/model/entity"
	"github.com/eulbyvan/go-user-management/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
	router *gin.RouterGroup
	usecase usecase.UserUsecase
}

func (c *UserController) GetAll(ctx *gin.Context) {
	res := c.usecase.GetAll()
	c.Success(ctx, 0, "", "Successfully retrieved all user data", res)
}

func (c *UserController) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError(""))
		return
	}

	res := c.usecase.GetOne(id)
	if res == nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("No user with id: %d", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully retrieved user with id: %d", id), res)
}

func (c *UserController) Add(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError(""))
		return
	}

	res := c.usecase.Add(&user)
	c.Success(ctx, http.StatusCreated, "", "Succesfully created new user", res)
}

func (c *UserController) Edit(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError(""))
		return
	}

	res := c.usecase.Edit(&user)
	if res == nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("No user with id: %d", user.ID)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Succesfully updated user with id: %d", user.ID), res)
}

func (c *UserController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError(""))
		return
	}

	res := c.usecase.Remove(id)
	if res == nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("No user with id: %d", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Succesfully deleted user with id: %d", id), res)
}

func NewUserController(r *gin.RouterGroup, u usecase.UserUsecase) *UserController {
	controller := UserController {
		router: r,
		usecase: u,
	}

	return &controller
}