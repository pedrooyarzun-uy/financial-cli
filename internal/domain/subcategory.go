package domain

type Subcategory struct {
	Id         int
	Name       string
	CategoryId int
	UserId     *int
}
