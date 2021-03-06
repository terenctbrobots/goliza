package eliza

import "testing"

func TestCompareStatementPass(t *testing.T) {
	if !CompareStatement("Bye", Goodbyes, false) {
		t.Error("CompareStatement should return true")
	}
}

func TestGreetings(t *testing.T) {
	message := Greetings()

	if !CompareStatement(message, Introductions, false) {
		t.Errorf("Greetings() = %s", message)
	}
}

func TestGoodbye(t *testing.T) {
	goodbye := Goodbye()

	if !CompareStatement(goodbye, Goodbyes, false) {
		t.Errorf("Goodbye() = %s", goodbye)
	}
}

func TestReplyTo(t *testing.T) {
	reply := ReplyTo("Saya mahu ketolongan")

	if reply != "Kenapa awak mahu ketolongan?" {
		t.Errorf("ReplyTo() = %s", reply)
	}
}
