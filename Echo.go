package main

import (  
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
)

func main() {
  server := echo.New()
  server.GET("/", index)
  server.Run(standard.New(":9292"))
}

func index(ctx echo.Context) error{
  //return ctx.String(200, "YEAYY");
  
}