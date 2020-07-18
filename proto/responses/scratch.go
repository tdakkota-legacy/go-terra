package responses

import (
	"github.com/tdakkota/go-terra/proto/structs"
	"github.com/tdakkota/go-terra/proto/structs/tile"
)

// Client -> Server
type ConnectRequest struct {
	Version string
}

// Server -> Client
type Disconnect struct {
	Reason tile.NetworkText
}

// Server -> Client
type SetUserSlot struct {
	PlayerID byte
}

// Server <-> Client (Sync)
type PlayerInfo struct {
	PlayerID        byte
	SkinVarient     byte
	Hair            byte
	Name            string
	HairDye         byte
	HideVisuals     byte
	HideVisuals2    byte
	HideMisc        byte
	HairColor       structs.Color
	SkinColor       structs.Color
	EyeColor        structs.Color
	ShirtColor      structs.Color
	UnderShirtColor structs.Color
	PantsColor      structs.Color
	ShoeColor       structs.Color
	DifficultyFlags byte
	TorchFlags      byte
}

// Server <-> Client (Sync)
type PlayerInventorySlot struct {
	PlayerID  byte
	SlotID    int16
	Stack     int16
	Prefix    byte
	ItemNetID int16
}

// Client -> Server
type RequestWorldData struct {
}

// Server -> Client
type WorldInfo struct {
	Time                        int32
	DayandMoonInfo              byte
	MoonPhase                   byte
	MaxTilesX                   int16
	MaxTilesY                   int16
	SpawnX                      int16
	SpawnY                      int16
	WorldSurface                int16
	RockLayer                   int16
	WorldID                     int32
	WorldName                   string
	GameMode                    byte
	WorldUniqueID               []byte
	WorldGeneratorVersion       uint64
	MoonType                    byte
	TreeBackground              byte
	CorruptionBackground        byte
	JungleBackground            byte
	SnowBackground              byte
	HallowBackground            byte
	CrimsonBackground           byte
	DesertBackground            byte
	OceanBackground             byte
	UnknownBackground1          byte
	UnknownBackground2          byte
	UnknownBackground3          byte
	UnknownBackground4          byte
	UnknownBackground5          byte
	IceBackStyle                byte
	JungleBackStyle             byte
	HellBackStyle               byte
	WindSpeedSet                float32
	CloudNumber                 byte
	Tree1                       int32
	Tree2                       int32
	Tree3                       int32
	TreeStyle1                  byte
	TreeStyle2                  byte
	TreeStyle3                  byte
	TreeStyle4                  byte
	CaveBack1                   int32
	CaveBack2                   int32
	CaveBack3                   int32
	CaveBackStyle1              byte
	CaveBackStyle2              byte
	CaveBackStyle3              byte
	CaveBackStyle4              byte
	Forest1TreeTopStyle         byte
	Forest2TreeTopStyle         byte
	Forest3TreeTopStyle         byte
	Forest4TreeTopStyle         byte
	CorruptionTreeTopStyle      byte
	JungleTreeTopStyle          byte
	SnowTreeTopStyle            byte
	HallowTreeTopStyle          byte
	CrimsonTreeTopStyle         byte
	DesertTreeTopStyle          byte
	OceanTreeTopStyle           byte
	GlowingMushroomTreeTopStyle byte
	UnderworldTreeTopStyle      byte
	Rain                        float32
	EventInfo                   byte
	EventInfo2                  byte
	EventInfo3                  byte
	EventInfo4                  byte
	EventInfo5                  byte
	EventInfo6                  byte
	EventInfo7                  byte
	CopperOreTier               int16
	IronOreTier                 int16
	SilverOreTier               int16
	GoldOreTier                 int16
	CobaltOreTier               int16
	MythrilOreTier              int16
	AdamantiteOreTier           int16
	InvasionType                int8
	LobbyID                     uint64
	SandstormSeverity           float32
}

// Client -> Server
type RequestEssentialTiles struct {
	X int32
	Y int32
}

// Server -> Client
type Status struct {
	StatusMax       int32
	StatusText      tile.NetworkText
	StatusTextFlags byte
}

// Server -> Client
type SendSection struct {
	Compressed      bool
	XStart          int32
	YStart          int32
	Width           int16
	Height          int16
	Tiles           []structs.Tile
	ChestCount      int16
	Chests          []structs.Chest
	SignCount       int16
	Signs           []structs.Sign
	TileEntityCount int16
	TileEntities    []structs.TileEntity
}

// Server -> Client
type SectionTileFrame struct {
	StartX int16
	StartY int16
	EndX   int16
	EndY   int16
}

