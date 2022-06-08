package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isOp(s byte) bool {
	return s == '#' || s == '$' || s == '@'
}

func Error(s string) {
	fmt.Println(s)
	os.Exit(0)
}

func Reverse(str string) string {
	if str != "" {
		return Reverse(str[1:]) + str[:1]
	}
	return ""
}

var Ans = map[string]bool{}

func findEnd(s string, pos int) int {
	var (
		ans     int
		balance = 1
	)
	for i := pos + 1; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		}
		if s[i] == ')' {
			balance--
		}
		if balance == 0 {
			ans = i + 1
			break
		}
	}
	return ans
}

func tokenize(s string) []string {
	pos := 1
	tokens := make([]string, 0, 0)
	if len(s) == 1 {
		return tokens
	}
	for s[pos] != ')' && pos < len(s) {
		if s[pos] == ' ' {
			pos++
		}
		if isOp(s[pos]) || unicode.IsLetter(rune(s[pos])) {
			tokens = append(tokens, s[pos:pos+1])
			pos++
		} else {
			//fmt.Println(pos)
			pos2 := findEnd(s, pos)
			tokens = append(tokens, s[pos:pos2])
			pos = pos2
		}
	}
	return tokens
}

func to_reverse_polish(s string) string {
	s = Reverse(s)
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if r[i] == ')' {
			r[i] = '('
		} else if r[i] == '(' {
			r[i] = ')'
		}
	}
	s = string(r)
	return s
}

func calculate(tokens []string) int {
	var s []int
	for i := range tokens {
		if len(tokens[i]) > 1 {
			t := parse(tokens[i])
			Ans[tokens[i]] = true
			s = append(s, t)
		} else if unicode.IsLetter(rune(tokens[i][0])) {
			s = append(s, int(tokens[i][0]-'0'))
		} else {
			if len(s) != 2 {
				Error("wrong expression")
			}
			operand1 := s[len(s)-1]
			s = s[:len(s)-1]
			operand2 := s[len(s)-1]
			s = s[:len(s)-1]
			if tokens[i][0] == '#' {
				s = append(s, operand1+operand2)
			}
			if tokens[i][0] == '$' {
				s = append(s, operand1-operand2)
			}
			if tokens[i][0] == '@' {
				s = append(s, operand1*operand2)
			}
		}
	}
	if len(s) != 1 {
		Error("wrong expression")
	}
	return s[0]
}

func parse(s string) int {
	tokens := tokenize(s)
	if len(tokens) != 0 {
		return calculate(tokens)
	}
	return 0
}

func main() {
	var s string
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	s = myscanner.Text()
	s = to_reverse_polish(s)
	if len(s) != 1 {
		Ans[s] = true
	}
	parse(s)
	fmt.Println(len(Ans))
}
