package config

import (
    //"fmt"
)
//func GetDb() string {

type Db struct {
    url, port, login, password, database string
}

//bob := User{"Bob", 25, -50, 4.2, 0.8, []string {"Beer","Eat","Sleep"}}

func (u Db) Getdb() string {
    return "fghfghfghf"+u.url
    //return u.login+":"+u.password+"@tcp("+u.url+":"+u.port+")/"+u.database
    //return fmt.Sprintf("%s:%s/@tcp(%s:%s)/%s", u.login, u.password, u.url, u.port, u.database)
}

func (u *Db) setNewName(newName string) {
    //u.Name = newName
    u.url="127.0.0.1"
    u.port="3306"
    u.login="root"
    u.password="root"
    u.database="rcmoney"
}

//}
