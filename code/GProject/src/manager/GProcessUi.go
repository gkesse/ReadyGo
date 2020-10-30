//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "bufio"
import "os"
//===============================================
type GProcessUiO struct {
    G_STATE string
}
//===============================================
var m_GProcessUiLock = &sync.Mutex{}
var m_GProcessUiO *GProcessUiO
//===============================================
func newGProcessUi() *GProcessUiO {
    lObj := &GProcessUiO{}
    return lObj
}
//===============================================
func GProcessUi() *GProcessUiO {
    if m_GProcessUiO == nil {
        m_GProcessUiLock.Lock()
        defer m_GProcessUiLock.Unlock()
        if m_GProcessUiO == nil {
            m_GProcessUiO = newGProcessUi()
        }
    }
    return m_GProcessUiO
}
//===============================================
func (obj *GProcessUiO) Run() {
    obj.G_STATE = "S_INIT"
    for {
        if obj.G_STATE == "S_INIT" { obj.run_INIT()
        } else if obj.G_STATE == "S_METHOD" { obj.run_METHOD()
        } else if obj.G_STATE == "S_CHOICE" { obj.run_CHOICE()
        //
        } else if obj.G_STATE == "S_SQLITE" { obj.run_SQLITE()
        } else if obj.G_STATE == "S_QT" { obj.run_QT()
        //
        } else if obj.G_STATE == "S_SAVE" { obj.run_SAVE()
        } else if obj.G_STATE == "S_LOAD" { obj.run_LOAD()
        } else { break }
    }
}
//===============================================
func (obj *GProcessUiO) run_INIT() {
    fmt.Printf("\n")
    fmt.Printf("GO_ADMIN !!!\n")
    fmt.Printf("\t%-2s : %s\n", "-q", "quitter l'application")
    fmt.Printf("\n")
    obj.G_STATE = "S_LOAD"
}
//===============================================
func (obj *GProcessUiO) run_METHOD() {
    fmt.Printf("GO_ADMIN :\n")
    fmt.Printf("\t%-2s : %s\n", "1", "S_SQLITE")
    fmt.Printf("\t%-2s : %s\n", "2", "S_QT")
    fmt.Printf("\n")
    obj.G_STATE = "S_CHOICE"
}
//===============================================
func (obj *GProcessUiO) run_CHOICE() {
    var lLast string = "1"
    fmt.Printf("GO_ADMIN (%s) ? ", lLast)
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "" { lAnswer = lLast }
    if lAnswer == "-q" { obj.G_STATE = "S_END"
    //
    } else if lAnswer == "1" { obj.G_STATE = "S_SQLITE" ; GConfig().SetData("GO_ADMIN_ID", lAnswer);
    } else if lAnswer == "2" { obj.G_STATE = "S_QT" ; GConfig().SetData("GO_ADMIN_ID", lAnswer);
    //
    }
}
//===============================================
func (obj *GProcessUiO) run_SQLITE() {
    GSQLiteUi().Run();
    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GProcessUiO) run_QT() {
    GQtUi().Run();
    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GProcessUiO) run_SAVE() {
    obj.G_STATE = "S_QUIT"
}
//===============================================
func (obj *GProcessUiO) run_LOAD() {
    obj.G_STATE = "S_METHOD"
}
//===============================================
