package dfa

import (
	"fmt"
	"reflect"
	"regex/tree"
	"strings"
)

type State struct {
	Name        string
	StateNumber []int
	Dtran       map[string]*State
	num 		int
	Receive		bool
}

var num = 1

var DictNames = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

type DFA struct {
	tree               tree.Ast
	FollowPos          [][]int
	InitialStateNumber []int
	InitialState       *State
	NumberStates	   int
	Alphabet 		   []string
}

func InitDFA(root tree.Ast, FollowPos [][]int, FirstPos []int) *DFA {
	return &DFA{tree: root, FollowPos: FollowPos, InitialStateNumber: FirstPos, NumberStates: 0}
}

func FindSeen(finded *State, mas []*State) bool {
	for _, a := range mas {
		if finded.Name == a.Name &&
			reflect.DeepEqual(finded.Name, a.Name) &&
			reflect.DeepEqual(finded.Dtran, a.Dtran) {
			return true
		}
	}
	return false
}

func Find(dfa *DFA, s string) bool {
	for _, v := range dfa.Alphabet {
		if v == s {
			return true
		}
	}
	return false
}

func Convert(dfa *DFA, node *tree.Node, LeafNodes map[string][]int) *State {
	dfa.InitialState = &State{Name: DictNames[num], StateNumber: dfa.InitialStateNumber}
	dfa.NumberStates++
	num++
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for k, v := range LeafNodes {
				var i []int
				for _, c := range state.StateNumber {
					if tree.FindMerge(v, c) {
						i = append(i, c)
					}
				}
				//fmt.Println(i)
				if len(i) != 0 {
					var nextStateNumber []int
					for _, c := range i {
						nextStateNumber = tree.Merge(nextStateNumber, dfa.FollowPos[c-1])
					}
					var broken = false
					for _, seen := range seenStates {
						if reflect.DeepEqual(nextStateNumber, seen.StateNumber) {
							if state.Dtran == nil {
								state.Dtran = make(map[string]*State)
							}
							state.Dtran[k] = seen
							broken = true
							break
						}
					}
					if !broken {
						var a, b = num, 1
						if num > 26 {
							a = num % 26
							b += num / 26
						}
						var nextState = &State{Name: strings.Repeat(DictNames[a], b),
							StateNumber: nextStateNumber}
						dfa.NumberStates++
						num++
						if state.Dtran == nil {
							state.Dtran = make(map[string]*State)
						}
						if !Find(dfa, k) {
							dfa.Alphabet = append(dfa.Alphabet, k)
						}
						state.Dtran[k] = nextState
						leftStates = append(leftStates, nextState)
					}
				}
			}
		}
	}
	return dfa.InitialState
}

func SetReceive(start *State) {
	var leftStates = []*State{start}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			k, _ := GetKeysValues(state.Dtran)
			if len(state.Dtran) == 0 || (len(state.Dtran)==1 && state.Dtran[k[0]] == state) {
				state.Receive = true
			} else {
				state.Receive = false
			}
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
		}
	}
}

func Print(start *State) {
	var leftStates = []*State{start}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			fmt.Printf("<%s  %v %v>     ", state.Name, state.StateNumber, state.Receive)
			for k, v := range state.Dtran {
				fmt.Printf("Dtran<%s>=%s    ", k, v.Name)
				leftStates = append(leftStates, v)
			}
			fmt.Printf("\n")
		}
	}
}

func CopyDFA(dfa *DFA) *DFA {
	num=1
	var tmp = &DFA{tree: dfa.tree, FollowPos: dfa.FollowPos, InitialStateNumber: dfa.InitialStateNumber}
	Convert(tmp, dfa.tree.Root, dfa.tree.LeafNodes)
	return tmp
}

func FindExit(st *State) (s []*State) {
	var leftStates = []*State{st}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			if state.Dtran == nil {
				s = append(s, state)
			}
		}
	}
	return s
}


