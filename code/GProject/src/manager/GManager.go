//===============================================
package manager
//===============================================
import "fmt"
import "sync"
//===============================================
type GManagerO struct {
    G_STATE string
}
//===============================================
var m_GManagerLock = &sync.Mutex{}
var m_GManagerO *GManagerO
//===============================================
func newGManager() *GManagerO {
    lObj := &GManagerO{}
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
func (obj *GManagerO) Run() {
	fmt.Printf("GManager_Run\n")
}
//===============================================
