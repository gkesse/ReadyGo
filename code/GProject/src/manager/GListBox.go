//===============================================
package manager
//===============================================
import "fmt"
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/gui"
//===============================================
// struct
//===============================================
type GListBox struct {
	GWidget
}
//===============================================
// interface
//===============================================
type GListBox_ITF interface {
	GWidget_ITF
	GListBox_PTR() *GListBox
}
//===============================================
// constructor
//===============================================
func (obj *GListBox) GListBox_PTR() *GListBox {
	return obj
}
//===============================================
func NewGListBox(parent widgets.QWidget_ITF) *GListBox {
    lParent := *NewGWidget(parent)
    
    lLabel := widgets.NewQLabel(nil, 0)
    lLabel.SetText("Bonjour");
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.Layout().AddWidget(lLabel)
    
    lParent.SetLayout(lMainLayout)
    
    lParent.ConnectMousePressEvent(mousePressEvent)
    
	return &GListBox {lParent}
}
//===============================================
// methods
//===============================================
func mousePressEvent(event *gui.QMouseEvent) {
    fmt.Printf("\n### mousePressEvent : GListBox\n\n")
}
//===============================================
