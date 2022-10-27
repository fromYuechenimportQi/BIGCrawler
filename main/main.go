package main

import (
	"github.com/yueyue970506/BIGCrawler/"
	"github.com/yueyue970506/BIGCrawler/"
	"flag"
	"fmt"
	"os"
	"time"
)

//func main() {
//	file, err := os.OpenFile("gwh.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	//db, err := gorm.Open(sqlite.Open("gwh.db"), &gorm.Config{
//	//	//Logger: logger.Default.LogMode(logger.Silent),
//	//})
//	//if err != nil {
//	//	panic("failed to connect database")
//	//}
//	//db.AutoMigrate(&src.NeedInfo{})
//	end := pageParser.GetTotalPage()
//	for start := 0; start <= end; start += 100 {
//		infos := pageParser.AccPageParser(src.RequstJson(0, start, 100))
//		for _, info := range infos {
//			//var temp src.NeedInfo
//			//db.First(&temp, "primayid = ?", info.Primayid)
//			//if temp.Primayid == "" {
//			//} else {
//			//	fmt.Println(temp.Primayid)
//			//	continue
//			//}
//			//db.Create(&info)
//			str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", info.Accession, info.ScientificName,
//				info.ReleaseDate, info.AssemblyLevel, info.GenomeRepresentation,
//				info.DNAFile, info.GFFFile, info.RNAFile, info.ProteinFile)
//			file.Write([]byte(str))
//		}
//		//db.Create(&infos)
//	}
//}

var path = flag.String("out", "./gwh_result.txt", "Give me a filename, I will write result into file")
var gwh = flag.Bool("gwh", false, "Run GWH Crawler")
var gsa = flag.Bool("gsa", false, "Run GSA Crawler")

func main() {

	flag.Parse()
	if *path == "" {
		fmt.Println("Invalid file path")
		fmt.Println()
		fmt.Printf("Type %v -h for help", os.Args[0])
		os.Exit(1)
	}
	if !*gwh && !*gsa {
		fmt.Printf("Usage:\n%v -gwh or %v -gsa\n", os.Args[0], os.Args[0])
		fmt.Printf("Type %v -h for help\n", os.Args[0])
		fmt.Printf("\nEnter to quit.....")
		fmt.Scanln()
		os.Exit(0)
	}

	if *gsa && *gwh {
		panic("Can not select both!")
	} else if *gwh {
		go func() {
			for {
				for _, v := range `\|/` {
					fmt.Printf("\r%c", v)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
		pageParser.GWHCrawler(*path)
	} else if *gsa {
		gsa := GsaCrawler.GSA{}
		gsa.GSACrawler(*path)
	}

}
