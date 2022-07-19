package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    var spykidsrule = false
    if knightIsAwake || archerIsAwake || prisonerIsAwake {
		spykidsrule = true
    }
	return spykidsrule
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    var result = prisonerIsAwake
    if prisonerIsAwake {
        result = !archerIsAwake  
    }
	return result
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var result = false
    if petDogIsPresent {
       result = !archerIsAwake
    } else if prisonerIsAwake && prisonerIsAwake != archerIsAwake && prisonerIsAwake != knightIsAwake {
		result = prisonerIsAwake            
    }
	return result
}
