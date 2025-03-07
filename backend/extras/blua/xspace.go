package blua

import (
	"database/sql"

	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/rs/zerolog"
)

/*

ROUTES:
 /z/startup/
 /z/preload/spaces/<space_type>/
 /z/pages/startup/
 /z/pages/portal/spaces/<space_type>/
 /z/spaces/<space_type>/
 /z/api/spaces/<space_type>/
TABLE:
 zSpaceXXX
 zspace_pid_xxx
 zBuddyXXX
 zbuddy_bid_xxx
 zCaptureXXX
 zcapture_chash_xxx
 zLogXXX
WORKING_DIR:
 /workin_space/<space_type>


*/

type BuilderOption struct {
	App xtypes.App
	DB  *database.DB
	SQL SQL

	Logger zerolog.Logger
}

type SQL interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) (*sql.Row, error)
	RunDDL(query string) error
	RunMigration(pid int64, name string, query string) error
	ListMigration(pid int64) ([]string, error)
	RemoveMigration(pid int64, name string) error

	// Transaction
	Begin() (*sql.Tx, error)
}

type Builder func(opt BuilderOption) (*XSpaceType, error)

type XSpaceField struct {
	Name       string   `json:"name"`
	KeyName    string   `json:"key_name"`
	Ftype      string   `json:"ftype"`
	Disabled   bool     `json:"disabled"`
	Options    []string `json:"options"`
	IsRequired bool     `json:"is_required"`
}

type XSpaceType struct {
	Name    string
	Slug    string
	Info    string
	Version string

	Fields         []XSpaceField
	EventHandlers  map[string]any
	AssetsProvider any
}
