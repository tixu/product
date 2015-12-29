package categories

import (
	"fmt"

	"github.com/tixu/productportfolio/form"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ADDCATEGORY = "categories/addcategory"
const LISTCATEGORIES = "categories/listcategories"

type Category struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	PID         string        `bson:"pid"`
	Description string        `bson:"description"`
}

func Add(ctx *macaron.Context, db *mgo.Database) {
	var results []Category
	results, err := getAll(db)
	if err != nil {
		panic(err)
	}
	ctx.Data["Categories"] = results
	ctx.HTML(200, ADDCATEGORY)
}

func Submit(categoryf form.CategoryForm, ctx *macaron.Context, db *mgo.Database) {
	category := Category{Name: categoryf.Name, PID: categoryf.Parentid, Description: categoryf.Description}
	categoriesCollection := db.C("categories")
	err := categoriesCollection.Insert(&category)
	if err != nil {
		panic(err)
	}
	ListAll(ctx, db)
}

func ListAll(ctx *macaron.Context, db *mgo.Database) {
	var results []Category
	results, err := getAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Categories collection size %s \n ", results)
	ctx.Data["Categories"] = results
	ctx.HTML(200, LISTCATEGORIES)
}

func getAll(db *mgo.Database) (results []Category, err error) {
	categoriesCollection := db.C("categories")
	err = categoriesCollection.Find(nil).All(&results)
	return results, err
}
