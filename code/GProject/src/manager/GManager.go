//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "github.com/therecipe/qt/widgets"
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
    sqlite *sGSQLite
}
//===============================================
type sGApp struct {
    app_name string
}
//===============================================
type sGQt struct {
    page_map *widgets.QStackedWidget
    page_id map[string]int
    current_page string
    address_key GWidget_ITF
    title_id *widgets.QLabel
    title_map map[string]string
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
func (obj *GManagerO) initObj() {
    // manager
    obj.mgr = &sGManager{}
    // app
    obj.mgr.qt = &sGQt{}
    obj.mgr.qt.page_id = make(map[string]int)
    obj.mgr.qt.title_map = make(map[string]string)
    obj.mgr.qt.current_page = ""
    // app
    obj.mgr.app = &sGApp{}
    obj.mgr.app.app_name = "ReadyApp"
    // sqlite
    obj.mgr.sqlite = &sGSQLite{}
    obj.mgr.sqlite.db_path = "C:\\Users\\Admin\\Downloads\\Programs\\ReadyBin\\win\\.CONFIG_O.dat"
}
//===============================================
func (obj *GManagerO) GetData() *sGManager {
    return obj.mgr
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
    if pageId == "" {return}
    lPageId,lOk := obj.mgr.qt.page_id[pageId]
    if lOk == false {return}
    if pageId == obj.mgr.qt.current_page {return}
    lTitle,lOk := obj.mgr.qt.title_map[pageId]
    if lOk == false {lTitle = ""}
    obj.mgr.qt.page_map.SetCurrentIndex(lPageId)
    obj.mgr.qt.address_key.SetContent(pageId)
    obj.mgr.qt.current_page = pageId
    obj.mgr.qt.title_id.SetText(lTitle)
}
//===============================================
