package structs

import (
	discord "github.com/WelcomerTeam/Sandwich-Daemon/next/discord/structs"
)

// SandwichMetadata represents the identification information that consumers will use.
type SandwichMetadata struct {
	Version    string
	Identifier string
	// ShardGroup ID, Shard ID, Shard Count
	Shard [3]int
}

// SandwichPayload represents the data that is sent to consumers.
type SandwichPayload struct {
	discord.GatewayPayload

	Data interface{}

	Extra    map[string]interface{}
	Metadata SandwichMetadata
	Trace    map[string]int
}
