package main

import (
	"fmt" // allows user input
	"os"  // exits programm
)// imports

type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}// state struct

type nfa struct{
	initial *state
	accept  *state
}// nfa struct



func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}// array of pointers

	for _, r :=  range pofix {
		switch r {
		case '.'://if the character = '.'
		    //pops the last nfa off the stack
			frag2 := nfastack[len(nfastack)-1]
			// sets stack to accept all but last element
			nfastack = nfastack[:len(nfastack)-1]
			//nfa stack is updated
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// frag1 points to frag2
			frag1.accept.edge1 = frag2.initial
			//push a new fragment to the stack,
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
			// case '.'
		case '|'://if the character = '|'
			//pops the last nfa off the stack
			frag2 := nfastack[len(nfastack)-1]
			// sets stack to accept all but last element
			nfastack = nfastack[:len(nfastack)-1]
			//nfa stack is updated
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			//state is now empty
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

			// case '|'
		case '*'://if the character = '*'
			//pops the last nfa off the stack
			frag := nfastack[len(nfastack)-1]
			// sets stack to accept all but last element
			nfastack = nfastack[:len(nfastack)-1]
			//state is now empty
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			// case '*'

		case '+'://if the character = '+'
			//pops the last nfa off the stack
			frag := nfastack[len(nfastack)-1]
			//state is now empty
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = &initial

			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

			// case '+'
		default:
			//empty state for accept
			accept := state{}
			initial := state{symbol: r ,edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}/// switch
	}// for 

	

	return nfastack[0]
}// poregtonfa

func addstate(l []*state, s *state, a *state) []*state{
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addstate(l, s.edge1, a)
		if s.edge2 != nil{
			l = addstate(l, s.edge2, a)
		}
	}// if
	return l
}

func pomatch(po string, s string) bool{
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addstate(current[:], ponfa.initial, ponfa.accept)


	for _, r := range s{
		for _, c := range current{
			if c.symbol == r {
				next = addstate(next[:], c.edge1, ponfa.accept)
			}// if
		}// inner for
		current, next = next, []*state{}
	}// for
	
	for _, c := range current {
		if c == ponfa.accept{
			ismatch = true
			break
		}// if
	}// for


	return ismatch
}// pomatch
// intopost method for shunting taken from shunt.go
func intopost(infix string) string{
	specials := map[rune]int{'*': 10,'.': 9,'|': 8}

	pofix , s:= []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
		s = append(s, r)
		//case (
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}// for
		s = s[:len(s)-1]
		//case )
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}// for
			s = append(s,r)
		//case specials[r]
		default:
			pofix = append(pofix,r)
		}// switch

	}// for
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}// for

	return string(pofix)
}// intopost


func runProg(){
	var condition string
	var test string
	//prompt for user input
	fmt.Println("Please enter the condition")
	
	fmt.Scanln(&condition)

	for _, r := range condition {
		if r == '(' || r == ')'{
			condition = intopost(condition)
			break
		}// if

	}// for
	fmt.Println("Please enter the test String")
	
	fmt.Scanln(&test)
	fmt.Print("\n\n")

	//Our test String
	fmt.Printf("The string %s condition is : %t\n\n", test, pomatch(condition, test))
	fmt.Println()
}// runProg
func main(){
	
	for {
		var input string
		//Menu options
		fmt.Println("Please choose an option")
		fmt.Println("1: Test a Condition")
		fmt.Println("2: Exit")

		fmt.Print("Option: ")

		fmt.Scanln(&input)

		switch input {
		case "1":
			//Run the test
			runProg()
		case "2":
			//Exits the program
			os.Exit(3)
		default:
			fmt.Println("Invalid input")
		}

	}
}