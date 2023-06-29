package kmaparser

import (
	"fmt"
	"strings"
)

type Info struct {
	Time string
	Station string
}

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

var (
	errInvaildResp = fmt.Errorf("The parser received an invalid response")
	errNoSameTimeStation = fmt.Errorf("The parser received an invalid response. - Exists different time or station in the one response")
)

func ParseWindProfilerRaderResp(resp string) (WindProfilerRaderDataset, error) {
	// Check response is valid
	if strings.Index(resp, "#") != 0 {
		return WindProfilerRaderDataset{}, errInvaildResp
	}

	dataset := []WindProfilerRaderData{}
	time := ""
	station := ""
	info := Info{time, station}
	result := WindProfilerRaderDataset{info, dataset}

	for _, line := range strings.Split(resp, "\n") {
		// 주석은 무시
		if strings.Index(line, "#") == 0 || len(line) == 0 {
			continue
		}

		data := strings.Fields(line)

		// 시간과 지점을 저장
		if time == "" && station == "" {
			time = data[0]
			station = data[1]
			result.Info.Time = time
			result.Info.Station = station
		}

		// 시간과 지점이 같지 않으면 에러
		if time != "" && station != "" && (time != data[0] || station != data[1]) {
			return result, errNoSameTimeStation
		}

		// 데이터 저장
		dataset = append(dataset, WindProfilerRaderData{
			Ht: data[2],
			Wd: data[3],
			Ws: data[4],
			U: data[5],
			V: data[6],
			W: data[7],
		})
	}

	result.Dataset = dataset
	
	return result, nil
}

