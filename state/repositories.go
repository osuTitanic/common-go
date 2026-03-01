package state

import (
	"github.com/osuTitanic/common-go/database"
	"github.com/osuTitanic/common-go/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	Users database.IUserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users: repositories.NewUserRepository(db),
	}
}
