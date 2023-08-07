## О сервисе
Данный сервис позволяет получать данные с веб-сервиса ЦБ РФ(http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx) в виде удобного json.  
Сервис завернут в контейнер, что удобно при использовании его в качестве источника данных для нескольких приложений или микросервисов.  
Зависимости от внешних SOAP-библиотек отсутствуют, лишние компоненты вроде SQL и NoSQL баз данных отсутствуют. Сервис задумывался как легковесная прокладка-конвертер между оргинальным веб-сервисом SOAP ЦБР и приложениями, нативно потребляющими json.
Сервис содержит в себе встроенный потоконезависимый кэш, что позволяет экономить на маршаллинге XML при запросе с сервиса ЦБР. 

## Настройки в config.env
  * `ADDRESS=cbrwsdltojson` - адрес, на котором поднимается сервис. При использовании в  контейнере менять не рекомендуется;  
  * `PORT=4000` - порт, на котором поднимается сервис;  
  * `SERVER_SHUTDOWN_TIMEOUT=30s` - таймаут для мягкого выключения сервиса(graceful shutdown);  
  * `CBR_WSDL_TIMEOUT=5s` - таймаут для запроса сервиса(целиком, включая анмаршаллинг и ответ);  
  * `CBR_WSDL_ADDRESS=http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx` - эндпоинт сервиса ЦБР, менять не рекомендуется, добавлено на будущее на случай переезда сервиса;  
  * `INFO_EXPIR_TIME=12h` - промежуток времени, по которому истекает актуальность хранения данных в кеше, если оно превышено, то запрос будет выполнен, минуя кэш(с обновлением кэша);   
  * `PERMITTED_REQUESTS=` - список разрешенных методов, если список пуст, то разрешены все методы, если нет, то выполняться будут только методы из списка(например, `PERMITTED_REQUESTS=GetCursOnDateXML Coins_baseXML` - названия методов необходимо разделять пробелами);  
  * `LOGGING_ON=true` - триггер логирования в текстовый файл, опционально, при высоких нагрузках можно выключать для выигрыша производительности;  

## Кэш
Кеширование данных происходит после первого запроса по данному методу после запуска сервиса.  
Время записи в кэш фиксируется, и по истечении периода, указанного в `INFO_EXPIR_TIME`, информация считается устаревшей и при очередном запросе информация в кэше обновляется.  
Для каждого метода есть возможность запросить принудительно данные напрямую, минуя кэш (данные в кэше после такого запроса также будут обновлены).  
Для принудительного прямого запроса надо выполнить запрос на хендлер вида `/GetMethodDataWithoutCache/[имя метода]` (например,  `/GetMethodDataWithoutCache/GetCursOnDateXML`)  

