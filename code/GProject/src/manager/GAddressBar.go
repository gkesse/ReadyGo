//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
//===============================================
// struct
//===============================================
type GAddressBar struct {
	GWidget
}
//===============================================
// interface
//===============================================
type GAddressBar_ITF interface {
	GWidget_ITF
	GAddressBar_PTR() *GAddressBar
}
//===============================================
// constructor
//===============================================
func (obj *GAddressBar) GAddressBar_PTR() *GAddressBar {
	return obj
}
//===============================================
func NewGAddressBar(parent widgets.QWidget_ITF) *GAddressBar {
    lObj := &GAddressBar{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GAddressBar")
    
    lIcon := widgets.NewQPushButton(nil)
    lIcon.SetObjectName("icon")
    lIcon.SetText("icon")
    
    lAddress := widgets.NewQLineEdit(nil)
    lAddress.SetObjectName("address")
    
    lGoTo := widgets.NewQPushButton(nil)
    lGoTo.SetObjectName("icon")
    lGoTo.SetText("icon")

    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lIcon, 0, 0)
    lMainLayout.AddWidget(lAddress, 1, core.Qt__AlignCenter)
    lMainLayout.AddWidget(lGoTo, 0, 0)
    
    lParent.SetLayout(lMainLayout)
    
	return lObj
}
//===============================================
// methods
//===============================================
