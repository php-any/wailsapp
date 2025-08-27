package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type RectContainsMethod struct {
	source *applicationsrc.Rect
}

func (h *RectContainsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Rect.Contains 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Point
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Point); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Point)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("Rect.Contains 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Point); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Point)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("Rect.Contains 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(applicationsrc.Point)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("Rect.Contains 参数类型不支持, index: 0"))
	}

	ret0 := h.source.Contains(arg0)
	return data.NewBoolValue(ret0), nil
}

func (h *RectContainsMethod) GetName() string            { return "contains" }
func (h *RectContainsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *RectContainsMethod) GetIsStatic() bool          { return true }
func (h *RectContainsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *RectContainsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *RectContainsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
