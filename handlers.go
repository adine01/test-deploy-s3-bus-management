package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Bus represents a bus in the system
type Bus struct {
	ID          int       `json:"id" db:"id"`
	PlateNumber string    `json:"plate_number" db:"plate_number"`
	Model       string    `json:"model" db:"model"`
	Capacity    int       `json:"capacity" db:"capacity"`
	Status      string    `json:"status" db:"status"` // active, maintenance, retired
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Staff represents a staff member
type Staff struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Position  string    `json:"position" db:"position"` // driver, conductor, mechanic
	LicenseNo string    `json:"license_no,omitempty" db:"license_no"`
	Status    string    `json:"status" db:"status"` // active, inactive
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Request structs
type CreateBusRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	Model       string `json:"model" binding:"required"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
}

type CreateStaffRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Position  string `json:"position" binding:"required"`
	LicenseNo string `json:"license_no,omitempty"`
}

// Bus handlers
func handleCreateBus(c *gin.Context) {
	var req CreateBusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bus := Bus{
		PlateNumber: req.PlateNumber,
		Model:       req.Model,
		Capacity:    req.Capacity,
		Status:      "active",
	}

	if err := CreateBus(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bus"})
		return
	}

	c.JSON(http.StatusCreated, bus)
}

func handleGetBuses(c *gin.Context) {
	buses, err := GetAllBuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve buses"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buses": buses, "count": len(buses)})
}

func handleGetBus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	bus, err := GetBusByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if bus == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bus not found"})
		return
	}

	c.JSON(http.StatusOK, bus)
}

func handleUpdateBus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	// Check if bus exists
	existingBus, err := GetBusByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if existingBus == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bus not found"})
		return
	}

	var req CreateBusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update bus fields
	existingBus.PlateNumber = req.PlateNumber
	existingBus.Model = req.Model
	existingBus.Capacity = req.Capacity

	if err := UpdateBus(existingBus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bus"})
		return
	}

	c.JSON(http.StatusOK, existingBus)
}

func handleDeleteBus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	// Check if bus exists
	existingBus, err := GetBusByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if existingBus == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bus not found"})
		return
	}

	if err := DeleteBus(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bus deleted successfully"})
}

// Staff handlers
func handleCreateStaff(c *gin.Context) {
	var req CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staffMember := Staff{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Position:  req.Position,
		LicenseNo: req.LicenseNo,
		Status:    "active",
	}

	if err := CreateStaff(&staffMember); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff member"})
		return
	}

	c.JSON(http.StatusCreated, staffMember)
}

func handleGetStaff(c *gin.Context) {
	staff, err := GetAllStaff()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve staff"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staff": staff, "count": len(staff)})
}

func handleGetStaffMember(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	staffMember, err := GetStaffByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if staffMember == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff member not found"})
		return
	}

	c.JSON(http.StatusOK, staffMember)
}

func handleUpdateStaff(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Check if staff member exists
	existingStaff, err := GetStaffByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if existingStaff == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff member not found"})
		return
	}

	var req CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update staff fields
	existingStaff.Name = req.Name
	existingStaff.Email = req.Email
	existingStaff.Phone = req.Phone
	existingStaff.Position = req.Position
	existingStaff.LicenseNo = req.LicenseNo

	if err := UpdateStaff(existingStaff); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update staff member"})
		return
	}

	c.JSON(http.StatusOK, existingStaff)
}

func handleDeleteStaff(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	// Check if staff member exists
	existingStaff, err := GetStaffByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if existingStaff == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff member not found"})
		return
	}

	if err := DeleteStaff(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete staff member"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff member deleted successfully"})
}
