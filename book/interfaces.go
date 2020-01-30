package book

import "fmt"

//interface:接口，就和Java的接口差不多

type PeopleInformation struct {
	PhoneNumber int64
	Name        string
	IdCard      string
}

type RemoteVerify interface {
	VerifyCode(p PeopleInformation)
}

type AliyunRemoteVerify struct {
}

func (a AliyunRemoteVerify) VerifyCode(p PeopleInformation) {
	fmt.Printf("来自于Aliyun的检测接口...\n")
}

type LocalRemoteVerify struct {
}

func (l LocalRemoteVerify) VerifyCode(p PeopleInformation) {
	fmt.Printf("来自于本地的检测接口...\n")
}

type ChuanglanRemoteVerify struct {
}

func (c ChuanglanRemoteVerify) VerifyCode(p PeopleInformation) {
	fmt.Printf("来自于Chuanglan的检测接口．．．\n")
}

func RemoteStrategy() {
	peopleInformation := PeopleInformation{
		PhoneNumber: 18270719298,
		Name:        "柳絮",
		IdCard:      "411481199002202457",
	}
	var strategy RemoteVerify
	strategy = AliyunRemoteVerify{}
	strategy.VerifyCode(peopleInformation)
	strategy = LocalRemoteVerify{}
	strategy.VerifyCode(peopleInformation)
	strategy = ChuanglanRemoteVerify{}
	strategy.VerifyCode(peopleInformation)
}
