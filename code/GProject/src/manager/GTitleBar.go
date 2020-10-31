//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
//===============================================
// struct
//===============================================
type GTitleBar struct {
    GWidget
}
//===============================================
type GTitleBar_ITF interface {
    GWidget_ITF
    GTitleBar_PTR() *GTitleBar
}
//===============================================
func (obj *GTitleBar) GTitleBar_PTR() *GTitleBar {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGTitleBar(parent widgets.QWidget_ITF) *GTitleBar {
    lObj := &GTitleBar{}
    
    lApp := GManager().mgr.app
    lQt := GManager().mgr.qt
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GTitleBar")
    
    lLogo := widgets.NewQPushButton(nil)
    lLogo.SetObjectName("logo")
    lLogo.SetText("logo");
    
    lAppName := widgets.NewQLabel(nil, 0)
    lAppName.SetObjectName("app_name")
    lAppName.SetText(lApp.app_name);
        
    lTitle := widgets.NewQLabel(nil, 0)
    lQt.title_id = lTitle
    lTitle.SetObjectName("title")
        
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
    lMainLayout.AddWidget(lTitle, 1, core.Qt__AlignCenter)
    lMainLayout.AddWidget(lMinimize, 0, 0)
    lMainLayout.AddWidget(lMaximize, 0, 0)
    lMainLayout.AddWidget(lClose, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
