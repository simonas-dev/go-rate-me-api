# Rate Me API
### Written in GO lang using Macaron web framework.

Use this for Go version management. https://github.com/moovweb/gvm

Dependencies
```
go get gopkg.in/macaron.v1
go get github.com/go-macaron/pongo2
go get github.com/jinzhu/gorm
go get github.com/go-macaron/binding
go get github.com/go-sql-driver/mysql
go get github.com/buger/jsonparser
```

### POST /rating
```
{
    "app_package_name": "com.nodomain.app_name",
    "rating": 1.2,
    "description": "very good app",
    "email": "lol@xd.lt",
    "device_id": "asdasd",
    "ip": "123.123.123.123"
}
```