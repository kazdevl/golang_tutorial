package parallel

import "fmt"

func Trace(name string) func() {
	fmt.Printf("[entered]: %s\n", name)
	return func() {
		fmt.Printf("[returned]: %s\n", name)
	}
}
