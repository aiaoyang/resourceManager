package resources

import (
	"fmt"
	"testing"
)

var (
	//stack = InitStack(100)
	queue = InitQueue(200)
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
func (q *Queue) empty() bool {
	if q.tail == q.head {
		return true
	}
	return false
}

// 队列实际可用空间为q.size-1,因为使用了一个空间判断队列满
func (q *Queue) full() bool {
	if (q.tail+1)%q.size == q.head {
		return true
	}
	return false
}

func (q *Queue) In(e *Vertex) {
	if q.full() {
		fmt.Println("queue is full")
		return
	}
	q.store[q.tail] = e
	q.tail = (q.tail + 1) % q.size
	//fmt.Printf("in queue now queue's head is : %d, tail is %d \n", q.head, q.tail)
}

func (q *Queue) Out() *Vertex {
	if q.empty() {
		fmt.Println("queue is empty")
		return nil
	}
	tmp := q.store[q.head]
	q.head = (q.head + 1) % q.size
	//fmt.Printf("out queue now queue's head is : %d, tail is %d \n", q.head, q.tail)
	return tmp
}

var (
	clock = 0
)

type ArcStatus int
type VexStatus int

const (
	undiscovered VexStatus = iota
	discovered
	visited
)
const (
	initail ArcStatus = iota
	tree
	forward
	backward
	cross
)

type Graph struct {
	vertexs []*Vertex
	arcs    [][]*Arc
}

type Vertex struct {
	data         int
	discoverTime int
	finishedTime int
	status       VexStatus
	inDegree     int
	outDegree    int
	child        []*Vertex
	parent       *Vertex
}
type Arc struct {
	data   int
	status ArcStatus
}

func reset(g *Graph) {
	for _, v := range g.vertexs {
		v.status = undiscovered
		v.discoverTime = 0
		v.finishedTime = 0
		v.child = make([]*Vertex, 1, 10)
		arcs := make([][]*Arc, len(g.vertexs))
		for i := 0; i < len(g.vertexs); i++ {
			arcs[i] = make([]*Arc, len(g.vertexs))
		}
		g.arcs = arcs
	}
}
func InitGraph(nodeNum int) *Graph {
	arcs := make([][]*Arc, nodeNum)
	for i := 0; i < nodeNum; i++ {
		arcs[i] = make([]*Arc, nodeNum)
	}
	return &Graph{
		vertexs: make([]*Vertex, nodeNum),
		arcs:    arcs,
	}
}
func DFS(g *Graph, vertexIndex int) {
	//clock := clockIn
	g.vertexs[vertexIndex].status = discovered
	//fmt.Printf("now start index %d dfs , clock is %d\n", vertexIndex, clock)
	queue.In(g.vertexs[vertexIndex])
	clock++
	//fmt.Println(clock)
	g.vertexs[vertexIndex].discoverTime = clock
	//寻找所有当前节点的邻居节点,如果邻居节点都已访问,则此节点标记为visited
	for nbrIndex, nbrStat := range g.arcs[vertexIndex] {
		// 找到邻居节点
		if nbrStat != nil {
			g.vertexs[vertexIndex].child = append(g.vertexs[vertexIndex].child, g.vertexs[nbrIndex])
			// 存在邻居节点则邻居节点入度+1
			g.vertexs[nbrIndex].inDegree++
			// 存在邻居节点则当前节点出度+1
			g.vertexs[vertexIndex].outDegree++
			//fmt.Printf("nbr index not null %d\n", nbrIndex)
			// 判断邻居节点状态
			switch g.vertexs[nbrIndex].status {
			case discovered:
				// 如果是已发现状态,标记弧为环路
				g.arcs[vertexIndex][nbrIndex].status = backward
				//fmt.Printf("discovered")
			case undiscovered:
				// 设置当前节点的孩子节点为邻居节点,邻居节点的父节点为当前节点
				//g.vertexs[nbrIndex].parent = g.vertexs[vertexIndex]
				// 如果邻居节点是未发现的状态，标记弧为tree
				DFS(g, nbrIndex)
				g.arcs[vertexIndex][nbrIndex].status = tree
			case visited:
				// 如果邻居节点是访问结束状态,根据发现时间标记弧为跨边或向前边
				if g.vertexs[vertexIndex].discoverTime > g.vertexs[nbrIndex].discoverTime {
					// 如果邻居节点的发现时间早于当前节点时间,则邻居节点是当前节点的跨边
					g.arcs[vertexIndex][nbrIndex].status = cross
				} else {
					// 如果邻居节点的发现时间晚于当前节点时间,则邻居节点是当前节点的向前边
					g.arcs[vertexIndex][nbrIndex].status = forward
				}
			}
		}
	}
	clock++
	g.vertexs[vertexIndex].status = visited
	g.vertexs[vertexIndex].finishedTime = clock
	//if vertex.status == undiscovered {
	//	vertex.status = discovered
	//	queue.In(vertex)
	//}

}

func Test_graph(t *testing.T) {
	g := InitGraph(10)
	for i := 0; i < 10; i++ {
		g.vertexs[i] = &Vertex{
			data:         i,
			discoverTime: 0,
			finishedTime: 0,
			status:       0,
		}
	}
	g.arcs[0][1] = &Arc{status: initail}
	g.arcs[6][2] = &Arc{status: initail}
	g.arcs[6][0] = &Arc{status: initail}
	g.arcs[2][3] = &Arc{status: initail}
	g.arcs[2][6] = &Arc{status: initail}
	g.arcs[1][3] = &Arc{status: initail}

	//clockTime := 0
	DFS(g, 6)
	for k, v := range g.vertexs {
		if v.status != visited {
			DFS(g, k)
		}
	}
	for !queue.empty() {
		vertex := queue.Out()
		//if vertex.parent == nil && vertex.child == nil {
		//	continue
		//}
		for _, v := range vertex.child {
			fmt.Printf("vertex :%v, child is : %v\n", vertex.data, v.data)
		}
		//fmt.Printf("vertex: %v, discover: %v, visited: %v \n", vertex.data, vertex.discoverTime, vertex.finishedTime)
	}
	//clock = 0
	//reset(g)
	//for k, v := range g.vertexs {
	//	if v.status != visited {
	//		DFS(g, k)
	//	}
	//}
	//for !queue.empty() {
	//	vertex := queue.Out()
	//	fmt.Printf("vertex: %v, discover: %v, visited: %v \n", vertex.data, vertex.discoverTime, vertex.finishedTime)
	//}

	//i := 0
	//AddInt(i)
	//fmt.Println(i)
}
func AddInt(i int) {
	if i < 10 {
		i++
		AddInt(i)
		fmt.Println(i)
	}
	return
}
