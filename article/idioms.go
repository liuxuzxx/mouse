package article

import (
	"fmt"
	"github.com/gogather/mahonia"
	"github.com/opesun/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//只做一个爬取成语的爬虫
//
const (
	rootUrl string = "http://chengyu.t086.com"
	sqlFile string = "/home/liuxu/Documents/idioms.sql"
)

type AttributeColumn struct {
	isFetch    bool
	columnName string
}

var (
	attributeNameMap = map[string]AttributeColumn{
		"词目": {isFetch: true, columnName: "term"},
		"发音": {isFetch: true, columnName: "pronunciation"},
		"释义": {isFetch: true, columnName: "interpretation"},
		"出处": {isFetch: true, columnName: "source"},
		"示例": {isFetch: true, columnName: "example"},
	}
	idiomsCount = 0
)

func SpiderIdioms() {
	fetchIdioms()
}

func fetchIdioms() {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, value := range letters {
		pageCount := 1
		url := rootUrl + "/list/" + value + "_" + strconv.Itoa(pageCount) + ".html"
		for content, statusCode := fetchIdiomsHtml(url); statusCode < 400; content, statusCode = fetchIdiomsHtml(url) {
			fmt.Printf("请求的url为:%s\n",url)
			parseHtml(content)
			pageCount++
			url = rootUrl + "/list/" + value + "_" + strconv.Itoa(pageCount) + ".html"
		}
	}
}

func fetchIdiomsHtml(url string) (content string, statusCode int) {
	response, responseError := http.Get(url)
	if responseError != nil {
		statusCode = 500
		return
	}
	defer response.Body.Close()
	responseBytes, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		statusCode = 500
		return
	}
	statusCode = response.StatusCode
	content = string(responseBytes)

	gbkEncoder := mahonia.NewDecoder("GBK")
	content = gbkEncoder.ConvertString(content)
	return content, statusCode
}

func parseHtml(content string) {
	nodes, err := goquery.ParseString(content)
	if err != nil {
		fmt.Printf("Parse html content is error! the error informaiton is:%s", err)
		return
	}
	idiomsContent := nodes.Find(".listw ul li a")
	for index := 0; index < idiomsContent.Length(); index++ {
		element := idiomsContent.Eq(index)
		href := element.Attr("href")
		url := rootUrl + href
		defineContent, defineStatusCode := fetchIdiomsHtml(url)
		if defineStatusCode < 400 {
			insertSql := parseExplanationHtml(defineContent)
			writeSqlResult(insertSql)
		}
	}
}

func parseExplanationHtml(content string) string {
	nodes, _ := goquery.ParseString(content)
	explanationTable := nodes.Find("#main table:first-child")
	tableTr := explanationTable.Find("tr")
	result := "insert into idioms(term,pronunciation,interpretation,source,example) values(%s);\n"
	var columnValue []string
	for index := 0; index < tableTr.Length(); index++ {
		trElement := tableTr.Eq(index)
		tdElement := trElement.Find("td")
		if tdElement.Length() == 2 {
			attributeName := tdElement.Eq(0).Text()
			column := attributeNameMap[attributeName]
			if column.isFetch {
				attributeValue := tdElement.Eq(1).Text()
				columnValue = append(columnValue, "'"+attributeValue+"'")
			}
		}
	}
	return fmt.Sprintf(result, strings.Join(columnValue, ","))
}

func writeSqlResult(sql string) {
	file, err := os.OpenFile(sqlFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Open file:%s fail,the error is:%s",sqlFile,err)
		return
	}
	defer file.Close()
	_, _ = file.WriteString(sql)
	idiomsCount++
	if idiomsCount%20==0{
		fmt.Printf("写入文件中的成语有:%d 个了！\n",idiomsCount)
	}
}
