package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmarcetic/gorchase/server/accounts"
)

const userKey = "ID"

//UserHandler specify methods for handling user requests
type UserHandler interface {
	Singup(c *gin.Context)
	Get(c *gin.Context)
}

type userHandler struct {
	service accounts.Service
}

// NewHandler init new userHandler
func NewHandler(svc accounts.Service) UserHandler {
	return &userHandler{
		service: svc,
	}
}

func (h userHandler) Singup(c *gin.Context) {
	var req RegisterReq
	if err := c.Bind(&req); err != nil {
		respErr(c, http.StatusBadRequest, ErrMissingLoginValues)
		return
	}
	user := accounts.User{
		Email:    req.Email,
		Password: []byte(req.Password),
	}
	id, err := h.service.Register(user)
	if err != nil {
		respErr(c, http.StatusBadRequest, err)
		return
	}
	respOK(c, gin.H{"id": id})
}

func (h userHandler) Get(c *gin.Context) {
	reqUser, err := getCurrentUser(c)
	if err != nil {
		respErr(c, http.StatusForbidden, err)
		return
	}
	user, err := h.service.Get(reqUser.ID)
	if err != nil {
		respErr(c, http.StatusNotFound, err)
		return
	}
	resUser := &resUser{
		ID:       user.ID,
		Email:    user.Email,
		Metadata: user.Metadata,
	}
	respOK(c, gin.H{"user": resUser})

}
