//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GWindow struct {
	GWidget
}
//===============================================
type GWindow_ITF interface {
	GWidget_ITF
	GWindow_PTR() *GWindow
}
//===============================================
func (obj *GWindow) GWindow_PTR() *GWindow {
	return obj
}
//===============================================
func NewGWindow(parent widgets.QWidget_ITF) *GWindow {
    lQt := GManager().mgr.qt
    
    lObj := &GWindow {}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GWindow")
    
    lTitleBar := CreateGWidget("titlebar", nil)
    lAddressBar := CreateGWidget("addressbar", nil)
    
    lAddressKey := CreateGWidget("addresskey", nil)
    lQt.address_key = lAddressKey
    
    lWorkspace := widgets.NewQStackedWidget(nil)
    lQt.page_map = lWorkspace
    
    lObj.AddPage(CreateGWidget("home", nil), "home", 1)
    lObj.AddPage(CreateGWidget("qt", nil), "home/qt", 0)
    lObj.AddPage(CreateGWidget("sqlite", nil), "home/sqlite", 0)
    lObj.AddPage(CreateGWidget("opencv", nil), "home/opencv", 0)
    lObj.AddPage(CreateGWidget("builder", nil), "home/builder", 0)
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lTitleBar, 0, 0)
    lMainLayout.AddWidget(lAddressBar, 0, 0)
    lMainLayout.AddWidget(lAddressKey, 0, 0)
    lMainLayout.AddWidget(lWorkspace, 1, 0)

    lParent.SetLayout(lMainLayout)
        
	return lObj
}
//===============================================
// methods
//===============================================
func (obj *GWindow) AddPage(widget widgets.QWidget_ITF, key string, defaultId int) {
    lQt := GManager().mgr.qt
    lPageId := lQt.page_map.Count()
    lQt.page_map.AddWidget(widget)
    lQt.page_id[key] = lPageId
    if defaultId == 1 {
        GManager().SetPage(key)
    }
}
//===============================================
