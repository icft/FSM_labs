package tree

import (
	"errors"
	"fmt"
	"unicode"
)

var ops = []string{"(", ")", "\\", "{", "}", "|", "+", "#", "."}
var id = 1

func Find(s string) bool {
	for i:=0; i < len(ops); i++ {
		if s == ops[i] {
			return true
		}
	}
	return false
}

func AddConcatenations(regex string) (NewRegex string) {
	NewRegex = ""
	var i = 0
	for i < len(regex)-1 {
		//ab->a.b
		//a(->a.(
		//a/->a./
		//)a->).a
		//)(->).(
		//}a->}.a
		//+a->+.a
		//+(->+.(
		//((->(.(
		//))->).)
		//)#->).#
		var Flag = Find(string(regex[i]))
		var Check = Find(string(regex[i+1]))
		if !Flag {
			if Check {
				if regex[i+1] == '(' {
					NewRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '\\' {
					NewRegex += string(regex[i]) + "."
					i++
					NewRegex += string(regex[i])
					i++
					for unicode.IsDigit(rune(regex[i])) {
						NewRegex += string(regex[i])
						i ++
					}
					NewRegex += "."
				} else if regex[i+1] == '{' {
					for regex[i] != '}' {
						NewRegex += string(regex[i])
						i++
					}
					if regex[i+1] != ')' {
						NewRegex += string(regex[i]) + "."
					} else {
						NewRegex += string(regex[i])
					}
					i++
				} else if regex[i+1] == '#' {
					NewRegex += string(regex[i]) + "."
					i++
				} else {
					NewRegex += string(regex[i])
					i++
				}
			} else {
				NewRegex += string(regex[i]) + "."
				i++
			}
		} else {
			if regex[i] == ')' {
				if regex[i+1] == '(' || !Find(string(regex[i+1])) {
					NewRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '#' || regex[i+1] == '\\' {
					NewRegex += string(regex[i]) + "."
					i++
				} else {
					NewRegex += string(regex[i])
					i++
				}
			} else if regex[i] == '+' {
				NewRegex += string(regex[i]) + "."
				i++
			} else if regex[i] == '#' {
				NewRegex += string(regex[i]) + string(regex[i+1]) + "."
				i += 2
			} else if regex[i] == '}' {
				if regex[i+1] != ')' {
					NewRegex += string(regex[i]) + "."
				} else {
					NewRegex += string(regex[i])
				}
			} else  {
				NewRegex += string(regex[i])
				i++
			}
		}
	}
	NewRegex += string(regex[len(regex)-1])
	var CopyNewRegex string
	var minus = 0
	for _, c := range NewRegex {
		CopyNewRegex += string(c)
	}
	var Bracket = false
	var IndBracket = -1
	for ind, char := range CopyNewRegex {
		if char == '(' {
			Bracket = true
			IndBracket = ind
		}
		if char == ')' {
			Bracket = false
			IndBracket = -1
		}
		if char == ':' && Bracket && (ind - IndBracket) > 0 {
			for j := IndBracket+-minus; j < ind-1-minus; j++ {
				if NewRegex[j] == '.' {
					NewRegex = NewRegex[:j]+NewRegex[j+1:]
					minus++
				}
			}
			NewRegex = NewRegex[:ind-minus-1]+":"+NewRegex[ind-minus+2:]
			minus += 2
		}
	}
	return
}

func CreateTokens(Regex string) (tokens []string) {
	var c = ""
	var FigBr = false
	var Shield = false
	for _, r := range Regex {
		if r == '(' || r == ')' && !FigBr {
			tokens = append(tokens, string(r))
		} else if (r == '.' || r == '|' || r == '+') && !FigBr && !Shield {
			tokens = append(tokens, string(r))
		} else if r == '{' {
			FigBr = true
			c += string(r)
		} else if r == '}' {
			FigBr = false
			c += string(r)
			tokens = append(tokens, c)
			c = ""
		} else if FigBr {
			c += string(r)
		} else if r == '\\' {
			c += string(r)
			Shield = true
		} else if Shield {
			if r == '.' {
				tokens = append(tokens, c)
				tokens = append(tokens, string(r))
				c = ""
				Shield = false
			} else {
				c += string(r)
			}
		} else if  !FigBr {
			tokens = append(tokens, string(r))
		}
	}
	var minus = 0
	var Bracket = false
	var indBracket = -1
	var copyTokens []string
	for _, ind := range tokens {
		copyTokens = append(copyTokens, ind)
	}
	for ind, char := range copyTokens {
		if char == "(" {
			Bracket = true
			indBracket = ind
		}
		if char == ")" {
			Bracket = false
			indBracket = -1
		}
		if char == ":" && Bracket && (ind - indBracket) > 0 {
			var s = ""
			var Ind = ind-minus
			var IndBracket = indBracket-minus
			var shift = 0-minus
			for j := IndBracket+1; j < Ind; j++ {
				s += tokens[j]
				minus++
			}
			minus--
			copy(tokens[IndBracket+2:], tokens[Ind:])
			tokens[IndBracket+1] = s
			tokens = tokens[:len(tokens)-(shift+minus)]
		}
	}
	return
}

func isDigit(str string) bool {
	for i := range str {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}

func CreateNodes(tokens []string) (nodes []*Node) {
	var i = 0
	for i < len(tokens) {
		var tok = tokens[i]
		if tok == "(" || tok == ")" {
			nodes = append(nodes, &Node{Type: Bracket, Val: tok})
			i++
		} else if len(tok) == 1 && !Find(tok) && tokens[i+1] != ":" {
			if tok != "^" {
				nodes = append(nodes, &Node{
					Id: 	  id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tok,
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: true,
				})
				id++
			} else {
				nodes = append(nodes, &Node{
					Id: 	  id,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     LeafNode,
					Val:      tok,
					FirstPos: []int{id},
					LastPos:  []int{id},
					Nullable: false,
				})
				id++
			}
			i++
		} else if len(tok) == 1 && Find(tok) && tokens[i+1] != ":" {
			if tok == "." {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Concat,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
			} else if tok == "|" {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Or,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
			} else if tok == "+" {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Plus,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
			} else if tok == "#" {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Sharp,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
			}
			i++
		} else {
			if tok[0] == '\\' {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Reference,
					Val:      tok[1:],
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if tok[0] == '{' {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Repeat,
					Val:      tok[1: len(tok)-1],
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i++
			} else if isDigit(tok) {
				nodes = append(nodes, &Node{
					Id: 	  -1,
					Parent:   nil,
					Left:     nil,
					Right:    nil,
					Type:     Group,
					Val:      tok,
					FirstPos: nil,
					LastPos:  nil,
					Nullable: false,
				})
				i += 2
			}
		}
	}
	return nodes
}

func ClosestBrackets(tokens []*Node) (first int, second int) {
	first = 0
	second = len(tokens)-1
	var currF = 0
	var currS = 0
	var priority = true
	for i, tok := range tokens {
		//fmt.Println(tok)
		if tok.Val == "(" && tokens[i+2].Val != ":" {
			currF = i
			priority = true
		} else if tok.Val == ")" {
			if priority {
				currS = i
				if currS-currF < second-first {
					second = currS
					first = currF
				}
				priority = false
			}
		}
	}
	return
}

func FindMerge(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Merge(first []int, second []int) []int {
	var res []int
	res = append(res, first...)
	for _, c := range second {
		if FindMerge(first, c) {
			res = append(res, c)
		}
	}
	return res
}

var m = make(map[string]*Node)

func CreateSubtree(nodes []*Node, first int, second int) ([]*Node, error) {
	var currNode *Node
	var i int
	_, ok := m[nodes[first+1].Val]
	var err error = nil
	/*fmt.Println(err)
	if !err {
		fmt.Println(m[nodes[first+1].Val])
	} else {
		fmt.Println(err)
	}*/
	//fmt.Println(nodes[first+1].Type)
	if nodes[first+1].Type == Group && !ok {
		i = first+2
		second--
		fmt.Println(second)
		for i <= second {
			if nodes[i].Type == Sharp {
				nodes[i].Left = nodes[i+1]
				nodes[i].Right = nil
				nodes[i].Nullable = nodes[i+1].Nullable
				nodes[i].FirstPos = nodes[i+1].FirstPos
				nodes[i].LastPos = nodes[i+1].LastPos
				nodes[i+1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i+1:], nodes[i+2:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i++
			}
			i++
		}
		i = first + 2
		for i <= second {
			if nodes[i].Type == Plus {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i].Nullable = true
				nodes[i].FirstPos = nodes[i-1].FirstPos
				nodes[i].LastPos = nodes[i-1].LastPos
				nodes[i-1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i--
			}
			i++
		}
		i = first + 2
		fmt.Println(second)
		for i <= second {
			if nodes[i].Type == Repeat {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i].Nullable = nodes[i-1].Nullable
				nodes[i].FirstPos = nodes[i-1].FirstPos
				nodes[i].LastPos = nodes[i-1].LastPos
				nodes[i-1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i--
			}
			i++
		}
		i = first+3
		for i <= second-1 {
			if nodes[i].Type == Concat && nodes[i].Left == nil && nodes[i].Right == nil {
				fmt.Println(i, second-1)
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nodes[i+1]
				nodes[i].Nullable = nodes[i-1].Nullable && nodes[i+1].Nullable
				if nodes[i-1].Nullable {
					nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
				} else {
					nodes[i].FirstPos = nodes[i-1].FirstPos
				}
				if nodes[i+1].Nullable {
					nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
				} else {
					nodes[i].LastPos = nodes[i+1].LastPos
				}
				nodes[i-1].Parent = nodes[i]
				nodes[i+1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i:], nodes[i+2:])
				nodes[i-1] = currNode
				nodes = nodes[:len(nodes)-2]
				Print(nodes)
				i -=2
				second = second-2
			}
			i++
		}
		i = first+3
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Or && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i].Nullable = nodes[i-1].Nullable || nodes[i+1].Nullable
					nodes[i].FirstPos = Merge(nodes[i-1].FirstPos, nodes[i+1].FirstPos)
					nodes[i].LastPos = Merge(nodes[i-1].LastPos, nodes[i+1].LastPos)
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					Print(nodes)
					i -=2
					second = second-2
				}
				i++
			}
			i++
		}
		nodes[first+1].Left = currNode
		nodes[first+1].Right = nil
		currNode.Parent = nodes[first+1]
		nodes[first] = nodes[first+1]
		copy(nodes[first+1:], nodes[second+2:])
		//Print(nodes)
		//Print(nodes)
		nodes = nodes[:len(nodes)-3]
		m[currNode.Val] = currNode.Left
		Print(nodes)
		//fmt.Println(1)
	} else if second-first != 2 {
		i = first+1
		second -= 1
		for i < second {
			if nodes[i].Type == Reference {
				var key = nodes[i].Val
				nodes[i].Left, ok = m[key]
				if !ok {
					err = errors.New("group number error")
					return nodes, err
				}

			}
		}
		i = first+1
		for i <= second {
			if nodes[i].Type == Sharp {
				nodes[i].Left = nodes[i+1]
				nodes[i].Right = nil
				nodes[i+1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i+1:], nodes[i+2:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i++
			}
			i++
		}
		i = first + 1
		for i <= second {
			if nodes[i].Type == Plus {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i-1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i--
			}
			i++
		}
		i = first + 1
		for i <= second {
			if nodes[i].Type == Repeat {
				nodes[i].Left = nodes[i-1]
				nodes[i].Right = nil
				nodes[i-1].Parent = nodes[i]
				currNode = nodes[i]
				copy(nodes[i-1:], nodes[i:])
				nodes = nodes[:len(nodes)-1]
				Print(nodes)
				second--
				i--
			}
			i++
		}
		i = first+1
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Concat && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					Print(nodes)
					i -=2
					second = second-2
				}
				i++
			}
			i++
		}
		i = first+1
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Or && nodes[i].Left == nil && nodes[i].Right == nil {
					nodes[i].Left = nodes[i-1]
					nodes[i].Right = nodes[i+1]
					nodes[i-1].Parent = nodes[i]
					nodes[i+1].Parent = nodes[i]
					currNode = nodes[i]
					copy(nodes[i:], nodes[i+2:])
					nodes[i-1] = currNode
					nodes = nodes[:len(nodes)-2]
					Print(nodes)
					i -=2
					second = second-2
				}
				i++
			}
			i++
		}
		//fmt.Println(first, second)
		copy(nodes[first+1:], nodes[second+2:])
		nodes[first] = currNode
		nodes = nodes[:len(nodes)-2]
		Print(nodes)
	} else {
		currNode = nodes[first+1]
		copy(nodes[first+1:], nodes[second+1:])
		nodes[first] = currNode
		nodes = nodes[:len(nodes)-2]
	}
	return nodes, err
}

