package resources

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

const (
	vertexNum = 100
)

var (
	stack = InitStack(100)
)

// Queue
type Queue struct {
	head  int
	tail  int
	size  int
	store []*Vertex
}

func InitQueue(size int) *Queue {
	q := &Queue{}
	q.head = 0
	q.tail = 0
	q.size = size
	q.store = make([]*Vertex, size)
	return q
}

//func (q *Queue) In(e *Vertex) {
//
//}
//func (q *Queue) Out() *Vertex {
//	//if q.head
//	e := q.store[q.head]
//	q.head = (q.head + 1) % q.size
//	return e
//}
// Stack
type Stack struct {
	top   int
	store []*Vertex
	size  int
}

func InitStack(size int) *Stack {
	s := &Stack{}
	store := make([]*Vertex, size)
	s.top = 0
	s.store = store
	s.size = size
	return s
}
func (s *Stack) Put(e *Vertex) {
	if s.top == s.size {
		log.Printf("stack is full\n")
		return
	}
	//fmt.Printf("elem %s is putted\n", e)
	s.store[s.top] = e
	s.top++
}
func (s *Stack) Pop() *Vertex {
	if s.top == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	s.top--
	tmp := s.store[s.top]
	s.store[s.top] = nil
	return tmp
}

// graph data structure
type Vertex struct {
	data    string
	fTime   int
	dTime   int
	visited bool
}
type Edge struct {
	data    int
	visited bool
}
type Graph struct {
	vertex   []*Vertex
	linkList [][]*Edge
	vexNum   int
}

func InitGraph() *Graph {
	g := &Graph{}
	g.vexNum = 10
	vertex := make([]*Vertex, 0, g.vexNum)
	for i := 0; i < g.vexNum; i++ {
		tmp := &Vertex{data: strconv.Itoa(i)}
		vertex = append(vertex, tmp)
	}
	linkList := make([][]*Edge, g.vexNum)
	for i := 0; i < g.vexNum; i++ {
		linkList[i] = make([]*Edge, g.vexNum)
	}
	g.vertex = vertex
	g.linkList = linkList
	return g
}
func Relation(g *Graph, e *Vertex) {

}
func (g *Graph) AddNode(e *Vertex) {
	fmt.Printf("add befor len is : %d\n", len(g.vertex))
	g.vertex = append(g.vertex, e)
	for _, v := range g.linkList {
		v = append(v, nil)
	}
	fmt.Printf("add befor len is : %d\n", len(g.vertex))
	g.linkList = append(g.linkList, make([]*Edge, len(g.linkList[0])))
}
func (g *Graph) AddArc(m, n *Vertex) {
	if m.data == n.data {
		fmt.Println("two equal value")
		return
	}
	var bm, bn bool
	for _, v := range g.vertex {
		if v.data == m.data {
			bm = true
			continue
		}
		if v.data == n.data {
			bn = true
			continue
		}
	}
	if !bm || !bn {
		fmt.Printf("one of the value (%v %v) does not exist\n", m, n)
		return
	}
	for k, v := range g.vertex {
		if v.data == m.data {
			for i, j := range g.vertex {
				if j.data == n.data {
					g.linkList[k][i] = &Edge{data: 1}
				}
			}
		}
	}
}
func (g *Graph) DFS(e *Vertex) {
	stack.Put(e)
	//for start, value := range g.vertex {
	//
	//}
	for indexOfe, v := range g.vertex {
		if v.data == e.data {
			// 在邻接矩阵中寻找与起始节点关联的关联节点
			for i, j := range g.linkList[indexOfe] {
				// 如果找到关联节点，则将关联节点入栈
				if j != nil {
					// 如果边已经被访问过，则终止这次操作
					if j.visited {
						break
					}
					j.visited = true
					g.DFS(g.vertex[i])
					//fmt.Printf("i is : %d\n", i)
					//break
				}
			}
			//}
			for stack.top != 0 {
				fmt.Printf(" <- %v", stack.Pop().data)
			}
			println()
		}
	}

}
func (g *Graph) String() {
	for k, v := range g.linkList {
		for i, j := range v {
			if j != nil {
				fmt.Printf("%s -> %s\n", g.vertex[k].data, g.vertex[i].data)
			}
		}
	}
}
func (g *Graph) test() {

}
func Test_resource(t *testing.T) {
	g := InitGraph()
	//g.AddArc(&Vertex{"2"}, &Vertex{"4"})
	g.AddArc(&Vertex{data: "2"}, &Vertex{data: "5"})
	g.AddArc(&Vertex{data: "2"}, &Vertex{data: "7"})
	g.AddArc(&Vertex{data: "5"}, &Vertex{data: "2"})
	//g.AddArc(&Vertex{data: "3"}, &Vertex{data: "7"})
	//g.AddArc(&Vertex{data: "2"}, &Vertex{data: "8"})

	g.String()
	g.DFS(&Vertex{data: "7"})
	//s := InitStack(10)
	//for i := 0; i < 10; i++ {
	//	s.Put(&Vertex{strconv.Itoa(i)})
	//}
	//for s.top != 0 {
	//	fmt.Println(s.Pop())
	//}
}
