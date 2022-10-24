package pageParser

import (
	"BIGCrawler/src"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"time"
)

const URL string = "https://ngdc.cncb.ac.cn/gwh/api/browse"

type FromTo struct {
	Start int
	End   int
}

func DivideInto4(n int, fromToChan chan FromTo) {
	var ret []FromTo
	var startEnd FromTo
	var unit = n / 4

	for i := 0; ; {
		if i-unit >= 0 {
			unit = i
			break
		}
		i += 100
	}

	for i, j := 0, unit; i < n; i, j = i+unit, j+unit {
		//fmt.Println(i)
		startEnd = FromTo{0, 0}
		if i != 0 {
			startEnd.Start = i + 100
		} else {
			startEnd.Start = i
		}
		startEnd.End = j
		ret = append(ret, startEnd)
	}
	if ret[len(ret)-1].End == 0 || ret[len(ret)-1].End >= n {
		ret[len(ret)-1].End = n
	}
	for _, v := range ret {
		fromToChan <- v
	}
	close(fromToChan)
}

func GetTotalPage() (totalPage int) {
	var gwh = src.GwhResponse{}
	var reqJson = src.RequstJson(0, 0, 10)
	var jsonStr = []byte(reqJson)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json")
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Wrong in get total page:", err)
		os.Exit(1)
	})
	c.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body, &gwh)
		if err != nil {
			fmt.Println("Wrong in get total page:", err)
			os.Exit(1)
		}
		totalPage = gwh.RecordsTotal
	})
	c.PostRaw(URL, jsonStr)
	return totalPage
}

func AccPageParser(fromtoChan chan FromTo, infoChan chan src.NeedInfo, exitChan chan bool) {

	c := colly.NewCollector()
	var gwh src.GwhResponse
	var info src.NeedInfo
	//var infos []src.NeedInfo
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json")
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("something went wrong", err)
	})
	c.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body, &gwh)
		if err != nil {
			fmt.Println("response went wrong", err)
		}
		//fmt.Println(string(response.Body))
		for _, v := range gwh.Data {
			info = src.NeedInfo{}
			info.Primayid = v.PrimaryId
			info.Accession = v.Accession
			info.ReleaseDate = v.ReleaseDate
			info.AssemblyLevel = v.AssemblyLevel
			info.ScientificName = v.ScientificName
			BaseDown := fmt.Sprintf("https://download.cncb.ac.cn/gwh/%s/", v.FtpDir)
			if v.Rna != "" {
				info.RNAFile = BaseDown + v.Rna
			}
			if v.Gff != "" {
				info.GFFFile = BaseDown + v.Gff
			}
			if v.Dna != "" {
				info.DNAFile = BaseDown + v.Dna
			}
			if v.Protein != "" {
				info.ProteinFile = BaseDown + v.Protein
			}
			info.GenomeRepresentation = v.GenomeRepresentation
			//infos = append(infos, info)
			infoChan <- info
		}
	})
	var start int
	var end int
	for {
		temp, ok := <-fromtoChan
		if !ok {
			break
		}
		start = temp.Start
		end = temp.End
		for ; start <= end; start += 100 {
			Json := src.RequstJson(0, start, 100)
			jsonStr := []byte(Json)
			c.PostRaw(URL, jsonStr)

		}
	}

	exitChan <- true
}
func GWHCrawler(path string) {
	fromToChan := make(chan FromTo, 4)
	infoChan := make(chan src.NeedInfo, 1000)
	exitChan := make(chan bool, 4)
	start := time.Now()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	go DivideInto4(GetTotalPage(), fromToChan)
	for i := 0; i < 4; i++ {
		go AccPageParser(fromToChan, infoChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(infoChan)
	}()

	for {
		info, ok := <-infoChan
		if !ok {
			break
		}
		str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", info.Accession, info.ScientificName,
			info.ReleaseDate, info.AssemblyLevel, info.GenomeRepresentation,
			info.DNAFile, info.GFFFile, info.RNAFile, info.ProteinFile)
		file.Write([]byte(str))
	}
	fmt.Println("Done! 用时: ", time.Since(start))
}
