package dfa

import (
	"fmt"
	"lab2/tree"
	"reflect"
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
	//fmt.Println(LeafNodes)
	//fmt.Println(dfa.InitialStateNumber)
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

func FindExit(state *State) (s *State) {
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
				s = state
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
	SyntaxTree.Root = tree.CreateTree(regex, SyntaxTree.ID, SyntaxTree.LeafNodes)
	tree.PrintTree(SyntaxTree.Root)
	fmt.Println(SyntaxTree.ID)
	//fmt.Println(SyntaxTree.ID)
	for i := 0; i < len(SyntaxTree.ID); i++ {
		SyntaxTree.FollowPos = append(SyntaxTree.FollowPos, nil)
	}
	tree.FindFollowPos(SyntaxTree.Root, SyntaxTree.FollowPos)
	fmt.Println(SyntaxTree.FollowPos)
	d = InitDFA(SyntaxTree, SyntaxTree.FollowPos, SyntaxTree.Root.FirstPos)
	var s = Convert(d, SyntaxTree.Root, SyntaxTree.LeafNodes)
	fmt.Println(d)
	fmt.Println(s)
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

func Minimization(dfa *DFA) {
	var states = ListOfStates(dfa)
	var exits = []string{FindExit(dfa.InitialState).Name}
	for i, v := range states {
		if v == exits[0] {
			copy(states[i:], states[i+1:])
			states[len(states)-1] = ""
			states = states[:len(states)-1]
		}
	}
	var piSplit [][]string
	piSplit = append(piSplit, states)
	piSplit = append(piSplit, exits)
	for k, v := range piSplit {
		
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

func GetSuccessors(dfa *DFA, name string) (pred []*State) {
	var leftStates = []*State{dfa.InitialState}
	var seenStates []*State
	for len(leftStates) != 0 {
		var state = leftStates[len(leftStates)-1]
		leftStates = leftStates[:len(leftStates)-1]
		if !FindSeen(state, seenStates) {
			seenStates = append(seenStates, state)
			for _, v := range state.Dtran {
				if state.Name == name {
					pred = append(pred, v)
				}
				leftStates = append(leftStates, v)
			}
		}
	}
	return
}

func CheckSelfLoop(s *State) bool {
	for _, v := range s.Dtran {
		if v == s {
			return true
		}
	}
	return false
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


func CreateRE(dfa *DFA) *State {
	var tmp = CopyDFA(dfa)
	var start = tmp.InitialState
	var exit = FindExit(start)
	var m = GetTransitions(start, nil)
	fmt.Println(m)
	for _, k := range ListOfStates(tmp) {
		if k == start.Name && k == exit.Name {
			continue
		}
		var p = GetPredecessors(tmp, k)
		var s = GetSuccessors(tmp, k)
		var loop = CheckSelfLoop(p[0].Dtran[k])
		var pTs, sTp, pLoop, sLoop string = "", "", "", ""
		for i:=0; i < len(p); i++ {
			for j:=0; j < len(s); j++ {
				//путь предыдущий->следующий
				pTs = GetTransition(p[i], s[j].Name)+GetTransition(p[i], k)
				if loop {
					pTs += GetTransition(p[i].Dtran[k], k) + "+"
				}
				pTs += GetTransition(p[i].Dtran[k], s[j].Name)
				//путь следующий предыдущий
				sTp = GetTransition(s[j], p[i].Name)+GetTransition(s[j], k)
				if loop {
					sTp += GetTransition(p[i].Dtran[k], k) + "+"
				}
				sTp += GetTransition(s[j].Dtran[k], p[i].Name)
				//луп для предыдущего
				var tr1 = GetTransition(p[i].Dtran[k], p[i].Name)
				if tr1 != "" {
					pLoop += GetTransition(p[i], k)
					if loop {
						pLoop += GetTransition(p[i].Dtran[k], k) + "+"
					}
					pLoop += tr1
				}
				//луп для следующего
				var tr2 = GetTransition(s[j].Dtran[k], s[j].Name)
				if tr2 != "" {
					sLoop += GetTransition(s[j], k)
					if loop {
						sLoop += GetTransition(s[j].Dtran[k], k) + "+"
					}
					sLoop += tr2
				}
				delete(m[p[i].Name], GetTransition(p[i], k))
				m[p[i].Name][pTs] = s[j]
				m[s[j].Name][sTp] = p[i]
				m[p[i].Name][pLoop] = p[i]
				m[s[j].Name][sLoop] = s[j]
				pTs, sTp, pLoop, sLoop = "", "", "", ""
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
	var ExitState = FindExit(dfa1.InitialState).Name + "," + FindExit(dfa2.InitialState).Name
	for _, v := range NewStateList {
		if v.Name != ExitState && v.Name != NewStateList[0].Name {
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
	var ExitState = FindExit(dfa1.InitialState).Name + "," + FindExit(dfa2.InitialState).Name
	for _, v := range stateList {
		if v.Name == ExitState {
			res.ReceiveStates = append(res.ReceiveStates, v)
		}
	}
	return res
}