// Client -> Server
type SpawnPlayer struct {
	PlayerID             byte
	SpawnX               int16
	SpawnY               int16
	RespawnTimeRemaining int32
	PlayerSpawnContext   byte
}

// Server <-> Client (Sync)
type UpdatePlayer struct {
	PlayerID          byte
	Control           byte
	Pulley            byte
	Misc              byte
	SleepingInfo      byte
	SelectedItem      byte
	PositionX         float32
	PositionY         float32
	VelocityX         float32
	VelocityY         float32
	OriginalPositionX float32
	OriginalPositionY float32
	HomePositionX     float32
	HomePositionY     float32
}

// Server -> Client
type PlayerActive struct {
	PlayerID byte
	Active   byte
}

// Never sent
type Null struct {
}

// Server <-> Client (Sync)
type PlayerHP struct {
	PlayerID byte
	HP       int16
	MaxHP    int16
}

// Server <-> Client (Sync)
type ModifyTile struct {
	Action byte
	TileX  int16
	TileY  int16
	Flags1 int16
	Flags2 byte
}

// Server -> Client
type Time struct {
	DayTime   bool
	TimeValue int32
	SunModY   int16
	MoonModY  int16
}

// Server <-> Client (Sync)
type DoorToggle struct {
	Action    byte
	TileX     int16
	TileY     int16
	Direction byte
}

// Server <-> Client (Sync)
type SendTileSquare struct {
	Size           uint16
	TileChangeType byte
	TileX          int16
	TileY          int16
	Tiles          []structs.Tile
}

// Server <-> Client (Sync)
type UpdateItemOwner struct {
	ItemID   int16
	PlayerID byte
}

// Server -> Client
type NPCUpdate struct {
	NPCID                                       int16
	PositionX                                   float32
	PositionY                                   float32
	VelocityX                                   float32
	VelocityY                                   float32
	Target                                      uint16
	NpcFlags1                                   byte
	NpcFlags2                                   byte
	AI                                          []float32
	NPCNetID                                    int16
	playerCountForMultiplayerDifficultyOverride byte
	StrengthMultiplier                          float32
	LifeBytes                                   byte
	Life                                        int32
	ReleaseOwner                                byte
}

// Server <-> Client (Sync)
type StrikeNPCwithHeldItem struct {
	NPCID    int16
	PlayerID byte
}

// Server <-> Client (Sync)
type ProjectileUpdate struct {
	ProjectileID   int16
	PositionX      float32
	PositionY      float32
	VelocityX      float32
	VelocityY      float32
	Owner          byte
	Type           int16
	ProjFlags      byte
	AI0            float32
	AI1            float32
	Damage         int16
	Knockback      float32
	OriginalDamage int16
	ProjUUID       int16
}

// Server <-> Client (Sync)
type NPCStrike struct {
	NPCID        int16
	Damage       int16
	Knockback    float32
	HitDirection byte
	Crit         bool
}

// Server <-> Client (Sync)
type DestroyProjectile struct {
	ProjectileID int16
	Owner        byte
}

// Server <-> Client (Sync)
type TogglePVP struct {
	PlayerID   byte
	PVPEnabled bool
}

// Client -> Server
type OpenChest struct {
	TileX int16
	TileY int16
}

// Server <-> Client (Sync)
type UpdateChestItem struct {
	ChestID   int16
	ItemSlot  byte
	Stack     int16
	Prefix    byte
	ItemNetID int16
}

// Server <-> Client (Sync)
type SyncActiveChest struct {
	ChestID    int16
	ChestX     int16
	ChestY     int16
	NameLength byte
	ChestName  string
}

// Server <-> Client
type PlaceChest struct {
	Action           byte
	TileX            int16
	TileY            int16
	Style            int16
	ChestIDtodestroy int16
}

// Server <-> Client (Sync)
type HealEffect struct {
	PlayerID   byte
	HealAmount int16
}

// Server <-> Client (Sync)
type PlayerZone struct {
	PlayerID   byte
	ZoneFlags1 byte
	ZoneFlags2 byte
	ZoneFlags3 byte
	ZoneFlags4 byte
}

// Server -> Client
type RequestPassword struct {
}

// Client -> Server
type SendPassword struct {
	Password string
}

// Server -> Client
type RemoveItemOwner struct {
	ItemIndex int16
}

// Server <-> Client (Sync)
type SetActiveNPC struct {
	PlayerID      byte
	NPCTalkTarget int16
}

// Server <-> Client (Sync)
type PlayerItemAnimation struct {
	PlayerID      byte
	ItemRotation  float32
	ItemAnimation int16
}

