package robot

import "testing"

func TestSendMsgToEmail(t *testing.T) {
	if err := SendMsgToEmail("test title", "mail body"); err != nil {
		t.Log(err)
	}

}
