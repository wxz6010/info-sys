package g

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
)

var (
	ginEngineOnce sync.Once
	ginEngine     *gin.Engine
	gormDb        *gorm.DB
	gormDbOnce    sync.Once
)
