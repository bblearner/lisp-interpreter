package main

func evaluate(n node) interface{} {

	stack := make([]interface{}, 0)

	for _, child := range n.childNodes {
		stack = append(stack, evaluate(*child))
	}

	switch n.val {
	case "+", "-", "*", "/", "%":
		if len(stack) == 2 {
			return arithmeticOperation(stack[0].(int), stack[1].(int), n.val)
		}
	case "<", ">", "<=", ">=", "==", "!=":
		if len(stack) == 2 {
			return comparisionOperation(stack[0].(int), stack[1].(int), n.val)
		}
	case "format":
		// Handle format
	case "let":
		// Handle let
	case "defun":
		// Handle defun
	case "defmacro":
		// Handle defmacro
	case "if":
		// Handle if
	case "string-upcase":
		// Handle string-upcase
	default:

	}

	return nil
}

func arithmeticOperation(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		panic("unsupported operation")
	}
}

func comparisionOperation(a, b int, op string) bool {
	switch op {
	case "<":
		return a < b
	case ">":
		return a > b
	case "<=":
		return a <= b
	case ">=":
		return a >= b
	case "==":
		return a == b
	case "!=":
		return a != b
	default:
		panic("unsupported comparison operation")
	}
}
