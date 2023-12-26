package getFollowersRepositories

import (
	"gorm.io/gorm"
)

type (
	GetFollowersRepositoriesService interface {
		// GetLastFollowers() ([]igtools.Followers, error)
	}

	getFollowersRepository struct {
		db *gorm.DB
	}
)

func NewGetFollowersRepository(db *gorm.DB) GetFollowersRepositoriesService {
	return &getFollowersRepository{db: db}
}

// func (r *getFollowersRepository) GetLastFollowers() ([]igtools.Followers, error) {
// }
