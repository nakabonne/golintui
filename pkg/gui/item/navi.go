package item

import (
	"fmt"

	"github.com/rivo/tview"
)

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
	globalNavi      = "[aqua]q[white]: quit, [aqua]r[white]: run, [aqua]j[white]: move down, [aqua]k[white]: move up, [aqua]l[white]: next panel, [aqua]h[white]: previous panel"
	lintersNavi     = "[aqua]space[white]: toggle enabled"
	sourceFilesNavi = "[aqua]space[white]: toggle selected, [aqua]o[white]: expand directory"
	commitsNavi     = "[aqua]space[white]: toggle selected"
	resultsNavi     = "[aqua]o[white]: open file"
)

func (n *Navi) Update(p tview.Primitive) {
	switch p.(type) {
	case *Linters:
		n.SetText(fmt.Sprintf("%s, %s", globalNavi, lintersNavi))
	case *SourceFiles:
		n.SetText(fmt.Sprintf("%s, %s", globalNavi, sourceFilesNavi))
	case *Commits:
		n.SetText(fmt.Sprintf("%s, %s", globalNavi, commitsNavi))
	case *Results:
		n.SetText(fmt.Sprintf("%s, %s", globalNavi, resultsNavi))
	}

}
