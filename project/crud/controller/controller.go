package controller

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
func AllTest(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:1000sean334@/aed_project")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name,phone FROM test`)

	if err != nil {
		log.Print(err)
	}

	var test model.Test
	var arrtest []model.Test

	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Name, &test.Phone)
		fmt.Println(test)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrtest = append(arrtest, test)
		}
	}
	fmt.Println(arrtest)
	context.JSON(200, gin.H{"arrtest": arrtest})

}

// InsertTest = Insert Test API
func InsertTest(context *gin.Context) {
	context.JSON(200, gin.H{})
	var test model.Test
	context.BindJSON(&test)
	fmt.Println(test)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/aed_project")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO test(name,phone) VALUES(?,?)`, test.Name, test.Phone)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateTest(context *gin.Context) {
	context.JSON(200, gin.H{})
	var test model.Test
	context.BindJSON(&test)

	fmt.Println(test)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/aed_project")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if test.Name != "" {

		_, err = db.Exec("UPDATE Test SET name=?,phone=? WHERE id=?", test.Name, test.Phone, test.Id)
	}

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteTest(context *gin.Context) {
	context.JSON(200, gin.H{})
	var test model.Test
	context.BindJSON(&test.Id)
	fmt.Println(test.Id)
	//var Id = 1
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/aed_project")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Test WHERE id=?", test.Id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
