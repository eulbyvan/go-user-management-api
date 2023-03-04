/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 10:11:22 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package res

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	c *gin.Context
	httpStatusCode int
	response ApiResponse
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, httpCode int, code string, msg string, data any) AppHttpResponse {
	httpStatusCode, res := NewSuccessMessage(httpCode, code, msg, data)
	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}

func NewErrorJsonResponse(c *gin.Context, httpCode int, code string, err error) AppHttpResponse {
	httpStatusCode, res := NewFailedMessage(httpCode, code, err)
	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}