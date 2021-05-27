package category

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
func AllCategory(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT c_id, name,floor FROM category`)

	if err != nil {
		log.Print(err)
	}

	var category model.Category
	var arrcate []model.Category

	for rows.Next() {
		err = rows.Scan(&category.C_id, &category.Name, &category.Floor)
		fmt.Println(category)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrcate = append(arrcate, category)
		}
	}
	fmt.Println(arrcate)
	context.JSON(200, gin.H{"arrcate": arrcate})

}

// InsertTest = Insert Test API
func InsertCategory(context *gin.Context) {
	context.JSON(200, gin.H{})
	var category model.Category
	context.BindJSON(&category)
	fmt.Println(category)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO category(name,floor) VALUES(?,?)`, category.Name, category.Floor)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateCategory(context *gin.Context) {
	context.JSON(200, gin.H{})
	var category model.Category
	context.BindJSON(&category)

	fmt.Println(category)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if category.Name != "" {

		_, err = db.Exec("UPDATE category SET name=?,floor=? WHERE c_id=?", category.Name, category.Floor, category.C_id)
	}

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteCategory(context *gin.Context) {
	context.JSON(200, gin.H{})
	var category model.Category
	context.BindJSON(&category.C_id)
	fmt.Println(category.C_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM category WHERE c_id=?", category.C_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
