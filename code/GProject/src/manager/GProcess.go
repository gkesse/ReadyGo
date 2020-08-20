//===============================================
package manager
//===============================================
import "fmt"
import "sync"
//===============================================
type GProcessO struct {
    G_STATE string
}
//===============================================
var m_GProcessLock = &sync.Mutex{}
var m_GProcessO *GProcessO
//===============================================
func newGProcess() *GProcessO {
    lObj := &GProcessO{}
    return lObj
}
//===============================================
func GProcess() *GProcessO {
    if m_GProcessO == nil {
        m_GProcessLock.Lock()
        defer m_GProcessLock.Unlock()
        if m_GProcessO == nil {
            m_GProcessO = newGProcess()
        }
    }
    return m_GProcessO
}
//===============================================
func (obj *GProcessO) Run() {
	fmt.Printf("GProcess_Run\n")
}
//===============================================
