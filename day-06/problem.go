package main

type Problem struct {
	Values   []int
	Operator string
}

func NewProblem(values []int, operator string) Problem {
	return Problem{Values: values, Operator: operator}
}

func Solve(problem Problem) int {
	switch problem.Operator {
	case "+":
		return Sum(problem.Values)
	case "*":
		return Product(problem.Values)
	default:
		panic("invalid operator: " + problem.Operator)
	}
}

func Sum(values []int) (sum int) {
	sum = 0
	for _, value := range values {
		sum += value
	}
	return
}

func Product(values []int) (product int) {
	product = 1
	for _, value := range values {
		product *= value
	}
	return
}
