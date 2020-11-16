//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GDatabaseView struct {
    GWidget
    workspace GWidget_ITF
}
//===============================================
type GDatabaseView_ITF interface {
    GWidget_ITF
    GDatabaseView_PTR() *GDatabaseView
}
//===============================================
func (obj *GDatabaseView) GDatabaseView_PTR() *GDatabaseView {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGDatabaseView(parent widgets.QWidget_ITF) *GDatabaseView {   
    lObj := &GDatabaseView{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GDatabaseView")
    
    lOpen := widgets.NewQPushButton(nil)
    lOpen.SetText("Open")
    
    lWorkspace := CreateGWidget("listbox", nil)
    lObj.workspace = lWorkspace
    
    lHeaderLayout := widgets.NewQHBoxLayout()
    lHeaderLayout.AddWidget(lOpen, 0, 0)
    lHeaderLayout.AddStretch(1)
    lHeaderLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lHeaderLayout.SetSpacing(5)
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddLayout(lHeaderLayout, 0)
    lMainLayout.AddWidget(lWorkspace, 0, 0)
    lMainLayout.AddStretch(1)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)

    lObj.loadTables()

    return lObj
}
//===============================================
// private
//===============================================
func (obj *GDatabaseView) loadTables() {
    lQuery := `
    select name from sqlite_master 
    where type='table'
    `
    lTables := GSQLite().QueryCol(lQuery)
    for i := 0; i < len(lTables); i++ {
        lTable := lTables[i]
        obj.workspace.AddContent(lTable, "table")
    }
}
//===============================================
