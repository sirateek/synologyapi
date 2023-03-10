package models

type CertificateList struct {
	Certificates []Certificate `json:"certificates"`
}

type Certificate struct {
	Desc               string               `json:"desc"`
	ID                 string               `json:"id"`
	IsBroken           bool                 `json:"is_broken"`
	IsDefault          bool                 `json:"is_default"`
	Issuer             CertificateIssuer    `json:"issuer"`
	KeyTypes           string               `json:"key_types"`
	Renewable          bool                 `json:"renewable"`
	Services           []CertificateService `json:"services"`
	SignatureAlgorithm string               `json:"signature_algorithm"`
	Subject            CertificateSubject   `json:"subject"`
	UserDeletable      bool                 `json:"user_deletable"`
	ValidFrom          string               `json:"valid_from"`
	ValidTill          string               `json:"valid_till"`
}

type CertificateSubject struct {
	City         string   `json:"city"`
	CommonName   string   `json:"common_name"`
	Country      string   `json:"country"`
	Organization string   `json:"organization"`
	SubAltName   []string `json:"sub_alt_name"`
}

type CertificateIssuer struct {
	CommonName   string `json:"common_name"`
	Country      string `json:"country"`
	Organization string `json:"organization"`
}

type CertificateService struct {
	DisplayName     string `json:"display_name"`
	DisplayNamei18n string `json:"display_name_i18n"`
	IsPkg           bool   `json:"isPkg"`
	MultipleCert    bool   `json:"multiple_cert"`
	Owner           string `json:"owner"`
	Service         string `json:"service"`
	Subscriber      string `json:"subscriber"`
	UserSetable     bool   `json:"user_setable"`
}

type CertificateCreateResponse struct {
	Id string `json:"id"`
}
