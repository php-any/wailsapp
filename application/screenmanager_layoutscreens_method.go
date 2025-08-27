package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ScreenManagerLayoutScreensMethod struct {
	source *applicationsrc.ScreenManager
}

func (h *ScreenManagerLayoutScreensMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.LayoutScreens 缺少参数, index: 0"))
	}

	arg0 := make([]*applicationsrc.Screen, 0)
	for _, v := range a0.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(**applicationsrc.Screen); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(*applicationsrc.Screen))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.LayoutScreens 参数类型不支持, index: 0"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(**applicationsrc.Screen); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(*applicationsrc.Screen))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("ScreenManager.LayoutScreens 参数类型不支持, index: 0"))
			}
		case *data.AnyValue:
			arg0 = append(arg0, elemVal.Value.(*applicationsrc.Screen))
		}
	}

	if err := h.source.LayoutScreens(arg0); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *ScreenManagerLayoutScreensMethod) GetName() string            { return "layoutScreens" }
func (h *ScreenManagerLayoutScreensMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ScreenManagerLayoutScreensMethod) GetIsStatic() bool          { return true }
func (h *ScreenManagerLayoutScreensMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "screens", 0, nil, nil),
	}
}

func (h *ScreenManagerLayoutScreensMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "screens", 0, nil),
	}
}

func (h *ScreenManagerLayoutScreensMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
