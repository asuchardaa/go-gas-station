package main

// Metoda pro získání volné pokladny
func findRegister(registers []*CashRegister) {
	for car := range buildingQueue {
		var bestRegister *CashRegister
		bestQueueLength := -1
		for _, register := range registers {
			queueLength := len(register.Queue)
			if bestQueueLength == -1 || queueLength < bestQueueLength {
				bestRegister = register
				bestQueueLength = queueLength
			}
		}
		bestRegister.Queue <- car
	}
	for _, register := range registers {
		close(register.Queue)
	}
}

// Metoda pro získání volného stojanu
func findStandRoutine(stands []*FuelStand) {
	for car := range arrivals {
		// Vyber nejkratší frontu
		var bestStand *FuelStand
		bestQueueLength := -1

		// Prohledávání stojanů
		for _, stand := range stands {
			if stand.Type == car.Fuel {
				queueLength := len(stand.Queue)
				if bestQueueLength == -1 || queueLength < bestQueueLength {
					bestStand = stand
					bestQueueLength = queueLength
				}
			}
		}

		// Rozdělení aut mezi stojany podle délky fronty
		for _, stand := range stands {
			if stand.Type == car.Fuel && len(stand.Queue) == bestQueueLength {
				bestStand = stand
				break
			}
		}

		bestStand.Queue <- car
	}

	for _, stand := range stands {
		close(stand.Queue)
	}
}
