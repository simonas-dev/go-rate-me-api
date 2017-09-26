package main

import(
  "gopkg.in/macaron.v1"
  "github.com/go-macaron/pongo2"
  "github.com/jinzhu/gorm"
  "github.com/go-macaron/binding"
  _ "github.com/go-sql-driver/mysql"
  "github.com/buger/jsonparser"
  "io/ioutil"
  "fmt"
)

type Rating struct {
  gorm.Model
  App_package_name    string `json:"app_package_name"`
  Rating              float32 `json:"rating"`
  Description         string `json:"description"`
  Email               string `json:"email"`
  Device_id           string `json:"device_id"`
  Ip                  string `json:"ip"`
}

func main() {
  config, err := ioutil.ReadFile("./config.json")

  mysql_user, _ := jsonparser.GetString(config, "mysql-user")
  mysql_pass, _ := jsonparser.GetString(config, "mysql-pass")
  mysql_database, _ := jsonparser.GetString(config, "mysql-database")
  server_port, _ := jsonparser.GetInt(config, "server-port")

  if (err != nil) {
    fmt.Println(err)
    return
  }

  // Init Web Server
  macaron.Env = macaron.PROD
  m := macaron.Classic()
  m.Use(macaron.Recovery())
  m.Use(pongo2.Pongoer())

  // Init Database ORM
  connection_params := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", mysql_user, mysql_pass, mysql_database)
  db, err := gorm.Open("mysql", connection_params)
  
  if err != nil {
    panic(err)
  }
  defer db.Close()
  // Migrate the schema
  db.AutoMigrate(&Rating{})

  // Server Status Endpoint
  m.Get("/", func() string {
    return "Alive!"
  })
  
  // Rating Insert Endpoint
  m.Post("/rating", binding.Bind(Rating{}), func(rating Rating) string {
    db.Create(&rating)
    return "Ok"
  })

  m.Run("0.0.0.0", server_port)
}
