package parkrecord

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
func AllParkrecord(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT p_id, carnumber,parktime,parkcard_id FROM parkrecord`)

	if err != nil {
		log.Print(err)
	}

	var parkrecord model.Parkrecord
	var arrparkrecord []model.Parkrecord

	for rows.Next() {
		err = rows.Scan(&parkrecord.P_id, &parkrecord.Carnum, &parkrecord.Parktime, &parkrecord.Packcard_id)
		fmt.Println(parkrecord)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrparkrecord = append(arrparkrecord, parkrecord)
		}
	}
	fmt.Println(arrparkrecord)
	context.JSON(200, gin.H{"arrparkrecord": arrparkrecord})

}

// InsertTest = Insert Test API
func InsertParkrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var parkrecord model.Parkrecord
	context.BindJSON(&parkrecord)
	fmt.Println(parkrecord)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO parkrecord(carnumber,parktime,parkcard_id) VALUES(?,?,?)`, parkrecord.Carnum, parkrecord.Parktime, parkrecord.Packcard_id)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateParkrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var parkrecord model.Parkrecord
	context.BindJSON(&parkrecord)

	fmt.Println(parkrecord)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if parkrecord.Carnum != "" {

		_, err = db.Exec("UPDATE Parkrecord SET carnumber=?,parktime=?,parkcard_id=? WHERE p_id=?", parkrecord.Carnum, parkrecord.Parktime, parkrecord.Packcard_id, parkrecord.P_id)
	}

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteParkrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var parkrecord model.Parkrecord
	context.BindJSON(&parkrecord.P_id)
	fmt.Println(parkrecord.P_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM parkrecord WHERE p_id=?", parkrecord.P_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
