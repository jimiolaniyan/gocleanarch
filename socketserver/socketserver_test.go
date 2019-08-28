package socketserver

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net"
	"strconv"
	"testing"
)

type FakeSocketService struct {
	connections int
	Message     string
}

var done = make(chan bool)

func (fss *FakeSocketService) serve(c net.Conn) {
	defer c.Close()

	fss.connections = fss.connections + 1
	done <- true
	//buf := make([]byte, 1024)
	//r := bufio.NewReader(c)
	//
	//n, err := r.Read(buf)
	//fmt.Println(n)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fss.Message = string(buf[:n])
	//fmt.Println(fss.Message)

}

type SocketServerSuite struct {
	suite.Suite
	server  *SocketServer
	service *FakeSocketService
	port    int
}

func (suite *SocketServerSuite) SetupTest() {
	suite.port = 8043
	suite.service = &FakeSocketService{}
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

//func (suite *SocketServerSuite) TestCanSendAndReceiveData() {
//	suite.server.start()
//	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(suite.port))
//	if err != nil {
//		fmt.Println(err)
//	}
//	_, err = conn.Write([]byte("hello"))
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(suite.service.Message)
//	suite.server.stop()
//	assert.Equal(suite.T(), "hello", suite.service.Message)
//}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSocketServerSuite(t *testing.T) {
	suite.Run(t, new(SocketServerSuite))
}
