package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// Cache keys prefixes
	EmployeeKeyPrefix = "employee:"
	ZipCodeKeyPrefix  = "zipcode:"

	// Cache expiration times
	EmployeeCacheExpiration = 24 * time.Hour
	ZipCodeCacheExpiration  = 24 * time.Hour
)

type CacheService struct {
	client *redis.Client
}

func NewCacheService(client *redis.Client) *CacheService {
	return &CacheService{
		client: client,
	}
}

// GetEmployee retrieves an employee from cache
func (s *CacheService) GetEmployee(ctx context.Context, employeeID string) ([]byte, error) {
	key := fmt.Sprintf("%s%s", EmployeeKeyPrefix, employeeID)
	return s.client.Get(ctx, key).Bytes()
}

// SetEmployee stores an employee in cache
func (s *CacheService) SetEmployee(ctx context.Context, employeeID string, data interface{}) error {
	key := fmt.Sprintf("%s%s", EmployeeKeyPrefix, employeeID)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, key, jsonData, EmployeeCacheExpiration).Err()
}

// DeleteEmployee removes an employee from cache
func (s *CacheService) DeleteEmployee(ctx context.Context, employeeID string) error {
	key := fmt.Sprintf("%s%s", EmployeeKeyPrefix, employeeID)
	return s.client.Del(ctx, key).Err()
}

// GetZipCode retrieves zipcode data from cache
func (s *CacheService) GetZipCode(ctx context.Context, zipCode string) ([]byte, error) {
	key := fmt.Sprintf("%s%s", ZipCodeKeyPrefix, zipCode)
	return s.client.Get(ctx, key).Bytes()
}

// SetZipCode stores zipcode data in cache
func (s *CacheService) SetZipCode(ctx context.Context, zipCode string, data interface{}) error {
	key := fmt.Sprintf("%s%s", ZipCodeKeyPrefix, zipCode)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, key, jsonData, ZipCodeCacheExpiration).Err()
}

// DeleteZipCode removes zipcode data from cache
func (s *CacheService) DeleteZipCode(ctx context.Context, zipCode string) error {
	key := fmt.Sprintf("%s%s", ZipCodeKeyPrefix, zipCode)
	return s.client.Del(ctx, key).Err()
}

// DeleteAllEmployees removes all employee cache entries
func (s *CacheService) DeleteAllEmployees(ctx context.Context) error {
	pattern := fmt.Sprintf("%s*", EmployeeKeyPrefix)
	iter := s.client.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		if err := s.client.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}
