package repository

import (
	"context"
	"fmt"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

func (repo *repository) DeleteEmployeeRepository(employeeID string) error {
	ctx := context.Background()

	if err := repo.db.Table(helpers.Employees).Where("id = ?", employeeID).Delete(&model.EmployeeDTO{}).Error; err != nil {
		return err
	}

	employeeKey := fmt.Sprintf("%s:%s", "id", employeeID)
	if err := repo.redis.Del(ctx, employeeKey, "employees:all").Err(); err != nil {
		fmt.Printf("Failed to invalidate cache: %v\n", err)
	}

	return nil
}
