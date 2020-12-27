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
	ID           ID        `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Modules      []Module  `json:"mods"`
	Resources    Resources `json:"res"`
}

// A Module is a component added to a Station.
type Module struct {
	ID         ID         `json:"id"`
	Type       ModuleType `json:"type"`
	Production Production `json:"prod"`
}

// A Quantity represents two things: an amount, and the time at which that amount was recorded.
type Quantity struct {
	Amount Amount    `json:"amt"`
	Time   time.Time `json:"time"`
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
	cm := ModuleCommand
	cm.ID = newID()

	return Station{
		ID:           newID(),
		CreationTime: time.Now(),
		Modules:      []Module{cm},
		Resources: Resources{
			ResourceCrew:  newQuantity(10),
			ResourceFunds: newQuantity(10000),
		},
	}
}

// NetProduction returns the aggregate production rates for the whole station.
func (s Station) NetProduction() Production {
	p := Production{}
	for _, m := range s.Modules {
		for res, rate := range m.Production {
			p[res] += rate
		}
	}
	return p
}

// CurrentResources returns the station's time-adjusted resources quantities.
func (s Station) CurrentResources() map[Resource]Amount {
	now := time.Now()
	rs := map[Resource]Amount{}
	for res, rate := range s.NetProduction() {
		startTime := s.Resources[res].Time
		if _, ok := s.Resources[res]; !ok {
			startTime = s.CreationTime
		}
		rs[res] = Amount(float64(s.Resources[res].Amount) + (float64(rate) * float64((now.Sub(startTime).Seconds()))))
	}
	for res, quant := range s.Resources {
		rs[res] = quant.Amount
	}
	return rs
}

func newQuantity(init Amount) Quantity {
	return Quantity{
		Amount: init,
		Time:   time.Now(),
	}
}

func newID() ID {
	return ID(ksuid.New().String())
}
