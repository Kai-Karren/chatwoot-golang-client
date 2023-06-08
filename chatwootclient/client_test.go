package chatwootclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateContact(t *testing.T) {

	// configure mocked chatwoot server

	createContactResponse := CreateContactResponse{
		Payload: Payload{
			Contact: Contact{
				ContactInboxes: []ContactInbox{
					{
						SourceID: "42",
					},
				},
			},
		},
	}

	body, _ := json.Marshal(createContactResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write(body)

	}))

	defer server.Close()

	// client set up

	client := ChatwootClient{
		BaseUrl:       server.URL,
		AccountId:     1,
		AgentBotToken: "",
		AgentToken:    "",
	}

	response, err := client.CreateContact(CreateContactRequest{
		InboxID: 1,
		Name:    "Unit Test Contact",
	})

	if err != nil {
		t.FailNow()
	}

	if response.Payload.Contact.ContactInboxes[len(response.Payload.Contact.ContactInboxes)-1].SourceID != "42" {
		t.FailNow()
	}

}
