package utils

import (
	"math"

	"go-echo-hexagonal/internal/core/domain"

	"gorm.io/gorm"
)

func Paginator(db *gorm.DB, page, limit int, model interface{}) (*domain.Paginator, error) {
	offset := (page - 1) * limit

	var totalRecord int64
	if err := db.Model(model).Count(&totalRecord).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(offset).Limit(limit).Find(model).Error; err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(limit)))

	var prevPage, nextPage int
	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = page
	}

	if page < totalPage {
		nextPage = page + 1
	} else {
		nextPage = page
	}

	return &domain.Paginator{
		TotalRecord: totalRecord,
		TotalPage:   totalPage,
		Records:     model,
		Offset:      offset,
		Limit:       limit,
		Page:        page,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}, nil
}