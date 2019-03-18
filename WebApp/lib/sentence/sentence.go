package sentence

import (
	_ "fmt"
	_ "github.com/liudng/godump"
	"fmt"
	"strings"
	)


type Node struct {
	Value  string
	Weight int
    Parent *Node
    Left   *Node
    Right  *Node
}

type Result struct {
	Str string
	Score int
}

type Tree struct {
	root *Node
}

type InputSubject struct {
	TotalString  string
	TargetString string
}

func (subject InputSubject) Analyse() *Tree {

	//subject := InputSubject{"The quick brown fox jumps over the lazy dog , and the brown dog besides watches the jump of the fox . I saw everything through my window and later on tell you the story of the brown fox",
	//                       "brown fox"}

	subjectArray := strings.Split(subject.TotalString, " ")
	targetArray  := strings.Split(subject.TargetString, " ")

	tree := &Tree{}
    tree  = tree.InitTree()

	flag := 0
	for index, element := range subjectArray{
		_ = index
		if ( flag > (len(targetArray) -1 )) {
			flag = 0
		}

		if (ManageTree(element, targetArray, flag, tree)) {
			flag ++
		}
	}

	// printTheTree(tree)
	// os.Exit(2)
	return tree
}

func (tree *Tree) ProcessResult(input InputSubject) []Result {
	lastNode := *tree.root

	TotalArr  := strings.Split(input.TotalString, " ")
	targetArr := strings.Split(input.TargetString, " ")
	targetLen := len(targetArr)

	_ = TotalArr

	flag := 0

	temp := make([]string, 0)
	tempWeight := 0
	lenArr := make([]string, 0)

	results := []Result{}

	for ( lastNode.Right != nil || lastNode.Left != nil ) {

		if (lastNode.Value != "" && lastNode.Parent  != nil) {
			lenArr = append(lenArr, lastNode.Value )
		}

		if ((lastNode.Value == "" && lastNode.Weight == 0 )|| lastNode.Left != nil) {
			left := *lastNode.Left
			temp = append(temp,left.Value)
			tempWeight += left.Weight
			flag ++
		}

		if (flag > 0 ) {
			if ( lastNode.Value != "") {
				temp = append(temp, lastNode.Value)
			}
			tempWeight += lastNode.Weight
		}

		if (flag > (targetLen - 1)) {
			flag = 0

			beforeStr   := strings.Join(lenArr, " ")
			indexLength := len(beforeStr)

			fmt.Println(indexLength)
			fmt.Println(beforeStr)
			fmt.Println(lenArr)

			lenArr = make([]string,0)

			results = append(results,Result{
				Str: strings.Join(temp, " "),
				Score: tempWeight ,
			})

			tempWeight = 0
			temp = make([]string, 0)
		}


		if (lastNode.Right == nil){
			break
		}

		lastNode = *lastNode.Right
	}

	return results
}

func ManageTree (subjectWord string , targetWord []string, flag int, tree *Tree) bool {
	if (flag == -1 ) {
		return false
	}
	if (targetWord[flag] == subjectWord) {
		tree.Insert(targetWord[flag], 0)
		return true
	} else {
		tree.Insert(subjectWord, 1)
		return false
	}
}

func (tree *Tree) InitTree () *Tree {
	if (tree.root == nil) {
		tree.root = &Node{
			Parent: nil,
			Value: "Root",
			Weight: -1,
			Left: nil,
			Right: nil,
		}
	}

	return tree
}

func (tree *Tree)Insert (value string, ori int) *Tree  {
	tree.root.insertNode(value, ori)
	return tree
}

func (n *Node)insertNode(value string, ori int)  {
	if (n == nil) {
		return
	}

	if (ori == 0) {
		zero := Node{
			Parent: n,
			Value: "",
			Weight: 0,
		}

		tmp := Node{
			Parent: &zero,
			Value: value,
			Weight: 0,
			Left: nil,
			Right: nil,
		}

		zero.Left  = &tmp
		zero.Right = nil


		if (n.Right == nil) {
			n.Right = &zero
			return
		}

		if (n.Right != nil) {
			n.Right.insertNode(value, ori)
			return
		}
	}

	if (ori == 1) {
		tmp := Node{
			Parent:n,
			Weight:1,
			Value: value,
			Left: nil,
			Right: nil,
		}

		if (n.Right == nil) {
			n.Right = &tmp
			return
		}

		if (n.Right != nil) {
			n.Right.insertNode(value, ori)
			return
		}
	}
}


func (tree *Tree) PrintTheTree()  {

	lastNode := *tree.root

	for (lastNode.Right != nil || lastNode.Left != nil ) {
		if (lastNode.Left != nil) {
			// Has a left
			temp := *lastNode.Left
			fmt.Println(temp.Value + "\t\t" + "\n")
			fmt.Println("\t\t\t\t\t\t\t" + "|" + "\n")
			fmt.Println("\t\t\t\t\t\t\t" + "|" + "\n")
		}
		fmt.Println("\t\t\t\t\t\t\t" + lastNode.Value + "\n")
		fmt.Println("\t\t\t\t\t\t\t" + "|" + "\n")
		fmt.Println("\t\t\t\t\t\t\t" + "|" + "\n")


		if (lastNode.Right == nil && lastNode.Left == nil) {
			break
		}
		lastNode = *lastNode.Right

	}
}
