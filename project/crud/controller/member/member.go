package member

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
func AllMember(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`select mem.m_id,mem.level,memd.name,memd.phone,memd.address 
	from member as mem,
	memberdetails as memd
	where mem.m_id = memd.m_id`)

	if err != nil {
		log.Print(err)
	}

	// var member model.Member
	// var memberdetail model.Memberdetail
	var memberall model.Memberall
	var arrmemberall []model.Memberall

	for rows.Next() {
		err = rows.Scan(&memberall.M_id, &memberall.Level, &memberall.Name, &memberall.Phone, &memberall.Address)
		//fmt.Println(employee)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrmemberall = append(arrmemberall, memberall)
		}
	}
	fmt.Println(arrmemberall)
	context.JSON(200, gin.H{"arrmemberall": arrmemberall})

}

// InsertTest = Insert Test API
func InsertMember(context *gin.Context) {
	context.JSON(200, gin.H{})
	var memberall model.Memberall
	context.BindJSON(&memberall)
	fmt.Println(memberall)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO member(level) VALUES(?)`, memberall.Level)
	_, err = db.Exec(`INSERT INTO memberdetails(name,phone,address) VALUES(?,?,?)`, memberall.Name, memberall.Phone, memberall.Address)

	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// // UpdateTest = Update Test API
func UpdateMember(context *gin.Context) {
	context.JSON(200, gin.H{})
	var memberall model.Memberall

	context.BindJSON(&memberall)

	fmt.Println(memberall)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE member SET level=? WHERE m_id=?", memberall.Level, memberall.M_id)
	_, err = db.Exec("UPDATE memberdetails SET name=?,phone=?,address=? WHERE m_id=?", memberall.Name, memberall.Phone, memberall.Address, memberall.M_id)

	if err != nil {
		log.Print(err)
	}

}

// // DeleteTest = Delete test API
func DeleteMember(context *gin.Context) {
	context.JSON(200, gin.H{})
	var memberall model.Memberall
	context.BindJSON(&memberall.M_id)
	fmt.Println(memberall.M_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM member WHERE m_id=?", memberall.M_id)
	_, err = db.Exec("DELETE FROM memberdetails WHERE m_id=?", memberall.M_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
