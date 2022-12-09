package main

import (
	_ "com.alex.infosys/app/api/system"
	"com.alex.infosys/pkg/g"
	"github.com/gin-gonic/gin"
	"net/http"
)

//@title gin cms api
//@version 0.0.1
//@BasePath /api/cms
func main() {
	g.Gin().Use(gin.Recovery())
	//g.Gin().Use(service.WithDbContext())
	g.Gin().GET("/api/openapi", func(c *gin.Context) {
		r, e := g.SwagJson2OpenApi("./docs/swagger.json")
		if e != nil {
			g.AbortErr(e, c)
			return
		}
		c.String(http.StatusOK, r)
	})

	g.Gin().Run(":8081")
}
