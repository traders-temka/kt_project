package main

import (
)

func main()
{
	repo := storage.NewStorage()

	go func()
	{
		handlers := server.NewHandlers(repo)
		router := server.NewRouter(handlers)

		log.Println("starting server on :8080 : ")
		if err := http.ListenAndServe(":8080", router); err != nil
		{
			log.Fatalf("Server failed: %s", err)
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		// Агент собирает данные
		stat, err := agent.FetchBinancePrice("BTCUSDT")
		if err != nil {
			log.Printf("Agent error: %v", err)
			continue
		}

		// Агент отправляет данные на локальный сервер через HTTP
		// (Даже если они в одном процессе, лучше слать через HTTP для практики)
		err = agent.SendStat("http://localhost:8080/update", stat)
		if err != nil {
			log.Printf("Send error: %v", err)
		}
	}
}
