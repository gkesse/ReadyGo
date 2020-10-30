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
    AddContent(text string)
    SetContent(text string)
    ClearContent()
    SlotItemClicked(ok bool)
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
    // window
    if key == "window" {return NewGWindow(parent)}
    if key == "home" {return NewGHome(parent)}
    //
	return NewGWidget(parent)
}
//===============================================
// methods
//===============================================
func (obj *GWidget) AddPage(widget widgets.QWidget_ITF) {}
func (obj *GWidget) AddContent(text string) {}
func (obj *GWidget) SetContent(text string) {}
func (obj *GWidget) ClearContent() {}
//===============================================
func (obj *GWidget) SlotItemClicked(ok bool) {}
//===============================================
