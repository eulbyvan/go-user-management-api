/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 9:51:39 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package entity

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required"`
}