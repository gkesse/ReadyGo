//===============================================
package manager
//===============================================
import "sync"
import "os"
import "github.com/therecipe/qt/widgets"
//===============================================
type GQtO struct {

}
//===============================================
var m_GQtLock = &sync.Mutex{}
var m_GQtO *GQtO
//===============================================
func newGQt() *GQtO {
    lObj := &GQtO{}
    return lObj
}
//===============================================
func GQt() *GQtO {
    if m_GQtO == nil {
        m_GQtLock.Lock()
        defer m_GQtLock.Unlock()
        if m_GQtO == nil {
            m_GQtO = newGQt()
        }
    }
    return m_GQtO
}
//===============================================
func (obj *GQtO) Open() {
    lApp := widgets.NewQApplication(len(os.Args), os.Args)
    GManager().SetStyle(lApp)
    lWindow := CreateGWidget("window", nil)
    lWindow.GWidget_PTR().Show()
    lApp.Exec()
}
//===============================================
