package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
)

const (
	ZipCodeCacheExpiration = 24 * time.Hour
)

func (r *repository) GetZipCodeCache(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *repository) UpdateCache(employee model.EmployeeDTO) error {
	key := fmt.Sprintf("id:%s", employee.ID)
	err := r.redis.Set(context.Background(), key, employee, 0).Err()
	if err != nil {
		return err
	}

	keyZipCode := fmt.Sprintf("zipcode:%s", employee.ZipCode)
	err = r.redis.RPush(context.Background(), keyZipCode, employee).Err()
	if err != nil {
		return err
	}

	return nil
}
