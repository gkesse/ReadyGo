//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "bufio"
import "os"
//===============================================
type GSQLiteO struct {
    G_STATE string
}
//===============================================
var m_GSQLiteLock = &sync.Mutex{}
var m_GSQLiteO *GSQLiteO
//===============================================
func newGSQLite() *GSQLiteO {
    lObj := &GSQLiteO{}
    return lObj
}
//===============================================
func GSQLite() *GSQLiteO {
    if m_GSQLiteO == nil {
        m_GSQLiteLock.Lock()
        defer m_GSQLiteLock.Unlock()
        if m_GSQLiteO == nil {
            m_GSQLiteO = newGSQLite()
        }
    }
    return m_GSQLiteO
}
//===============================================
func (obj *GSQLiteO) Run() {
    obj.G_STATE = "S_INIT"
    for {
        if obj.G_STATE == "S_ADMIN" { obj.run_ADMIN()
        } else if obj.G_STATE == "S_INIT" { obj.run_INIT()
        } else if obj.G_STATE == "S_METHOD" { obj.run_METHOD()
        } else if obj.G_STATE == "S_CHOICE" { obj.run_CHOICE()
        //
        } else if obj.G_STATE == "S_CONFIG_SHOW" { obj.run_CONFIG_SHOW()
        //
        } else if obj.G_STATE == "S_SAVE" { obj.run_SAVE()
        } else if obj.G_STATE == "S_LOAD" { obj.run_LOAD()
        } else if obj.G_STATE == "S_QUIT" { obj.run_QUIT()
        } else { break }
    }
}
//===============================================
func (obj *GSQLiteO) run_ADMIN() {
    GProcess().Run()
    obj.G_STATE = "S_END"
}
//===============================================
func (obj *GSQLiteO) run_INIT() {
	fmt.Printf("\n")
	fmt.Printf("GO_SQLITE !!!\n")
	fmt.Printf("\t%-2s : %s\n", "-q", "quitter l'application")
	fmt.Printf("\t%-2s : %s\n", "-i", "reinitialiser l'application")
	fmt.Printf("\t%-2s : %s\n", "-a", "redemarrer l'application")
	fmt.Printf("\t%-2s : %s\n", "-v", "valider les configurations")
	fmt.Printf("\n")
    obj.G_STATE = "S_LOAD"
}
//===============================================
func (obj *GSQLiteO) run_METHOD() {
	fmt.Printf("GO_SQLITE :\n")
	fmt.Printf("\t%-2s : %s\n", "1", "S_CONFIG_SHOW")
	fmt.Printf("\n")
    obj.G_STATE = "S_CHOICE"
}
//===============================================
func (obj *GSQLiteO) run_CHOICE() {
    var lLast string = "1"
	fmt.Printf("GO_SQLITE (%s) ? ", lLast)
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "" { lAnswer = lLast }
    if lAnswer == "-q" { obj.G_STATE = "S_END" 
    } else if lAnswer == "1" { obj.G_STATE = "S_CONFIG_SHOW" ; GConfig().SetData("GO_SQLITE_ID", lAnswer);
    }
}
//===============================================
func (obj *GSQLiteO) run_CONFIG_SHOW() {
    obj.G_STATE = "S_SAVE"
	fmt.Printf("\n")
}
//===============================================
func (obj *GSQLiteO) run_SAVE() {
    obj.G_STATE = "S_QUIT"
}
//===============================================
func (obj *GSQLiteO) run_LOAD() {
    obj.G_STATE = "S_METHOD"
}
//===============================================
func (obj *GSQLiteO) run_QUIT() {
	fmt.Printf("GO_QUIT (Oui/[N]on) ? ")
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "-q" { obj.G_STATE = "S_END" 
    } else if lAnswer == "-i" { obj.G_STATE = "S_INIT" 
    } else if lAnswer == "-a" { obj.G_STATE = "S_ADMIN" 
    } else if lAnswer == "o" { obj.G_STATE = "S_END" 
    } else if lAnswer == "n" { obj.G_STATE = "S_INIT" 
    } else if lAnswer == "" { obj.G_STATE = "S_INIT" 
    }
}
//===============================================
