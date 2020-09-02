package proto

//procm:use=derive_binary
type Tag byte

const (
	ConnectRequest Tag = iota + 1
	Disconnect
	SetUserSlot
	PlayerInfo
	PlayerInventorySlot
	RequestWorldData
	WorldInfo
	RequestEssentialTiles
	Status
	SendSection
	SectionTileFrame
	SpawnPlayer
	UpdatePlayer
	PlayerActive
	Null
	PlayerHP
	ModifyTile
	Time
	DoorToggle
	SendTileSquare
	UpdateItemDrop
	UpdateItemOwner
	NPCUpdate
	StrikeNPCwithHeldItem
	Null2
	Null3
	ProjectileUpdate
	NPCStrike
	DestroyProjectile
	TogglePVP
	OpenChest
	UpdateChestItem
	SyncActiveChest
	PlaceChest
	HealEffect
	PlayerZone
	RequestPassword
	SendPassword
	RemoveItemOwner
	SetActiveNPC
	PlayerItemAnimation
	PlayerMana
	ManaEffect
	Null4
	PlayerTeam
	RequestSign
	UpdateSign
	SetLiquid
	CompleteConnectionAndSpawn
	UpdatePlayerBuff
	SpecialNPCEffect
	Unlock
	AddNPCBuff
	UpdateNPCBuff
	AddPlayerBuff
	UpdateNPCName
	UpdateGoodEvil
	PlayMusicItem
	HitSwitch
	NPCHomeUpdate
	SpawnBossInvasion
	PlayerDodge
	PaintTile
	PaintWall
	PlayerNPCTeleport
	HealOtherPlayer
	Placeholder
	ClientUUID
	GetChestName
	CatchNPC
	ReleaseNPC
	TravellingMerchantInventory
	TeleportationPotion
	AnglerQuest
	CompleteAnglerQuestToday
	NumberOfAnglerQuestsCompleted
	CreateTemporaryAnimation
	ReportInvasionProgress
	PlaceObject
	SyncPlayerChestIndex
	CreateCombatText
	LoadNetModule
	SetNPCKillCount
	SetPlayerStealth
	ForceItemIntoNearestChest
	UpdateTileEntity
	PlaceTileEntity
	TweakItem
	PlaceItemFrame
	UpdateItemDrop2
	SyncEmoteBubble
	SyncExtraValue
	SocialHandshake
	Deprecated
	KillPortal
	PlayerTeleportPortal
	NotifyPlayerNPCKilled
	NotifyPlayerOfEvent
	UpdateMinionTarget
	NPCTeleportPortal
	UpdateShieldStrengths
	NebulaLevelUp
	MoonLordCountdown
	NPCShopItem
	GemLockToggle
	PoofofSmoke
	SmartTextMessage
	WiredCannonShot
	MassWireOperation
	MassWireOperationConsume
	ToggleBirthdayParty
	GrowFX
	CrystalInvasionStart
	CrystalInvasionWipeAll
	MinionAttackTargetUpdate
	CrystalInvasionSendWaitTime
	PlayerHurtV2
	PlayerDeathV2
	CombatTextString
	Emoji
	TEDisplayDollItemSync
	RequestTileEntityInteraction
	WeaponsRackTryPlacing
	TEHatRackItemSync
	SyncTilePicking
	SyncRevengeMarker
	RemoveRevengeMarker
	LandGolfBallInCup
	FinishedConnectingToServer
	FishOutNPC
	TamperWithNPC
	PlayLegacySound
	FoodPlatterTryPlacing
	UpdatePlayerLuckFactors
	DeadPlayer
	SyncCavernMonsterType
	RequestNPCBuffRemoval
	ClientFinishedInventoryChangesOnThisTick
	SetCountsAsHostForGameplay
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
