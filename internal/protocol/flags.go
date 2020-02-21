package protocol

type GamePlayerSlot uint32

type GamePlayerIndex uint32

func (i GamePlayerIndex) GamePlayerSlot() GamePlayerSlot {
	return GamePlayerSlot(((i / 5) << 7) + (i % 5))
}

type BuildingState uint32
