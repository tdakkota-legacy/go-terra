package structs

//procm:use=derive_binary
type NetworkTextMode byte

const (
	Literal NetworkTextMode = iota
	Formattable
	LocalizationKey
)

//procm:use=derive_binary
type NetworkText struct {
	Mode             NetworkTextMode
	Text             string
	SubstitutionList []NetworkText
}
