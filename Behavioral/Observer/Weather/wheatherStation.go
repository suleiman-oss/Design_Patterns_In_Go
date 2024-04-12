package weather

import "fmt"

type WeatherObserver interface {
	Update(temp float32, press float32)
}

type Mobile struct{}

func (m Mobile) Update(temp float32, press float32) {
	fmt.Printf("The temperature is %fC and pressure is %fATP\treporting from Mobile\n", temp, press)
}

type Console struct{}

func (c Console) Update(temp float32, press float32) {
	fmt.Printf("The temperature is %fC and pressure is %fATP\treporting from Console\n", temp, press)
}

type WeatherStation struct {
	observers []WeatherObserver
	temp      float32
	press     float32
}

func (w *WeatherStation) Register(obs WeatherObserver) {
	w.observers = append(w.observers, obs)
}
func (w *WeatherStation) SetMeasurements(temperature float32, pressure float32) {
	w.temp = temperature
	w.press = pressure
	w.notifyObservers()
}
func (w *WeatherStation) notifyObservers() {
	for _, obs := range w.observers {
		obs.Update(w.temp, w.press)
	}
}
