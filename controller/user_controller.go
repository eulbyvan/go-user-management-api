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
	 router  *gin.RouterGroup
	 usecase usecase.UserUsecase
 }
 
 func NewUserController(r *gin.RouterGroup, u usecase.UserUsecase) *UserController {
	 controller := UserController{
		 router:  r,
		 usecase: u,
	 }
 
	 // define routes
	 r.GET("/users", controller.GetAll)
	 r.GET("/users/:id", controller.GetOne)
	 r.POST("/users", controller.Add)
	 r.PUT("/users/:id", controller.Edit)
	 r.DELETE("/users/:id", controller.Remove)
 
	 return &controller
 }
 
 func (c *UserController) GetAll(ctx *gin.Context) {
	 res, err := c.usecase.GetAll()
	 if err != nil {
		 c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", "Successfully retrieved all user data", res)
 }
 
 func (c *UserController) GetOne(ctx *gin.Context) {
	 id, err := strconv.Atoi(ctx.Param("id"))
	 if err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		 return
	 }
 
	 res, err := c.usecase.GetOne(id)
	 if err != nil {
		 c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("user with id %d not found", id)))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully retrieved user with ID %d", id), res)
 }
 
func (c *UserController) Add(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}

	res, err := c.usecase.Add(&user)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create user"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully created new user", res)
}

 
 func (c *UserController) Edit(ctx *gin.Context) {
	 var user entity.User
	 if err := ctx.BindJSON(&user); err != nil {
		 c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid request body"))
		 return
	 }
 
	 res, err := c.usecase.Edit(&user)
	 if err != nil {
		 c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("user with id %d not found", user.ID)))
		 return
	 }
 
	 c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully updated user with ID %d", user.ID), res)
 }
 
 func (c *UserController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	err = c.usecase.Remove(id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("user with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully removed user with ID %d", id), nil)
}