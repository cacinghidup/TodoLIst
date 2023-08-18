package repository

import (
	"Moonlay/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SubList interface {
	FindAll(int, int, int, string, string) ([]models.SubList, error)
	FindByID(ID int) (models.SubList, error)
	CreateSubList(subList models.SubList) (models.SubList, error)
	UpdateSubList(subList models.SubList) (models.SubList, error)
	DeleteSubList(ID int) (string, error)
	GetTotalRecordsSubList() (int, error)
}

type subListRepository struct {
	db *gorm.DB
}

func SubListRepository(db *gorm.DB) *subListRepository {
	return &subListRepository{db}
}

func (r *subListRepository) FindAll(page, totalRecords, pageSize int, title, deskripsi string) ([]models.SubList, error) {
	var sublists []models.SubList
	offset := (page - 1) * pageSize
	// Menghitung totalRecords
	var count int64
	r.db.Model(&models.SubList{}).Count(&count)
	totalRecords = int(count)

	query := r.db.Preload("List")

	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}

	if deskripsi != "" {
		query = query.Where("LOWER(deskripsi) LIKE ?", "%"+strings.ToLower(deskripsi)+"%")
	}

	err := query.Limit(pageSize).Offset(offset).Find(&sublists).Error

	return sublists, err
}

func (r *subListRepository) FindByID(ID int) (models.SubList, error) {
	var subList models.SubList
	err := r.db.Preload("List").Where("id = ?", ID).First(&subList).Error

	return subList, err
}

func (r *subListRepository) CreateSubList(subList models.SubList) (models.SubList, error) {
	err := r.db.Preload("List").Create(&subList).Error

	return subList, err
}

func (r *subListRepository) UpdateSubList(subList models.SubList) (models.SubList, error) {
	err := r.db.Save(&subList).Error

	return subList, err
}

func (r *subListRepository) DeleteSubList(ID int) (string, error) {
	var subList models.SubList
	err := r.db.Where("id = ?", ID).Delete(&subList).Error

	return `ID:` + strconv.Itoa(ID) + ``, err
}

func (r *subListRepository) GetTotalRecordsSubList() (int, error) {
	var count int64
	err := r.db.Model(&models.SubList{}).Count(&count).Error
	totalRecords := int(count)
	return totalRecords, err
}
