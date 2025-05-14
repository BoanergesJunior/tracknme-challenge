package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
	"github.com/redis/go-redis/v9"
)

func (r *repository) GetEmployeesByZipCodeRepository(zipCode string) ([]*model.EmployeeDTO, error) {
	ctx := context.Background()
	key := fmt.Sprintf("zipcode:%s", zipCode)

	employeesCache, err := r.redis.LRange(ctx, key, 0, -1).Result()
	if err == nil && len(employeesCache) > 0 {
		employees := make([]*model.EmployeeDTO, 0, len(employeesCache))
		for _, empStr := range employeesCache {
			var employee model.EmployeeDTO
			if err := json.Unmarshal([]byte(empStr), &employee); err != nil {
				fmt.Printf("Failed to unmarshal employee from cache: %v\n", err)
				continue
			}
			employees = append(employees, &employee)
		}
		if len(employees) > 0 {
			return employees, nil
		}
	} else if err != nil && err != redis.Nil {
		fmt.Printf("Cache error: %v\n", err)
	}

	var employees []*model.EmployeeDTO
	if err := r.db.Table(helpers.Employees).Where("zip_code = ?", zipCode).Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}
