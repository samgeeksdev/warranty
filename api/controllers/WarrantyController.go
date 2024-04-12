package controllers

//
//func GetAllWarranties(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	warranties, err := services.GetAllWarrantiesService(ctx, database.GetDB())
//	if err != nil {
//		switch err {
//		case database.ErrRecordNotFound:
//			c.JSON(http.StatusNotFound, gin.H{"error": "No warranties found"})
//		default:
//			// Handle other unexpected errors (log for debugging)
//			log.Printf("Error getting all warranties: %v\n", err)
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		}
//		return
//	}
//
//	c.JSON(http.StatusOK, warranties)
//}
//
//func GetWarrantyByID(c *gin.Context) {
//	ctx := c.Request.Context()
//	warrantyIDStr := c.Param("id")
//
//	warrantyID, err := strconv.ParseUint(warrantyIDStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID format"})
//		return
//	}
//
//	warranty, err := services.GetWarrantyByIDService(ctx, database.GetDB(), uint(warrantyID))
//	if err != nil {
//		if errors.Is(err, errs.ErrWarrantyNotFound) {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Warranty not found"})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//
//	c.JSON(http.StatusOK, warranty)
//}
//
//func CreateWarranty(c *gin.Context) {
//	ctx := c.Request.Context()
//
//	var warranty models.Warranty
//	if err := c.BindJSON(&warranty); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty data"})
//		return
//	}
//
//	newWarranty, err := services.CreateWarrantyService(ctx, database.GetDB(), &warranty)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusCreated, newWarranty)
//}
//
//func UpdateWarranty(c *gin.Context) {
//	ctx := c.Request.Context()
//	warrantyIDStr := c.Param("id")
//
//	warrantyID, err := strconv.ParseUint(warrantyIDStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID format"})
//		return
//	}
//
//	var updatedWarranty models.Warranty
//	if err := c.BindJSON(&updatedWarranty); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty data"})
//		return
//	}
//	updatedWarranty.ID = int64(uint(warrantyID)) // Assuming ID is stored as int64
//
//	updatedWarranty, err = services.UpdateWarrantyService(ctx, database.GetDB(), &updatedWarranty)
//	if err != nil {
//		if errors.Is(err, errs.ErrWarrantyNotFound) {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Warranty not found"})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//
//	c.JSON(http.StatusOK, updatedWarranty)
//}
//
//func DeleteWarranty(c *gin.Context) {
//	ctx := c.Request.Context()
//	warrantyIDStr := c.Param("id")
//
//	warrantyID, err := strconv.ParseUint(warrantyIDStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID format"})
//		return
//	}
//
//	err = services.DeleteWarrantyService(ctx, database.GetDB(), uint(warrantyID))
//	if err != nil {
//		if errors.Is(err, errs.ErrWarrantyNotFound) {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Warranty not found"})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//
//	c.JSON(http.StatusNoContent, nil) // Respond with No Content (204) on successful deletion
//}
