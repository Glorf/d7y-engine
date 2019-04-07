package main

import (
	"flag"
	"fmt"
	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
	//"github.com/go-redis/redis"
	"log"
	"net/smtp"
	"os"
)

type Mailer struct {
	sendServer string
	sendUser string
	sendPassword string
	sendPort string
	hostname string
}

func (m Mailer) sendMail(sender string, recipent string, body string, subject string, replyto string) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		m.sendUser,
		m.sendPassword,
		m.sendServer,
	)

	err := smtp.SendMail(
		m.sendServer+":"+m.sendPort,
		auth,
		sender,
		[]string{recipent},
		[]byte("To:"+recipent+"\r\nSubject: "+subject+"\r\nReply-To: "+replyto+"\r\n\r\n" + body),
	)
	if err != nil {
		log.Fatal(err)
	}
}


var DiplomacyProcessor = func() backends.Decorator {
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskValidateRcpt {

					//TODO: check for user in redis by uuid of receiver
					//e.RcptTo ==

					//if not found:
					/* return backends.NewResult(
					   response.Canned.FailNoSenderDataCmd),
					   backends.NoSuchUser
					*/
					// if no error:
					return p.Process(e, task)
				} else if task == backends.TaskSaveMail {

					//TODO: check before saving to redis, then save to queue

					// if you want your processor to do some processing after
					// receiving the email, continue here.
					// if want to stop processing, return
					// errors.New("Something went wrong")
					// return backends.NewBackendResult(fmt.Sprintf("554 Error: %s", err)), err
					// call the next processor in the chain
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}

func (m Mailer) startListener() {
	cfg := &guerrilla.AppConfig{}
	sc := guerrilla.ServerConfig{
		ListenInterface: "0.0.0.0:587",
		IsEnabled:true,
	}
	cfg.Servers = append(cfg.Servers, sc)
	bcfg := backends.BackendConfig{
		"save_workers_size":  3,
		"save_process":      "HeadersParser|Header|Debugger|Diplomacy",
		"log_received_mails": true,
		"primary_mail_host" : m.hostname,

	}

	cfg.BackendConfig = bcfg

	d := guerrilla.Daemon{Config: cfg}
	d.AddProcessor("Diplomacy", DiplomacyProcessor)
	err := d.Start()

	if err != nil {
		fmt.Println("Server error!")
	}
}

func main() {
	mailer := Mailer{
		sendServer: *flag.String("sender-server", "localhost", "Setup outbound smtp server"),
		sendPort: *flag.String("sender-port", "25", "Setup outbound smtp port"),
		sendUser: os.Getenv("DIPLOMACY_SMTP_USER"),
		sendPassword: os.Getenv("DIPLOMACY_SMTP_PASSWORD"),
		hostname: *flag.String("host", "diplomacy.mbien.pl", "Game server hostname"),

	}

	flag.Parse()
	mailer.sendMail("no-reply@diplomacy.mbien.pl","michal@mbien.pl", "Hello", "Diplomacy", "some-uuid@diplomacy.mbien.pl")
	mailer.startListener()
}