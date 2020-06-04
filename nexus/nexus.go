package nexus

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/apex/log"
)

type Session struct {
	Credential *Credential
	HttpClient *http.Client
}

type Credential struct {
	Username    string
	Userpass    string
	Certificate string
}

type Artifact struct {
	Url         string
	File        string
	ContentType string
}

func New(creds Credential) (*Session, error) {
	var httpClient *http.Client

	cer := creds.Certificate
	if cer == "" {
		cer = certsDefault
	}
	pool, err := x509.SystemCertPool()
	if err != nil {
		if runtime.GOOS == "windows" {
			// on windows ignore errors
			pool = x509.NewCertPool()
		} else {
			return nil, err
		}
	}
	pool.AppendCertsFromPEM([]byte(cer))
	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
			},
		},
	}

	return &Session{
		HttpClient: httpClient,
		Credential: &Credential{
			Username:    creds.Username,
			Userpass:    creds.Userpass,
			Certificate: creds.Certificate,
		},
	}, nil
}

func (session *Session) PutArtifact(artifact Artifact) error {
	f, err := os.Open(artifact.File)
	if err != nil {
		return err
	}
	defer f.Close()

	creds := session.Credential
	var headers = map[string]string{}
	if artifact.ContentType != "" {
		headers["Content-Type"] = artifact.ContentType
	}
	req, err := request(artifact.Url, http.MethodPut, f, creds.Username, creds.Userpass, headers)
	if err != nil {
		return err
	}

	// log.Debugf("executing request: %s %s (headers: %v)", req.Method, req.URL, req.Header)
	response, err := session.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("fail PUT status '%s'", response.Status)
	} else {
		log.Infof("uploaded successful %s", artifact.File)
	}

	return nil
}

func request(url, method string, data io.Reader, username, userpass string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	//req.ContentLength = a.Size
	if username != "" && userpass != "" {
		req.SetBasicAuth(username, userpass)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

func (session *Session) GetArtifact(artifact Artifact) error {
	creds := session.Credential
	req, err := request(artifact.Url, http.MethodGet, nil, creds.Username, creds.Userpass, map[string]string{})
	if err != nil {
		return err
	}

	response, err := session.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// log.Debugf("Fail GET '%s'", response.Status)
		return fmt.Errorf("Fail GET status '%s'", response.Status)
	} else {
		log.Infof("get successful - ContentLength %d", response.ContentLength)
	}

	return nil
}
