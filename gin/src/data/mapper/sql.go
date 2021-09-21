package mapper

import (
	"think/gin/src/db"

	"gorm.io/gorm"
)

// SqlMapper 通用sql
type SqlMapper struct {
	// sql语句
	Sql string
	// 参数
	Args []interface{}
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{Sql: sql, Args: args}
}

// Query 查询
func (sm *SqlMapper) Query() *gorm.DB {
	return db.Orm.Raw(sm.Sql, sm.Args...)
}

// Exec 执行编辑操作
func (sm *SqlMapper) Exec() *gorm.DB {
	return db.Orm.Exec(sm.Sql, sm.Args...)
}

// Mapper 转化SqlMapper
func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err)
	}
	return NewSqlMapper(sql, args)
}
