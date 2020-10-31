//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GAddressBar struct {
    GWidget
    address *widgets.QLineEdit
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
    lQt := GManager().mgr.qt
    
    lObj := &GAddressBar{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GAddressBar")
    
    lIcon := widgets.NewQPushButton(nil)
    lIcon.SetObjectName("icon")
    lIcon.SetText("icon")
    
    lAddress := widgets.NewQLineEdit(nil)
    lObj.address = lAddress
    lQt.address_edit = lAddress
    lAddress.SetObjectName("address")
    
    lGoTo := widgets.NewQPushButton(nil)
    lGoTo.SetObjectName("goto")
    lGoTo.SetText("goto")

    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lIcon, 0, 0)
    lMainLayout.AddWidget(lAddress, 1, 0)
    lMainLayout.AddWidget(lGoTo, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)
    
    lAddress.ConnectReturnPressed(lObj.SlotItemClick)
    lGoTo.ConnectClicked(lObj.SlotItemClicked)
    
    return lObj
}
//===============================================
// slots
//===============================================
func (obj *GAddressBar) SlotItemClick() {
    lAddress := obj.address.Text()
    GManager().SetPage(lAddress)
}
//===============================================
func (obj *GAddressBar) SlotItemClicked(ok bool) {
    obj.SlotItemClick()
}
//===============================================
