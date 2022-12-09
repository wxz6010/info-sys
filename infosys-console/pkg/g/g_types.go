package g

import (
	"com.alex.infosys/pkg/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Context struct {
	Gin *gin.Context
	Uid uint
}

type HandleFun func(ctx *Context) (interface{}, error)

func (c *Context) BindParams(req interface{}) error {

	if err := c.Gin.ShouldBind(req); err != nil {
		return errs.AddErrorContext(errs.BadRequest.Wrapf(err, "参数"), "uri", c.Gin.Request.RequestURI)
	}
	return nil
}

type Response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Total      int         `json:"total"`
} //@name Response

type LoginClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}
