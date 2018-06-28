package cmd

import (
	"net/http"
	"net/http/httptest"
	"net/rpc"

	"github.com/qri-io/qri/config"
	"github.com/qri-io/qri/lib"
	"github.com/qri-io/qri/p2p"
	"github.com/qri-io/qri/repo"
	"github.com/qri-io/qri/repo/test"
	"github.com/qri-io/registry/regclient"
)

// TestFactory is an implementation of the Factory interface for testing purposes
type TestFactory struct {
	IOStreams
	// QriRepoPath is the path to the QRI repository
	qriRepoPath string
	// IpfsFsPath is the path to the IPFS repo
	ipfsFsPath string

	// Configuration object
	config *config.Config
	node   *p2p.QriNode
	repo   repo.Repo
	rpc    *rpc.Client
}

// NewTestFactory creates TestFactory object with an in memory test repo
func NewTestFactory() (tf TestFactory, err error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockResponse)
	}))
	rc := regclient.NewClient(&regclient.Config{Location: server.URL})
	repo, err := test.NewTestRepo(rc)
	if err != nil {
		return
	}

	cfg := config.DefaultConfig()

	return TestFactory{
		qriRepoPath: "",
		ipfsFsPath:  "",

		repo:   repo,
		rpc:    nil,
		config: cfg,
		node:   nil,
	}, nil
}

// Config returns from internal state
func (t TestFactory) Config() (*config.Config, error) {
	return t.config, nil
}

// IpfsFsPath returns from internal state
func (t TestFactory) IpfsFsPath() string {
	return t.ipfsFsPath
}

// QriRepoPath returns from internal state
func (t TestFactory) QriRepoPath() string {
	return t.qriRepoPath
}

// Repo returns from internal state
func (t TestFactory) Repo() (repo.Repo, error) {
	return t.repo, nil
}

// RPC returns from internal state
func (t TestFactory) RPC() *rpc.Client {
	return nil
}

// DatasetRequests generates a lib.DatasetRequests from internal state
func (t TestFactory) DatasetRequests() (*lib.DatasetRequests, error) {
	return lib.NewDatasetRequests(t.repo, t.rpc), nil
}

// RegistryRequests generates a lib.RegistryRequests from internal state
func (t TestFactory) RegistryRequests() (*lib.RegistryRequests, error) {
	return lib.NewRegistryRequests(t.repo, t.rpc), nil
}

// HistoryRequests generates a lib.HistoryRequests from internal state
func (t TestFactory) HistoryRequests() (*lib.HistoryRequests, error) {
	return lib.NewHistoryRequests(t.repo, t.rpc), nil
}

// PeerRequests generates a lib.PeerRequests from internal state
func (t TestFactory) PeerRequests() (*lib.PeerRequests, error) {
	return lib.NewPeerRequests(nil, t.rpc), nil
}

// ProfileRequests generates a lib.ProfileRequests from internal state
func (t TestFactory) ProfileRequests() (*lib.ProfileRequests, error) {
	return lib.NewProfileRequests(t.repo, t.rpc), nil
}

// SelectionRequests creates a lib.SelectionRequests from internal state
func (t TestFactory) SelectionRequests() (*lib.SelectionRequests, error) {
	return lib.NewSelectionRequests(t.repo, t.rpc), nil
}

// SearchRequests generates a lib.SearchRequests from internal state
func (t TestFactory) SearchRequests() (*lib.SearchRequests, error) {
	return lib.NewSearchRequests(t.repo, t.rpc), nil
}

// RenderRequests generates a lib.RenderRequests from internal state
func (t TestFactory) RenderRequests() (*lib.RenderRequests, error) {
	return lib.NewRenderRequests(t.repo, t.rpc), nil
}

