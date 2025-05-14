package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/errors"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (repo *repository) GetEmployeeRepository(employeeID string) (*model.EmployeeDTO, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%s", "id", employeeID)

	employeeCache, err := repo.redis.Get(ctx, key).Result()
	if err == nil {
		var employee model.EmployeeDTO
		if err := json.Unmarshal([]byte(employeeCache), &employee); err == nil {
			return &employee, nil
		}
	} else if err != redis.Nil {
		fmt.Printf("Cache error: %v\n", err)
	}

	employee := new(model.EmployeeDTO)
	repo.db.Table(helpers.Employees).Where("id = ?", employeeID).First(&employee)
	if employee.ID == uuid.Nil {
		return nil, errors.NewAppError(http.StatusNotFound, "Employee not found", errors.ErrNotFound)
	}

	return employee, nil
}
