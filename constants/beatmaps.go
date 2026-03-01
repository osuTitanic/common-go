package constants

type BeatmapStatus int

const (
	BeatmapStatusInactive  BeatmapStatus = -3
	BeatmapStatusGraveyard BeatmapStatus = -2
	BeatmapStatusWIP       BeatmapStatus = -1
	BeatmapStatusPending   BeatmapStatus = 0
	BeatmapStatusRanked    BeatmapStatus = 1
	BeatmapStatusApproved  BeatmapStatus = 2
	BeatmapStatusQualified BeatmapStatus = 3
	BeatmapStatusLoved     BeatmapStatus = 4
)

type BeatmapServer int

const (
	BeatmapServerBancho  BeatmapServer = 0
	BeatmapServerTitanic BeatmapServer = 1
)

type BeatmapResourceType int

const (
	BeatmapResourceTypeOsz        BeatmapResourceType = 0
	BeatmapResourceTypeOszNoVideo BeatmapResourceType = 1
	BeatmapResourceTypeBeatmap    BeatmapResourceType = 2
	BeatmapResourceTypeThumbnail  BeatmapResourceType = 3
	BeatmapResourceTypeBackground BeatmapResourceType = 4
	BeatmapResourceTypeAudio      BeatmapResourceType = 5
)

type BeatmapGenre int

const (
	BeatmapGenreAny         BeatmapGenre = 0
	BeatmapGenreUnspecified BeatmapGenre = 1
	BeatmapGenreVideoGame   BeatmapGenre = 2
	BeatmapGenreAnime       BeatmapGenre = 3
	BeatmapGenreRock        BeatmapGenre = 4
	BeatmapGenrePop         BeatmapGenre = 5
	BeatmapGenreOther       BeatmapGenre = 6
	BeatmapGenreNovelty     BeatmapGenre = 7
	BeatmapGenreHipHop      BeatmapGenre = 9
	BeatmapGenreElectronic  BeatmapGenre = 10
	BeatmapGenreMetal       BeatmapGenre = 11
	BeatmapGenreClassical   BeatmapGenre = 12
	BeatmapGenreFolk        BeatmapGenre = 13
	BeatmapGenreJazz        BeatmapGenre = 14
)

type BeatmapLanguage int

const (
	BeatmapLanguageAny          BeatmapLanguage = 0
	BeatmapLanguageUnspecified  BeatmapLanguage = 1
	BeatmapLanguageEnglish      BeatmapLanguage = 2
	BeatmapLanguageJapanese     BeatmapLanguage = 3
	BeatmapLanguageChinese      BeatmapLanguage = 4
	BeatmapLanguageInstrumental BeatmapLanguage = 5
	BeatmapLanguageKorean       BeatmapLanguage = 6
	BeatmapLanguageFrench       BeatmapLanguage = 7
	BeatmapLanguageGerman       BeatmapLanguage = 8
	BeatmapLanguageSwedish      BeatmapLanguage = 9
	BeatmapLanguageSpanish      BeatmapLanguage = 10
	BeatmapLanguageItalian      BeatmapLanguage = 11
	BeatmapLanguageRussian      BeatmapLanguage = 12
	BeatmapLanguagePolish       BeatmapLanguage = 13
	BeatmapLanguageOther        BeatmapLanguage = 14
)
