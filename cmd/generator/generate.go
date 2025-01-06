package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type CommandData struct {
	ModuleName    string
	ModuleTitle   string
	ModuleLower   string
	ModulePlural  string
	GeneratedBy   string
	GeneratedTime string
}

const (
	modulesDirPath = "internal/modules"
)

func main() {
	// Define commands
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Flags for init command
	initModuleName := initCmd.String("m", "", "Module name")
	withTests := initCmd.Bool("tests", false, "Generate test files")
	withMocks := initCmd.Bool("mocks", false, "Generate mock files for testing")
	force := initCmd.Bool("force", false, "Force overwrite if module exists")
	skipPrompt := initCmd.Bool("y", false, "Skip confirmation prompts")

	// Flags for delete command
	deleteModuleName := deleteCmd.String("m", "", "Module name to delete")
	skipDeletePrompt := deleteCmd.Bool("y", false, "Skip confirmation prompts")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		if *initModuleName == "" {
			fmt.Println("Please provide a module name using -m flag")
			os.Exit(1)
		}
		if err := validateModuleName(*initModuleName); err != nil {
			fmt.Printf("Invalid module name: %v\n", err)
			os.Exit(1)
		}
		generateModule(*initModuleName, *withTests, *withMocks, *force, *skipPrompt)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteModuleName == "" {
			fmt.Println("Please provide a module name to delete using -m flag")
			os.Exit(1)
		}
		deleteModule(*deleteModuleName, *skipDeletePrompt)

	case "help":
		printUsage()

	default:
		fmt.Println("Expected 'init', 'delete' or 'help' subcommand")
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  Generate a new module:")
	fmt.Println("    go run cmd/generator/generate.go init -m <module-name> [options]")
	fmt.Println("    Options:")
	fmt.Println("      -tests    Generate test files")
	fmt.Println("      -mocks    Generate mock files for testing")
	fmt.Println("      -force    Force overwrite if module exists")
	fmt.Println("      -y        Skip confirmation prompts")
	fmt.Println("\n  Delete a module:")
	fmt.Println("    go run cmd/generator/generate.go delete -m <module-name> [options]")
	fmt.Println("    Options:")
	fmt.Println("      -y        Skip confirmation prompts")
	fmt.Println("\n  Show this help:")
	fmt.Println("    go run cmd/generator/generate.go help")
}

func validateModuleName(name string) error {
	if len(name) < 2 {
		return fmt.Errorf("module name must be at least 2 characters long")
	}
	if !strings.HasPrefix(strings.ToLower(name), name) {
		return fmt.Errorf("module name must start with a lowercase letter")
	}
	return nil
}

func confirmAction(prompt string, skipPrompt bool) bool {
	if skipPrompt {
		return true
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", prompt)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

func deleteModule(moduleName string, skipPrompt bool) {
	modulePath := filepath.Join(modulesDirPath, strings.ToLower(moduleName))
	
	// Check if module exists
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		fmt.Printf("Module '%s' does not exist at path: %s\n", moduleName, modulePath)
		return
	}

	// Count files to be deleted
	var fileCount int
	filepath.Walk(modulePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileCount++
		}
		return nil
	})

	// Confirm deletion
	if !confirmAction(
		fmt.Sprintf("Are you sure you want to delete module '%s' containing %d files?", moduleName, fileCount),
		skipPrompt,
	) {
		fmt.Println("Operation cancelled")
		return
	}

	// Delete the module directory
	err := os.RemoveAll(modulePath)
	if err != nil {
		fmt.Printf("Error deleting module '%s': %v\n", moduleName, err)
		return
	}

	fmt.Printf("Successfully deleted module '%s' from %s\n", moduleName, modulePath)
}