func Compile(regex string) (d *DFA) {
	var SyntaxTree tree.Ast
	regex = "(("+regex+")$)"
	SyntaxTree.ID = make(map[int]string)
	SyntaxTree.LeafNodes = make(map[string][]int)
	SyntaxTree.Root, SyntaxTree.FollowPos = tree.CreateTree(regex, SyntaxTree.ID, SyntaxTree.LeafNodes, SyntaxTree.FollowPos)
	d = InitDFA(SyntaxTree, SyntaxTree.FollowPos, SyntaxTree.Root.FirstPos)
	Convert(d, SyntaxTree.Root, SyntaxTree.LeafNodes)
	SetReceive(d.InitialState)
	return
}

var startState *State

func Search(pattern interface{}, str string) string {
	var dfa *DFA
	if reflect.TypeOf(pattern) == reflect.TypeOf("11") {
		dfa = Compile(pattern.(string))
	} else {
		dfa = pattern.(*DFA)
	}
	startState = dfa.InitialState
	return search(dfa.InitialState, str, 0, "")
}

func search(state *State, str string, ind int, copystr string) string {
	var finded bool
	var startInd = -1
	for ind < len(str) {
		finded = false
		for k, v := range state.Dtran {
			if string(str[ind]) == k {
				if startInd == -1 {
					startInd = ind
				}
				copystr += k
				finded = true
				state = v
				break
			}
		}
		ind++
		if state.Receive {
			return copystr
		}
		if !finded && state != startState {
			startInd = -1
			for _, v := range state.Dtran {
				return search(v, str, ind, copystr)
			}
		}
	}
	return ""
}

func SameStates(s1 *State, s2 *State) bool {
	var tr = false
	for k, v := range s1.Dtran {
		for k1,v1 := range s2.Dtran {
			if k == k1 {
				tr = true
				if v != v1 {
					return false
				}
			}
		}
		if !tr {
			return false
		}
	}
	return true
}


func MakeTransition(state *State, tr string) *State {
	for k, v := range state.Dtran {
		if k == tr {
			return v
		}
	}
	return nil
}

func Equal(mas1 []*State, mas2 []*State) bool {
	if len(mas1) != len(mas2) {
		return false
	}
	for i:=0; i < len(mas1); i++ {
		if mas1[i].Name == mas2[i].Name && reflect.DeepEqual(mas1[i].Dtran, mas2[i].Dtran) {
			continue
		} else {
			return false
		}
	}
	return true
}

func Minimization(dfa *DFA) {
	var states = ListOfStates(dfa)
	var exits []*State
	for _, v := range FindExit(dfa.InitialState) {
		exits = append(exits, v)
	}
	for i, v := range states {
		for _, val := range exits {
			if v == val {
				copy(states[i:], states[i+1:])
				states[len(states)-1] = nil
				states = states[:len(states)-1]
			}
		}
	}
	var split, new_split [][]*State
	split = append(split, states)
	split = append(split, exits)
	var table = make(map[*State][]*State)
	for _, i := range split {
		for _, j := range i {
			table[j] = i
		}
	}
	for len(new_split) != len(split) {
		var new_table = make(map[*State][]*State)
		if len(new_split) != 0 {
			split = new_split
			new_split = nil
		}
		for _, v := range split {
			if len(v) == 1 {
				new_split = append(new_split, v)
				if _, err := new_table[v[0]]; !err {
					new_table[v[0]] = v
				}
				continue
			}
			for j, a := range v {
				if _, err := new_table[a]; !err {
					var tmp []*State
					tmp = append(tmp, a)
					new_table[a] = tmp
					new_split = append(new_split, tmp)
				}
				for k := j + 1; k < len(v); k++ {
					var s = v[k]
					var add = true
					if _, err := new_table[s]; !err {
						for _, m := range dfa.Alphabet {
							var r1 = MakeTransition(a, m)
							var r2 = MakeTransition(s, m)
							if r1 != nil && r2 != nil {
								if Equal(new_table[r1], new_table[r2]) {
									add = false
									break
								}
							} else if !(r1 == nil && r2 == nil) {
								add = false
								break
							}
						}
					}
					if add {
						if _, err := new_table[s]; !err {
							new_table[a] = append(new_table[a], s)
							new_table[s] = new_table[a]
						}
					}
				}
			}
		}
		table = new_table
		new_table = nil
	}
	split = new_split
	for _, v := range split {
 		if len(v) > 1 {
 			var m =  v[0]
 			for i:=1; i < len(v); i++ {
 				var predecessors = GetPredecessors(dfa, v[i].Name)
 				var successors = GetSuccessors(dfa, v[i].Name)
 				for _, state := range predecessors {
 					var k string
 					for key, val := range state.Dtran {
 						if val == v[i] {
 							delete(state.Dtran, key)
 							k = key
 							break
						}
					}
					 state.Dtran[k] = m
				}
				for _, state := range successors {
					var k string
					for key, val := range m.Dtran {
						if val == state {
							delete(m.Dtran, k)
							k = key
							break
						}
					}
					m.Dtran[k] = state
				}
			}
		}
	}
}

