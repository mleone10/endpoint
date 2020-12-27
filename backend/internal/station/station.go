package station

import (
	"time"

	"github.com/segmentio/ksuid"
)

// An ID is a unique identifier used to identify a station.
type ID string

// A ModuleType is a top-level module category.
type ModuleType string

// A Resource is something the station can have, produce, or consume.
type Resource string

// A Rate is the speed at which a Resource's Amount changes every second.
type Rate float64

// An Amount represents how much of a given resource a station has.
type Amount float64

// A Timestamp is a moment in time.
type Timestamp time.Time

// Production describes the various Rates at which a module produces or consumes a resource.
type Production map[Resource]Rate

// Resources represent how much of each Resource a station has.
type Resources map[Resource]Quantity

// Station represents a top-level game state.
type Station struct {
	ID        ID        `json:"id"`
	Modules   []Module  `json:"mods"`
	Resources Resources `json:"res"`
}

// A Module is a component added to a Station.
type Module struct {
	Type       ModuleType `json:"type"`
	Production Production `json:"prod"`
}

// A Quantity represents two things: an amount, and the time at which that amount was recorded.
type Quantity struct {
	Amount Amount
	Time   time.Time
}

// Resources are things which the station has, produces, or consumes.
const (
	ResourceCrew     = Resource("crew")
	ResourceFood     = Resource("food")
	ResourceFunds    = Resource("funds")
	ResourceOxygen   = Resource("oxy")
	ResourceResearch = Resource("sci")
	ResourceWater    = Resource("h20")
)

// ModuleTypes are top-level module categories.
const (
	ModuleTypeCommand = ModuleType("command")
)

var (
	// ModuleCommand is the initial station module.
	ModuleCommand = Module{
		Type: ModuleTypeCommand,
		Production: Production{
			ResourceResearch: Rate(5.0),
		},
	}
)

// New returns an newly created Station with initial configuration.
func New() Station {
	return Station{
		ID: ID(ksuid.New().String()),
		Modules: []Module{
			ModuleCommand,
		},
		Resources: Resources{
			ResourceCrew:  newQuantity(10),
			ResourceFunds: newQuantity(10000),
		},
	}
}

func newQuantity(init Amount) Quantity {
	return Quantity{
		Amount: init,
		Time:   time.Now(),
	}
}
