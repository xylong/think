package bird

import "think/mode/model"

// 鸟
type Bird struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func New(a ...model.Attr) model.AnimalCreate {
	return func(a ...model.Attr) interface{} {
		bird := &Bird{}
		model.Attrs(a).Apply(bird)
		return bird
	}
}
