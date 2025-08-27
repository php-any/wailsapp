package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructResolvesAliasesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructResolvesAliasesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.ResolvesAliases 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.ResolvesAliases(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructResolvesAliasesMethod) GetName() string { return "resolvesAliases" }
func (h *OpenFileDialogStructResolvesAliasesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructResolvesAliasesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructResolvesAliasesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "resolvesAliases", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructResolvesAliasesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "resolvesAliases", 0, nil),
	}
}

func (h *OpenFileDialogStructResolvesAliasesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
