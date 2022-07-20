package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/livoir/goose-migration/internal/core/domain/api"
	"github.com/livoir/goose-migration/internal/core/ports"
	"net/http"
	"strconv"
)

type APIHandler struct {
	UserService ports.UserService
}

func NewAPIHandler(userService ports.UserService) *APIHandler {
	return &APIHandler{UserService: userService}
}

func (handler *APIHandler) GetAll(c *gin.Context) {
	users, err := handler.UserService.GetAll(c)
	if err != nil {
		res := gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "Internal Error",
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := api.ResponseAPI{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   users,
	}
	fmt.Println(res)

	c.JSON(http.StatusOK, res)

}

func (handler *APIHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		res := gin.H{
			"code":    http.StatusBadRequest,
			"status":  "Bad Request",
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	users, err := handler.UserService.GetById(c, id)

	if err != nil {
		var code int
		var status, message string
		if err == sql.ErrNoRows {
			code = http.StatusNotFound
			status = "Not Found"
			message = fmt.Sprintf("User id %d not found", id)
		} else {
			code = http.StatusInternalServerError
			status = "Internal Error"
			message = err.Error()
		}

		res := gin.H{
			"code":    code,
			"status":  status,
			"message": message,
		}
		c.AbortWithStatusJSON(code, res)
		return
	}

	res := api.ResponseAPI{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   users,
	}

	c.JSON(http.StatusOK, res)

}
