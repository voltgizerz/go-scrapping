package main

import (
    "encoding/json"
    "github.com/PuerkitoBio/goquery"
    "log"
	"net/http"
	"strings"
)

type Article struct {
    Title    string
}

func main() {

	res, err := http.Get("https://www.ninjasaga.com/game-info/all_clan.php")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	type Clan struct {
		ClanRank string `json:"clan_rank"`
		ClanName string `json:"clan_name"`
		ClanMaster string `json:"clan_master"`
		ClanMember string `json:"clan_member"`
		Reputation string `json:"reputation"`
	}

	var row []string
	var rows [][]Clan
	var check []string
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find each table
	doc.Find("tBody").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				if len(strings.TrimSpace(tablecell.Text()))<50  {
					row = append(row, strings.TrimSpace(tablecell.Text()))
					check = append(check, tablecell.Text())
				}
			})


			if len(row)>4 {
				test := []Clan{
					Clan {ClanRank: row[0],ClanName: row[1],ClanMaster:row[2],ClanMember: row[3], Reputation: row[4]},
				}
				rows = append(rows,test)
			}
			row = nil
		})
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(string(bts))
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
	   if a == str {
		  return true
	   }
	}
	return false
 }

