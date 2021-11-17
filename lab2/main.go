package main

import (
	"fmt"
	"lab2/dfa"
	"lab2/tree"
)

func main() {
	var s = "(((a|b)abb)$)"
	//var s = "(a|b+.(7:fg)\\7b{1,2}j)"
	var str = tree.AddConcatenations(s)
	fmt.Println(str)
	var d = dfa.Compile(s)
	dfa.Minimization(d)
//	dfa.Search(s, "fepfk")
	/*var regex = tree.ReplaceRepeat(str)
	fmt.Println(regex)
	var tokens = tree.CreateTokens(regex)
	fmt.Println(tokens)
	*///var root = tree.CreateTree(s)
	/*fmt.Printf("\n\n\n")
	tree.PrintTree(root)
	var a []int = nil
	var followpos [][]int
	for i := 0; i < len(tree.ID); i++ {
		followpos = append(followpos, a)
	}
	tree.FindFollowPos(root, followpos)
	//fmt.Println(tree.FollowPos)
	fmt.Printf("\n\n\n\n")*/
	//fmt.Println(root.FirstPos)
	//var avtomata = dfa.InitDFA(root, tree.FollowPos, root.FirstPos)
	//var state = dfa.Convert(avtomata, root)
	//fmt.Println("\n", "\n", state, "\n\n")
	//dfa.Print(avtomata)
}
