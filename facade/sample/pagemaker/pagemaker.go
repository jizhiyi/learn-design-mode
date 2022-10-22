package pagemaker

import (
	"fmt"
	"learn/adapter/a2"
)

func MakeWelcomePage(mailAddr, fileName string) (err error) {
	var mailProp a2.FileIO
	mailProp, err = getProperties("maildata")
	if err != nil {
		return err
	}
	userName := mailProp.GetValue(mailAddr)
	var hw *htmlWriter
	hw, err = NewHtmlWriter(fileName)
	if err != nil {
		return err
	}
	hw.title("Welcome to " + userName + "'s page!")
	hw.paragraph("欢迎来到" + userName + "的主页。")
	hw.paragraph("等着你的邮件哦！")
	hw.mailto(mailAddr, userName)
	hw.close()
	fmt.Println(fileName + " is created for " + mailAddr + " (" + userName + ")")
	return nil
}
