package helpers

var time []float64 = []float64{0, 10, 30, 60, 127, 153, 237, 324, 330, 973}
var pitch []float64 = []float64{90, 85, 75, 60, 50, 45, 35, 25, 20, 5}

func InterPolatePitch(t float64) float64 {

	if t <= time[0] {
		return pitch[0]
	}
	if t >= time[len(time)-1] {
		return pitch[len(pitch)-1]
	}
	for i := 0; i < len(time)-1; i++ {
		if t >= time[i] && t <= time[i+1] {
			t0, t1 := time[i], time[i+1]
			p0, p1 := pitch[i], pitch[i+1]

			return p0 + (t-t0)*(p1-p0)/(t1-t0)
		}
	}
	return 0

}
