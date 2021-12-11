package main

import (
	"regex/dfa"
)

func main() {
	var s = "(((5:aas)aa)$)"
	//var str = "((a)$)"
	//var s = "(a|b+.(7:fg)\\7b{1,2}j)"
	//fmt.Println(tree.ReplaceRepeat(tree.AddConcatenations(s)))
	//dfa.Minimization(d)
	dfa.Compile(s)
	//var d2 = dfa.Compile(str)
	//var m = dfa.GetTransitions(d.InitialState, nil)
	//fmt.Println(dfa.CreateRE(dfa.CopyDFA(d), d.InitialState, m))
	//fmt.Println(dfa.Difference(d,d2))
	//	dfa.Search(s, "fepfk")
}

