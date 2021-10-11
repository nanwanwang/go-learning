package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
  router:= gin.Default()
  //router.GET("/ping", func(context *gin.Context) {
	//  context.JSON(http.StatusOK,gin.H{
	//  	"message":"pong",
	//  })
  //})

  todoV1:=router.Group("/api/v1/todos")
  {
  	todoV1.POST("/",createTodo)
  	todoV1.GET("/",fetchAllTo)
  	todoV1.GET("/:id",fetchTodo)
  	todoV1.PUT("/:id",updateTodo)
  	todoV1.DELETE("/:id",deleteTodo)

  }




  router.Run()
}


func createTodo(c *gin.Context){

}
