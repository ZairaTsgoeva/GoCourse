package models

// Base ...
type Base struct {
	Brand           string
	Year            int
	BootVolume      float32
	IsEngineStarted bool
	IsWindowOpend   bool
	BootFill        float32
}

// Car ...
type Car struct {
	Base      Base
	ClassName string
}

// Truck ...
type Truck struct {
	Base       Base
	WheelCount int
}
