//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "path"
import "path/filepath"
import "io/ioutil"
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/gui"
//===============================================
// struct
//===============================================
type GManagerO struct {
    mgr *sGManager    
}
//===============================================
type sGManager struct {
    app *sGApp
    qt *sGQt
    size *sGSize
    sqlite *sGSQLite
}
//===============================================
type sGApp struct {
    app_name string
    font_path string
    icon_path string
    icon_map map[string]string
    style_path string
}
//===============================================
type sGQt struct {
    // window
    win *GWindow
    // page
    page_map *widgets.QStackedWidget
    page_id map[string]int
    current_page string
    // address
    address_key GWidget_ITF
    address_edit *widgets.QLineEdit
    // title
    title_id *widgets.QLabel
    title_map map[string]string
    // signal
    widget_id string
}
//===============================================
type sGSize struct {
    logo int
    icon int
}
//===============================================
type sGSQLite struct {
    db_path string
}
//===============================================
// constructor
//===============================================
var m_GManagerLock = &sync.Mutex{}
var m_GManagerO *GManagerO
//===============================================
func newGManager() *GManagerO {
    lObj := &GManagerO{}
    lObj.initObj();
    lObj.loadIconObj();
    lObj.loadFontObj();
    return lObj
}
//===============================================
func GManager() *GManagerO {
    if m_GManagerO == nil {
        m_GManagerLock.Lock()
        defer m_GManagerLock.Unlock()
        if m_GManagerO == nil {
            m_GManagerO = newGManager()
        }
    }
    return m_GManagerO
}
//===============================================
// obj
//===============================================
func (obj *GManagerO) GetData() *sGManager {
    return obj.mgr
}
//===============================================
func (obj *GManagerO) initObj() {
    // manager
    obj.mgr = &sGManager{}
    // app
    obj.mgr.app = &sGApp{}
    obj.mgr.app.app_name = "ReadyApp"
    obj.mgr.app.font_path = "data/font"
    obj.mgr.app.icon_path = "data/img/fa/white/64"
    obj.mgr.app.icon_map = make(map[string]string)
    obj.mgr.app.style_path = "data/css/style.css"
    // qt
    obj.mgr.qt = &sGQt{}
    obj.mgr.qt.page_id = make(map[string]int)
    obj.mgr.qt.title_map = make(map[string]string)
    obj.mgr.qt.current_page = ""
    // size
    obj.mgr.size = &sGSize{}
    obj.mgr.size.logo = 40
    obj.mgr.size.icon = 20
    // sqlite
    obj.mgr.sqlite = &sGSQLite{}
    obj.mgr.sqlite.db_path = "C:/Users/Admin/Downloads/Programs/ReadyBin/win/.CONFIG_O.dat"
}
//===============================================
func (obj *GManagerO) loadIconObj() {
    lPath := obj.mgr.app.icon_path
    lDir, _ := ioutil.ReadDir(lPath)
    for _, lFile := range lDir {
        if lFile.IsDir() == false {
            lFilename := lFile.Name()
            lFullname := lPath + "/" + lFilename
            lKey := obj.Basename(lFilename)
            obj.mgr.app.icon_map[lKey] = lFullname
        }
    }
}
//===============================================
func (obj *GManagerO) loadFontObj() {
    lPath := obj.mgr.app.font_path
    lDir, _ := ioutil.ReadDir(lPath)
    for _, lFile := range lDir {
        if lFile.IsDir() == false {
            lFilename := lFile.Name()
            lFullname := lPath + "/" + lFilename
            lFont := gui.NewQFontDatabase()
            lFont.AddApplicationFont(lFullname)
        }
    }
}
//===============================================
// manager
//===============================================
func (obj *GManagerO) ShowData() {
    lWidth := -40
    // app
    fmt.Printf("### app :\n")
    fmt.Printf("%*s : %s\n", lWidth, "obj.mgr.app.app_name", obj.mgr.app.app_name)
}
//===============================================
// go_map
//===============================================
func (obj *GManagerO) ClearMap(mapId map[widgets.QWidget_ITF]string) {
    for lItem := range mapId {
        delete(mapId, lItem)
    }
}
//===============================================
// go_string
//===============================================
func (obj *GManagerO) Basename(pathId string) string {
	lFilename := path.Base(pathId)
    lExtension := filepath.Ext(lFilename)
    lBasename := lFilename[0:len(lFilename)-len(lExtension)]
	return lBasename
}
//===============================================
// qt_layout
//===============================================
func (obj *GManagerO) ClearLayout(layoutId widgets.QLayout_ITF) {
    lCount := layoutId.QLayout_PTR().Count()
    for i := 0; i < lCount; i++ {
        lItem := layoutId.QLayout_PTR().TakeAt(0)
        lWidget := lItem.Widget()
        lWidget.DestroyQWidget()
        lItem.DestroyQLayoutItem()
    }
}
//===============================================
// qt_sender
//===============================================
func (obj *GManagerO) GetSender(widgetId map[widgets.QWidget_ITF]string) widgets.QWidget_ITF {
    for lWidget,_ := range widgetId {
        lSender := lWidget.QObject_PTR().Sender().Pointer()
        if lSender == nil {continue}
        return lWidget
    }
    return nil
}
//===============================================
// qt_page
//===============================================
func (obj *GManagerO) SetPage(pageId string) {
    if pageId == "" {obj.SetAddress() ; return}
    lPageId,lOk := obj.mgr.qt.page_id[pageId]
    if lOk == false {obj.SetAddress() ; return}
    if pageId == obj.mgr.qt.current_page {obj.SetAddress() ; return}
    lTitle,lOk := obj.mgr.qt.title_map[pageId]
    if lOk == false {lTitle = ""}
    obj.mgr.qt.page_map.SetCurrentIndex(lPageId)
    obj.mgr.qt.address_key.SetContent(pageId)
    obj.mgr.qt.current_page = pageId
    obj.mgr.qt.title_id.SetText(lTitle)
    obj.mgr.qt.address_edit.SetText(pageId)
}
//===============================================
// qt_address
//===============================================
func (obj *GManagerO) SetAddress() {
    lAddress := obj.mgr.qt.current_page
    obj.mgr.qt.address_edit.SetText(lAddress)
}
//===============================================
// qt_style
//===============================================
func (obj *GManagerO) SetStyle(app *widgets.QApplication) {
    lFile := core.NewQFile2(obj.mgr.app.style_path)
    lFile.Open(core.QIODevice__ReadOnly)
    lStyle := lFile.ReadAll().Data()
    app.SetStyleSheet(lStyle)
}
//===============================================
