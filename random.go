package ocean

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Random(max int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var err error
	var min int

	nargs := len(os.Args)
	switch nargs {
	case 1:
		max = 100
	case 2:
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	case 3:
		min, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		max, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println((min + rand.Intn(max-min)))
}
