package model

import (
	"fmt"

	"github.com/assignment/resultview/config"
	"github.com/assignment/resultview/entities"
	"github.com/gchaincl/dotsql"
)

type StudentModel struct{}

func (*StudentModel) FindAll(student *entities.Student) ([]entities.Student, error) {

	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		fmt.Println("Connection established.... hoy nay")
		//return , err
	}
	// Loads queries from file
	dot, err2 := dotsql.LoadFromFile("queries.sql")
	if err2 != nil {
		panic(err)
	} else {

		// Run queries for create table
		//fmt.Println(student.RollID)
		//yearStr := strconv.FormatInt(student.Year, 10)
		//tab := student.Degree + yearStr
		//fmt.Println(dot)
		//rows, err4 := db.Query("Select * from "+tab+" WHERE roll_id=? AND board=?", student.RollID, student.Board)
		rows, err4 := dot.Query(db, "find-users-by-degree-year-board-roll", student.RollID,
			student.Board, student.Degree, student.Year)
		if err4 != nil {
			panic(err4)
			//log.Fatal(err)

		} else {
			//fmt.Println("Ayse")
			var students []entities.Student
			for rows.Next() {

				var student entities.Student
				rows.Scan(&student.ID, &student.RollID, &student.Name, &student.FatherName, &student.MotherName,
					&student.Degree, &student.Year, &student.Board, &student.Institute, &student.Bangla,
					&student.GeneralMath, &student.English, &student.History, &student.CreatedAt)
				students = append(students, student)
			}
			//	fmt.Println(students)
			return students, nil
		}
	}

}
