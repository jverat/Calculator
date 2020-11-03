package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// TODO
// comprobar la validez de las expresiones

type Op struct {
	pos   int
	grade int
}

//Organiza descendentemente los elementos del arreglo según su grade
func organizeOps(operations []Op) []Op {
	var result []Op

	for i, currentGrade := 0, 0; len(operations) != len(result); i++ {
		if operations[i].grade == currentGrade {
			result = append(result, operations[i])
		}
		if i == len(operations)-1 {
			i = -1
			currentGrade++
		}
	}

	return result
}

//Borra los espacios vacios del slice
func trimN(numbers []int) []int {

	for i, n := range numbers {
		if n == math.MaxInt64 {
			for j := i; j < len(numbers)-1; j++ {
				numbers[j] = numbers[j+1]
			}
			numbers = numbers[0 : len(numbers)-1]
		}
	}

	return numbers
}

func trimOp(positionsArray []Op) []Op {

	for i := 0; i < len(positionsArray)-1; i++ {
		positionsArray[i] = (positionsArray)[i+1]
	}

	positionsArray = positionsArray[0 : len(positionsArray)-1]
	return positionsArray
}

func fixArrays(operations []string, positionsArray []Op) ([]string, []Op) {

	var pos int

	for i, n := range operations {
		if n == "" {
			pos = i
			for j := i; j < len(operations)-1; j++ {
				operations[j] = operations[j+1]
			}
			operations = operations[0 : len(operations)-1]
		}
	}

	//fmt.Println("pos: ", pos)

	for i := range positionsArray {
		if positionsArray[i].pos > pos {
			positionsArray[i].pos--
		}
	}

	/*fmt.Println("++++++++++++")
	  fmt.Println("positionsArray")
	  for i, posOp := range positionsArray {
	    fmt.Println("#", i, ": ", posOp)
	  }

	  fmt.Println("++++++++++++")
	  fmt.Println("operationsArray")
	  for i, op := range operations {
	    fmt.Println("#", i, ": ", op)
	  }*/

	return operations, positionsArray
}


func hierarchicalExecution(numbers []int, operations []string, positionsArray []Op) int {

	switch operations[positionsArray[0].pos] {
	case "!":
		{
			n := 0

			for i := numbers[positionsArray[0].pos]; i > 0; i-- {
				if n == 0 {
					n += i
				} else {
					n *= i
				}
			}

			numbers[positionsArray[0].pos] = n
		}
	case "^":
		{
			numbers[positionsArray[0].pos] = int(math.Pow(float64(numbers[positionsArray[0].pos]), float64(numbers[positionsArray[0].pos])))
			numbers[(positionsArray)[0].pos] = math.MaxInt64
			trimN(numbers)
		}
	case "*":
		{
			numbers[positionsArray[0].pos] *= numbers[positionsArray[0].pos+1]
			numbers[positionsArray[0].pos+1] = math.MaxInt64
			trimN(numbers)
		}
	case "/":
		{
			numbers[positionsArray[0].pos] /= numbers[positionsArray[0].pos+1]
			numbers[positionsArray[0].pos+1] = math.MaxInt64
			trimN(numbers)
		}
	case "%":
		{
			numbers[positionsArray[0].pos] %= numbers[positionsArray[0].pos+1]
			numbers[positionsArray[0].pos+1] = math.MaxInt64
			trimN(numbers)
		}
	case "+":
		{
			numbers[positionsArray[0].pos] += numbers[positionsArray[0].pos+1]
			numbers[positionsArray[0].pos+1] = math.MaxInt64
			trimN(numbers)
		}
	case "-":
		{
			numbers[positionsArray[0].pos] -= numbers[positionsArray[0].pos+1]
			numbers[positionsArray[0].pos+1] = math.MaxInt64
			trimN(numbers)
		}
	}

	if len(positionsArray) == 1 {
		return numbers[0]
	} else {
		operations[positionsArray[0].pos] = ""
		trimmedPosArr := trimOp(positionsArray)
		ops, posArr := fixArrays(operations, trimmedPosArr)
		organizeOps(posArr)
		return hierarchicalExecution(numbers, ops, posArr)
	}
}

