package proto

//procm:use=derive_binary
type Tag byte

const (
	ConnectRequest                           Tag = 1
	Disconnect                                   = 2
	SetUserSlot                                  = 3
	PlayerInfo                                   = 4
	PlayerInventorySlot                          = 5
	RequestWorldData                             = 6
	WorldInfo                                    = 7
	RequestEssentialTiles                        = 8
	Status                                       = 9
	SendSection                                  = 10
	SectionTileFrame                             = 11
	SpawnPlayer                                  = 12
	UpdatePlayer                                 = 13
	PlayerActive                                 = 14
	Null                                         = 15
	PlayerHP                                     = 16
	ModifyTile                                   = 17
	Time                                         = 18
	DoorToggle                                   = 19
	SendTileSquare                               = 20
	UpdateItemDrop                               = 21
	UpdateItemOwner                              = 22
	NPCUpdate                                    = 23
	StrikeNPCwithHeldItem                        = 24
	Null2                                        = 25
	Null3                                        = 26
	ProjectileUpdate                             = 27
	NPCStrike                                    = 28
	DestroyProjectile                            = 29
	TogglePVP                                    = 30
	OpenChest                                    = 31
	UpdateChestItem                              = 32
	SyncActiveChest                              = 33
	PlaceChest                                   = 34
	HealEffect                                   = 35
	PlayerZone                                   = 36
	RequestPassword                              = 37
	SendPassword                                 = 38
	RemoveItemOwner                              = 39
	SetActiveNPC                                 = 40
	PlayerItemAnimation                          = 41
	PlayerMana                                   = 42
	ManaEffect                                   = 43
	Null4                                        = 44
	PlayerTeam                                   = 45
	RequestSign                                  = 46
	UpdateSign                                   = 47
	SetLiquid                                    = 48
	CompleteConnectionAndSpawn                   = 49
	UpdatePlayerBuff                             = 50
	SpecialNPCEffect                             = 51
	Unlock                                       = 52
	AddNPCBuff                                   = 53
	UpdateNPCBuff                                = 54
	AddPlayerBuff                                = 55
	UpdateNPCName                                = 56
	UpdateGoodEvil                               = 57
	PlayMusicItem                                = 58
	HitSwitch                                    = 59
	NPCHomeUpdate                                = 60
	SpawnBossInvasion                            = 61
	PlayerDodge                                  = 62
	PaintTile                                    = 63
	PaintWall                                    = 64
	PlayerNPCTeleport                            = 65
	HealOtherPlayer                              = 66
	Placeholder                                  = 67
	ClientUUID                                   = 68
	GetChestName                                 = 69
	CatchNPC                                     = 70
	ReleaseNPC                                   = 71
	TravellingMerchantInventory                  = 72
	TeleportationPotion                          = 73
	AnglerQuest                                  = 74
	CompleteAnglerQuestToday                     = 75
	NumberOfAnglerQuestsCompleted                = 76
	CreateTemporaryAnimation                     = 77
	ReportInvasionProgress                       = 78
	PlaceObject                                  = 79
	SyncPlayerChestIndex                         = 80
	CreateCombatText                             = 81
	LoadNetModule                                = 82
	SetNPCKillCount                              = 83
	SetPlayerStealth                             = 84
	ForceItemIntoNearestChest                    = 85
	UpdateTileEntity                             = 86
	PlaceTileEntity                              = 87
	TweakItem                                    = 88
	PlaceItemFrame                               = 89
	UpdateItemDrop2                              = 90
	SyncEmoteBubble                              = 91
	SyncExtraValue                               = 92
	SocialHandshake                              = 93
	Deprecated                                   = 94
	KillPortal                                   = 95
	PlayerTeleportPortal                         = 96
	NotifyPlayerNPCKilled                        = 97
	NotifyPlayerOfEvent                          = 98
	UpdateMinionTarget                           = 99
	NPCTeleportPortal                            = 100
	UpdateShieldStrengths                        = 101
	NebulaLevelUp                                = 102
	MoonLordCountdown                            = 103
	NPCShopItem                                  = 104
	GemLockToggle                                = 105
	PoofofSmoke                                  = 106
	SmartTextMessage                             = 107
	WiredCannonShot                              = 108
	MassWireOperation                            = 109
	MassWireOperationConsume                     = 110
	ToggleBirthdayParty                          = 111
	GrowFX                                       = 112
	CrystalInvasionStart                         = 113
	CrystalInvasionWipeAll                       = 114
	MinionAttackTargetUpdate                     = 115
	CrystalInvasionSendWaitTime                  = 116
	PlayerHurtV2                                 = 117
	PlayerDeathV2                                = 118
	CombatTextString                             = 119
	Emoji                                        = 120
	TEDisplayDollItemSync                        = 121
	RequestTileEntityInteraction                 = 122
	WeaponsRackTryPlacing                        = 123
	TEHatRackItemSync                            = 124
	SyncTilePicking                              = 125
	SyncRevengeMarker                            = 126
	RemoveRevengeMarker                          = 127
	LandGolfBallInCup                            = 128
	FinishedConnectingToServer                   = 129
	FishOutNPC                                   = 130
	TamperWithNPC                                = 131
	PlayLegacySound                              = 132
	FoodPlatterTryPlacing                        = 133
	UpdatePlayerLuckFactors                      = 134
	DeadPlayer                                   = 135
	SyncCavernMonsterType                        = 136
	RequestNPCBuffRemoval                        = 137
	ClientFinishedInventoryChangesOnThisTick     = 138
	SetCountsAsHostForGameplay                   = 139
)

