package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/helpers"
)

const (
	ZipCodeCacheExpiration = 24 * time.Hour
)

func (r *repository) GetZipCodeCache(ctx context.Context, key string) (string, error) {
	zipCodeCache, err := r.redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return "", err
	}

	if len(zipCodeCache) > 0 {
		return zipCodeCache[0], nil
	}

	return "", nil
}

func (r *repository) UpdateCache(employee model.EmployeeDTO, oldZipCode ...string) error {
	ctx := context.Background()
	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", helpers.EmployeeKeyPrefix, employee.ID)
	err = r.redis.Set(ctx, key, employeeJSON, 0).Err()
	if err != nil {
		return err
	}

	if len(oldZipCode) > 0 && oldZipCode[0] != "" && oldZipCode[0] != employee.ZipCode {
		oldKeyZipCode := fmt.Sprintf("%s:%s", helpers.AddressKeyPrefix, oldZipCode[0])
		employees, err := r.redis.LRange(ctx, oldKeyZipCode, 0, -1).Result()
		if err == nil {
			for _, empStr := range employees {
				var emp model.EmployeeDTO
				if err := json.Unmarshal([]byte(empStr), &emp); err == nil && emp.ID == employee.ID {
					r.redis.LRem(ctx, oldKeyZipCode, 0, empStr).Result()
					break
				}
			}
		}
	}

	keyZipCode := fmt.Sprintf("%s:%s", helpers.AddressKeyPrefix, employee.ZipCode)
	err = r.redis.RPush(ctx, keyZipCode, employeeJSON).Err()
	if err != nil {
		return err
	}

	r.redis.Del(ctx, helpers.AllEmployeesCacheKey).Err()

	return nil
}

func (r *repository) DeleteFromCache(employee model.EmployeeDTO) error {
	ctx := context.Background()

	key := fmt.Sprintf("%s:%s", helpers.EmployeeKeyPrefix, employee.ID)
	r.redis.Del(ctx, key).Err()

	keyZipCode := fmt.Sprintf("%s:%s", helpers.AddressKeyPrefix, employee.ZipCode)
	employeeJSON, _ := json.Marshal(employee)
	r.redis.LRem(ctx, keyZipCode, 0, string(employeeJSON)).Result()

	r.redis.Del(ctx, helpers.AllEmployeesCacheKey).Err()

	return nil
}
