// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sandwich

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SandwichClient is the client API for Sandwich service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandwichClient interface {
	// FetchConsumerConfiguration returns the Consumer Configuration.
	FetchConsumerConfiguration(ctx context.Context, in *FetchConsumerConfigurationRequest, opts ...grpc.CallOption) (*FetchConsumerConfigurationResponse, error)
	// FetchGuildChannels returns guilds based on the guildID.
	// Takes either query or channelIDs. Empty query and empty channelIDs will return all.
	FetchGuildChannels(ctx context.Context, in *FetchGuildChannelsRequest, opts ...grpc.CallOption) (*ChannelsResponse, error)
	// FetchGuildEmojis returns emojis based on the guildID.
	// Takes either query or emojiIDs. Empty query and empty emojiIDs will return all.
	FetchGuildEmojis(ctx context.Context, in *FetchGuildEmojisRequest, opts ...grpc.CallOption) (*EmojisResponse, error)
	// FetchGuildMembers returns guild members based on the guildID.
	// Takes either query or userIDs. Empty query and empty userIDs will return all.
	FetchGuildMembers(ctx context.Context, in *FetchGuildMembersRequest, opts ...grpc.CallOption) (*GuildMembersResponse, error)
	// FetchGuild returns guilds based on the guildIDs.
	FetchGuild(ctx context.Context, in *FetchGuildRequest, opts ...grpc.CallOption) (*GuildResponse, error)
	// FetchGuildRoles returns roles based on the roleIDs.
	// Takes either query or roleIDs. Empty query and empty roleIDs will return all.
	FetchGuildRoles(ctx context.Context, in *FetchGuildRolesRequest, opts ...grpc.CallOption) (*GuildRolesResponse, error)
	// FetchMutualGuilds returns a list of all mutual guilds based on userID.
	// Populates guildIDs with a list of snowflakes of all guilds.
	// If expand is passed and True, will also populate guilds with the guild object.
	FetchMutualGuilds(ctx context.Context, in *FetchMutualGuildsRequest, opts ...grpc.CallOption) (*GuildsResponse, error)
	// RequestGuildChunk sends a guild chunk request.
	// Returns once the guild has been chunked.
	RequestGuildChunk(ctx context.Context, in *RequestGuildChunkRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	// SendWebsocketMessage manually sends a websocket message.
	SendWebsocketMessage(ctx context.Context, in *SendWebsocketMessageRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	// WhereIsGuild returns a list of WhereIsGuildLocations based on guildId.
	// WhereIsGuildLocations contains the manager, shardGroup and shardId.
	WhereIsGuild(ctx context.Context, in *WhereIsGuildRequest, opts ...grpc.CallOption) (*WhereIsGuildResponse, error)
}

type sandwichClient struct {
	cc grpc.ClientConnInterface
}

func NewSandwichClient(cc grpc.ClientConnInterface) SandwichClient {
	return &sandwichClient{cc}
}

func (c *sandwichClient) FetchConsumerConfiguration(ctx context.Context, in *FetchConsumerConfigurationRequest, opts ...grpc.CallOption) (*FetchConsumerConfigurationResponse, error) {
	out := new(FetchConsumerConfigurationResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchConsumerConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchGuildChannels(ctx context.Context, in *FetchGuildChannelsRequest, opts ...grpc.CallOption) (*ChannelsResponse, error) {
	out := new(ChannelsResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchGuildChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchGuildEmojis(ctx context.Context, in *FetchGuildEmojisRequest, opts ...grpc.CallOption) (*EmojisResponse, error) {
	out := new(EmojisResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchGuildEmojis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchGuildMembers(ctx context.Context, in *FetchGuildMembersRequest, opts ...grpc.CallOption) (*GuildMembersResponse, error) {
	out := new(GuildMembersResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchGuildMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchGuild(ctx context.Context, in *FetchGuildRequest, opts ...grpc.CallOption) (*GuildResponse, error) {
	out := new(GuildResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchGuildRoles(ctx context.Context, in *FetchGuildRolesRequest, opts ...grpc.CallOption) (*GuildRolesResponse, error) {
	out := new(GuildRolesResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchGuildRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) FetchMutualGuilds(ctx context.Context, in *FetchMutualGuildsRequest, opts ...grpc.CallOption) (*GuildsResponse, error) {
	out := new(GuildsResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/FetchMutualGuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) RequestGuildChunk(ctx context.Context, in *RequestGuildChunkRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/RequestGuildChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) SendWebsocketMessage(ctx context.Context, in *SendWebsocketMessageRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/SendWebsocketMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandwichClient) WhereIsGuild(ctx context.Context, in *WhereIsGuildRequest, opts ...grpc.CallOption) (*WhereIsGuildResponse, error) {
	out := new(WhereIsGuildResponse)
	err := c.cc.Invoke(ctx, "/sandwich.Sandwich/WhereIsGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandwichServer is the server API for Sandwich service.
// All implementations must embed UnimplementedSandwichServer
// for forward compatibility
type SandwichServer interface {
	// FetchConsumerConfiguration returns the Consumer Configuration.
	FetchConsumerConfiguration(context.Context, *FetchConsumerConfigurationRequest) (*FetchConsumerConfigurationResponse, error)
	// FetchGuildChannels returns guilds based on the guildID.
	// Takes either query or channelIDs. Empty query and empty channelIDs will return all.
	FetchGuildChannels(context.Context, *FetchGuildChannelsRequest) (*ChannelsResponse, error)
	// FetchGuildEmojis returns emojis based on the guildID.
	// Takes either query or emojiIDs. Empty query and empty emojiIDs will return all.
	FetchGuildEmojis(context.Context, *FetchGuildEmojisRequest) (*EmojisResponse, error)
	// FetchGuildMembers returns guild members based on the guildID.
	// Takes either query or userIDs. Empty query and empty userIDs will return all.
	FetchGuildMembers(context.Context, *FetchGuildMembersRequest) (*GuildMembersResponse, error)
	// FetchGuild returns guilds based on the guildIDs.
	FetchGuild(context.Context, *FetchGuildRequest) (*GuildResponse, error)
	// FetchGuildRoles returns roles based on the roleIDs.
	// Takes either query or roleIDs. Empty query and empty roleIDs will return all.
	FetchGuildRoles(context.Context, *FetchGuildRolesRequest) (*GuildRolesResponse, error)
	// FetchMutualGuilds returns a list of all mutual guilds based on userID.
	// Populates guildIDs with a list of snowflakes of all guilds.
	// If expand is passed and True, will also populate guilds with the guild object.
	FetchMutualGuilds(context.Context, *FetchMutualGuildsRequest) (*GuildsResponse, error)
	// RequestGuildChunk sends a guild chunk request.
	// Returns once the guild has been chunked.
	RequestGuildChunk(context.Context, *RequestGuildChunkRequest) (*BaseResponse, error)
	// SendWebsocketMessage manually sends a websocket message.
	SendWebsocketMessage(context.Context, *SendWebsocketMessageRequest) (*BaseResponse, error)
	// WhereIsGuild returns a list of WhereIsGuildLocations based on guildId.
	// WhereIsGuildLocations contains the manager, shardGroup and shardId.
	WhereIsGuild(context.Context, *WhereIsGuildRequest) (*WhereIsGuildResponse, error)
	mustEmbedUnimplementedSandwichServer()
}

// UnimplementedSandwichServer must be embedded to have forward compatible implementations.
type UnimplementedSandwichServer struct {
}

func (UnimplementedSandwichServer) FetchConsumerConfiguration(context.Context, *FetchConsumerConfigurationRequest) (*FetchConsumerConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchConsumerConfiguration not implemented")
}
func (UnimplementedSandwichServer) FetchGuildChannels(context.Context, *FetchGuildChannelsRequest) (*ChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuildChannels not implemented")
}
func (UnimplementedSandwichServer) FetchGuildEmojis(context.Context, *FetchGuildEmojisRequest) (*EmojisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuildEmojis not implemented")
}
func (UnimplementedSandwichServer) FetchGuildMembers(context.Context, *FetchGuildMembersRequest) (*GuildMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuildMembers not implemented")
}
func (UnimplementedSandwichServer) FetchGuild(context.Context, *FetchGuildRequest) (*GuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuild not implemented")
}
func (UnimplementedSandwichServer) FetchGuildRoles(context.Context, *FetchGuildRolesRequest) (*GuildRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuildRoles not implemented")
}
func (UnimplementedSandwichServer) FetchMutualGuilds(context.Context, *FetchMutualGuildsRequest) (*GuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMutualGuilds not implemented")
}
func (UnimplementedSandwichServer) RequestGuildChunk(context.Context, *RequestGuildChunkRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestGuildChunk not implemented")
}
func (UnimplementedSandwichServer) SendWebsocketMessage(context.Context, *SendWebsocketMessageRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWebsocketMessage not implemented")
}
func (UnimplementedSandwichServer) WhereIsGuild(context.Context, *WhereIsGuildRequest) (*WhereIsGuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhereIsGuild not implemented")
}
func (UnimplementedSandwichServer) mustEmbedUnimplementedSandwichServer() {}

// UnsafeSandwichServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandwichServer will
// result in compilation errors.
type UnsafeSandwichServer interface {
	mustEmbedUnimplementedSandwichServer()
}

func RegisterSandwichServer(s grpc.ServiceRegistrar, srv SandwichServer) {
	s.RegisterService(&Sandwich_ServiceDesc, srv)
}

func _Sandwich_FetchConsumerConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchConsumerConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchConsumerConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchConsumerConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchConsumerConfiguration(ctx, req.(*FetchConsumerConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchGuildChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGuildChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchGuildChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchGuildChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchGuildChannels(ctx, req.(*FetchGuildChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchGuildEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGuildEmojisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchGuildEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchGuildEmojis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchGuildEmojis(ctx, req.(*FetchGuildEmojisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchGuildMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGuildMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchGuildMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchGuildMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchGuildMembers(ctx, req.(*FetchGuildMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchGuild(ctx, req.(*FetchGuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchGuildRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGuildRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchGuildRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchGuildRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchGuildRoles(ctx, req.(*FetchGuildRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_FetchMutualGuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchMutualGuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).FetchMutualGuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/FetchMutualGuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).FetchMutualGuilds(ctx, req.(*FetchMutualGuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_RequestGuildChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGuildChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).RequestGuildChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/RequestGuildChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).RequestGuildChunk(ctx, req.(*RequestGuildChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_SendWebsocketMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendWebsocketMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).SendWebsocketMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/SendWebsocketMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).SendWebsocketMessage(ctx, req.(*SendWebsocketMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandwich_WhereIsGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhereIsGuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandwichServer).WhereIsGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwich.Sandwich/WhereIsGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandwichServer).WhereIsGuild(ctx, req.(*WhereIsGuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sandwich_ServiceDesc is the grpc.ServiceDesc for Sandwich service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sandwich_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sandwich.Sandwich",
	HandlerType: (*SandwichServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchConsumerConfiguration",
			Handler:    _Sandwich_FetchConsumerConfiguration_Handler,
		},
		{
			MethodName: "FetchGuildChannels",
			Handler:    _Sandwich_FetchGuildChannels_Handler,
		},
		{
			MethodName: "FetchGuildEmojis",
			Handler:    _Sandwich_FetchGuildEmojis_Handler,
		},
		{
			MethodName: "FetchGuildMembers",
			Handler:    _Sandwich_FetchGuildMembers_Handler,
		},
		{
			MethodName: "FetchGuild",
			Handler:    _Sandwich_FetchGuild_Handler,
		},
		{
			MethodName: "FetchGuildRoles",
			Handler:    _Sandwich_FetchGuildRoles_Handler,
		},
		{
			MethodName: "FetchMutualGuilds",
			Handler:    _Sandwich_FetchMutualGuilds_Handler,
		},
		{
			MethodName: "RequestGuildChunk",
			Handler:    _Sandwich_RequestGuildChunk_Handler,
		},
		{
			MethodName: "SendWebsocketMessage",
			Handler:    _Sandwich_SendWebsocketMessage_Handler,
		},
		{
			MethodName: "WhereIsGuild",
			Handler:    _Sandwich_WhereIsGuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}
