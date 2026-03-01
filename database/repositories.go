package database

import "github.com/osuTitanic/common-go/schemas"

type IUserRepository interface {
	Create(user *schemas.User) error
	Save(user *schemas.User) error
	Delete(user *schemas.User) error

	ById(id int) (*schemas.User, error)
	ByName(name string) (*schemas.User, error)
	BySafeName(safeName string) (*schemas.User, error)
	ByEmail(email string) (*schemas.User, error)
	ByDiscordId(discordId int64) (*schemas.User, error)

	ManyById(userIds []int) ([]*schemas.User, error)
	ManyByName(names []string) ([]*schemas.User, error)
	ManyByRank(limit int, ascending bool) ([]*schemas.User, error)
	ManyByCreationDate(limit int, ascending bool) ([]*schemas.User, error)

	GetUsername(id int) (string, error)
	GetUserId(name string) (int, error)
	GetAvatarChecksum(id int) (string, error)
	GetCount() (int, error)
}
