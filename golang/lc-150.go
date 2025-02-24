package main

import "strconv"

func evalRPN(tokens []string) int {
	sta := []int{}
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	for i := 0; i < len(tokens); i++ {
		if operators[tokens[i]] == true {
			num1, num2 := sta[len(sta)-1], sta[len(sta)-2]
			sta = sta[:len(sta)-1]
			sta = sta[:len(sta)-1]
			tmp := 0
			switch tokens[i] {
			case "+":
				tmp = num1 + num2
			case "-":
				tmp = num2 - num1
			case "*":
				tmp = num2 * num1
			case "/":
				tmp = num2 / num1
			}
			sta = append(sta, tmp)
		} else {
			num, _ := strconv.Atoi(tokens[i])
			sta = append(sta, num)
		}
	}
	return sta[0]
}
