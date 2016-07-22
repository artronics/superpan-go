package ieee8021504

const (
	//The number of symbols forming a superframe slot when the superframe
	// order is equal to zero, as described in 5.1.1.1.
	ABaseSlotDuration = 60
	//The number of slots contained in any superframe.
	ANumSuperframeSlots = 16
	//The number of symbols forming a superframe when the superframe order is equal to zero.
	ABaseSuperframeDuration = ABaseSlotDuration * ANumSuperframeSlots
)
