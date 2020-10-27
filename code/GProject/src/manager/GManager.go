//===============================================
package manager
//===============================================
import "fmt"
import "sync"
//===============================================
type GManagerO struct {
    m_mgr *sGManager    
}
//===============================================
type sGManager struct {
    app *sGApp
    sqlite *sGSQLite
}
//===============================================
type sGApp struct {
    app_name string
}
//===============================================
type sGSQLite struct {
    db_path string
}
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
	obj.m_mgr = &sGManager{}
    // app
	obj.m_mgr.app = &sGApp{}
	obj.m_mgr.app.app_name = "ReadyAppp"
    // sqlite
	obj.m_mgr.sqlite = &sGSQLite{}
	obj.m_mgr.sqlite.db_path = "C:/Users/Admin/Downloads/Programs/ReadyBin/win/.CONFIG_O.dat"
}
//===============================================
func (obj *GManagerO) GetData() *sGManager {
	return obj.m_mgr
}
//===============================================
func (obj *GManagerO) ShowData() {
    lWidth := -40
	// app
    fmt.Printf("### app :\n")
    fmt.Printf("%*s : %s\n", lWidth, "obj.m_mgr.app.app_name", obj.m_mgr.app.app_name)
}
//===============================================
