package main

import (
	"log"
)
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var le *walk.LineEdit
var sport *walk.CheckBox

func main() {
	if _, err := MainWindow1.Run(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(Bind("enabledCB.Checked"))
}

var MainWindow1 = MainWindow{
	Title:    "MainWindow",
	MinSize:  Size{300, 200},
	Layout:   VBox{},
	Children: widget,
}

var widget = []Widget{
	LineEdit1,
	CheckBoxSport,
	PushButtonOK,
	hboxLayout,
	TableView{
		CheckBoxes:       true,
		ColumnsOrderable: true,
		MultiSelection:   true,
		Columns: []TableViewColumn{
			{Title: "编号"},
			{Title: "名称"},
			{Title: "价格"},
		},
	},
}

var le2 *walk.LineEdit

var hboxLayout = Composite{
	Layout: HBox{},
	Children: []Widget{
		LinkLabel{
			Text: "linkd label",
		},
		PushButton{
			Text: "OK",
			OnClicked: func() {
				log.Println("click.")
				le.SetText(le2.Text())
			},
		},
		LineEdit{
			AssignTo: &le2,
			Text:     "hello world.",
		},
	},
}
var LineEdit1 = LineEdit{
	AssignTo: &le,
}

var CheckBoxSport = CheckBox{
	AssignTo: &sport,
	Text:     "喜欢运动",
	Checked:  true,
}

var PushButtonOK = PushButton{
	Text:      "OK",
	OnClicked: OK_Clicked,
}

func OK_Clicked() {
	if sport.Checked() {
		le.SetText("喜欢运动")
	} else {
		le.SetText("不喜欢运动")
	}
	log.Println("on clicked")
}
