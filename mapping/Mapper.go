package mapping

type HierarchyNode interface {
	AppendChild(interface{}) interface{}
}

type MappingNode struct {
	name     string
	nodeType string
	value    string
	mappedTo []string
	filters  []string
	children []*MappingNode
	parent   *MappingNode
}

type Mapping struct {
	root *MappingNode
}

type path struct {
	levelsUp int
	pathDown []string
	nodeName string
}

func (parent *MappingNode) AppendChild(child *MappingNode) *MappingNode {
	return child
}

func (m *Mapping) AppendAt(child *MappingNode, path path) *MappingNode {}

func (m *Mapping) find(path) *MappingNode {}

func (m *Mapping) parsePath(pathString string) (path path) {
	return
}