func GetPredecessors(dfa *DFA, name string) (pred []*State) {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				if v.Name == name {
					pred = append(pred, state)
				}
				leftStates = append(leftStates, v)
			}
		}
	}
	return
}

func GetSuccessors(dfa *DFA, name string) (s []*State) {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				if state.Name == name && v.Name != name {
					s = append(s, v)
				}
				leftStates = append(leftStates, v)
			}
		}
	}
	return
}

func ListOfStates(dfa *DFA) (names []*State)  {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			names = append(names, state)
		}
	}
	return
}

func GetTransitions(start1 *State, start2 *State) map[string]map[string]*State {
	var res = make(map[string]map[string]*State)
	var leftStates = []*State{start1}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			res[state.Name] = state.Dtran
		}
	}
	leftStates = []*State{start2}
	seenStates = nil
	if start2 == nil {
		return res
	} else {
		for len(leftStates) != 0 {
			var state = leftStates[len(leftStates)-1]
			leftStates = leftStates[:len(leftStates)-1]
			if !FindSeen(state, seenStates) {
				seenStates = append(seenStates, state)
				for _, v := range state.Dtran {
					leftStates = append(leftStates, v)
				}
				res[state.Name] = state.Dtran
			}
		}
	}
	return res
}

func GetTransition(state *State, next string) string {
	for k, v := range state.Dtran {
		if v.Name == next {
			return k
		}
	}
	return ""
}

func GetState(dfa *DFA, name string) *State {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			if state.Name == name {
				return state
			}
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
		}
	}
	return nil
}

func CheckSelfLoop(dfa *DFA, name string) bool {
	var s = GetState(dfa, name)
	for _, v := range s.Dtran {
		if v.Name == name {
			return true
		}
	}
	return false
}

func GetKeysValues(m map[string]*State) (str []string, s []*State) {
	for k, v := range m {
		str = append(str, k)
		s = append(s, v)
	}
	return
}

type Pair struct {
	str string
	state *State
}

