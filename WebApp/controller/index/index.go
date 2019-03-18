package index

import (
	"encoding/json"
	"github.com/CloudyKit/jet"
	"github.com/huzhaer/qianxun/lib/flight"
	"github.com/liudng/godump"
	"go.bug.st/serial.v1"
	"strconv"

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
	router.Get("/", Index)
	router.Get("/portlist", PortList)
	router.Post("/motor",Motor)
	router.Get("/led", Led)
}
//
func Testing(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)

	_ = c

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

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)

	ViewsObj := c.View
	ViewsObj.SetDevelopmentMode(true)

	tVars := make(jet.VarMap)
	tVars.Set("title","Robot!")

	thisView, err := ViewsObj.GetTemplate("pages/index.jet")
	var resp bytes.Buffer
	if err = thisView.Execute(&resp, tVars, nil); err != nil {
		w.WriteHeader(503)
		fmt.Fprintf(w, "Error when executing template: %+v", err.Error())
		return
	}
	w.WriteHeader(200)
	_ , _ = w.Write(resp.Bytes())

	return
}

func Motor(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)

	_ = c

    num := r.FormValue("num")
    pos := r.FormValue("pos")
    ori := r.FormValue("ori")

    mark := "0000" + ori + "0" + num + pos

	mode := &serial.Mode{
		BaudRate: 115200,
	}
	portConn, _ := serial.Open("COM7", mode)
	_, err := portConn.Write([]byte(mark))

	buff := make([]byte, 4)

	n, err := portConn.Read(buff)

	godump.Dump(err);
	godump.Dump(mark);
	godump.Dump(string(buff[:]));
	godump.Dump(n);


	_ = portConn.Close()

	return
}

func PortList(w http.ResponseWriter, r *http.Request) {
	c := flight.Context( w, r)

	_ = c

	Ports, err := serial.GetPortsList()
	_ = err

	js, err := json.Marshal(Ports)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ , _  = w.Write(js)

	godump.Dump(Ports)
	return
}

func Led(w http.ResponseWriter, r *http.Request) {

	state ,_ :=r.URL.Query()["state"]


	stateInt, _ := strconv.ParseInt(state[0], 0, 64);
	godump.Dump(stateInt)
	var ledStr string
	if (stateInt == 1) {
		ledStr = "0000O000"
	}
	if (stateInt == 0) {
		ledStr = "0000F000"
	}

	mode := &serial.Mode{
		BaudRate: 115200,
	}

	portConn, er := serial.Open("COM7", mode)
	godump.Dump(er);

	n, _ := portConn.Write([]byte(ledStr))

	fmt.Printf("Sent %v bytes\n", n)

	_ = portConn.Close()

	return
}