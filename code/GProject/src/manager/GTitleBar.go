//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GTitleBar struct {
	GWidget
}
//===============================================
// interface
//===============================================
type GTitleBar_ITF interface {
	GWidget_ITF
	GTitleBar_PTR() *GTitleBar
}
//===============================================
// constructor
//===============================================
func (obj *GTitleBar) GTitleBar_PTR() *GTitleBar {
	return obj
}
//===============================================
func NewGTitleBar(parent widgets.QWidget_ITF) *GTitleBar {
    lParent := *NewGWidget(parent)
    lParent.SetObjectName("GTitleBar")
    
    lLogo := widgets.NewQPushButton(nil)
    lLogo.SetObjectName("logo")
    lLogo.SetText("logo");
    
    lAppName := widgets.NewQLabel(nil, 0)
    lAppName.SetObjectName("app_name")
    lAppName.SetText("app_name");
        
    lTitle := widgets.NewQLabel(nil, 0)
    lTitle.SetObjectName("title")
    lTitle.SetText("title");
        
    lMinimize := widgets.NewQPushButton(nil)
    lMinimize.SetObjectName("minimize")
    lMinimize.SetText("minimize");

    lMaximize := widgets.NewQPushButton(nil)
    lMaximize.SetObjectName("maximize")
    lMaximize.SetText("maximize");

    lClose := widgets.NewQPushButton(nil)
    lClose.SetObjectName("close")
    lClose.SetText("close");

    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lLogo, 0, 0)
    lMainLayout.AddWidget(lAppName, 0, 0)
    lMainLayout.AddWidget(lTitle, 1, 0)
    lMainLayout.AddWidget(lMinimize, 0, 0)
    lMainLayout.AddWidget(lMaximize, 0, 0)
    lMainLayout.AddWidget(lClose, 0, 0)
    
    lParent.SetLayout(lMainLayout)
    
	return &GTitleBar {lParent}
}
//===============================================
// methods
//===============================================

//===============================================