## Список поддерживаемых методов, примеры json запросов и ответов
Поля с дробной частью возвращаются в виде строк, чтобы максимально сохранить оригинальную сигнатуру методов.  
  
  <details><summary><b>GetCursOnDateXML</b></summary>
    <ul>
        <li>request: {"OnDate":"2023-06-22"} </li>
        <li>response: {"OnDate":"20230622","ValuteCursOnDate":[{"Vname":"Австралийский доллар","Vnom":1,"Vcurs":"57.1445","Vcode":"36","VchCode":"AUD"},{"Vname":"Азербайджанский манат","Vnom":1,"Vcurs":"49.5569","Vcode":"944","VchCode":"AZN"},{"Vname":"Фунт стерлингов Соединенного королевства","Vnom":1,"Vcurs":"107.2882","Vcode":"826","VchCode":"GBP"},{"Vname":"Армянский драм","Vnom":100,"Vcurs":"21.8165","Vcode":"51","VchCode":"AMD"},{"Vname":"Белорусский рубль","Vnom":1,"Vcurs":"28.2073","Vcode":"933","VchCode":"BYN"},{"Vname":"Болгарский лев","Vnom":1,"Vcurs":"47.0941","Vcode":"975","VchCode":"BGN"},{"Vname":"Бразильский реал","Vnom":1,"Vcurs":"17.5781","Vcode":"986","VchCode":"BRL"},{"Vname":"Венгерский форинт","Vnom":100,"Vcurs":"24.7799","Vcode":"348","VchCode":"HUF"},{"Vname":"Вьетнамский донг","Vnom":10000,"Vcurs":"35.5067","Vcode":"704","VchCode":"VND"},{"Vname":"Гонконгский доллар","Vnom":1,"Vcurs":"10.7815","Vcode":"344","VchCode":"HKD"},{"Vname":"Грузинский лари","Vnom":1,"Vcurs":"32.1995","Vcode":"981","VchCode":"GEL"},{"Vname":"Датская крона","Vnom":1,"Vcurs":"12.3649","Vcode":"208","VchCode":"DKK"},{"Vname":"Дирхам ОАЭ","Vnom":1,"Vcurs":"22.9368","Vcode":"784","VchCode":"AED"},{"Vname":"Доллар США","Vnom":1,"Vcurs":"84.2467","Vcode":"840","VchCode":"USD"},{"Vname":"Евро","Vnom":1,"Vcurs":"92.0014","Vcode":"978","VchCode":"EUR"},{"Vname":"Египетский фунт","Vnom":10,"Vcurs":"27.2655","Vcode":"818","VchCode":"EGP"},{"Vname":"Индийская рупия","Vnom":10,"Vcurs":"10.2348","Vcode":"356","VchCode":"INR"},{"Vname":"Индонезийская рупия","Vnom":10000,"Vcurs":"56.0151","Vcode":"360","VchCode":"IDR"},{"Vname":"Казахстанский тенге","Vnom":100,"Vcurs":"18.7925","Vcode":"398","VchCode":"KZT"},{"Vname":"Канадский доллар","Vnom":1,"Vcurs":"63.6256","Vcode":"124","VchCode":"CAD"},{"Vname":"Катарский риал","Vnom":1,"Vcurs":"23.1447","Vcode":"634","VchCode":"QAR"},{"Vname":"Киргизский сом","Vnom":100,"Vcurs":"96.4979","Vcode":"417","VchCode":"KGS"},{"Vname":"Китайский юань","Vnom":1,"Vcurs":"11.7059","Vcode":"156","VchCode":"CNY"},{"Vname":"Молдавский лей","Vnom":10,"Vcurs":"46.8829","Vcode":"498","VchCode":"MDL"},{"Vname":"Новозеландский доллар","Vnom":1,"Vcurs":"51.9718","Vcode":"554","VchCode":"NZD"},{"Vname":"Норвежская крона","Vnom":10,"Vcurs":"78.2300","Vcode":"578","VchCode":"NOK"},{"Vname":"Польский злотый","Vnom":1,"Vcurs":"20.7137","Vcode":"985","VchCode":"PLN"},{"Vname":"Румынский лей","Vnom":1,"Vcurs":"18.5431","Vcode":"946","VchCode":"RON"},{"Vname":"СДР (специальные права заимствования)","Vnom":1,"Vcurs":"112.7305","Vcode":"960","VchCode":"XDR"},{"Vname":"Сингапурский доллар","Vnom":1,"Vcurs":"62.6929","Vcode":"702","VchCode":"SGD"},{"Vname":"Таджикский сомони","Vnom":10,"Vcurs":"77.1942","Vcode":"972","VchCode":"TJS"},{"Vname":"Таиландский бат","Vnom":10,"Vcurs":"24.1945","Vcode":"764","VchCode":"THB"},{"Vname":"Турецкая лира","Vnom":10,"Vcurs":"35.7005","Vcode":"949","VchCode":"TRY"},{"Vname":"Новый туркменский манат","Vnom":1,"Vcurs":"24.0705","Vcode":"934","VchCode":"TMT"},{"Vname":"Узбекский сум","Vnom":10000,"Vcurs":"73.3218","Vcode":"860","VchCode":"UZS"},{"Vname":"Украинская гривна","Vnom":10,"Vcurs":"22.8114","Vcode":"980","VchCode":"UAH"},{"Vname":"Чешская крона","Vnom":10,"Vcurs":"38.7965","Vcode":"203","VchCode":"CZK"},{"Vname":"Шведская крона","Vnom":10,"Vcurs":"78.0040","Vcode":"752","VchCode":"SEK"},{"Vname":"Швейцарский франк","Vnom":1,"Vcurs":"93.7429","Vcode":"756","VchCode":"CHF"},{"Vname":"Сербский динар","Vnom":100,"Vcurs":"78.4473","Vcode":"941","VchCode":"RSD"},{"Vname":"Южноафриканский рэнд","Vnom":10,"Vcurs":"45.9696","Vcode":"710","VchCode":"ZAR"},{"Vname":"Вон Республики Корея","Vnom":1000,"Vcurs":"65.2064","Vcode":"410","VchCode":"KRW"},{"Vname":"Японская иена","Vnom":100,"Vcurs":"59.4963","Vcode":"392","VchCode":"JPY"}]}</li>
    </ul>
   </details>
   <details><summary><b>BiCurBaseXML</b></summary>
    <ul>
        <li>request: {"FromDate":"2023-06-22","ToDate":"2023-06-23"}</li>
        <li>response: {"BCB":[{"D0":"2023-06-22T00:00:00+03:00","VAL":"87.736315"},{"D0":"2023-06-23T00:00:00+03:00","VAL":"87.358585"}]}</li>
    </ul>
   </details>
   <details><summary><b>BliquidityXML</b></summary>
    <ul>
        <li>request: {"FromDate":"2023-06-22","ToDate":"2023-06-23"}</li>
        <li>response: {"BL":[{"DT":"2023-06-23T00:00:00+03:00","StrLiDef":"-1022.50","claims":"1533.70","actionBasedRepoFX":"1378.40","actionBasedSecureLoans":"0.00","standingFacilitiesRepoFX":"0.00","standingFacilitiesSecureLoans":"155.30","liabilities":"-2890.20","depositAuctionBased":"-1828.30","depositStandingFacilities":"-1061.90","CBRbonds":"0.00","netCBRclaims":"334.10"},{"DT":"2023-06-22T00:00:00+03:00","StrLiDef":"-980.70","claims":"1558.80","actionBasedRepoFX":"1378.40","actionBasedSecureLoans":"0.00","standingFacilitiesRepoFX":"0.00","standingFacilitiesSecureLoans":"180.40","liabilities":"-2873.00","depositAuctionBased":"-1828.30","depositStandingFacilities":"-1044.60","CBRbonds":"0.00","netCBRclaims":"333.40"}]}</li>
    </ul>
   </details>
   <details><summary><b>DepoDynamicXML</b></summary>
    <ul>
        <li>request: {"FromDate":"2023-06-22","ToDate":"2023-06-23"}</li>
        <li>response: {"Depo":[{"DateDepo":"2023-06-22T00:00:00+03:00","Overnight":"6.50"},{"DateDepo":"2023-06-23T00:00:00+03:00","Overnight":"6.50"}]}</li>
    </ul>
   </details>
   <details><summary><b>DragMetDynamicXML</b></summary>
    <ul>
        <li>request: {"FromDate":"2023-06-22","ToDate":"2023-06-23"}</li>
        <li>response: {"DrgMet":[{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"1","price":"5228.8000"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"2","price":"64.3800"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"3","price":"2611.0800"},{"DateMet":"2023-06-22T00:00:00+03:00","CodMet":"4","price":"3786.6100"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"1","price":"5176.2400"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"2","price":"62.0300"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"3","price":"2550.9600"},{"DateMet":"2023-06-23T00:00:00+03:00","CodMet":"4","price":"3610.0500"}]}</li>
    </ul>
   </details>
   <details><summary><b>DVXML</b></summary>
    <ul>
        <li>request: {"FromDate":"2023-06-22","ToDate":"2023-06-23"}</li>
        <li>response: {"DV":[{"Date":"2023-06-22T00:00:00+03:00","VOvern":"0.0000","VLomb":"9051.4000","VIDay":"281.3800","VOther":"504831.8300","Vol_Gold":"0.0000","VIDate":"2023-06-21T00:00:00+03:00"},{"Date":"2023-06-23T00:00:00+03:00","VOvern":"0.0000","VLomb":"8851.4000","VIDay":"118.5300","VOther":"480499.1600","Vol_Gold":"0.0000","VIDate":"2023-06-22T00:00:00+03:00"}]}</li>
    </ul>
   </details>
   <details><summary><b>EnumReutersValutesXML</b></summary>
    <ul>
        <li>request: - </li>
        <li>response: {"EnumRValutes":[{"num_code":8,"char_code":"ALL ","Title_ru":"Албанский лек","Title_en":"Albanian Lek"},{"num_code":12,"char_code":"DZD ","Title_ru":"Алжирский динар","Title_en":"Algerian Dinar"},{"num_code":32,"char_code":"ARS ","Title_ru":"Аргентинское песо","Title_en":"Argentine Peso"},{"num_code":44,"char_code":"BSD ","Title_ru":"Багамский доллар","Title_en":"Bahamian Dollar"},{"num_code":48,"char_code":"BHD ","Title_ru":"Бахрейнский динар","Title_en":"Bahraini Dinar"},{"num_code":50,"char_code":"BDT ","Title_ru":"Бангладешская така","Title_en":"Bangladeshi Taka"},{"num_code":52,"char_code":"BBD ","Title_ru":"Барбадосский доллар","Title_en":"Barbados Dollar"},{"num_code":60,"char_code":"BMD ","Title_ru":"Бермудский доллар","Title_en":"Bermudian Dollar"},{"num_code":64,"char_code":"BTN ","Title_ru":"Бутанский нгултрум","Title_en":"Bhutan Ngultrum"},{"num_code":68,"char_code":"BOB ","Title_ru":"Боливийский боливиано","Title_en":"Bolivian Boliviano"},{"num_code":72,"char_code":"BWP ","Title_ru":"Ботсванская пула","Title_en":"Botswana Pula"},{"num_code":84,"char_code":"BZD ","Title_ru":"Белизский доллар","Title_en":"Belize Dollar"},{"num_code":90,"char_code":"SBD ","Title_ru":"Доллар Соломоновых Островов","Title_en":"Solomon Is. Dollar"},{"num_code":96,"char_code":"BND ","Title_ru":"Брунейский доллар","Title_en":"Brunei Dollar"},{"num_code":108,"char_code":"BIF ","Title_ru":"Бурундийский франк","Title_en":"Burundi Franc"},{"num_code":116,"char_code":"KHR ","Title_ru":"Камбоджийский риель","Title_en":"Cambodia Riel"},{"num_code":132,"char_code":"CVE ","Title_ru":"Эскудо Кабо-Верде","Title_en":"Cabo Verde Escudo"},{"num_code":144,"char_code":"LKR ","Title_ru":"Шри-Ланкийская рупия","Title_en":"Sri Lanka Rupee"},{"num_code":152,"char_code":"CLP ","Title_ru":"Чилийское песо","Title_en":"Chilean Peso"},{"num_code":170,"char_code":"COP ","Title_ru":"Колумбийское песо","Title_en":"Colombian Peso"},{"num_code":174,"char_code":"KMF ","Title_ru":"Коморский франк","Title_en":"Comorian Franc"},{"num_code":188,"char_code":"CRC ","Title_ru":"Костариканский колон","Title_en":"Costa Rican Colon"},{"num_code":191,"char_code":"HRK ","Title_ru":"Хорватская куна","Title_en":"Croatian Kuna"},{"num_code":192,"char_code":"CUP ","Title_ru":"Кубинское песо","Title_en":"Cuban Peso"},{"num_code":214,"char_code":"DOP ","Title_ru":"Доминиканское песо","Title_en":"Dominican Peso"},{"num_code":222,"char_code":"SVC ","Title_ru":"Сальвадорский колон","Title_en":"El Salvador Colon"},{"num_code":230,"char_code":"ETB ","Title_ru":"Эфиопский быр","Title_en":"Ethiopian Birr"},{"num_code":232,"char_code":"ERN ","Title_ru":"Эритрейская накфа","Title_en":"Eritrea Nakfa"},{"num_code":238,"char_code":"FKP ","Title_ru":"Фунт Фолклендских островов","Title_en":"Falkland Islands Pound"},{"num_code":242,"char_code":"FJD ","Title_ru":"Доллар Фиджи","Title_en":"Fiji Dollar"},{"num_code":262,"char_code":"DJF ","Title_ru":"Франк Джибути","Title_en":"Djibouti Franc"},{"num_code":270,"char_code":"GMD ","Title_ru":"Гамбийский даласи","Title_en":"Gambian Dalasi"},{"num_code":292,"char_code":"GIP ","Title_ru":"Гибралтарский фунт","Title_en":"Gibraltar Pound"},{"num_code":320,"char_code":"GTQ ","Title_ru":"Гватемальский кетсаль","Title_en":"Guatemala Quetzal"},{"num_code":324,"char_code":"GNF ","Title_ru":"Гвинейский франк","Title_en":"Guinea Franc"},{"num_code":328,"char_code":"GYD ","Title_ru":"Гайанский доллар","Title_en":"Guyana Dollar"},{"num_code":332,"char_code":"HTG ","Title_ru":"Гаитский гурд","Title_en":"Haiti Gourde"},{"num_code":340,"char_code":"HNL ","Title_ru":"Гондурасская лемпира","Title_en":"Honduras Lempira"},{"num_code":344,"char_code":"HKD ","Title_ru":"Гонконгский доллар","Title_en":"Hong Kong Dollar"},{"num_code":352,"char_code":"ISK ","Title_ru":"Исландская крона","Title_en":"Iceland Krona"},{"num_code":360,"char_code":"IDR ","Title_ru":"Индонезийская рупия","Title_en":"Indonesian Rupiah"},{"num_code":364,"char_code":"IRR ","Title_ru":"Иранский риал","Title_en":"Iranian Rial"},{"num_code":368,"char_code":"IQD ","Title_ru":"Иракский динар","Title_en":"Iraqi Dinar"},{"num_code":376,"char_code":"ILS ","Title_ru":"Новый израильский шекель","Title_en":"New Israeli Sheqel"},{"num_code":388,"char_code":"JMD ","Title_ru":"Ямайский доллар","Title_en":"Jamaican Dollar"},{"num_code":400,"char_code":"JOD ","Title_ru":"Иорданский динар","Title_en":"Jordanian Dinar"},{"num_code":404,"char_code":"KES ","Title_ru":"Кенийский шиллинг","Title_en":"Kenyan Shilling"},{"num_code":408,"char_code":"KPW ","Title_ru":"Северокорейская вона","Title_en":"North Korean Won"},{"num_code":414,"char_code":"KWD ","Title_ru":"Кувейтский динар","Title_en":"Kuwaiti Dinar"},{"num_code":418,"char_code":"LAK ","Title_ru":"Лаосский кип","Title_en":"Lao Kip"},{"num_code":422,"char_code":"LBP ","Title_ru":"Ливанский фунт","Title_en":"Lebanese Pound"},{"num_code":430,"char_code":"LRD ","Title_ru":"Либерийский доллар","Title_en":"Liberian Dollar"},{"num_code":434,"char_code":"LYD ","Title_ru":"Ливийский динар","Title_en":"Libyan Dinar"},{"num_code":446,"char_code":"MOP ","Title_ru":"Патака Макао","Title_en":"Macao Pataca"},{"num_code":454,"char_code":"MWK ","Title_ru":"Малавийская квача","Title_en":"Malawi Kwacha"},{"num_code":458,"char_code":"MYR ","Title_ru":"Малайзийский ринггит","Title_en":"Malaysian Ringgit"},{"num_code":462,"char_code":"MVR ","Title_ru":"Мальдивская руфия","Title_en":"Maldives Rufiyaa"},{"num_code":478,"char_code":"MRO ","Title_ru":"Мавританская угия","Title_en":"Mauritania Ouguiya"},{"num_code":480,"char_code":"MUR ","Title_ru":"Маврикийская рупия","Title_en":"Mauritius Rupee"},{"num_code":484,"char_code":"MXN ","Title_ru":"Мексиканское песо","Title_en":"Mexican Peso"},{"num_code":496,"char_code":"MNT ","Title_ru":"Монгольский тугрик","Title_en":"Mongolia Tugrik"},{"num_code":504,"char_code":"MAD ","Title_ru":"Марокканский дирхам","Title_en":"Moroccan Dirham"},{"num_code":512,"char_code":"OMR ","Title_ru":"Оманский риал","Title_en":"Rial Omani"},{"num_code":516,"char_code":"NAD ","Title_ru":"Доллар Намибии","Title_en":"Namibia Dollar"},{"num_code":524,"char_code":"NPR ","Title_ru":"Непальская рупия","Title_en":"Nepalese Rupee"},{"num_code":533,"char_code":"AWG ","Title_ru":"Арубанский флорин","Title_en":"Aruban Florin"},{"num_code":548,"char_code":"VUV ","Title_ru":"Вануатский вату","Title_en":"Vanuatu Vatu"},{"num_code":554,"char_code":"NZD ","Title_ru":"Новозеландский доллар","Title_en":"New Zealand Dollar"},{"num_code":558,"char_code":"NIO ","Title_ru":"Никарагуанская золотая кордоба","Title_en":"Cordoba Oro"},{"num_code":566,"char_code":"NGN ","Title_ru":"Нигерийская найра","Title_en":"Nigerian Naira"},{"num_code":586,"char_code":"PKR ","Title_ru":"Пакистанская рупия","Title_en":"Pakistan Rupee"},{"num_code":590,"char_code":"PAB ","Title_ru":"Панамский бальбоа","Title_en":"Panama Balboa"},{"num_code":598,"char_code":"PGK ","Title_ru":"Кина Папуа-Новой Гвинеи","Title_en":"Papua New Guinean Kina"},{"num_code":600,"char_code":"PYG ","Title_ru":"Парагвайский гуарани","Title_en":"Paraguay Guarani"},{"num_code":604,"char_code":"PEN ","Title_ru":"Перуанский соль","Title_en":"Peru Sol"},{"num_code":608,"char_code":"PHP ","Title_ru":"Филиппинское писо","Title_en":"Philippine Piso"},{"num_code":634,"char_code":"QAR ","Title_ru":"Катарский риал","Title_en":"Qatari Rial"},{"num_code":646,"char_code":"RWF ","Title_ru":"Франк Руанды","Title_en":"Rwanda Franc"},{"num_code":654,"char_code":"SHP ","Title_ru":"Фунт Св. Елены","Title_en":"St Helena Pound"},{"num_code":678,"char_code":"STD ","Title_ru":"Добра Сан-Томе и Принсипи","Title_en":"Sao Tome \u0026 Principe Dobra"},{"num_code":682,"char_code":"SAR ","Title_ru":"Саудовский риял","Title_en":"Saudi Riyal"},{"num_code":690,"char_code":"SCR ","Title_ru":"Сейшельская рупия","Title_en":"Seychelles Rupee"},{"num_code":694,"char_code":"SLL ","Title_ru":"Сьерра-Леонский леоне","Title_en":"Sierra Leone Leone"},{"num_code":704,"char_code":"VND ","Title_ru":"Вьетнамский донг","Title_en":"Vietnam Dong"},{"num_code":706,"char_code":"SOS ","Title_ru":"Сомалийский шиллинг","Title_en":"Somali Shilling"},{"num_code":748,"char_code":"SZL ","Title_ru":"Свазилендский лилангени","Title_en":"Swaziland Lilangeni"},{"num_code":760,"char_code":"SYP ","Title_ru":"Сирийский фунт","Title_en":"Syrian Pound"},{"num_code":764,"char_code":"THB ","Title_ru":"Таиландский бат","Title_en":"Thai Baht"},{"num_code":776,"char_code":"TOP ","Title_ru":"Паанга Королевства Тонга","Title_en":"Tonga Pa'anga"},{"num_code":780,"char_code":"TTD ","Title_ru":"Доллар Тринидада и Тобаго","Title_en":"Trinidad and Tobago Dollar"},{"num_code":784,"char_code":"AED ","Title_ru":"Дирхам ОАЭ","Title_en":"UAE Dirham"},{"num_code":788,"char_code":"TND ","Title_ru":"Тунисский динар","Title_en":"Tunisian Dinar"},{"num_code":800,"char_code":"UGX ","Title_ru":"Угандийский шиллинг","Title_en":"Uganda Shilling"},{"num_code":807,"char_code":"MKD ","Title_ru":"Денар Республики Македония","Title_en":"Macedonian Denar"},{"num_code":818,"char_code":"EGP ","Title_ru":"Египетский фунт","Title_en":"Egyptian Pound"},{"num_code":834,"char_code":"TZS ","Title_ru":"Танзанийский шиллинг","Title_en":"Tanzanian Shilling"},{"num_code":858,"char_code":"UYU ","Title_ru":"Уругвайское песо","Title_en":"Peso Uruguayo"},{"num_code":886,"char_code":"YER ","Title_ru":"Йеменский риал","Title_en":"Yemeni Rial"},{"num_code":901,"char_code":"TWD ","Title_ru":"Новый тайваньский доллар","Title_en":"New Taiwan Dollar"},{"num_code":928,"char_code":"VES ","Title_ru":"Венесуэльский боливар cоберано","Title_en":"Venezuela Bolivar Soberano"},{"num_code":929,"char_code":"MRU ","Title_ru":"Мавританская угия","Title_en":"Mauritania Ouguiya"},{"num_code":930,"char_code":"STN ","Title_ru":"Добра Сан-Томе и Принсипи","Title_en":"Sao Tome \u0026 Principe Dobra"},{"num_code":936,"char_code":"GHS ","Title_ru":"Ганский седи","Title_en":"Ghana Cedi"},{"num_code":937,"char_code":"VEF ","Title_ru":"Венесуэльский боливар","Title_en":"Venezuela Bolivar"},{"num_code":938,"char_code":"SDG ","Title_ru":"Суданский фунт","Title_en":"Sudanese Pound"},{"num_code":941,"char_code":"RSD ","Title_ru":"Сербский динар","Title_en":"Serbian Dinar"},{"num_code":943,"char_code":"MZN ","Title_ru":"Мозамбикский метикал","Title_en":"Mozambique Metical"},{"num_code":950,"char_code":"XAF ","Title_ru":"Франк КФА ВЕАС","Title_en":"CFA Franc BEAC"},{"num_code":951,"char_code":"XCD ","Title_ru":"Восточно - карибский доллар","Title_en":"East Caribbean Dollar"},{"num_code":952,"char_code":"XOF ","Title_ru":"Франк КФА ВСЕАО","Title_en":"CFA Franc BCEAO"},{"num_code":967,"char_code":"ZMW ","Title_ru":"Замбийская квача","Title_en":"Zambian Kwacha"},{"num_code":968,"char_code":"SRD ","Title_ru":"Суринамский доллар","Title_en":"Surinam Dollar"},{"num_code":969,"char_code":"MGA ","Title_ru":"Малагасийский ариари","Title_en":"Malagasy Ariary"},{"num_code":971,"char_code":"AFN ","Title_ru":"Афганский афгани","Title_en":"Afghan Afghani"},{"num_code":973,"char_code":"AOA ","Title_ru":"Ангольская кванза","Title_en":"Angolan Kwanza"},{"num_code":976,"char_code":"CDF ","Title_ru":"Конголезский франк","Title_en":"Congolese Franc"},{"num_code":977,"char_code":"BAM ","Title_ru":"Конвертируемая марка","Title_en":"Convertible Mark"},{"num_code":981,"char_code":"GEL ","Title_ru":"Грузинский лари","Title_en":"Georgian Lari"}]}</li>
    </ul>
   </details>
  
