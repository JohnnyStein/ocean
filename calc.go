package ocean

import (
	"fmt"
)

func Calc(operador string, num1 int, num2 int) {
	if operador == "+" {
		fmt.Println(num1 + num2)
	} else if operador == "-" {
		fmt.Println(num1 - num2)
	} else if operador == "*" {
		fmt.Println(num1 * num2)
	} else if operador == "/" {
		fmt.Println(num1 / num2)
	} else if operador == "%" {
		fmt.Println(num1 % num2)
	}
}
