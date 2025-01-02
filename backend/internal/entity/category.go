package entity

const TABLE_CATEGORIES = "categories"

type Category struct {
	Name string `json:"name" db:"name"`
}
