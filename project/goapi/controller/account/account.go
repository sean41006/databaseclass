package account

import (
	"crud/model"
	"fmt"
	"log"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// AllTest = Select Test API
func AllAccount(context *gin.Context) {

	db, err := sql.Open("mysql",
		"root:password@/department")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id,account,password FROM account`)

	if err != nil {
		log.Print(err)
	}

	var account model.Account
	var arraccount []model.Account

	for rows.Next() {
		err = rows.Scan(&account.Id, &account.Account, &account.Password)
		fmt.Println(account)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arraccount = append(arraccount, account)
		}
	}
	fmt.Println(arraccount)
	context.JSON(200, gin.H{"arraccount": arraccount})

}
