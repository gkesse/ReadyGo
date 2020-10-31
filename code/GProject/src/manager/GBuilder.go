//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GBuilder struct {
    GWidget
    contentWidget GWidget_ITF
}
//===============================================
type GBuilder_ITF interface {
    GWidget_ITF
    GBuilder_PTR() *GBuilder
}
//===============================================
func (obj *GBuilder) GBuilder_PTR() *GBuilder {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGBuilder(parent widgets.QWidget_ITF) *GBuilder {
    lObj := &GBuilder{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GBuilder")
    
    lContentWidget := CreateGWidget("vscroll", nil)
    lObj.contentWidget = lContentWidget
    
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    lObj.AddBuilder("databaseview")
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lContentWidget, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
func (obj *GBuilder) AddBuilder(key string) {
    lWidget := CreateGWidget(key, nil)
    obj.contentWidget.AddScroll(lWidget)
}
//===============================================
