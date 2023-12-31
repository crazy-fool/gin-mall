package main

import (
	"gin-mall/pkg/db"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

// dataMap mapping relationship
var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	// int mapping
	"int": func(columnType gorm.ColumnType) (dataType string) { return "int32" },

	// bool mapping
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		ct, _ := columnType.ColumnType()
		if strings.HasPrefix(ct, "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},
	//"datetime": func(columnType gorm.ColumnType) (dataType string) {
	//},
}

func main() {

	//logger := log.GetLog()
	//conf := config.GetConfig()
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/gen/query",
		Mode:    gen.WithQueryInterface,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		//FieldCoverable: true,
		// if you want generate field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	g.UseDB(db.GetDB())

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// 想对已有的model生成crud等基础方法可以直接指定model struct ，例如model.User{}
	// 如果是想直接生成表的model和crud方法，则可以指定表的名称，例如g.GenerateModel("company")
	// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	//g.ApplyBasic(model.Category{})

	//g.WithJSONTagNameStrategy()

	g.WithFileNameStrategy(func(tableName string) (fileName string) {
		return strings.TrimPrefix(tableName, "hmh_")
	})
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		tableName = strings.TrimPrefix(tableName, "hmh_")
		tableName = strings.Replace(tableName, "_", " ", -1)
		tableName = strings.Title(tableName)
		return strings.Replace(tableName, " ", "", -1)
	})
	g.WithDataTypeMap(dataMap)
	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))
	//g.GenerateAllTable()
	g.ApplyBasic(g.GenerateAllTable()...)
	// execute the action of code generation
	g.Execute()

}
