package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo"

	"crud-echo/models"
)


type Token struct{
	Token string `json:"Token"`
}

var customer = &models.Customer{
	Id   :  1,
	Name : "Thomas",
	Email: "thomas@gmail.com",
	Alamat: "jln. amarta raya blok ce no.5",
	Nohp  : +6281218573272,
}

const SESSION_ID = "id"

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgres:123456@127.0.0.1:5433/universal?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}


func Handler(c echo.Context) error {

    store := newPostgresStore()
    session, _ := store.Get(c.Request(), SESSION_ID)
    
    if len(session.Values) == 0 {
        return c.String(http.StatusOK, "empty result")
    }
    sess := &models.Session{
        Session : session.Values["session"].(string),
    }

    sesValues := sess

    var customerList = &models.ListCustomer{
        List :  []*models.Customer{
            customer,
        },
        Session: []*models.Session{
            sesValues,
        },
    }

    return c.JSON(http.StatusOK, customerList)
}



func HandlerSet(c echo.Context) error {

    store := newPostgresStore()
    
    session, _ := store.Get(c.Request(), SESSION_ID)
    session.Values["session"] = "asd session"
    sess := &models.Session{
        Session : session.Values["session"].(string),
    }

    sesValues := sess
    
    session.Save(c.Request(), c.Response())
    
    var result = &models.ListCustomer{
        Session :  []*models.Session{
            sesValues,
        },
    }

    return c.JSON(http.StatusOK, result)
}

func Handler2(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func HandlerDeleteSession(c echo.Context) error {
    store := newPostgresStore()

	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())

    return c.JSON(http.StatusOK, "Deleted Session")
}