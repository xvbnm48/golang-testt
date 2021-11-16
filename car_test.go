package go_testing_baru

import "testing"

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
			if got := e.MaxSPeed(); got != tt.want {
				t.Errorf("MaxSpeed() = #{got} , want #{tt.want}")

			}
		})
	}
}
