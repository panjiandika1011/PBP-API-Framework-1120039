package main

import (
	"fmt"
	"net/http"

	Controller "PBP-API-Framework-1120039/Controller"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/render"
)

func main() {
	//utilization static file serving and routing
	m := martini.Classic()

	//memakai penggunaan layout dan template pada html
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	//req to index on template
	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Get("/schedule", Controller.GetAllSchedule)
	m.Get("/medstaffschedule", Controller.GetMedStaffSchedule)
	m.Post("/add", Controller.InsertSchedule)
	m.Put("/update", Controller.UpdateSchedule)
	m.Delete("/delete/:id", Controller.DeleteSchedule)

	http.Handle("/", m)
	fmt.Println("Connect to port 8080")
	m.RunOnAddr(":8080")
}
