package api

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nmarcetic/gorchase/server/accounts"
)

var (
	identityKey      = "ID"
	identityEmail    = "Email"
	identityMetadata = "Metadata"
	issuer           = "nmarcetic"
	//TODO: load from env
	defSecret     = "nmarcetic-GORCHASE-IIq52JIyECl4loVRc8MHopW70erYHqwrZP2YCEXvWIm07LttOf"
	defJWTTimeout = time.Hour * 3 // 3 hours
)

// UserContext defines user object which we propagates throu the context
// After successful authentication object is avaiblable in *gin.Context as User
type UserContext struct {
	ID       string
	Email    string
	Metadata accounts.Metadata
}

type usersHandler struct {
	service accounts.Service
}

// AuthMiddleware returns auth middleware
func AuthMiddleware(svc accounts.Service) *jwt.GinJWTMiddleware {

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       issuer,
		Key:         []byte(defSecret),
		Timeout:     defJWTTimeout,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserContext); ok {
				return jwt.MapClaims{
					identityKey:      v.ID,
					identityEmail:    v.Email,
					identityMetadata: v.Metadata,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			var u UserContext
			u.ID = claims[identityKey].(string)
			u.Email = claims[identityEmail].(string)
			u.Metadata = claims[identityMetadata].(map[string]interface{})
			return u
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals LoginReq
			if err := c.Bind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userEmail := loginVals.Email
			userPassword := loginVals.Password

			if userEmail == "" || userPassword == "" {
				return "", jwt.ErrMissingLoginValues
			}
			user, err := svc.Login(userEmail, userPassword)
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			return &UserContext{
				ID:       user.ID,
				Email:    user.Email,
				Metadata: user.Metadata,
			}, nil
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	if v, ok := data.(*UserContext); ok && v.UserName == "admin" {
		// 		return true
		// 	}

		// 	return false
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware

}
