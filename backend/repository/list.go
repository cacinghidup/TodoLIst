package repository

import (
	"Moonlay/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type List interface {
	FindAll(int, int, int, string, string) ([]models.List, error)
	FindByID(ID int) (models.List, error)
	CreateList(list models.List) (models.List, error)
	UpdateList(list models.List) (models.List, error)
	DeleteList(ID int) (string, error)
	GetTotalRecords() (int, error)
}

type listRepository struct {
	db *gorm.DB
}

func ListRepository(db *gorm.DB) *listRepository {
	return &listRepository{db}
}

func (r *listRepository) FindAll(page, totalRecords, pageSize int, title, deskripsi string) ([]models.List, error) {
	var lists []models.List
	offset := (page - 1) * pageSize
	// 	// Menghitung totalRecords
	var count int64
	r.db.Model(&models.List{}).Count(&count)
	totalRecords = int(count)

	query := r.db.Preload("SubList.List")

	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}

	if deskripsi != "" {
		query = query.Where("LOWER(deskripsi) LIKE ?", "%"+strings.ToLower(deskripsi)+"%")
	}

	err := query.Limit(pageSize).Offset(offset).Find(&lists).Error

	return lists, err
}

func (r *listRepository) FindByID(ID int) (models.List, error) {
	var list models.List
	err := r.db.Preload("SubList.List").Where("id = ?", ID).First(&list).Error

	return list, err
}

func (r *listRepository) CreateList(list models.List) (models.List, error) {
	err := r.db.Create(&list).Error

	return list, err
}

func (r *listRepository) UpdateList(list models.List) (models.List, error) {
	err := r.db.Save(&list).Error

	return list, err
}

func (r *listRepository) DeleteList(ID int) (string, error) {
	var list models.List
	err := r.db.Where("id = ?", ID).Delete(&list).Error

	return `ID: ` + strconv.Itoa(ID) + ``, err
}

func (r *listRepository) GetTotalRecords() (int, error) {
	var count int64
	err := r.db.Model(&models.List{}).Count(&count).Error
	totalRecords := int(count)
	return totalRecords, err
}
