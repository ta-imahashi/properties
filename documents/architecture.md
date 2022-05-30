# Architecture


```plantuml

actor user
node Nextjs
node "Nestjs[with Apollo]" as Nestjs
node "API(Go/Gin)" as api
database dynamodb

user -> Nextjs: 1.access
user <- Nextjs: 2.return html/js
user -d-> Nestjs: 3.Query/Mutation
Nestjs -d-> api: 4.GET/POST/PUT/DELETE
api -d-> dynamodb: 5.Scan/Put/Update/Delete

```
