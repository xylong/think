package mapper

import (
	"think/gin/src/model/LogModel"

	"github.com/Masterminds/squirrel"
)

type LogMapper struct {
}

// AddLog 添加日志
func (lm *LogMapper) AddLog(log *LogModel.Log) *SqlMapper {
	return Mapper(squirrel.Insert(log.TableName()).
		Columns("name", "created_at").
		Values(log.Name, log.CreatedAt).
		ToSql())
}
