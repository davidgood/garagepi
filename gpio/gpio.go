package gpio

import (
	"fmt"
	"strconv"

	"github.com/davidgood/garagepi/os"
	"github.com/pivotal-golang/lager"
	"github.com/stianeikeland/go-rpio"
)

//go:generate counterfeiter . Gpio

// Gpio Interface
type Gpio interface {
	Read(pin uint) (string, error)
	WriteLow(pin uint) error
	WriteHigh(pin uint) error
}

type gpio struct {
	osHelper os.OSHelper
	logger   lager.Logger
}

// NewGpio ctor
func NewGpio(
	osHelper os.OSHelper,
	logger lager.Logger,
) Gpio {
	return &gpio{
		osHelper: osHelper,
		logger:   logger,
	}
}

func (g gpio) Read(pin uint) (string, error) {
	g.logger.Debug("reading from pin", lager.Data{"pin": pin})

	rpin := rpio.Pin(pin)

	err := rpio.Open()
	if err != nil {
		return "", err
	}
	defer rpio.Close()

	state := rpin.Read()
	return fmt.Sprintf("%v", state), err
}

func writePin(g gpio, pin uint, state rpio.State) error {
	g.logger.Debug("writing low to pin", lager.Data{"pin": pin})

	rpin := rpio.Pin(pin)

	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	rpin.Output()
	if state == rpio.Low {
		rpin.Low()
	} else {
		rpin.High()
	}
	return nil
}

func (g gpio) WriteLow(pin uint) error {
	return writePin(g, pin, rpio.Low)
}

func (g gpio) WriteHigh(pin uint) error {
	return writePin(g, pin, rpio.High)
}

func tostr(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
