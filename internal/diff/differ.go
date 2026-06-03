package diff

import (
	"fmt"
	"github.com/zinuo-xu/prism/internal/model"
)

func Diff(a, b *model.Node) []model.Delta {
	var deltas []model.Delta
	compareNodes("$", a, b, &deltas)
	return deltas
}

func compareNodes(path string, a, b *model.Node, deltas *[]model.Delta) {
	if a == nil && b == nil {
		return
	}
	if a == nil {
		*deltas = append(*deltas, model.Delta{
			Path: path, Type: model.Added, NewValue: b.Value,
			Message: "Field added",
		})
		return
	}
	if b == nil {
		*deltas = append(*deltas, model.Delta{
			Path: path, Type: model.Removed, OldValue: a.Value,
			Message: "Field removed",
		})
		return
	}
	if a.NodeType != b.NodeType {
		*deltas = append(*deltas, model.Delta{
			Path: path, Type: model.TypeChanged,
			OldValue: a.Value, NewValue: b.Value,
			Message: fmt.Sprintf("Type changed"),
		})
		return
	}
	if a.NodeType == model.ValueNode && a.Value != b.Value {
		*deltas = append(*deltas, model.Delta{
			Path: path, Type: model.Changed,
			OldValue: a.Value, NewValue: b.Value,
			Message: fmt.Sprintf("%v -> %v", a.Value, b.Value),
		})
	}
	// Recursively compare children for objects and arrays
	for i := 0; i < len(a.Children) || i < len(b.Children); i++ {
		var childA, childB *model.Node
		childPath := path
		if i < len(a.Children) {
			childA = a.Children[i]
			childPath = path + "." + childA.Key
		}
		if i < len(b.Children) {
			childB = b.Children[i]
			if childA == nil {
				childPath = path + "." + childB.Key
			}
		}
		compareNodes(childPath, childA, childB, deltas)
	}
}
