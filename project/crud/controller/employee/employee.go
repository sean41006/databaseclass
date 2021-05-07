package employee

import (
	"crud/config"
	"crud/model"
	"fmt"
	"log"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// AllTest = Select Test API
func AllEmployee(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT e_id, name,phone,account,parkcard_id,b_id FROM employee`)

	if err != nil {
		log.Print(err)
	}

	var employee model.Employee
	var arremployee []model.Employee

	for rows.Next() {
		err = rows.Scan(&employee.E_id, &employee.Name, &employee.Phone, &employee.Account, &employee.Packcard_id, &employee.B_id)
		fmt.Println(employee)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arremployee = append(arremployee, employee)
		}
	}
	fmt.Println(arremployee)
	context.JSON(200, gin.H{"arremployee": arremployee})

}

// InsertTest = Insert Test API
func InsertEmployee(context *gin.Context) {
	context.JSON(200, gin.H{})
	var employee model.Employee
	context.BindJSON(&employee)
	fmt.Println(employee)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO employee(name,phone,account,parkcard_id,b_id) VALUES(?,?,?,?,?)`, employee.Name, employee.Phone, employee.Account, employee.Packcard_id, employee.B_id)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateEmployee(context *gin.Context) {
	context.JSON(200, gin.H{})
	var employee model.Employee
	context.BindJSON(&employee)

	fmt.Println(employee)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if employee.Name != "" {

		_, err = db.Exec("UPDATE employee SET name=?,phone=?,account=?,parkcard_id=?,b_id=? WHERE e_id=?", employee.Name, employee.Phone, employee.Account, employee.Packcard_id, employee.B_id, employee.E_id)
	}

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteEmployee(context *gin.Context) {
	context.JSON(200, gin.H{})
	var employee model.Employee
	context.BindJSON(&employee.E_id)
	fmt.Println(employee.E_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM employee WHERE e_id=?", employee.E_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
