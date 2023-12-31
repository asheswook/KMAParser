package stations

import "errors"

var (
	errInvaildStationName = errors.New("The station name is invalid")
	errInvaildStationNumber = errors.New("The station number is invalid")
)

// 지점명 - 지점번호 매핑
var stations = map[string]string{
	"충주": "47127",
	"서산": "47129",
	"울진": "47130",
	"청주": "47131",
	"대전": "47133",
	"추풍령": "47135",
	"안동": "47136",
	"상주": "47137",
	"포항": "47138",
	"군산": "47140",
	"대구": "47143",
	"전주": "47146",
	"울산": "47152",
	"창원": "47155",
	"광주": "47156",
	"부산": "47159",
	"통영": "47162",
	"목포": "47165",
	"여수": "47168",
	"흑산도": "47169",
	"완도": "47170",
	"순천": "47174",
	"제주": "47184",
	"고산": "47185",
	"국가태풍센터": "47186",
	"성산": "47188",
	"서귀포": "47189",
	"진주": "47192",
	"강화": "47201",
	"양평": "47202",
	"인천1": "47203",
	"인제": "47211",
	"홍천": "47212",
	"태백": "47216",
	"제천": "47221",
	"보은": "47226",
	"천안": "47232",
	"보령": "47235",
	"부여": "47236",
	"금산": "47238",
	"부안": "47243",
	"임실": "47244",
	"정읍": "47245",
	"남원": "47247",
	"장수": "47248",
	"장흥": "47260",
	"속초": "47090",
	"철원": "47095",
	"동두천": "47098",
	"파주": "47099",
	"대관령": "47100",
	"춘천": "47101",
	"백령도": "47102",
	"북강릉": "47104",
	"강릉": "47105",
	"동해": "47106",
	"서울": "47108",
	"인천2": "47112",
	"원주": "47114",
	"울릉도": "47115",
	"수원": "47119",
	"영월": "47121",
	"오산": "47122",
	"해남": "47261",
	"고흥": "47262",
	"봉화": "47271",
	"영주": "47272",
	"문경": "47273",
	"영덕": "47277",
	"의성": "47278",
	"구미": "47279",
	"영천": "47281",
	"거창": "47284",
	"합천": "47285",
	"밀양": "47288",
	"산청": "47289",
	"거제": "47294",
	"남해": "47295",
}

func GetName(stnCode string) (string, error) {
	for name, code := range stations {
		if code == stnCode {
			return name, nil
		}
	}
	return "", errInvaildStationNumber
}

func GetCode(stnName string) (string, error) {
	value, ok := stations[stnName]
	if ok {
		return value, nil
	}
	return "", errInvaildStationName
}