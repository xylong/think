package mode

import (
	"regexp"
	"strconv"
)

type IExpression interface {
	Interpret(*User) bool
}

type User struct {
	ID   int
	Name string
	Age  int
}

// UserFilter 用户切片过滤器
type UserFilter struct {
	Expression IExpression
}

func NewUserFilter(rule string) *UserFilter {
	list := regexp.MustCompile("\\s+").Split(rule, -1)
	if len(list) != 3 {
		panic("error rule")
	}
	if list[0] == "age" {
		operator := list[1]
		value, err := strconv.Atoi(list[2])
		if err != nil {
			panic("error rule value")
		}

		return &UserFilter{
			Expression: &AgeExpression{
				operator: operator,
				value:    value,
			},
		}
	}

	panic("found no expression")
}

// Filter 过滤切片
func (uf *UserFilter) Filter(users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if uf.Expression.Interpret(user) {
			result = append(result, user)
		}
	}
	return result
}

// AgeExpression 年龄表达式
type AgeExpression struct {
	operator string
	value    int
}

func (ae *AgeExpression) Interpret(user *User) bool {
	switch ae.operator {
	case ">":
		return user.Age > ae.value
	case ">=":
		return user.Age >= ae.value
	case "<":
		return user.Age < ae.value
	case "<=":
		return user.Age <= ae.value
	default:
		return user.Age == ae.value
	}
}