//Se ejecutan las operaciones según el orden en el que aparecen en la expresión
/*func sequentialExecution(numbers *[]int, operations []string) int {

	indexO := 0

	for i, n := range *numbers {

		if i == len(*numbers)-1 && operations[len(operations)-1] != "!" {
			break
		}

		switch operations[indexO] {
		case "+":
			(*numbers)[i+1] += n
		case "-":
			(*numbers)[i+1] -= n
		case "*":
			(*numbers)[i+1] *= n
		case "/":
			(*numbers)[i+1] = n / (*numbers)[i+1]
		case "%":
			(*numbers)[i+1] = n % (*numbers)[i+1]
		case "^":
			(*numbers)[i+1] = int(math.Pow(float64(n), float64((*numbers)[i+1])))
		case "!":
			{
				if (*numbers)[i] == 0 {
					(*numbers)[i] = 0
				} else {
					for j := 1; j < n; j++ {
						(*numbers)[i] *= n - j
					}
				}
			}
		default:
			{
				//fmt.Println("idk how you got in here but here is what is wrong i guess: " + operations[indexO])
			}
		}
		indexO++
	}

	return (*numbers)[len(*numbers)-1]
}*/

// Se convierten los números desde el string para posteriormente ejecutar las operaciones. Retorna los resultados de ambos modos de ejecución
func operation(processableExpression string) int {

	var numbers []int
	var operations []string
	var positionsArray []Op
	var indexN int

	if !strings.ContainsAny(processableExpression, "1234567890+-*/%^!") {
		fmt.Println("Invalid expression")
		os.Exit(2)
	}

	for i, ch := range processableExpression {
		//Se obtiene el número digito a digito
		if strings.ContainsAny(string(ch), "1234567890") {
			var n string
			for j := i; j < len(processableExpression) && strings.ContainsAny(string(processableExpression[j]), "1234567890"); j++ {
				n += string(processableExpression[j])
			}

			//string a int
			number, err := strconv.Atoi(n)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			numbers = append(numbers, number)
			indexN++
		} else if strings.ContainsAny(string(ch), "+-*/%^!") {
			//Se obtienen los operadores
			operations = append(operations, string(ch))
		}
	}

	for i, o := range operations {
		switch o {
		case "!":
			positionsArray = append(positionsArray, Op{i, 0})
		case "^":
			positionsArray = append(positionsArray, Op{i, 1})
		case "*":
			positionsArray = append(positionsArray, Op{i, 2})
		case "/":
			positionsArray = append(positionsArray, Op{i, 2})
		case "%":
			positionsArray = append(positionsArray, Op{i, 2})
		case "+":
			positionsArray = append(positionsArray, Op{i, 3})
		case "-":
			positionsArray = append(positionsArray, Op{i, 3})
		}
	}

	a := organizeOps(positionsArray)

	return hierarchicalExecution(numbers, operations, a)//, sequentialExecution(&numbers, operations)
}

func expressionProcessing(expression string) string {

	//Operaciones de signos
	expression = strings.ReplaceAll(expression, "--", "+")
	expression = strings.ReplaceAll(expression, "++", "+")
	expression = strings.ReplaceAll(expression, "+-", "-")
	expression = strings.ReplaceAll(expression, "-+", "-")
	fmt.Println("Flag: ", expression)

	openingParenthesisPos := -1

	if !strings.ContainsAny(expression, "()") {
		return strconv.Itoa(operation(expression))
	} else {

		//Se separa la expresión en diferentes partes según como los paréntesis las discriminan
		for i, ch := range expression {
			switch string(ch) {
			case "(":
				{
					openingParenthesisPos = i
				}
			case ")":
				{
					expression = strings.Replace(expression, expression[openingParenthesisPos : i+1], strconv.Itoa(operation(expression[openingParenthesisPos+1 : i])), 1)
					if strings.ContainsAny(expression, "!^*/%+-") {
						return expressionProcessing(expression)
					} else {
						return expression
					}
				}
			}
		}

		if strings.ContainsAny(expression, "!^*/%+-") {
			return expressionProcessing(expression)
		} else {
			return expression
		}
	}
}

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa tu operación:")
		expression, _ := reader.ReadString('\n')

		if expression == "q" {
			break
		} else {
			fmt.Println(expressionProcessing(expression))
		}
	}
}
