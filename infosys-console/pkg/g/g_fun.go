package g

import (
	"com.alex.infosys/conf"
	"com.alex.infosys/pkg/errs"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

func SuccessResponse(r *Response) *Response {
	r.Success = true
	r.StatusCode = http.StatusOK
	return r
}

func WarpContext(handle HandleFun) gin.HandlerFunc {
	return func(g *gin.Context) {
		data, err := handle(&Context{Gin: g})
		if err != nil {
			AbortErr(err, g)
			return
		}
		if r, ok := data.(*Response); ok {
			g.AbortWithStatusJSON(http.StatusOK, SuccessResponse(r))
		} else {
			g.AbortWithStatusJSON(http.StatusOK, SuccessResponse(&Response{
				Data: data,
			}))
		}
	}
}

func Gin() *gin.Engine {
	ginEngineOnce.Do(func() {
		ginEngine = gin.Default()
	})
	return ginEngine
}

func DB() *gorm.DB {
	gormDbOnce.Do(func() {
		db, err := gorm.Open(mysql.Open(conf.MYSQLNDS), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true,
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					Colorful:                  false,       // Disable color
				}),
		})
		if err != nil {
			panic(errs.Wrap(err, "初始数据库错误"))
		}
		gormDb = db
	})
	return gormDb
}

func HandleErrs(err error) *Response {

	r := &Response{
		Message: err.Error(),
	}
	switch errs.GetType(err) {
	case errs.BadRequest:
		r.StatusCode = http.StatusBadRequest
	case errs.Unauthorized:
		r.StatusCode = http.StatusUnauthorized
	default:
		r.StatusCode = http.StatusExpectationFailed
	}
	return r
}

func AbortErr(err error, c *gin.Context) {
	r := HandleErrs(err)
	c.AbortWithStatusJSON(r.StatusCode, r)
}
