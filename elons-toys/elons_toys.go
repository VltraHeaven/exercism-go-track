package elon

import (
    "strconv"
)

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
    if c.battery >= c.batteryDrain {
        c.battery -= c.batteryDrain
        c.distance += c.speed
    }
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
    return "Driven " + strconv.Itoa(c.distance) + " meters"
}
// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
    return "Battery at " + strconv.Itoa(c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
    return c.battery / c.batteryDrain * c.speed >= trackDistance
}