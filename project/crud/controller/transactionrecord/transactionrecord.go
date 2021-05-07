package transactionrecord

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
func AllTransactionrecord(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT t.t_id,t.time,t.product,t.amount,m.name,b.name
	from transactionrecord as t,
	brand as b,
	memberdetails as m
	where t.m_id = m.m_id && t.b_id = b.b_id`)

	if err != nil {
		log.Print(err)
	}

	var transname model.Transname
	var arrtrans []model.Transname

	for rows.Next() {

		err = rows.Scan(&transname.T_id, &transname.Time, &transname.Product, &transname.Amount, &transname.M_name, &transname.B_name)
		fmt.Println(transname)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrtrans = append(arrtrans, transname)
		}
	}
	fmt.Println(arrtrans)
	context.JSON(200, gin.H{"arrtrans": arrtrans})

}

// InsertTest = Insert Test API
func InsertTransactionrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var transactionrecord model.Transactionrecord
	context.BindJSON(&transactionrecord)
	fmt.Println(transactionrecord)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO transactionrecord(time,product,amount,m _id,b_id) VALUES(?,?,?,?,?)`, transactionrecord.Time, transactionrecord.Product, transactionrecord.Amount, transactionrecord.M_id, transactionrecord.B_id)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print("Insert data to database")
}

// UpdateTest = Update Test API
func UpdateTransactionrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var transactionrecord model.Transactionrecord
	context.BindJSON(&transactionrecord)

	fmt.Println(transactionrecord)
	// fmt.Println(id)
	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE transactionrecord SET time=?,product=?,amount=?,m_id=?,b_id=? WHERE t_id=?", transactionrecord.Time, transactionrecord.Product, transactionrecord.Amount, transactionrecord.M_id, transactionrecord.B_id)

	if err != nil {
		log.Print(err)
	}

}

// DeleteTest = Delete test API
func DeleteTransactionrecord(context *gin.Context) {
	context.JSON(200, gin.H{})
	var transactionrecord model.Transactionrecord
	context.BindJSON(&transactionrecord.T_id)
	fmt.Println(transactionrecord.T_id)

	db := config.Connect()

	db, err := sql.Open("mysql",
		"root:1000sean334@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM employee WHERE t_id=?", transactionrecord.T_id)

	if err != nil {
		log.Print(err)
		return
	}

	fmt.Print("Delete data successfully")

}
