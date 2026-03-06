import "github.com/redis/go-redis/v9"

func ConnectRedis() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Адрес, который мы пробросили в Docker
        Password: "",           // По умолчанию пароля нет
        DB:       0,            // Номер базы
    })
}
