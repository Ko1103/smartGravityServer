

package main
import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/go-xorm/xorm"
    "github.com/go-sql-driver/mysql"
    "time"
    "model"
)


var engine *xorm.Engine //DBを変数の定義

func main() {
  var err error
  engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8") //engineの作成 engineはORM DBのこと

  router := gin.Default()
  v1 := router.Group("/api/v1/todos")
  {
    v1.POST("/", createTodo)
    v1.GET("/", fetchAllTodo)
    v1.GET("/:id", fetchSingleTodo)
    v1.PUT("/:id", updateTodo)
    v1.DELETE("/:id", deleteTodo)
  }
  router.Run()
  
  //init mysql DB
  engine, err := xorm.NewEngine("mysql", )
}

type User struct {
  Id int64 `json:"Id"  xorm:"`id`"`
  Name string `json:"name" xorm:"`name`"`
  Password string `json:"password" xorm:"`password`"`
}

func deleteWeight() {
  var weight 
}s