package main

import (
	"fmt"

	"github.com/go-macaron/binding"
	"github.com/tixu/productportfolio/categories"
	"github.com/tixu/productportfolio/form"
	"github.com/tixu/productportfolio/products"
	"github.com/tixu/productportfolio/server"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
)

func main() {
	session := connecttoDB("127.0.0.1/poducts")
	fmt.Println(session)
	// functions used in the  route
	m := server.GetServer()
	m.Use(func(ctx *macaron.Context) {
		//****
		//cloning the session to connect to the DB in the different route.
		//*****
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
		m.Get("/deleteAll", products.DeleteAll)
		m.Get("/", products.ListAll)
		m.Get("/add", products.Add)
		m.Get("/delete/:name", products.Delete)
		m.Post("/submit", binding.Bind(form.ProductForm{}), products.Submit)
	})

	m.Group("/category", func() {
		m.Get("/add", categories.Add)
		m.Get("/", categories.ListAll)
		m.Post("/submit", binding.Bind(form.CategoryForm{}), categories.Submit)
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

// func getSession(ctx *macaron.Context, session *mgo.Session) {
//****
//cloning the session to connect to the DB in the different route.
//*****
//	reqsession := session.Clone()
//	ctx.Map(reqsession.DB("product"))
//	defer reqsession.Close()
//	ctx.Next()
//}
