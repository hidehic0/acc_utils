package types

type ContestInfomation struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Directory struct {
	Path    string `json:"path"`
	TestDir string `json:"testdir"`
	Submit  string `json:"submit"`
}

type TaskInfomation struct {
	Id        string    `json:"id"`
	Lable     string    `json:"label"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Directory Directory `json:"directory"`
}

type Infomation struct {
	Contest ContestInfomation `json:"contest"`
	Tasks   []TaskInfomation  `json:"tasks"`
}
