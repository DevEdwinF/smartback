package services

import "log"

func RunCronJob() {

	c := cron.New()

	c.AddFunc("0 0 * * *", func() {
		err := syncData()
		if err != nil {
			log.Println("Error al sincronizar datos:", err)
		}
	})

	c.Start()

	select {}
}
