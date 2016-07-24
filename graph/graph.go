package graph

import (
	"fmt"
	"sync"
	"bytes"
	"io"
	"encoding/json"
)

// ID is unique identifier.
type ID interface {
	// String returns the string ID.
	String() string
}

type StringID string

func (s StringID) String() string {
	return string(s)
}

// Vertex is a graph vertex. The ID must be unique within the graph.
type Vertex interface {
	// ID returns the ID.
	ID() ID
	String() string
}

type vertex struct {
	id string
}

func NewVertex(id string) Vertex {
	return &vertex{
		id: id,
	}
}

func (v *vertex) ID() ID {
	return StringID(v.id)
}

func (v *vertex) String() string {
	return v.id
}

// Edge connects between two Vertices.
type Edge interface {
	Source() Vertex
	Target() Vertex
	Weight() float64
	String() string
}


// edge is an Edge from Source to Target.
type edge struct {
	src Vertex
	tgt Vertex
	wgt float64
}

func NewEdge(src, tgt Vertex, wgt float64) Edge {
	return &edge{
		src: src,
		tgt: tgt,
		wgt: wgt,
	}
}

func (e *edge) Source() Vertex {
	return e.src
}

func (e *edge) Target() Vertex {
	return e.tgt
}

func (e *edge) Weight() float64 {
	return e.wgt
}

func (e *edge) String() string {
	return fmt.Sprintf("%s -- %.3f -→ %s\n", e.src, e.wgt, e.tgt)
}

type EdgeSlice []Edge

func (e EdgeSlice) Len() int {
	return len(e)
}
func (e EdgeSlice) Less(i, j int) bool {
	return e[i].Weight() < e[j].Weight()
}
func (e EdgeSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Graph interface {
	// Init initializes a Graph.
	Init()

	// GetVertexCount returns the total number of vertices.
	GetVertexCount() int

	// GetVertex finds the Vertex. It returns nil if the Vertex
	// does not exist in the graph.
	GetVertex(id ID) Vertex

	// GetVertices returns a map from vertex ID to
	// empty struct value. Graph does not allow duplicate
	// vertex ID or name.
	GetVertices() map[ID]Vertex

	// AddVertex adds a vertex to a graph, and returns false
	// if the vertex already existed in the graph.
	AddVertex(nd Vertex) bool

	// DeleteVertex deletes a vertex from a graph.
	// It returns true if it got deleted.
	// And false if it didn't get deleted.
	DeleteVertex(id ID) bool

	// AddEdge adds an edge from nd1 to nd2 with the weight.
	// It returns error if a vertex does not exist.
	AddEdge(id1, id2 ID, weight float64) error

	// ReplaceEdge replaces an edge from id1 to id2 with the weight.
	ReplaceEdge(id1, id2 ID, weight float64) error

	// DeleteEdge deletes an edge from id1 to id2.
	DeleteEdge(id1, id2 ID) error

	// GetWeight returns the weight from id1 to id2.
	GetWeight(id1, id2 ID) (float64, error)

	// GetSources returns the map of parent Vertices.
	// (Vertices that come towards the argument vertex.)
	GetSources(id ID) (map[ID]Vertex, error)

	// GetTargets returns the map of child Vertices.
	// (Vertices that go out of the argument vertex.)
	GetTargets(id ID) (map[ID]Vertex, error)

	// String describes the Graph.
	String() string

}

// graph is an internal default graph type that
// implements all methods in Graph interface.
type graph struct {
	mu              sync.RWMutex // guards the following

				     // idToVertices stores all vertices.
	idToVertices    map[ID]Vertex

				     // vettexToSources maps a Vertex identifer to sources(parents) with edge weights.
	vertexToSources map[ID]map[ID]float64
				     // vertexToTargets maps a Vertex identifer to targets(children) with edge weights.
	vertexToTargets   map[ID]map[ID]float64
}

// newGraph returns a new graph.
func newGraph() *graph {
	return &graph{
		idToVertices:     make(map[ID]Vertex),
		vertexToSources: make(map[ID]map[ID]float64),
		vertexToTargets: make(map[ID]map[ID]float64),
		//
		// without this
		// panic: assignment to entry in nil map
	}
}

// NewGraph returns a new graph.
func NewGraph() Graph {
	return newGraph()
}

func (g *graph) Init() {
	// (X) g = newGraph()
	// this only updates the pointer
	//
	//
	// (X) *g = *newGraph()
	// assignment copies lock value

	g.idToVertices = make(map[ID]Vertex)
	g.vertexToSources = make(map[ID]map[ID]float64)
	g.vertexToTargets = make(map[ID]map[ID]float64)
}

func (g *graph) GetVertexCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return len(g.idToVertices)
}

func (g *graph) GetVertex(id ID) Vertex {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.idToVertices[id]
}

func (g *graph) GetVertices() map[ID]Vertex {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.idToVertices
}

func (g *graph) unsafeExistID(id ID) bool {
	_, ok := g.idToVertices[id]
	return ok
}

func (g *graph) AddVertex(nd Vertex) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.unsafeExistID(nd.ID()) {
		return false
	}

	id := nd.ID()
	g.idToVertices[id] = nd
	return true
}

func (g *graph) DeleteVertex(id ID) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id) {
		return false
	}

	delete(g.idToVertices, id)

	delete(g.vertexToTargets, id)
	for _, smap := range g.vertexToTargets {
		delete(smap, id)
	}

	delete(g.vertexToSources, id)
	for _, smap := range g.vertexToSources {
		delete(smap, id)
	}

	return true
}

