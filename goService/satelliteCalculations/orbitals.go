package satellitecalculations

import "github.com/Anandhu3301/satelliteGohelper/structs"

func OrbitalsCalc( channelStream chan <- structs.OrbitalStruct ) {
	for i := 0; i < 8; i++ {
		channelStream <- structs.OrbitalStruct{Data: i+1, Flag: true}
	}
	channelStream <- structs.OrbitalStruct{Data: "END",Flag: false}
}