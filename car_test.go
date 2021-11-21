package go_testing_baru

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "default engine should have 50 max speeed",
			want: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DefaultEngine{}
			if got := e.MaxSpeed(); got != tt.want {
				t.Errorf("MaxSpeed() = %v , want %v", got, tt.want)

			}
		})
	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	test := []struct {
		name string
		want int
	}{
		{
			name: "should have 100 as max speed",
			want: 100,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			e := TurboEngine{}
			if got := e.MaxSpeed(); got != tt.want {
				t.Errorf("MaxSpeed() = %v , want %v", got, tt.want)
			}
		})
	}
}

type FakeEngine struct {
}

func (e FakeEngine) MaxSpeed() int {
	return 5
}

type MockEngine struct {
	mock.Mock
}

func (m MockEngine) MaxSpeed() int {
	args := m.Called()
	return args.Get(0).(int)
}

func TestCar_Speed_WithMock(t *testing.T) {
	mock := new(MockEngine)
	car := Car{
		Speeder: mock,
	}

	mock.On("MaxSpeed").Return(9)
	speed := car.Speed()
	assert.Equal(t, 20, speed)

}
func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "speed should be 50 when use default engine",
			fields: fields{Speeder: &DefaultEngine{}},
			want:   50,
		}, {
			name:   "speed should be 80 when use Turbo engine",
			fields: fields{Speeder: &TurboEngine{}},
			want:   80,
		}, {
			name:   "speed should be 20 when use max speed is less than 10",
			fields: fields{Speeder: &FakeEngine{}},
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Speeder: tt.fields.Speeder,
			}

			if got := c.Speed(); got != tt.want {
				t.Errorf("speed() = %v , want %v", got, tt.want)
			}
		})
	}
}
