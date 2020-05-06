package item

import "github.com/rivo/tview"

type Navi struct {
	*tview.TextView
}

func NewNavi() *Navi {
	n := &Navi{
		TextView: tview.NewTextView().SetTextAlign(tview.AlignLeft).SetDynamicColors(true),
	}
	n.SetTitleAlign(tview.AlignLeft)
	return n
}

const (
	globalNavi = "[yellow::b]r[white]: run, [yellow::b]j[white]: move down, [yellow]k[white]: move up, [yellow]q[white]: quit"
)

func (n *Navi) Update(p tview.Primitive) {
	switch p.(type) {
	case *Linters:
		n.SetText(globalNavi)
	case *SourceFiles:
		n.SetText("")
	case *Results:
	}

}
