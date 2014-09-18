package entity

type GraphNode struct {
	Id				int
	Data			map[string]interface{} `json:"data"`
	Extensions		map[string]interface{} `json:"extensions"`
}

func (n *GraphNode) GetRelation(int) (*GraphRelation) {
	return &GraphRelation{}
}
