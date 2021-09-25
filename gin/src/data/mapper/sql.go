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
	// 事物执行db
	db *gorm.DB
}

func NewSqlMapper(sql string, args []interface{}) *SqlMapper {
	return &SqlMapper{Sql: sql, Args: args}
}

// Query 查询
func (sm *SqlMapper) Query() *gorm.DB {
	if sm.db != nil {
		return sm.db.Raw(sm.Sql, sm.Args...)
	}
	return db.Orm.Raw(sm.Sql, sm.Args...)
}

// Exec 执行编辑操作
func (sm *SqlMapper) Exec() *gorm.DB {
	if sm.db != nil {
		return sm.db.Exec(sm.Sql, sm.Args...)
	}
	return db.Orm.Exec(sm.Sql, sm.Args...)
}

// 设置事物执行db
func (sm *SqlMapper) setDB(db *gorm.DB) {
	sm.db = db
}

// Mapper 转化SqlMapper
func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err)
	}
	return NewSqlMapper(sql, args)
}

type SqlMappers []*SqlMapper

// ?Exec 执行事务
// * f 事务内容
func (sm SqlMappers) Exec(f func() error) error {
	return db.Orm.Transaction(func(tx *gorm.DB) error {
		// 将要执行事务的sql的db设置成同一个
		sm.apply(tx)
		return f()
	})
}

func (sm SqlMappers) apply(db *gorm.DB) {
	for _, item := range sm {
		item.setDB(db)
	}
}

// Mappers 将多个sql放入sql集合
func Mappers(sqlMappers ...*SqlMapper) SqlMappers {
	return sqlMappers
}
