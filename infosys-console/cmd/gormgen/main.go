package main

import (
	"com.alex.infosys/pkg/g"
	"gorm.io/gen"
)

func main() {
	ge := gen.NewGenerator(gen.Config{
		OutPath:      "./app/dal/query",
		ModelPkgPath: "./app/dal/model",
		Mode:         gen.WithoutContext,
	})
	//db, _ := gorm.Open(mysql.Open(conf.MYSQLNDS))
	ge.UseDB(g.DB())
	ge.ApplyBasic(ge.GenerateAllTable()...)

	//ge.ApplyInterface(func(method bus.Method) {}, ge.GenerateModel("users"))
	ge.Execute()
}