func generateModule(name string, withTests, withMocks, force, skipPrompt bool) {
	modulePath := filepath.Join(modulesDirPath, strings.ToLower(name))

	// Check if module already exists
	if _, err := os.Stat(modulePath); err == nil && !force {
		if !confirmAction(
			fmt.Sprintf("Module '%s' already exists. Overwrite?", name),
			skipPrompt,
		) {
			fmt.Println("Operation cancelled")
			return
		}
	}

	data := CommandData{
		ModuleName:    name,
		ModuleTitle:   strings.Title(strings.ToLower(name)),
		ModuleLower:   strings.ToLower(name),
		ModulePlural:  strings.ToLower(name) + "s",
		GeneratedBy:   "go-backend generator",
		GeneratedTime: time.Now().Format(time.RFC3339),
	}

	// Create module directory structure
	dirs := []string{
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/entity"),
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/repository"),
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/service"),
		filepath.Join(modulesDirPath, data.ModuleLower, "handlers"),
		filepath.Join(modulesDirPath, data.ModuleLower, "dto"),
	}

	if withTests {
		dirs = append(dirs,
			filepath.Join(modulesDirPath, data.ModuleLower, "tests"),
		)
	}

	if withMocks {
		dirs = append(dirs,
			filepath.Join(modulesDirPath, data.ModuleLower, "mocks"),
		)
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Generate files
	templates := map[string]string{
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/entity", data.ModuleLower+".go"):          entityTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/repository", "repository.go"):             repositoryTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "domain/service", "service.go"):                   serviceTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "handlers", "handler.go"):                         handlerTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "dto", "dto.go"):                                 dtoTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "module.go"):                                      moduleTemplate,
		filepath.Join(modulesDirPath, data.ModuleLower, "routes.go"):                                      routesTemplate,
	}

	if withTests {
		templates[filepath.Join(modulesDirPath, data.ModuleLower, "tests", "handler_test.go")] = handlerTestTemplate
		templates[filepath.Join(modulesDirPath, data.ModuleLower, "tests", "service_test.go")] = serviceTestTemplate
	}

	if withMocks {
		templates[filepath.Join(modulesDirPath, data.ModuleLower, "mocks", "repository_mock.go")] = repositoryMockTemplate
		templates[filepath.Join(modulesDirPath, data.ModuleLower, "mocks", "service_mock.go")] = serviceMockTemplate
	}

	for path, tmpl := range templates {
		t, err := template.New(filepath.Base(path)).Parse(tmpl)
		if err != nil {
			fmt.Printf("Error parsing template %s: %v\n", path, err)
			return
		}

		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", path, err)
			return
		}
		defer file.Close()

		if err := t.Execute(file, data); err != nil {
			fmt.Printf("Error executing template %s: %v\n", path, err)
			return
		}

		fmt.Printf("Generated %s\n", path)
	}

	fmt.Printf("\nSuccessfully generated module '%s'\n", name)
	if withTests {
		fmt.Println("Test files generated")
	}
	if withMocks {
		fmt.Println("Mock files generated")
	}
}

var entityTemplate = `package entity

import (
	"time"
	"gorm.io/gorm"
)

type {{.ModuleTitle}} struct {
	ID          uint           ` + "`json:\"id\" gorm:\"primaryKey\"`" + `
	Name        string         ` + "`json:\"name\" gorm:\"not null\"`" + `
	Description string         ` + "`json:\"description\"`" + `
	UserID      uint          ` + "`json:\"user_id\" gorm:\"not null\"`" + `
	CreatedAt   time.Time      ` + "`json:\"created_at\"`" + `
	UpdatedAt   time.Time      ` + "`json:\"updated_at\"`" + `
	DeletedAt   gorm.DeletedAt ` + "`json:\"-\" gorm:\"index\"`" + `
}`

