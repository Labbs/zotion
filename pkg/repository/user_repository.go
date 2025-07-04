package repository

import (
	"github.com/labbs/zotion/pkg/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository and expose the functions available to the user
// repository. It takes a gorm.DB instance as a parameter to interact with the database.
// The userRepository struct implements the UserRepository interface defined in the models package.
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// GetByEmailOrUsername retrieves a user from the database by their email or username.
// It takes an emailOrUsername string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetByEmailOrUsername(emailOrUsername string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR username = ?", emailOrUsername, emailOrUsername).First(&user).Error
	return user, err
}

// GetByEmail returns a user from the database by their email.
// It takes an email string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
// The email is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
// The email is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetById retrieves a user from the database by their ID.
// It takes an id string as a parameter and returns a user and an error.
// If the user is found, it returns the user and a nil error. If not, it returns an empty user and an error.
// The error is nil if the user is found, otherwise it contains the error message.
// The ID is the unique identifier for the user in the database.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is found, otherwise it contains the error message.
func (r *userRepository) GetById(id string) (models.User, error) {
	var user models.User
	err := r.db.Debug().Select("id, name, email, avatar_url, created_at, updated_at").Where("id = ?", id).First(&user).Error
	return user, err
}

// Create creates a new user in the database.
// It takes a user pointer as a parameter and returns an error.
// If the user is created successfully, it returns a nil error. If not, it returns an error.
// The user is passed as a pointer to avoid copying the entire struct.
// The user is returned as a models.User struct.
// The error is returned as a gorm.Error type, which contains information about the error that occurred.
// The error is nil if the user is created successfully, otherwise it contains the error message.
// The user is created in the database using the GORM Create method.
// The Create method takes a pointer to the user struct and creates a new record in the database.
func (r *userRepository) Create(user *models.User) error {
	err := r.db.Debug().Create(user).Error
	return err
}

// Update updates a user
// It takes a user pointer as a parameter and returns the updated user and an error.
func (r *userRepository) Update(user *models.User) (models.User, error) {
	return *user, r.db.Debug().Save(user).Error
}

// Delete deletes a user
// It takes an id string as a parameter and returns an error.
func (r *userRepository) Delete(id string) error {
	return r.db.Debug().Where("id = ?", id).Delete(&models.User{}).Error
}

// GetGroups returns the groups of a user
// It retrieves the groups associated with a user by their ID.
func (r *userRepository) GetGroupsByUserId(id string) ([]models.Group, error) {
	var groups []models.Group
	err := r.db.Debug().Unscoped().Table("group").Select("id", "role").
		Joins("JOIN user_group ON user_group.group_id = id").
		Where("user_group.user_id = ?", id).
		Find(&groups).Error
	return groups, err
}

// GetAllUsers returns all users
// It retrieves all active users from the database, selecting only the id, name, email, and avatar_url fields.
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Debug().Select("id, name, email, avatar_url").Where("active = true").Find(&users).Error
	return users, err
}

// GetAllInactiveUsers returns all inactive users
// It retrieves all inactive users from the database, selecting only the id, name, email, and avatar_url fields.
func (r *userRepository) GetAllInactiveUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Debug().Select("id, name, email, avatar_url").Where("active = false").Find(&users).Error
	return users, err
}

// GetPreferencesById retrieves the preferences of a user by their ID.
func (r *userRepository) GetPreferencesById(id string) (models.JSONB, error) {
	var user models.User
	err := r.db.Debug().Select("preferences").Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.JSONB{}, err
	}
	return user.Preferences, nil
}

// UpdatePreferences updates the preferences of a user by their ID.
func (r *userRepository) UpdatePreferences(id string, preferences models.JSONB) error {
	return r.db.Debug().Model(&models.User{}).Where("id = ?", id).Update("preferences", preferences).Error
}

// GetUserWithGroups retrieves a user by their ID, including their groups.
func (r *userRepository) GetUserWithGroups(id string) (models.User, error) {
	var user models.User
	err := r.db.Debug().Preload("Groups").Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUsersWithGroups retrieves all users with their associated groups.
func (r *userRepository) GetUsersWithGroups() ([]models.User, error) {
	var users []models.User
	err := r.db.Debug().Preload("Groups").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
