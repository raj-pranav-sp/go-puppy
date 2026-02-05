package gopuppy

import (
	"fmt"

	godog "github.com/raj-pranav-sp/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Sit() string {
	return "Sitting"
}

func BigBark() string {
	return godog.WhenGrownUp(Bark()) // calling function from go-dog module
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}
