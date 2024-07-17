package repositories

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    var users []models.User
    if err := r.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
    var user models.User
    if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
    user.ID = uuid.New()
    return r.DB.Create(user).Error
}

func (r *UserRepository) Update(id string, user *models.User) error {
    return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) Delete(id string) error {
    return r.DB.Delete(&models.User{}, "id = ?", id).Error
}

func (r *UserRepository) UsernameExists(username string) bool {
    var user models.User
    err := r.DB.Where("username = ?", username).First(&user).Error
    return err == nil
}

func (r *UserRepository) EmailExists(email string) bool {
    var user models.User
    err := r.DB.Where("email = ?", email).First(&user).Error
    return err == nil
}
