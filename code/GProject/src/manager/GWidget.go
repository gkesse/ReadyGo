//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import _ "github.com/therecipe/qt/core"
//===============================================
// struct
//===============================================
type GWidget struct {
    widgets.QFrame
    // signals
    _ func() `signal:"emitItemClick,auto"`
}
//===============================================
// interface
//===============================================
type GWidget_ITF interface {
    widgets.QFrame_ITF
    GWidget_PTR() *GWidget
    // methods
    AddPage(widget widgets.QWidget_ITF, key string, title string, defaultId int)
    AddContent(key string, text string)
    AddBuilder(key string)
    AddScroll(widget widgets.QWidget_ITF)
    SetContent(text string)
    ClearContent()
    // slots
    SlotItemClick()
    SlotItemClicked(ok bool)
    SlotTextChanged(text string)
}
//===============================================
// constructor
//===============================================
func (obj *GWidget) GWidget_PTR() *GWidget {
    return obj
}
//===============================================
func NewGWidget(parent widgets.QWidget_ITF) *GWidget {
    lObj := &GWidget{}
    
    lParent := *widgets.NewQFrame(parent, 0)
    lObj.QFrame = lParent
    
    return lObj
}
//===============================================
func CreateGWidget(key string, parent widgets.QWidget_ITF) GWidget_ITF {
    // widget
    if key == "widget" {return NewGWidget(parent)}
    if key == "listbox" {return NewGListBox(parent)}
    if key == "titlebar" {return NewGTitleBar(parent)}
    if key == "addressbar" {return NewGAddressBar(parent)}
    if key == "addresskey" {return NewGAddressKey(parent)}
    if key == "databaseview" {return NewGDatabaseView(parent)}
    if key == "hscroll" {return NewGHScroll(parent)}
    if key == "vscroll" {return NewGVScroll(parent)}
    // window
    if key == "window" {return NewGWindow(parent)}
    if key == "home" {return NewGHome(parent)}
    if key == "builder" {return NewGBuilder(parent)}
    //
    return NewGWidget(parent)
}
//===============================================
// methods
//===============================================
func (obj *GWidget) AddPage(widget widgets.QWidget_ITF, key string, title string, defaultId int) {}
func (obj *GWidget) AddContent(text string, key string) {}
func (obj *GWidget) AddBuilder(key string) {}
func (obj *GWidget) AddScroll(widget widgets.QWidget_ITF) {}
func (obj *GWidget) SetContent(text string) {}
func (obj *GWidget) ClearContent() {}
//===============================================
// slots
//===============================================
func (obj *GWidget) SlotItemClick() {}
func (obj *GWidget) SlotItemClicked(ok bool) {}
func (obj *GWidget) SlotTextChanged(text string) {}
//===============================================
