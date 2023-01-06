package synologyapi

type certificate struct {
	BaseApi
	Api     string
	cgiPath string
}

type Certificate interface {
}

func NewCertificate(baseApi BaseApi) Certificate {
	return *&certificate{
		BaseApi: baseApi,
		Api:     "SYNO.Core.Certificate.CRT",
		cgiPath: "",
	}
}
