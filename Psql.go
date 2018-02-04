package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)


func main(){
    db, err := gorm.Open("postgres", 
        "host=localhost user=rsuh1 dbname=opencart sslmode=disable password=qwer1234")
        if err != nil{
            panic(err.Error())
        }
    defer db.Close()
	 //
     //db.CreateTable(&Post{})
     //db.DropTableIfExists(&Post{})
     //post := Post{
     //    Title:"asd",
     //    Body: "asdasd",
     //}
     //db.Debug().Delete(&post)

    var p Post
    db.Debug().First(&p)
    fmt.Println(p)

}

//type Post struct {
//    gorm.Model
//    Title  string
//    Body  string
//
//}