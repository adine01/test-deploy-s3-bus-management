package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB initializes the database connection pool
func InitDB() error {
	var err error
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Create connection pool
	db, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return err
	}

	// Test the connection
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	log.Println("Connected to Supabase database")

	// Create tables if they don't exist
	if err := createTables(); err != nil {
		return err
	}

	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// createTables creates the buses and staff tables if they don't exist
func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS buses (
		id SERIAL PRIMARY KEY,
		plate_number VARCHAR(20) UNIQUE NOT NULL,
		model VARCHAR(100) NOT NULL,
		capacity INTEGER NOT NULL CHECK (capacity > 0),
		status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'maintenance', 'retired')),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS staff (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR(20) NOT NULL,
		position VARCHAR(50) NOT NULL CHECK (position IN ('driver', 'conductor', 'mechanic')),
		license_no VARCHAR(50),
		status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	-- Create indexes for better performance
	CREATE INDEX IF NOT EXISTS idx_buses_plate_number ON buses(plate_number);
	CREATE INDEX IF NOT EXISTS idx_buses_status ON buses(status);
	CREATE INDEX IF NOT EXISTS idx_staff_email ON staff(email);
	CREATE INDEX IF NOT EXISTS idx_staff_position ON staff(position);
	CREATE INDEX IF NOT EXISTS idx_staff_status ON staff(status);
	`

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return err
	}

	log.Println("Bus and Staff tables created successfully")
	return nil
}

// Bus database operations

// CreateBus inserts a new bus into the database
func CreateBus(bus *Bus) error {
	query := `
		INSERT INTO buses (plate_number, model, capacity, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := db.QueryRow(context.Background(), query, bus.PlateNumber, bus.Model, bus.Capacity, bus.Status).
		Scan(&bus.ID, &bus.CreatedAt, &bus.UpdatedAt)

	return err
}

// GetBusByID retrieves a bus by ID
func GetBusByID(id int) (*Bus, error) {
	bus := &Bus{}
	query := `
		SELECT id, plate_number, model, capacity, status, created_at, updated_at
		FROM buses
		WHERE id = $1
	`

	err := db.QueryRow(context.Background(), query, id).
		Scan(&bus.ID, &bus.PlateNumber, &bus.Model, &bus.Capacity, &bus.Status, &bus.CreatedAt, &bus.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Bus not found
		}
		return nil, err
	}

	return bus, nil
}

// GetAllBuses retrieves all buses from the database
func GetAllBuses() ([]Bus, error) {
	var buses []Bus
	query := `
		SELECT id, plate_number, model, capacity, status, created_at, updated_at
		FROM buses
		ORDER BY created_at DESC
	`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bus Bus
		err := rows.Scan(&bus.ID, &bus.PlateNumber, &bus.Model, &bus.Capacity, &bus.Status, &bus.CreatedAt, &bus.UpdatedAt)
		if err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}

	return buses, nil
}

// UpdateBus updates an existing bus
func UpdateBus(bus *Bus) error {
	query := `
		UPDATE buses
		SET plate_number = $1, model = $2, capacity = $3, status = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING updated_at
	`

	err := db.QueryRow(context.Background(), query, bus.PlateNumber, bus.Model, bus.Capacity, bus.Status, bus.ID).
		Scan(&bus.UpdatedAt)

	return err
}

// DeleteBus deletes a bus by ID
func DeleteBus(id int) error {
	query := `DELETE FROM buses WHERE id = $1`
	_, err := db.Exec(context.Background(), query, id)
	return err
}

// Staff database operations

// CreateStaff inserts a new staff member into the database
func CreateStaff(staff *Staff) error {
	query := `
		INSERT INTO staff (name, email, phone, position, license_no, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := db.QueryRow(context.Background(), query, staff.Name, staff.Email, staff.Phone,
		staff.Position, staff.LicenseNo, staff.Status).
		Scan(&staff.ID, &staff.CreatedAt, &staff.UpdatedAt)

	return err
}

// GetStaffByID retrieves a staff member by ID
func GetStaffByID(id int) (*Staff, error) {
	staff := &Staff{}
	query := `
		SELECT id, name, email, phone, position, license_no, status, created_at, updated_at
		FROM staff
		WHERE id = $1
	`

	err := db.QueryRow(context.Background(), query, id).
		Scan(&staff.ID, &staff.Name, &staff.Email, &staff.Phone, &staff.Position,
			&staff.LicenseNo, &staff.Status, &staff.CreatedAt, &staff.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Staff not found
		}
		return nil, err
	}

	return staff, nil
}

// GetAllStaff retrieves all staff members from the database
func GetAllStaff() ([]Staff, error) {
	var staff []Staff
	query := `
		SELECT id, name, email, phone, position, license_no, status, created_at, updated_at
		FROM staff
		ORDER BY created_at DESC
	`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Staff
		err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Phone, &s.Position,
			&s.LicenseNo, &s.Status, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		staff = append(staff, s)
	}

	return staff, nil
}

// UpdateStaff updates an existing staff member
func UpdateStaff(staff *Staff) error {
	query := `
		UPDATE staff
		SET name = $1, email = $2, phone = $3, position = $4, license_no = $5, status = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
		RETURNING updated_at
	`

	err := db.QueryRow(context.Background(), query, staff.Name, staff.Email, staff.Phone,
		staff.Position, staff.LicenseNo, staff.Status, staff.ID).
		Scan(&staff.UpdatedAt)

	return err
}

// DeleteStaff deletes a staff member by ID
func DeleteStaff(id int) error {
	query := `DELETE FROM staff WHERE id = $1`
	_, err := db.Exec(context.Background(), query, id)
	return err
}
