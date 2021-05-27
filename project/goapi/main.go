package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"crud/controller/account"
	"crud/controller/brand"
	"crud/controller/category"
	"crud/controller/employee"
	"crud/controller/member"
	"crud/controller/parkrecord"
	"crud/controller/transactionrecord"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/getCategory", category.AllCategory)
	router.POST("/insertCategory", category.InsertCategory)
	router.PUT("/updateCategory", category.UpdateCategory)
	router.DELETE("/deleteCategory", category.DeleteCategory)
	router.GET("/getBrand", brand.AllBrand)
	router.POST("/insertBrand", brand.InsertBrand)
	router.PUT("/updateBrand", brand.UpdateBrand)
	router.DELETE("/deleteBrand", brand.DeleteBrand)
	router.GET("/getEmployee", employee.AllEmployee)
	router.POST("/insertEmployee", employee.InsertEmployee)
	router.PUT("/updateEmployee", employee.UpdateEmployee)
	router.DELETE("/deleteEmployee", employee.DeleteEmployee)
	router.GET("/getParkrecord", parkrecord.AllParkrecord)
	router.POST("/insertParkrecord", parkrecord.InsertParkrecord)
	router.PUT("/updateParkrecord", parkrecord.UpdateParkrecord)
	router.DELETE("/deleteParkrecord", parkrecord.DeleteParkrecord)
	router.GET("/getMember", member.AllMember)
	router.POST("/insertMember", member.InsertMember)
	router.DELETE("/deleteMember", member.DeleteMember)
	router.PUT("/updateMember", member.UpdateMember)
	router.GET("/getTransactionrecord", transactionrecord.AllTransactionrecord)
	router.POST("/insertTransactionrecord", transactionrecord.InsertTransactionrecord)
	router.DELETE("/deleteTransactionrecord", transactionrecord.DeleteTransactionrecord)
	router.PUT("/updateTransactionrecord", transactionrecord.UpdateTransactionrecord)
	router.GET("/getAccount", account.AllAccount)
	router.Run()
}