var repositoryTemplate = `package repository

import (
	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
	"gorm.io/gorm"
)

type {{.ModuleTitle}}Repository interface {
	Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error
	GetByID(id uint) (*entity.{{.ModuleTitle}}, error)
	GetAll() ([]entity.{{.ModuleTitle}}, error)
	Update({{.ModuleLower}} *entity.{{.ModuleTitle}}) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error)
}

type {{.ModuleLower}}Repository struct {
	db *gorm.DB
}

func New{{.ModuleTitle}}Repository(db *gorm.DB) {{.ModuleTitle}}Repository {
	return &{{.ModuleLower}}Repository{db: db}
}

func (r *{{.ModuleLower}}Repository) Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	return r.db.Create({{.ModuleLower}}).Error
}

func (r *{{.ModuleLower}}Repository) GetByID(id uint) (*entity.{{.ModuleTitle}}, error) {
	var {{.ModuleLower}} entity.{{.ModuleTitle}}
	err := r.db.First(&{{.ModuleLower}}, id).Error
	return &{{.ModuleLower}}, err
}

func (r *{{.ModuleLower}}Repository) GetAll() ([]entity.{{.ModuleTitle}}, error) {
	var {{.ModulePlural}} []entity.{{.ModuleTitle}}
	err := r.db.Find(&{{.ModulePlural}}).Error
	return {{.ModulePlural}}, err
}

func (r *{{.ModuleLower}}Repository) Update({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	return r.db.Save({{.ModuleLower}}).Error
}

func (r *{{.ModuleLower}}Repository) Delete(id uint) error {
	return r.db.Delete(&entity.{{.ModuleTitle}}{}, id).Error
}

func (r *{{.ModuleLower}}Repository) GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error) {
	var {{.ModulePlural}} []entity.{{.ModuleTitle}}
	err := r.db.Where("user_id = ?", userID).Find(&{{.ModulePlural}}).Error
	return {{.ModulePlural}}, err
}`

var serviceTemplate = `package service

import (
	"errors"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/repository"
)

type {{.ModuleTitle}}Service interface {
	Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error
	GetByID(id uint) (*entity.{{.ModuleTitle}}, error)
	GetAll() ([]entity.{{.ModuleTitle}}, error)
	Update({{.ModuleLower}} *entity.{{.ModuleTitle}}, userID uint) error
	Delete(id, userID uint) error
	GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error)
}

type {{.ModuleLower}}Service struct {
	repo repository.{{.ModuleTitle}}Repository
}

func New{{.ModuleTitle}}Service(repo repository.{{.ModuleTitle}}Repository) {{.ModuleTitle}}Service {
	return &{{.ModuleLower}}Service{repo: repo}
}

func (s *{{.ModuleLower}}Service) Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	return s.repo.Create({{.ModuleLower}})
}

func (s *{{.ModuleLower}}Service) GetByID(id uint) (*entity.{{.ModuleTitle}}, error) {
	return s.repo.GetByID(id)
}

func (s *{{.ModuleLower}}Service) GetAll() ([]entity.{{.ModuleTitle}}, error) {
	return s.repo.GetAll()
}

func (s *{{.ModuleLower}}Service) Update({{.ModuleLower}} *entity.{{.ModuleTitle}}, userID uint) error {
	existing, err := s.repo.GetByID({{.ModuleLower}}.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only update your own {{.ModulePlural}}")
	}

	return s.repo.Update({{.ModuleLower}})
}

func (s *{{.ModuleLower}}Service) Delete(id, userID uint) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only delete your own {{.ModulePlural}}")
	}

	return s.repo.Delete(id)
}

func (s *{{.ModuleLower}}Service) GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error) {
	return s.repo.GetByUserID(userID)
}`

var handlerTemplate = `package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/service"
	"go-backend/internal/modules/{{.ModuleLower}}/dto"
)

type Response struct {
	Status  int         ` + "`json:\"status\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   string      ` + "`json:\"error,omitempty\"`" + `
}

type {{.ModuleTitle}}Handler struct {
	service service.{{.ModuleTitle}}Service
}

func New{{.ModuleTitle}}Handler(service service.{{.ModuleTitle}}Service) *{{.ModuleTitle}}Handler {
	return &{{.ModuleTitle}}Handler{service: service}
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func (h *{{.ModuleTitle}}Handler) Create(c *gin.Context) {
	var req dto.Create{{.ModuleTitle}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	{{.ModuleLower}} := &entity.{{.ModuleTitle}}{
		Name:        req.Name,
		Description: req.Description,
		UserID:      userID.(uint),
	}

	if err := h.service.Create({{.ModuleLower}}); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to create {{.ModuleLower}}", nil, err.Error()))
		return
	}

	response := dto.{{.ModuleTitle}}Response{
		ID:          {{.ModuleLower}}.ID,
		Name:        {{.ModuleLower}}.Name,
		Description: {{.ModuleLower}}.Description,
		UserID:      {{.ModuleLower}}.UserID,
	}

	c.JSON(http.StatusCreated, formatResponse(http.StatusCreated, "{{.ModuleTitle}} created successfully", response, ""))
}

func (h *{{.ModuleTitle}}Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	{{.ModuleLower}}, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, formatResponse(http.StatusNotFound, "{{.ModuleTitle}} not found", nil, err.Error()))
		return
	}

	response := dto.{{.ModuleTitle}}Response{
		ID:          {{.ModuleLower}}.ID,
		Name:        {{.ModuleLower}}.Name,
		Description: {{.ModuleLower}}.Description,
		UserID:      {{.ModuleLower}}.UserID,
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "{{.ModuleTitle}} retrieved successfully", response, ""))
}

