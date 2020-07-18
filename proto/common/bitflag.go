package common

type ByteBitFlag byte

func (b ByteBitFlag) Set(flag ByteBitFlag) ByteBitFlag { return b | flag }
func (b ByteBitFlag) Has(flag ByteBitFlag) bool        { return b&flag != 0 }

type Uint16BitFlag byte

func (b Uint16BitFlag) Set(flag Uint16BitFlag) Uint16BitFlag { return b | flag }
func (b Uint16BitFlag) Has(flag Uint16BitFlag) bool          { return b&flag != 0 }
