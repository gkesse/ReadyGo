//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GDatabaseView struct {
    GWidget
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
    
    lWorkspace := widgets.NewQLabel(nil, 0)
    lWorkspace.SetText("GDatabaseView")
    
    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lWorkspace, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)

    return lObj
}
//===============================================
