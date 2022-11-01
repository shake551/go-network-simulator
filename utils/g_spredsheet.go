package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GSpread struct {
	ID           string
	SheetName    string
	SheetService *sheets.Service
	Result       *[][]interface{}
}

func NewGSpread(packetRate float64) GSpread {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}

	sheetID := os.Getenv("SHEET_ID")
	credentialFilePath := "./credential/erudite-course-366605-30e48fc4772f.json"

	client, err := httpClient(credentialFilePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetService, err := sheets.New(client)
	if err != nil {
		fmt.Println(err)
	}

	gs := GSpread{
		ID:           sheetID,
		SheetService: sheetService,
	}

	gs.SheetName = fmt.Sprintf("%f", packetRate)

	_, err = gs.sheetId(gs.SheetName)
	if err != nil {
		gs.CreateNewSheet(gs.SheetName)
	}

	statisticsData := gs.readStatisticsData(gs.SheetName)
	gs.Result = &statisticsData

	return gs
}

func (g GSpread) readStatisticsData(sheetName string) [][]interface{} {
	statisticsData, err := g.SheetService.Spreadsheets.Values.Get(g.ID, sheetName+"!A:D").Do()
	if err != nil {
		fmt.Println(err)
	}

	return statisticsData.Values
}

func (g GSpread) AppendNewStatistics(packetCount float64, stayTime float64, packetLoss float64) {
	*g.Result = append(*g.Result, []interface{}{packetCount, stayTime, packetLoss})
}

func (g GSpread) Insert() {
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         *g.Result,
	}
	_, err := g.SheetService.Spreadsheets.Values.Update(g.ID, g.SheetName+"!A:D", valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to write value. %v", err)
	}
}

func (g GSpread) CreateNewSheet(sheetName string) {
	req := sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: sheetName + strconv.FormatInt(time.Now().UnixNano(), 10),
			},
		},
	}

	rbb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&req},
	}

	_, err := g.SheetService.Spreadsheets.BatchUpdate(g.ID, rbb).Do()
	if err != nil {
		fmt.Println(err)
	}

	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values: [][]interface{}{
			{"システム内平均パケット数", "平均システム滞在時間", "パケット廃棄率"},
		},
	}
	_, err = g.SheetService.Spreadsheets.Values.Update(g.ID, g.SheetName+"!A:D", valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to write value. %v", err)
	}
}

func httpClient(credentialFilePath string) (*http.Client, error) {
	data, err := ioutil.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext), nil
}

func (g GSpread) sheetId(sheetName string) (int64, error) {
	spreadsheet, err := g.SheetService.Spreadsheets.Get(g.ID).Do()
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == sheetName {
			return sheet.Properties.SheetId, nil
		}
	}
	return 0, errors.New(sheetName + " is not exists.")
}
