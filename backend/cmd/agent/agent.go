package agent

import(
	"backend/internal/models"
	//"backend/internal/agent/collector"
	"backend/internal/agent/sender"
)

func RunAgent(exchanges []models.Exchange, coins []string) {
		for _, coin := range coins {
			for _, ex := range exchanges {
				stat := ex.GetStat(coin)
				sender.SendStat("http://localhost:8080/update", &stat)
			}
		}
}
