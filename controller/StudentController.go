package controller

import (
	//"fmt"

	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/assignment/resultview/entities"
	"github.com/assignment/resultview/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseGlob("views/Index.html"))
	tmp.ExecuteTemplate(w, "Index.html", nil)
}

func ShowResult(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseGlob("views/show_result.html"))
	tmp.ExecuteTemplate(w, "show_result.html", nil)
}

func ViewResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var student entities.Student
	student.Year, _ = strconv.ParseInt(r.Form.Get("year"), 10, 64)
	student.Board = r.Form.Get("board")
	student.RollID, _ = strconv.ParseInt(r.Form.Get("roll"), 10, 64)
	student.Degree = r.Form.Get("degree")
	fmt.Println(student.RollID)
	var studentModel model.StudentModel
	students, _ := studentModel.FindAll(&student)
	data := map[string]interface{}{
		"students": students,
	}
	//fmt.Println(data)
	tmp := template.Must(template.ParseGlob("views/view_result.html"))
	tmp.ExecuteTemplate(w, "view_result.html", data)

}
