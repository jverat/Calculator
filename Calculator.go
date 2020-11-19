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

//Quita la primera posición del slice
func trimOp(positionsArray []Op) []Op {

	for i := 0; i < len(positionsArray)-1; i++ {
		positionsArray[i] = (positionsArray)[i+1]
	}

	positionsArray = positionsArray[0 : len(positionsArray)-1]
	return positionsArray
}

//Elimina del slice de operaciones aquellas que ya se ejecutaron, y repara el slice que determina el orden de ejecución para que coincida con las nuevas posiciones
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

	for i := range positionsArray {
		if positionsArray[i].pos > pos {
			positionsArray[i].pos--
		}
	}

	return operations, positionsArray
}

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, fmt.Errorf("factorial de número negativo no existe")
	}

	if n > 50 {
		return 0, fmt.Errorf("número demasiado grande para que el resultado sea computable")
	}

	ans := 1

	for i := 1; i <= n; i++ {
		ans *= i
	}

	return ans, nil
}

func hierarchicalExecution(numbers []int, operations []string, positionsArray []Op) (int, error) {

	switch operations[positionsArray[0].pos] {
	case "!":
		{
			ans , err := factorial(numbers[positionsArray[0].pos])

			if err != nil {
				return 0, err
			}

			numbers[positionsArray[0].pos] = ans
		}
	case "^":
		{
			numbers[positionsArray[0].pos] = int(math.Pow(float64(numbers[positionsArray[0].pos]), float64(numbers[positionsArray[0].pos])))
			numbers[positionsArray[0].pos] = math.MaxInt64
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
		return numbers[0], nil
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

// Se convierten los números desde el string para posteriormente ejecutar las operaciones
func operation(processableExpression string) (int, error) {

	var numbers []int
	var operations []string
	var positionsArray []Op
	var indexN int

	if !strings.ContainsAny(processableExpression, "1234567890+-*/%^!") {
		return 0, fmt.Errorf("la expresión no contiene caracteres computables: %s", processableExpression)
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
				return 0, err
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

//Se recibe la expresión. Es donde se garantiza que los parentesis se resuelvan primero
func expressionProcessing(expression string) (string, error) {

	//Operaciones de signos
	expression = strings.ReplaceAll(expression, "--", "+")
	expression = strings.ReplaceAll(expression, "++", "+")
	expression = strings.ReplaceAll(expression, "+-", "-")
	expression = strings.ReplaceAll(expression, "-+", "-")

	openingParenthesisPos := -1

	if !strings.ContainsAny(expression, "()") {
		ans, err := operation(expression)

		if err != nil {
			return "", err
		}

		return strconv.Itoa(ans), nil
	} else {

		//Se ejecutan los parentesis antes que cualquier otra parte de la expresión
		for i, ch := range expression {
			switch string(ch) {
			case "(":
				{
					openingParenthesisPos = i
				}
			case ")":
				{
					ans, err := operation(expression[openingParenthesisPos+1 : i])
					if err != nil {
						return "", err
					}
					expression = strings.Replace(expression, expression[openingParenthesisPos : i+1], strconv.Itoa(ans), 1)
					if strings.ContainsAny(expression, "!^*/%+-") {
						return expressionProcessing(expression)
					} else {
						return expression, nil
					}
				}
			}
		}

		if strings.ContainsAny(expression, "!^*/%+-") {
			return expressionProcessing(expression)
		} else {
			return expression, nil
		}
	}
}

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa tu operación:")
		expression, err := reader.ReadString('\n')

		if err != nil {
			fmt.Errorf(err.Error())
			os.Exit(2)
		}

		if expression == "q" {
			break
		} else {
			fmt.Println(expressionProcessing(expression))
		}
	}
}
