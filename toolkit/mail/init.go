package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

var (
	sendCh chan *Email
	config string
)

type Mail struct {
}

func init() {
	queueSize := 100
	host := "smtp.exmail.qq.com"
	port := 25
	username := "happs.lives@6617.com"
	password := "Userpass12"
	from := "happs.lives@6617.com"
	config = fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d,"from":"%s"}`, username, password, host, port, from)

	sendCh = make(chan *Email, queueSize)

	go func() {
		for {
			select {
			case m, ok := <-sendCh:
				if !ok {
					return
				}
				if err := m.Send(); err != nil {
					fmt.Println("发送email出错:", err.Error())
				}
			}
		}
	}()
}

func SendMail(address, name, subject, content string) bool {
	mail := NewEMail(config)
	mail.To = []string{address}
	mail.Subject = subject
	mail.HTML = content
	select {
	case sendCh <- mail:
		return true
	case <-time.After(time.Second * 3):
		return false
	}
}

var MailTpl *template.Template

func init() {
	MailTpl, _ = template.New("mail_tpl").Parse(`
	你好, <br/>

<p>以下是数据抓取任务执行结果：</p>

<p>
任务ID：{{.task_id}}<br/>
任务名称：{{.task_name}}<br/>       
执行时间：{{.start_time}}<br />
执行状态：{{.status}}<br />
</p>
<p>-------------以下是任务执行输出-------------</p>
<p>{{.output}}</p>
<p>
--------------------------------------------<br />

</p>
`)
}

func temple() (sendContent string) {
	data := make(map[string]interface{})
	// data["username"] = "负责人"
	data["task_id"] = 1
	data["task_name"] = "baozi"
	data["start_time"] = time.Unix(12323, 0).Format("2006-01-02 03:04:05 PM")
	data["status"] = "we"
	data["output"] = "fa"
	content := new(bytes.Buffer)
	MailTpl.Execute(content, data)
	sendContent = content.String()
	return
}
