package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = errors.New("you don't have permission to access this resource")

	// ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrMissingLoginValues = errors.New("missing Username or Password")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = errors.New("incorrect Username or Password")
)

type resUser struct {
	ID       string                 `json:"id"`
	Email    string                 `json:"email"`
	Metadata map[string]interface{} `json:"metadata"`
}

type respBody struct {
	Code  int         `json:"code"` // 0 or 1, 0 is ok, 1 is error
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func respErr(c *gin.Context, code int, err error) {
	glog.Warningln(err)

	c.JSON(code, &respBody{
		Code:  code,
		Error: err.Error(),
	})
}

func respOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &respBody{
		Code: http.StatusOK,
		Data: data,
	})
}
