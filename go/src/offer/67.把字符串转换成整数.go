package main

import (
	"fmt"
	"math"
	"strconv"
)

type Automaton struct {
	state string
	ans   int
	sign  int
	table map[string][]string
}

func newAutomaton() *Automaton {
	return &Automaton{
		state: "start",
		ans:   0,
		sign:  1,
		table: map[string][]string{
			"start":  {"start", "sign", "number", "end"},
			"sign":   {"end", "end", "number", "end"},
			"number": {"end", "end", "number", "end"},
			"end":    {"end", "end", "end", "end"},
		},
	}
}

func (a *Automaton) getNextState(in string) int {
	switch in {
	case " ":
		return 0
	case "+", "-":
		return 1
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return 2
	default:
		return 3
	}
}

func (a *Automaton) transfer(in string) {
	a.state = a.table[a.state][a.getNextState(in)]
	if a.state == "number" {
		// 	computer
		num, _ := strconv.Atoi(in)
		a.ans = a.ans*10 + num
		if a.sign == 1 && a.ans > math.MaxInt32 {
			a.ans = math.MaxInt32
		} else if a.sign == -1 && -a.ans < math.MinInt32 {
			a.ans = -math.MaxInt32
		}
	} else if a.state == "sign" {
		if in == "-" {
			a.sign = -1
		}
	}
}

func (a *Automaton) getState() string {
	return a.state
}

func (a *Automaton) getAns() int {
	return a.ans
}

func (a *Automaton) getSign() int {
	return a.sign
}

func strToInt(str string) int {
	auto := newAutomaton()
	for i := range str {
		auto.transfer(string(str[i]))
		if auto.getState() == "end" {
			break
		}
	}
	return auto.getAns() * auto.getSign()
}

func main() {
	n := strToInt("   -42")
	fmt.Println(n)
}
