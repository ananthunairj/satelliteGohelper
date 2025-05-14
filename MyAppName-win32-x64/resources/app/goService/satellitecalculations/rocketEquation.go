package satellitecalculations

import "math"

func EscapeVelocityFinder(altitude float64) float64 {
	//v = sq.root of G (Gravitational contant) / R (Radius of the Earth) * r (altitude from earth center to the orbit)
	var centertoSurfaceDistance float64 = 6371e3 //approx km in sreehariKota
	var G float64 = 6.6743e-11;
	var M float64 = 5.97219e24;
	var r float64 = altitude + centertoSurfaceDistance;
	var velocity float64 = math.Sqrt((G * M) / r)
	return velocity / 1000  //return in km/s
}

//vf need for LEO = 9.5 , 18e2 km,Geostatioanry Transfer Orbit(GTO) 10.3 -> 34e3, Escape(moon / mars) 11.2 
//ratio distribution of velocities 45%, 35%, 20%
//lvm3 Isp => s200 -> 274s, L110 -> 316s, C25 -> 442s
//s200 2045e2 kg each, total mass -> 236e3 kg each
//L110 115e3 total mass 1256e2
//c25 28e3 total 313e2
