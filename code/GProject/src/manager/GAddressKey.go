//===============================================
package manager
//===============================================
import "strings"
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
//===============================================
// struct
//===============================================
type GAddressKey struct {
	GWidget
    mainLayout *widgets.QHBoxLayout
}
//===============================================
// interface
//===============================================
type GAddressKey_ITF interface {
	GWidget_ITF
	GAddressKey_PTR() *GAddressKey
}
//===============================================
// constructor
//===============================================
func (obj *GAddressKey) GAddressKey_PTR() *GAddressKey {
	return obj
}
//===============================================
func NewGAddressKey(parent widgets.QWidget_ITF) *GAddressKey {
    lObj := &GAddressKey{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GAddressKey")
    
    lMainLayout := widgets.NewQHBoxLayout()
    lObj.mainLayout = lMainLayout
    
    lParent.SetLayout(lMainLayout)
    
	return lObj
}
//===============================================
// methods
//===============================================
func (obj *GAddressKey) SetContent(text string) {
    lMap := strings.Split(text, "/")
    for i,lText := range lMap {
        if i != 0 {
            lSep := widgets.NewQPushButton(nil)
            lSep.SetText(">")
            obj.mainLayout.AddWidget(lSep, 0, core.Qt__AlignLeft)
        }
        lButton := widgets.NewQPushButton(nil)
        lButton.SetText(lText)
        obj.mainLayout.AddWidget(lButton, 0, core.Qt__AlignLeft)
    }
}
//===============================================
