package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type NewFunction struct{}

func NewNewFunction() data.FuncStmt {
	return &NewFunction{}
}

func (h *NewFunction) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("New 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Options
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Options); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Options)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("New 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Options); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Options)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("New 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(applicationsrc.Options)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("New 参数类型不支持, index: 0"))
	}

	ret0 := applicationsrc.New(arg0)
	return data.NewClassValue(NewAppClassFrom(ret0), ctx), nil
}

func (h *NewFunction) GetName() string            { return "wails\\application\\new" }
func (h *NewFunction) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *NewFunction) GetIsStatic() bool          { return true }
func (h *NewFunction) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "appOptions", 0, nil, nil),
	}
}
func (h *NewFunction) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "appOptions", 0, nil),
	}
}
func (h *NewFunction) GetReturnType() data.Types { return data.NewBaseType("void") }
