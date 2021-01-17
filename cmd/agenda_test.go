package cmd

import (
	"bytes"
	"gocal/models"
	"io/ioutil"
	"testing"
)

type googlecalendarServiceMock struct{}

func (gs googlecalendarServiceMock) GetEvents() []*models.Event {
	return []*models.Event{{ID: "1"}}
}

func Test_agendaCmd(t *testing.T) {
	b := new(bytes.Buffer)
	cmd := agendaCmd(googlecalendarServiceMock{})
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "1" {
		t.Fatalf("expected \"%s\" got \"%s\"", "1", string(out))
	}
}
