package wowapi

// Class will just match the class number to a string
var Class = map[int]string{
	0:  "None",
	1:  "Warrior",
	2:  "Paladin",
	3:  "Hunter",
	4:  "Rogue",
	5:  "Priest",
	6:  "Death Knight",
	7:  "Shaman",
	8:  "Mage",
	9:  "Warlock",
	10: "Monk",
	11: "Druid",
}

// CharacterProfile is the struct that will contain the json RESP from
// the API request
type CharacterProfile struct {
	LastModified      int    `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Class             int    `json:"class"`
	Race              int    `json:"race"`
	Gender            int    `json:"gender"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievementPoints"`
	Thumbnail         string `json:"thumbnail"`
	CalcClass         string `json:"calcClass"`
	Faction           int    `json:"faction"`
	TotalHK           int    `json:"totalHonorableKills"`
}

type CharacterTalents struct {
	CharacterProfile
	Talents []Talent `json:"talents"`
}

type CharacterItems struct {
	CharacterProfile
	Items ItemList `json:"items"`
}

type ItemList struct {
	AverageItemLevel         int        `json:"averageItemLevel"`
	AverageItemLevelEquipped int        `json:"averageItemLevelEquipped"`
	Head                     ItemDetail `json:"head"`
	Neck                     ItemDetail `json:"neck"`
	Shoulder                 ItemDetail `json:"shoulder"`
	Back                     ItemDetail `json:"back"`
	Chest                    ItemDetail `json:"chest"`
	Shirt                    ItemDetail `json:"shirt"`
	Wrist                    ItemDetail `json:"wrist"`
	Hands                    ItemDetail `json:"hands"`
	Waist                    ItemDetail `json:"waist"`
	Legs                     ItemDetail `json:"legs"`
	Feet                     ItemDetail `json:"feet"`
	Finger1                  ItemDetail `json:"finger1"`
	Finger2                  ItemDetail `json:"finger2"`
	Trinket1                 ItemDetail `json:"trinket1"`
	Trinket2                 ItemDetail `json:"trinket2"`
	MainHand                 ItemDetail `json:"mainHand"`
	OffHand                  ItemDetail `json:"offHand"`
}

type ItemDetail struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	Icon          string           `json:"icon"`
	Quality       int           `json:"quality"`
	ItemLevel     int           `json:"itemLevel"`
	ToolTipParams ItemToolTipParam `json:"tooltipParams"`
	Stats         []ItemStats      `json:"stats"`
	Armor         int              `json:"armor"`
	WeaponInfo    ItemWeaponInfo   `json:"weaponInfo"`
	Context       string           `json:"context"`
	BonusLists    []int            `json:"bonusLists"`
}

type ItemWeaponInfo struct {
	Damage      ItemDamage `json:"damage"`
	WeaponSpeed float32    `json:"weaponSpeed"`
	DPS         float32    `json:"dps"`
}

type ItemDamage struct {
	Min      int     `json:"min"`
	Max      int     `json:"max"`
	ExactMin float32 `json:"exactMin"`
	ExactMax float32 `json:"exactMax"`
}

type ItemStats struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

type ItemToolTipParam struct {
	Set			   []int       `json:"set"`
	Gem0		   int		   `json:"gem0"`
	Gem1		   int		   `json:"gem1"`
	Gem2		   int		   `json:"gem2"`
	Gem3		   int		   `json:"gem3"`
	Tinker		   int         `json:"tinker"`
	Enchant        int         `json:"enchant"`
	TransmogItem   int         `json:"transmogItem"`
	Upgrade        ItemUpgrade `json:"upgrade"`
	TimeWakerLevel int         `json:"timewakerLevel"`
}

type ItemUpgrade struct {
	Current            int `json:"current"`
	Total              int `json:"total"`
	ItemLevelIncrement int `json:"itemLevelIncrement"`
}

type Talent struct {
	Selected   bool          `json:"selected,omitempty"`
	Talents    []TalentInner `json:"talents"`
	Glyphs     Glyphs        `json:"glyphs"`
	Spec       Spec          `json:"spec"`
	CalcTalent string        `json:"calcTalent"`
	CalcSpec   string        `json:"calcSpec"`
	CalcGlyph  string        `json:"calcGlyph"`
}

type Spec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

type Glyphs struct {
	Major []Glyph `json:"major"`
	Minor []Glyph `json:"minor"`
}

type Glyph struct {
	Glyph int    `json:"glyph"`
	Item  int    `json:"item"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
}

type TalentInner struct {
	Tier   int   `json:"tier"`
	Column int   `json:"tier"`
	Spell  Spell `json:"spell"`
}

type Spell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
	Cooldown    string `json:"cooldown"`
}
