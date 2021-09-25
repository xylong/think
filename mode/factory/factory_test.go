package factory

import (
	"testing"
	"think/mode/model/bird"
	"think/mode/model/fish"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	Convey("简单工厂模式", t, func() {
		shark, parrot := "鲨鱼", "鹦鹉"

		fish, ok := CreateAnimal(Fish)(fish.Name(shark)).(*fish.Fish)
		So(ok, ShouldEqual, true)
		So(fish.Name, ShouldEqual, shark)

		bird, ok := CreateAnimal(Bird)(bird.Name(parrot)).(*bird.Bird)
		So(ok, ShouldEqual, true)
		So(bird.Name, ShouldEqual, parrot)
	})
}
