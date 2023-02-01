package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
	"reflect"
)

// SetLocal puts custom data to fiber handler context.
func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

// GetLocal retrieves custom data stored in fiber handler context.
func GetLocal[T any](c *fiber.Ctx, key string) (T, error) {
	var result T
	val := c.Locals(key)
	if val == nil {
		return result, eris.New("Could not get value from fiber context")
	}
	result = c.Locals(key).(T)
	if reflect.ValueOf(&result).Elem().IsZero() {
		return result, eris.New("Value in fiber context is of wrong type")
	}
	return result, nil
}
