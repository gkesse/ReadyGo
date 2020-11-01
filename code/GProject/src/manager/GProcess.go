//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "os"
import "github.com/therecipe/qt/widgets"
//===============================================
type GProcessO struct {
    
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
    lArgc := len(os.Args)
    lKey := "ui"
    if lArgc > 1 { lKey = os.Args[1] }
    if lKey == "ui" { obj.ui() ; return }
    if lKey == "test" { obj.test() ; return }
    obj.ui()
}
//===============================================
func (obj *GProcessO) test() {
    fmt.Printf("\n### GO_TEST\n\n")
    lApp := widgets.NewQApplication(len(os.Args), os.Args)
    GManager().SetStyle(lApp, "data/css/styles.css")
    lWindow := CreateGWidget("window", nil)
    lWindow.GWidget_PTR().Show()
    lApp.Exec()
}
//===============================================
func (obj *GProcessO) ui() {
    GProcessUi().Run()
}
//===============================================
