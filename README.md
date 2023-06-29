# KMAParser

기상청 고층관측 - 연직 바람 관측 장비 (Windprofiler) API의 데이터 parser입니다.

## Usage

```go
code, _ := kmaparser.stations.GetCode("창원") // 창원 지역의 관측소 코드를 가져옵니다.

p := kmaparser.NewParams("", code, "L", true) // Time, Station, Mode, Help(결과와 관계 없음)
p.SetAuthKey("API KEY")

resp, err := GetWindProfilerRaderResp(*p)

result, err := ParseWindProfilerRaderResp(resp)
```

결과값은 아래와 같이 나타납니다.

```go
type WindProfilerRaderDataset struct {
	Info Info
	Dataset []WindProfilerRaderData
}

type WindProfilerRaderData struct {
	Ht string // 높이 (m)
	Wd string // 풍향 (deg)
	Ws string // 풍속 (m/s)
	U string // U 동서 바람 성분 (m/s)
	V string // V 남북 바람 성분 (m/s)
	W string // W 연직 바람 성분 (m/s)
}

result.Info.Time // 관측 시간
result.Info.Station // 관측소

result.Dataset // 관측 데이터 Array
```
