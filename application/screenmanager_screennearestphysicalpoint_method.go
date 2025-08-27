package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenManagerScreenNearestPhysicalPointMethod struct {
	source *applicationsrc.ScreenManager
}

func (h *ScreenManagerScreenNearestPhysicalPointMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.ScreenNearestPhysicalPoint 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.ScreenNearestPhysicalPoint 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.ScreenNearestPhysicalPoint 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(applicationsrc.Point)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.ScreenNearestPhysicalPoint 参数类型不支持, index: 0"))
	}

	ret0 := h.source.ScreenNearestPhysicalPoint(arg0)
	return data.NewClassValue(NewScreenClassFrom(ret0), ctx), nil
}

func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetName() string {
	return "screenNearestPhysicalPoint"
}
func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetIsStatic() bool { return true }
func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "physicalPoint", 0, nil, nil),
	}
}

func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "physicalPoint", 0, nil),
	}
}

func (h *ScreenManagerScreenNearestPhysicalPointMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
