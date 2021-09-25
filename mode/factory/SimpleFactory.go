package factory

import (
	"think/mode/model"
	"think/mode/model/bird"
	"think/mode/model/fish"
)

const (
	Fish = iota
	Bird
)

// 用户类型
type AnimalType int

// CreateAnimal 创建动物
func CreateAnimal(at AnimalType) model.AnimalCreate {
	switch at {
	case Fish:
		return fish.New()
	case Bird:
		return bird.New()
	default:
		return nil
	}
}
