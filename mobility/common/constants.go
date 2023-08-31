package common

var PathPriority = struct {
	Recommend string
	Time      string
	Distance  string
}{
	Recommend: "RECOMMEND",
	Time:      "TIME",
	Distance:  "DISTANCE",
}

var CarFuel = struct {
	Gasoline string
	Diesel   string
	LPG      string
}{
	Gasoline: "GASOLINE",
	Diesel:   "DIESEL",
	LPG:      "LPG",
}

// 0	길찾기 성공
// 1	길찾기 결과를 찾을 수 없음
// 101	경유지 지점 주변의 도로를 탐색할 수 없음
// 102	시작 지점 주변의 도로를 탐색할 수 없음
// 103	도착 지점 주변의 도로를 탐색할 수 없음
// 104	출발지와 도착지가 5 m 이내로 설정된 경우 경로를 탐색할 수 없음
// 105	시작 지점 주변의 도로에 유고 정보(교통 장애)가 있음
// 106	도착 지점 주변의 도로에 유고 정보(교통 장애)가 있음
// 107	경유지 주변의 도로에 유고 정보(교통 장애)가 있음.
// result_message에 경유지의 순번이 표시되며 번호는 1번부터 시작함
//
// 예시)
// result_code: 107
// result_message: 경유지에 유고 정보 존재: 1번 경유지
// 201	다중 출발지: 출발지가 탐색 영역에 포함되지 않음
// 202	다중 출발지: 출발지 최대 개수 초과 도로 선택 실패
// 203	다중 출발지: 목적지 도로 선택 실패
// 204	다중 출발지: 경로 탐색 처리 시간 제한
// 205	다중 출발지: 출발지 주변의 유고 정보(교통 장애)로 인한 통행 불가
// 206	다중 출발지: 목적지 주변의 유고 정보(교통 장애)로 인한 통행 불가
// 207	다중 출발지: 출발지가 설정한 길찾기 반경 범위를 벗어남
// 301	다중 목적지: 출발지 도로 선택 실패
// 302	다중 목적지: 목적지 도로 선택 실패
// 303	다중 목적지: 목적지 최대 개수 초과로 인해 경로 탐색 실패
// 304	다중 목적지: 목적지가 설정한 길찾기 반경 범위를 벗어남
var NavigationResultCode = struct {
	Success string
}{
	Success: "0",
}
