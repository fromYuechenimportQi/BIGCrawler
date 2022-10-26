package GsaCrawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var SearchURL string = "https://ngdc.cncb.ac.cn/gsa/browse/"

var DownloadPath string

func GetPageURL(postfix string) string {
	return fmt.Sprintf("%s%s", SearchURL, postfix)
}

func (this *GSA) GetTotalPage(url string) (TotalPage string) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `.\.gov`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  12,
	})
	if err != nil {
		log.Println(err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("get total page visiting:%s......\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Println("Something went wrong", err, r.StatusCode)
		}
	})

	c.OnResponse(func(r *colly.Response) {

	})

	c.OnHTML(".container", func(ele *colly.HTMLElement) {
		var cnt = 1
		ele.ForEachWithBreak("li.total", func(_ int, element *colly.HTMLElement) bool {
			str := strings.TrimSpace(element.Text)
			if cnt == 4 {
				TotalPage = strings.Split(str, "/")[1]

				return false
			}
			cnt += 1
			return true
		})
	})

	var requestData map[string]string
	requestData = make(map[string]string, 5)
	requestData["pageSize"] = "50"
	requestData["pageNo"] = "1"
	c.Post(url, requestData)
	return TotalPage
}

func (this *GSA) SearchPageParse(searchChan chan GSA) {
	TotalPage := this.GetTotalPage(SearchURL)
	var requestData map[string]string
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `.\.gov`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  12,
	})
	if err != nil {
		log.Println(err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("search page visiting:%spageNo=%s...\n", r.URL, requestData["pageNo"])
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Println("Something went wrong", err, r.StatusCode)
		}
	})

	c.OnResponse(func(r *colly.Response) {

	})

	c.OnHTML(".tab-content", func(ele *colly.HTMLElement) {
		ele.ForEach("td", func(_ int, elem *colly.HTMLElement) {
			if this.GSAAceesion != "" && this.BioProject != "" {
				this.GSAAceesion = ""
				this.BioProject = ""
			}
			str := strings.TrimSpace(elem.Text)
			if strings.HasPrefix(str, "CRA") {
				if this.GSAAceesion == "" {
					this.GSAAceesion = str
				}
			} else if strings.HasPrefix(str, "PRJ") {
				this.BioProject = str
				searchChan <- *this
			}

		})
	})
	requestData = make(map[string]string, 5)
	requestData["pageSize"] = "50"
	if TotalPage == "" {
		requestData["pageNo"] = "1"
		c.Post(SearchURL, requestData)
	} else {
		end, err := strconv.Atoi(TotalPage)
		fmt.Println(end)
		if err != nil {
			fmt.Println("Atoi wrong", err)
		}
		for i := 1; i <= end; i++ {
			requestData["pageNo"] = strconv.Itoa(i)
			c.Post(SearchURL, requestData)
		}
	}
	close(searchChan)
}