var mockResponse = []byte(`{"data":[{"Type":"","ID":"ramfox/test","Value":null},{"Type":"","ID":"b5/test","Value":{"commit":{"path":"/","signature":"JI/VSNqMuFGYVEwm3n8ZMjZmey+W2mhkD5if2337wDp+kaYfek9DntOyZiILXocW5JuOp48EqcsWf/BwhejQiYZ2utaIzR8VcMPo7u7c5nz2G6JTsoW+u9UUaKRVtl30jh6Kg1O2bnhYh9v4qW9VQgxOYfdhBl6zT4cYcjm1UkrblEe/wh494k9NziM5Bi2ATGRE2m71Lsf/TEDoNI549SebLQ1dsWXr1kM7lCeFqlDVjgbKQmGXowqcK/P9v+RBIRCnArnwFe/BQq4i1wmmnMEqpuUnfWR3xfJTE1DUMVaAid7U0jTWGVxROUdKk6mmTzlb1PiNdfruP+SFhjyQwQ==","timestamp":"2018-05-25T13:44:54.692493401Z","title":"created dataset"},"meta":{"citations":[{"url":"https://api.github.com/repos/qri-io/qri/releases"}],"qri":"md:0"},"path":"/ipfs/QmPi5wrPsY4xPwy2oRr7NRZyfFxTeupfmnrVDubzoABLNP","qri":"","structure":{"checksum":"QmQXfdYYubCvr9ePJtgABpp7N1fNsAnnywUJPYAJTvDhrn","errCount":0,"entries":7,"format":"json","length":19116,"qri":"","schema":{"type":"array"}},"transform":{"config":{"org":"qri-io","repo":"qri"},"path":"/","syntax":"skylark"},"Handle":"b5","Name":"test","PublicKey":"CAASpgIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC/W17VPFX+pjyM1MHI1CsvIe+JYQX45MJNITTd7hZZDX2rRWVavGXhsccmVDGU6ubeN3t6ewcBlgCxvyewwKhmZKCAs3/0xNGKXK/YMyZpRVjTWw9yPU9gOzjd9GuNJtL7d1Hl7dPt9oECa7WBCh0W9u2IoHTda4g8B2mK92awLOZTjXeA7vbhKKX+QVHKDxEI0U2/ooLYJUVxEoHRc+DUYNPahX5qRgJ1ZDP4ep1RRRoZR+HjGhwgJP+IwnAnO5cRCWUbZvE1UBJUZDvYMqW3QvDp+TtXwqUWVvt69tp8EnlBgfyXU91A58IEQtLgZ7klOzdSEJDP+S8HIwhG/vbTAgMBAAE="}},{"Type":"","ID":"EDGI/fib_6","Value":{"commit":{"path":"/","signature":"dG6NoEFlQ9ILFjVDrecSDlUbDPRSiwK9kFQ3vjTh/4tpfgrT5EOw6eGDx75lklx0DWx51s3AC2Qqytll2JwwCB6SMVVl0I9qnZJ4XQVG+MX4hGeIJ4crGCSts85unDvmiQCfc4EVqPYZKLzaVqjXa43zv5mJPfRA2ktTew3VGkOcg8RmyU6e2XWrZpkILcZg1jt2apKs5qHslx4klKBVPtQIr53/U61OW4tzME9kz08FYrk2F7I5FHWy45W7VU8DpzCbhw6kxJXu2KYD1QstsZGCKH93sZY3agP4XGY15HeEOTib465LK6+nsoBtrsroQSOTBHzVgyUZACNom5SUvQ==","timestamp":"2018-05-23T19:50:03.307982846Z","title":"created dataset"},"meta":{"description":"test of skylark as a transformation language","qri":"md:0","title":"Fibonacci(6)"},"path":"/ipfs/QmS6jJSEJYxZvCeo8cZqzVa7Ybu9yNQeFYfNZAHxM4eyDK","qri":"","structure":{"checksum":"QmdhDDZTAWoifKCWwFrZqzKUyXM6rN7ThdP1u67rkeJvTj","errCount":0,"entries":6,"format":"cbor","length":7,"qri":"","schema":{"type":"array"}},"transform":{"syntax":"skylark"},"Handle":"EDGI","Name":"fib_6","PublicKey":"CAASpgIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCmTFRx/6dKmoxje8AG+jFv94IcGUGnjrupa7XEr12J/c4ZLn3aPrD8F0tjRbstt1y/J+bO7Qb69DGiu2iSIqyE21nl2oex5+14jtxbupRq9jRTbpUHRj+y9I7uUDwl0E2FS1IQpBBfEGzDPIBVavxbhguC3O3XA7Aq7vea2lpJ1tWpr0GDRYSNmJAybkHS6k7dz1eVXFK+JE8FGFJi/AThQZKWRijvWFdlZvb8RyNFRHzpbr9fh38bRMTqhZpw/YGO5Ly8PNSiOOE4Y5cNUHLEYwG2/lpT4l53iKScsaOazlRkJ6NmkM1il7riCa55fcIAQZDtaAx+CT5ZKfmek4P5AgMBAAE="}}],"meta":{"code":200}}`)