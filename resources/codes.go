package resources

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"perx-go-test/models"
)

var (
	ErrCodeNotFound    = errors.New("code not found")
	ErrCodeAlreadyUsed = errors.New("code already used")
)

type CodesResource struct {
	db       *gorm.DB
	alphabet []rune
}

func NewCodesResource(db *gorm.DB, alphabet []rune) *CodesResource {
	return &CodesResource{
		db:       db,
		alphabet: alphabet,
	}
}

func (c *CodesResource) Create(size int) (*models.Code, error) {
	for i := 0; i < 1000; i++ {
		code := models.Code{
			Code: c.generateString(size),
		}

		tx := c.db.Begin()

		if result := tx.First(&models.Code{}, code); result.Error != nil {
			if !result.RecordNotFound() {
				tx.Rollback()
				return nil, errors.Wrap(result.Error, "create a new code failed")
			}
		} else {
			tx.Rollback()
			continue
		}

		if err := tx.Create(&code).Error; err != nil {
			tx.Rollback()
			return nil, errors.Wrap(err, "create a new code failed")
		}
		tx.Commit()

		return &code, nil
	}

	return nil, fmt.Errorf("max retries exceeded")
}

func (c *CodesResource) generateString(size int) string {
	result := make([]rune, size)
	for i := range result {
		result[i] = c.alphabet[rand.Intn(len(c.alphabet))]
	}

	return string(result)
}

func (c *CodesResource) Check(codeStr string) (bool, error) {
	tx := c.db.Begin()

	var code models.Code
	if result := tx.Where(models.Code{Code: codeStr}).First(&code); result.Error != nil {
		tx.Rollback()
		if result.RecordNotFound() {
			return false, ErrCodeNotFound
		}
		return false, errors.Wrap(result.Error, "a code check failed")
	}

	if code.UsedAt != nil {
		tx.Rollback()
		return false, ErrCodeAlreadyUsed
	}

	now := time.Now()
	code.UsedAt = &now

	if err := tx.Save(code).Error; err != nil {
		tx.Rollback()
		return false, errors.Wrap(err, "update code failed")
	}

	tx.Commit()

	return true, nil
}
