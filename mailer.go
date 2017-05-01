package hydrocarbon

import "log"

// A Mailer sends mail
type Mailer interface {
	Send(email string, body string) error
	RootDomain() string
}

type MockMailer struct {
	Mails []string
}

func (mm *MockMailer) Send(email string, body string) error {
	return nil
}
func (mm *MockMailer) RootDomain() string {
	return "http://localhost"
}

type StdoutMailer struct {
	Domain string
}

func (*StdoutMailer) Send(email string, body string) error {
	log.Println("hydrocarbon: new mail to", email, "\n", body)
	return nil
}

func (sm *StdoutMailer) RootDomain() string {
	return sm.Domain
}