func (this *GSA) DetailPageParse(searchChan chan GSA, resultChan chan GSA, exitChan chan bool) {
	//func (this *GSA) DetailPageParse(URL string) {

	var requestData map[string]string
	requestData = make(map[string]string, 5)
	requestData["pageSize"] = "50"
	requestData["pageNo"] = "1"
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `.\.gov`,
		RandomDelay:  500 * time.Millisecond,
		Parallelism:  12,
	})
	if err != nil {
		log.Println(err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Detail page visiting:%s/pageNo=%s......\n", r.URL, requestData["pageNo"])
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Println("Something went wrong", err, r.StatusCode)
			log.Println("Detail page wrong ", err, r.StatusCode, r.Request.URL)
		}
	})

	c.OnResponse(func(r *colly.Response) {

	})

	c.OnHTML(".container", func(ele *colly.HTMLElement) {

		tempStr := ele.ChildText("div[class='panel-body ']")
		tempStrLst := strings.Split(tempStr, "\n")
		var temp []string
		for _, v := range tempStrLst {
			str := strings.TrimSpace(v)
			if str != "" {
				temp = append(temp, str)
			}
		}
		for _, v := range temp {
			//fmt.Printf("%v:\n%v\n", i, v)
			if strings.HasPrefix(v, "Title") {
				this.GSATitle = strings.TrimSpace(strings.Split(v, ":")[1])
			} else if strings.HasPrefix(v, "Release") {
				this.ReleaseDate = strings.TrimSpace(strings.Split(v, ":")[1])
			} else if strings.HasPrefix(v, "HTTPS") {
				DownloadPath = strings.TrimSpace(strings.Split(v, "：")[1]) + "/"
				tempList := strings.Split(v, "/")
				this.GSAAceesion = tempList[len(tempList)-1]
			} else if strings.HasPrefix(v, "HTTPS") {
				this.BioProject = v
			}
		}
		ele.ForEach("tr", func(_ int, subele *colly.HTMLElement) {
			if subele.Attr("class") == "experiment" {
				var temp []string
				this.ExperimentAccession = subele.ChildText("td[class='experiments']")
				str := subele.Text
				strList := strings.Split(str, "\n")
				for _, v := range strList {
					if strings.TrimSpace(v) != "" {
						temp = append(temp, strings.TrimSpace(v))
					}
				}
				this.ExperimentTitle = temp[1]
				this.TaxonName = temp[2]
				this.PlatForm = temp[3]
				this.SampleAccession = temp[4]

			} else if subele.Attr("class") == "runTr" {
				var temp2 []string
				this.RunAccession = subele.ChildText("td[class='runs']")
				str := subele.Text
				strList := strings.Split(str, "\n")
				for _, v := range strList {
					if strings.TrimSpace(v) != "" {
						temp2 = append(temp2, strings.TrimSpace(v))
					}
				}
				//this.FilePath = ""
				tempPath := ""
				for _, v := range temp2 {

					if strings.HasPrefix(v, "CRR") {
						continue
					} else if strings.HasPrefix(v, "File:") {
						strList := strings.Split(v, ":")
						filename := strings.TrimSpace(strList[1])
						filepath := DownloadPath + this.RunAccession + "/" + filename + ";"
						tempPath += filepath
						this.FilePath = tempPath

					} else {
						this.RunAlias = v
					}
				}
				//this.FilePath = pathTemp
				//fmt.Println(this.ExperimentTitle)
				resultChan <- *this
			}
		})
		//fmt.Println(GSAarr)
	})
	for {
		search, ok := <-searchChan
		if !ok {
			break
		}
		cra := search.GSAAceesion
		if cra == "" {
			break
		}

		fmt.Println(cra)
		TotalPage := this.GetTotalPage(SearchURL + cra)
		end, err := strconv.Atoi(TotalPage)
		if err != nil {
			fmt.Println("Atoi wrong", err)
		}

		for i := 1; i <= end; i++ {
			requestData["pageNo"] = strconv.Itoa(i)

			c.Post(SearchURL+cra, requestData)
		}

	}
	exitChan <- true
}

func (this *GSA) GSACrawler(path string) {
	searchChan := make(chan GSA, 5000)
	resultChan := make(chan GSA, 5000)
	exitChan := make(chan bool, 4)
	start := time.Now()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	go this.SearchPageParse(searchChan)
	for i := 0; i < 4; i++ {
		go this.DetailPageParse(searchChan, resultChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(resultChan)
	}()

	for {
		gwh, ok := <-resultChan
		if !ok {
			break
		}
		str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", gwh.RunAccession, gwh.RunAlias, gwh.FilePath, gwh.ExperimentAccession,
			gwh.ExperimentTitle, gwh.TaxonName, gwh.PlatForm, gwh.SampleAccession, gwh.GSAAceesion,
			gwh.GSATitle, gwh.BioProject, gwh.ReleaseDate)
		file.Write([]byte(str))
	}
	fmt.Println("Done! 用时: ", time.Since(start))
}
