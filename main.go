package main

import (
	"flag"
	"fmt"
	"github.com/pedroalbanese/beeep"
	"github.com/pedroalbanese/dlgs"
	"log"
	"os"
	"time"
)

var (
	beep     = flag.Bool("beep", false, "Beep alarm.")
	date     = flag.Bool("date", false, "Date selection dialog box.")
	erro     = flag.Bool("error", false, "Error dialog box.")
	file     = flag.Bool("file", false, "File selection dialog box.")
	folder   = flag.Bool("folder", false, "Folder selection dialog box.")
	info     = flag.Bool("info", false, "Info dialog box.")
	input    = flag.Bool("input", false, "Text input box.")
	password = flag.Bool("pass", false, "Password input box.")
	question = flag.Bool("quest", false, "Question dialog box. (Check the Exit code)")
	subtitle = flag.String("sub", "Subtitle", "Box subtitle.")
	title    = flag.String("title", "Title", "Box title.")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *info == true {
		_, err := dlgs.Info(*title, *subtitle)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *erro == true {
		_, err := dlgs.Error(*title, *subtitle)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *input == true {
		msg, _, err := dlgs.Entry(*title, *subtitle, "")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}

	if *password == true {
		passwd, _, err := dlgs.Password(*title, *subtitle+": ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(passwd)
	}

	if *question == true {
		qst, err := dlgs.Question(*title, *subtitle, true)
		if err != nil {
			log.Fatal(err)
		}
		if qst == true {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if *folder == true {
		folder, _, err := dlgs.File("Select folder", "*", true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(folder)
	}

	if *file == true {
		file, _, err := dlgs.File("Select file", "*", false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(file)
	}

	if *date == true {
		dat, _, err := dlgs.Date(*title, *subtitle, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dat.Format("2006-01-02"))
	}

	if *beep == true {
		_ = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	}
}
