package form

/*

Basic form
*/

type ProductForm struct {
	Name  string `form:"name" binding:"Required"`
	Email string `form:"email"`
}