func (h *{{.ModuleTitle}}Handler) GetAll(c *gin.Context) {
	{{.ModulePlural}}, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve {{.ModulePlural}}", nil, err.Error()))
		return
	}

	response := make([]dto.{{.ModuleTitle}}Response, len({{.ModulePlural}}))
	for i, {{.ModuleLower}} := range {{.ModulePlural}} {
		response[i] = dto.{{.ModuleTitle}}Response{
			ID:          {{.ModuleLower}}.ID,
			Name:        {{.ModuleLower}}.Name,
			Description: {{.ModuleLower}}.Description,
			UserID:      {{.ModuleLower}}.UserID,
		}
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "{{.ModulePlural}} retrieved successfully", response, ""))
}

func (h *{{.ModuleTitle}}Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	var req dto.Update{{.ModuleTitle}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	{{.ModuleLower}} := &entity.{{.ModuleTitle}}{
		ID:          uint(id),
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.Update({{.ModuleLower}}, userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to update {{.ModuleLower}}", nil, err.Error()))
		return
	}

	response := dto.{{.ModuleTitle}}Response{
		ID:          {{.ModuleLower}}.ID,
		Name:        {{.ModuleLower}}.Name,
		Description: {{.ModuleLower}}.Description,
		UserID:      userID.(uint),
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "{{.ModuleTitle}} updated successfully", response, ""))
}

func (h *{{.ModuleTitle}}Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, formatResponse(http.StatusBadRequest, "Invalid ID", nil, "Invalid ID format"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	if err := h.service.Delete(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to delete {{.ModuleLower}}", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "{{.ModuleTitle}} deleted successfully", nil, ""))
}

func (h *{{.ModuleTitle}}Handler) GetByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, formatResponse(http.StatusUnauthorized, "User not authenticated", nil, "User not authenticated"))
		return
	}

	{{.ModulePlural}}, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, formatResponse(http.StatusInternalServerError, "Failed to retrieve user's {{.ModulePlural}}", nil, err.Error()))
		return
	}

	response := make([]dto.{{.ModuleTitle}}Response, len({{.ModulePlural}}))
	for i, {{.ModuleLower}} := range {{.ModulePlural}} {
		response[i] = dto.{{.ModuleTitle}}Response{
			ID:          {{.ModuleLower}}.ID,
			Name:        {{.ModuleLower}}.Name,
			Description: {{.ModuleLower}}.Description,
			UserID:      {{.ModuleLower}}.UserID,
		}
	}

	c.JSON(http.StatusOK, formatResponse(http.StatusOK, "User's {{.ModulePlural}} retrieved successfully", response, ""))
}`

var dtoTemplate = `package dto

