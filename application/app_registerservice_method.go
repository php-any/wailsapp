package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type AppRegisterServiceMethod struct {
	source *applicationsrc.App
}

func (h *AppRegisterServiceMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("App.RegisterService 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Service
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Service); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Service)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("App.RegisterService 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Service); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Service)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("App.RegisterService 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(applicationsrc.Service)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("App.RegisterService 参数类型不支持, index: 0"))
	}

	h.source.RegisterService(arg0)
	return nil, nil
}

func (h *AppRegisterServiceMethod) GetName() string            { return "registerService" }
func (h *AppRegisterServiceMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *AppRegisterServiceMethod) GetIsStatic() bool          { return true }
func (h *AppRegisterServiceMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "service", 0, nil, nil),
	}
}

func (h *AppRegisterServiceMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "service", 0, nil),
	}
}

func (h *AppRegisterServiceMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
