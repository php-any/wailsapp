package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type BrowserManagerOpenURLMethod struct {
	source *applicationsrc.BrowserManager
}

func (h *BrowserManagerOpenURLMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("BrowserManager.OpenURL 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	if err := h.source.OpenURL(arg0); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *BrowserManagerOpenURLMethod) GetName() string            { return "openURL" }
func (h *BrowserManagerOpenURLMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *BrowserManagerOpenURLMethod) GetIsStatic() bool          { return true }
func (h *BrowserManagerOpenURLMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "url", 0, nil, nil),
	}
}

func (h *BrowserManagerOpenURLMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "url", 0, nil),
	}
}

func (h *BrowserManagerOpenURLMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
