//go:build integration
// +build integration

package integrationtests

import (
	"bytes"
	"context"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"

	datastructures "github.com/skolzkyi/cbrwsdltojson/internal/datastructures"
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
	Method           string
	Handler          string
	Request          string
	OutputControl    string
	UnmControlMethod func(t *testing.T, data []byte)
	Mode             int
}

func (atc *AllTestCases) Init() {
	atc.Cases = make([]TestCase, 0)
	var curCase TestCase
	curCase = TestCase{
		Method:        "GetCursOnDate",
		Handler:       "/GetCursOnDateXML",
		Request:       `{"OnDate":"2023-06-22"}`,
		OutputControl: `{"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "BiCurBaseXML",
		Handler:       "/BiCurBaseXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"BCB":[{"D0":"2023-06-22T00:00:00+03:00","VAL":"87.736315"},{"D0":"2023-06-23T00:00:00+03:00","VAL":"87.358585"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "BliquidityXML",
		Handler:       "/BliquidityXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"BL":[{"DT":"2023-06-23T00:00:00+03:00","StrLiDef":"-1022.50","claims":"1533.70","actionBasedRepoFX":"1378.40","actionBasedSecureLoans":"0.00","standingFacilitiesRepoFX":"0.00","standingFacilitiesSecureLoans":"155.30","liabilities":"-2890.20","depositAuctionBased":"-1828.30","depositStandingFacilities":"-1061.90","CBRbonds":"0.00","netCBRclaims":"334.10"},{"DT":"2023-06-22T00:00:00+03:00","StrLiDef":"-980.70","claims":"1558.80","actionBasedRepoFX":"1378.40","actionBasedSecureLoans":"0.00","standingFacilitiesRepoFX":"0.00","standingFacilitiesSecureLoans":"180.40","liabilities":"-2873.00","depositAuctionBased":"-1828.30","depositStandingFacilities":"-1044.60","CBRbonds":"0.00","netCBRclaims":"333.40"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "DepoDynamicXML",
		Handler:       "/DepoDynamicXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"Depo":[{"DateDepo":"2023-06-22T00:00:00+03:00","Overnight":"6.50"},{"DateDepo":"2023-06-23T00:00:00+03:00","Overnight":"6.50"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "DragMetDynamicXML",
		Handler:       "/DragMetDynamicXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"DrgMet":[{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"1","price":"5228.8000"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"2","price":"64.3800"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"3","price":"2611.0800"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"4","price":"3786.6100"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"1","price":"5176.2400"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"2","price":"62.0300"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"3","price":"2550.9600"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"4","price":"3610.0500"}]}`,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "DVXML",
		Handler:       "/DVXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"DV":[{"Date":"2023-06-22T00:00:00+03:00","VOvern":"0.0000","VLomb":"9051.4000","VIDay":"281.3800","VOther":"504831.8300","Vol_Gold":"0.0000","VIDate":"2023-06-21T00:00:00+03:00"},{"Date":"2023-06-23T00:00:00+03:00","VOvern":"0.0000","VLomb":"8851.4000","VIDay":"118.5300","VOther":"480499.1600","Vol_Gold":"0.0000","VIDate":"2023-06-22T00:00:00+03:00"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "EnumReutersValutesXML",
		Handler:       "/EnumReutersValutesXML",
		Request:       "",
		OutputControl: `{"EnumRValutes":[{"num_code":8,"char_code":"ALL ","Title_ru":"Албанский лек","Title_en":"Albanian Lek"},{"num_code":12,"char_code":"DZD ","Title_ru":"Алжирский динар","Title_en":"Algerian Dinar"},{"num_code":32,"char_code":"ARS ","Title_ru":"Аргентинское песо","Title_en":"Argentine Peso"},{"num_code":44,"char_code":"BSD ","Title_ru":"Багамский доллар","Title_en":"Bahamian Dollar"},{"num_code":48,"char_code":"BHD ","Title_ru":"Бахрейнский динар","Title_en":"Bahraini Dinar"},{"num_code":50,"char_code":"BDT ","Title_ru":"Бангладешская така","Title_en":"Bangladeshi Taka"},{"num_code":52,"char_code":"BBD ","Title_ru":"Барбадосский доллар","Title_en":"Barbados Dollar"},{"num_code":60,"char_code":"BMD ","Title_ru":"Бермудский доллар","Title_en":"Bermudian Dollar"},{"num_code":64,"char_code":"BTN ","Title_ru":"Бутанский нгултрум","Title_en":"Bhutan Ngultrum"},{"num_code":68,"char_code":"BOB ","Title_ru":"Боливийский боливиано","Title_en":"Bolivian Boliviano"},{"num_code":72,"char_code":"BWP ","Title_ru":"Ботсванская пула","Title_en":"Botswana Pula"},{"num_code":84,"char_code":"BZD ","Title_ru":"Белизский доллар","Title_en":"Belize Dollar"},{"num_code":90,"char_code":"SBD ","Title_ru":"Доллар Соломоновых Островов","Title_en":"Solomon Is. Dollar"},{"num_code":96,"char_code":"BND ","Title_ru":"Брунейский доллар","Title_en":"Brunei Dollar"},{"num_code":108,"char_code":"BIF ","Title_ru":"Бурундийский франк","Title_en":"Burundi Franc"},{"num_code":116,"char_code":"KHR ","Title_ru":"Камбоджийский риель","Title_en":"Cambodia Riel"},{"num_code":132,"char_code":"CVE ","Title_ru":"Эскудо Кабо-Верде","Title_en":"Cabo Verde Escudo"},{"num_code":144,"char_code":"LKR ","Title_ru":"Шри-Ланкийская рупия","Title_en":"Sri Lanka Rupee"},{"num_code":152,"char_code":"CLP ","Title_ru":"Чилийское песо","Title_en":"Chilean Peso"},{"num_code":170,"char_code":"COP ","Title_ru":"Колумбийское песо","Title_en":"Colombian Peso"},{"num_code":174,"char_code":"KMF ","Title_ru":"Коморский франк","Title_en":"Comorian Franc"},{"num_code":188,"char_code":"CRC ","Title_ru":"Костариканский колон","Title_en":"Costa Rican Colon"},{"num_code":191,"char_code":"HRK ","Title_ru":"Хорватская куна","Title_en":"Croatian Kuna"},{"num_code":192,"char_code":"CUP ","Title_ru":"Кубинское песо","Title_en":"Cuban Peso"},{"num_code":214,"char_code":"DOP ","Title_ru":"Доминиканское песо","Title_en":"Dominican Peso"},{"num_code":222,"char_code":"SVC ","Title_ru":"Сальвадорский колон","Title_en":"El Salvador Colon"},{"num_code":230,"char_code":"ETB ","Title_ru":"Эфиопский быр","Title_en":"Ethiopian Birr"},{"num_code":232,"char_code":"ERN ","Title_ru":"Эритрейская накфа","Title_en":"Eritrea Nakfa"},{"num_code":238,"char_code":"FKP ","Title_ru":"Фунт Фолклендских островов","Title_en":"Falkland Islands Pound"},{"num_code":242,"char_code":"FJD ","Title_ru":"Доллар Фиджи","Title_en":"Fiji Dollar"},{"num_code":262,"char_code":"DJF ","Title_ru":"Франк Джибути","Title_en":"Djibouti Franc"},{"num_code":270,"char_code":"GMD ","Title_ru":"Гамбийский даласи","Title_en":"Gambian Dalasi"},{"num_code":292,"char_code":"GIP ","Title_ru":"Гибралтарский фунт","Title_en":"Gibraltar Pound"},{"num_code":320,"char_code":"GTQ ","Title_ru":"Гватемальский кетсаль","Title_en":"Guatemala Quetzal"},{"num_code":324,"char_code":"GNF ","Title_ru":"Гвинейский франк","Title_en":"Guinea Franc"},{"num_code":328,"char_code":"GYD ","Title_ru":"Гайанский доллар","Title_en":"Guyana Dollar"},{"num_code":332,"char_code":"HTG ","Title_ru":"Гаитский гурд","Title_en":"Haiti Gourde"},{"num_code":340,"char_code":"HNL ","Title_ru":"Гондурасская лемпира","Title_en":"Honduras Lempira"},{"num_code":344,"char_code":"HKD ","Title_ru":"Гонконгский доллар","Title_en":"Hong Kong Dollar"},{"num_code":352,"char_code":"ISK ","Title_ru":"Исландская крона","Title_en":"Iceland Krona"},{"num_code":360,"char_code":"IDR ","Title_ru":"Индонезийская рупия","Title_en":"Indonesian Rupiah"},{"num_code":364,"char_code":"IRR ","Title_ru":"Иранский риал","Title_en":"Iranian Rial"},{"num_code":368,"char_code":"IQD ","Title_ru":"Иракский динар","Title_en":"Iraqi Dinar"},{"num_code":376,"char_code":"ILS ","Title_ru":"Новый израильский шекель","Title_en":"New Israeli Sheqel"},{"num_code":388,"char_code":"JMD ","Title_ru":"Ямайский доллар","Title_en":"Jamaican Dollar"},{"num_code":400,"char_code":"JOD ","Title_ru":"Иорданский динар","Title_en":"Jordanian Dinar"},{"num_code":404,"char_code":"KES ","Title_ru":"Кенийский шиллинг","Title_en":"Kenyan Shilling"},{"num_code":408,"char_code":"KPW ","Title_ru":"Северокорейская вона","Title_en":"North Korean Won"},{"num_code":414,"char_code":"KWD ","Title_ru":"Кувейтский динар","Title_en":"Kuwaiti Dinar"},{"num_code":418,"char_code":"LAK ","Title_ru":"Лаосский кип","Title_en":"Lao Kip"},{"num_code":422,"char_code":"LBP ","Title_ru":"Ливанский фунт","Title_en":"Lebanese Pound"},{"num_code":430,"char_code":"LRD ","Title_ru":"Либерийский доллар","Title_en":"Liberian Dollar"},{"num_code":434,"char_code":"LYD ","Title_ru":"Ливийский динар","Title_en":"Libyan Dinar"},{"num_code":446,"char_code":"MOP ","Title_ru":"Патака Макао","Title_en":"Macao Pataca"},{"num_code":454,"char_code":"MWK ","Title_ru":"Малавийская квача","Title_en":"Malawi Kwacha"},{"num_code":458,"char_code":"MYR ","Title_ru":"Малайзийский ринггит","Title_en":"Malaysian Ringgit"},{"num_code":462,"char_code":"MVR ","Title_ru":"Мальдивская руфия","Title_en":"Maldives Rufiyaa"},{"num_code":478,"char_code":"MRO ","Title_ru":"Мавританская угия","Title_en":"Mauritania Ouguiya"},{"num_code":480,"char_code":"MUR ","Title_ru":"Маврикийская рупия","Title_en":"Mauritius Rupee"},{"num_code":484,"char_code":"MXN ","Title_ru":"Мексиканское песо","Title_en":"Mexican Peso"},{"num_code":496,"char_code":"MNT ","Title_ru":"Монгольский тугрик","Title_en":"Mongolia Tugrik"},{"num_code":504,"char_code":"MAD ","Title_ru":"Марокканский дирхам","Title_en":"Moroccan Dirham"},{"num_code":512,"char_code":"OMR ","Title_ru":"Оманский риал","Title_en":"Rial Omani"},{"num_code":516,"char_code":"NAD ","Title_ru":"Доллар Намибии","Title_en":"Namibia Dollar"},{"num_code":524,"char_code":"NPR ","Title_ru":"Непальская рупия","Title_en":"Nepalese Rupee"},{"num_code":533,"char_code":"AWG ","Title_ru":"Арубанский флорин","Title_en":"Aruban Florin"},{"num_code":548,"char_code":"VUV ","Title_ru":"Вануатский вату","Title_en":"Vanuatu Vatu"},{"num_code":554,"char_code":"NZD ","Title_ru":"Новозеландский доллар","Title_en":"New Zealand Dollar"},{"num_code":558,"char_code":"NIO ","Title_ru":"Никарагуанская золотая кордоба","Title_en":"Cordoba Oro"},{"num_code":566,"char_code":"NGN ","Title_ru":"Нигерийская найра","Title_en":"Nigerian Naira"},{"num_code":586,"char_code":"PKR ","Title_ru":"Пакистанская рупия","Title_en":"Pakistan Rupee"},{"num_code":590,"char_code":"PAB ","Title_ru":"Панамский бальбоа","Title_en":"Panama Balboa"},{"num_code":598,"char_code":"PGK ","Title_ru":"Кина Папуа-Новой Гвинеи","Title_en":"Papua New Guinean Kina"},{"num_code":600,"char_code":"PYG ","Title_ru":"Парагвайский гуарани","Title_en":"Paraguay Guarani"},{"num_code":604,"char_code":"PEN ","Title_ru":"Перуанский соль","Title_en":"Peru Sol"},{"num_code":608,"char_code":"PHP ","Title_ru":"Филиппинское писо","Title_en":"Philippine Piso"},{"num_code":634,"char_code":"QAR ","Title_ru":"Катарский риал","Title_en":"Qatari Rial"},{"num_code":646,"char_code":"RWF ","Title_ru":"Франк Руанды","Title_en":"Rwanda Franc"},{"num_code":654,"char_code":"SHP ","Title_ru":"Фунт Св. Елены","Title_en":"St Helena Pound"},{"num_code":678,"char_code":"STD ","Title_ru":"Добра Сан-Томе и Принсипи","Title_en":"Sao Tome \u0026 Principe Dobra"},{"num_code":682,"char_code":"SAR ","Title_ru":"Саудовский риял","Title_en":"Saudi Riyal"},{"num_code":690,"char_code":"SCR ","Title_ru":"Сейшельская рупия","Title_en":"Seychelles Rupee"},{"num_code":694,"char_code":"SLL ","Title_ru":"Сьерра-Леонский леоне","Title_en":"Sierra Leone Leone"},{"num_code":704,"char_code":"VND ","Title_ru":"Вьетнамский донг","Title_en":"Vietnam Dong"},{"num_code":706,"char_code":"SOS ","Title_ru":"Сомалийский шиллинг","Title_en":"Somali Shilling"},{"num_code":748,"char_code":"SZL ","Title_ru":"Свазилендский лилангени","Title_en":"Swaziland Lilangeni"},{"num_code":760,"char_code":"SYP ","Title_ru":"Сирийский фунт","Title_en":"Syrian Pound"},{"num_code":764,"char_code":"THB ","Title_ru":"Таиландский бат","Title_en":"Thai Baht"},{"num_code":776,"char_code":"TOP ","Title_ru":"Паанга Королевства Тонга","Title_en":"Tonga Pa'anga"},{"num_code":780,"char_code":"TTD ","Title_ru":"Доллар Тринидада и Тобаго","Title_en":"Trinidad and Tobago Dollar"},{"num_code":784,"char_code":"AED ","Title_ru":"Дирхам ОАЭ","Title_en":"UAE Dirham"},{"num_code":788,"char_code":"TND ","Title_ru":"Тунисский динар","Title_en":"Tunisian Dinar"},{"num_code":800,"char_code":"UGX ","Title_ru":"Угандийский шиллинг","Title_en":"Uganda Shilling"},{"num_code":807,"char_code":"MKD ","Title_ru":"Денар Республики Македония","Title_en":"Macedonian Denar"},{"num_code":818,"char_code":"EGP ","Title_ru":"Египетский фунт","Title_en":"Egyptian Pound"},{"num_code":834,"char_code":"TZS ","Title_ru":"Танзанийский шиллинг","Title_en":"Tanzanian Shilling"},{"num_code":858,"char_code":"UYU ","Title_ru":"Уругвайское песо","Title_en":"Peso Uruguayo"},{"num_code":886,"char_code":"YER ","Title_ru":"Йеменский риал","Title_en":"Yemeni Rial"},{"num_code":901,"char_code":"TWD ","Title_ru":"Новый тайваньский доллар","Title_en":"New Taiwan Dollar"},{"num_code":928,"char_code":"VES ","Title_ru":"Венесуэльский боливар cоберано","Title_en":"Venezuela Bolivar Soberano"},{"num_code":929,"char_code":"MRU ","Title_ru":"Мавританская угия","Title_en":"Mauritania Ouguiya"},{"num_code":930,"char_code":"STN ","Title_ru":"Добра Сан-Томе и Принсипи","Title_en":"Sao Tome \u0026 Principe Dobra"},{"num_code":936,"char_code":"GHS ","Title_ru":"Ганский седи","Title_en":"Ghana Cedi"},{"num_code":937,"char_code":"VEF ","Title_ru":"Венесуэльский боливар","Title_en":"Venezuela Bolivar"},{"num_code":938,"char_code":"SDG ","Title_ru":"Суданский фунт","Title_en":"Sudanese Pound"},{"num_code":941,"char_code":"RSD ","Title_ru":"Сербский динар","Title_en":"Serbian Dinar"},{"num_code":943,"char_code":"MZN ","Title_ru":"Мозамбикский метикал","Title_en":"Mozambique Metical"},{"num_code":950,"char_code":"XAF ","Title_ru":"Франк КФА ВЕАС","Title_en":"CFA Franc BEAC"},{"num_code":951,"char_code":"XCD ","Title_ru":"Восточно - карибский доллар","Title_en":"East Caribbean Dollar"},{"num_code":952,"char_code":"XOF ","Title_ru":"Франк КФА ВСЕАО","Title_en":"CFA Franc BCEAO"},{"num_code":967,"char_code":"ZMW ","Title_ru":"Замбийская квача","Title_en":"Zambian Kwacha"},{"num_code":968,"char_code":"SRD ","Title_ru":"Суринамский доллар","Title_en":"Surinam Dollar"},{"num_code":969,"char_code":"MGA ","Title_ru":"Малагасийский ариари","Title_en":"Malagasy Ariary"},{"num_code":971,"char_code":"AFN ","Title_ru":"Афганский афгани","Title_en":"Afghan Afghani"},{"num_code":973,"char_code":"AOA ","Title_ru":"Ангольская кванза","Title_en":"Angolan Kwanza"},{"num_code":976,"char_code":"CDF ","Title_ru":"Конголезский франк","Title_en":"Congolese Franc"},{"num_code":977,"char_code":"BAM ","Title_ru":"Конвертируемая марка","Title_en":"Convertible Mark"},{"num_code":981,"char_code":"GEL ","Title_ru":"Грузинский лари","Title_en":"Georgian Lari"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "EnumValutesXML",
		Handler:       "/EnumValutesXML",
		Request:       `{"Seld":false}`,
		OutputControl: `{"EnumValutes":[{"Vcode":"R01010","Vname":"Австралийский доллар","VEngname":"Australian Dollar","Vnom":1,"VcommonCode":"R01010","VnumCode":36,"VcharCode":"AUD"},{"Vcode":"R01015","Vname":"Австрийский шиллинг","VEngname":"Austrian Shilling","Vnom":1000,"VcommonCode":"R01015","VnumCode":40,"VcharCode":"ATS"},{"Vcode":"R01020A","Vname":"Азербайджанский манат","VEngname":"Azerbaijan Manat","Vnom":1,"VcommonCode":"R01020","VnumCode":944,"VcharCode":"AZN"},{"Vcode":"R01035","Vname":"Фунт стерлингов Соединенного королевства","VEngname":"British Pound Sterling","Vnom":1,"VcommonCode":"R01035","VnumCode":826,"VcharCode":"GBP"},{"Vcode":"R01040F","Vname":"Ангольская новая кванза","VEngname":"Angolan new Kwanza","Vnom":100000,"VcommonCode":"R01040","VnumCode":24,"VcharCode":"AON"},{"Vcode":"R01060","Vname":"Армянский драм","VEngname":"Armenia Dram","Vnom":1000,"VcommonCode":"R01060","VnumCode":51,"VcharCode":"AMD"},{"Vcode":"R01090B","Vname":"Белорусский рубль","VEngname":"Belarussian Ruble","Vnom":1,"VcommonCode":"R01090","VnumCode":933,"VcharCode":"BYN"},{"Vcode":"R01095","Vname":"Бельгийский франк","VEngname":"Belgium Franc","Vnom":1000,"VcommonCode":"R01095","VnumCode":56,"VcharCode":"BEF"},{"Vcode":"R01100","Vname":"Болгарский лев","VEngname":"Bulgarian lev","Vnom":1,"VcommonCode":"R01100","VnumCode":975,"VcharCode":"BGN"},{"Vcode":"R01115","Vname":"Бразильский реал","VEngname":"Brazil Real","Vnom":1,"VcommonCode":"R01115","VnumCode":986,"VcharCode":"BRL"},{"Vcode":"R01135","Vname":"Венгерский форинт","VEngname":"Hungarian Forint","Vnom":100,"VcommonCode":"R01135","VnumCode":348,"VcharCode":"HUF"},{"Vcode":"R01150","Vname":"Вьетнамский донг","VEngname":"Vietnam Dong","Vnom":10000,"VcommonCode":"R01150","VnumCode":704,"VcharCode":"VND"},{"Vcode":"R01200","Vname":"Гонконгский доллар","VEngname":"Hong Kong Dollar","Vnom":10,"VcommonCode":"R01200","VnumCode":344,"VcharCode":"HKD"},{"Vcode":"R01205","Vname":"Греческая драхма","VEngname":"Greek Drachma","Vnom":10000,"VcommonCode":"R01205","VnumCode":300,"VcharCode":"GRD"},{"Vcode":"R01210","Vname":"Грузинский лари","VEngname":"Georgia Lari","Vnom":1,"VcommonCode":"R01210","VnumCode":981,"VcharCode":"GEL"},{"Vcode":"R01215","Vname":"Датская крона","VEngname":"Danish Krone","Vnom":10,"VcommonCode":"R01215","VnumCode":208,"VcharCode":"DKK"},{"Vcode":"R01230","Vname":"Дирхам ОАЭ","VEngname":"UAE Dirham","Vnom":10,"VcommonCode":"R01230","VnumCode":784,"VcharCode":"AED"},{"Vcode":"R01235","Vname":"Доллар США","VEngname":"US Dollar","Vnom":1,"VcommonCode":"R01235","VnumCode":840,"VcharCode":"USD"},{"Vcode":"R01239","Vname":"Евро","VEngname":"Euro","Vnom":1,"VcommonCode":"R01239","VnumCode":978,"VcharCode":"EUR"},{"Vcode":"R01240","Vname":"Египетский фунт","VEngname":"Egyptian Pound","Vnom":10,"VcommonCode":"R01240","VnumCode":818,"VcharCode":"EGP"},{"Vcode":"R01270","Vname":"Индийская рупия","VEngname":"Indian Rupee","Vnom":100,"VcommonCode":"R01270","VnumCode":356,"VcharCode":"INR"},{"Vcode":"R01280","Vname":"Индонезийская рупия","VEngname":"Indonesian Rupiah","Vnom":10000,"VcommonCode":"R01280","VnumCode":360,"VcharCode":"IDR"},{"Vcode":"R01305","Vname":"Ирландский фунт","VEngname":"Irish Pound","Vnom":100,"VcommonCode":"R01305","VnumCode":372,"VcharCode":"IEP"},{"Vcode":"R01310","Vname":"Исландская крона","VEngname":"Iceland Krona","Vnom":10000,"VcommonCode":"R01310","VnumCode":352,"VcharCode":"ISK"},{"Vcode":"R01315","Vname":"Испанская песета","VEngname":"Spanish Peseta","Vnom":10000,"VcommonCode":"R01315","VnumCode":724,"VcharCode":"ESP"},{"Vcode":"R01325","Vname":"Итальянская лира","VEngname":"Italian Lira","Vnom":100000,"VcommonCode":"R01325","VnumCode":380,"VcharCode":"ITL"},{"Vcode":"R01335","Vname":"Казахстанский тенге","VEngname":"Kazakhstan Tenge","Vnom":100,"VcommonCode":"R01335","VnumCode":398,"VcharCode":"KZT"},{"Vcode":"R01350","Vname":"Канадский доллар","VEngname":"Canadian Dollar","Vnom":1,"VcommonCode":"R01350","VnumCode":124,"VcharCode":"CAD"},{"Vcode":"R01355","Vname":"Катарский риал","VEngname":"Qatari Riyal","Vnom":10,"VcommonCode":"R01355","VnumCode":634,"VcharCode":"QAR"},{"Vcode":"R01370","Vname":"Киргизский сом","VEngname":"Kyrgyzstan Som","Vnom":100,"VcommonCode":"R01370","VnumCode":417,"VcharCode":"KGS"},{"Vcode":"R01375","Vname":"Китайский юань","VEngname":"China Yuan","Vnom":10,"VcommonCode":"R01375","VnumCode":156,"VcharCode":"CNY"},{"Vcode":"R01390","Vname":"Кувейтский динар","VEngname":"Kuwaiti Dinar","Vnom":10,"VcommonCode":"R01390","VnumCode":414,"VcharCode":"KWD"},{"Vcode":"R01405","Vname":"Латвийский лат","VEngname":"Latvian Lat","Vnom":1,"VcommonCode":"R01405","VnumCode":428,"VcharCode":"LVL"},{"Vcode":"R01420","Vname":"Ливанский фунт","VEngname":"Lebanese Pound","Vnom":100000,"VcommonCode":"R01420","VnumCode":422,"VcharCode":"LBP"},{"Vcode":"R01435","Vname":"Литовский лит","VEngname":"Lithuanian Lita","Vnom":1,"VcommonCode":"R01435","VnumCode":440,"VcharCode":"LTL"},{"Vcode":"R01436","Vname":"Литовский талон","VEngname":"Lithuanian talon","Vnom":1,"VcommonCode":"R01435","VnumCode":0,"VcharCode":""},{"Vcode":"R01500","Vname":"Молдавский лей","VEngname":"Moldova Lei","Vnom":10,"VcommonCode":"R01500","VnumCode":498,"VcharCode":"MDL"},{"Vcode":"R01510","Vname":"Немецкая марка","VEngname":"Deutsche Mark","Vnom":1,"VcommonCode":"R01510","VnumCode":276,"VcharCode":"DEM"},{"Vcode":"R01510A","Vname":"Немецкая марка","VEngname":"Deutsche Mark","Vnom":100,"VcommonCode":"R01510","VnumCode":280,"VcharCode":"DEM"},{"Vcode":"R01523","Vname":"Нидерландский гульден","VEngname":"Netherlands Gulden","Vnom":100,"VcommonCode":"R01523","VnumCode":528,"VcharCode":"NLG"},{"Vcode":"R01530","Vname":"Новозеландский доллар","VEngname":"New Zealand Dollar","Vnom":1,"VcommonCode":"R01530","VnumCode":554,"VcharCode":"NZD"},{"Vcode":"R01535","Vname":"Норвежская крона","VEngname":"Norwegian Krone","Vnom":10,"VcommonCode":"R01535","VnumCode":578,"VcharCode":"NOK"},{"Vcode":"R01565","Vname":"Польский злотый","VEngname":"Polish Zloty","Vnom":1,"VcommonCode":"R01565","VnumCode":985,"VcharCode":"PLN"},{"Vcode":"R01570","Vname":"Португальский эскудо","VEngname":"Portuguese Escudo","Vnom":10000,"VcommonCode":"R01570","VnumCode":620,"VcharCode":"PTE"},{"Vcode":"R01585","Vname":"Румынский лей","VEngname":"Romanian Leu","Vnom":10000,"VcommonCode":"R01585","VnumCode":642,"VcharCode":"ROL"},{"Vcode":"R01585F","Vname":"Румынский лей","VEngname":"Romanian Leu","Vnom":10,"VcommonCode":"R01585","VnumCode":946,"VcharCode":"RON"},{"Vcode":"R01589","Vname":"СДР (специальные права заимствования)","VEngname":"SDR","Vnom":1,"VcommonCode":"R01589","VnumCode":960,"VcharCode":"XDR"},{"Vcode":"R01625","Vname":"Сингапурский доллар","VEngname":"Singapore Dollar","Vnom":1,"VcommonCode":"R01625","VnumCode":702,"VcharCode":"SGD"},{"Vcode":"R01665A","Vname":"Суринамский доллар","VEngname":"Surinam Dollar","Vnom":1,"VcommonCode":"R01665","VnumCode":968,"VcharCode":"SRD"},{"Vcode":"R01670","Vname":"Таджикский сомони","VEngname":"Tajikistan Ruble","Vnom":10,"VcommonCode":"R01670","VnumCode":972,"VcharCode":"TJS"},{"Vcode":"R01675","Vname":"Таиландский бат","VEngname":"Thai Baht","Vnom":100,"VcommonCode":"R01675","VnumCode":764,"VcharCode":"THB"},{"Vcode":"R01700J","Vname":"Турецкая лира","VEngname":"Turkish Lira","Vnom":1,"VcommonCode":"R01700","VnumCode":949,"VcharCode":"TRY"},{"Vcode":"R01710","Vname":"Туркменский манат","VEngname":"Turkmenistan Manat","Vnom":10000,"VcommonCode":"R01710","VnumCode":795,"VcharCode":"TMM"},{"Vcode":"R01710A","Vname":"Новый туркменский манат","VEngname":"New Turkmenistan Manat","Vnom":1,"VcommonCode":"R01710","VnumCode":934,"VcharCode":"TMT"},{"Vcode":"R01717","Vname":"Узбекский сум","VEngname":"Uzbekistan Sum","Vnom":1000,"VcommonCode":"R01717","VnumCode":860,"VcharCode":"UZS"},{"Vcode":"R01720","Vname":"Украинская гривна","VEngname":"Ukrainian Hryvnia","Vnom":10,"VcommonCode":"R01720","VnumCode":980,"VcharCode":"UAH"},{"Vcode":"R01720A","Vname":"Украинский карбованец","VEngname":"Ukrainian Hryvnia","Vnom":1,"VcommonCode":"R01720","VnumCode":0,"VcharCode":""},{"Vcode":"R01740","Vname":"Финляндская марка","VEngname":"Finnish Marka","Vnom":100,"VcommonCode":"R01740","VnumCode":246,"VcharCode":"FIM"},{"Vcode":"R01750","Vname":"Французский франк","VEngname":"French Franc","Vnom":1000,"VcommonCode":"R01750","VnumCode":250,"VcharCode":"FRF"},{"Vcode":"R01760","Vname":"Чешская крона","VEngname":"Czech Koruna","Vnom":10,"VcommonCode":"R01760","VnumCode":203,"VcharCode":"CZK"},{"Vcode":"R01770","Vname":"Шведская крона","VEngname":"Swedish Krona","Vnom":10,"VcommonCode":"R01770","VnumCode":752,"VcharCode":"SEK"},{"Vcode":"R01775","Vname":"Швейцарский франк","VEngname":"Swiss Franc","Vnom":1,"VcommonCode":"R01775","VnumCode":756,"VcharCode":"CHF"},{"Vcode":"R01790","Vname":"ЭКЮ","VEngname":"ECU","Vnom":1,"VcommonCode":"R01790","VnumCode":954,"VcharCode":"XEU"},{"Vcode":"R01795","Vname":"Эстонская крона","VEngname":"Estonian Kroon","Vnom":10,"VcommonCode":"R01795","VnumCode":233,"VcharCode":"EEK"},{"Vcode":"R01805","Vname":"Югославский новый динар","VEngname":"Yugoslavian Dinar","Vnom":1,"VcommonCode":"R01804","VnumCode":890,"VcharCode":"YUN"},{"Vcode":"R01805F","Vname":"Сербский динар","VEngname":"Serbian Dinar","Vnom":100,"VcommonCode":"R01804","VnumCode":941,"VcharCode":"RSD"},{"Vcode":"R01810","Vname":"Южноафриканский рэнд","VEngname":"S.African Rand","Vnom":10,"VcommonCode":"R01810","VnumCode":710,"VcharCode":"ZAR"},{"Vcode":"R01815","Vname":"Вон Республики Корея","VEngname":"South Korean Won","Vnom":1000,"VcommonCode":"R01815","VnumCode":410,"VcharCode":"KRW"},{"Vcode":"R01820","Vname":"Японская иена","VEngname":"Japanese Yen","Vnom":100,"VcommonCode":"R01820","VnumCode":392,"VcharCode":"JPY"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "KeyRateXML",
		Handler:       "/KeyRateXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"KR":[{"DT":"2023-06-23T00:00:00+03:00","Rate":"7.50"},{"DT":"2023-06-22T00:00:00+03:00","Rate":"7.50"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:  "MainInfoXML",
		Handler: "/MainInfoXML",
		Request: "",
		UnmControlMethod: func(t *testing.T, data []byte) {
			t.Helper()
			testStruct := datastructures.MainInfoXMLResult{}
			XMLToStructDecoder(t, data, "RegData", &testStruct)
		},
		Mode: 2,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "mrrf7DXML",
		Handler:       "/mrrf7DXML",
		Request:       `{"FromDate":"2023-06-15","ToDate":"2023-06-23"}`,
		OutputControl: `{"mr":[{"D0":"2023-06-16T00:00:00+03:00","val":"587.50"},{"D0":"2023-06-23T00:00:00+03:00","val":"586.90"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "mrrfXML",
		Handler:       "/mrrfXML",
		Request:       `{"FromDate":"2023-05-01","ToDate":"2023-06-23"}`,
		OutputControl: `{"mr":[{"D0":"2023-05-01T00:00:00+03:00","p1":"595787.00","p2":"447187.00","p3":"418628.00","p4":"23559.00","p5":"5000.00","p6":"148599.00"},{"D0":"2023-06-01T00:00:00+03:00","p1":"584175.00","p2":"438344.00","p3":"410313.00","p4":"23127.00","p5":"4903.00","p6":"145832.00"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "NewsInfoXML",
		Handler:       "/NewsInfoXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"News":[{"Doc_id":35498,"DocDate":"2023-06-22T19:10:00.07+03:00","Title":"О развитии банковского сектора Российской Федерации в мае 2023 года","Url":"/analytics/bank_sector/develop/#a_48876"},{"Doc_id":35495,"DocDate":"2023-06-22T09:35:00+03:00","Title":"Указание Банка России от 10.01.2023 № 6356-У","Url":"/Queries/UniDbQuery/File/90134/2803"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:  "OmodInfoXML",
		Handler: "/OmodInfoXML",
		Request: "",
		UnmControlMethod: func(t *testing.T, data []byte) {
			t.Helper()
			testStruct := datastructures.OmodInfoXMLResult{}
			XMLToStructDecoder(t, data, "OMO", &testStruct)
		},
		Mode: 2,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "OstatDepoNewXML",
		Handler:       "/OstatDepoNewXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"odn":[{"DT":"2023-06-22T00:00:00+03:00","TOTAL":"2872966.59","AUC_1W":"1828340.00","OV_P":"1044626.59"},{"DT":"2023-06-23T00:00:00+03:00","TOTAL":"2890199.16","AUC_1W":"1828340.00","OV_P":"1061859.16"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "OstatDepoXML",
		Handler:       "/OstatDepoXML",
		Request:       `{"FromDate":"2022-12-29","ToDate":"2022-12-30"}`,
		OutputControl: `{"odr":[{"D0":"2022-12-29T00:00:00+03:00","D1_7":"1747362.67","D8_30":"2515151.15","total":"4262513.81"},{"D0":"2022-12-30T00:00:00+03:00","D1_7":"1387715.38","D8_30":"2515151.15","total":"3897866.53"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "OstatDynamicXML",
		Handler:       "/OstatDynamicXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"Ostat":[{"DateOst":"2023-06-22T00:00:00+03:00","InRuss":"3756300.00","InMoscow":"3528600.00"},{"DateOst":"2023-06-23T00:00:00+03:00","InRuss":"3688300.00","InMoscow":"3441000.00"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "OvernightXML",
		Handler:       "/OvernightXML",
		Request:       `{"FromDate":"2023-07-22","ToDate":"2023-08-16"}`,
		OutputControl: `{"OB":[{"date":"2023-07-24T00:00:00+03:00","stavka":"9.50"},{"date":"2023-08-15T00:00:00+03:00","stavka":"13.00"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "RepoDebtXML",
		Handler:       "/RepoDebtXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"RD":[{"Date":"2023-06-22T00:00:00+03:00","debt":"1378387.6","debt_auc":"1378387.6","debt_fix":"0.0"},{"Date":"2023-06-23T00:00:00+03:00","debt":"1378379.7","debt_auc":"1378379.7","debt_fix":"0.0"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "RepoDebtUSDXML",
		Handler:       "/RepoDebtUSDXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"rd":[{"D0":"2023-06-22T00:00:00+03:00","TP":0},{"D0":"2023-06-22T00:00:00+03:00","TP":1},{"D0":"2023-06-23T00:00:00+03:00","TP":0},{"D0":"2023-06-23T00:00:00+03:00","TP":1}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "ROISfixXML",
		Handler:       "/ROISfixXML",
		Request:       `{"FromDate":"2022-02-27","ToDate":"2023-03-02"}`,
		OutputControl: `{"rf":[{"D0":"2022-02-28T00:00:00+03:00","R1W":"17.83","R2W":"18.00","R1M":"20.65","R2M":"21.96","R3M":"23.23","R6M":"24.52"},{"D0":"2022-03-01T00:00:00+03:00","R1W":"19.85","R2W":"19.91","R1M":"22.63","R2M":"23.79","R3M":"24.49","R6M":"25.71"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "RuoniaSVXML",
		Handler:       "/RuoniaSVXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"ra":[{"DT":"2023-06-22T00:00:00+03:00","RUONIA_Index":"2.65003371140540","RUONIA_AVG_1M":"7.33031817626889","RUONIA_AVG_3M":"7.28023580262342","RUONIA_AVG_6M":"7.34479164787354"},{"DT":"2023-06-23T00:00:00+03:00","RUONIA_Index":"2.65055282759819","RUONIA_AVG_1M":"7.32512579295002","RUONIA_AVG_3M":"7.27890778428907","RUONIA_AVG_6M":"7.34359578515310"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "RuoniaXML",
		Handler:       "/RuoniaXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"ro":[{"D0":"2023-06-22T00:00:00+03:00","ruo":"7.1500","vol":"367.9500","DateUpdate":"2023-06-23T14:09:39.6+03:00"},{"D0":"2023-06-23T00:00:00+03:00","ruo":"7.1300","vol":"388.4500","DateUpdate":"2023-06-26T14:08:26.15+03:00"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SaldoXML",
		Handler:       "/SaldoXML",
		Request:       `{"FromDate":"2023-06-22","ToDate":"2023-06-23"}`,
		OutputControl: `{"So":[{"Dt":"2023-06-22T00:00:00+03:00","DEADLINEBS":"1044.60"},{"Dt":"2023-06-23T00:00:00+03:00","DEADLINEBS":"1061.30"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapDayTotalXML",
		Handler:       "/SwapDayTotalXML",
		Request:       `{"FromDate":"2022-02-25","ToDate":"2022-02-28"}`,
		OutputControl: `{"SDT":[{"DT":"2022-02-28T00:00:00+03:00","Swap":"0.0"},{"DT":"2022-02-25T00:00:00+03:00","Swap":"24120.4"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapDynamicXML",
		Handler:       "/SwapDynamicXML",
		Request:       `{"FromDate":"2022-02-25","ToDate":"2022-02-28"}`,
		OutputControl: `{"Swap":[{"DateBuy":"2022-02-25T00:00:00+03:00","DateSell":"2022-02-28T00:00:00+03:00","BaseRate":"96.8252","SD":"0.0882","TIR":"10.5000","Stavka":"-0.576000","Currency":1},{"DateBuy":"2022-02-25T00:00:00+03:00","DateSell":"2022-02-28T00:00:00+03:00","BaseRate":"87.1154","SD":"0.0748","TIR":"10.5000","Stavka":"0.050000","Currency":0}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapInfoSellUSDVolXML",
		Handler:       "/SwapInfoSellUSDVolXML",
		Request:       `{"FromDate":"2022-02-24","ToDate":"2022-02-28"}`,
		OutputControl: `{"SSUV":[{"DT":"2022-02-25T00:00:00+03:00","TODTOMrubvol":"435577.0","TODTOMusdvol":"5000.0","TOMSPTrubvol":"128974.3","TOMSPTusdvol":"1480.5"},{"DT":"2022-02-24T00:00:00+03:00","TODTOMrubvol":"403236.5","TODTOMusdvol":"5000.0","TOMSPTrubvol":"32299.2","TOMSPTusdvol":"400.5"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapInfoSellUSDXML",
		Handler:       "/SwapInfoSellUSDXML",
		Request:       `{"FromDate":"2022-02-25","ToDate":"2022-02-28"}`,
		OutputControl: `{"SSU":[{"DateBuy":"2022-02-25T00:00:00+03:00","DateSell":"2022-02-28T00:00:00+03:00","DateSPOT":"2022-03-01T00:00:00+03:00","Type":1,"BaseRate":"87.115400","SD":"0.016500","TIR":"8.5000","Stavka":"1.5500","limit":"2.0000"},{"DateBuy":"2022-02-25T00:00:00+03:00","DateSell":"2022-02-25T00:00:00+03:00","DateSPOT":"2022-02-28T00:00:00+03:00","Type":0,"BaseRate":"87.115400","SD":"0.049600","TIR":"8.5000","Stavka":"1.5500","limit":"5.0000"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapInfoSellVolXML",
		Handler:       "/SwapInfoSellVolXML",
		Request:       `{"FromDate":"2023-05-05","ToDate":"2023-05-10"}`,
		OutputControl: `{"SSUV":[{"DT":"2023-05-10T00:00:00+03:00","Currency":2,"type":0,"VOL_FC":"1113.5","VOL_RUB":"12512.6"},{"DT":"2023-05-05T00:00:00+03:00","Currency":2,"type":0,"VOL_FC":"4583.7","VOL_RUB":"51606.0"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapInfoSellXML",
		Handler:       "/SwapInfoSellXML",
		Request:       `{"FromDate":"2023-06-20","ToDate":"2023-06-21"}`,
		OutputControl: `{"SSU":[{"Currency":2,"DateBuy":"2023-06-21T00:00:00+03:00","DateSell":"2023-06-21T00:00:00+03:00","DateSPOT":"2023-06-26T00:00:00+03:00","Type":0,"BaseRate":"11.764246","SD":"0.003375","TIR":"6.5000","Stavka":"4.3440","limit":"10.0000"},{"Currency":2,"DateBuy":"2023-06-20T00:00:00+03:00","DateSell":"2023-06-20T00:00:00+03:00","DateSPOT":"2023-06-21T00:00:00+03:00","Type":0,"BaseRate":"11.730496","SD":"0.000626","TIR":"6.5000","Stavka":"4.4890","limit":"10.0000"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:        "SwapMonthTotalXML",
		Handler:       "/SwapMonthTotalXML",
		Request:       `{"FromDate":"2022-02-11","ToDate":"2022-02-24"}`,
		OutputControl: `{"SMT":[{"D0":"2022-02-11T00:00:00+03:00","RUB":"41208.1","USD":"553.3"},{"D0":"2022-02-24T00:00:00+03:00","RUB":"24113.5","USD":"299.0"}]}`,
		Mode:          0,
	}
	atc.Cases = append(atc.Cases, curCase)
	curCase = TestCase{
		Method:  "AllDataInfoXML",
		Handler: "/AllDataInfoXML",
		Request: "",
		UnmControlMethod: func(t *testing.T, data []byte) {
			t.Helper()
			testStruct := datastructures.AllDataInfoXMLResult{}
			XMLToStructDecoder(t, data, "AllData", &testStruct)
		},
		Mode: 2,
	}
	atc.Cases = append(atc.Cases, curCase)
}

func XMLToStructDecoder(t *testing.T, data []byte, startNodeName string, pointerToStruct interface{}) {
	t.Helper()
	xmlData := bytes.NewBuffer(data)

	d := xml.NewDecoder(xmlData)

	for token, _ := d.Token(); token != nil; token, _ = d.Token() {
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == startNodeName {
				err := d.DecodeElement(pointerToStruct, &se)
				require.NoError(t, err)
			}
		}
	}

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

			url = "http://" + config.GetServerURL() + "/GetMethodDataWithoutCache" + curCase.Handler

			resp, err = http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
			require.NoError(t, err)

			respBodyWC, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			resp.Body.Close()
			switch curCase.Mode {
			case 0:
				require.Equal(t, curCase.OutputControl, string(respBody))
				require.Equal(t, curCase.OutputControl, string(respBodyWC))
			case 2:
				curCase.UnmControlMethod(t, respBody)
				curCase.UnmControlMethod(t, respBodyWC)
			default:
				require.Fail(t, "unsupported Mode")
			}

		})
	}
}
