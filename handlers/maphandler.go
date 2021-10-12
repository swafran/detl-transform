package handlers

import (
	"fmt"
	"time"

	"github.com/swafran/detl-common/parsers"
	"github.com/swafran/detl-common/system"
	"github.com/swafran/detl-transform/maps"
)

//MapHandler is the default handler, applies map to data; implements Handler
type MapHandler struct {
	Mapping map[string]interface{}
	Parser  parsers.Parser
	MapTree *maps.MapTree
}

//Handle processes messages
func (h *MapHandler) Handle(m string) {

	for system.IsMaxedOut() {

		//TODO alert and log that we're blocking

		time.Sleep(5 * time.Second)
	}

	go h.resolve(m)
}

func (h *MapHandler) resolve(m string) {
	input := h.Parser.Parse(m)

	n := maps.MapNode{
		Name:     "root",
		NodeType: "object",
	}

	h.traverse(input, &n, h.Mapping)
}

func (h *MapHandler) traverse(input interface{}, n *maps.MapNode,
	mappingLevel interface{}) {

	parent := n
	n = &maps.MapNode{}

	switch v := input.(type) {
	case bool:
		//
	case float64:
		fmt.Println(v)
	case string:
		n.Name = input.(string)
		n.NodeType = "string"
		n.Value = input.(string)

	case []interface{}:

		mappingLevel = mappingLevel.([]interface{})

		// (json arrays)
	case map[string]interface{}:
		// json objects
	case nil:
		// json null
	default:
		// shouldn't get here; TODO decide on error handling
	}

	if len(n.Children) > 0 {
		h.placeTargets(n, mappingLevel)
	} else {
		parent.AppendChild(n)
	}

	h.applyFilters(n, mappingLevel)

	//h.MapTree.place(&n, &parent)

}

func (h *MapHandler) applyFilters(n *maps.MapNode, mappingLevel interface{}) {}
func (h *MapHandler) placeTargets(n *maps.MapNode, mappingLevel interface{}) {}
