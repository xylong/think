package fish

import "think/mode/model"

// 鱼
type Fish struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func New(a ...model.Attr) model.AnimalCreate {
	return func(a ...model.Attr) interface{} {
		fish := &Fish{}
		model.Attrs(a).Apply(fish)
		return fish
	}
}
