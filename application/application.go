package application

import (
	tl "github.com/JoelOtter/termloop"
)

type MovingText struct {
	*tl.Text
}

func (m *MovingText) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := m.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x += 1
		case tl.KeyArrowLeft:
			x -= 1
		case tl.KeyArrowUp:
			y -= 1
		case tl.KeyArrowDown:
			y += 1
		}
		m.SetPosition(x, y)
	}
}

func RunApp() {
	g := tl.NewGame()
	g.Screen().SetFps(30)
	g.Screen().AddEntity(&MovingText{tl.NewText(0, 0, "api-server", tl.ColorWhite, tl.ColorBlue)})
	g.Screen().AddEntity(&MovingText{tl.NewText(2, 1, "control manager", tl.ColorWhite, tl.ColorCyan)})
	g.Screen().AddEntity(&MovingText{tl.NewText(4, 2, "kube-agent", tl.ColorWhite, tl.ColorGreen)})
	g.Screen().AddEntity(&MovingText{tl.NewText(6, 3, "etcd", tl.ColorWhite, tl.ColorYellow)})
	g.Screen().AddEntity(&MovingText{tl.NewText(6, 4, "CRE", tl.ColorWhite, tl.ColorRed)})
	g.Screen().AddEntity(&MovingText{tl.NewText(4, 5, "cloud control", tl.ColorWhite, tl.ColorMagenta)})
	g.Screen().AddEntity(&MovingText{tl.NewText(2, 6, "scheduler", tl.ColorWhite, tl.ColorWhite)})
	g.Screen().AddEntity(&MovingText{tl.NewText(0, 7, "kube-proxy", tl.ColorWhite, tl.ColorBlack)})

	g.Screen().AddEntity(&MovingText{tl.NewText(60, 0, "master", tl.ColorBlack, tl.ColorRed)})
	g.Screen().AddEntity(&MovingText{tl.NewText(120, 0, "worker", tl.ColorWhite, tl.ColorBlack)})

	g.Start()

}
