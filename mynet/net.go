package mynet

import "time"

type Net interface {
	Connect() (error, time.Duration)
	Ping()
}
