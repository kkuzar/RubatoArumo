package analyse

import (
	"github.com/huzhaer/qianxun/lib/flight"
	//"net/http"
	"github.com/huzhaer/teamlite_core/router"

	"bytes"
	"fmt"
	"github.com/huzhaer/qianxun/model/users"
	_ "github.com/liudng/godump"
	"net/http"
)


// Load the routes.
func Load() {
	router.Get("/", Testing)
}
//
func Testing(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)


	users.Test(c.DB)

	Views := c.View
	Views.SetDevelopmentMode(true)
	view, err := Views.GetTemplate("base.jet")

	var resp bytes.Buffer
	if err = view.Execute(&resp, nil, nil); err != nil {
		w.WriteHeader(503)
		fmt.Fprintf(w, "Error when executing template: %+v", err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write(resp.Bytes())

	// fmt.Fprint(w,"Testing!!!!")
	return
}