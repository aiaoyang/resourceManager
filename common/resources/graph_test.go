package resources

import (
	"fmt"
	"testing"
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
}
type Arc struct {
	data   int
	status ArcStatus
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
func DFS(g *Graph, vertexIndex int, clock int) {
	fmt.Printf("now start index %d dfs \n", vertexIndex)
	queue.In(g.vertexs[vertexIndex])
	clock++
	g.vertexs[vertexIndex].discoverTime = clock
	//寻找所有当前节点的邻居节点,如果邻居节点都已访问,则此节点标记为visited
	for nbrIndex, nbrStat := range g.arcs[vertexIndex] {
		// 找到邻居节点
		if nbrStat != nil {
			fmt.Printf("nbr index not null %d\n", nbrIndex)
			// 判断邻居节点状态
			switch g.vertexs[nbrIndex].status {
			case discovered:
				// 如果是已发现状态,标记弧为环路
				g.arcs[vertexIndex][nbrIndex].status = backward
				fmt.Printf("discovered")
			case undiscovered:
				// 如果邻居节点是未发现的状态，标记弧为tree
				DFS(g, nbrIndex, clock)
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
	g := InitGraph(20)
	for i := 0; i < 20; i++ {
		g.vertexs[i] = &Vertex{
			data:         i,
			discoverTime: 0,
			finishedTime: 0,
			status:       0,
		}
	}
	g.arcs[1][2] = &Arc{status: initail}
	g.arcs[2][3] = &Arc{status: initail}
	g.arcs[3][4] = &Arc{status: initail}
	g.arcs[4][5] = &Arc{status: initail}
	g.arcs[6][7] = &Arc{status: initail}
	clock := 0
	DFS(g, 0, clock)
	for k, v := range g.vertexs {
		if v.status != visited {
			DFS(g, k, clock)
		}
	}
	for !queue.empty() {
		fmt.Printf("%v -> ", queue.Out())
	}
}
