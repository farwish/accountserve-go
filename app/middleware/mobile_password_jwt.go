package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/ginger/app/param"
	"github.com/gopher-lego/ginger/app/repository"
	"github.com/gopher-lego/response"
	"log"
	"time"
)

func MobilePasswordJwtMiddleware() *jwt.GinJWTMiddleware {
	identityKey := "member_id"

	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 "member zone",
		SigningAlgorithm:      "",
		Key:                   []byte("member secret key!"),
		Timeout:               time.Hour * 24 * 7,
		MaxRefresh:            time.Hour * 24 * 7,
		IdentityKey:           identityKey,

		// the data returned from the authenticator is passed in as a parameter into the PayloadFunc
		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var params param.MobilePassword
			if err := c.ShouldBind(&params); err != nil {
				return nil, errors.New("参数有误")
			}

			if memberResultData, error := repository.MemberMobilePasswordQuery(params.Mobile, params.Password); error == nil {
				return memberResultData, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// LoginHandler will call Authenticator, it is Member Login logic.
		// This function is called after having successfully authenticated (logged in)
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*repository.MemberResultData); ok {
				return jwt.MapClaims{
					identityKey: v.MemberId,
				}
			}

			return jwt.MapClaims{}
		},
		// After having successfully authenticated with Authenticator
		LoginResponse: func(context *gin.Context, code int, token string, i2 time.Time) {
			response.Success(context, token)
		},

		IdentityHandler:       nil,
		Authorizator:          nil,

		LogoutResponse:        nil,

		RefreshResponse:       nil,

		Unauthorized:          nil,

		// TokenLookup:           "",
		// TokenHeadName:         "",
		// TimeFunc:              nil,
		// HTTPStatusMessageFunc: nil,
		// PrivKeyFile:           "",
		// PubKeyFile:            "",
		// SendCookie:            false,
		// CookieMaxAge:          0,
		// SecureCookie:          false,
		// CookieHTTPOnly:        false,
		// CookieDomain:          "",
		// SendAuthorization:     false,
		// DisabledAbort:         false,
		// CookieName:            "",
		// CookieSameSite:        0,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return jwtMiddleware
}