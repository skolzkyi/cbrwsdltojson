//go:build integration
// +build integration

package integrationtests

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"

	"github.com/skolzkyi/cbrwsdltojson/internal/logger"
	"github.com/stretchr/testify/require"
)

var (
	configFilePath string
	config         Config
	log            *logger.LogWrap
)

type AllTestCases struct {
	Cases []TestCase
}

type TestCase struct {
	Method        string
	Handler       string
	Request       string
	OutputControl string
}

func (atc *AllTestCases) Init() {
	atc.Cases = make([]TestCase, 0)
	var curCase TestCase
	curCase = TestCase{
		Method:        "GetCursOnDate",
		Handler:       "/GetCursOnDateXML",
		Request:       `{"OnDate":"2023-06-22"}`,
		OutputControl: `{"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}`,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "GetCursOnDate",
		Handler:       "/GetMethodDataWithoutCache/GetCursOnDateXML",
		Request:       `{"OnDate":"2023-06-22"}`,
		OutputControl: `{"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}`,
	}
	atc.Cases = append(atc.Cases, curCase)
}

func init() {
	flag.StringVar(&configFilePath, "config", "./configs/", "Path to config.env")
}

func TestMain(m *testing.M) {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config = NewConfig()
	err := config.Init(configFilePath)
	if err != nil {
		fmt.Println(err)
	}

	log, err = logger.New(config.Logger.Level, true)
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Info("Integration tests down with error")
			os.Exit(1) //nolint:gocritic
		default:
			log.Info("Integration tests up")
			exitCode := m.Run()
			log.Info("exitCode:" + strconv.Itoa(exitCode))
			log.Info("Integration tests down succesful")
			os.Exit(exitCode) //nolint:gocritic
		}
	}
}

func TestAllIntegrationCases(t *testing.T) {
	AllCases := AllTestCases{}
	AllCases.Init()
	for _, curCase := range AllCases.Cases {
		curCase := curCase
		t.Run(curCase.Method+":"+curCase.Handler, func(t *testing.T) {
			url := "http://" + config.GetServerURL() + curCase.Handler

			jsonStr := []byte(curCase.Request)
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
			require.NoError(t, err)

			respBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			resp.Body.Close()

			require.Equal(t, curCase.OutputControl, string(respBody))
		})
	}
}

/*
func TestGetCursOnDate(t *testing.T) {
	t.Run("GetCursOnDate_Positive", func(t *testing.T) {
		url := helpers.StringBuild("http://", config.GetServerURL(), "/GetCursOnDate")

		jsonStr := []byte(`{"OnDate":"2023-06-22"}`)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		require.Equal(t, string(respBody), `{"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}`)

		url = helpers.StringBuild("http://", config.GetServerURL(), "/GetMethodDataWithoutCache/GetCursOnDate")

		jsonStr = []byte(`{"OnDate":"2023-06-22"}`)
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)
		defer resp.Body.Close()

		respBody, err = io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, string(respBody), `{"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}`)
	})
}
*/