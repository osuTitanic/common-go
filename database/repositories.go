package database

type IUserRepository interface {
	Create(user *User) error
	Save(user *User) error
	Delete(user *User) error

	ById(id int) (*User, error)
	ByName(name string) (*User, error)
	BySafeName(safeName string) (*User, error)
	ByEmail(email string) (*User, error)
	ByDiscordId(discordId string) (*User, error)

	ManyById(userIds []int) ([]*User, error)
	ManyByName(names []string) ([]*User, error)
	ManyByRank(limit int, ascending bool) ([]*User, error)
	ManyByCreationDate(limit int, ascending bool) ([]*User, error)

	GetUsername(id int) (string, error)
	GetUserId(name string) (int, error)
	GetAvatarChecksum(id int) (string, error)
	GetCount() (int, error)
}
