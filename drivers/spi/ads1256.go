package spi

import (
	"strconv"

	"gobot.io/x/gobot"
)

// ADS1256DriverMaxChannel is the number of channels of this A/D converter.
const ADS1256DriverMaxChannel = 8

// ADS1256Driver is a driver for the ADS1256 A/D converter.
type ADS1256Driver struct {
	name       string
	connector  Connector
	connection Connection
	Config
	gobot.Commander
}

// NewADS1256Driver creates a new Gobot Driver for ADS1256Driver A/D converter.
//
// Params:
//     a *Adapter - the Adapter to use with this Driver
//
// Optional params:
//     spi.WithBus(int):     bus to use with this driver
//     spi.WithChip(int):    chip to use with this driver
//     spi.WithMode(int):    mode to use with this driver
//     spi.WithBits(int):    number of bits to use with this driver
//     spi.WithSpeed(int64): speed in Hz to use with this driver
//
func NewADS1256Driver(a Connector, options ...func(Config)) *ADS1256Driver {
	d := &ADS1256Driver{
		name:      gobot.DefaultName("ADS1256"),
		connector: a,
		Config:    NewConfig(),
	}
	for _, option := range options {
		option(d)
	}
	return d
}

// Name returns the name of the device.
func (d *ADS1256Driver) Name() string { return d.name }

// SetName sets the name of the device.
func (d *ADS1256Driver) SetName(n string) { d.name = n }

// Connection returns the Connection of the device.
func (d *ADS1256Driver) Connection() gobot.Connection { return d.connection.(gobot.Connection) }

// Start initializes the driver.
func (d *ADS1256Driver) Start() (err error) {
	bus := d.GetBusOrDefault(d.connector.GetSpiDefaultBus())
	chip := d.GetChipOrDefault(d.connector.GetSpiDefaultChip())
	mode := d.GetModeOrDefault(d.connector.GetSpiDefaultMode())
	bits := d.GetBitsOrDefault(d.connector.GetSpiDefaultBits())
	maxSpeed := d.GetSpeedOrDefault(d.connector.GetSpiDefaultMaxSpeed())

	d.connection, err = d.connector.GetSpiConnection(bus, chip, mode, bits, maxSpeed)
	if err != nil {
		return err
	}
	return nil
}

// Halt stops the driver.
func (d *ADS1256Driver) Halt() (err error) {
	d.connection.Close()
	return
}

// Read reads the current analog data for the desired channel.
func (d *ADS1256Driver) Read(channel int) (result int, err error) {
	tx := make([]byte, 3)
	tx[0] = 0x01
	tx[1] = 0x80 + (byte(channel) << 4)
	tx[2] = 0x00

	d.connection.Tx(tx, nil)

	return 0, nil
}

// AnalogRead returns value from analog reading of specified pin
func (d *ADS1256Driver) AnalogRead(pin string) (value int, err error) {
	channel, _ := strconv.Atoi(pin)
	value, err = d.Read(channel)

	return
}
