package go_testing_baru

type Speeder interface {
	MaxSpeed() int
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	defaultSpeed := 80
	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}
	return defaultSpeed
}

type DefaultEngine struct {
}

func (e DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct {
}

func (e TurboEngine) MaxSpeed() int {
	return 100
}
