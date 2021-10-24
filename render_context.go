package toothpaste

type RenderContext struct {
	variables map[string]string
}

func NewRenderContext() *RenderContext {
	return &RenderContext{
		variables: map[string]string{},
	}
}

func (r *RenderContext) getVariable(key string) (string, bool) {
	v, f := r.variables[key]
	return v, f
}

func (r *RenderContext) SetVariable(name string, value string)  {
	r.variables[name] = value
}
