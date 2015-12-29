package products

import (
	"github.com/tixu/productportfolio/form"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ADDPRODUCT = "products/addproduct"
const LISTPRODUCTS = "products/listproducts"

type Product struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Responsible string        `bson:"responsible_name"`
	Category    string        `bson:"category_name"`
}

func Add(ctx *macaron.Context, db *mgo.Database) {
	ctx.HTML(200, ADDPRODUCT)
}

func DeleteAll(ctx *macaron.Context, db *mgo.Database) {
	productsCollection := db.C("products")
	_, err := productsCollection.RemoveAll(nil)
	if err != nil {
		panic(err)
	}
}

func ListAll(ctx *macaron.Context, db *mgo.Database) {

	productsCollection := db.C("products")
	var results []Product
	err := productsCollection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	ctx.Data["Products"] = results
	ctx.HTML(200, LISTPRODUCTS)
}

func Delete(ctx *macaron.Context, db *mgo.Database) {
	productsCollection := db.C("products")
	id := ctx.Params("name")

	err := productsCollection.Remove(bson.M{"name": id})
	if err != nil {
		panic(err)
	}

	ListAll(ctx, db)
}
func Submit(productf form.ProductForm, ctx *macaron.Context, db *mgo.Database) {
	product := Product{Name: productf.Name, Responsible: productf.Email}
	productsCollection := db.C("products")
	err := productsCollection.Insert(&product)
	if err != nil {
		panic(err)
	}
	ListAll(ctx, db)

}
