package selfhealing

type SHProcesses struct{
	First string `json:"first"`
	Href string `json:"href"`
	Last string `json:"last"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	next string `json:"offset"`
	Processes []SHProcess `json:"processes"`
	Total int `json:"total_count"`
}

type SHProcess struct{
	Href string `json:"href"`
	Id string `json:"id"`
	Status string `json:"status"`
	Threads HrefStruct `json:"threads"`
	Type string `json:"type"`
}

type HrefStruct struct{
	Href string `json:"href"`
}

type PNPErrors struct{
	Errors []PNPError `json:"errors"`
}

type PNPError struct{
	Code string `json:"code"`
	Detail string `json:"detail"`
	Message string `json:"message"`
}

type SHThreads struct{
	First string `json:"first"`
	Href string `json:"href"`
	Last string `json:"last"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	next string `json:"offset"`
	Threads []SHThread `json:"processes"`
	Total int `json:"total_count"`
}

type SHThread struct{
	Href string `json:"href"`
	Id string `json:"id"`
	Status string `json:"status"`
	Logs HrefStruct `json:"logs"`
}

type LogContent struct {
	Content string `json:"content"`
}