// Server <-> Client (Sync)
type PlayerMana struct {
	PlayerID byte
	Mana     int16
	MaxMana  int16
}

// Server <-> Client (Sync)
type ManaEffect struct {
	PlayerID   byte
	ManaAmount int16
}

// Server <-> Client (Sync)
type PlayerTeam struct {
	PlayerID byte
	Team     byte
}

// Client -> Server
type RequestSign struct {
	X int16
	Y int16
}

// Updates sign if sent from client otherwise displays sign to client.
type UpdateSign struct {
	SignID    int16
	X         int16
	Y         int16
	Text      string
	PlayerID  byte
	SignFlags byte
}

// Server <-> Client (Sync)
type SetLiquid struct {
	X          int16
	Y          int16
	Liquid     byte
	LiquidType byte
}

// Server -> Client
type CompleteConnectionAndSpawn struct {
	//	- -
}

// Server <-> Client (Sync)
type UpdatePlayerBuff struct {
	PlayerID byte
	BuffType [22]uint16
}

// Server <-> Client (Sync)
type SpecialNPCEffect struct {
	PlayerID byte
	Type     byte
}

// Server <-> Client (Sync)
type Unlock struct {
	Type byte
	X    int16
	Y    int16
}

// Server <-> Client (Sync)
type AddNPCBuff struct {
	NPCID int16
	Buff  uint16
	Time  int16
}

// Server -> Client
type UpdateNPCBuff struct {
	NPCID   int16
	BuffID1 uint16
	Time1   int16
	BuffID2 uint16
	Time2   int16
	BuffID3 uint16
	Time3   int16
	BuffID4 uint16
	Time4   int16
	BuffID5 uint16
	Time5   int16
}

// Server <-> Client (Sync)
type AddPlayerBuff struct {
	PlayerID byte
	Buff     uint16
	Time     int32
}

// Server <-> Client (Sync)
type UpdateNPCName struct {
	NPCID                 int16
	Name                  string
	TownNpcVariationIndex int32
}

// Server -> Client
type UpdateGoodEvil struct {
	Good    byte
	Evil    byte
	Crimson byte
}

// Server <-> Client (Sync)
type PlayMusicItem struct {
	PlayerID byte
	Note     float32
}

// Server <-> Client (Sync)
type HitSwitch struct {
	X int16
	Y int16
}

// Server <-> Client (Sync)
type NPCHomeUpdate struct {
	NPCID     int16
	HomeTileX int16
	HomeTileY int16
	Homeless  byte
}

// Client -> Server
type SpawnBossInvasion struct {
	PlayerID int16
	Type     int16
}

// Server <-> Client (Sync)
type PlayerDodge struct {
	PlayerID byte
	Flag     byte
}

// Server <-> Client (Sync)
type PaintTile struct {
	X     int16
	Y     int16
	Color byte
}

// Server <-> Client (Sync)
type PaintWall struct {
	X     int16
	Y     int16
	Color byte
}

// Server <-> Client (Sync)
type PlayerNPCTeleport struct {
	Flags     byte
	TargetID  int16
	X         float32
	Y         float32
	Style     byte
	ExtraInfo int32
}

// Server <-> Client (Sync)
type HealOtherPlayer struct {
	PlayerID   byte
	HealAmount int16
}

// Does not exist in the official client. Exists solely for the purpose of being used by custom clients and servers.
type Placeholder struct{}

// Client -> Server
type ClientUUID struct {
	UUID string
}

// Server <-> Client (Sync)
type GetChestName struct {
	ChestID int16
	ChestX  int16
	ChestY  int16
	Name    string
}

// Client -> Server
type CatchNPC struct {
	NPCID    int16
	PlayerID byte
}

// Client -> Server
type ReleaseNPC struct {
	X       int32
	Y       int32
	NPCType int16
	Style   byte
}

// Server -> Client
type TravellingMerchantInventory struct {
	Items [40]int16
}

// Server <-> Client
type TeleportationPotion struct {
	Type byte
}

// Server -> Client
type AnglerQuest struct {
	Quest     byte
	Completed bool
}

// Client -> Server
type CompleteAnglerQuestToday struct {
}

// Client -> Server
type NumberOfAnglerQuestsCompleted struct {
	PlayerID              byte
	AnglerQuestsCompleted int32
	GolferScore           int32
}

// Server -> Client
type CreateTemporaryAnimation struct {
	AnimationType int16
	TileType      uint16
	X             int16
	Y             int16
}

