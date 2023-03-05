/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 10:11:22 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package app_error

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorMessage string
	ErrorType    int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("type: %d, err: %s", e.ErrorType, e.ErrorMessage)
}

func InvalidError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "invalid input",
			ErrorType: http.StatusBadRequest,
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
			ErrorType: http.StatusBadRequest,
		}
	}
}

func UnauthorizedError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "unauthorized user",
			ErrorType: http.StatusUnauthorized,
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
			ErrorType: http.StatusUnauthorized,
		}
	}
}

func DataNotFoundError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "no data found",
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
		}
	}
}

func InternalServerError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "something went wrong",
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
		}
	}
}