//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "bufio"
import "os"
//===============================================
type GSQLiteUiO struct {
    G_STATE string
}
//===============================================
var m_GSQLiteUiLock = &sync.Mutex{}
var m_GSQLiteUiO *GSQLiteUiO
//===============================================
func newGSQLiteUi() *GSQLiteUiO {
    lObj := &GSQLiteUiO{}
    return lObj
}
//===============================================
func GSQLiteUi() *GSQLiteUiO {
    if m_GSQLiteUiO == nil {
        m_GSQLiteUiLock.Lock()
        defer m_GSQLiteUiLock.Unlock()
        if m_GSQLiteUiO == nil {
            m_GSQLiteUiO = newGSQLiteUi()
        }
    }
    return m_GSQLiteUiO
}
//===============================================
func (obj *GSQLiteUiO) Run() {
    obj.G_STATE = "S_INIT"
    for {
        if obj.G_STATE == "S_ADMIN" { obj.run_ADMIN()
        } else if obj.G_STATE == "S_INIT" { obj.run_INIT()
        } else if obj.G_STATE == "S_METHOD" { obj.run_METHOD()
        } else if obj.G_STATE == "S_CHOICE" { obj.run_CHOICE()
        //
        } else if obj.G_STATE == "S_BUILDER" { obj.run_BUILDER()
        } else if obj.G_STATE == "S_TABLES_SHOW" { obj.run_TABLES_SHOW()
        //
        } else if obj.G_STATE == "S_SAVE" { obj.run_SAVE()
        } else if obj.G_STATE == "S_LOAD" { obj.run_LOAD()
        } else if obj.G_STATE == "S_QUIT" { obj.run_QUIT()
        } else { break }
    }
}
//===============================================
func (obj *GSQLiteUiO) run_ADMIN() {
    GProcess().Run()
    obj.G_STATE = "S_END"
}
//===============================================
func (obj *GSQLiteUiO) run_INIT() {
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
func (obj *GSQLiteUiO) run_METHOD() {
	fmt.Printf("GO_SQLITE :\n")
	fmt.Printf("\t%-2s : %s\n", "0", "lancer le builder")
	fmt.Printf("\n")
	fmt.Printf("\t%-2s : %s\n", "1", "afficher les tables")
	fmt.Printf("\t%-2s : %s\n", "2", "afficher la table CONFIG_GO")
	fmt.Printf("\n")
    obj.G_STATE = "S_CHOICE"
}
//===============================================
func (obj *GSQLiteUiO) run_CHOICE() {
    var lLast string = "1"
	fmt.Printf("GO_SQLITE (%s) ? ", lLast)
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "" { lAnswer = lLast }
    if lAnswer == "-q" { obj.G_STATE = "S_END" 
    //
    } else if lAnswer == "0" { obj.G_STATE = "S_BUILDER" ; GConfig().SetData("GO_SQLITE_ID", lAnswer);
    //
    } else if lAnswer == "1" { obj.G_STATE = "S_TABLES_SHOW" ; GConfig().SetData("GO_SQLITE_ID", lAnswer);
    //
    }
}
//===============================================
func (obj *GSQLiteUiO) run_BUILDER() {
	fmt.Printf("\n")
    //lDb, lErr :=  sql.Open("sqlite3", "./foo.db")
    lApp := GManager().GetData().app;
    lApp.app_name = "ReadyGo";
    GManager().ShowData()
    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GSQLiteUiO) run_TABLES_SHOW() {
	fmt.Printf("\n")

    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GSQLiteUiO) run_SAVE() {
    obj.G_STATE = "S_QUIT"
}
//===============================================
func (obj *GSQLiteUiO) run_LOAD() {
    obj.G_STATE = "S_METHOD"
}
//===============================================
func (obj *GSQLiteUiO) run_QUIT() {
	fmt.Printf("\n")
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
