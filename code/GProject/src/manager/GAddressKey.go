//===============================================
package manager
//===============================================
import "strings"
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/gui"
import "github.com/therecipe/qt/core"
//===============================================
// struct
//===============================================
type GAddressKey struct {
    GWidget
    mainLayout *widgets.QHBoxLayout
    widgetId map[widgets.QWidget_ITF]string
}
//===============================================
type GAddressKey_ITF interface {
    GWidget_ITF
    GAddressKey_PTR() *GAddressKey
}
//===============================================
func (obj *GAddressKey) GAddressKey_PTR() *GAddressKey {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGAddressKey(parent widgets.QWidget_ITF) *GAddressKey {
    lObj := &GAddressKey{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GAddressKey")
    
    lObj.widgetId = make(map[widgets.QWidget_ITF]string)
        
    lMainLayout := widgets.NewQHBoxLayout()
    lObj.mainLayout = lMainLayout
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
// methods
//===============================================
func (obj *GAddressKey) SetContent(text string) {
    lApp := GManager().mgr.app
    lSize := GManager().mgr.size

    obj.ClearContent()
    lMap := strings.Split(text, "/")
    lKey := ""
    
    for i,lText := range lMap {
        if i != 0 {
            lSep := widgets.NewQPushButton(nil)
            lSep.SetIcon(gui.NewQIcon5(lApp.icon_map["angle-right"]));
            lSep.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
            obj.mainLayout.AddWidget(lSep, 0, 0)
        }
        
        lButton := widgets.NewQPushButton(nil)
        lButton.SetText(lText)
        obj.mainLayout.AddWidget(lButton, 0, 0)
        
        if i != 0 {lKey += "/"}
        lKey += lText
        obj.widgetId[lButton] = lKey

        lButton.ConnectClicked(obj.SlotItemClicked)
    }
    
    lSpacer := widgets.NewQLabel(nil, 0)
    obj.mainLayout.AddWidget(lSpacer, 1, 0)
}
//===============================================
func (obj *GAddressKey) ClearContent() {
    GManager().ClearLayout(obj.mainLayout)
    GManager().ClearMap(obj.widgetId)
}
//===============================================
// slot
//===============================================
func (obj *GAddressKey) SlotItemClicked(ok bool) {
    lWidget := GManager().GetSender(obj.widgetId)
    lWidgetId := obj.widgetId[lWidget]
    GManager().SetPage(lWidgetId)
}
//===============================================
