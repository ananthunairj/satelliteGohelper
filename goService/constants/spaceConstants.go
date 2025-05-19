package constants

import "github.com/Anandhu3301/satelliteGohelper/helpers"


type RocketComponentMap map[helpers.StageType]map[helpers.RocketComponents]float64


const (
	G0      = 9.81
	sizeIsp = 3
	Cd      = 0.5   //drag coeefficient
	Rho0    = 1.225 //air density
	H       = 8500  //height for atmosphere(m)

)
const (
	Thrust helpers.RocketComponents = iota + 1
	Isp
	DryMass
	Fuel
	Burntime
	Diameter
)
const (
	StageOne   helpers.StageType = "booster"
	StageTwo   helpers.StageType = "propellant"
	StageThree helpers.StageType = "cryogenic"
)

var rocketData RocketComponentMap = RocketComponentMap{
	StageOne: {
		Thrust:   5150e3 * 2,
		Isp:      274,
		DryMass:  62e3 * 2,
		Fuel:     205e3 * 2,
		Burntime: 133,
		Diameter: 5,
	},
	StageTwo: {
		Thrust:   1588e3 ,
		Isp:      281,
		DryMass:  9e3,
		Fuel:     116e3,
		Burntime: 203,
		Diameter: 4,
	},
	StageThree: {
		Thrust:   186360,
		Isp:      434,
		DryMass:  5e3,
		Fuel:     28e3,
		Burntime: 637,
		Diameter: 4,
	},
}

func RocketDataSelector(stageChecker helpers.RocketDataFetcher) float64 {
	 return rocketData[stageChecker.Stage][stageChecker.RocketDataSpecific]
}

func GetAllRocketData() [][]float64 {
	var alldata[][]float64
	for _,component := range rocketData {
		stageData := []float64 {
			component[Thrust],
			component[Isp],
			component[DryMass],
			component[Fuel],
			component[Burntime],
			component[Diameter],
		}
		alldata  = append(alldata, stageData)
	}
	return alldata
}
