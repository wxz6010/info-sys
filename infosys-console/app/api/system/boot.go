package system

import (
	"com.alex.infosys/app/dal/query"
	"com.alex.infosys/pkg/g"
)

func init() {
	db := g.DB()
	q := query.Use(db)
	orgs := &orgs{
		db: db,
		q:  q,
	}
	orgsGroup := g.Gin().Group("/api/system/orgs")
	orgsGroup.POST("/updates", g.WarpContext(orgs.createOrupdate))
	orgsGroup.POST("/query", g.WarpContext(orgs.conditionQuery))

	orgsType := &orgType{
		db: db,
		q:  q,
	}
	orgsTypeGroup := orgsGroup.Group("/type")
	orgsTypeGroup.POST("/updates", g.WarpContext(orgsType.createOrupdate))
	orgsTypeGroup.POST("/query", g.WarpContext(orgsType.conditionQuery))

	dictionary := &dictionary{
		db: db,
		q:  q,
	}

	dicGroup := g.Gin().Group("/api/system/dic/")
	dicGroup.POST("/updates", g.WarpContext(dictionary.createOrupdate))
	dicGroup.POST("/query", g.WarpContext(dictionary.conditionQuery))

	sysRoles := &roles{
		db: db,
		q:  q,
	}
	rolesGroup := g.Gin().Group("/api/system/roles")
	rolesGroup.POST("/updates", g.WarpContext(sysRoles.createOrupdate))
	rolesGroup.POST("/query", g.WarpContext(sysRoles.conditionQuery))
}
