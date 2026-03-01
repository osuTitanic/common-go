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
