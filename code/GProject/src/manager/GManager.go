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
}
//===============================================
type sGApp struct {
    app_name string
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
}
//===============================================
func (obj *GManagerO) GetData() *sGManager {
	return obj.m_mgr
}
//===============================================
func (obj *GManagerO) ShowData() {
	// app
    fmt.Printf("### app :\n");
    fmt.Printf("%*s : %s\n", -40, "obj.m_mgr.app.app_name", obj.m_mgr.app.app_name)
}
//===============================================
