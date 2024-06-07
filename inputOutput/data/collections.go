package data

import "fmt"

var (
	Countries [10]string
	Slice     []int
	Codes     map[int]bool
)

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"

	// qty := len(Countries)

	fmt.Println("Countries saved")
}