func CreateTransitions(tmp *DFA, k *State, flag bool) *State {
	if k.Name == tmp.InitialState.Name {
		return k
	}
	var p = GetPredecessors(tmp, k.Name)[0]
	var s = GetSuccessors(tmp, k.Name)[0]
	var loop = CheckSelfLoop(tmp, k.Name)
	/*if k.Name == "T" {
		fmt.Println(p, s, loop)
	}*/
	var pTs, sTp, pLoop, sLoop string = "", "", "", ""
	if !flag {
		if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
			GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
			GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
			GetTransition(p, k.Name) == "#" {
			pTs = "#"+GetTransition(p, k.Name)
		} else  {
			pTs = GetTransition(p, k.Name)
		}
	} else {
		if GetTransition(p, s.Name) == "+" || GetTransition(p, s.Name) == "|" ||
			GetTransition(p, s.Name) == "{" || GetTransition(p, s.Name) == "}" ||
			GetTransition(p, s.Name) == "\\" || GetTransition(p, s.Name) == "." ||
			GetTransition(p, s.Name) == "#" {
			pTs += "#"+GetTransition(p, s.Name)
		} else  {
			pTs += GetTransition(p, s.Name)
		}
		if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
			GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
			GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
			GetTransition(p, k.Name) == "#" {
			pTs += "#"+GetTransition(p, k.Name)
		} else  {
			pTs += GetTransition(p, k.Name)
		}
	}
	if loop {
		if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
			GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
			GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
			GetTransition(k, k.Name) == "#" {
			pTs += "#"+GetTransition(k, k.Name) + "*"
		} else  {
			pTs += GetTransition(k, k.Name) + "*"
		}
	}
	if GetTransition(k, s.Name) == "+" || GetTransition(k, s.Name) == "|" ||
		GetTransition(k, s.Name) == "{" || GetTransition(k, s.Name) == "}" ||
		GetTransition(k, s.Name) == "\\" || GetTransition(k, s.Name) == "." ||
		GetTransition(k, s.Name) == "#" {
		pTs += "#"+GetTransition(k, s.Name)
	} else  {
		pTs += GetTransition(k, s.Name)
	}
	if GetTransition(s, p.Name) == "+" || GetTransition(s, p.Name) == "|" ||
		GetTransition(s, p.Name) == "{" || GetTransition(s, p.Name) == "}" ||
		GetTransition(s, p.Name) == "\\" || GetTransition(s, p.Name) == "." ||
		GetTransition(s, p.Name) == "#" {
		sTp = "#" + GetTransition(s, p.Name)
	} else  {
		sTp = GetTransition(s, p.Name)
	}
	if GetTransition(s, k.Name) == "+" || GetTransition(s, k.Name) == "|" ||
		GetTransition(s, k.Name) == "{" || GetTransition(s, k.Name) == "}" ||
		GetTransition(s, k.Name) == "\\" || GetTransition(s, k.Name) == "." ||
		GetTransition(s, k.Name) == "#" {
		sTp += "#" + GetTransition(s, k.Name)
	} else  {
		sTp += GetTransition(s, k.Name)
	}
	if sTp != "" {
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				sTp += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				sTp += GetTransition(k, k.Name) + "*"
			}
		}
		if GetTransition(k, p.Name) == "+" || GetTransition(k, p.Name) == "|" ||
			GetTransition(k, p.Name) == "{" || GetTransition(k, p.Name) == "}" ||
			GetTransition(k, p.Name) == "\\" || GetTransition(k, p.Name) == "." ||
			GetTransition(k, p.Name) == "#" {
			sTp += "#" + GetTransition(k, p.Name)
		} else  {
			sTp += GetTransition(k, p.Name)
		}
	}
	var tr1_1, tr1_2 string
	if GetTransition(k, p.Name) == "+" || GetTransition(k, p.Name) == "|" ||
		GetTransition(k, p.Name) == "{" || GetTransition(k, p.Name) == "}" ||
		GetTransition(k, p.Name) == "\\" || GetTransition(k, p.Name) == "." ||
		GetTransition(k, p.Name) == "#" {
		tr1_1 = "#"+GetTransition(k, p.Name)
	} else  {
		tr1_1 = GetTransition(k, p.Name)
	}
	if GetTransition(p, k.Name) == "+" || GetTransition(p, k.Name) == "|" ||
		GetTransition(p, k.Name) == "{" || GetTransition(p, k.Name) == "}" ||
		GetTransition(p, k.Name) == "\\" || GetTransition(p, k.Name) == "." ||
		GetTransition(p, k.Name) == "#" {
		tr1_2 = "#"+GetTransition(p, k.Name)
	} else  {
		tr1_2 = GetTransition(p, k.Name)
	}
	if tr1_1 != "" && tr1_2 != "" {
		pLoop += tr1_2
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				pLoop += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				pLoop += GetTransition(k, k.Name) + "*"
			}
		}
		pLoop += tr1_1
	}
	var tr2_1, tr2_2 string
	if GetTransition(k, s.Name) == "+" || GetTransition(k, s.Name) == "|" ||
		GetTransition(k, s.Name) == "{" || GetTransition(k, s.Name) == "}" ||
		GetTransition(k, s.Name) == "\\" || GetTransition(k, s.Name) == "." ||
		GetTransition(k, s.Name) == "#" {
		tr2_1 = "#"+GetTransition(k, p.Name)
	} else  {
		tr2_1 = GetTransition(k, p.Name)
	}
	if GetTransition(s, k.Name) == "+" || GetTransition(s, k.Name) == "|" ||
		GetTransition(s, k.Name) == "{" || GetTransition(s, k.Name) == "}" ||
		GetTransition(s, k.Name) == "\\" || GetTransition(s, k.Name) == "." ||
		GetTransition(s, k.Name) == "#" {
		tr2_2 = "#"+GetTransition(p, k.Name)
	} else  {
		tr1_2 = GetTransition(p, k.Name)
	}
	if tr2_1 != "" && tr2_2 != "" {
		sLoop += tr2_2
		if loop {
			if GetTransition(k, k.Name) == "+" || GetTransition(k, k.Name) == "|" ||
				GetTransition(k, k.Name) == "{" || GetTransition(k, k.Name) == "}" ||
				GetTransition(k, k.Name) == "\\" || GetTransition(k, k.Name) == "." ||
				GetTransition(k, k.Name) == "#" {
				sLoop += "#" + GetTransition(k, k.Name) + "*"
			} else  {
				sLoop += GetTransition(k, k.Name) + "*"
			}
		}
		sLoop += tr2_1
	}
	delete(p.Dtran, GetTransition(p, k.Name))
	//fmt.Printf("%s %s %s %s\n", pTs, sTp, pLoop, sLoop)
	if pTs != "" {
		p.Dtran[pTs] = s
	}
	if sTp != "" {
		s.Dtran[sTp] = p
	}
	if pLoop != "" {
		p.Dtran[pLoop] = p
	}
	if sLoop != "" {
		s.Dtran[sLoop] = s
	}
	return s
}