type Create{{.ModuleTitle}}Request struct {
	Name        string ` + "`json:\"name\" binding:\"required\"`" + `
	Description string ` + "`json:\"description\"`" + `
}

type Update{{.ModuleTitle}}Request struct {
	Name        string ` + "`json:\"name\" binding:\"required\"`" + `
	Description string ` + "`json:\"description\"`" + `
}

type {{.ModuleTitle}}Response struct {
	ID          uint   ` + "`json:\"id\"`" + `
	Name        string ` + "`json:\"name\"`" + `
	Description string ` + "`json:\"description\"`" + `
	UserID      uint   ` + "`json:\"user_id\"`" + `
}`

var moduleTemplate = `package {{.ModuleLower}}

import (
	"go-backend/internal/modules/{{.ModuleLower}}/domain/repository"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/service"
	"go-backend/internal/modules/{{.ModuleLower}}/handlers"
	"gorm.io/gorm"
)

type Module struct {
	Handler *handlers.{{.ModuleTitle}}Handler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.New{{.ModuleTitle}}Repository(db)
	svc := service.New{{.ModuleTitle}}Service(repo)
	handler := handlers.New{{.ModuleTitle}}Handler(svc)

	return &Module{
		Handler: handler,
	}
}`

var routesTemplate = `package {{.ModuleLower}}

import (
	"github.com/gin-gonic/gin"
	"go-backend/internal/infrastructure/middleware"
)

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	{{.ModulePlural}} := router.Group("/{{.ModulePlural}}")
	{
		// Public routes
		{{.ModulePlural}}.GET("", m.Handler.GetAll)
		{{.ModulePlural}}.GET("/:id", m.Handler.GetByID)

		// Protected routes
		protected := {{.ModulePlural}}.Use(middleware.JWTAuth(middleware.AccessToken))
		{
			protected.POST("", m.Handler.Create)
			protected.PUT("/:id", m.Handler.Update)
			protected.DELETE("/:id", m.Handler.Delete)
		}
	}
}`

var handlerTestTemplate = `package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/service"
	"go-backend/internal/modules/{{.ModuleLower}}/dto"
	"go-backend/internal/modules/{{.ModuleLower}}/handlers"
	"go-backend/internal/modules/{{.ModuleLower}}/mocks"
)

func TestCreate{{.ModuleTitle}}(t *testing.T) {
	mockService := new(mocks.Mock{{.ModuleTitle}}Service)
	handler := handlers.New{{.ModuleTitle}}Handler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/{{.ModulePlural}}", handler.Create)

	tests := []struct {
		name           string
		input         dto.Create{{.ModuleTitle}}Request
		setupMock     func()
		expectedCode  int
		expectedError string
	}{
		{
			name: "Success",
			input: dto.Create{{.ModuleTitle}}Request{
				Name:        "Test {{.ModuleTitle}}",
				Description: "Test Description",
			},
			setupMock: func() {
				mockService.On("Create", mock.AnythingOfType("*entity.{{.ModuleTitle}}")).
					Return(nil)
			},
			expectedCode: http.StatusCreated,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/{{.ModulePlural}}", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			mockService.AssertExpectations(t)
		})
	}
}`

var serviceTestTemplate = `package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/service"
	"go-backend/internal/modules/{{.ModuleLower}}/mocks"
)

func TestCreate{{.ModuleTitle}}Service(t *testing.T) {
	mockRepo := new(mocks.Mock{{.ModuleTitle}}Repository)
	svc := service.New{{.ModuleTitle}}Service(mockRepo)

	tests := []struct {
		name          string
		input         *entity.{{.ModuleTitle}}
		setupMock     func()
		expectedError error
	}{
		{
			name: "Success",
			input: &entity.{{.ModuleTitle}}{
				Name:        "Test {{.ModuleTitle}}",
				Description: "Test Description",
			},
			setupMock: func() {
				mockRepo.On("Create", mock.AnythingOfType("*entity.{{.ModuleTitle}}")).
					Return(nil)
			},
			expectedError: nil,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			err := svc.Create(tt.input)
			assert.Equal(t, tt.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}`

var repositoryMockTemplate = `package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
)

type Mock{{.ModuleTitle}}Repository struct {
	mock.Mock
}

func (m *Mock{{.ModuleTitle}}Repository) Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	args := m.Called({{.ModuleLower}})
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Repository) GetByID(id uint) (*entity.{{.ModuleTitle}}, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.{{.ModuleTitle}}), args.Error(1)
}

func (m *Mock{{.ModuleTitle}}Repository) GetAll() ([]entity.{{.ModuleTitle}}, error) {
	args := m.Called()
	return args.Get(0).([]entity.{{.ModuleTitle}}), args.Error(1)
}

func (m *Mock{{.ModuleTitle}}Repository) Update({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	args := m.Called({{.ModuleLower}})
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Repository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Repository) GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.{{.ModuleTitle}}), args.Error(1)
}`

var serviceMockTemplate = `package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/internal/modules/{{.ModuleLower}}/domain/entity"
)

type Mock{{.ModuleTitle}}Service struct {
	mock.Mock
}

func (m *Mock{{.ModuleTitle}}Service) Create({{.ModuleLower}} *entity.{{.ModuleTitle}}) error {
	args := m.Called({{.ModuleLower}})
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Service) GetByID(id uint) (*entity.{{.ModuleTitle}}, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.{{.ModuleTitle}}), args.Error(1)
}

func (m *Mock{{.ModuleTitle}}Service) GetAll() ([]entity.{{.ModuleTitle}}, error) {
	args := m.Called()
	return args.Get(0).([]entity.{{.ModuleTitle}}), args.Error(1)
}

func (m *Mock{{.ModuleTitle}}Service) Update({{.ModuleLower}} *entity.{{.ModuleTitle}}, userID uint) error {
	args := m.Called({{.ModuleLower}}, userID)
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Service) Delete(id, userID uint) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

func (m *Mock{{.ModuleTitle}}Service) GetByUserID(userID uint) ([]entity.{{.ModuleTitle}}, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.{{.ModuleTitle}}), args.Error(1)
}`

var middlewareTemplate = `package middleware

import (
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-backend/internal/modules/user/domain/repository"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Response struct {
	Status  int         ` + "`json:\"status\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   string      ` + "`json:\"error,omitempty\"`" + `
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

// RateLimiter stores IP-based rate limiters
type RateLimiter struct {
	visitors map[string]*rate.Limiter
	mu      sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*rate.Limiter),
	}
}

