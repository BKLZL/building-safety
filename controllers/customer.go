package controllers

import (
	"building-safety/config"
	"building-safety/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "客户不存在"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "客户不存在"})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Customer{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "客户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
