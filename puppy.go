package puppy

import "github.com/tinashe-mundangepfupfu/dog"

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() string {
	return "this is version v1.1.0"
}

func From12() string {
	return "this is version v1.2.0"
}
