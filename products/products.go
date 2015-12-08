package products

import (
	"fmt"

	"github.com/tixu/productportfolio/form"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Responsible string        `bson:"responsible"`
}

func Delete(ctx *macaron.Context, db *mgo.Database) {
	productsCollection := db.C("products")
	_, err := productsCollection.RemoveAll(nil)
	if err != nil {
		panic(err)
	}
}
func ListAll(ctx *macaron.Context, db *mgo.Database) {
	ctx.Data["title"] = "journal"
	ctx.HTML(200, "product")
	productsCollection := db.C("products")
	var results []Product
	err := productsCollection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

}

func Submit(productf form.ProductForm, ctx *macaron.Context, db *mgo.Database) {
	product := Product{Name: productf.Name, Responsible: productf.Email}
	productsCollection := db.C("products")
	err := productsCollection.Insert(&product)
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
}
