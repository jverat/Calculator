package calculator

import (
	"fmt"
	"math"
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
func trimN(numbers []int, index int) []int {

	if index == 0 {
		numbers = numbers[1:]
	} else if index == len(numbers)-1 {
		numbers = numbers[:index]
	} else {
		numbers = append(numbers[:index], numbers[index+1:]...)
	}

	return numbers
}

//Quita la primera posición del slice
func trimOp(positionsArray []Op) []Op {
	positionsArray = positionsArray[1:]
	return positionsArray
}

//Elimina del slice de operaciones aquellas que ya se ejecutaron, y repara el slice que determina el orden de ejecución para que coincida con las nuevas posiciones
func fixArrays(operations []string, index int, positionsArray []Op) ([]string, []Op) {

	/*for j := index; j < len(operations)-1; j++ {
		operations[j] = operations[j+1]
	}

	operations = operations[0 : len(operations)-1]*/

	if index == 0 {
		operations = operations[1:]
	} else if index == len(operations) {
		operations = operations[:len(operations)-1]
	} else {
		operations = append(operations[:index], operations[index+1:]...)
	}

	for i := range positionsArray {
		if positionsArray[i].pos > index {
			positionsArray[i].pos--
		}
	}

	return operations, positionsArray
}

func fixArraysMulti(operations []string, index []int, positionsArray []Op) ([]string, []Op) {
	for _, n := range index {
		operations, positionsArray = fixArrays(operations, n, positionsArray)
	}

	return operations, positionsArray
}

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, fmt.Errorf("factorial de número negativo no existe")
	}

	if n > 20 {
		return 0, fmt.Errorf("número demasiado grande para que el resultado sea computable")
	}

	ans := 1

	for i := 1; i <= n; i++ {
		ans *= i
	}

	return ans, nil
}

//Se ejecutan las operaciones siguiendo el orden determinado en positionsArray
func hierarchicalExecution(numbers []int, operations []string, positionsArray []Op) (int, error) {

	switch operations[positionsArray[0].pos] {
	case "!":
		{
			fmt.Printf("%d!\n", numbers[positionsArray[0].pos])

			ans, err := factorial(numbers[positionsArray[0].pos])

			if err != nil {
				return 0, err
			}

			numbers[positionsArray[0].pos] = ans
		}
	case "^":
		{
			fmt.Printf("%f ^ %f = %f\n", float64(numbers[positionsArray[0].pos]), float64(numbers[positionsArray[0].pos+1]), math.Pow(float64(numbers[positionsArray[0].pos]), float64(numbers[positionsArray[0].pos])))

			numbers[positionsArray[0].pos] = int(math.Pow(float64(numbers[positionsArray[0].pos]), float64(numbers[positionsArray[0].pos])))
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	case "*":
		{
			fmt.Printf("%d %s %d\n", numbers[positionsArray[0].pos], operations[positionsArray[0].pos], numbers[positionsArray[0].pos+1])

			numbers[positionsArray[0].pos] *= numbers[positionsArray[0].pos+1]
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	case "/":
		{
			fmt.Printf("%d %s %d\n", numbers[positionsArray[0].pos], operations[positionsArray[0].pos], numbers[positionsArray[0].pos+1])

			numbers[positionsArray[0].pos] /= numbers[positionsArray[0].pos+1]
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	case "%":
		{
			fmt.Printf("%d %s %d\n", numbers[positionsArray[0].pos], operations[positionsArray[0].pos], numbers[positionsArray[0].pos+1])

			numbers[positionsArray[0].pos] %= numbers[positionsArray[0].pos+1]
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	case "+":
		{
			fmt.Printf("%d %s %d\n", numbers[positionsArray[0].pos], operations[positionsArray[0].pos], numbers[positionsArray[0].pos+1])

			numbers[positionsArray[0].pos] += numbers[positionsArray[0].pos+1]
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	case "-":
		{
			fmt.Printf("%d %s %d\n", numbers[positionsArray[0].pos], operations[positionsArray[0].pos], numbers[positionsArray[0].pos+1])

			numbers[positionsArray[0].pos] -= numbers[positionsArray[0].pos+1]
			numbers = trimN(numbers, positionsArray[0].pos+1)
		}
	}

	if len(positionsArray) == 1 {
		return numbers[0], nil
	} else {
		//operations[positionsArray[0].pos] = ""
		trimmedPosArr := trimOp(positionsArray)
		ops, posArr := fixArrays(operations, positionsArray[0].pos, trimmedPosArr)
		posArr = organizeOps(posArr)

		return hierarchicalExecution(numbers, ops, posArr)
	}
}

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
		} else {
			return 0, fmt.Errorf("la expresion contiene carácteres ilícitos: %s", string(ch))
		}
	}

	//Se computan las operaciones determinando su jerarquía para posteriormente organizarlas por orden de ejecución
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

	return hierarchicalExecution(numbers, operations, a) //, sequentialExecution(&numbers, operations)
}

//Se recibe la expresión. Es donde se garantiza que los parentesis se resuelvan primero
func ExpressionProcessing(expression string) (string, error) {

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
					expression = strings.Replace(expression, expression[openingParenthesisPos:i+1], strconv.Itoa(ans), 1)
					if strings.ContainsAny(expression, "!^*/%+-") {
						return ExpressionProcessing(expression)
					} else {
						return expression, nil
					}
				}
			}
		}

		if strings.ContainsAny(expression, "!^*/%+-") {
			return ExpressionProcessing(expression)
		} else {
			return expression, nil
		}
	}
}
