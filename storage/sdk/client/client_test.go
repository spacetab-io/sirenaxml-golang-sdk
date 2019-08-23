package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/sirenaxml-golang-sdk/configuration"
	"github.com/tmconsulting/sirenaxml-golang-sdk/strings"
)

var (
	sc sirenaXML.Config
)

func tearUp() {
	clientID, _ := strings.String2Uint16(os.Getenv("CLIENT_ID"))
	requestHandlersNum, _ := strings.String2Int32(os.Getenv("MAX_CONNECTIONS"))

	sc = sirenaXML.Config{
		ClientID:                 clientID,
		Environment:              os.Getenv("ENV"),
		Ip:                       os.Getenv("IP"),
		MaxConnections:           requestHandlersNum,
		ClientPublicKey:          os.Getenv("CLIENT_PUBLIC_KEY"),
		ClientPrivateKey:         os.Getenv("CLIENT_PRIVATE_KEY"),
		ServerPublicKey:          os.Getenv("SERVER_PUBLIC_KEY"),
		ClientPrivateKeyPassword: os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD"),
		ZippedMessaging:          true,
		MaxConnectTries:          3,
	}
	err := sc.PrepareKeys()
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	tearUp()
	addr, _ := sc.GetAddr()
	//if err != nil {
	//	return nil, err
	//}
	//logger := logs.NewNullLog()
	logger := logrus.New()
	//t.Run("success", func(t *testing.T) {
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c, err := New(
		SetKeys(sc.GetKeys()),
		SetAddr(addr),
		SetLogger(logger),
		SetUseZip(sc.ZippedMessaging),
		SetClientID(sc.ClientID),
		SetMaxConnections(sc.MaxConnections),
	)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	errCh := make(chan error)
	go func() {
		errCh <- c.Connect(ctx)
	}()
	go func() {
		for {
			select {
			case <-errCh:
				ctx.Done()
			}
		}
	}()
	err = c.WaitKeySign(ctx)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	kd := c.GetKeyData()
	if !assert.NotZero(t, kd.ID) {
		t.FailNow()
	}

	//reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
	//if !assert.NoError(t, err) {
	//	t.FailNow()
	//}
	//respData, err := c.SendMsg(ctx, reqXML)
	//if !assert.NoError(t, err) {
	//	t.FailNow()
	//}
	//var resp structs.KeyInfoResponse
	//err = xml.Unmarshal(respData, &resp)
	//if !assert.NoError(t, err) {
	//	t.FailNow()
	//}

	//})

	//t.Run("throttling", func(t *testing.T) {
	//	c, err := New(&sc, logger)
	//	if !assert.NoError(t, err) {
	//		t.FailNow()
	//	}
	//	assert.Panics(t, func() {
	//		c.recover()
	//	})
	//})
}

//func TestChannel_Reconnect(t *testing.T) {
//	tearUp()
//	//logger := logs.NewNullLog()
//	logger := logrus.New()
//	t.Run("reconnect", func(t *testing.T) {
//		ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
//		c, err := New(&sc, SetLogger(logger))
//		if !assert.NoError(t, err, "check channel initialisation success") {
//			t.FailNow()
//		}
//		err = c.Connect(ctx)
//		if !assert.NoError(t, err, "check connect success") {
//			t.FailNow()
//		}
//
//		if !assert.NotEmpty(t, c.conn.KeyType.ID, "check key data is set correctly") {
//			t.FailNow()
//		}
//		// setting bad data to emulate symmetric key expiration
//		badKeyData := message.KeyType{
//			ID:  1917653,
//			Key: []byte("Kp4MXPqs"),
//		}
//		c.conn.KeyType = badKeyData
//
//		// try to get some data
//		reqXML, err := xml.Marshal(&structs.KeyInfoRequest{})
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		resp, err := c.SendMsg(ctx, reqXML)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		if !assert.Nil(t, resp, string(resp)) {
//			t.FailNow()
//		}
//
//		time.Sleep(2 * time.Second)
//
//		if !assert.NotEqual(t, badKeyData.Key, c.conn.KeyType.Key) {
//			t.FailNow()
//		}
//		// try to get some data again
//		resp, err = c.SendMsg(ctx, reqXML)
//		if !assert.NoError(t, err) {
//			t.FailNow()
//		}
//		if !assert.NotNil(t, resp) {
//			t.FailNow()
//		}
//	})
//}
