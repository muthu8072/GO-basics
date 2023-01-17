package main

import (
	"database/sql"
	  "fmt"
	"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "807248"
	dbname   = "calhounio_demo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into  Employee("Name","EmpId")values("rohith",21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	//insertstmt:="insert into "Employee"("Name","EmpId")values($1,$2)"
	//_,e:=db.Exec(insertstmt,"krish",03)
	//CheckError(e)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
