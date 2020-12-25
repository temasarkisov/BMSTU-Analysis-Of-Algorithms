package ant

import (
	"fmt"
	"math/rand"
	"os"
)

// GenData - data generator
func GenData(n int) {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				str := fmt.Sprintf("%d ", rand.Intn(10)+1)
				file.WriteString(str)
			} else {
				str := fmt.Sprintf("%d ", 0)
				file.WriteString(str)
			}
		}
		str := fmt.Sprintf("\n")
		file.WriteString(str)
	}
	file.Close()
}