// Server -> Client
type ReportInvasionProgress struct {
	Progress    int32
	MaxProgress int32
	Icon        int8
	Wave        int8
}

// Server <-> Client
type PlaceObject struct {
	X         int16
	Y         int16
	Type      int16
	Style     int16
	Alternate byte
	Random    int8
	Direction bool
}

// Server -> Client
type SyncPlayerChestIndex struct {
	Player byte
	Chest  int16
}

// Server -> Client
type CreateCombatText struct {
	X          float32
	Y          float32
	Color      structs.Color
	HealAmount int32
}

// Client -> Server
type LoadNetModule struct {
	ModuleID uint16
	//	Payload ??
}

// Server -> Client
type SetNPCKillCount struct {
	NPCType   int16
	KillCount int32
}

// Server <-> Client (Sync)
type SetPlayerStealth struct {
	Player  byte
	Stealth float32
}

// Client -> Server
type ForceItemIntoNearestChest struct {
	InventorySlot byte
}

// Server -> Client
type UpdateTileEntity struct {
	TileEntityId   int32
	UpdateTileFlag bool
	TileEntityType byte
	X              int16
	Y              int16
}

// Client -> Server
type PlaceTileEntity struct {
	X              int16
	Y              int16
	TileEntityType byte
}

// Server -> Client
type TweakItem struct {
	ItemIndex        int16
	Flags1           byte
	PackedColorValue uint32
	Damage           uint16
	Knockback        float32
	UseAnimation     uint16
	UseTime          uint16
	Shoot            int16
	ShootSpeed       float32
	Flags2           byte
	Width            int16
	Height           int16
	Scale            float32
	Ammo             int16
	UseAmmo          int16
	NotAmmo          bool
}

// Client -> Server
type PlaceItemFrame struct {
	X      int16
	Y      int16
	ItemId int16
	Prefix byte
	Stack  int16
}

// Server <-> Client (Sync)
type UpdateItemDrop2 struct {
	ItemID    int16
	PositionX float32
	PositionY float32
	VelocityX float32
	VelocityY float32
	StackSize int16
	Prefix    byte
	NoDelay   byte
	ItemNetID int16
}

// Server -> Client
type SyncEmoteBubble struct {
	EmoteID       int32
	AnchorType    byte
	PlayerID      uint16
	EmoteLifeTime uint16
	Emote         byte
	EmoteMetaData int16
}

// Server <-> Client (Sync)
type SyncExtraValue struct {
	NPCIndex   int16
	ExtraValue int32
	X          float32
	Y          float32
}

// Not used
type SocialHandshake struct {
}

// Not used
type Deprecated struct {
}

// Client -> Server
type KillPortal struct {
	ProjectileOwner uint16
	ProjectileAI    byte
}

// Server <-> Client
type PlayerTeleportPortal struct {
	PlayerID         byte
	PortalColorIndex int16
	NewPositionX     float32
	NewPositionY     float32
	VelocityX        float32
	VelocityY        float32
}

// Server -> Client
type NotifyPlayerNPCKilled struct {
	NPCID int16
}

// Server -> Client
type NotifyPlayerOfEvent struct {
	EventID int16
}

// Server <-> Client (Sync)
type UpdateMinionTarget struct {
	PlayerID byte
	TargetX  float32
	TargetY  float32
}

// Server <-> Client
type NPCTeleportPortal struct {
	NPCID            uint16
	PortalColorIndex int16
	NewPositionX     float32
	NewPositionY     float32
	VelocityX        float32
	VelocityY        float32
}

// Server -> Client
type UpdateShieldStrengths struct {
	SolarTowerShieldStrength    uint16
	VortexTowerShieldStrength   uint16
	NebulaTowerShieldStrength   uint16
	StardustTowerShieldStrength uint16
}

// Server <-> Client (Sync)
type NebulaLevelUp struct {
	PlayerID    byte
	LevelUpType uint16
	OriginX     float32
	OriginY     float32
}

// Server -> Client
type MoonLordCountdown struct {
	MoonLordCountdown int32
}

// Server -> Client
type NPCShopItem struct {
	Slot     byte
	ItemType int16
	Stack    int16
	Prefix   byte
	Value    int32
	Flags    byte
}

// Client -> Server
type GemLockToggle struct {
	X  int16
	Y  int16
	On bool
}

// Server -> Client
type PoofofSmoke struct {
	PackedVector uint32
}

// Server -> Client
type SmartTextMessage struct {
	MessageColor  structs.Color
	Message       tile.NetworkText
	MessageLength int16
}

