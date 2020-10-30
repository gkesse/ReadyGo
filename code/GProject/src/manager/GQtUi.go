//===============================================
package manager
//===============================================
import "fmt"
import "sync"
import "os"
import "bufio"
//===============================================
type GQtUiO struct {
    G_STATE string
}
//===============================================
var m_GQtUiLock = &sync.Mutex{}
var m_GQtUiO *GQtUiO
//===============================================
func newGQtUi() *GQtUiO {
    lObj := &GQtUiO{}
    return lObj
}
//===============================================
func GQtUi() *GQtUiO {
    if m_GQtUiO == nil {
        m_GQtUiLock.Lock()
        defer m_GQtUiLock.Unlock()
        if m_GQtUiO == nil {
            m_GQtUiO = newGQtUi()
        }
    }
    return m_GQtUiO
}
//===============================================
func (obj *GQtUiO) Run() {
    obj.G_STATE = "S_INIT"
    for {
        if obj.G_STATE == "S_ADMIN" { obj.run_ADMIN()
        } else if obj.G_STATE == "S_INIT" { obj.run_INIT()
        } else if obj.G_STATE == "S_METHOD" { obj.run_METHOD()
        } else if obj.G_STATE == "S_CHOICE" { obj.run_CHOICE()
        //
        } else if obj.G_STATE == "S_OPEN" { obj.run_OPEN()
        } else if obj.G_STATE == "S_CLOSE" { obj.run_CLOSE()
        //
        } else if obj.G_STATE == "S_SAVE" { obj.run_SAVE()
        } else if obj.G_STATE == "S_LOAD" { obj.run_LOAD()
        } else if obj.G_STATE == "S_QUIT" { obj.run_QUIT()
        } else { break }
    }
}
//===============================================
func (obj *GQtUiO) run_ADMIN() {
    GProcess().Run()
    obj.G_STATE = "S_END"
}
//===============================================
func (obj *GQtUiO) run_INIT() {
	fmt.Printf("\n")
	fmt.Printf("GO_QT !!!\n")
	fmt.Printf("\t%-2s : %s\n", "-q", "quitter l'application")
	fmt.Printf("\t%-2s : %s\n", "-i", "reinitialiser l'application")
	fmt.Printf("\t%-2s : %s\n", "-a", "redemarrer l'application")
	fmt.Printf("\t%-2s : %s\n", "-v", "valider les configurations")
	fmt.Printf("\n")
    obj.G_STATE = "S_LOAD"
}
//===============================================
func (obj *GQtUiO) run_METHOD() {
	fmt.Printf("GO_QT :\n")
	fmt.Printf("\t%-2s : %s\n", "1", "ouvrir l'application")
	fmt.Printf("\t%-2s : %s\n", "2", "fermer l'application")
	fmt.Printf("\n")
    obj.G_STATE = "S_CHOICE"
}
//===============================================
func (obj *GQtUiO) run_CHOICE() {
    var lLast string = "1"
	fmt.Printf("GO_QT (%s) ? ", lLast)
    lReader := bufio.NewReader(os.Stdin) ; lAnswer, _ := lReader.ReadString('\n')
    lAnswer = lAnswer[:len(lAnswer)-2]
    if lAnswer == "" { lAnswer = lLast }
    if lAnswer == "-q" { obj.G_STATE = "S_END" 
    } else if lAnswer == "-i" { obj.G_STATE = "S_INIT" 
    } else if lAnswer == "-a" { obj.G_STATE = "S_ADMIN" 
    //
    } else if lAnswer == "1" { obj.G_STATE = "S_OPEN" ; GConfig().SetData("GO_QT_ID", lAnswer);
    } else if lAnswer == "2" { obj.G_STATE = "S_CLOSE" ; GConfig().SetData("GO_QT_ID", lAnswer);
    //
    }
}
//===============================================
func (obj *GQtUiO) run_OPEN() {
	fmt.Printf("\n")
    GQt().Open()
    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GQtUiO) run_CLOSE() {
	fmt.Printf("\n")

    obj.G_STATE = "S_SAVE"
}
//===============================================
func (obj *GQtUiO) run_SAVE() {
    obj.G_STATE = "S_QUIT"
}
//===============================================
func (obj *GQtUiO) run_LOAD() {
    obj.G_STATE = "S_METHOD"
}
//===============================================
func (obj *GQtUiO) run_QUIT() {
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
