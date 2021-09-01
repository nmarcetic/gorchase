package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginReq defines User login credentials
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterReq defins User registration payload
type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getCurrentUser(c *gin.Context) (u UserContext, err error) {
	data, ok := c.Get(userKey)
	if !ok {
		return UserContext{}, fmt.Errorf("Can't get current user from context")
	}
	usr, ok := data.(UserContext)
	if !ok {
		return UserContext{}, fmt.Errorf(" 2 Can't get current user from context")
	}

	return usr, nil
}
