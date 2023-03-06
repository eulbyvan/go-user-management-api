/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 9:48:33 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package delivery

import (
	"fmt"
	"log"

	"github.com/eulbyvan/go-user-management/config"
	"github.com/eulbyvan/go-user-management/controller"
	"github.com/eulbyvan/go-user-management/manager"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (p *AppServer) v1() {
	v1Routes := p.engine.Group("/v1")
	p.userController(v1Routes)
}

func (p *AppServer) userController(rg *gin.RouterGroup) {
	controller.NewUserController(rg, p.usecaseManager.UserUsecase())
}

func (p *AppServer) Run() {
	p.v1()
	err := p.engine.Run(p.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
