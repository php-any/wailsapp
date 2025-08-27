package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSnapAssistMethod struct {
	source applicationsrc.Window
}

func (h *WindowSnapAssistMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.SnapAssist()
	return nil, nil
}

func (h *WindowSnapAssistMethod) GetName() string            { return "snapAssist" }
func (h *WindowSnapAssistMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSnapAssistMethod) GetIsStatic() bool          { return true }
func (h *WindowSnapAssistMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowSnapAssistMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowSnapAssistMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
