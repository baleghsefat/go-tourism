package models

import (
	"tourism/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Company struct {
	gorm.Model
	IsActive *bool  `gorm:"default:true" json:"is_active"`
	Slug     string `gorm:"unique; not null; type:nvarchar(50)" json:"slug"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Company{})
}

func GetAllCompanies() []Company {
	var companies []Company
	db.Find(&companies)
	return companies
}

func GetCompanyById(companyId int) Company {
	var company Company
	db.First(&company, companyId)
	return company
}

func CreateCompany(company *Company) *gorm.DB {
	result := db.Create(&company)

	return result
}

func UpdateCompany(company *Company, companyId int) *gorm.DB {
	newData := Company{
		Slug:     company.Slug,
		IsActive: company.IsActive,
	}

	result := db.Model(Company{}).Where("ID = ?", companyId).Updates(&newData)

	return result
}

func DeleteCompany(companyId int) {
	db.Delete(&Company{}, companyId)
}