func CreateTree(regex string) *Node {
	var nodes = CreateNodes(CreateTokens(AddConcatenations(regex)))
	var first, second = ClosestBrackets(nodes)
	Print(nodes)
	var err error
	//fmt.Println(first, second)
	for second-first > 1 {
		//fmt.Println(first,second)
		nodes, err = CreateSubtree(nodes, first, second)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		first, second = ClosestBrackets(nodes)
	}
	return nodes[0]
}

func Print(nodes []*Node) {
	for i := 0; i < len(nodes); i++ {
		//fmt.Printf("\n%d\n", i)
		if nodes[i].Type == 1 {
			fmt.Printf("concat ")
		}
		if nodes[i].Type == 2 {
			fmt.Printf("or ")
		}
		if nodes[i].Type == 3 {
			fmt.Printf("plus ")
		}
		if nodes[i].Type == 4 {
			fmt.Printf("sharp ")
		}
		if nodes[i].Type == 5 {
			fmt.Printf("group ")
		}
		if nodes[i].Type == 6 {
			fmt.Printf("shield ")
		}
		if nodes[i].Type == 7 {
			fmt.Printf("repeat ")
		}
		if nodes[i].Type == 8 {
			fmt.Printf("leaf ")
		}
		if nodes[i].Type == 9 {
			fmt.Printf("bracket ")
		}
	}
	fmt.Println()
}