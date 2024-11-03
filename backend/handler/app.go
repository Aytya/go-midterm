package handler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"myproject/backend/domain"
	"myproject/backend/repository"
	"time"
)

// App struct
type App struct {
	ctx  context.Context
	repo *repository.Repository
}

// NewApp creates a new App application struct
func NewApp(repo *repository.Repository) *App {
	return &App{
		repo: repo,
	}
}

var mySigningKey = []byte("secret")

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (a *App) Register(username, password string) (string, error) {
	if username == "" {
		return "", fmt.Errorf("username is required")
	}
	if password == "" {
		return "", fmt.Errorf("password is required")
	}

	userID, err := GenerateRandomID(16)
	if err != nil {
		return "", fmt.Errorf("failed to generate user ID: %v", err)
	}

	existingUser, err := a.repo.User.GetByUsername(context.Background(), username)
	if err == nil && existingUser != nil {
		return "", fmt.Errorf("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	user := domain.User{
		UserID:   userID,
		Username: username,
		Role:     "user",
		Password: string(passwordHash),
	}

	if err := a.repo.User.Create(context.Background(), &user); err != nil {
		return "", fmt.Errorf("could not register user: %v", err)
	}

	return "User registered successfully", nil
}

func (a *App) Login(username, password string) (map[string]string, error) {
	user, err := a.repo.User.GetByUsername(context.Background(), username)
	if err != nil || user == nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	token, err := a.GenerateJWT(user.UserID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("Error generating token")
	}

	return map[string]string{
		"token": token,
		"role":  user.Role,
	}, nil
}

func (a *App) GenerateJWT(userId, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func (a *App) Auth(tokenString string, requiredRole string) (bool, error) {
	claims, err := a.ValidateJWT(tokenString)
	if err != nil {
		return false, err
	}

	userRole := claims["role"].(string)
	if userRole != requiredRole {
		return false, fmt.Errorf("forbidden: insufficient permissions")
	}

	return true, nil
}

func (a *App) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func (a *App) CreateTodo(token string, title, priority string, dateTime time.Time) (*domain.Todo, error) {
	ctx, err := a.AuthMiddleware(token)
	if err != nil {
		return nil, err
	}

	userRole := ctx.Value("role").(string)
	if !CheckPermission(userRole, "create") {
		return nil, errors.New("permission denied")
	}

	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	id := uuid.New().String()
	activeAt := time.Now().UTC()

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		DateTime: dateTime,
		ActiveAt: activeAt,
		Status:   false,
		Priority: priority,
	}

	err = a.repo.Todo.Create(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) UpdateTodo(id, title, priority string, dateTime time.Time) (*domain.Todo, error) {
	if title == "" {
		log.Println("Error: Title is empty")
		return nil, errors.New("title cannot be empty")
	}
	if priority == "" {
		log.Println("Error: Date or time is empty")
		return nil, errors.New("date and time cannot be empty")
	}

	todo := &domain.Todo{
		ID:       id,
		Title:    title,
		DateTime: dateTime,
		Priority: priority,
	}

	err := a.repo.Todo.Update(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (a *App) GetTodoByID(id string) (*domain.Todo, error) {
	return a.repo.Todo.GetByID(context.Background(), id)
}

func (a *App) GetAllTodos() ([]*domain.Todo, error) {
	ctx := context.Background()
	todos, err := a.repo.Todo.GetAll(ctx)
	if err != nil {
		log.Printf("Error retrieving todos: %v", err)
		return nil, err
	}
	if todos == nil {
		log.Println("No todos found, returning empty slice.")
		return []*domain.Todo{}, nil
	}
	log.Println("Retrieved todos:", todos)
	return todos, nil
}

func (a *App) DeleteTodo(id string) error {
	return a.repo.Todo.Delete(context.Background(), id)
}

func (a *App) CheckTodo(id string) error {
	return a.repo.Todo.CheckTodo(context.Background(), id)
}
