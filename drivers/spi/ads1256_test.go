package spi

import (
	"testing"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/gobottest"
)

var _ gobot.Driver = (*ADS1256Driver)(nil)

// must implement the AnalogReader interface
var _ aio.AnalogReader = (*ADS1256Driver)(nil)

func initTestADS1256Driver() *ADS1256Driver {
	d := NewADS1256Driver(&TestConnector{})
	return d
}

func TestADS1256DriverStart(t *testing.T) {
	d := initTestADS1256Driver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestADS1256DriverHalt(t *testing.T) {
	d := initTestADS1256Driver()
	d.Start()
	gobottest.Assert(t, d.Start(), nil)
}

func TestADS1256DriverRead(t *testing.T) {
	d := initTestADS1256Driver()
	d.Start()
	
	// TODO: actual read test
}
