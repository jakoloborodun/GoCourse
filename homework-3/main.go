package main

import (
	"GoCourse/homework-3/cars"
	"fmt"
	"math/rand"
)

func main() {
	var lada = cars.Passenger{
		Cars: cars.Cars{
			Brand:             "Lada",
			Model:             "Largus",
			YearOfManufacture: 2015,
			IsEngineRunning:   false,
			IsWindowsOpen:     true,
			Fullness:          156.7,
		},
		TrunkVolume: 300,
	}

	var zil = cars.Truck{
		Cars: cars.Cars{
			Brand:             "ZiL",
			Model:             "ZIL-130",
			YearOfManufacture: 1977,
			IsEngineRunning:   true,
			IsWindowsOpen:     true,
			Fullness:          8560,
		},
		BodyVolume: 6000,
	}

	fmt.Printf("passenger - %T %v\n", lada, lada)
	fmt.Printf("truck - %T %v\n", zil, zil)

	fmt.Println(lada.Brand, lada.Model)

	fmt.Printf("Год производства: %v. %v\n", lada.YearOfManufacture, getEngineStatus(lada.IsEngineRunning))

	for i := 0; i < 5; i++ {
		lada.IsEngineRunning = randBool()
		if lada.IsEngineRunning {
			fmt.Println("Ура, завелась!")
			break
		} else {
			fmt.Println("Будь проклят тот день, когда я сел за баранку этого пылесоса!")
		}
	}

	fmt.Println("Грузовик", zil.Brand, zil.Model, "загружен углем на", zil.Fullness, "килограмм")
	zil.Fullness += 230
	overload := percentage(zil.BodyVolume, zil.Fullness)
	fmt.Println("Загрузили еще угля, процент загруженности - ", overload)

}

func percentage(old, new float64) (delta float64) {
	delta = (new / old) * 100
	return
}

func getEngineStatus(isEngineRunning bool) (engineStatus string) {
	if isEngineRunning {
		engineStatus = "Двигатель работает!"
	} else {
		engineStatus = "Двигатель не заводится."
	}
	return
}

func randBool() bool {
	return rand.Int()%2 == 0
}
