// Package env reads the application settings.
package env

import (
	"encoding/json"

	"github.com/huzhaer/teamlite_core/asset"
	"github.com/huzhaer/teamlite_core/email"
	"github.com/huzhaer/teamlite_core/form"
	"github.com/huzhaer/teamlite_core/generate"
	"github.com/huzhaer/teamlite_core/jsonconfig"
	"github.com/huzhaer/teamlite_core/storage/driver/mysql"
	"github.com/huzhaer/teamlite_core/view"

	"github.com/huzhaer/qianxun/core/server"
	"github.com/huzhaer/qianxun/core/session"

	"github.com/huzhaer/qianxun/lib/pgsql"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

// Info structures the application settings.
type Info struct {
	Asset      asset.Info    `json:"Asset"`
	Email      email.Info    `json:"Email"`
	Form       form.Info     `json:"Form"`
	Generation generate.Info `json:"Generation"`
	MySQL      mysql.Info    `json:"MySQL"`
	PgSQL      pgsql.Info    `json:"Postgres"`
	Server     server.Info   `json:"Server"`
	Session    session.Info  `json:"Session"`
	Template   view.Template `json:"Template"`
	View       view.Info     `json:"View"`
	path       string
}

// Path returns the env.json path
func (c *Info) Path() string {
	return c.path
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// New returns a instance of the application settings.
func New(path string) *Info {
	return &Info{
		path: path,
	}
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Create a new configuration with the path to the file
	config := New(configFile)

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}
