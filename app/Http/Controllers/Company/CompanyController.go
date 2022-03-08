package Controller

import (
	"net/http"
	"strconv"
	CompanyModel "tourism/app/Models/Company"

	"github.com/labstack/echo/v4"
)

var NewCompany CompanyModel.Company

func Index(context echo.Context) error {
	companies := CompanyModel.GetAllCompanies()

	return context.JSON(http.StatusOK, companies)
}

func Show(context echo.Context) error {
	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	companyId := int(number)
	company := CompanyModel.GetCompanyById(companyId)

	return context.JSON(http.StatusOK, company)
}

func Store(context echo.Context) error {
	var newCompany CompanyModel.Company
	context.Bind(&newCompany)
	CompanyModel.CreateCompany(&newCompany)

	return context.JSON(http.StatusCreated, newCompany)
}

func Update(context echo.Context) error {
	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	companyId := int(number)

	context.Bind(&NewCompany)
	CompanyModel.UpdateCompany(&NewCompany, companyId)

	return context.JSON(http.StatusOK, NewCompany)
}

func Destroy(context echo.Context) error {
	number, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	companyId := int(number)

	CompanyModel.DeleteCompany(companyId)

	return context.NoContent(http.StatusNoContent)
}