// Server -> Client
type WiredCannonShot struct {
	Damage    int16
	Knockback float32
	X         int16
	Y         int16
	Angle     int16
	Ammo      int16
	PlayerID  byte
}

// Client -> Server
type MassWireOperation struct {
	StartX   int16
	StartY   int16
	EndX     int16
	EndY     int16
	ToolMode byte
}

// Server -> Client
type MassWireOperationConsume struct {
	ItemType int16
	Quantity int16
	PlayerID byte
}

// Client -> Server
type ToggleBirthdayParty struct {
}

// Server <-> Client (Sync)
type GrowFX struct {
	EffectFlags byte
	X           int32
	Y           int32
	Data        byte
	TreeGore    int16
}

// Client -> Server
type CrystalInvasionStart struct {
	X int16
	Y int16
}

// Server -> Client
type CrystalInvasionWipeAll struct {
}

// Client -> Server
type MinionAttackTargetUpdate struct {
	PlayerID           byte
	MinionAttackTarget int16
}

// Server -> Client
type CrystalInvasionSendWaitTime struct {
	TimeUntilNextWave int32
}

// Client -> Server
type PlayerHurtV2 struct {
	PlayerID          byte
	PlayerDeathReason tile.PlayerDeathReason
	Damage            int16
	HitDirection      byte
	Flags             byte
	CooldownCounter   int8
}

// Client -> Server
type PlayerDeathV2 struct {
	PlayerID          byte
	PlayerDeathReason tile.PlayerDeathReason
	Damage            int16
	HitDirection      byte
	Flags             byte
}

// Client <-> Server
type CombatTextString struct {
	X          float32
	Y          float32
	Color      structs.Color
	CombatText tile.NetworkText
}

// Client -> Server
type Emoji struct {
	PlayerIndex byte
	EmoticonID  byte
}

// Client <-> Server
type TEDisplayDollItemSync struct {
	PlayerID     byte
	TileEntityID int32
	ItemIndex    byte
	ItemID       uint16
	Stack        uint16
	Prefix       byte
}

// Client <-> Server
type RequestTileEntityInteraction struct {
	TileEntityID int32
	PlayerID     byte
}

// Client -> Server
type WeaponsRackTryPlacing struct {
	X      int16
	Y      int16
	NetID  int16
	Prefix byte
	Stack  int16
}

// Client <-> Server
type TEHatRackItemSync struct {
	PlayerID     byte
	TileEntityID int32
	ItemIndex    byte
	ItemID       uint16
	Stack        uint16
	Prefix       byte
}

// Client <-> Server
type SyncTilePicking struct {
	PlayerID   byte
	X          int16
	Y          int16
	PickDamage byte
}

// Server -> Client
type SyncRevengeMarker struct {
	UniqueID          int32
	X                 float32
	Y                 float32
	NPCID             int32
	NPCHPPercent      float32
	NPCType           int32
	NPCAI             int32
	CoinValue         int32
	BaseValue         float32
	SpawnedFromStatue bool
}

// Server -> Client
type RemoveRevengeMarker struct {
	UniqueID int32
}

// Client <-> Server
type LandGolfBallInCup struct {
	PlayerID     byte
	X            uint16
	Y            uint16
	NumberofHits uint16
	ProjID       uint16
}

// Server -> Client
type FinishedConnectingToServer struct {
}

// Client -> Server
type FishOutNPC struct {
	X     uint16
	Y     uint16
	NPCID int16
}

// Server -> Client
type TamperWithNPC struct {
	NPCID            uint16
	SetNPCImmunity   byte
	ImmunityTime     int32
	ImmunityPlayerID int16
}

// Server -> Client
type PlayLegacySound struct {
	X          float32
	Y          float32
	SoundID    uint16
	SoundFlags byte
}

// Client -> Server
type FoodPlatterTryPlacing struct {
	X      int16
	Y      int16
	ItemID int16
	Prefix byte
	Stack  int16
}

// Client <-> Server
type UpdatePlayerLuckFactors struct {
	PlayerID                 byte
	LadybugLuckTimeRemaining int32
	TorchLuck                float32
	LuckPotion               byte
	HasGardenGnomeNearby     bool
}

// Server -> Client
type DeadPlayer struct {
	PlayerID byte
}

// Client <-> Server
type SyncCavernMonsterType [6]uint16

// Client -> Server
type RequestNPCBuffRemoval struct {
	NPCID  int16
	BuffID uint16
}

// Client -> Server
type ClientFinishedInventoryChangesOnThisTick struct {
}

// Server -> Client
type SetCountsAsHostForGameplay struct {
	PlayerID     byte
	CountsAsHost bool
}
