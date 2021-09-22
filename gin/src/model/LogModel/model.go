package LogModel

import "time"

// Log 日志
type Log struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;" json:"id" form:"id"`
	Name      string    `gorm:"type:varchar(20);not null;comment:'名称'" json:"name" form:"name"`
	CreatedAt time.Time `gorm:"type:datetime;not null;comment:'创建时间'" json:"created_at"`
}

func (u *Log) TableName() string {
	return "logs"
}

func New(attrs ...Attr) *Log {
	log := &Log{}
	Attrs(attrs).Apply(log)
	return log
}
