package scribe

import (
	"errors"
	"net"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/cloudaice/scribe-go/facebook/scribe"
)

type ScribeLoger struct {
	transport *thrift.TFramedTransport
	client    *scribe.ScribeClient
}

func NewScribeLoger(host, port string) (*ScribeLoger, error) {
	Ttransport, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		return nil, err
	}
	transport := thrift.NewTFramedTransport(Ttransport)

	protocol := thrift.NewTBinaryProtocol(transport, false, false)

	client := scribe.NewScribeClientProtocol(transport, protocol, protocol)
	if err := transport.Open(); err != nil {
		return nil, err
	}
	return &ScribeLoger{
		transport: transport,
		client:    client,
	}, nil
}

func (this *ScribeLoger) SendOne(category, message string) (bool, error) {
	logEntry := &scribe.LogEntry{category, message}
	result, err := this.client.Log([]*scribe.LogEntry{logEntry})
	if err != nil {
		return false, err
	}
	return this.dealResult(result)
}

func (this *ScribeLoger) SendArray(category string, messages []string) (bool, error) {
	logEntrys := []*scribe.LogEntry{}

	for _, message := range messages {
		logEntry := &scribe.LogEntry{category, message}
		logEntrys = append(logEntrys, logEntry)
	}
	result, err := this.client.Log(logEntrys)
	if err != nil {
		return false, err
	}
	return this.dealResult(result)
}

func (this *ScribeLoger) dealResult(result scribe.ResultCode) (bool, error) {
	ok := false
	var err error
	switch result {
	case scribe.ResultCode_OK:
		ok = true
	case scribe.ResultCode_TRY_LATER:
		ok = false
	default:
		err = errors.New(result.String())
	}
	return ok, err
}

func (this *ScribeLoger) Close() error {
	return this.transport.Close()
}
