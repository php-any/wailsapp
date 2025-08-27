package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerSaveFileMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerSaveFileMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.SaveFile()
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *DialogManagerSaveFileMethod) GetName() string            { return "saveFile" }
func (h *DialogManagerSaveFileMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerSaveFileMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerSaveFileMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerSaveFileMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerSaveFileMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
