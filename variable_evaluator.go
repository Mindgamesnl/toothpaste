package toothpaste

import "C"
import (
	"fmt"
	"strconv"
)

func evaluateVariableValue(value interface{}, c *RenderContext) string {
	switch v := value.(type) {
	default:
		return fmt.Sprintf("unexpected type %T", v)
	case func(ctx *RenderContext) string:
		return value.(func(ctx *RenderContext) string)(c)
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(value.(int))
	case int16:
		return strconv.Itoa(value.(int))
	case int64:
		return strconv.Itoa(value.(int))
	case string:
		return value.(string)
	}
}
