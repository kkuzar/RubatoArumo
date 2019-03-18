// Package login handles the user login.
package login

import (
	"bytes"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/huzhaer/qianxun/lib/flight"
	"github.com/huzhaer/teamlite_core/router"
	"net/http"

	"math/rand"
	"time"
)

// Load the routes.
func Load() {
	router.Post("/login", LoginPost)
	router.Get("/login", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)

	ViewsObj := c.View
	ViewsObj.SetDevelopmentMode(true)

	tVars := make(jet.VarMap)
	tVars.Set("title","Template Testing!")

	thisView, err := ViewsObj.GetTemplate("pages/admin_login.jet")
	var resp bytes.Buffer
	if err = thisView.Execute(&resp, tVars, nil); err != nil {
		w.WriteHeader(503)
		fmt.Fprintf(w, "Error when executing template: %+v", err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write(resp.Bytes())
}

func LoginPost (w http.ResponseWriter, r *http.Request)  {
	c := flight.Context( w, r)

	c.LoginAttempt(w , r)
	_ = c
}
//
//func Login(w http.ResponseWriter, r *http.Request){
//	c := flight.Context( w, r)
//
//	res := utils.ProcessPostByDict(r, usersservice.UserEntity)
//	username := res["username"]
//	password := res["password"]
//
//	t_User, u_err := usersservice.GetUserByUsername(c.PgDB, username)
//
//	if u_err != nil {
//		fmt.Println(u_err)
//	}
//
//	if (passhash.MatchString( t_User.Password , password)) {
//		// When the password did match
//		fmt.Println("Success!")
//		c.Sess.Values["isLogin"]  = true
//		c.Sess.Values["username"] = t_User.Username
//		c.Sess.Values["email"]    = t_User.Email
//		c.Sess.Values["mobile"]   = t_User.Mobile
//
//		godump.Dump(c.Sess)
//
//		jsonData, _ := json.Marshal(res)
//		resultJ := map[string]string {
//			"status" :"OK",
//		}
//		resultJson,_ := json.Marshal(resultJ)
//		w.Header().Set("Content-Type","application/json")
//		fmt.Fprint(w,string(jsonData[:]))
//		fmt.Fprint(w,string(resultJson[:]))
//		return
//	}
//
//	fail := map[string]string {
//		"status": "False",
//	}
//	jsonData, _ := json.Marshal(fail)
//	w.Header().Set("Content-Type","application/json")
//	fmt.Fprint(w,string(jsonData[:]))
//	return
//}


//func TestLogin (w http.ResponseWriter, r *http.Request)  {
//	 c := flight.Context(w,r)
//	//var netClient = &http.Client{
//	//	Timeout: time.Millisecond * 800,
//	//}
//	//res, geterr := netClient.Get("http://localhost:3000/login")
//	//
//	//if geterr == nil {
//	//	 defer res.Body.Close()
//	//	 body, _ := ioutil.ReadAll(res.Body)
//	//	 w.Write(body)
//	//
//	//	 return
//	//}
//	//fmt.Println(geterr)
//	//http.Redirect(w, r, "/", http.StatusNotFound)
//	users.Test(c.PgDB)
//
//}

// Index displays the login page.
//func Index(w http.ResponseWriter, r *http.Request) {
//	 c := flight.Context(w, r)
//	//
//	//v := c.View.New("login/index")
//	//form.Repopulate(r.Form, v.Vars, "email")
//	//v.Render(w, r)
//
//	u := map[string]string {
//		"username" : "Csrftest",
//		"family_name" : RandStringRunes(4),
//		"first_name" : RandStringRunes(6),
//		"password" : "huzhaer123456",
//		"mobile" : "213543216",
//		"email" : RandStringRunes(5)+ "@"+ RandStringRunes(4)+".com",
//		"role" : "",
//		"country_id" : "",
//		"city_id" : "",
//		"created_at" : "",
//		"updated_at" : "",
//		"last_login" : "",
//		"remark" : "",
//		"isActive" : "",
//		"isDeleted" : "",
//		"isAbnormal" : "",
//		"hash_token" : "",
//		"token" : "",
//		"old_token" : "",
//	}
//
//	res, err :=usersservice.CreateUser(c.PgDB, u)
//	godump.Dump(err)
//	println(res)

	// res, err := users.CreateUser(c.PgDB,"hank"+RandStringRunes(3),RandStringRunes(16),"123@"+RandStringRunes(6)+".com")

	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	//
	//js, err := json.Marshal(profile)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(js)
//}


func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

