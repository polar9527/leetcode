package offer

import "fmt"

type state map[byte]int

func isNumber(s string) bool {
	// state{' ': -1, 'd': -1, 's': -1, '.': -1, 'e': -1},
	p0 := state{' ': 0, 'd': 2, 's': 1, '.': 4, 'e': -1}
	p1 := state{' ': -1, 'd': 2, 's': -1, '.': 4, 'e': -1}
	p2 := state{' ': 8, 'd': 2, 's': -1, '.': 3, 'e': 5}
	p3 := state{' ': 8, 'd': 3, 's': -1, '.': -1, 'e': 5}
	p4 := state{' ': -1, 'd': 3, 's': -1, '.': -1, 'e': -1}
	p5 := state{' ': -1, 'd': 7, 's': 6, '.': -1, 'e': -1}
	p6 := state{' ': -1, 'd': 7, 's': -1, '.': -1, 'e': -1}
	p7 := state{' ': 8, 'd': 7, 's': -1, '.': -1, 'e': -1}
	p8 := state{' ': 8, 'd': -1, 's': -1, '.': -1, 'e': -1}

	fsm := []state{p0, p1, p2, p3, p4, p5, p6, p7, p8}
	p := 0
	var t byte
	for i := 0; i < len(s); i++ {
		switch true {
		case s[i] == ' ':
			t = ' '
		case s[i] >= '0' && s[i] <= '9':
			t = 'd'
		case s[i] == '-' || s[i] == '+':
			t = 's'
		case s[i] == '.':
			t = '.'
		case s[i] == 'e' || s[i] == 'E':
			t = 'e'
		default:
			t = '?'
		}
		fmt.Println("i : ", i)
		fmt.Println("s[i] : ", string(s[i]))
		fmt.Println("t : ", string(t))
		fmt.Println("fsm[p][t] : ", fsm[p][t])
		nextP, ok := fsm[p][t]
		if !ok {
			return false
		}
		if fsm[p][t] == -1 {
			return false
		}
		p = nextP
	}
	// p in [2,3,7,8]
	switch p {
	case 2, 3, 7, 8:
		return true
	default:
		return false
	}
}
