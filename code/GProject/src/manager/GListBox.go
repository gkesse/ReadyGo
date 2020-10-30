//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GListBox struct {
	GWidget
    mainLayout *widgets.QVBoxLayout
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
    
    lMainLayout := widgets.NewQVBoxLayout()
    lObj.mainLayout = lMainLayout
    
    lParent.SetLayout(lMainLayout)
        
	return lObj
}
//===============================================
// methods
//===============================================
func (obj *GListBox) AddContent(text string) {
    lButton := widgets.NewQPushButton(nil)
    lButton.SetText(text)
    obj.mainLayout.AddWidget(lButton, 0, 0)
}
//===============================================
