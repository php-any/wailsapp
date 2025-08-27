package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructPromptForMultipleSelectionMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, err := h.source.PromptForMultipleSelection()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	return data.NewAnyValue(ret0), nil
}

func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetName() string {
	return "promptForMultipleSelection"
}
func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *OpenFileDialogStructPromptForMultipleSelectionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
