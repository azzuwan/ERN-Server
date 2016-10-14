package main

import(
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
  rdb "gopkg.in/dancannon/gorethink.v2"
  "fmt"
) 

func main(){
  //Database
  session, err := rdb.Connect(rdb.ConnectOpts{
      Address: "localhost:28015",
      Database: "ERN",
    })

  if err !=nil {
    fmt.Println(err)
  }

  res,err1 := rdb.DB("ERN").Table("alerts").Run(session)
  if err1 != nil {
    fmt.Println(err)
  }

  var rows []interface{}
  err2 := res.All(&rows)

  if err2 !=nil {
    fmt.Println(err)
  }

  for _, row := range rows{
    fmt.Printf("%v",row)
  }


  
  //Configuration
  iris.Static("/public", "./public", 1)
  iris.Static("/data", "./data", 1)
  iris.Config.IsDevelopment = true
  iris.UseTemplate(html.New(html.Config{ Layout: "layout.html"})).Directory("./templates", ".html")
  iris.Config.Websocket.Endpoint = "/sock"

  //Routes
  iris.Get("/", index)
  iris.Get("/node", node)  
  iris.Get("/video", video)
  iris.Get("/test/video", downloadVideo)
  iris.Get("/stream", stream)

  //Websocket
  iris.Websocket.OnConnection( onConnection)

  iris.Listen("0.0.0.0:9292")
}

func index(ctx *iris.Context){
  ctx.MustRender("index.html", nil)
}

func node(ctx *iris.Context){
  ctx.MustRender("node.html", nil)
}

func video(ctx *iris.Context){
  ctx.MustRender("video.html", nil)
}

func downloadVideo(ctx *iris.Context){
  file := "./data/sample.mp4"
  ctx.SendFile(file, "sample.mp4")
}

func stream(ctx *iris.Context){
  err := ctx.Render("stream.html", nil, iris.RenderOptions{"layout": iris.NoLayout}) 
  if err != nil {
    println(err.Error())
  }
}

func onConnection(c iris.WebsocketConnection){
  channel := "ern-channel"
  c.Join(channel)
  c.On("monitor", func(message string) {
      // all connections which are inside this room will receive this message
    fmt.Println("Message from websocket client! ")
      //c.To(channel).Emit("chat", "From: "+c.ID()+": "+message)
  })

  c.OnDisconnect(func() {
      fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
  })
}
