package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Rafiur/wallet_app/internal/config"
	"github.com/Rafiur/wallet_app/internal/config/database/postgres"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/repo_postgres"
	"github.com/Rafiur/wallet_app/internal/infrastructure/repository/schema"
	"github.com/Rafiur/wallet_app/internal/security"
	"github.com/Rafiur/wallet_app/internal/service"
)

func main() {
	fmt.Println("Starting database seeder...")

	// Initialize configuration
	config.Init()
	dynamicConfig := config.GetDynamicConfig()
	if dynamicConfig == nil {
		log.Fatal("Failed to load dynamic config")
	}

	// Initialize database
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories and services
	userRepo := repo_postgres.NewUserRepo(db)
	passwordService := security.NewPasswordService()
	userService := service.NewUserService(userRepo, passwordService)

	ctx := context.Background()

	// Seed default admin user
	adminUser := &schema.User{
		FullName: "Super Admin",
		Email:    "admin@walletapp.com",
		Password: "admin123", // This will be hashed
	}

	// Check if admin already exists
	existingUser, err := userService.GetByEmail(ctx, adminUser.Email)
	if err == nil && existingUser != nil {
		fmt.Printf("Admin user already exists: %s\n", existingUser.Email)
		return
	}

	// Create admin user
	createdUser, err := userService.Create(ctx, adminUser)
	if err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Printf("Admin user created successfully!\n")
	fmt.Printf("Email: %s\n", createdUser.Email)
	fmt.Printf("Password: admin123\n")
	fmt.Println("Please change the password after first login.")
}
