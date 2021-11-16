package go_testing_baru

type Speeder interface {
	MaxSPeed() int
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
	if defaultSpeed > c.Speeder.MaxSPeed() {
		return c.Speeder.MaxSPeed()
	}
	return defaultSpeed
}

type DefaultEngine struct {
}

func (e DefaultEngine) MaxSPeed() int {
	return 51
}

type TurboEngine struct {
}

func (e TurboEngine) MaxSPeed() int {
	return 100
}
