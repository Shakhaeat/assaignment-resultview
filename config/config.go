package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Con *sql.DB

func Connect() (*sql.DB, error) {

	//db, err := sql.Open("mysql", "root:@/resultview")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/result")
	if err != nil {
		//fmt.Println("paynay")
		panic(err)

	}

	// Loads queries from file
	// dot, err2 := dotsql.LoadFromFile("queries.sql")
	// if err2 != nil {
	// 	panic(err)
	// }
	// Run queries for create table
	// _, err = dot.Exec(db, "create-students-table")
	// if err != nil {
	
	// 	panic(err4)
	// 	//log.Fatal(err)

	// }

	Con = db
	//defer db.Close()
	return db, nil
}

// func TestCreateTable(t *testing.T) {
// 	db, _ := Connect()
// 	defer db.Close()

// 	row := db.QueryRow(
// 		"SELECT name FROM students WHERE type='table' AND name='students'",
// 	)

// 	var table string
// 	row.Scan(&table)

// 	if table != "users" {
// 		t.Errorf("Table 'users' has not been created")
// 	}
// }
