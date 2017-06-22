package geomiddle

import (
	"math"
	"testing"
)

func Test_CalculateMidPoint(t *testing.T) {
	type args struct {
		locations []Location
	}
	tests := []struct {
		name string
		args args
		want Location
	}{
		{"same location", args{[]Location{Location{Lat: 51.6978, Long: 5.3037}, Location{Lat: 51.6978, Long: 5.3037}}}, Location{Lat: 51.6978, Long: 5.3037}},
		{"amsterdam, den bosch -> middle = nieuwegein", args{[]Location{Location{Lat: 51.6978, Long: 5.3037}, Location{Lat: 52.3702, Long: 4.8952}}}, Location{Lat: 52.034176, Long: 5.10098}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMidPoint(tt.args.locations); !(math.Abs(got.Lat-tt.want.Lat) < 0.0001) || !(math.Abs(got.Long-tt.want.Long) < 0.0001) {
				t.Errorf("getMidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
