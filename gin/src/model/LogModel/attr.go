package LogModel

import "time"

type Attrs []Attr

func (a Attrs) Apply(log *Log) {
	for _, f := range a {
		f(log)
	}
}

type Attr func(*Log)

// WithName 设置名称
func WithName(name string) Attr {
	return func(l *Log) {
		l.Name = name
	}
}

// WithCreatedAt 设置创建时间
func WithCreatedAt(createdAt time.Time) Attr {
	return func(l *Log) {
		l.CreatedAt = createdAt
	}
}