var stack []*State

func CreateRE(tmp *DFA) (regex string) {
	var start = tmp.InitialState
	for true {
		if start.Receive {
			break
		}
		if len(start.Dtran) == 2 && CheckSelfLoop(tmp, start.Name) {
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) > 1  {
			for _, v := range start.Dtran {
				stack = append(stack, start)
				CreateSubOr(tmp, v)
			}
			for _, v := range start.Dtran {
				stack = append(stack, start)
				CreateSubOr(tmp, v)
			}
			var newMap = make(map[string]*State)
			var newKey = "("
			var _, end = GetKeysValues(start.Dtran)
			for k, _ := range start.Dtran {
				if k == "+" || k == "|" || k== "{" || k == "}" ||
					k == "\\" || k == "." || k == "#"{
					newKey += "(#" + k + ")|"
				} else {
					newKey += "(" + k + ")|"
				}
			}
			newKey = newKey[:len(newKey)-1]
			newKey += ")"
			newMap[newKey] = end[0]
			start.Dtran = newMap
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) == 1 && start != tmp.InitialState {
			start = CreateTransitions(tmp, start, true)
		} else if len(start.Dtran) == 1 && start == tmp.InitialState {
			_, v := GetKeysValues(start.Dtran)
			start = v[0]
		}
	}
	k, _ := GetKeysValues(tmp.InitialState.Dtran)
	return k[0]
}

func CreateSubOr(tmp *DFA, state *State) {
	fmt.Println(state)
	if len(GetPredecessors(tmp, state.Name)) > 1 {
		stack = stack[:len(stack)-1]
		return
	}
	if state.Receive {
		return
	}
	for len(stack) != 0 {
		//fmt.Println(state)
		if len(GetPredecessors(tmp, state.Name)) > 1 {
			stack = stack[:len(stack)-1]
			return
		}
		if len(state.Dtran) > 1 {
			for _, v := range state.Dtran {
				stack = append(stack, state)
				CreateSubOr(tmp, v)
			}
			var newMap = make(map[string]*State)
			var newKey = "("
			var _, end = GetKeysValues(state.Dtran)
			for k, _ := range state.Dtran {
				if k == "+" || k == "|" || k== "{" || k == "}" ||
					k == "\\" || k == "." || k == "#"{
					newKey += "(#" + k + ")|"
				} else {
					newKey += "(" + k + ")|"
				}
			}
			newKey = newKey[:len(newKey)-1]
			newKey += ")"
			newMap[newKey] = end[0]
			state.Dtran = newMap
			state = CreateTransitions(tmp, state, true)
		} else if len(state.Dtran) == 1 && state != tmp.InitialState {
			state = CreateTransitions(tmp, state, false)
		}
	}
	return
}

