package AST

var ops = []rune{'(', ')', '\\', '{', '}', '|', '+'}

func Find(s rune, ops []rune) bool {
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
		//	fmt.Println(string(regex[i]), " ", i)
		var Flag = Find(rune(regex[i]), ops)
		var Check = Find(rune(regex[i+1]), ops)
		if !Flag {
			if Check {
				if regex[i+1] == '(' {
					NewRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '\\' {
					NewRegex += string(regex[i]) + "."
					NewRegex += string(regex[i+1]) + string(regex[i+2]) + "."
					i += 3
				} else if regex[i+1] == '{' {
					for regex[i] != '}' {
						NewRegex += string(regex[i])
						i++
					}
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
				if  regex[i+1] == '(' || !Find(rune(regex[i+1]), ops) {
					NewRegex += string(regex[i]) + "."
					i++
				} else {
					NewRegex += string(regex[i])
					i++
				}
			} else if regex[i] == '+' {
				NewRegex += string(regex[i]) + "."
				i++
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
	//fmt.Println(CopyNewRegex)
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
			//fmt.Println(NewRegex)
			for j := IndBracket+-minus; j < ind-1-minus; j++ {
				if NewRegex[j] == '.' {
					//fmt.Println(j)
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
			//fmt.Println(Shield)
		} else if Shield {
			if r == '.' {
				//fmt.Println(11111111111111)
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
	//fmt.Println("Исходное:", tokens)
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
			//fmt.Println(tokens)
			//fmt.Println(minus)
			//fmt.Println(IndBracket, ind)
			var s = ""
			//fmt.Println(tokens)
			//fmt.Println(indBracket-minus, ind-minus, minus)

			var Ind = ind-minus
			var IndBracket = indBracket-minus
			var shift = 0-minus
			for j := IndBracket+1; j < Ind; j++ {
				s += tokens[j]
				minus++
			}
			//fmt.Println(s)
			minus--
			/*fmt.Println(shift, minus)
			fmt.Println(len(tokens))*/
			copy(tokens[IndBracket+2:], tokens[Ind:])
			/*fmt.Println(len(tokens[IndBracket+2:]), len(tokens[Ind:]))
			fmt.Println(len(tokens))
			*/tokens[IndBracket+1] = s
			tokens = tokens[:len(tokens)-(shift+minus)]
		}
	}
	return
}

func CreateRPN(tokens []string) (tokensRPN []string) {
	var stack []string
	for _, tok := range tokens {
		if tok == "(" {
			stack = append(stack, tok)
		} else if tok == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				tokensRPN = append(tokensRPN, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)]
		} else if tok == "+" {
			stack = append(stack, tok)
		} else if tok == "." {
			for len(stack) > 0 && stack[len(stack)-1] != "+" {
				tokensRPN = append(tokensRPN, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tok)
		} else if tok == "|" {
			for len(stack) > 0 && (stack[len(stack)-1] != "+" || stack[len(stack)-1] == ".") {
				tokensRPN = append(tokensRPN, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tok)
		} else {
			tokensRPN = append(tokensRPN, tok)
		}
	}
	for len(stack) > 0 {
		tokens = append(tokensRPN, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return
}
