package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/antcebolla/web-server/src/database"
	"gorm.io/gorm"
)

var ApptStatuses = [3]string{"pending", "approved", "rejected"}

type Appointment struct {
	gorm.Model
	Date               string `json:"date" gorm:"type:date not null"`
	StartTime          string `json:"start_time" gorm:"type:time not null"`
	EndTime            string `json:"end_time" gorm:"type:time not null"`
	Status             string `json:"status" gorm:"not null"`
	VeterinarianID     sql.NullInt64
	VeterinaryCenterID uint
	Animals            []Animal `gorm:"many2many:appointment_animals;"`
}

func (a *Appointment) ValidateAndFormat() error {
	//check if stats is valid
	var isValidStatus bool
	for _, status := range ApptStatuses {
		if status == a.Status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return gorm.ErrInvalidData
	}

	// Definimos los formatos estándar que envía un navegador
	const dateFormat = "2006-01-02" // YYYY-MM-DD
	const timeFormat = "15:04"      // HH:MM

	// 1. Unimos la fecha y la hora para crear objetos time.Time completos
	// Esto nos permite comparar turnos de forma precisa (día y hora)
	startStr := fmt.Sprintf("%s %s", a.Date, a.StartTime)
	endStr := fmt.Sprintf("%s %s", a.Date, a.EndTime)

	newApptStart, err := time.Parse(fmt.Sprintf("%s %s", dateFormat, timeFormat), startStr)
	if err != nil {
		return fmt.Errorf("invalid start date/time format: %w", err)
	}

	newApptEnd, err := time.Parse(fmt.Sprintf("%s %s", dateFormat, timeFormat), endStr)
	if err != nil {
		return fmt.Errorf("invalid end date/time format: %w", err)
	}

	// 2. Verificamos que el inicio no sea igual o después del final
	if !newApptStart.Before(newApptEnd) {
		return gorm.ErrInvalidData
	}

	var appts []Appointment
	// Buscamos los turnos existentes del mismo veterinario, día y centro
	database.DB.Where(&Appointment{
		VeterinarianID:     a.VeterinarianID,
		Date:               a.Date,
		VeterinaryCenterID: a.VeterinaryCenterID,
	}).Find(&appts)

	// 3. Iteramos sobre los turnos existentes para buscar solapamientos
	for _, appt := range appts {
		// Creamos objetos time.Time para el turno existente
		existingApptStart, err := time.Parse(fmt.Sprintf("%s %s", dateFormat, timeFormat), fmt.Sprintf("%s %s", appt.Date, appt.StartTime))
		if err != nil {
			return fmt.Errorf("invalid existing appointment start date/time format: %w", err)
		}
		existingApptEnd, err := time.Parse(fmt.Sprintf("%s %s", dateFormat, timeFormat), fmt.Sprintf("%s %s", appt.Date, appt.EndTime))
		if err != nil {
			return fmt.Errorf("invalid existing appointment end date/time format: %w", err)
		}

		// Verificamos el solapamiento con los métodos del paquete time
		if !(existingApptStart.After(newApptEnd) || existingApptEnd.Before(newApptStart)) {
			return gorm.ErrInvalidData
		}
	}

	return nil
}
