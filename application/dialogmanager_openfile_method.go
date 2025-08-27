package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerOpenFileMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerOpenFileMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.OpenFile()
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *DialogManagerOpenFileMethod) GetName() string            { return "openFile" }
func (h *DialogManagerOpenFileMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerOpenFileMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerOpenFileMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerOpenFileMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerOpenFileMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
