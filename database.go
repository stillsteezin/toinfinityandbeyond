package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	// Import postgres driver.
	_ "github.com/lib/pq"
)

var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s <db URL> <mountpoint>\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		usage()
		os.Exit(2)
	}

	dbURL, mountPoint := flag.Arg(0), flag.Arg(1)

	// Open DB connection first.
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	if err := initSchema(db); err != nil {
		log.Fatal(err)
	}

	cfs := CFS{db}
	// Mount filesystem.
	c, err := fuse.Mount(
		mountPoint,
		fuse.FSName("CockroachFS"),
		fuse.Subtype("CockroachFS"),
		fuse.LocalVolume(),
		fuse.VolumeName(""),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = c.Close()
	}()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		for range sig {
			if err := fuse.Unmount(mountPoint); err != nil {
				log.Printf("Signal received, but could not unmount: %s", err)
			} else {
				break
			}
		}
	}()

	// Serve root.
	err = fs.Serve(c, cfs)
	if err != nil {
		log.Fatal(err)
	}

	// check if the mount process has an error to report
	<-c.Ready
	if err := c.MountError; err != nil {
		log.Fatal(err)
	}
}
