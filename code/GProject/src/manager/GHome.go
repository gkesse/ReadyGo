//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GHome struct {
    GWidget
}
//===============================================
type GHome_ITF interface {
    GWidget_ITF
    GHome_PTR() *GHome
}
//===============================================
func (obj *GHome) GHome_PTR() *GHome {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGHome(parent widgets.QWidget_ITF) *GHome {
    lObj := &GHome{}

    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    
    lListBox := CreateGWidget("listbox", nil)
    lListBox.AddContent("Qt", "home/qt")
    lListBox.AddContent("SQLite", "home/sqlite")
    lListBox.AddContent("OpenCV", "home/opencv")
    lListBox.AddContent("Builder", "home/builder")
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lListBox, 1, 0)
    
    lParent.SetLayout(lMainLayout)
        
    return lObj
}
//===============================================
// methods
//===============================================
