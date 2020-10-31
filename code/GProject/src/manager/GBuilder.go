//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GBuilder struct {
    GWidget
    mainLayout *widgets.QVBoxLayout
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
            
    lMainLayout := widgets.NewQVBoxLayout()
    lObj.mainLayout = lMainLayout
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
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
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
func (obj *GBuilder) AddBuilder(key string) {
    lWidget := CreateGWidget(key, nil)
    obj.mainLayout.AddWidget(lWidget, 0, 0)
}
//===============================================
