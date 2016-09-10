package controllers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/modules/db/app"
	"github.com/revel/revel"
	"log"
	"mainApp/app/models"
	//"time"
	//"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

//set Global Dbm variable
var Dbm *gorp.DbMap

//func InitDB() *gorp.DbMap {
func InitDB() {
	//Initalize database, in db.spec location, determined by app conf
	db.Init()

	// construct a gorp DbMap
	Dbm := &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	//Set Databse sizes
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}

	}

	//add a table, setting the table name to 'users' and
	// specifying that the Id property is an auto incrementing Pk
	t := Dbm.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 20,
		"Name":     100,
	})

	//send sql logging to the logger
	Dbm.TraceOn("[gorp]", revel.INFO)

	err := Dbm.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln(err)
	}
	/*
		bcryptPassword, _ := bcrypt.GenerateFromPassword(
			[]byte("demo"), bcrypt.DefaultCost)
		demoUser := &models.User{5, 1000, "Demo User", "demo", "demo", bcryptPassword}
		if err := Dbm.Insert(demoUser); err != nil {
			panic(err)
		}*/
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}

}

func init() {
	revel.OnAppStart(InitDB)
}

type UserAuth struct {
	*revel.Controller
}

func (c UserAuth) Login() revel.Result {
	return c.Render()
}
