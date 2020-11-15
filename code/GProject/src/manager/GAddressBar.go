//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/gui"
import "github.com/therecipe/qt/core"
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
    lApp := GManager().mgr.app
    lSize := GManager().mgr.size

    lQt := GManager().mgr.qt
    
    lObj := &GAddressBar{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GAddressBar")
    
    lIcon := widgets.NewQPushButton(nil)
    lIcon.SetObjectName("icon")
    lIcon.SetIcon(gui.NewQIcon5(lApp.icon_map["home"]));
    lIcon.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lIcon.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
    
    lAddress := widgets.NewQLineEdit(nil)
    lObj.address = lAddress
    lQt.address_edit = lAddress
    lAddress.SetObjectName("address")
    
    lGoTo := widgets.NewQPushButton(nil)
    lGoTo.SetObjectName("goto")
    lGoTo.SetIcon(gui.NewQIcon5(lApp.icon_map["paper-plane-o"]));
    lGoTo.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lGoTo.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))

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
