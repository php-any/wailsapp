package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ButtonOnClickMethod struct {
	source *applicationsrc.Button
}

func (h *ButtonOnClickMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Button.OnClick 缺少参数, index: 0"))
	}

	var arg0 func()
	switch fnv := a0.(type) {
	case *data.FuncValue:
		arg0 = func() { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("Button.OnClick 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OnClick(arg0)
	return data.NewClassValue(NewButtonClassFrom(ret0), ctx), nil
}

func (h *ButtonOnClickMethod) GetName() string            { return "onClick" }
func (h *ButtonOnClickMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ButtonOnClickMethod) GetIsStatic() bool          { return true }
func (h *ButtonOnClickMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "callback", 0, nil, nil),
	}
}

func (h *ButtonOnClickMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "callback", 0, nil),
	}
}

func (h *ButtonOnClickMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
