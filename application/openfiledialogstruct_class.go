package application

import (
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

func NewOpenFileDialogStructClass() data.ClassStmt {
	return &OpenFileDialogStructClass{
		source:                          nil,
		construct:                       &OpenFileDialogStructConstructMethod{source: nil},
		addFilter:                       &OpenFileDialogStructAddFilterMethod{source: nil},
		allowsOtherFileTypes:            &OpenFileDialogStructAllowsOtherFileTypesMethod{source: nil},
		attachToWindow:                  &OpenFileDialogStructAttachToWindowMethod{source: nil},
		canChooseDirectories:            &OpenFileDialogStructCanChooseDirectoriesMethod{source: nil},
		canChooseFiles:                  &OpenFileDialogStructCanChooseFilesMethod{source: nil},
		canCreateDirectories:            &OpenFileDialogStructCanCreateDirectoriesMethod{source: nil},
		canSelectHiddenExtension:        &OpenFileDialogStructCanSelectHiddenExtensionMethod{source: nil},
		hideExtension:                   &OpenFileDialogStructHideExtensionMethod{source: nil},
		promptForMultipleSelection:      &OpenFileDialogStructPromptForMultipleSelectionMethod{source: nil},
		promptForSingleSelection:        &OpenFileDialogStructPromptForSingleSelectionMethod{source: nil},
		resolvesAliases:                 &OpenFileDialogStructResolvesAliasesMethod{source: nil},
		setButtonText:                   &OpenFileDialogStructSetButtonTextMethod{source: nil},
		setDirectory:                    &OpenFileDialogStructSetDirectoryMethod{source: nil},
		setMessage:                      &OpenFileDialogStructSetMessageMethod{source: nil},
		setOptions:                      &OpenFileDialogStructSetOptionsMethod{source: nil},
		setTitle:                        &OpenFileDialogStructSetTitleMethod{source: nil},
		showHiddenFiles:                 &OpenFileDialogStructShowHiddenFilesMethod{source: nil},
		treatsFilePackagesAsDirectories: &OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod{source: nil},
	}
}

func NewOpenFileDialogStructClassFrom(source *applicationsrc.OpenFileDialogStruct) data.ClassStmt {
	return &OpenFileDialogStructClass{
		source:                          source,
		construct:                       &OpenFileDialogStructConstructMethod{source: source},
		addFilter:                       &OpenFileDialogStructAddFilterMethod{source: source},
		allowsOtherFileTypes:            &OpenFileDialogStructAllowsOtherFileTypesMethod{source: source},
		attachToWindow:                  &OpenFileDialogStructAttachToWindowMethod{source: source},
		canChooseDirectories:            &OpenFileDialogStructCanChooseDirectoriesMethod{source: source},
		canChooseFiles:                  &OpenFileDialogStructCanChooseFilesMethod{source: source},
		canCreateDirectories:            &OpenFileDialogStructCanCreateDirectoriesMethod{source: source},
		canSelectHiddenExtension:        &OpenFileDialogStructCanSelectHiddenExtensionMethod{source: source},
		hideExtension:                   &OpenFileDialogStructHideExtensionMethod{source: source},
		promptForMultipleSelection:      &OpenFileDialogStructPromptForMultipleSelectionMethod{source: source},
		promptForSingleSelection:        &OpenFileDialogStructPromptForSingleSelectionMethod{source: source},
		resolvesAliases:                 &OpenFileDialogStructResolvesAliasesMethod{source: source},
		setButtonText:                   &OpenFileDialogStructSetButtonTextMethod{source: source},
		setDirectory:                    &OpenFileDialogStructSetDirectoryMethod{source: source},
		setMessage:                      &OpenFileDialogStructSetMessageMethod{source: source},
		setOptions:                      &OpenFileDialogStructSetOptionsMethod{source: source},
		setTitle:                        &OpenFileDialogStructSetTitleMethod{source: source},
		showHiddenFiles:                 &OpenFileDialogStructShowHiddenFilesMethod{source: source},
		treatsFilePackagesAsDirectories: &OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod{source: source},
	}
}

type OpenFileDialogStructClass struct {
	node.Node
	source                          *applicationsrc.OpenFileDialogStruct
	construct                       data.Method
	addFilter                       data.Method
	allowsOtherFileTypes            data.Method
	attachToWindow                  data.Method
	canChooseDirectories            data.Method
	canChooseFiles                  data.Method
	canCreateDirectories            data.Method
	canSelectHiddenExtension        data.Method
	hideExtension                   data.Method
	promptForMultipleSelection      data.Method
	promptForSingleSelection        data.Method
	resolvesAliases                 data.Method
	setButtonText                   data.Method
	setDirectory                    data.Method
	setMessage                      data.Method
	setOptions                      data.Method
	setTitle                        data.Method
	showHiddenFiles                 data.Method
	treatsFilePackagesAsDirectories data.Method
}

func (s *OpenFileDialogStructClass) GetValue(ctx data.Context) (data.GetValue, data.Control) {
	return data.NewProxyValue(NewOpenFileDialogStructClassFrom(&applicationsrc.OpenFileDialogStruct{}), ctx.CreateBaseContext()), nil
}
func (s *OpenFileDialogStructClass) GetName() string {
	return "wails\\application\\OpenFileDialogStruct"
}
func (s *OpenFileDialogStructClass) GetExtend() *string                         { return nil }
func (s *OpenFileDialogStructClass) GetImplements() []string                    { return nil }
func (s *OpenFileDialogStructClass) AsString() string                           { return "OpenFileDialogStruct{}" }
func (s *OpenFileDialogStructClass) GetSource() any                             { return s.source }
func (s *OpenFileDialogStructClass) GetProperty(_ string) (data.Property, bool) { return nil, false }
func (s *OpenFileDialogStructClass) GetProperties() map[string]data.Property    { return nil }
func (s *OpenFileDialogStructClass) GetMethod(name string) (data.Method, bool) {
	switch name {
	case "addFilter":
		return s.addFilter, true
	case "allowsOtherFileTypes":
		return s.allowsOtherFileTypes, true
	case "attachToWindow":
		return s.attachToWindow, true
	case "canChooseDirectories":
		return s.canChooseDirectories, true
	case "canChooseFiles":
		return s.canChooseFiles, true
	case "canCreateDirectories":
		return s.canCreateDirectories, true
	case "canSelectHiddenExtension":
		return s.canSelectHiddenExtension, true
	case "hideExtension":
		return s.hideExtension, true
	case "promptForMultipleSelection":
		return s.promptForMultipleSelection, true
	case "promptForSingleSelection":
		return s.promptForSingleSelection, true
	case "resolvesAliases":
		return s.resolvesAliases, true
	case "setButtonText":
		return s.setButtonText, true
	case "setDirectory":
		return s.setDirectory, true
	case "setMessage":
		return s.setMessage, true
	case "setOptions":
		return s.setOptions, true
	case "setTitle":
		return s.setTitle, true
	case "showHiddenFiles":
		return s.showHiddenFiles, true
	case "treatsFilePackagesAsDirectories":
		return s.treatsFilePackagesAsDirectories, true
	}
	return nil, false
}

func (s *OpenFileDialogStructClass) GetMethods() []data.Method {
	return []data.Method{
		s.addFilter,
		s.allowsOtherFileTypes,
		s.attachToWindow,
		s.canChooseDirectories,
		s.canChooseFiles,
		s.canCreateDirectories,
		s.canSelectHiddenExtension,
		s.hideExtension,
		s.promptForMultipleSelection,
		s.promptForSingleSelection,
		s.resolvesAliases,
		s.setButtonText,
		s.setDirectory,
		s.setMessage,
		s.setOptions,
		s.setTitle,
		s.showHiddenFiles,
		s.treatsFilePackagesAsDirectories,
	}
}

func (s *OpenFileDialogStructClass) GetConstruct() data.Method { return s.construct }

type OpenFileDialogStructConstructMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructConstructMethod) Call(ctx data.Context) (data.GetValue, data.Control) {
	return nil, nil
}

func (h *OpenFileDialogStructConstructMethod) GetName() string            { return "construct" }
func (h *OpenFileDialogStructConstructMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *OpenFileDialogStructConstructMethod) GetIsStatic() bool          { return false }
func (h *OpenFileDialogStructConstructMethod) GetParams() []data.GetValue { return []data.GetValue{} }
func (h *OpenFileDialogStructConstructMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}
func (h *OpenFileDialogStructConstructMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
