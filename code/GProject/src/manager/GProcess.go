//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "bufio"
import "os"
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
    obj.G_STATE = "S_INIT"
    for {
        if obj.G_STATE == "S_ADMIN" { obj.run_ADMIN()
        } else if obj.G_STATE == "S_INIT" { obj.run_INIT()
        } else if obj.G_STATE == "S_METHOD" { obj.run_METHOD()
        } else if obj.G_STATE == "S_CHOICE" { obj.run_CHOICE()
        } else if obj.G_STATE == "S_SQLITE" { obj.run_SQLITE()
        } else if obj.G_STATE == "S_SAVE" { obj.run_SAVE()
        } else if obj.G_STATE == "S_LOAD" { obj.run_LOAD()
        } else if obj.G_STATE == "S_QUIT" { obj.run_QUIT()
        } else { break }
    }
}
//===============================================
func (obj *GProcessO) run_ADMIN() {
    obj.G_STATE = "S_END"
}
//===============================================
func (obj *GProcessO) run_INIT() {
    fmt.Printf("\n")
    fmt.Printf("GO_ADMIN !!!\n")
    fmt.Printf("\t%-2s : %s\n", "-q", "quitter l'application")
    fmt.Printf("\n")
    obj.G_STATE = "S_LOAD"
}
//===============================================
func (obj *GProcessO) run_METHOD() {
    fmt.Printf("GO_ADMIN :\n")
    fmt.Printf("\t%-2s : %s\n", "1", "S_SQLITE")
    fmt.Printf("\n")
    obj.G_STATE = "S_CHOICE"
}
//===============================================
func (obj *GProcessO) run_CHOICE() {
    var lLast string = "1"
    fmt.Printf("GO_ADMIN (%s) ? ", lLast)
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "" { lAnswer = lLast }
    if lAnswer == "-q" { obj.G_STATE = "S_END" 
    } else if lAnswer == "1" { obj.G_STATE = "S_SQLITE" ; GConfig().SetData("GO_ADMIN_ID", lAnswer);
    }
}
//===============================================
func (obj *GProcessO) run_SQLITE() {
    GSQLiteUi().Run();
    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GProcessO) run_SAVE() {
    obj.G_STATE = "S_QUIT"
}
//===============================================
func (obj *GProcessO) run_LOAD() {
    obj.G_STATE = "S_METHOD"
}
//===============================================
func (obj *GProcessO) run_QUIT() {
    obj.G_STATE = "S_END"
}
//===============================================
