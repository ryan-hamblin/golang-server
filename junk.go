package main

import(
    "fmt"
    r "github.com/dancannon/gorethink"
)

type User struct {
    Id string `gorethink:"id,omitempty"`
    Name string `gorethink:"name"`
}

func main(){
    session, err := r.Connect(r.ConnectOpts{
        Address: "localhost:28015",
        Database: "rtsupport",

    })
    if err != nil{
        fmt.Println(err)
        return
    }

    // response, err := r.Table("user").
    //     Insert(user).
    //     RunWrite(session)
    //     if err != nil{
    //         fmt.Println(err)
    //         return
    //     }
    r.Table("user").get("4e163787-17e9-43eb-ba62-92e3e27efcb1").
        Update(user).
        RunWrite(session)
    user := User{
        Name: "Riggins",
    }
        fmt.Printf("%#v\n", response)

}