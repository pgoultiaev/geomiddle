package geomiddle

import (
	"math"
	"testing"
)

func Test_getMidPoint(t *testing.T) {
	type args struct {
		locations []location
	}
	tests := []struct {
		name string
		args args
		want location
	}{
		{"same location", args{[]location{location{Lat: 51.6978, Long: 5.3037}, location{Lat: 51.6978, Long: 5.3037}}}, location{Lat: 51.6978, Long: 5.3037}},
		{"amsterdam, den bosch -> middle = nieuwegein", args{[]location{location{Lat: 51.6978, Long: 5.3037}, location{Lat: 52.3702, Long: 4.8952}}}, location{Lat: 52.034176, Long: 5.10098}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMidPoint(tt.args.locations); !(math.Abs(got.Lat-tt.want.Lat) < 0.0001) || !(math.Abs(got.Long-tt.want.Long) < 0.0001) {
				t.Errorf("getMidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
