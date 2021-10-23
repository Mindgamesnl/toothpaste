package toothpaste

import "sort"

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

func (r *Renderer) Render(renderContext *RenderContext, input string) string {
	return r.RecursiveRender(renderContext, input, nil)
}

func (r *Renderer) RecursiveRender(renderContext *RenderContext, input string, parent *Node) string {
	var elements = findNodesIn(input)

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].start > elements[j].start
	})

	for i := range elements {
		var value, err = elements[i].evaluate(input, renderContext, r, parent)
		if err != nil {
			input = input[:elements[i].start] + "<b>ERROR: " + err.Error() + "</b>" + input[elements[i].end:]
		} else {
			input = input[:elements[i].start] + value + input[elements[i].end:]
		}
	}

	return input
}
