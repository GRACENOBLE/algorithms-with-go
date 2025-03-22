package main

// "github.com/GRACENOBLE/algorithms-with-go/sorting"
import (
	"fmt"

	"github.com/GRACENOBLE/algorithms-with-go/substrings"
)

func main() {
	count := substrings.CountVowelSubstrings("cuaieuouac")
	fmt.Printf("The count is : %v", count)
}
