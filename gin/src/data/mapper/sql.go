package mapper

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

// Mapper 转化SqlMapper
func Mapper(sql string, args []interface{}, err error) *SqlMapper {
	if err != nil {
		panic(err)
	}
	return NewSqlMapper(sql, args)
}
