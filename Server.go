package main

import(
  "github.com/kataras/iris"
  "github.com/kataras/go-template/html"
) 

func main(){
  //api := iris.New()

  //Configuration
  iris.Static("/public", "./public", 1)
  iris.Static("/data", "./data", 1)
  iris.Config.IsDevelopment = true
  iris.UseTemplate(html.New(html.Config{
        Layout: "layout.html",
    })).Directory("./templates", ".html")
  

  //Routes
  iris.Get("/", index)
  iris.Get("/node", node)
  iris.Get("/video", video)
  iris.Get("/test/video", downloadVideo)
  iris.Get("/stream", stream)  
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
