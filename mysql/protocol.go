package mysql

import (
	"bytes"
	"encoding/binary"
)

//
//主要放置Slave和Master交互的协议对象
//

type GreetingProtocol struct {
	payloadLength            int32
	sequenceId               int32
	version                  int32
	serverVersionInformation string
	threadId                 int32
	scrambleDataPart1        string
	emptyInt                 int32
	serverCharset            int32
	serverStatus             int32
	serverCapability         int32
	pktScrambleLength        int32
	unUseInformation         []byte
	scrambleDataPart2        string
	scramblePluginName       string
}

func (g *GreetingProtocol) parse(source []byte) {
	index := 0
	_ = binary.Read(bytes.NewBuffer(append(source[0:3], byte(0))), binary.LittleEndian, &g.payloadLength)
	g.sequenceId = int32(source[3])
	g.version = int32(source[4])
	index = 5
	serverVersionBytes := make([]byte, 10)
	for start := index; start < len(source); start = start + 1 {
		if source[start] == byte(0) {
			index = start
			break
		}
		serverVersionBytes = append(serverVersionBytes, source[start])
	}
	g.serverVersionInformation = string(serverVersionBytes)
	_ = binary.Read(bytes.NewBuffer(source[index+1:index+5]), binary.LittleEndian, &g.threadId)

	index = index + 5
	scrambleDataPart1Byte := make([]byte, 20)
	for start := index; start < len(source); start = start + 1 {
		if source[start] == byte(0) {
			index = start
			break
		}
		scrambleDataPart1Byte = append(scrambleDataPart1Byte, source[start])
	}
	g.scrambleDataPart1 = string(scrambleDataPart1Byte)

	index = index + 1
	_ = binary.Read(bytes.NewBuffer(source[index:index+2]), binary.LittleEndian, &g.emptyInt)
	index = index + 2
	g.serverCharset = int32(source[index])
	index = index + 1
	_ = binary.Read(bytes.NewBuffer(append(source[index:index+2], byte(0), byte(0))), binary.LittleEndian, &g.serverStatus)
	index = index + 2
	_ = binary.Read(bytes.NewBuffer(append(source[index:index+2], byte(0), byte(0))), binary.LittleEndian, &g.serverCapability)

	index = index + 2
	g.pktScrambleLength = int32(source[index])

	index = index + 1
	g.unUseInformation = source[index : index+10]

	index = index + 10
	scrambleDataPart2Byte := make([]byte, 20)
	for start := index; start < len(source); start = start + 1 {
		if source[start] == byte(0) {
			index = start
			break
		}
		scrambleDataPart2Byte = append(scrambleDataPart2Byte, source[start])
	}
	g.scrambleDataPart2 = string(scrambleDataPart2Byte)
	index = index + 1

	scramblePluginNameByte := make([]byte, 20)
	for start := index; start < len(source); start = start + 1 {
		if source[start] == byte(0) {
			index = start
			break
		}
		scramblePluginNameByte = append(scramblePluginNameByte, source[start])
	}
	g.scramblePluginName = string(scramblePluginNameByte)
}

type CertificationRequest struct {
	payloadLength      int32
	sequenceId         int32
	clientFlag         int32
	maxPacketSize      int32
	clientCharset      int32
	unUseInformation   [23]byte
	userName           string
	clientCapability   int32
	scrambleData       [20]byte
	databaseName       string
	scramblePluginName string
}