func DekartMul(mas1 []*State, mas2 []*State) (res []*State, m map[string]bool) {
	for _, i := range mas1 {
		for _, j := range mas2 {
			var str =  i.Name+","+j.Name
			res = append(res, &State{Name: i.Name+","+j.Name})
			if m == nil {
				m = make(map[string]bool)
			}
			m[str] = false
		}
	}
	return
}

func Mul(dfa1 *DFA, dfa2 *DFA) ([]*State, map[string]bool) {
	var names1 = ListOfStates(dfa1)
	var names2 = ListOfStates(dfa2)
	var stateList, reach = DekartMul(names1, names2)
	var trans = GetTransitions(dfa1.InitialState, dfa2.InitialState)
	reach[stateList[0].Name] = true
	for i:=0; i < len(stateList); i++ {
		var m = make(map[string]map[string]*State)
		var n []string
		var k string = ""
		for j, v := range stateList[i].Name {
			if v != ',' {
				k += string(v)
			} else {
				m[k] = trans[k]
				n = append(n, k)
				k = ""
			}
			if j == len(stateList[i].Name)-1 {
				m[k] = trans[k]
				n = append(n, k)
			}
		}
		var nextName string = ""
		for key, v := range m[n[0]] {
			nextName = v.Name+","
			for z:=1; z < len(n); z++ {
				for kz, vz := range m[n[z]] {
					if key == kz {
						nextName += vz.Name+","
						break
					}
				}
			}
			nextName = nextName[:len(nextName)-1]
			for a:=0; a < len(stateList); a++ {
				if stateList[a].Name == nextName {
					if stateList[i].Dtran == nil {
						stateList[i].Dtran = make(map[string]*State)
					}
					reach[stateList[a].Name] = reach[stateList[i].Name]
					stateList[i].Dtran[key] = stateList[a]
					break
				}
			}
		}
	}
	return stateList, reach
}

func FindInExits(exits []*State, name string) bool {
	for _, v := range exits {
		if v.Name == name {
			return true
		}
	}
	return false
}

func FindInExitsString(exits []string, name string) bool {
	for _, v := range exits {
		if v == name {
			return true
		}
	}
	return false
}

func Difference(dfa1 *DFA, dfa2 *DFA) *DFA {
	var stateList, reach = Mul(dfa1, dfa2)
	var NewStateList []*State
	for _, v := range stateList {
		if reach[v.Name] {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &DFA{InitialState: NewStateList[0]}
	var Exits []string
	var exits1, exits2, list2 = FindExit(dfa1.InitialState), FindExit(dfa2.InitialState), ListOfStates(dfa2)
	var ex []*State
	for _, v := range list2 {
		if !FindInExits(exits2, v.Name) {
			ex = append(ex, v)
		}
	}
	for _, v := range exits1 {
		for _, f := range ex {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
		if FindInExitsString(Exits, v.Name) && v.Name != NewStateList[0].Name && reach[v.Name] == true {
			v.Receive = true
		}
	}
	var leftStates = []*State{res.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for k, v := range state.Dtran {
				leftStates = append(leftStates, v)
				if !Find(res, k) {
					res.Alphabet = append(res.Alphabet, k)
				}
			}
		}
	}
	return res
}

func Intersection(dfa1 *DFA, dfa2 *DFA) *DFA {
	var stateList, reach = Mul(dfa1, dfa2)
	var NewStateList []*State
	for _, v := range stateList {
		if reach[v.Name] {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &DFA{InitialState: NewStateList[0]}
	var Exits []string
	for _, v := range FindExit(dfa1.InitialState) {
		for _, f := range FindExit(dfa2.InitialState) {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
		if FindInExitsString(Exits, v.Name) && v.Name != NewStateList[0].Name && reach[v.Name] == true {
			v.Receive = true
		}
	}
	var leftStates = []*State{res.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for k, v := range state.Dtran {
				leftStates = append(leftStates, v)
				if !Find(res, k) {
					res.Alphabet = append(res.Alphabet, k)
				}
			}
		}
	}
	return res
}
