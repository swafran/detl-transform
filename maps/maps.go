package maps

type MapNode struct {
	Name     string
	NodeType string
	Value    string
	MappedTo []string
	Filters  []string
	Children []*MapNode
	Parent   *MapNode
}

type MapTree struct {
	Root *MapNode
}

type path struct {
	levelsUp int
	pathDown []string
	nodeName string
}

func (parent *MapNode) AppendChild(child *MapNode) *MapNode {
	return child
}

func (m *MapTree) AppendAt(child *MapNode, path path) *MapNode {
	return &MapNode{}
}

func (m *MapTree) find(path) *MapNode {
	return &MapNode{}
}

func (m *MapTree) parsePath(pathString string) (path path) {
	return
}

func (m *MapTree) place(n *MapNode, parent *MapNode) {

}
