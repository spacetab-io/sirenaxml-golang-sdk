package client

import (
	"io"
	"log"
	"time"

	"github.com/pkg/errors"

	"github.com/tmconsulting/sirenaxml-golang-sdk/logs"
)

const MaxTries = 3

func (client *SirenaClient) sendToSocket(request *Request) error {
	client.socketWriteMux.Lock()
	logs.Log.Debugf("[%d] sending message to socket", request.Header.MessageID)
	var data []byte
	data = append(data, request.Header.ToBytes()...)
	if len(request.SubHeader) > 0 {
		data = append(data, request.SubHeader...)
	}
	data = append(data, request.Message...)
	if len(request.MessageSignature) > 0 {
		data = append(data, request.MessageSignature...)
	}
	if _, err := client.Conn.Write(data); err != nil {
		client.socketWriteMux.Unlock()
		return errors.Wrapf(err, errorFormat, "request write error")
	}

	client.socketWriteMux.Unlock()
	return nil
}

func (client *SirenaClient) getMessageFromSocket(i int) {
	logs.Log.Debugf("[%d] getting socket answer", i)
	//err := client.Conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	//if err != nil {
	//	client.errChan[ReadingSocketError] <- &RequestError{
	//		Error: errors.Wrap(err, "SetReadDeadline error"),
	//	}
	//	//continue
	//	return
	//
	//}

	logs.Log.Debug("getting response header")
	//connReader := bufio.NewReader(client.Conn)
	responseHeaderBytes := make([]byte, 100)
	if _, err := client.Conn.Read(responseHeaderBytes); err != nil {
		logs.Log.Debugf("ошибка получения заголовка ответа: %v", err)
		//client.errChan[ReadingSocketError] <- &RequestError{
		//	Error: errors.Wrapf(err, errorFormat, "response header read error"),
		//}
		//continue
		return
	}
	logs.Log.Debug("parsing response header")
	responseHeader := ParseHeader(responseHeaderBytes)
	if responseHeader.MessageLength == 0 {
		log.Fatalf("[%d] пустое тело запроса", responseHeader.MessageID)
	}

	if responseHeader.RequestNoHandled {
		logs.Log.Warnf("[%d] не обработался запрос. Запрос существует? %v", responseHeader.MessageID, client.respChanExists(responseHeader.MessageID))
		if client.respChanExists(responseHeader.MessageID) {
			go client.resendRequest(responseHeader.MessageID)
			return
		}
	}

	//
	//if err != nil {
	//	logs.Log.Fatalf(err.Error())
	//	client.errChan[ReadingSocketError] <- &RequestError{
	//		Error: errors.Wrapf(err, errorFormat, "response header parse error"),
	//	}
	//	return
	//}

	if ok := client.respChanExists(responseHeader.MessageID); !ok {
		logs.Log.Fatalf("еблысь, тебя налево! наличие respChan: %v", ok)
		if !client.waitHasNoResult(responseHeader.MessageID) {
			return
		}
	}

	if responseHeader.MessageID == 0 {
		logs.Log.Fatal("messageID нет совсем. Пичаль...")
	}

	responseMessageBytes := client.getResponseMessage(responseHeader)

	logs.Log.Debugf("[%d] sending response message to response channel", responseHeader.MessageID)
	client.sendResponseToRespChan(responseHeader, responseMessageBytes)
	logs.Log.Debug("done! once again!")
}

func (client *SirenaClient) waitHasNoResult(msgID uint32) bool {
	for i := 1; i < MaxTries; i++ {
		time.Sleep(time.Duration(i*10) * time.Millisecond)
		if client.respChanExists(msgID) {
			return true
		}
	}
	return false
}

func (client *SirenaClient) getResponseMessage(responseHeader *Header) []byte {
	logs.Log.Debugf("[%d] getting response message with length %d", responseHeader.MessageID, responseHeader.MessageLength)
	responseMessageBytes := make([]byte, responseHeader.MessageLength)
	if _, err := io.ReadFull(client, responseMessageBytes); err != nil {
		if client.respChanExists(responseHeader.MessageID) && responseHeader.MessageID != 0 {
			client.errChan[responseHeader.MessageID] <- &RequestError{
				MessageID: responseHeader.MessageID,
				Error:     errors.Wrapf(err, errorFormat, "response read error"),
			}
		} else {
			client.errChan[ReadingSocketError] <- &RequestError{
				MessageID: responseHeader.MessageID,
				Error:     errors.Wrapf(err, errorFormat, "response read error"),
			}
		}
		//return
	}
	return responseMessageBytes
}

func (client *SirenaClient) respChanExists(msgID uint32) bool {
	client.respMux.RLock()
	_, ok := client.respChan[msgID]
	client.respMux.RUnlock()
	return ok
}

func (client *SirenaClient) sendResponseToRespChan(responseHeader *Header, responseMessageBytes []byte) {
	client.respMux.Lock()
	client.respChan[responseHeader.MessageID] <- &Response{
		MessageID: responseHeader.MessageID,
		Message:   responseMessageBytes,
		Header:    responseHeader,
	}
	client.respMux.Unlock()
}

func (client *SirenaClient) listenSocketContinuously() {
	logs.Log.Debugf("start listening socket continuously")
	go func() {
		var i = 0
		for {
			//time.Sleep(5*time.Millisecond)
			client.getMessageFromSocket(i)
			i++
		}
	}()

	select {
	case err := <-client.errChan[ReadingSocketError]:
		logs.Log.WithError(err.Error).WithField("messageID", err.MessageID).Fatal(err.Message)
	}
}