// GetLimiter returns a rate limiter for an IP
func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.visitors[ip]
	if !exists {
		// Allow 100 requests per minute
		limiter = rate.NewLimiter(rate.Every(time.Minute/100), 100)
		rl.visitors[ip] = limiter
	}

	return limiter
}

// RateLimitMiddleware limits the number of requests from an IP
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rl.GetLimiter(ip)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, formatResponse(
				http.StatusTooManyRequests,
				"Rate limit exceeded",
				nil,
				"Too many requests. Please try again later.",
			))
			c.Abort()
			return
		}
		c.Next()
	}
}

// TokenType represents the type of token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// JWTAuth middleware for protecting routes
func JWTAuth(tokenType TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Authorization header is required",
			))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token format. Use 'Bearer <token>'",
			))
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		claims := jwt.MapClaims{}

		// Validate token signature and claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token signature",
			))
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token claims",
			))
			c.Abort()
			return
		}

		// Validate token type
		if tokenTypeClaim, ok := claims["token_type"].(string); !ok || TokenType(tokenTypeClaim) != tokenType {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token type",
			))
			c.Abort()
			return
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.JSON(http.StatusUnauthorized, formatResponse(
					http.StatusUnauthorized,
					"Authentication failed",
					nil,
					"Token has expired",
				))
				c.Abort()
				return
			}
		}

		// Get user from database and verify token
		userID := uint(claims["user_id"].(float64))
		db := c.MustGet("db").(*gorm.DB)
		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.GetByID(userID)

		if err != nil {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"User not found",
			))
			c.Abort()
			return
		}

		if user.Token == nil || *user.Token != tokenString {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Token has been revoked",
			))
			c.Abort()
			return
		}

		// Store user information in context
		c.Set("user_id", userID)
		c.Set("user_role", claims["role"])
		c.Next()
	}
}

// RequireAuth protects routes that require authentication
func RequireAuth(enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enabled {
			c.Next()
			return
		}

		_, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication required",
				nil,
				"You must be logged in to access this resource",
			))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRole checks if the authenticated user has the required role
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusForbidden, formatResponse(
				http.StatusForbidden,
				"Access denied",
				nil,
				"Role information not found",
			))
			c.Abort()
			return
		}

		if userRole != role {
			c.JSON(http.StatusForbidden, formatResponse(
				http.StatusForbidden,
				"Access denied",
				nil,
				"Insufficient permissions to access this resource",
			))
			c.Abort()
			return
		}

		c.Next()
	}
}`
