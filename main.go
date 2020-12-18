package main

import (
	"bufio"
	"fmt"
	"os"

	"dependencies/calculator"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa tu operación:")
		expression, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		if expression == "q" {
			break
		} else {
			fmt.Println(calculator.ExpressionProcessing(expression))
		}
	}
}
