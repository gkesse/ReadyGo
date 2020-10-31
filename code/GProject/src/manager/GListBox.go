//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
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
    
    lObj.widgetId = make(map[widgets.QWidget_ITF]string)
    
    lContentLayout := widgets.NewQVBoxLayout()
    lObj.contentLayout = lContentLayout
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddLayout(lContentLayout, 0)
    lMainLayout.AddStretch(1)
    
    lParent.SetLayout(lMainLayout)
        
    return lObj
}
//===============================================
// methods
//===============================================
func (obj *GListBox) AddContent(text string, key string) {
    lButton := widgets.NewQPushButton(nil)
    lButton.SetText(text)
    obj.contentLayout.AddWidget(lButton, 0, 0)
    obj.widgetId[lButton] = key
    lButton.ConnectClicked(obj.SlotItemClicked)
}
//===============================================
// slot
//===============================================
func (obj *GListBox) SlotItemClicked(ok bool) {
    lWidget := GManager().GetSender(obj.widgetId)
    lWidgetId := obj.widgetId[lWidget]
    GManager().SetPage(lWidgetId)
}
//===============================================
