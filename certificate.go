package synologyapi

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"reflect"

	"github.com/sirateek/synologyapi/models"
	"github.com/sirupsen/logrus"
)

type certificateApi struct {
	baseApi *BaseApi
	Api     string
}

type CertificateApi interface {
	ListCertificate() error
	UploadCertificate(uploadDetail CertificateUploadDetail) error
}

func NewCertificate(baseApi *BaseApi) CertificateApi {
	return &certificateApi{
		baseApi: baseApi,
		Api:     "SYNO.Core.Certificate.CRT",
	}
}

func (c *certificateApi) ListCertificate() error {
	value := url.Values{}
	value.Add("method", "list")

	req, err := c.baseApi.GetNewHttpRequest(GET, c.Api, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery += value.Encode()

	var targetResponse models.Response[models.CertificateList]
	err = c.baseApi.SendRequest(req, &targetResponse)
	if err != nil {
		return err
	}
	return nil
}

type CertificateUploadDetail struct {
	Certificate             string `form:"cert" contentType:"application/x-x509-ca-cert" filename:"cert.cert"`
	PrivateKey              string `form:"key" contentType:"application/octet-stream" filename:"key.key"`
	IntermediateCertificate string `form:"inter_cert" omitempty:"true"`
	SetDefault              string `form:"as_default" default:"False"`
	Id                      string `form:"id"`
	Desc                    string `form:"desc"`
}

func (c *certificateApi) UploadCertificate(uploadDetail CertificateUploadDetail) error {
	value := url.Values{}
	value.Add("method", "import")

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	// Decode the struct value.
	fieldStruct := reflect.TypeOf(uploadDetail)
	valueStruct := reflect.ValueOf(uploadDetail)
	for i := 0; i < fieldStruct.NumField(); i++ {
		writerKey := fieldStruct.Field(i).Tag.Get("form")
		valueStructField := fmt.Sprint(valueStruct.Field(i))
		// Check if the value is empty first time. Set it to default value
		if valueStructField == "" {
			valueStructField = fieldStruct.Field(i).Tag.Get("default")
		}
		// If it still empty, Check the omitempty field
		if valueStructField == "" {
			omitempty := fieldStruct.Field(i).Tag.Get("omitempty")
			// If omitempty field is set, Don't add it to the writer fields
			if omitempty != "" {
				continue
			}
		}

		h := make(textproto.MIMEHeader)
		contentDisposition := fmt.Sprintf(`form-data; name="%s"`, writerKey)
		fileName := fieldStruct.Field(i).Tag.Get("filename")
		if fileName != "" {
			contentDisposition += fmt.Sprintf(`; filename="%s"`, fileName)
		}
		h.Set("Content-Disposition", contentDisposition)
		contentType := fieldStruct.Field(i).Tag.Get("contentType")
		if contentType != "" {
			h.Set("Content-Type", contentType)
		}

		ioWritter, err := writer.CreatePart(h)
		if err != nil {
			return err
		}

		ioWritter.Write([]byte(valueStructField))
	}
	writer.Close()

	req, err := c.baseApi.GetNewHttpRequest(POST, "SYNO.Core.Certificate", &body)
	if err != nil {
		return err
	}

	req.URL.RawQuery += value.Encode()
	// Headers
	req.Header.Add("Content-Type", fmt.Sprint("multipart/form-data; charset=utf-8; ", "boundary=", writer.Boundary()))

	var targetResponse models.Response[models.CertificateCreateResponse]
	err = c.baseApi.SendRequest(req, &targetResponse)
	if err != nil {
		return err
	}
	logrus.Info(targetResponse)
	return nil
}
