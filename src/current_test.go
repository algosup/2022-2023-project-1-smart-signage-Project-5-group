package main

import (
	"testing"
)

func TestSensorValue(t *testing.T) {
	t.Run("Sensor value is high", func(t *testing.T) {
		sensorValue := 100
		if sensorValue < 10 {
			t.Errorf("The sensor value is high")
		}
	})
	t.Run("Sensor value is low", func(t *testing.T) {
		sensorValue := 9
		if sensorValue >= 10 {
			t.Errorf("The sensor value is low")
		}
	})
}