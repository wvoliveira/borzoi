package main

import "flag"

var (
	fMigrate = flag.Bool("migrate", false, "Enable GORM migration")
)
