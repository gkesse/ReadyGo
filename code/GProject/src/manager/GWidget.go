//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GWidget struct {
	widgets.QFrame
}
//===============================================
// interface
//===============================================
type GWidget_ITF interface {
	widgets.QFrame_ITF
    GWidgetI
	GWidget_PTR() *GWidget
}
//===============================================
type GWidgetI interface {
    AddPage(widget widgets.QWidget_ITF)
    SetContent(text string)
}
//===============================================
// constructor
//===============================================
func (obj *GWidget) GWidget_PTR() *GWidget {
	return obj
}
//===============================================
func NewGWidget(parent widgets.QWidget_ITF) *GWidget {
    lParent := *widgets.NewQFrame(parent, 0)
	return &GWidget {lParent}
}
//===============================================
func CreateGWidget(key string, parent widgets.QWidget_ITF) GWidget_ITF {
    // widget
    if key == "widget" {return NewGWidget(parent)}
    if key == "listbox" {return NewGListBox(parent)}
    if key == "titlebar" {return NewGTitleBar(parent)}
    if key == "addressbar" {return NewGAddressBar(parent)}
    if key == "addresskey" {return NewGAddressKey(parent)}
    //
    if key == "window" {return NewGWindow(parent)}
    //
	return NewGWidget(parent)
}
//===============================================
// methods
//===============================================
func (obj *GWidget) AddPage(widget widgets.QWidget_ITF) {}
func (obj *GWidget) SetContent(text string) {}
//===============================================
