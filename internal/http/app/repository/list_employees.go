package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/redis/go-redis/v9"
)

const AllEmployeesCacheKey = "employees:all"

func (repo *repository) ListEmployeesRepository() ([]*model.EmployeeDTO, error) {
	ctx := context.Background()

	employeesCache, err := repo.redis.Get(ctx, AllEmployeesCacheKey).Result()
	if err == nil {
		var employees []*model.EmployeeDTO
		if err := json.Unmarshal([]byte(employeesCache), &employees); err == nil {
			return employees, nil
		}
		fmt.Printf("Failed to unmarshal cache: %v\n", err)
	} else if err != redis.Nil {
		fmt.Printf("Cache error: %v\n", err)
	}

	employees := []*model.EmployeeDTO{}
	if err := repo.db.Table(helpers.Employees).Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}
