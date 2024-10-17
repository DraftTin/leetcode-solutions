package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	res := []int{}
	for i := 0; i < len(expression); i++ {
		oper := expression[i]
		if oper == '+' || oper == '-' || oper == '*' {
			s1 := diffWaysToCompute(expression[:i])
			s2 := diffWaysToCompute(expression[i+1:])
			for _, n1 := range s1 {
				for _, n2 := range s2 {
					if oper == '+' {
						res = append(res, n1+n2)
					} else if oper == '-' {
						res = append(res, n1-n2)
					} else {
						res = append(res, n1*n2)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		n, _ := strconv.Atoi(expression)
		res = append(res, n)
	}
	return res
}
