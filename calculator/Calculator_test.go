package calculator

import (
	"fmt"
	"testing"
)

func TestOrganizeOpsBasic(t *testing.T) {
	operations := []Op{{0, 7},
		{1, 6},
		{2, 5},
		{3, 4},
		{4, 3},
		{5, 2},
		{6, 1},
		{7, 0}}

	expectedResult := []Op{{7, 0},
		{6, 1},
		{5, 2},
		{4, 3},
		{3, 4},
		{2, 5},
		{1, 6},
		{0, 7}}

	operations = organizeOps(operations)

	for i, op := range operations {
		if op != expectedResult[i] {
			t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
				if op != expectedResult[i] {
					t.Errorf("Mismatch: %+v != %+v", op, expectedResult[i])
				}
			})
		}
	}
}

func TestTrimN(t *testing.T) {
	numbers := []int{1, 0, 2, 0, 3, 0}

	expectedResult := []int{1, 2, 3}

	numbers = trimN(numbers, 5)
	numbers = trimN(numbers, 3)
	numbers = trimN(numbers, 1)

	if len(numbers) != len(expectedResult) {
		t.Errorf("Untrimmed")
	}

	for i, n := range numbers {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			if n != expectedResult[i] {
				t.Errorf("Wrong result #%d: Answer = %+d != %+d", i, n, expectedResult[i])
			}
		})
	}
}

func TestTrimOp(t *testing.T) {
	operations := []Op{{0, 0}, {1, 1}, {2, 2}}

	expectedResult := []Op{{1, 1}, {2, 2}}

	operations = trimOp(operations)

	for i, op := range operations {
		t.Run(fmt.Sprintf("Case #%d", i), func(t *testing.T) {
			if op != expectedResult[i] {
				t.Errorf("Mismatch: %+v != %+v", op, expectedResult[i])
			}
		})
	}
}

