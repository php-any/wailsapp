package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructPromptForSingleSelectionMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructPromptForSingleSelectionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, err := h.source.PromptForSingleSelection()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return data.NewStringValue(ret0), nil
}

func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetName() string {
	return "promptForSingleSelection"
}
func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *SaveFileDialogStructPromptForSingleSelectionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
