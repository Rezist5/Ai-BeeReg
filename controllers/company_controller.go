// controllers/company_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/Rezist5/Ai-BeeReg/services"
	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	CompanyService *services.CompanyService
	UserService    *services.UserService
}

func NewCompanyController(companyService *services.CompanyService, userService *services.UserService) *CompanyController {
	return &CompanyController{CompanyService: companyService, UserService: userService}
}

func (ctrl *CompanyController) CreateCompany(c *gin.Context) {
	var req struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := ctrl.CompanyService.CreateCompany(req.Name, req.Description, req.Images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
}

func (ctrl *CompanyController) GetAllCompanies(c *gin.Context) {
	companies, err := ctrl.CompanyService.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении компаний"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (ctrl *CompanyController) GetCompanyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := ctrl.CompanyService.GetCompanyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (ctrl *CompanyController) UpdateCompany(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var req struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := ctrl.CompanyService.GetCompanyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	updatedCompany, err := ctrl.CompanyService.UpdateCompany(company, req.Name, req.Description, req.Images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCompany)
}

func (ctrl *CompanyController) DeleteCompany(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := ctrl.CompanyService.GetCompanyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	if err := ctrl.CompanyService.DeleteCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении компании"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Компания успешно удалена"})
}

func (c *CompanyController) AssignAdmin(ctx *gin.Context) {
	var input struct {
		AdminID uint `json:"adminId" binding:"required"`
	}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Преобразование id из строки в uint
	companyID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компании"})
		return
	}

	company, err := c.CompanyService.GetCompanyByID(uint(companyID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена"})
		return
	}

	admin, err := c.UserService.GetUserByID(input.AdminID) // Предполагается, что у вас есть метод для получения пользователя по ID
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Администратор не найден"})
		return
	}

	// Добавляем администратора к компании
	company.Admins = append(company.Admins, *admin) // Разыменовываем admin для получения значения

	if err := c.CompanyService.AssignAdmin(company, admin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при привязке администратора к компании"})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// RemoveAdmin отвязывает администратора от компании
func (c *CompanyController) RemoveAdmin(ctx *gin.Context) {
	var input struct {
		AdminID uint `json:"adminId" binding:"required"`
	}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Преобразование id из строки в uint
	companyID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компании"})
		return
	}

	// Получаем компанию
	company, err := c.CompanyService.GetCompanyByID(uint(companyID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена"})
		return
	}

	// Удаляем администратора из списка администраторов компании
	for i, admin := range company.Admins {
		if admin.ID == input.AdminID {
			company.Admins = append(company.Admins[:i], company.Admins[i+1:]...) // Удаление администратора
			break
		}
	}

	if err := c.CompanyService.RemoveAdmin(company, input.AdminID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отвязывании администратора от компании"})
		return
	}

	ctx.JSON(http.StatusOK, company)
}
