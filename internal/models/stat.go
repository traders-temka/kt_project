import "time"

type Stat struct
{
	Currency string 'json:"currency"'
	Price float64 'json:"price"'
	Source string 'json:"source"'
	Timedump time.Time 'json:"timedump"'
}
