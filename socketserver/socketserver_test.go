package socketserver

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net"
	"strconv"
	"testing"
)

var done = make(chan bool)

type TestSocketServicer interface {
	doService(c net.Conn)
}

type TestSocketService struct {
	TestSocketServicer
}

func (tss *TestSocketService) serve(c net.Conn) {
	defer c.Close()
	tss.doService(c)
	done <- true
}

func NewTestSocketService(servicer TestSocketServicer) *TestSocketService {
	return &TestSocketService{servicer}
}

type ClosingSocketService struct {
	*TestSocketService
	connections int
}

func (css *ClosingSocketService) doService(c net.Conn) {
	css.connections++
}

func NewClosingSocketService() *ClosingSocketService {
	closingService := new(ClosingSocketService)
	closingService.TestSocketService = NewTestSocketService(closingService)
	return closingService
}

type SocketServerSuite struct {
	suite.Suite
	server  *SocketServer
	service *ClosingSocketService
	port    int
}

func (suite *SocketServerSuite) SetupTest() {
	suite.port = 8042
	suite.service = NewClosingSocketService()
	var err error
	suite.server, err = NewSocketServer(suite.port, suite.service)

	if err != nil {
		fmt.Println(err)
	}
}

func (suite *SocketServerSuite) TearDownTest() {
	suite.server.stop()
}

func (suite *SocketServerSuite) TestInstantiate() {
	assert.Equal(suite.T(), suite.port, suite.server.Port())
	assert.Equal(suite.T(), suite.service, suite.server.Service())
}

func (suite *SocketServerSuite) TestCanStartAndStopServer() {
	suite.server.start()
	assert.True(suite.T(), suite.server.Running())
	suite.server.stop()
	assert.False(suite.T(), suite.server.Running())
}

func (suite *SocketServerSuite) TestAcceptsAnIncomingConnection() {
	suite.server.start()
	_, err := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
	<-done

	if err != nil {
		fmt.Println(err)
	}
	suite.server.stop()
	assert.Equal(suite.T(), 1, suite.service.connections)
}

func (suite *SocketServerSuite) TestAcceptsMultipleIncomingConnections() {
	suite.server.start()
	_, err := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
	<-done
	done = make(chan bool)
	_, err2 := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
	<-done
	if err != nil {
		fmt.Println(err)
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	suite.server.stop()
	assert.Equal(suite.T(), 2, suite.service.connections)
}

type ReadingSocketService struct {
	*TestSocketService
	Message string
}

func (rss *ReadingSocketService) doService(c net.Conn) {
	buf := make([]byte, 1024)
	r := bufio.NewReader(c)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	rss.Message = string(buf[:n])
}

func NewReadingSocketService() *ReadingSocketService {
	readingService := new(ReadingSocketService)
	readingService.TestSocketService = NewTestSocketService(readingService)
	return readingService
}

type ReadingSocketServerTestSuite struct {
	suite.Suite
	readingService *ReadingSocketService
	server         *SocketServer
	port           int
}

func (suite *ReadingSocketServerTestSuite) SetupTest() {
	suite.port = 8043
	suite.readingService = NewReadingSocketService()

	var err error
	suite.server, err = NewSocketServer(suite.port, suite.readingService)

	if err != nil {
		fmt.Println(err)
	}
}

func (suite *ReadingSocketServerTestSuite) TearDownTest() {
	suite.server.stop()
}

func (suite *ReadingSocketServerTestSuite) TestCanSendAndReceiveData() {
	suite.server.start()
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Write([]byte("hello"))
	<-done
	if err != nil {
		fmt.Println(err)
	}

	suite.server.stop()
	assert.Equal(suite.T(), "hello", suite.readingService.Message)
}

type EchoSocketService struct {
	*TestSocketService
	Message string
}

func (ess *EchoSocketService) doService(c net.Conn) {
}

func NewEchoSocketService() *EchoSocketService {
	echoService := new(EchoSocketService)
	echoService.TestSocketService = NewTestSocketService(echoService)
	return echoService
}

type EchoSocketServerTestSuite struct {
	suite.Suite
	echoService *EchoSocketService
	server      *SocketServer
	port        int
}

func (suite *EchoSocketServerTestSuite) SetupTest() {
	suite.port = 8043
	suite.echoService = NewEchoSocketService()

	var err error
	suite.server, err = NewSocketServer(suite.port, suite.echoService)

	if err != nil {
		fmt.Println(err)
	}
}

func (suite *EchoSocketServerTestSuite) TearDownTest() {
	suite.server.stop()
}

func (suite *EchoSocketServerTestSuite) TestCanSendAndReceiveData() {
	//suite.server.start()
	//conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//_, err = conn.Write([]byte("hello"))
	//<-done
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//suite.server.stop()
	//assert.Equal(suite.T(), "hello", suite.readingService.Message)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSocketServerSuite(t *testing.T) {
	suite.Run(t, new(SocketServerSuite))
}

func TestWithReadingSocketService(t *testing.T) {
	suite.Run(t, new(ReadingSocketServerTestSuite))
}

func TestWithEchoSocketService(t *testing.T) {
	suite.Run(t, new(EchoSocketServerTestSuite))
}
