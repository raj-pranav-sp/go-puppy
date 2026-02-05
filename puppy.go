package gopuppy

import (
	godog "github.com/raj-pranav-sp/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Sit() string {
	return "Sitting"
}

func BigBark() string {
	return godog.WhenGrownUp(Bark())
}
