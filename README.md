ε-NFA: Epsilon-Nondeterministic finite automaton
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/e-nfa/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/e-nfa?status.svg)](https://godoc.org/github.com/kkdai/nfa)  [![Build Status](https://travis-ci.org/kkdai/e-nfa.svg?branch=master)](https://travis-ci.org/kkdai/e-nfa)



![image](https://upload.wikimedia.org/wikipedia/commons/thumb/0/0e/NFAexample.svg/250px-NFAexample.svg.png)



What is Nondeterministic finite automaton
=============

`ε-NFA`: Epsilon-Nondeterministic finite automaton (so call:Nondeterministic finite automaton with ε-moves)


 (sited from [here](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton))


Looking for DFA implement?
=============

I also write a DFA implenent in Go here. [https://github.com/kkdai/dfa](https://github.com/kkdai/dfa)

Looking for NFA implement?
=============

I also write a NFA implenent in Go here. [https://github.com/kkdai/dfa](https://github.com/kkdai/nfa)


Installation and Usage
=============


Install
---------------

```go

go get github.com/kkdai/e-nfa

```

Usage
---------------

```go

package main

import (
    "github.com/kkdai/enfa"
    "fmt"
)

func main() {
	nfa := NewNFA(0, false)
	nfa.AddState(1, false)
	nfa.AddState(2, true)

	nfa.AddTransition(0, "a", 1, 2)
	nfa.AddTransition(1, "b", 0, 2)

	var inputs []string
	inputs = append(inputs, "a")
	inputs = append(inputs, "b")
	fmt.Println("If input a, b will go to final?", nfa.VerifyInputs(inputs) )
}

```

Inspired By
=============

- [NFA: Wiki](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton)
- [Coursera: Automata](https://class.coursera.org/automata-004/)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
