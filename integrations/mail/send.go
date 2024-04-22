package mail

func Send(receiver string, subject string, template string, items interface{}) error {
	r := newMailRequest([]string{receiver}, subject)
	return r.SendMail(template, items)

}
