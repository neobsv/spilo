package docs

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Import sqllite gorm
)

// DB is the database connection to read data
var DB *gorm.DB

// MasterDB is the database connection to write data
var MasterDB *gorm.DB

// Docs is a struct which holds documentation
// for a website in an ORM model.
type Docs struct {
	gorm.Model
	Identifier string `json:"id"`
	Section    string `json:"section"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

// Doc is an unpacked representation of the gorm model
type Doc struct {
	ID         int64  `json:"ID"`
	CreatedAt  string `json:"CreatedAt"`
	UpdatedAt  string `json:"UpdatedAt"`
	DeletedAt  string `json:"DeletedAt"`
	Identifier string `json:"id"`
	Section    string `json:"section"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}

// DeleteDoc removes a document from the database.
func DeleteDoc(c *fiber.Ctx) {
	var docs []Docs
	id := c.Params("id")

	dbc := DB.Find(&docs, Docs{Identifier: id})
	if dbc.Error != nil {
		fmt.Println("Error! could not find docs " + dbc.Error.Error())
		c.Send("Error: Could not find docs")
		return
	}

	for _, document := range docs {
		dbc := MasterDB.Delete(&document)
		if dbc.Error != nil {
			fmt.Println("Error! could not delete docs " + dbc.Error.Error())
			c.Send("Error: Could not delete docs")
			return
		}
	}

	c.JSON(docs)
}

// GetDocs uses the ORM to retrieve all the documents in the database.
func GetDocs(c *fiber.Ctx) {
	var docs []Docs
	dbc := DB.Find(&docs)
	if dbc.Error != nil {
		fmt.Println("Error! could not read database " + dbc.Error.Error())
		c.Send("Error! 404")
		return
	}

	c.JSON(docs)
}

// GetDoc retrieves a single document from the database.
func GetDoc(c *fiber.Ctx) {
	var doc []Docs
	id := c.Params("id")
	dbc := DB.Find(&doc, Docs{Identifier: id})
	if dbc.Error != nil {
		fmt.Println("Error! could not read database " + dbc.Error.Error())
		c.Send("Error! 404")
		return
	}

	c.JSON(doc)
}

// PostDoc creates a new document in the database.
func PostDoc(c *fiber.Ctx) {
	doc := new(Docs)

	err := c.BodyParser(doc)
	if err != nil {
		fmt.Println("Error! could not parse request")
		return
	}

	dbc := MasterDB.Create(&doc)
	if dbc.Error != nil {
		fmt.Println("Error! could not create record " + dbc.Error.Error())
		c.Send("Error! 400")
		return
	}

	c.JSON(doc)
}

// TestDoc inserts test data into the database
func TestDoc(c *fiber.Ctx) {
	dbc0 := MasterDB.Create(&Docs{Identifier: "0", Section: "1.2", Title: "Sub1", Text: "This is a paragraph. A wall of text with no real meaning, possibly beacause this is part of a placeholder. The actual contents of this page will be described in the following contents. We have a splendid API as well which we will be describing in a separate page. I need to type more things into this box, to fill it up. A lot of information will be provided, in order to allow people to understand what exactly we are trying to do and how we will be able to provide those services."})
	if dbc0.Error != nil {
		fmt.Println("Error! could not create record " + dbc0.Error.Error())
		c.Send("Error! 400")
		return
	}

	dbc1 := MasterDB.Create(&Docs{Identifier: "1", Section: "1.2", Title: "Sub1", Text: "This is a paragraph. A wall of text with no real meaning, possibly beacause this is part of a placeholder. The actual contents of this page will be described in the following contents. We have a splendid API as well which we will be describing in a separate page. I need to type more things into this box, to fill it up. A lot of information will be provided, in order to allow people to understand what exactly we are trying to do and how we will be able to provide those services."})
	if dbc1.Error != nil {
		fmt.Println("Error! could not create record " + dbc1.Error.Error())
		c.Send("Error! 400")
		return
	}

	c.Send("Two docs inserted successfully.")
}

// TestServer returns a string without hitting the database
func TestServer(c *fiber.Ctx) {
	c.Send("Hello world!")
}

// TestJSON returns a json without hitting the database
func TestJSON(c *fiber.Ctx) {
	var ds []*Doc
	d1 := Doc{
		ID:         1,
		CreatedAt:  "2020-06-03T01:53:53.439138Z",
		UpdatedAt:  "2020-06-03T01:53:53.439138Z",
		DeletedAt:  "null",
		Identifier: "0",
		Section:    "1.1",
		Title:      "Sub1",
		Text:       "This is a paragraph. A wall of text with no real meaning, possibly beacause this is part of a placeholder. The actual contents of this page will be described in the following contents. We have a splendid API as well which we will be describing in a separate page. I need to type more things into this box, to fill it up. A lot of information will be provided, in order to allow people to understand what exactly we are trying to do and how we will be able to provide those services.",
	}
	d2 := Doc{
		ID:         2,
		CreatedAt:  "2020-06-03T01:53:53.464223Z",
		UpdatedAt:  "2020-06-03T01:53:53.464223Z",
		DeletedAt:  "null",
		Identifier: "1",
		Section:    "1.2",
		Title:      "Sub1",
		Text:       "This is a paragraph. A wall of text with no real meaning, possibly beacause this is part of a placeholder. The actual contents of this page will be described in the following contents. We have a splendid API as well which we will be describing in a separate page. I need to type more things into this box, to fill it up. A lot of information will be provided, in order to allow people to understand what exactly we are trying to do and how we will be able to provide those services.",
	}

	ds = append(ds, &d1, &d2)

	c.JSON(ds)
}
