package main

import (
	"github.com/gin-gonic/gin"
	"restapi.com/hello/db"
	h "restapi.com/hello/helpers"
	p "restapi.com/hello/person"
)

func main() {
	db := db.New()
	r := gin.Default()
	r.PUT("/api/create", func(c *gin.Context) {
		var p p.Person
		if c.BindJSON(&p) != nil {
			return
		}
		id := db.Create(&p)
		h.JOk(c, id)
	})
	r.GET("/api/read", func(c *gin.Context) {
		var id int
		if c.BindJSON(&id) != nil {
			return
		}
		if p, err := db.Read(id); err != nil {
			h.JErr(c, err.Error())
		} else {
			h.JOk(c, p)
		}
	})
	r.PUT("/api/update", func(c *gin.Context) {
		query := &struct {
			ID     int
			Person p.Person
		}{}
		if c.BindJSON(query) != nil {
			return
		}
		if err := db.Update(query.ID, &query.Person); err != nil {
			h.JErr(c, err.Error())
		} else {
			h.JDone(c)
		}
	})
	r.DELETE("/api/delete", func(c *gin.Context) {
		var id int
		if c.BindJSON(&id) != nil {
			return
		}
		if err := db.Delete(id); err != nil {
			h.JErr(c, err.Error())
		} else {
			h.JDone(c)
		}
	})
	r.GET("/api/search", func(c *gin.Context) {
		var sub string
		if c.BindJSON(&sub) != nil {
			return
		}
		res := db.Search(sub)
		h.JOk(c, res)
	})
	r.Run("127.0.0.1:8080")
}
