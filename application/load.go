package application

import (
	"github.com/php-any/origami/data"
)

func Load(vm data.VM) {
	// 添加顶级函数
	for _, fun := range []data.FuncStmt{
		NewNewFunction(),
		NewNewAppMenuFunction(),
		NewNewMenuFunction(),
		NewNewWindowEventFunction(),
	} {
		vm.AddFunc(fun)
	}

	// 添加类
	vm.AddClass(NewAppClass())
	vm.AddClass(NewAssetOptionsClass())
	vm.AddClass(NewBrowserManagerClass())
	vm.AddClass(NewButtonClass())
	vm.AddClass(NewClipboardManagerClass())
	vm.AddClass(NewContextMenuClass())
	vm.AddClass(NewContextMenuDataClass())
	vm.AddClass(NewContextMenuManagerClass())
	vm.AddClass(NewCustomEventClass())
	vm.AddClass(NewDialogManagerClass())
	vm.AddClass(NewDropZoneDetailsClass())
	vm.AddClass(NewEnvironmentInfoClass())
	vm.AddClass(NewEnvironmentManagerClass())
	vm.AddClass(NewEventManagerClass())
	vm.AddClass(NewKeyBindingManagerClass())
	vm.AddClass(NewLRTBClass())
	vm.AddClass(NewLinuxOptionsClass())
	vm.AddClass(NewLinuxWindowClass())
	vm.AddClass(NewMacLiquidGlassClass())
	vm.AddClass(NewMacOptionsClass())
	vm.AddClass(NewMacTitleBarClass())
	vm.AddClass(NewMacWebviewPreferencesClass())
	vm.AddClass(NewMacWindowClass())
	vm.AddClass(NewMenuClass())
	vm.AddClass(NewMenuBarThemeClass())
	vm.AddClass(NewMenuItemClass())
	vm.AddClass(NewMenuManagerClass())
	vm.AddClass(NewMessageDialogClass())
	vm.AddClass(NewMessageDialogOptionsClass())
	vm.AddClass(NewOpenFileDialogOptionsClass())
	vm.AddClass(NewOpenFileDialogStructClass())
	vm.AddClass(NewOptionsClass())
	vm.AddClass(NewPointClass())
	vm.AddClass(NewRGBAClass())
	vm.AddClass(NewRectClass())
	vm.AddClass(NewSaveFileDialogOptionsClass())
	vm.AddClass(NewSaveFileDialogStructClass())
	vm.AddClass(NewScreenClass())
	vm.AddClass(NewScreenManagerClass())
	vm.AddClass(NewSingleInstanceOptionsClass())
	vm.AddClass(NewSizeClass())
	vm.AddClass(NewSystemTrayClass())
	vm.AddClass(NewSystemTrayManagerClass())
	vm.AddClass(NewTextThemeClass())
	vm.AddClass(NewThemeSettingsClass())
	vm.AddClass(NewWebviewWindowClass())
	vm.AddClass(NewWebviewWindowOptionsClass())
	vm.AddClass(NewWindowEventClass())
	vm.AddClass(NewWindowEventContextClass())
	vm.AddClass(NewWindowManagerClass())
	vm.AddClass(NewWindowThemeClass())
	vm.AddClass(NewWindowsOptionsClass())
	vm.AddClass(NewWindowsWindowClass())
}
