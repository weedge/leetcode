package main

import "strconv"

func evalRPN(tokens []string) int {
	var s []int
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			n1 := s[len(s)-1]
			s = s[:len(s)-1]
			n2 := s[len(s)-1]
			s = s[:len(s)-1]
			switch v {
			case "+":
				s = append(s, n1+n2)
			case "-":
				s = append(s, n2-n1)
			case "*":
				s = append(s, n2*n1)
			case "/":
				s = append(s, n2/n1)
			}
		default:
			i, _ := strconv.Atoi(v)
			s = append(s, i)
		}
	}
	res := s[len(s)-1]
	return res
}
