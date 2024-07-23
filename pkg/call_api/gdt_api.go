package call_api

type Gdt struct {
	apiBaseUrl string
}

func NewGdtCallInstance() (gdt *Gdt) {
	gdt = &Gdt{}

	return
}

func (gdt *Gdt) Call() (err error) {
	//resp, err := resty.New().R().Get("https://httpbin.org/get")
	//if err != nil {
	//	return
	//}
	return
}
