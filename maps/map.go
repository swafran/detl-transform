package maps

type MapNode struct {
	name     string
	nodeType string
	value    string
	mappedTo []string
	filters  []string
	children []*MapNode
	parent   *MapNode
}

type Map struct {
	root *MapNode
}

type path struct {
	levelsUp int
	pathDown []string
	nodeName string
}

func (parent *MapNode) AppendChild(child *MapNode) *MapNode {
	return child
}

func (m *Map) AppendAt(child *MapNode, path path) *MapNode {}

func (m *Map) find(path) *MapNode {}

func (m *Map) parsePath(pathString string) (path path) {
	return
}
