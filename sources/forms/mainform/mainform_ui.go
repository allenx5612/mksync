// WARNING! All changes made in this file will be lost!
package mainform

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIMainformForm struct {
	GridLayoutWidget *widgets.QWidget
	GridLayout *widgets.QGridLayout
	ToolButton3 *widgets.QToolButton
	ToolButton2 *widgets.QToolButton
	ToolButton *widgets.QToolButton
}

func (this *UIMainformForm) SetupUI(Form *widgets.QWidget) {
	Form.SetObjectName("Form")
	Form.SetGeometry(core.NewQRect4(0, 0, 449, 300))
	this.GridLayoutWidget = widgets.NewQWidget(Form, core.Qt__Widget)
	this.GridLayoutWidget.SetObjectName("GridLayoutWidget")
	this.GridLayoutWidget.SetGeometry(core.NewQRect4(0, 0, 451, 51))
	this.GridLayout = widgets.NewQGridLayout(this.GridLayoutWidget)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout.SetSpacing(0)
	this.ToolButton3 = widgets.NewQToolButton(this.GridLayoutWidget)
	this.ToolButton3.SetObjectName("ToolButton3")
	this.GridLayout.AddWidget3(this.ToolButton3, 0, 0, 1, 1, 0)
	this.ToolButton2 = widgets.NewQToolButton(this.GridLayoutWidget)
	this.ToolButton2.SetObjectName("ToolButton2")
	this.GridLayout.AddWidget3(this.ToolButton2, 0, 1, 1, 1, 0)
	this.ToolButton = widgets.NewQToolButton(this.GridLayoutWidget)
	this.ToolButton.SetObjectName("ToolButton")
	this.GridLayout.AddWidget3(this.ToolButton, 0, 2, 1, 1, 0)


    this.RetranslateUi(Form)

}

func (this *UIMainformForm) RetranslateUi(Form *widgets.QWidget) {
    _translate := core.QCoreApplication_Translate
	Form.SetWindowTitle(_translate("Form", "Form", "", -1))
	this.ToolButton3.SetText(_translate("Form", "...", "", -1))
	this.ToolButton2.SetText(_translate("Form", "...", "", -1))
	this.ToolButton.SetText(_translate("Form", "...", "", -1))
}
