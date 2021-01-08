package tools

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"strconv"
	"strings"
)
//验证邮箱密码
func CheckEmailPassword(server string, email string, password string) bool {
	if !strings.Contains(server, ":") {
		return false
	}
	var c *client.Client
	serverSlice := strings.Split(server, ":")
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		return false
	}

	// 不要忘了退出
	//defer c.Logout()

	// 登陆
	c=connect(server,email,password)
	if c==nil{
		return false
	}
	return true
}
//获取连接
func connect(server string, email string, password string)*client.Client{
	var c *client.Client
	var err error
	serverSlice := strings.Split(server, ":")
	uri := serverSlice[0]
	port, _ := strconv.Atoi(serverSlice[1])
	if port != 993 && port != 143 {
		return nil
	}
	if port == 993 {
		c, err = client.DialTLS(fmt.Sprintf("%s:%d", uri, port), nil)
	} else {
		c, err = client.Dial(fmt.Sprintf("%s:%d", uri, port))
	}
	if err != nil {
		return nil
	}

	// 登陆
	if err := c.Login(email, password); err != nil {
		return nil
	}
	return c
}
//获取邮件夹
func GetFolders(server string, email string, password string)[]string{
	var c *client.Client
	//defer c.Logout()
	c=connect(server,email,password)
	if c==nil{
		return nil
	}
	// 列邮箱
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	// 存储邮件夹
	var folders []string
	for m := range mailboxes {
		folders=append(folders,m.Name)
	}
	return folders
}