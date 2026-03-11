# handler
принимает json Обьект, вносит его в структуру. Валидирует аднные, хеширует, отправляет в usecase

## decerialization
- *.json to models.Stat
## validation
- checks for nullptr
- checks if all data is correct
## security
- if we'll add hashing, it would check data signs with middlewares
- used for new users or user authorization
