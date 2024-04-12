package main

import weather "github.com/suleiman-oss/Design_Patterns_in_Go/Behavioral/Observer/Weather"

func main() {
	m := weather.Mobile{}
	c := weather.Console{}
	w := weather.WeatherStation{}
	w.Register(m)
	w.Register(c)
	w.SetMeasurements(32, 106)
}
