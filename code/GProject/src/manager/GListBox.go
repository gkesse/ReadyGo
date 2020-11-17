//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/gui"
//===============================================
// struct
//===============================================
type GListBox struct {
    GWidget
    contentLayout *widgets.QVBoxLayout
    widgetId map[widgets.QWidget_ITF]string
}
//===============================================
type GListBox_ITF interface {
    GWidget_ITF
    GListBox_PTR() *GListBox
}
//===============================================
func (obj *GListBox) GListBox_PTR() *GListBox {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGListBox(parent widgets.QWidget_ITF) *GListBox {   
    lObj := &GListBox{}

    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GListBox")

    lObj.widgetId = make(map[widgets.QWidget_ITF]string)
    
    lContentLayout := widgets.NewQVBoxLayout()
    lObj.contentLayout = lContentLayout
    lContentLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lContentLayout.SetSpacing(5)

    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddLayout(lContentLayout, 0)
    lMainLayout.AddStretch(1)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
    lParent.SetLayout(lMainLayout)
        
    return lObj
}
//===============================================
// methods
//===============================================
func (obj *GListBox) AddContent(text string, key string) {
    lApp := GManager().mgr.app
    lSize := GManager().mgr.size

    lButton := widgets.NewQPushButton(nil)
    lButton.SetObjectName("item")
    lButton.SetText(text)
    lButton.SetIcon(gui.NewQIcon5(lApp.icon_map["book"]));
    lButton.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lButton.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
    obj.contentLayout.AddWidget(lButton, 0, 0)
    obj.widgetId[lButton] = key
    lButton.ConnectClicked(obj.SlotItemClicked)
}
//===============================================
// slot
//===============================================
func (obj *GListBox) SlotItemClicked(ok bool) {
    lQt := GManager().mgr.qt 
    lWidget := GManager().GetSender(obj.widgetId)
    lQt.widget_id = obj.widgetId[lWidget]
    //obj.EmitItemClick()
}
//===============================================
