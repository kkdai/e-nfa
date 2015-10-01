package enfa

import "testing"

func TestBasic(t *testing.T) {
	nfa := NewENFA(0, false)
	nfa.AddState(1, false)
	nfa.AddState(2, true)

	nfa.AddTransition(0, "a", 1)
	nfa.AddTransition(1, "b", 2)

	if ret := nfa.Input("a"); ret[0] != 1 {
		t.Errorf("Expect 1, but get %d\n", ret)
	}

	if ret := nfa.Input("b"); ret[0] != 2 {
		t.Errorf("Expect 2, but get %d\n", ret)
	}

	if !nfa.Verify() {
		t.Errorf("Verify is failed")
	}
}

func TestVerifyInputs(t *testing.T) {
	nfa := NewENFA(0, false)
	nfa.AddState(1, false)
	nfa.AddState(2, true)

	nfa.AddTransition(0, "a", 1)
	nfa.AddTransition(1, "b", 2)

	var inputs []string
	inputs = append(inputs, "a")
	inputs = append(inputs, "b")
	if !nfa.VerifyInputs(inputs) {
		t.Errorf("Verify Inputs is failed")
	}

	nfa.PrintTransitionTable()
}

func TestAdvancenfa(t *testing.T) {

	nfa := NewENFA(0, false)
	nfa.AddState(1, true)
	nfa.AddState(2, false)

	nfa.AddTransition(0, "0", 0)
	nfa.AddTransition(0, "1", 1)
	nfa.AddTransition(1, "0", 0)
	nfa.AddTransition(1, "1", 2)
	nfa.AddTransition(2, "0", 2)
	nfa.AddTransition(2, "1", 2)

	nfa.PrintTransitionTable()
	inputs := []string{"0", "0", "1", "0", "1"}

	if !nfa.VerifyInputs(inputs) {
		t.Errorf("Verify inputs is failed")
	}

	//Reset the nfa for another verification
	nfa.Reset()

	//Test go to dead state 2

	inputs2 := []string{"1", "1", "0", "0", "0"}

	if nfa.VerifyInputs(inputs2) {
		t.Errorf("Verify inputs is failed")
	}
}

func TestNFA(t *testing.T) {
	nfa := NewENFA(0, false)
	nfa.AddState(1, true)
	nfa.AddState(2, false)

	nfa.AddTransition(0, "0", 0, 1)
	nfa.AddTransition(0, "1", 1)
	nfa.AddTransition(1, "0", 0)
	nfa.AddTransition(1, "1", 2)
	nfa.AddTransition(2, "0", 2)
	nfa.AddTransition(2, "1", 2, 0)
	nfa.PrintTransitionTable()
	inputs := []string{"0", "0", "1", "0", "1"}

	if !nfa.VerifyInputs(inputs) {
		t.Errorf("Verify inputs is failed")
	}

	inputs2 := []string{"0", "0", "0", "0", "1"}

	if !nfa.VerifyInputs(inputs2) {
		t.Errorf("Verify inputs2 is failed")

	}

	inputs3 := []string{"0", "1", "2"}

	if nfa.VerifyInputs(inputs3) {
		t.Errorf("Verify inputs3 is failed")
	}
}

func TestEpsilonNFA(t *testing.T) {

	nfa := NewENFA(0, false)
	nfa.AddState(1, false)
	nfa.AddState(2, false)
	nfa.AddState(3, true)
	nfa.AddState(4, false)
	nfa.AddState(5, false)

	nfa.AddTransition(0, "1", 1)
	nfa.AddTransition(0, "0", 4)

	nfa.AddTransition(1, "1", 2)
	nfa.AddTransition(1, "", 3) //epsilon
	nfa.AddTransition(2, "1", 3)
	nfa.AddTransition(4, "0", 5)
	nfa.AddTransition(4, "", 1, 2) //E -> epsilon B C
	nfa.AddTransition(5, "0", 3)

	nfa.PrintTransitionTable()

	if !nfa.VerifyInputs([]string{"1"}) {
		t.Errorf("Verify inputs is failed")
	}

	nfa.Reset()

	if !nfa.VerifyInputs([]string{"1", "1", "1"}) {
		t.Errorf("Verify inputs is failed")
	}

	nfa.Reset()

	if !nfa.VerifyInputs([]string{"0", "1"}) {
		t.Errorf("Verify inputs is failed")
	}

	nfa.Reset()
	if !nfa.VerifyInputs([]string{"0", "0", "0"}) {
		t.Errorf("Verify inputs is failed")
	}
}
