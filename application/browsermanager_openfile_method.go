package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type BrowserManagerOpenFileMethod struct {
	source *applicationsrc.BrowserManager
}

func (h *BrowserManagerOpenFileMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("BrowserManager.OpenFile 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	if err := h.source.OpenFile(arg0); err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return nil, nil
}

func (h *BrowserManagerOpenFileMethod) GetName() string            { return "openFile" }
func (h *BrowserManagerOpenFileMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *BrowserManagerOpenFileMethod) GetIsStatic() bool          { return true }
func (h *BrowserManagerOpenFileMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "path", 0, nil, nil),
	}
}

func (h *BrowserManagerOpenFileMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "path", 0, nil),
	}
}

func (h *BrowserManagerOpenFileMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
