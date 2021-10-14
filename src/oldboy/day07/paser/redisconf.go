package main



type RedisConfig struct{
   Host  string  `ini:"host"`
   Port   int   `ini:"port"`
   Password  string `ini:"password"`
   Database  int `ini:"database`
	
}