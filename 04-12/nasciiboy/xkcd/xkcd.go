package xkcd

const (
  baseName = "https://xkcd.com"
  jsonRequest = "/info.0.json"
)

type Comic struct {
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	ImageURL   string `json:"img"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
}

type Comics []Comic

type dbInfo struct {
  Init  int
  Last  int
  OK    int
  Fail []*FailRequest `json:"fail_request"`
}

type FailRequest struct {
  Num  int
  Json bool
  Img  bool
}
