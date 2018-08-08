// This file is not a data structure or algorithm about graph, it's a helper data structure to implement graph algorithms.

package graph

// a navie and ad-hoc implementation of disjoint set
type disjointSet struct {
	// elem -> set
	set map[*Vertex]*Vertex
}

func (s *disjointSet) makeSet(v *Vertex) {
	s.set[v] = v
}

func (s *disjointSet) findSet(v *Vertex) *Vertex {
	if set, ok := s.set[v]; ok {
		return set
	}
	return nil
}

func (s *disjointSet) union(u, v *Vertex) {
	setU, setV := s.findSet(u), s.findSet(v)

	// merge setV to setU
	for elem, set := range s.set {
		if set == setV {
			s.set[elem] = setU
		}
	}
}
