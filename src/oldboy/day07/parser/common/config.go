package common


type Config struct{
     MysqlConfig	`ini:"mysql"`
	 RedisConfig    `ini:"redis"`
}