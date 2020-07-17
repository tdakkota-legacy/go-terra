package tile

type NetworkTextMode byte

const (
	Literal NetworkTextMode = iota
	Formattable
	LocalizationKey
)

type NetworkText struct {
	Mode                   NetworkTextMode
	Text                   string
	SubstitutionListLength byte
	SubstitutionList       []NetworkText
}
