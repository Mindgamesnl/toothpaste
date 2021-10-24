package toothpaste

import (
	"sort"
)

type Renderer struct {
	components map[string]string
}

func NewRenderer() *Renderer {
	return &Renderer{
		components: map[string]string{},
	}
}

func (r *Renderer) RegisterComponent(name string, value string)  {
	r.components[name] = value
}

func (r *Renderer) Render(renderContext *RenderContext, input string) (string, error) {
	// render tree nodes (if, for, etc)
	var treeNodes = findTreeNodes(input)

	var rootTreeNode = &TreeNode{
		start: 0,
		end: len(input),
		hasTextCache: false,
		isRoot: true,
	}

	sort.Slice(treeNodes, func(i, j int) bool {
		return treeNodes[i].start < treeNodes[j].start
	})

	for i := range treeNodes {
		var tn = &treeNodes[i]
		tn.nodeType = tn.parseType(input)
		tn.tag = tn.getText(input)
		if tn.nodeType == TREE_NODE_TYPE_END {
			if len(rootTreeNode.children) == 0 {
				rootTreeNode.contentStart = rootTreeNode.end
				rootTreeNode.contentEnd = tn.start
				rootTreeNode.content = input[(rootTreeNode.contentStart):(rootTreeNode.contentEnd)]
			}
			rootTreeNode = rootTreeNode.parent
			rootTreeNode.addChildren(tn)
		} else {
			tn.parent = rootTreeNode
			rootTreeNode.addChildren(tn)
			rootTreeNode = tn
		}
	}

	var treeResult = rootTreeNode.render(renderContext)
	sort.Slice(treeResult, func(i, j int) bool {
		return treeResult[i].start > treeResult[j].start
	})

	var removedBytes = 0
	for i := range treeResult {
		var r = treeResult[i]
		var a = input
		input = input[:r.start - removedBytes] + r.newValue + a[r.end-removedBytes:]
	}

	// render plain nodes (include, and variable types)
	var plainNodes = findPlainNodes(input)

	sort.Slice(plainNodes, func(i, j int) bool {
		return plainNodes[i].start > plainNodes[j].start
	})

	for i := range plainNodes {
		var value, err = plainNodes[i].evaluate(input, renderContext, r)
		if err != nil {
			input = input[:plainNodes[i].start] + "<b>ERROR: " + err.Error() + "</b>" + input[plainNodes[i].end:]
		} else {
			input = input[:plainNodes[i].start] + value + input[plainNodes[i].end:]
		}
	}
	return input, rootTreeNode.failure
}
