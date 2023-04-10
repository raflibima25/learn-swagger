package main

import (
	// "DTS-Materi/learn-swagger-materi/controllers"
	"DTS-Materi/learn-swagger-materi/routers"
)

func main() {
	// controllers.GetOneCars()
	// controllers.GetAllCars()
	// controllers.CreateCars()
	// controllers.UpdateCars()
	// controllers.DeleteCars()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