func (g *graph) AddEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.vertexToTargets[id1]; ok {
		if v, ok2 := g.vertexToTargets[id1][id2]; ok2 {
			g.vertexToTargets[id1][id2] = v + weight
		} else {
			g.vertexToTargets[id1][id2] = weight
		}
	} else {
		tmap := make(map[ID]float64)
		tmap[id2] = weight
		g.vertexToTargets[id1] = tmap
	}
	if _, ok := g.vertexToSources[id2]; ok {
		if v, ok2 := g.vertexToSources[id2][id1]; ok2 {
			g.vertexToSources[id2][id1] = v + weight
		} else {
			g.vertexToSources[id2][id1] = weight
		}
	} else {
		tmap := make(map[ID]float64)
		tmap[id1] = weight
		g.vertexToSources[id2] = tmap
	}

	return nil
}

func (g *graph) ReplaceEdge(id1, id2 ID, weight float64) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.vertexToTargets[id1]; ok {
		g.vertexToTargets[id1][id2] = weight
	} else {
		tmap := make(map[ID]float64)
		tmap[id2] = weight
		g.vertexToTargets[id1] = tmap
	}
	if _, ok := g.vertexToSources[id2]; ok {
		g.vertexToSources[id2][id1] = weight
	} else {
		tmap := make(map[ID]float64)
		tmap[id1] = weight
		g.vertexToSources[id2] = tmap
	}
	return nil
}

func (g *graph) DeleteEdge(id1, id2 ID) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.unsafeExistID(id1) {
		return fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.vertexToTargets[id1]; ok {
		if _, ok := g.vertexToTargets[id1][id2]; ok {
			delete(g.vertexToTargets[id1], id2)
		}
	}
	if _, ok := g.vertexToSources[id2]; ok {
		if _, ok := g.vertexToSources[id2][id1]; ok {
			delete(g.vertexToSources[id2], id1)
		}
	}
	return nil
}

func (g *graph) GetWeight(id1, id2 ID) (float64, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id1) {
		return 0, fmt.Errorf("%s does not exist in the graph.", id1)
	}
	if !g.unsafeExistID(id2) {
		return 0, fmt.Errorf("%s does not exist in the graph.", id2)
	}

	if _, ok := g.vertexToTargets[id1]; ok {
		if v, ok := g.vertexToTargets[id1][id2]; ok {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf("there is no edge from %s to %s", id1, id2)
}

func (g *graph) GetSources(id ID) (map[ID]Vertex, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	rs := make(map[ID]Vertex)
	if _, ok := g.vertexToSources[id]; ok {
		for n := range g.vertexToSources[id] {
			rs[n] = g.idToVertices[n]
		}
	}
	return rs, nil
}

func (g *graph) GetTargets(id ID) (map[ID]Vertex, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if !g.unsafeExistID(id) {
		return nil, fmt.Errorf("%s does not exist in the graph.", id)
	}

	rs := make(map[ID]Vertex)
	if _, ok := g.vertexToTargets[id]; ok {
		for n := range g.vertexToTargets[id] {
			rs[n] = g.idToVertices[n]
		}
	}
	return rs, nil
}

func (g *graph) String() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	buf := new(bytes.Buffer)
	for id1, nd1 := range g.idToVertices {
		nmap, _ := g.GetTargets(id1)
		for id2, nd2 := range nmap {
			weight, _ := g.GetWeight(id1, id2)
			fmt.Fprintf(buf, "%s -- %.3f -→ %s\n", nd1, weight, nd2)
		}
	}
	return buf.String()
}

// NewGraphFromJSON returns a new Graph from a JSON file.
// Here's the sample JSON data:
//
//	{
//	    "graph_00": {
//	        "S": {
//	            "A": 100,
//	            "B": 14,
//	            "C": 200
//	        },
//	        "A": {
//	            "S": 15,
//	            "B": 5,
//	            "D": 20,
//	            "T": 44
//	        },
//	        "B": {
//	            "S": 14,
//	            "A": 5,
//	            "D": 30,
//	            "E": 18
//	        },
//	        "C": {
//	            "S": 9,
//	            "E": 24
//	        },
//	        "D": {
//	            "A": 20,
//	            "B": 30,
//	            "E": 2,
//	            "F": 11,
//	            "T": 16
//	        },
//	        "E": {
//	            "B": 18,
//	            "C": 24,
//	            "D": 2,
//	            "F": 6,
//	            "T": 19
//	        },
//	        "F": {
//	            "D": 11,
//	            "E": 6,
//	            "T": 6
//	        },
//	        "T": {
//	            "A": 44,
//	            "D": 16,
//	            "F": 6,
//	            "E": 19
//	        }
//	    },
//	}
//
func NewGraphFromJSON(rd io.Reader, graphID string) (Graph, error) {
	js := make(map[string]map[string]map[string]float64)
	dec := json.NewDecoder(rd)
	for {
		if err := dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	if _, ok := js[graphID]; !ok {
		return nil, fmt.Errorf("%s does not exist", graphID)
	}
	gmap := js[graphID]

	g := newGraph()
	for id1, mm := range gmap {
		nd1 := g.GetVertex(StringID(id1))
		if nd1 == nil {
			nd1 = NewVertex(id1)
			if ok := g.AddVertex(nd1); !ok {
				return nil, fmt.Errorf("%s already exists", nd1)
			}
		}
		for id2, weight := range mm {
			nd2 := g.GetVertex(StringID(id2))
			if nd2 == nil {
				nd2 = NewVertex(id2)
				if ok := g.AddVertex(nd2); !ok {
					return nil, fmt.Errorf("%s already exists", nd2)
				}
			}
			g.ReplaceEdge(nd1.ID(), nd2.ID(), weight)
		}
	}

	return g, nil
}
