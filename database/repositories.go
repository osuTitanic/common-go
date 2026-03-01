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

type IStatsRepository interface {
	Create(stats *schemas.Stats) error
	Save(stats *schemas.Stats) error
	Delete(stats *schemas.Stats) error

	ByMode(userId int, mode int) (*schemas.Stats, error)
	ManyByUserId(userId int) ([]*schemas.Stats, error)
}

type IRelationshipRepository interface {
	Create(relationship *schemas.Relationship) error
	Save(relationship *schemas.Relationship) error
	Delete(relationship *schemas.Relationship) error

	ByUserAndTarget(userId int, targetId int) (*schemas.Relationship, error)
	ManyByUserId(userId int) ([]*schemas.Relationship, error)
	ManyByTargetId(targetId int) ([]*schemas.Relationship, error)
	CountByUserId(userId int) (int, error)
	CountByTargetId(targetId int) (int, error)
}

type IBadgeRepository interface {
	Create(badge *schemas.Badge) error
	Save(badge *schemas.Badge) error
	Delete(badge *schemas.Badge) error

	ById(id int) (*schemas.Badge, error)
	ManyByUserId(userId int) ([]*schemas.Badge, error)
}

type INameRepository interface {
	Create(name *schemas.Name) error
	Save(name *schemas.Name) error
	Delete(name *schemas.Name) error

	ById(id int) (*schemas.Name, error)
	ByName(name string) (*schemas.Name, error)
	ManyByUserId(userId int) ([]*schemas.Name, error)
}

type IInfringementRepository interface {
	Create(infringement *schemas.Infringement) error
	Save(infringement *schemas.Infringement) error
	Delete(infringement *schemas.Infringement) error

	ById(id int) (*schemas.Infringement, error)
	ManyByUserId(userId int) ([]*schemas.Infringement, error)
}

type IReportRepository interface {
	Create(report *schemas.Report) error
	Save(report *schemas.Report) error
	Delete(report *schemas.Report) error

	ById(id int) (*schemas.Report, error)
	ManyByTargetId(targetId int) ([]*schemas.Report, error)
	ManyBySenderId(senderId int) ([]*schemas.Report, error)
}

type IVerificationRepository interface {
	Create(verification *schemas.Verification) error
	Save(verification *schemas.Verification) error
	Delete(verification *schemas.Verification) error

	ById(id int) (*schemas.Verification, error)
	ByToken(token string) (*schemas.Verification, error)
	ManyByUserId(userId int) ([]*schemas.Verification, error)
	DeleteByToken(token string) error
}

type IGroupRepository interface {
	Create(group *schemas.Group) error
	Save(group *schemas.Group) error
	Delete(group *schemas.Group) error

	ById(id int) (*schemas.Group, error)
	Many(includeHidden bool) ([]*schemas.Group, error)
}

type IGroupEntryRepository interface {
	Create(entry *schemas.GroupEntry) error
	Save(entry *schemas.GroupEntry) error
	Delete(entry *schemas.GroupEntry) error

	ByUserAndGroup(userId int, groupId int) (*schemas.GroupEntry, error)
	ManyByUserId(userId int) ([]*schemas.GroupEntry, error)
	ManyByGroupId(groupId int) ([]*schemas.GroupEntry, error)
}

type IUserPermissionRepository interface {
	Create(permission *schemas.UserPermission) error
	Save(permission *schemas.UserPermission) error
	Delete(permission *schemas.UserPermission) error

	ById(id int) (*schemas.UserPermission, error)
	ManyByUserId(userId int) ([]*schemas.UserPermission, error)
}

type IGroupPermissionRepository interface {
	Create(permission *schemas.GroupPermission) error
	Save(permission *schemas.GroupPermission) error
	Delete(permission *schemas.GroupPermission) error

	ById(id int) (*schemas.GroupPermission, error)
	ManyByGroupId(groupId int) ([]*schemas.GroupPermission, error)
}

type INotificationRepository interface {
	Create(notification *schemas.Notification) error
	Save(notification *schemas.Notification) error
	Delete(notification *schemas.Notification) error

	ById(id int64) (*schemas.Notification, error)
	ManyByUserId(userId int) ([]*schemas.Notification, error)
	CountByUserId(userId int) (int, error)
}

type IAchievementRepository interface {
	Create(achievement *schemas.Achievement) error
	Save(achievement *schemas.Achievement) error
	Delete(achievement *schemas.Achievement) error

	ManyByUserId(userId int) ([]*schemas.Achievement, error)
}

type IBeatmapFavouriteRepository interface {
	Create(favourite *schemas.BeatmapFavourite) error
	Save(favourite *schemas.BeatmapFavourite) error
	Delete(favourite *schemas.BeatmapFavourite) error

	ByUserAndSet(userId int, setId int) (*schemas.BeatmapFavourite, error)
	ManyByUserId(userId int) ([]*schemas.BeatmapFavourite, error)
	CountByUserId(userId int) (int, error)
	CountBySetId(setId int) (int, error)
}
