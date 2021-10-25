
package AST

import "unicode"

var ops = []rune{'(', ')', '\\', '{', '}', '|', '+', '#'}

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
		//)#->).#
		var Flag = Find(rune(regex[i]), ops)
		var Check = Find(rune(regex[i+1]), ops)
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
					NewRegex += string(regex[i]) + "."
					i++
				} else if regex[i+1] == '#' {
					NewRegex += string(regex[i]) + "." + string(regex[i+1]) + string(regex[i+2]) + "."
					i += 3
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

func Pop(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

/*func CreateRPN(tokens []string) (tokensRPN []string) {

}*/
