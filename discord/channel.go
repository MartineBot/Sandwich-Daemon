package discord

// channel.go contains the information relating to channels

// ChannelType represents a channel's type.
type ChannelType uint8

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
	_
	_
	_
	ChannelTypeGuildNewsThread
	ChannelTypeGuildPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
)

// VideoQualityMode represents the quality of the video.
type VideoQualityMode uint8

const (
	VideoQualityModeAuto VideoQualityMode = 1 + iota
	VideoqualityModeFull
)

// StageChannelPrivacyLevel represents the privacy level of a stage channel.
type StageChannelPrivacyLevel uint8

const (
	StageChannelPrivacyLevelPublic StageChannelPrivacyLevel = 1 + iota
	StageChannelPrivacyLevelGuildOnly
)

// Channel represents a Discord channel.
type Channel struct {
	OwnerID                    *Snowflake           `json:"owner_id,omitempty"`
	GuildID                    *Snowflake           `json:"guild_id,omitempty"`
	Permissions                Int64                `json:"permissions"`
	ThreadMember               *ThreadMember        `json:"member,omitempty"`
	ThreadMetadata             *ThreadMetadata      `json:"thread_metadata,omitempty"`
	VideoQualityMode           VideoQualityMode     `json:"video_quality_mode"`
	LastPinTimestamp           *Timestamp           `json:"last_pin_timestamp"`
	ParentID                   *Snowflake           `json:"parent_id,omitempty"`
	ApplicationID              *Snowflake           `json:"application_id,omitempty"`
	RTCRegion                  string               `json:"rtc_region"`
	Topic                      string               `json:"topic"`
	Icon                       string               `json:"icon"`
	Name                       string               `json:"name"`
	LastMessageID              *string              `json:"last_message_id"`
	PermissionOverwrites       ChannelOverwriteList `json:"permission_overwrites"`
	Recipients                 UserList             `json:"recipients"`
	ID                         Snowflake            `json:"id"`
	UserLimit                  int32                `json:"user_limit"`
	Bitrate                    int32                `json:"bitrate"`
	MessageCount               int32                `json:"message_count"`
	MemberCount                int32                `json:"member_count"`
	RateLimitPerUser           int32                `json:"rate_limit_per_user"`
	Position                   int32                `json:"position"`
	DefaultAutoArchiveDuration int32                `json:"default_auto_archive_duration"`
	NSFW                       bool                 `json:"nsfw"`
	Type                       ChannelType          `json:"type"`
}

// ChannelOverwrite represents a permission overwrite for a channel.
type ChannelOverwrite struct {
	Type  ChannelOverrideType `json:"type"`
	ID    Snowflake           `json:"id"`
	Allow Int64               `json:"allow"`
	Deny  Int64               `json:"deny"`
}

// ChannelOverrideType represents the target of a channel override.
type ChannelOverrideType uint8

const (
	ChannelOverrideTypeRole ChannelOverrideType = iota
	ChannelOverrideTypeMember
)

// ThreadMetadata contains thread-specific channel fields.
type ThreadMetadata struct {
	ArchiveTimestamp    Timestamp `json:"archive_timestamp"`
	AutoArchiveDuration int32     `json:"auto_archive_duration"`
	Archived            bool      `json:"archived"`
	Locked              bool      `json:"locked"`
}

// ThreadMember is used to indicate whether a user has joined a thread or not.
type ThreadMember struct {
	ID            *Snowflake `json:"id,omitempty"`
	UserID        *Snowflake `json:"user_id,omitempty"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	JoinTimestamp Timestamp  `json:"join_timestamp"`
	Flags         int32      `json:"flags"`
}

// StageInstance represents a stage channel instance.
type StageInstance struct {
	PrivacyLabel         *StageChannelPrivacyLevel `json:"privacy_level"`
	Topic                string                    `json:"topic"`
	ID                   Snowflake                 `json:"id"`
	GuildID              Snowflake                 `json:"guild_id"`
	ChannelID            Snowflake                 `json:"channel_id"`
	DiscoverableDisabled bool                      `json:"discoverable_disabled"`
}

// FollowedChannel represents a followed channel.
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// ChannelPermissionsParams represents the arguments to modify guild channel permissions.
type ChannelPermissionsParams struct {
	ID              Snowflake `json:"id"`
	Position        int32     `json:"position"`
	LockPermissions bool      `json:"lock_permissions"`
	ParentID        Snowflake `json:"parent_id"`
}
