package store

import "testing"

func TestGenerateDAG(t *testing.T) {
	size := 100000
	edges := size - 1

	graph := GenerateGraph(size)

	foundSize := graph.CountVertex()
	if foundSize != size {
		t.Fatalf("expected vertices count %d, fount %d", size, foundSize)
	}

	foundEdges := graph.CountEdge()
	if foundEdges != edges {
		t.Fatalf("expected edges count %d, fount %d", edges, foundEdges)
	}
}

func BenchmarkGenerateDAG(t *testing.B) {
	size := 100000

	_ = GenerateGraph(size)
}

func BenchmarkReach(t *testing.B) {
	size := 100000
	graph := GenerateGraph(size)

	v := getVertex(graph, t)

	graph.Reach(v.ID)
}

func BenchmarkConditionalReach(t *testing.B) {
	size := 100000
	graph := GenerateGraph(size)

	v := getVertex(graph, t)

	graph.ConditionalReach(v.ID, false)
}

func BenchmarkList(t *testing.B) {
	size := 100000
	graph := GenerateGraph(size)

	v := getVertex(graph, t)

	graph.List(v.ID)
}

func BenchmarkConditionalList(t *testing.B) {
	size := 100000
	graph := GenerateGraph(size)

	v := getVertex(graph, t)

	graph.ConditionalList(v.ID, false)
}

func getVertex(graph *DAG, t *testing.B) *Vertex {
	var index int = 0
	// index := rand.Intn(graph.CountVertex())

	v, err := graph.GetVertexByIndex(index)
	if err != nil {
		t.Fatal(err)
	}

	return v
}
