package constants

type RankingType string

const (
	RankingTypeGlobal      RankingType = "global"
	RankingTypeCountry     RankingType = "country"
	RankingTypeTotalScore  RankingType = "tscore"
	RankingTypeRankedScore RankingType = "rscore"
	RankingTypePPv1        RankingType = "ppv1"
	RankingTypeClears      RankingType = "clears"
)
