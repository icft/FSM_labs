package AST

import (
	"fmt"
	"unicode"
)

var ops = []string{"(", ")", "\\", "{", "}", "|", "+", "#", "."}
var numberGroups []int

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
					NewRegex += string(regex[i])
					i++
				} else if regex[i+1] == '#' {
					NewRegex += string(regex[i]) + "."
					i ++
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
				if  regex[i+1] == '(' || !Find(string(regex[i+1])) {
					NewRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '#' {
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
				NewRegex += string(regex[i]) + string(regex[i+1])+"."
				i+=2
			} else {
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
	var Sharp = false
	for _, r := range Regex {
		if r == '(' || r == ')' && !FigBr {
			tokens = append(tokens, string(r))
		} else if (r == '.' || r == '|' || r == '+') && !FigBr && !Shield && !Sharp {
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
		} else if r == '#' {
			Sharp = true
			c += string(r)
		} else if Sharp {
			c += string(r)
			Sharp = false
			tokens = append(tokens, c)
			c = ""
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
		} else if  !FigBr || !Sharp {
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
			nodes = append(nodes, &Node{Type:LeafNode, Val:tok})
			i++
		} else if len(tok) == 1 && Find(tok) && tokens[i+1] != ":" {
			if tok == "." {
				nodes = append(nodes, &Node{Type:Concat, Val: tok})
			} else if tok == "|" {
				nodes = append(nodes, &Node{Type:Or, Val: tok})
			} else if tok == "." {
				nodes = append(nodes, &Node{Type:Concat, Val: tok})
			} else if tok == "+" {
				nodes = append(nodes, &Node{Type:Plus, Val: tok})
			} else if tok == "." {
				nodes = append(nodes, &Node{Type:Concat, Val: tok})
			}
			i++
		} else {
			if tok[0] == '#' {
				nodes = append(nodes, &Node{Type: Sharp, Val: string(tok[1])})
				i++
			} else if tok[0] == '\\' {
				nodes = append(nodes, &Node{Type: Shield, Val: string(tok[1:])})
				i++
			} else if tok[0] == '{' {
				nodes = append(nodes, &Node{Type: Repeat, Val: tok[1 : len(tok)-1]})
				i++
			} else if isDigit(tok) {
				nodes = append(nodes, &Node{Type: Group, Val: tok})
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
	var isGroup = false
	var sum = 0
	var priority = true
	for i, tok := range tokens {
		if tok.Val == "(" && tokens[i+2].Val != ":" {
			currF = i
			priority = true
		} else if tok.Val == "(" && tokens[i+2].Val == ":" && priority {
			isGroup = true
			sum++
		} else if tok.Val == ")" {
			if priority && !isGroup {
				currS = i
				if currS-currF < second-first {
					second = currS
					first = currF
				}
				priority = false
			} else if priority && isGroup {
				sum--
				if sum <= 0 {
					isGroup = false
				}
			}
		}
	}
	return
}

func CreateSubtree(nodes []*Node, first int, second int) *Node {
	var currNode *Node
	if nodes[first+1].Type == Group {
		
	} else {
		var i = first+1
		second -= 1
		//fmt.Println(i, second)
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
		//fmt.Println(i, second)
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
		//fmt.Println(i, second)
		for i <= second-1 {
			if i >= first+2 {
				if nodes[i].Type == Concat && nodes[i].Left == nil && nodes[i].Right == nil {
					//fmt.Println(i)
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
					//fmt.Println(i)
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
		copy(nodes[first+1:], nodes[second+1:])
		nodes[first] = currNode
		nodes = nodes[:len(nodes)-2]
		Print(nodes)
	}
	return currNode
}

func Print(nodes []*Node) {
	for i := 0; i < len(nodes); i++ {
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
