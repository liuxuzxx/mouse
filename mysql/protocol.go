package mysql

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

//
//主要放置Slave和Master交互的协议对象
//

type GreetingProtocol struct {
	payloadLength            int32
	sequenceId               int32
	protocolVersion          int32
	serverVersionInformation string
	threadId                 int32
	scrambleDataPart1        string
	capabilityFlags1         int32
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
	g.protocolVersion = int32(source[4])
	index = 5
	serverVersionBytes := make([]byte, 0)
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
	g.scrambleDataPart1 = string(source[index : index+8])
	index = index + 1 + 8
	_ = binary.Read(bytes.NewBuffer(source[index:index+2]), binary.LittleEndian, &g.capabilityFlags1)
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
	payloadLength    int32
	sequenceId       int32
	clientFlag       int32
	maxPacketSize    int32
	clientCharset    uint8
	unUseInformation [23]byte
	userName         string
	databaseName     string
	password         string
	scrambleData     []byte
	pluginName       string
}

func (c *CertificationRequest) encode() []byte {
	c.sequenceId = 0
	c.clientFlag = 1<<19 + 8
	c.maxPacketSize = 82
	c.clientCharset = 8
	c.unUseInformation = [23]byte{0}

	loginRequestByte := make([]byte, 0)
	loginRequestByte = append(loginRequestByte, byte(c.sequenceId))

	loginRequestByte = append(loginRequestByte, c.convertLittleEndian(c.clientFlag)...)
	loginRequestByte = append(loginRequestByte, c.convertLittleEndian(c.maxPacketSize)...)
	loginRequestByte = append(loginRequestByte, c.clientCharset)
	loginRequestByte = append(loginRequestByte, c.unUseInformation[0:]...)
	loginRequestByte = append(loginRequestByte, append([]byte(c.userName), byte(0))...)
	loginRequestByte = append(loginRequestByte, append(c.mysqlNativePassword(), byte(0))...)
	loginRequestByte = append(loginRequestByte, append([]byte(c.databaseName), byte(0))...)

	protocolLength := len(loginRequestByte)
	lengthBytes := c.convertLittleEndian(int32(protocolLength))

	protocolBytes := make([]byte, 0)
	protocolBytes = append(protocolBytes, lengthBytes[0:3]...)
	protocolBytes = append(protocolBytes, loginRequestByte...)

	return protocolBytes
}

func (c *CertificationRequest) convertLittleEndian(source int32) []byte {
	targetBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(targetBytes, uint32(source))
	return targetBytes
}

func (c *CertificationRequest) mysqlNativePassword() []byte {
	sha1Crypt := sha1.New()
	sha1Crypt.Write([]byte(c.password))
	hashStageFirst := sha1Crypt.Sum(nil)
	fmt.Printf("%x\n", hashStageFirst)

	sha1Crypt.Reset()
	sha1Crypt.Write(hashStageFirst)
	hashStageSecond := sha1Crypt.Sum(nil)

	sha1Crypt.Reset()
	sha1Crypt.Write(append(c.scrambleData, hashStageSecond...))
	hashStateThird := sha1Crypt.Sum(nil)

	for index := 0; index < len(hashStateThird); index = index + 1 {
		hashStateThird[index] = hashStateThird[index] ^ hashStageFirst[index]
	}
	return hashStateThird
}
