package main


type MysqlConfig struct{
	Address string `ini:"address"`
	Port    int  `ini:"port"`
	Username  string  `ini:"username"`
	Password  string   `ini:"password"`
}