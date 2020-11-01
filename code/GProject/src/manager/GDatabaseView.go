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
    
    lOpen := widgets.NewQPushButton(nil)
    lOpen.SetText("Open")
        
    lHeaderLayout := widgets.NewQHBoxLayout()
    lHeaderLayout.AddWidget(lOpen, 0, 0)
    lHeaderLayout.AddStretch(1)
    lHeaderLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lHeaderLayout.SetSpacing(5)
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddLayout(lHeaderLayout, 0)
    lHeaderLayout.AddStretch(1)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)

    return lObj
}
//===============================================
