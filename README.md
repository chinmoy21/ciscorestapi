A REST API in Golang to showcase DB access, insert, update and delete.

The API is modeled based on the guidelines in the below link:

[Building Golang with Echo](https://medium.com/@kyawmyintthein/building-golang-restful-api-with-echo-framework-1-422abc78e3a7), echo framework

The two third-party packages used are:
- [echo](https://echo.labstack.com/), HTTP web server
- [go-sqlite3](https://github.com/mattn/go-sqlite3), database driver handlers

Package execution:

    1) Download the zip from Github link.
    2) Install go package.
    3) Run "go get github.com/chin/ciscorestapi
    4) Move to "$GOPATH/github.com/chin/ciscorestapi"
    5) Run "go build main.go"
    6) On the web browser, run "http://localhost:1200/"

The available endpoints are:

- HOME
  - GET `/`  to print readme

- People
  - GET `/people/` to list all the people entries
  - GET `/people/:id` to show a particular people entry
  - POST `/people/` to create a people entry
  - PUT `/people/:id` to update a people entry
  - DELETE `/people/:id` to remove a people entry

- Vehicles
  - GET `/vehicles/` to list all the vehicles entries
  - GET `/vehicles/:id` to show a particular vehicles entry

- Starships
  - GET `/starships/` to list all the starships entries
  - GET `/starships/:id` to show a particular starships entry

---


