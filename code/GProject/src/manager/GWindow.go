//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/gui"
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
// constructor
//===============================================
func NewGWindow(parent widgets.QWidget_ITF) *GWindow {
    lApp := GManager().mgr.app
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
    
    lObj.AddPage(CreateGWidget("home", nil), "home", "Accueil", 1)
    lObj.AddPage(CreateGWidget("qt", nil), "home/qt", "Qt", 0)
    lObj.AddPage(CreateGWidget("sqlite", nil), "home/sqlite", "SQLite", 0)
    lObj.AddPage(CreateGWidget("opencv", nil), "home/opencv", "OpenCV", 0)
    lObj.AddPage(CreateGWidget("builder", nil), "home/builder", "Builder", 1)
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lTitleBar, 0, 0)
    lMainLayout.AddWidget(lAddressBar, 0, 0)
    lMainLayout.AddWidget(lAddressKey, 0, 0)
    lMainLayout.AddWidget(lWorkspace, 1, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(10, 10, 10, 10)
    lMainLayout.SetSpacing(10)

    lParent.SetLayout(lMainLayout)
    
    lParent.SetWindowTitle(lApp.app_name);
    lParent.SetWindowIcon(gui.NewQIcon5(lApp.icon_map["wifi"]));
    lParent.SetWindowFlags(core.Qt__Window | core.Qt__FramelessWindowHint)
    lParent.Resize2(640, 480)
    
    lQt.win = lObj
    
    return lObj
}
//===============================================
// methods
//===============================================
func (obj *GWindow) AddPage(widget widgets.QWidget_ITF, key string, title string, defaultId int) {
    lQt := GManager().mgr.qt
    lPageId := lQt.page_map.Count()
    lQt.page_map.AddWidget(widget)
    lQt.page_id[key] = lPageId
    lQt.title_map[key] = title
    if defaultId == 1 {
        GManager().SetPage(key)
    }
}
//===============================================
