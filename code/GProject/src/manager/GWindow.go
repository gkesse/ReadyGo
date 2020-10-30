//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GWindow struct {
	GWidget
    workspace *widgets.QStackedWidget
}
//===============================================
// interface
//===============================================
type GWindow_ITF interface {
	GWidget_ITF
	GWindow_PTR() *GWindow
}
//===============================================
// constructor
//===============================================
func (obj *GWindow) GWindow_PTR() *GWindow {
	return obj
}
//===============================================
func NewGWindow(parent widgets.QWidget_ITF) *GWindow {
    lObj := &GWindow {}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GWindow")
    
    lTitleBar := CreateGWidget("titlebar", nil)
    lAddressBar := CreateGWidget("addressbar", nil)
    
    lAddressKey := CreateGWidget("addresskey", nil)
    lAddressKey.SetContent("home/builder/opencv")
    
    lWorkspace := widgets.NewQStackedWidget(nil)
    lObj.workspace = lWorkspace
    
    lObj.AddPage(CreateGWidget("titlebar", nil))
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lTitleBar, 0, 0)
    lMainLayout.AddWidget(lAddressBar, 0, 0)
    lMainLayout.AddWidget(lAddressKey, 0, 0)
    lMainLayout.AddWidget(lWorkspace, 1, 0)

    lParent.SetLayout(lMainLayout)
    
	return lObj
}
//===============================================
// methods
//===============================================
func (obj *GWindow) AddPage(widget widgets.QWidget_ITF) {
    obj.workspace.AddWidget(widget)
}
//===============================================
