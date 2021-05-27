package brand

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
func AllBrand(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT b_id, name,floor,position,c_id FROM brand`)

	if err != nil {
		log.Print(err)
	}

	var brand model.Brand
	var arrbrand []model.Brand

	for rows.Next() {
		err = rows.Scan(&brand.B_id, &brand.Name, &brand.Floor, &brand.Position, &brand.C_id)
		fmt.Println(brand)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrbrand = append(arrbrand, brand)
		}
	}
	fmt.Println(arrbrand)
	context.JSON(200, gin.H{"arrbrand": arrbrand})

}

// InsertTest = Insert Test API
func InsertBrand(context *gin.Context) {
	context.JSON(200, gin.H{})
	var brand model.Brand
	context.BindJSON(&brand)
	fmt.Println(brand)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO brand(name,floor,position,c_id) VALUES(?,?,?,?)`, brand.Name, brand.Floor, brand.Position, brand.C_id)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateBrand(context *gin.Context) {
	context.JSON(200, gin.H{})
	var brand model.Brand
	context.BindJSON(&brand)

	fmt.Println(brand)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if brand.Name != "" {

		_, err = db.Exec("UPDATE brand SET name=?,floor=?,position=?,c_id=? WHERE b_id=?", brand.Name, brand.Floor, brand.Position, brand.C_id, brand.B_id)
	}

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteBrand(context *gin.Context) {
	context.JSON(200, gin.H{})
	var brand model.Brand
	context.BindJSON(&brand.B_id)
	fmt.Println(brand.B_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM brand WHERE b_id=?", brand.B_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
