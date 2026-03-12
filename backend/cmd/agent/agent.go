package agent

import(
	"fmt"
	"os"
	"sync"
	"backend/internal/models"
	"backend/internal/agent/sender"
)

func RunAgent(exchanges []models.Exchange, coins []string) {
	var wg sync.WaitGroup

	for i, bcoin := range coins {
		for j, qcoin := range coins {
			if i == j {
				continue
			}
			for _, ex := range exchanges {
				wg.Add(1)
				go func(b string, q string, e models.Exchange) {
					defer wg.Done()
					stat, err := ex.GetStat(b, q)
					if err != nil {
						fmt.Fprintf(os.Stderr, "ERROR: (%v)\n", err)
						return
					}
					fmt.Printf("I want SEND\n")
					err = sender.SendStat("http://localhost:8080/update", &stat)
										if err != nil {
						fmt.Fprintf(os.Stderr, "ERROR: (%v)\n", err)
						return
					}
					fmt.Printf("I SEND\n")
				}(bcoin, qcoin, ex)
			}
		}
	}
	wg.Wait()
	fmt.Printf("end \n")
}
