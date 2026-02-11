package schemas

import (
	"time"

	"github.com/osuTitanic/common-go/constants"
)

type User struct {
	Id               int                   `gorm:"column:id;primaryKey;autoIncrement"`
	Name             string                `gorm:"column:name;unique"`
	SafeName         string                `gorm:"column:safe_name;unique"`
	Email            string                `gorm:"column:email;unique"`
	DiscordId        *int64                `gorm:"column:discord_id;unique"`
	Bcrypt           string                `gorm:"column:pw"`
	IsBot            bool                  `gorm:"column:bot;default:false"`
	Country          string                `gorm:"column:country"`
	SilenceEnd       *time.Time            `gorm:"column:silence_end"`
	CreatedAt        time.Time             `gorm:"column:created_at;autoCreateTime"`
	LatestActivity   time.Time             `gorm:"column:latest_activity;autoCreateTime"`
	Restricted       bool                  `gorm:"column:restricted;default:false"`
	Activated        bool                  `gorm:"column:activated;default:false"`
	PreferredMode    constants.Mode        `gorm:"column:preferred_mode;default:0"`
	PreferredRanking constants.RankingType `gorm:"column:preferred_ranking;default:global"`
	Playstyle        int                   `gorm:"column:playstyle;default:0"`
	IrcToken         string                `gorm:"column:irc_token;default:encode(gen_random_bytes(5),'hex')"`
	AvatarHash       *string               `gorm:"column:avatar_hash"`
	AvatarLastUpdate time.Time             `gorm:"column:avatar_last_changed;autoCreateTime"`
	FriendOnlyDMs    bool                  `gorm:"column:friendonly_dms;default:false"`
	// TODO: Add playstyle to constants

	Userpage  *string `gorm:"column:userpage_about"`
	Signature *string `gorm:"column:userpage_signature"`
	Title     *string `gorm:"column:userpage_title"`
	Banner    *string `gorm:"column:userpage_banner"`
	Website   *string `gorm:"column:userpage_website"`
	Discord   *string `gorm:"column:userpage_discord"`
	Twitter   *string `gorm:"column:userpage_twitter"`
	Location  *string `gorm:"column:userpage_location"`
	Interests *string `gorm:"column:userpage_interests"`
}

func (User) TableName() string {
	return "users"
}
