package synologyapi

import (
	"mime/multipart"
	"net/textproto"
	"net/url"

	"github.com/sirateek/synologyapi/models"
)

type certificateApi struct {
	BaseApi
	Api string
}

type CertificateApi interface {
	ListCertificate() error
}

func NewCertificate(baseApi *BaseApi) CertificateApi {
	return &certificateApi{
		BaseApi: *baseApi,
		Api:     "SYNO.Core.Certificate.CRT",
	}
}

func (c *certificateApi) ListCertificate() error {
	value := url.Values{}
	value.Add("method", "list")
	value.Add("api", c.Api)

	req, err := c.GetNewHttpRequest(GET, c.Api)
	if err != nil {
		return err
	}
	req.URL.RawQuery += value.Encode()

	var targetResponse models.Response[models.CertificateList]
	err = c.SendRequest(req, &targetResponse)
	if err != nil {
		return err
	}
	return nil
}

func (c *certificateApi) UploadAndReplaceExistingCertificate(privateKey []byte, certificate []byte, intermediateCertificate []byte, certId string) error {

	return nil
}

func (c *certificateApi) UploadAndCreateNewCertificate(privateKey []byte, certificate []byte, intermediateCertificate []byte, setDefault bool) error {
	value := url.Values{}
	value.Add("method", "import")
	value.Add("api", "SYNO.Core.Certificate")

	req, err := c.GetNewHttpRequest(POST, c.Api)
	if err != nil {
		return err
	}

	req.ParseMultipartForm(10 << 20)
	req.FormFile("key")
	req.MultipartForm.File["key"] = []*multipart.FileHeader{{
		Header: textproto.MIMEHeader{},
	}}

	return nil
}