func TestFixArrays(t *testing.T) {
	operations := []string{"a", "", "b", "c"}
	positions := []Op{{0, 0}, {2, 0}, {3, 0}}

	expectedOps := []string{"a", "b", "c"}
	expectedPos := []Op{{0, 0}, {1, 0}, {2, 0}}

	operations, positions = fixArrays(operations, 1, positions)

	t.Run("sizes", func(t *testing.T) {
		if len(operations) > len(positions) {
			t.Errorf("Empty spaces remains")
		} else if len(operations) < len(positions) {
			t.Errorf("Too trimmed")
		}
	})

	for i, op := range operations {
		t.Run(fmt.Sprintf("operations_test#%d", i), func(t *testing.T) {
			if op != expectedOps[i] {
				t.Errorf("Mismatch: %s != %s", op, expectedOps[i])
			}
		})
		t.Run(fmt.Sprintf("Positions Test #%d", i), func(t *testing.T) {
			if positions[i] != expectedPos[i] {
				t.Errorf("Mismatch: %+v != %+v", op, expectedOps[i])
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	numbers := [5]int{3, 6, 9, 12, 15}

	expectedResult := [5]int{6, 720, 362880, 479001600, 1307674368000}

	for i, n := range numbers {
		numbers[i], _ = factorial(n)
		t.Run(fmt.Sprintf("Factorial test #%d", i), func(t *testing.T) {
			if numbers[i] != expectedResult[i] {
				t.Errorf("Mismatch: %d != %d", numbers[i], expectedResult[i])
			}
		})
	}
}

func TestHierarchicalExecutionSameGradeOps(t *testing.T) {
	numbers := []int{3, 6, 9, 12, 15}
	operations := []string{"+", "-", "+", "+"}
	hierarchicalOrder := []Op{{0, 3}, {1, 3}, {2, 3}, {3, 3}}

	expectedResult := 27

	ans, err := hierarchicalExecution(numbers, operations, hierarchicalOrder)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if ans != expectedResult {
		t.Errorf("Mismatch: %d != %d", ans, expectedResult)
	}
}

func TestHierarchicalExecutionVariousGradesOps(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 2}
	operations := []string{"+", "*", "+", "/"}
	hierarchicalOrder := []Op{{1, 2}, {3, 2}, {0, 3}, {2, 3}}

	expectedResult := 9

	ans, err := hierarchicalExecution(numbers, operations, hierarchicalOrder)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if ans != expectedResult {
		t.Errorf("Mismatch: %d != %d", ans, expectedResult)
	}
}

func TestHierarchicalExecutionVariousGradesOps2(t *testing.T) {
	numbers := []int{2, 2, 3, 4, 3}
	operations := []string{"^", "*", "+", "-", "!"}
	//2^2*3+4-3!
	//2^
	hierarchicalOrder := []Op{{4, 0}, {0, 1}, {1, 2}, {2, 3}, {3, 3}}

	expectedResult := 10

	ans, err := hierarchicalExecution(numbers, operations, hierarchicalOrder)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if ans != expectedResult {
		t.Errorf("Mismatch: %d != %d", ans, expectedResult)
	}
}

func TestHierarchicalExecutionVariousGradesOps3(t *testing.T) {
	numbers := []int{2, 2, 6, 3}
	operations := []string{"^", "+", "/"}
	hierarchicalOrder := []Op{{0, 1}, {2, 2}, {1, 3}}
	//2^2+6/3

	expectedResult := 6

	ans, err := hierarchicalExecution(numbers, operations, hierarchicalOrder)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if ans != expectedResult {
		t.Errorf("Mismatch: %d != %d", ans, expectedResult)
	}
}

func TestHierarchicalExecutionVariousGradesOps4(t *testing.T) {
	numbers := []int{3, 2, 2}
	operations := []string{"^", "+"}
	hierarchicalOrder := []Op{{0, 1}, {1, 3}}
	//3^2+2

	expectedResult := 11

	ans, err := hierarchicalExecution(numbers, operations, hierarchicalOrder)

	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if ans != expectedResult {
		t.Errorf("Mismatch: %d != %d", ans, expectedResult)
	}
}

func TestHierarchicalExecution(t *testing.T) {
	t.Run("sameGradesOps", TestHierarchicalExecutionSameGradeOps)
	t.Run("variousGradesOps", TestHierarchicalExecutionVariousGradesOps)
	t.Run("variousGradesOps2", TestHierarchicalExecutionVariousGradesOps2)
	t.Run("variousGradesOps3", TestHierarchicalExecutionVariousGradesOps3)
}

func TestOperation(t *testing.T) {
	processableExpressions := [8]string{"3!+5",
		"1*2*3",
		"1+2+3",
		"2^2+6/3",
		"4-2*2/2",
		"3^2+2",
		"3+3",
		"4!+11*6"}

	expectedResult := [8]int{11, 6, 6, 6, 2, 11, 6, 90}

	for i, op := range processableExpressions {
		t.Run(fmt.Sprintf("operationTest_%d", i), func(t *testing.T) {
			ans, err := operation(op)

			if err != nil {
				t.Errorf("Error: %s", err.Error())
			}
			if ans != expectedResult[i] {
				t.Errorf("Mismatch: %d != %d", ans, expectedResult[i])
			}
		})
	}
}

func TestExpressionProcessing(t *testing.T) {
	expressions := [5]string{"(1+2)+3*(4-2)",
		"(4/(6/3))/2",
		"(4!+(3^2+2)*(3+3))",
		"(6-5)*4/(3-1)",
		"10*2+(3^(4-2))"}

	expectedResult := [5]string{"9", "1", "123", "2", "29"}

	for i, op := range expressions {
		t.Run(fmt.Sprintf("expressionProcessingTest_%d", i), func(t *testing.T) {
			ans, err := ExpressionProcessing(op)

			if err != nil {
				t.Errorf("Error: %s", err.Error())
			}
			if ans != expectedResult[i] {
				t.Errorf("Mismatch: %s != %s", ans, expectedResult[i])
			}
		})
	}
}