func (t Tag) String() string {
	switch t {
	case ConnectRequest:
		return "ConnectRequest"
	case Disconnect:
		return "Disconnect"
	case SetUserSlot:
		return "SetUserSlot"
	case PlayerInfo:
		return "PlayerInfo"
	case PlayerInventorySlot:
		return "PlayerInventorySlot"
	case RequestWorldData:
		return "RequestWorldData"
	case WorldInfo:
		return "WorldInfo"
	case RequestEssentialTiles:
		return "RequestEssentialTiles"
	case Status:
		return "Status"
	case SendSection:
		return "SendSection"
	case SectionTileFrame:
		return "SectionTileFrame"
	case SpawnPlayer:
		return "SpawnPlayer"
	case UpdatePlayer:
		return "UpdatePlayer"
	case PlayerActive:
		return "PlayerActive"
	case Null:
		return "Null"
	case PlayerHP:
		return "PlayerHP"
	case ModifyTile:
		return "ModifyTile"
	case Time:
		return "Time"
	case DoorToggle:
		return "DoorToggle"
	case SendTileSquare:
		return "SendTileSquare"
	case UpdateItemDrop:
		return "UpdateItemDrop"
	case UpdateItemOwner:
		return "UpdateItemOwner"
	case NPCUpdate:
		return "NPCUpdate"
	case StrikeNPCwithHeldItem:
		return "StrikeNPCwithHeldItem"
	case Null2:
		return "Null2"
	case Null3:
		return "Null3"
	case ProjectileUpdate:
		return "ProjectileUpdate"
	case NPCStrike:
		return "NPCStrike"
	case DestroyProjectile:
		return "DestroyProjectile"
	case TogglePVP:
		return "TogglePVP"
	case OpenChest:
		return "OpenChest"
	case UpdateChestItem:
		return "UpdateChestItem"
	case SyncActiveChest:
		return "SyncActiveChest"
	case PlaceChest:
		return "PlaceChest"
	case HealEffect:
		return "HealEffect"
	case PlayerZone:
		return "PlayerZone"
	case RequestPassword:
		return "RequestPassword"
	case SendPassword:
		return "SendPassword"
	case RemoveItemOwner:
		return "RemoveItemOwner"
	case SetActiveNPC:
		return "SetActiveNPC"
	case PlayerItemAnimation:
		return "PlayerItemAnimation"
	case PlayerMana:
		return "PlayerMana"
	case ManaEffect:
		return "ManaEffect"
	case Null4:
		return "Null4"
	case PlayerTeam:
		return "PlayerTeam"
	case RequestSign:
		return "RequestSign"
	case UpdateSign:
		return "UpdateSign"
	case SetLiquid:
		return "SetLiquid"
	case CompleteConnectionAndSpawn:
		return "CompleteConnectionAndSpawn"
	case UpdatePlayerBuff:
		return "UpdatePlayerBuff"
	case SpecialNPCEffect:
		return "SpecialNPCEffect"
	case Unlock:
		return "Unlock"
	case AddNPCBuff:
		return "AddNPCBuff"
	case UpdateNPCBuff:
		return "UpdateNPCBuff"
	case AddPlayerBuff:
		return "AddPlayerBuff"
	case UpdateNPCName:
		return "UpdateNPCName"
	case UpdateGoodEvil:
		return "UpdateGoodEvil"
	case PlayMusicItem:
		return "PlayMusicItem"
	case HitSwitch:
		return "HitSwitch"
	case NPCHomeUpdate:
		return "NPCHomeUpdate"
	case SpawnBossInvasion:
		return "SpawnBossInvasion"
	case PlayerDodge:
		return "PlayerDodge"
	case PaintTile:
		return "PaintTile"
	case PaintWall:
		return "PaintWall"
	case PlayerNPCTeleport:
		return "PlayerNPCTeleport"
	case HealOtherPlayer:
		return "HealOtherPlayer"
	case Placeholder:
		return "Placeholder"
	case ClientUUID:
		return "ClientUUID"
	case GetChestName:
		return "GetChestName"
	case CatchNPC:
		return "CatchNPC"
	case ReleaseNPC:
		return "ReleaseNPC"
	case TravellingMerchantInventory:
		return "TravellingMerchantInventory"
	case TeleportationPotion:
		return "TeleportationPotion"
	case AnglerQuest:
		return "AnglerQuest"
	case CompleteAnglerQuestToday:
		return "CompleteAnglerQuestToday"
	case NumberOfAnglerQuestsCompleted:
		return "NumberOfAnglerQuestsCompleted"
	case CreateTemporaryAnimation:
		return "CreateTemporaryAnimation"
	case ReportInvasionProgress:
		return "ReportInvasionProgress"
	case PlaceObject:
		return "PlaceObject"
	case SyncPlayerChestIndex:
		return "SyncPlayerChestIndex"
	case CreateCombatText:
		return "CreateCombatText"
	case LoadNetModule:
		return "LoadNetModule"
	case SetNPCKillCount:
		return "SetNPCKillCount"
	case SetPlayerStealth:
		return "SetPlayerStealth"
	case ForceItemIntoNearestChest:
		return "ForceItemIntoNearestChest"
	case UpdateTileEntity:
		return "UpdateTileEntity"
	case PlaceTileEntity:
		return "PlaceTileEntity"
	case TweakItem:
		return "TweakItem"
	case PlaceItemFrame:
		return "PlaceItemFrame"
	case UpdateItemDrop2:
		return "UpdateItemDrop2"
	case SyncEmoteBubble:
		return "SyncEmoteBubble"
	case SyncExtraValue:
		return "SyncExtraValue"
	case SocialHandshake:
		return "SocialHandshake"
	case Deprecated:
		return "Deprecated"
	case KillPortal:
		return "KillPortal"
	case PlayerTeleportPortal:
		return "PlayerTeleportPortal"
	case NotifyPlayerNPCKilled:
		return "NotifyPlayerNPCKilled"
	case NotifyPlayerOfEvent:
		return "NotifyPlayerOfEvent"
	case UpdateMinionTarget:
		return "UpdateMinionTarget"
	case NPCTeleportPortal:
		return "NPCTeleportPortal"
	case UpdateShieldStrengths:
		return "UpdateShieldStrengths"
	case NebulaLevelUp:
		return "NebulaLevelUp"
	case MoonLordCountdown:
		return "MoonLordCountdown"
	case NPCShopItem:
		return "NPCShopItem"
	case GemLockToggle:
		return "GemLockToggle"
	case PoofofSmoke:
		return "PoofofSmoke"
	case SmartTextMessage:
		return "SmartTextMessage"
	case WiredCannonShot:
		return "WiredCannonShot"
	case MassWireOperation:
		return "MassWireOperation"
	case MassWireOperationConsume:
		return "MassWireOperationConsume"
	case ToggleBirthdayParty:
		return "ToggleBirthdayParty"
	case GrowFX:
		return "GrowFX"
	case CrystalInvasionStart:
		return "CrystalInvasionStart"
	case CrystalInvasionWipeAll:
		return "CrystalInvasionWipeAll"
	case MinionAttackTargetUpdate:
		return "MinionAttackTargetUpdate"
	case CrystalInvasionSendWaitTime:
		return "CrystalInvasionSendWaitTime"
	case PlayerHurtV2:
		return "PlayerHurtV2"
	case PlayerDeathV2:
		return "PlayerDeathV2"
	case CombatTextString:
		return "CombatTextString"
	case Emoji:
		return "Emoji"
	case TEDisplayDollItemSync:
		return "TEDisplayDollItemSync"
	case RequestTileEntityInteraction:
		return "RequestTileEntityInteraction"
	case WeaponsRackTryPlacing:
		return "WeaponsRackTryPlacing"
	case TEHatRackItemSync:
		return "TEHatRackItemSync"
	case SyncTilePicking:
		return "SyncTilePicking"
	case SyncRevengeMarker:
		return "SyncRevengeMarker"
	case RemoveRevengeMarker:
		return "RemoveRevengeMarker"
	case LandGolfBallInCup:
		return "LandGolfBallInCup"
	case FinishedConnectingToServer:
		return "FinishedConnectingToServer"
	case FishOutNPC:
		return "FishOutNPC"
	case TamperWithNPC:
		return "TamperWithNPC"
	case PlayLegacySound:
		return "PlayLegacySound"
	case FoodPlatterTryPlacing:
		return "FoodPlatterTryPlacing"
	case UpdatePlayerLuckFactors:
		return "UpdatePlayerLuckFactors"
	case DeadPlayer:
		return "DeadPlayer"
	case SyncCavernMonsterType:
		return "SyncCavernMonsterType"
	case RequestNPCBuffRemoval:
		return "RequestNPCBuffRemoval"
	case ClientFinishedInventoryChangesOnThisTick:
		return "ClientFinishedInventoryChangesOnThisTick"
	case SetCountsAsHostForGameplay:
		return "SetCountsAsHostForGameplay"
	}

	return "Unknown"
}
