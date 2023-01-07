package synologyapi

import (
	"net/url"

	"github.com/sirateek/synologyapi/models"
	"github.com/sirupsen/logrus"
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
	value.Add("version", "1")
	value.Add("method", "list")
	value.Add("api", c.Api)

	req, err := c.GetNewHttpRequest(GET, c.Api)
	if err != nil {
		return err
	}
	req.URL.RawQuery = value.Encode()

	var targetResponse models.Response[models.CertificateList]
	err = c.SendRequest(req, &targetResponse)
	if err != nil {
		return err
	}
	logrus.Info(targetResponse)
	return nil
}
