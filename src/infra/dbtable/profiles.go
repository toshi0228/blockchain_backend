package dbtable

import (
	_ "embed"
	"fmt"
)

const TableNameProfile = "profiles"

//go:embed profiles.sql
var profilesTableSql string

func Profiles() string {

	cmd := fmt.Sprintf(profilesTableSql, TableNameProfile)
	return cmd
}
