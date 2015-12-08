package main

import (
	"github.com/go-macaron/binding"
	"github.com/tixu/productportfolio/form"
	"github.com/tixu/productportfolio/products"
	"github.com/tixu/productportfolio/server"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
)

func main() {
	session := connecttoDB("127.0.0.1/poducts")
	// functions used in the  route
	m := server.GetServer()
	m.Use(func(ctx *macaron.Context) {

		reqsession := session.Clone()

		ctx.Map(reqsession.DB("product"))
		defer reqsession.Close()
		ctx.Next()
	})
	m.Get("/index", func(ctx *macaron.Context) {

		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Get("/hello", func(ctx *macaron.Context) {

		ctx.HTML(200, "hello") // 200 is the response code.
	})

	// grouping all the route

	m.Group("/product", func() {
		m.Get("/delete", products.Delete)
		m.Get("/", products.ListAll)
		m.Post("/submit", binding.Bind(form.ProductForm{}), products.Submit)
	})
	m.Run()
}

func connecttoDB(dbname string) *mgo.Session {
	session, err := mgo.Dial("127.0.0.1/product")
	// Optional. Switch the session to a monotonic behavior.
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}
