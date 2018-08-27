package utils/pop3

//original from github.com\reyoung\pop3

import (
	"bytes"

	//	"crypto/tls"
	"fmt"
	//	"github.com/reyoung/pop3"
	"log"
	"net/smtp"
	"strings"
	"time"
	"peterSZW/goutils/pop3"
)

/*
 *	user : example@example.com login smtp server user
 *	password: xxxxx login smtp server password
 *	host: smtp.example.com:port   smtp.163.com:25
 *	to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

//"crypto/tls

//func without_cert() {
//    tr := &http.Transport{
//        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//    }
//    client := &http.Client{Transport: tr}
//    _, err := client.Get("https://golang.org/")
//    if err != nil {
//        fmt.Println(err)
//    }
//}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func DemoSendmail() {
	user := "peta@163.com"
	password := "####"
	host := "smtp.163.com:25"
	to := "wilson@track4win.com"

	subject := "Test send email by golang"

	body := `
	<html>
	<body>
	<h3>
	"Test send email by golang"
	</h3>
	</body>
	</html>
	`
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}

func DemoGetmail() {

	//addr := "pop.gmail.com:995"
	//msg_buffer := make([]byte, 1024, 1024)
	//conn, err := tls.Dial("tcp", addr, nil)
	//sz, err := conn.Read(msg_buffer)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//} else if sz < 3 || string(msg_buffer[0:3]) != "+OK" {

	//	fmt.Println("OKKKK")
	//	//err = errors.New("Welcome Message Error")
	//	return
	//} else {
	//	fmt.Println(string(msg_buffer))
	//}

	cli, err := pop3.NewClient("mail.track4win.com:110", "wilson@track4win.com", "cy$2012")

	//cli, err := NewClient("pop.gmail.com:995", "peter.zw.song@gmail.com", "")

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	list, err := cli.UIDL()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("UIDL", list)

	time.Sleep(10e8 * 1)

	if len(list) > 0 {
		last := list[len(list)-1]

		msg, err := cli.GetMail(last.Id)
		if err != nil {
			log.Fatal(err)
		} else {
			if string(msg.UID) != last.UID {
				log.Println("Get Email Size Is Not as expect", string(msg.UID), last.UID)
			}
			log.Print(msg.Message.Header)
			log.Print(string(msg.UID))
			log.Print(msg.Message.Header.Get("from"))
			log.Print(msg.Message.Header.Get("to"))
			log.Print(msg.Message.Header.Get("subject"))

			buf := new(bytes.Buffer)
			buf.ReadFrom(msg.Message.Body)
			s := buf.String() // Does a complete copy of the bytes in the buffer.
			//defer msg.Message.Body.Close()
			//body, err := ioutil.ReadAll(msg.Message.Body)
			log.Print(s)

		}
	}
}
