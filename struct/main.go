package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

//Product is my first struck <3
type Product struct {
	Id    int    `json:"-"`
	Name  string `json:"nameOfTheProduct"` //This is the name
	Stock int
}

func (p Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

type Empty struct{}

func (e Empty) Calc() int {
	return 10
}

type Node struct {
	Value    *bool
	Operator *string
	Left     *Node
	Right    *Node
}

func (n *Node) Eval() bool {
	if n.Value != nil {
		return *n.Value
	}
	switch *n.Operator {
	case "AND":
		return n.Left.Eval() && n.Right.Eval()
	case "OR":
		return n.Left.Eval() || n.Right.Eval()
	}
	return false //TODO : error handling

}

type Node2 struct {
	Name     string
	Children []*Node2
}

type MapItem struct {
	A int
}

func main() {
	m := make(map[string]int)
	m["a"] = 5
	m["b"] = 2

	if _, found := m["b"]; found {
		fmt.Println("found")
	}

	items := make(map[int]*MapItem)

	items[1] = &MapItem{
		A: 1,
	}
	items[1].A = 2
	spew.Dump(items)
}

// for true {
// 	var id string
// 	if _, found := nodes[id]; !found {
// 		nodes[id] = Node2{
// 			Name:     "",
// 			Children: make([]*Node2, 0),
// 		}
// 	}
// 	nodes[id].Children = append(nodes[id].Children)
// }

/*
	data, _ := os.ReadFile("2.json")
	fmt.Println(string(data))

	root := Node{}
	json.Unmarshal(data, &root)
	fmt.Println(root)
	spew.Dump(root)
	fmt.Println(root.Eval())
*/

// prd := Product{
// 	Id:    5,
// 	Name:  "PC",
// 	Stock: 13,
// }

// data, _ := json.Marshal(prd)
// fmt.Println(string(data))

// prd.Name = "asd"

// _ = json.Unmarshal(data, &prd)
// fmt.Println(prd)

// e := Empty{}
// fmt.Println(e.Calc())

// fmt.Println("1", prd.GetName())
// prd.SetName("Joska")
// fmt.Println("2", prd.GetName())

// if result := e.Calc(); result > 1000 {
// 	fmt.Println(result)
// }

// arr := make([]int, 0)
// arr2 := []int{1, 2, 3}

// for i := 0; i < 10; i++ {
// 	arr = append(arr, i*2)
// }

// for i, elem := range arr {
// 	arr[i] = 5
// 	fmt.Println(elem)
// 	elem = 6
// }

// switch prd.id {
// case 1:
// 	fmt.Println("1. case")
// case 2:
// 	fmt.Println("2. case")
// default:
// 	fmt.Println("default")
// }
// prd.id = 1000
// switch {
// case prd.id > 10:
// 	fmt.Println("1. case")
// 	fallthrough
// case prd.id == 100:
// 	fmt.Println("2. case")
// }

// fmt.Println(arr, arr2)

// if (prd.id == 5 && prd.Name == "asd") || prd.id == 4 {

// } else if prd.id == 6 {

// } else {

// }
