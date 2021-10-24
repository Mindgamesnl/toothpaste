package toothpaste

type RenderContext struct {
	variables map[string]interface{}
}

func NewRenderContext() *RenderContext {
	return &RenderContext{
		variables: map[string]interface{}{},
	}
}

func (r *RenderContext) getVariable(key string) (interface{}, bool) {
	v, f := r.variables[key]
	return v, f
}

func (r *RenderContext) SetVariable(name string, value interface{})  {
	r.variables[name] = value
}
