package main

import (
	"backend-comee/api"
	"backend-comee/internal/env"
)

func main() {
	c := env.NewConfiguration()
	/*port := os.Getenv("PORT")

	puerto, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error fatal in asignature port")
	}*/
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
