package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"warranty/api/errs"
	"warranty/api/models"
	"warranty/api/services"
	"warranty/database"

	"github.com/gin-gonic/gin"
)

// Customer handlers

func GetCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomersService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}

	newCustomer, err := services.CreateCustomerService(c.Request.Context(), database.GetDB(), &customer)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newCustomer)
}

func GetCustomerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer, err := services.GetCustomerByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}
	customer.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedCustomer, err := services.UpdateCustomerService(c.Request.Context(), database.GetDB(), &customer)
	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedCustomer)
}

func DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	err = services.DeleteCustomerService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
