//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GAddressBar struct {
    GWidget
}
//===============================================
type GAddressBar_ITF interface {
    GWidget_ITF
    GAddressBar_PTR() *GAddressBar
}
//===============================================
func (obj *GAddressBar) GAddressBar_PTR() *GAddressBar {
    return obj
}
//===============================================
// constructor
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
    lMainLayout.AddWidget(lAddress, 1, 0)
    lMainLayout.AddWidget(lGoTo, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
// methods
//===============================================
