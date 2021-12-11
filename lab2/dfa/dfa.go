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
	num int
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
				//fmt.Println(k, v, i)
				if len(i) != 0 {
					var nextStateNumber []int
					for _, c := range i {
						nextStateNumber = tree.Merge(nextStateNumber, dfa.FollowPos[c-1])
						//				fmt.Println(nextStateNumber, c)
					}
					var broken = false
					/*for _, v := range seenStates {
						fmt.Printf("%s ", v.Name)
					}
					fmt.Println(nextStateNumber)*/
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
						//				fmt.Println(state.Name, state.Dtran)
						if state.Dtran == nil {
							state.Dtran = make(map[string]*State)
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

func Print(start *State) {
	var leftStates = []*State{start}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			fmt.Printf("<%s  %v>     ", state.Name, state.StateNumber)
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

func FindExit(state *State) (s []*State) {
	var leftStates = []*State{state}
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

func FindAnother(state *State) (mas []*State) {
	var leftStates = []*State{state}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				leftStates = append(leftStates, v)
			}
			if state.Dtran != nil {
				mas = append(mas, state)
			}
		}
	}
	return mas
}

func Compile(regex string) (d *DFA) {
	var SyntaxTree tree.Ast
	SyntaxTree.ID = make(map[int]string)
	SyntaxTree.LeafNodes = make(map[string][]int)
	SyntaxTree.Root, SyntaxTree.FollowPos = tree.CreateTree(regex, SyntaxTree.ID, SyntaxTree.LeafNodes, SyntaxTree.FollowPos)
	tree.PrintTree(SyntaxTree.Root)
	//fmt.Println(SyntaxTree.ID)
	//fmt.Println(SyntaxTree.FollowPos)
	//fmt.Println(SyntaxTree.LeafNodes)
	d = InitDFA(SyntaxTree, SyntaxTree.FollowPos, SyntaxTree.Root.FirstPos)
	Convert(d, SyntaxTree.Root, SyntaxTree.LeafNodes)
	//fmt.Println(d)
	Print(d.InitialState)
	return
}

func Search(pattern interface{}, str string) string {
	var dfa *DFA
	if reflect.TypeOf(pattern) == reflect.TypeOf("11") {
		dfa = Compile(pattern.(string))
	} else {
		dfa = pattern.(*DFA)
	}
	var state = dfa.InitialState.Dtran
	var i = 0
	var subStr = false
	var finded bool
	var startInd= -1
	for i < len(str) {
		finded = false
		for k, v := range state {
			if string(str[i]) == k {
				if startInd == -1 {
					startInd = i
				}
				finded = true
				subStr = true
				state = v.Dtran
			}
		}
		if !finded {
			startInd = -1
			state = dfa.InitialState.Dtran
		}
		if subStr && state == nil {
			return str[startInd:i+1]
		}
		i++
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

func Minimization(dfa *DFA) {
	var states = ListOfStates(dfa)
	var exits []string
	for _, v := range FindExit(dfa.InitialState) {
		exits = append(exits, v.Name)
	}
	for i, v := range states {
		for _, val := range exits {
			if v == val {
				copy(states[i:], states[i+1:])
				states[len(states)-1] = ""
				states = states[:len(states)-1]
			}
		}
	}
	var piSplit [][]string
	piSplit = append(piSplit, states)
	piSplit = append(piSplit, exits)
	/*for k, v := range piSplit {

	}*/
	fmt.Println(piSplit)
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
			//fmt.Println(state.Name, state.Dtran)
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

func ListOfStates(dfa *DFA) (names []string)  {
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
			names = append(names, state.Name)
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


func CreateRE(tmp *DFA, s *State, m map[string]map[string]*State) *State {
	var start = s
	var exit = FindExit(start)
	var exitFlag bool
	fmt.Println(start, exit)
	for _, k := range ListOfStates(tmp) {
		if k == start.Name {
			continue
		}
		exitFlag = false
		for _, v := range exit {
			if k == v.Name {
				exitFlag = true
			}
		}
		if exitFlag {
			continue
		}
		var p = GetPredecessors(tmp, k)
		var s = GetSuccessors(tmp, k)
		var loop = CheckSelfLoop(tmp, k)
		var state = GetState(tmp, k)
		//fmt.Println(p[0].Name, s[0].Name, loop, state.Name)
		//fmt.Println(m)
		var pTs, sTp, pLoop, sLoop string = "", "", "", ""
		if len(s) > 1 {
			for _, v := range s {
				CreateRE(tmp, v, m)
			}
		} else {
			for i := 0; i < len(p); i++ {
				for j := 0; j < len(s); j++ {
					//fmt.Println(i, j, p, s)
					//путь предыдущий->следующий
					pTs = GetTransition(p[i], s[j].Name) + GetTransition(p[i], k)
					//fmt.Println(pTs)
					if loop {
						pTs += GetTransition(state, k) + "*"
					}
					pTs += GetTransition(state, s[j].Name)
					//fmt.Println(pTs)
					//путь следующий предыдущий
					//fmt.Println(s[j].Name, p[i].Name, k)
					sTp = GetTransition(s[j], p[i].Name) + GetTransition(s[j], k)
					if sTp != "" {
						if loop {
							sTp += GetTransition(state, k) + "*"
						}
						sTp += GetTransition(state, p[i].Name)
					}
					//луп для предыдущего
					var tr1_1 = GetTransition(state, p[i].Name)
					var tr1_2 = GetTransition(p[i], state.Name)
					if tr1_1 != "" && tr1_2 != "" {
						pLoop += tr1_2
						if loop {
							pLoop += GetTransition(state, k) + "*"
						}
						pLoop += tr1_1
					}
					//луп для следующего
					var tr2_1 = GetTransition(state, s[j].Name)
					var tr2_2 = GetTransition(s[j], state.Name)
					if tr2_1 != "" && tr2_2 != "" {
						pLoop += tr2_2
						if loop {
							pLoop += GetTransition(state, k) + "*"
						}
						pLoop += tr2_1
					}
					//fmt.Println(pTs)
					delete(m[p[i].Name], GetTransition(p[i], k))
					if pTs != "" {
						m[p[i].Name][pTs] = s[j]
					}
					if sTp != "" {
						m[s[j].Name][sTp] = p[i]
					}
					if pLoop != "" {
						m[p[i].Name][pLoop] = p[i]
					}
					if sLoop != "" {
						m[s[j].Name][sLoop] = s[j]
					}
					//fmt.Printf("pTs=%v, sTp=%v, pL=%v, sL=%v\n", pTs, sTp, pLoop, sLoop)
					//fmt.Println(m)
					pTs, sTp, pLoop, sLoop = "", "", "", ""
				}
			}
		}
	}
	return start
}

type ResState struct {
	Name        string
	Dtran       map[string]*ResState
	Reachable bool
}

func DekartMul(mas1 []string, mas2 []string) (res []*ResState) {
	for _, i := range mas1 {
		for _, j := range mas2 {
			res = append(res, &ResState{Name: i+","+j, Reachable: false})
		}
	}
	return
}

func Mul(dfa1 *DFA, dfa2 *DFA) []*ResState {
	var names1 = ListOfStates(dfa1)
	var names2 = ListOfStates(dfa2)
	var stateList = DekartMul(names1, names2)
	var trans = GetTransitions(dfa1.InitialState, dfa2.InitialState)
	stateList[0].Reachable = true
	for i:=0; i < len(stateList); i++ {
		var m = make(map[string]map[string]*State)
		var n []string
		var k string = ""
		//fmt.Println(stateList[i])
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
					//fmt.Println(key, kz)
					if key == kz {
						nextName += vz.Name+","
						break
					}
				}
			}
			nextName = nextName[:len(nextName)-1]
			//fmt.Println(nextName)
			for a:=0; a < len(stateList); a++ {
				//fmt.Println(stateList[a])
				if stateList[a].Name == nextName {
					if stateList[i].Dtran == nil {
						stateList[i].Dtran = make(map[string]*ResState)
					}
					stateList[a].Reachable = stateList[i].Reachable
					stateList[i].Dtran[key] = stateList[a]
					break
				}
			}
		}
	}
	return stateList
}

type ResultDfa struct {
	InitialState *ResState
	ReceiveStates []*ResState
}


func FindInExits(exits []string, name string) bool {
	for _, v := range exits {
		if v == name {
			return true
		}
	}
	return false
}


func Difference(dfa1 *DFA, dfa2 *DFA) *ResultDfa {
	var stateList = Mul(dfa1, dfa2)
	for _, v := range stateList {
		fmt.Println(v)
	}
	var NewStateList []*ResState
	for _, v := range stateList {
		if v.Reachable {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &ResultDfa{InitialState: NewStateList[0], ReceiveStates: nil}
	var Exits []string
	for _, v := range FindExit(dfa1.InitialState) {
		for _, f := range FindExit(dfa2.InitialState) {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
		if FindInExits(Exits, v.Name) && v.Name != NewStateList[0].Name && v.Reachable == true {
			res.ReceiveStates = append(res.ReceiveStates, v)
		}
	}
	for _, v := range res.ReceiveStates {
		fmt.Println(v)
	}
	return res
}

func Intersection(dfa1 *DFA, dfa2 *DFA) *ResultDfa {
	var stateList = Mul(dfa1, dfa2)
	for _, v := range stateList {
		fmt.Println(v)
	}
	var NewStateList []*ResState
	for _, v := range stateList {
		if v.Reachable {
			NewStateList = append(NewStateList, v)
		}
	}
	var res = &ResultDfa{InitialState: NewStateList[0], ReceiveStates: nil}
	var Exits []string
	for _, v := range FindExit(dfa1.InitialState) {
		for _, f := range FindExit(dfa2.InitialState) {
			Exits = append(Exits, v.Name+","+f.Name)
		}
	}
	for _, v := range NewStateList {
		if FindInExits(Exits, v.Name) && v.Name != NewStateList[0].Name && v.Reachable == true {
			res.ReceiveStates = append(res.ReceiveStates, v)
		}
	}
	return res
}
