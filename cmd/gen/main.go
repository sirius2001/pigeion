package main

import (
	"pigeon2/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const DSN = "root:root@(127.0.0.1:3306)/pigeon2?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.WorkFlow{}, &model.WorkNode{}, &model.WorkOrder{})

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:       "pkg/persistence",                                                  // 生成代码的输出目录
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	})

	// 设置数据库连接
	g.UseDB(db)

	g.ApplyBasic(&model.WorkFlow{}, &model.WorkNode{}, &model.WorkOrder{})

	// 执行生成
	g.Execute()
}
