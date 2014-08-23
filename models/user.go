package models

import (
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type User struct {
	ID      string
	Email   string
	Name    string
	Created time.Time
}

func Create(r *http.Request, email string) {

	context := appengine.NewContext(r)
	user := User{
		Email:   email,
		Created: time.Now(),
	}
}

func (u *User) Read(r *http.Request, email string) (User, errors) {

	c := appengine.NewContext(r)
	q := datastore.NewQuery("User").
		Filter("Email =", email).
		Limit(1)

	var usr []User

	_, err := q.GetAll(c, &usr)

	if err != nil {
		return nil, err
	}

	if len(usr) == 0 {
		return nil, nil
	} else {
		return usr[0], nil
	}
}

func SignIn(w *http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {

		url, err := user.LoginURL(c, r.URL.String())

		if err != nil {
			return nil, err
		}

		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	account, err := Read(r, u.Email)

	if err != nil {
		return nil, err
	}

	if account == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	return account, nil
}
