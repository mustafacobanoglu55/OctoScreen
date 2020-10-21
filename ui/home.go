package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
)

var homePanelInstance *homePanel

type homePanel struct {
	CommonPanel
}

func HomePanel(ui *UI, parent Panel) Panel {
	if homePanelInstance == nil {
		m := &homePanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		homePanelInstance = m
	}

	return homePanelInstance
}

func (m *homePanel) initialize() {
	defer m.Initialize()

	m.AddButton(m.createMoveButton("XYZ Sıfırla", "home.svg",
		octoprint.XAxis, octoprint.YAxis, octoprint.ZAxis,
	))

	m.AddButton(m.createMoveButton("X Sıfırla", "home-x.svg", octoprint.XAxis))
	m.AddButton(m.createMoveButton("Y Sıfırla", "home-y.svg", octoprint.YAxis))
	m.AddButton(m.createMoveButton("Z Sıfırla", "home-z.svg", octoprint.ZAxis))
}

func (m *homePanel) createMoveButton(label, image string, axes ...octoprint.Axis) gtk.IWidget {
	return MustButtonImageStyle(label, image, "color2", func() {
		cmd := &octoprint.PrintHeadHomeRequest{Axes: axes}
		Logger.Warningf("Homing the print head in %s axes", axes)
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})
}
