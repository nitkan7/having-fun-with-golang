package main

import "github.com/gofiber/fiber/v2"



func main(){
app:=fiber.New() //creates a new fiber instance 

//app.Method
app.Get("/",func(c *fiber.Ctx)error{
return c.SendString("Hello, World!")
})

app.Listen(":3000")
}
