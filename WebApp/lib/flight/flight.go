// Package flight provides access to the application settings safely.
package flight

import (
	"fmt"
	"github.com/huzhaer/qianxun/core/passhash"
	"github.com/huzhaer/qianxun/core/session"
	"github.com/huzhaer/qianxun/lib/utils"
	"github.com/huzhaer/qianxun/model/users"
	"github.com/huzhaer/qianxun/services/users"
	"log"
	"net/http"
	"sync"

	"github.com/huzhaer/qianxun/lib/env"

	"github.com/CloudyKit/jet"
	"github.com/CloudyKit/jet/loaders/multi"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/huzhaer/teamlite_core/flash"
	"github.com/huzhaer/teamlite_core/form"
	"github.com/huzhaer/teamlite_core/router"
	"github.com/jmoiron/sqlx"
)

var (
	configInfo env.Info
	//dbInfo     interface{}
	dbInfo     *sqlx.DB
	pgInfo     *sqlx.DB
	mutex 		sync.RWMutex
)

var ViewsVar = jet.NewHTMLSetLoader(multi.NewLoader(
	jet.NewOSFileSystemLoader("./views"),
	// httpfs.NewLoader(templates.Assets),
))


// StoreConfig stores the application settings so controller functions can
//access them safely.
func StoreConfig(ci env.Info) {
	mutex.Lock()
	configInfo = ci
	mutex.Unlock()
}

// StoreDB stores the database connection settings so controller functions can
// access them safely.
func StoreDB(db *sqlx.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

//func StoreDB(db interface{}) {
//	mutex.Lock()
//	dbInfo = db
//	mutex.Unlock()
//}

func StorePqsqlPoolConnection (conn *sqlx.DB) {
	mutex.Lock()
	pgInfo = conn
	mutex.Unlock()
}

// Info structures the application settings.
type Info struct {
	Config env.Info
	Sess   *sessions.Session
	UserID string
	W      http.ResponseWriter
	R      *http.Request
	//DB     interface{}
	DB     *sqlx.DB
	PgDB   *sqlx.DB
	View  *jet.Set
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {
	var id string

	// Get the session
	sess, err := configInfo.Session.Instance(r)

	// If the session is valid
	if err == nil {
		// Get the user id
		id = fmt.Sprintf("%v", sess.Values["id"])
	}

	// Setting up the CSRF header
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	mutex.RLock()
	i := Info{
		Config: configInfo,
		Sess:   sess,
		UserID: id,
		W:      w,
		R:      r,
		DB:     dbInfo,
		PgDB:   pgInfo,
		View:   ViewsVar,
	}
	mutex.RUnlock()

	return i
}

// Reset will delete all package globals
func Reset() {
	mutex.Lock()
	configInfo = env.Info{}
	dbInfo = &sqlx.DB{}
	pgInfo = &sqlx.DB{}
	mutex.Unlock()
}

// Param gets the URL parameter.
func (c *Info) Param(name string) string {
	return router.Param(c.R, name)
}

// Redirect sends a temporary redirect.
func (c *Info) Redirect(urlStr string) {
	http.Redirect(c.W, c.R, urlStr, http.StatusFound)
}

// FormValid determines if the user submitted all the required fields and then
// saves an error flash. Returns true if form is valid.
func (c *Info) FormValid(fields ...string) bool {
	if valid, missingField := form.Required(c.R, fields...); !valid {
		c.Sess.AddFlash(flash.Info{"Field missing: " + missingField, flash.Warning})
		c.Sess.Save(c.R, c.W)
		return false
	}

	return true
}

// Repopulate fills the forms on the page after the user submits.
func (c *Info) Repopulate(v map[string]interface{}, fields ...string) {
	form.Repopulate(c.R.Form, v, fields...)
}

// FlashSuccess saves a success flash.
func (c *Info) FlashSuccess(message string) {
	c.Sess.AddFlash(flash.Info{message, flash.Success})
	c.Sess.Save(c.R, c.W)
}

// FlashNotice saves a notice flash.
func (c *Info) FlashNotice(message string) {
	c.Sess.AddFlash(flash.Info{message, flash.Notice})
	c.Sess.Save(c.R, c.W)
}

// FlashWarning saves a warning flash.
func (c *Info) FlashWarning(message string) {
	c.Sess.AddFlash(flash.Info{message, flash.Warning})
	c.Sess.Save(c.R, c.W)
}

// FlashError saves an error flash and logs the error.
func (c *Info) FlashError(err error) {
	log.Println(err)
	c.Sess.AddFlash(flash.Info{err.Error(), flash.Error})
	c.Sess.Save(c.R, c.W)
}

// FlashErrorGeneric saves a generic error flash and logs the error.
func (c *Info) FlashErrorGeneric(err error) {
	log.Println(err)
	c.Sess.AddFlash(flash.Info{"An error occurred on the server. Please try again later.", flash.Error})
	c.Sess.Save(c.R, c.W)
}

// Auth Functions
func (c *Info) LoginAttempt (w http.ResponseWriter, r *http.Request) (res bool ) {

	userInfo := utils.ProcessPostByDict(r, users.UserAuthDict)
	username, _ := userInfo["username"]
	password, _ := userInfo["password"]

	tUser := users_service.GetUserWithUserName(username, c.DB)

	if (passhash.MatchString(tUser.Password, password)) {
		c.Sess.Values["isLogin"]  = true
		c.Sess.Values["username"] = tUser.Username
		c.Sess.Values["email"]    = tUser.Email
		c.Sess.Values["mobile"]   = tUser.Mobile

		c.Sess.Save( r, w)
		return true
	}

	return false
}

func (c *Info) Logout (w http.ResponseWriter, r *http.Request) {
	session.Empty(c.Sess)
}