//===============================================
package manager
//===============================================
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/core"
import "github.com/therecipe/qt/gui"
//===============================================
// struct
//===============================================
type GTitleBar struct {
    GWidget
    widgetId map[widgets.QWidget_ITF]string
}
//===============================================
type GTitleBar_ITF interface {
    GWidget_ITF
    GTitleBar_PTR() *GTitleBar
}
//===============================================
func (obj *GTitleBar) GTitleBar_PTR() *GTitleBar {
    return obj
}
//===============================================
// constructor
//===============================================
func NewGTitleBar(parent widgets.QWidget_ITF) *GTitleBar {
    lApp := GManager().mgr.app
    lQt := GManager().mgr.qt
    lSize := GManager().mgr.size

    lObj := &GTitleBar{}

    lParent := *NewGWidget(parent)
    lObj.GWidget = lParent
    lParent.SetObjectName("GTitleBar")
    
    lObj.widgetId = make(map[widgets.QWidget_ITF]string)

    lLogo := widgets.NewQPushButton(nil)
    lLogo.SetObjectName("logo")
    lLogo.SetIcon(gui.NewQIcon5(lApp.icon_map["wifi"]));
    lLogo.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lObj.widgetId[lLogo] = "logo"

    lAppName := widgets.NewQLabel(nil, 0)
    lAppName.SetObjectName("app_name")
    lAppName.SetText(lApp.app_name);
        
    lTitle := widgets.NewQLabel(nil, 0)
    lQt.title_id = lTitle
    lTitle.SetObjectName("title")

    lMinimize := widgets.NewQPushButton(nil)
    lMinimize.SetObjectName("minimize")
    lMinimize.SetIcon(gui.NewQIcon5(lApp.icon_map["window-minimize"]));
    lMinimize.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lMinimize.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
    lObj.widgetId[lMinimize] = "minimize"

    lMaximize := widgets.NewQPushButton(nil)
    lMaximize.SetObjectName("maximize")
    lMaximize.SetIcon(gui.NewQIcon5(lApp.icon_map["window-maximize"]));
    lMaximize.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lMaximize.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
    lObj.widgetId[lMaximize] = "maximize"

    lClose := widgets.NewQPushButton(nil)
    lClose.SetObjectName("close")
    lClose.SetIcon(gui.NewQIcon5(lApp.icon_map["times"]));
    lClose.SetIconSize(core.NewQSize2(lSize.icon, lSize.icon));
    lClose.SetCursor(gui.NewQCursor2(core.Qt__PointingHandCursor))
    lObj.widgetId[lClose] = "close"
    
    lMainLayout := widgets.NewQHBoxLayout()
    lMainLayout.AddWidget(lLogo, 0, 0)
    lMainLayout.AddWidget(lAppName, 0, 0)
    lMainLayout.AddWidget(lTitle, 1, core.Qt__AlignCenter)
    lMainLayout.AddWidget(lMinimize, 0, 0)
    lMainLayout.AddWidget(lMaximize, 0, 0)
    lMainLayout.AddWidget(lClose, 0, 0)
    lMainLayout.QLayout_PTR().SetContentsMargins(0, 0, 0, 0)
    lMainLayout.SetSpacing(5)
    
    lParent.SetLayout(lMainLayout)
    
    lLogo.ConnectClicked(lObj.SlotItemClicked)
    lMaximize.ConnectClicked(lObj.SlotItemClicked)
    lMinimize.ConnectClicked(lObj.SlotItemClicked)
    lClose.ConnectClicked(lObj.SlotItemClicked)

    return lObj
}
//===============================================
// slot
//===============================================
func (obj *GTitleBar) SlotItemClicked(ok bool) {
    lQt := GManager().mgr.qt
    
    lWidget := GManager().GetSender(obj.widgetId)
    lWidgetId := obj.widgetId[lWidget]
    if lWidgetId == "close" {
        lQt.win.Close()
    }
}
//===============================================
