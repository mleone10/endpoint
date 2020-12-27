package station_test

import (
	"testing"

	"github.com/mleone10/endpoint/internal/station"
	test "github.com/mleone10/endpoint/internal/testing"
)

func TestNewStation(t *testing.T) {
	s := station.New()
	test.AssertNotEquals(t, "", s.ID)
	test.AssertEquals(t, 1, len(s.Modules))
	test.AssertEquals(t, station.ModuleTypeCommand, s.Modules[0].Type)
}
