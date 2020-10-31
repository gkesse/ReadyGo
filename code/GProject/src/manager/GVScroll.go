//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GVScroll struct {
    GWidget
    scrollLayout *widgets.QVBoxLayout
}
//===============================================
type GVScroll_ITF interface {
    GWidget_ITF
    GVScroll_PTR() *GVScroll
}
//===============================================
func (obj *GVScroll) GVScroll_PTR() *GVScroll {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGVScroll(parent widgets.QWidget_ITF) *GVScroll {
    lObj := &GVScroll{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GVScroll")
    
    lScrollLayout := widgets.NewQVBoxLayout()
    lObj.scrollLayout = lScrollLayout
    
    lScrollWidget := widgets.NewQWidget(nil, 0)
    lScrollWidget.SetLayout(lScrollLayout)
    
    lScrollArea := widgets.NewQScrollArea(nil)
    lScrollArea.SetWidget(lScrollWidget)
    
    lMainLayout := widgets.NewQVBoxLayout()
    lMainLayout.AddWidget(lScrollArea, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
func (obj *GVScroll) AddScroll(widget widgets.QWidget_ITF) {
    obj.scrollLayout.AddWidget(widget, 0, 0)
}
//===============================================
