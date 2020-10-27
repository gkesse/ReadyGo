//===============================================
package manager
//===============================================
import "sync"
//===============================================
type GConfigO struct {
    m_dataMap map[string]string
}
//===============================================
var m_GConfigLock = &sync.Mutex{}
var m_GConfigO *GConfigO
//===============================================
func newGConfig() *GConfigO {
    lObj := &GConfigO{}
    lObj.m_dataMap = make(map[string]string)
    return lObj
}
//===============================================
func GConfig() *GConfigO {
    if m_GConfigO == nil {
        m_GConfigLock.Lock()
        defer m_GConfigLock.Unlock()
        if m_GConfigO == nil {
            m_GConfigO = newGConfig()
        }
    }
    return m_GConfigO
}
//===============================================
func (obj *GConfigO) SetData(keyIn string, valueIn string) {
    obj.m_dataMap[keyIn] = valueIn
}
//===============================================
func (obj *GConfigO) GetData(keyIn string) string {
    return obj.m_dataMap[keyIn]
}
//===============================================
