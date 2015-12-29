package form

/*

Basic form
*/

type ProductForm struct {
	Name  string `form:"name" binding:"Required"`
	Email string `form:"email"`
}

type CategoryForm struct {
	Name        string `form:"name" binding:"Required"`
	Parentid    string `form:"parentid" binding:"Required"`
	Description string `form:"description" binding:"Required"`
}
