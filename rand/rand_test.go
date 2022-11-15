package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	choices := "thisistest"
	fmt.Println(RandFromChoices(choices, 6))
	fmt.Println(RandLetters(6))
	fmt.Println(RandLetterNumbers(6))
	fmt.Println(RandLowLetterNumber(6))
	fmt.Println(RandomString(6))
	fmt.Println(RandomUA())
}
