package parser

import (
	"encoding/json"
	"fmt"
	"github.com/zinuo-xu/prism/internal/model"
)

type JSONParser struct{}

func (p *JSONParser) Parse(data []byte) (*model.Node, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return convertToNode("$", v), nil
}

func convertToNode(key string, v interface{}) *model.Node {
	switch val := v.(type) {
	case map[string]interface{}:
		node := &model.Node{Key: key, NodeType: model.ObjectNode}
		for k, childV := range val {
			node.Children = append(node.Children, convertToNode(k, childV))
		}
		return node
	case []interface{}:
		node := &model.Node{Key: key, NodeType: model.ArrayNode}
		for i, childV := range val {
			node.Children = append(node.Children,
				convertToNode(fmt.Sprintf("[%d]", i), childV))
		}
		return node
	default:
		return &model.Node{Key: key, NodeType: model.ValueNode, Value: val}
	}
}
