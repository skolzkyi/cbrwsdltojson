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
  
