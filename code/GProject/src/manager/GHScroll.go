//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
//===============================================
// struct
//===============================================
type GHScroll struct {
    GWidget
    scrollLayout *widgets.QHBoxLayout
}
//===============================================
type GHScroll_ITF interface {
    GWidget_ITF
    GHScroll_PTR() *GHScroll
}
//===============================================
func (obj *GHScroll) GHScroll_PTR() *GHScroll {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGHScroll(parent widgets.QWidget_ITF) *GHScroll {
    lObj := &GHScroll{}
    
    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GHScroll")
    
    lScrollLayout := widgets.NewQHBoxLayout()
    lObj.scrollLayout = lScrollLayout
    
    lScrollWidget := widgets.NewQWidget(nil, 0)
    lScrollWidget.SetLayout(lScrollLayout)
    
    lScrollArea := widgets.NewQScrollArea(nil)
    lScrollArea.SetWidget(lScrollWidget)
    
    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lScrollArea, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(0)
    
    lParent.SetLayout(lMainLayout)
    
    return lObj
}
//===============================================
func (obj *GHScroll) AddScroll(widget widgets.QWidget_ITF) {
    obj.scrollLayout.AddWidget(widget, 0, 0)
}
//===============================================
