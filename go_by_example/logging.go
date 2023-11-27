package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}

/*
electornic data processing (unit1)
logical vs physical file (unit 1)
databse vs file system (unit 2)
ER diagram (unit 3)
realtionship (unit 3)
SQL , select , from (unit 4)
entity set (unitunit 4)
intregrity constraints ( unit 5)
Third Normal form (3NF) vs Boyce-Codd Normal Form (BCNF) (unit 5)
*/
