package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructPromptForSingleSelectionMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructPromptForSingleSelectionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, err := h.source.PromptForSingleSelection()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return data.NewStringValue(ret0), nil
}

func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetName() string {
	return "promptForSingleSelection"
}
func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *OpenFileDialogStructPromptForSingleSelectionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
