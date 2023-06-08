package chatwootclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Please note that certain functions like to add labels or assign agents are blocked when using an Agent Bot Token
// therefore an AgentToken has to be provided. The client uses the AgentBotToken wherever possible.
type ChatwootClient struct {
	BaseUrl       string
	AccountId     int
	AgentBotToken string
	AgentToken    string
}

func NewChatwootClient(baseUrl string, accountId int, agentBotToken string) ChatwootClient {
	return ChatwootClient{
		baseUrl,
		accountId,
		agentBotToken,
		"",
	}
}

func NewChatwootClientWithAgentToken(baseUrl string, accountId int, agentBotToken string, agentToken string) ChatwootClient {
	return ChatwootClient{
		baseUrl,
		accountId,
		agentBotToken,
		agentToken,
	}
}

type CreateContactRequest struct {
	InboxID          int         `json:"inbox_id"`
	Name             string      `json:"name,omitempty"`
	EMail            string      `json:"email,omitempty"`
	PhoneNumber      string      `json:"phone_number,omitempty"`
	Avatar           string      `json:"avatar,omitempty"`
	AvatarUrl        string      `json:"avatar_url,omitempty"`
	Identifier       string      `json:"identifier,omitempty"`
	CustomAttributes interface{} `json:"custom_attributes,omitempty"`
}

type CreateContactResponse struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	ID             int            `json:"id"`
	ContactInboxes []ContactInbox `json:"contact_inboxes"`
}

type ContactInbox struct {
	SourceID string `json:"source_id"`
}

func (client *ChatwootClient) CreateContact(createContactRequest CreateContactRequest) (CreateContactResponse, error) {

	url := fmt.Sprintf("%s/api/v1/accounts/%v/contacts", client.BaseUrl, client.AccountId)

	requestJSON, err := json.Marshal(createContactRequest)

	if err != nil {
		return CreateContactResponse{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentToken)

	if err != nil {
		return CreateContactResponse{}, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return CreateContactResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return CreateContactResponse{}, err
	}

	var createContactResponse CreateContactResponse

	if err := json.Unmarshal(body, &createContactResponse); err != nil {
		return CreateContactResponse{}, err
	}

	return createContactResponse, nil

}

type CreateNewConversationRequest struct {
	SourceID  string `json:"source_id"`
	InboxID   string `json:"inbox_id"`
	ContactID string `json:"contact_id"`
	Status    string `json:"status"`
}

type CreateNewConversationResponse struct {
	ID        int `json:"id"`
	AccountId int `json:"account_id"`
	InboxId   int `json:"inbox_id"`
}

func (client *ChatwootClient) CreateNewConversation(sourceId string, inboxId string, contactId string, status string) (CreateNewConversationResponse, error) {

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations", client.BaseUrl, client.AccountId)

	requestBody := CreateNewConversationRequest{
		SourceID:  sourceId,
		InboxID:   inboxId,
		ContactID: contactId,
		Status:    status,
	}

	requestJSON, err := json.Marshal(requestBody)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentBotToken)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	var createNewConversationResponse CreateNewConversationResponse

	if err := json.Unmarshal(body, &createNewConversationResponse); err != nil {
		return CreateNewConversationResponse{}, err
	}

	return createNewConversationResponse, nil

}

func (client *ChatwootClient) CreateNewConversationWithSourceIdAndInboxId(source_id string, inbox_id string) (CreateNewConversationResponse, error) {

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations", client.BaseUrl, client.AccountId)

	requestBody := fmt.Sprintf(`{"source_id": "%s", "inbox_id": "%s"}`, source_id, inbox_id)

	requestJSON, err := json.Marshal(requestBody)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentBotToken)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return CreateNewConversationResponse{}, err
	}

	var createNewConversationResponse CreateNewConversationResponse

	if err := json.Unmarshal(body, &createNewConversationResponse); err != nil {
		return CreateNewConversationResponse{}, err
	}

	return createNewConversationResponse, nil

}

type GetMessagesResponse struct {
	Meta    interface{}      `json:"meta"`
	Payload ChatwootMessages `json:"payload"`
}

type ChatwootMessages []struct {
	Id          int         `json:"id"`
	Content     string      `json:"content"`
	ContentType string      `json:"content_type,omitempty"`
	Private     bool        `json:"private,omitempty"`
	Sender      interface{} `json:"sender,omitempty"`
}

func (client *ChatwootClient) GetMessages(conversationId string) (ChatwootMessages, error) {

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/messages", client.BaseUrl, client.AccountId, conversationId)

	request, _ := http.NewRequest(http.MethodGet, url, nil)

	request.Header.Add("api_access_token", client.AgentToken)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var getMessagesResponse GetMessagesResponse

	if err := json.Unmarshal(responseBody, &getMessagesResponse); err != nil {
		return nil, err
	}

	return getMessagesResponse.Payload, nil

}

// Struct that allows to build minimal create message requests.
type CreateNewMessageRequest struct {
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
	Private     bool   `json:"private"`
}

type CreateNewMessageResponse struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
	Private     bool   `json:"private"`
}

func NewCreateNewMessageRequest(content string, messageType string, private bool) CreateNewMessageRequest {
	return CreateNewMessageRequest{
		Content:     content,
		MessageType: messageType,
		Private:     private,
	}
}

func (client *ChatwootClient) CreateNewMessage(conversationId int, createMessageRequest CreateNewMessageRequest) (CreateNewMessageResponse, error) {

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/messages", client.BaseUrl, client.AccountId, conversationId)

	requestBodyJSON, err := json.Marshal(createMessageRequest)

	if err != nil {
		return CreateNewMessageResponse{}, err
	}

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentBotToken)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return CreateNewMessageResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return CreateNewMessageResponse{}, err
	}

	var createNewMessageResponse CreateNewMessageResponse

	if err := json.Unmarshal(body, &createNewMessageResponse); err != nil {
		return CreateNewMessageResponse{}, err
	}

	return createNewMessageResponse, nil

}

func (client *ChatwootClient) CreateOutgoingMessage(conversationId int, content string) (CreateNewMessageResponse, error) {

	return client.CreateNewMessage(conversationId, NewCreateNewMessageRequest(
		content,
		"outgoing",
		false,
	))

}

func (client *ChatwootClient) CreateOutgoingPrivateMessage(conversationId int, content string) (CreateNewMessageResponse, error) {

	return client.CreateNewMessage(conversationId, NewCreateNewMessageRequest(
		content,
		"outgoing",
		true,
	))

}

func (client *ChatwootClient) CreateIncomingMessage(conversationId int, content string) (CreateNewMessageResponse, error) {

	return client.CreateNewMessage(conversationId, NewCreateNewMessageRequest(
		content,
		"incoming",
		false,
	))

}

func (client *ChatwootClient) CreateIncomingPrivateMessage(conversationId int, content string) (CreateNewMessageResponse, error) {

	return client.CreateNewMessage(conversationId, NewCreateNewMessageRequest(
		content,
		"incoming",
		true,
	))

}

type AddLabelsRequest struct {
	Labels []string `json:"labels"`
}

func (client *ChatwootClient) AddLabels(conversationId int, labels []string) error {

	if client.AgentToken == "" {
		return errors.New("agentToken is empty. Adding labels requires a Chatwoot agent token")
	}

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/labels", client.BaseUrl, client.AccountId, conversationId)

	requestBody := AddLabelsRequest{
		Labels: labels,
	}

	requestBodyJSON, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentToken)

	_, err = http.DefaultClient.Do(request)

	return err

}

func (client *ChatwootClient) AddLabel(conversationId int, label string) error {

	if client.AgentToken == "" {
		return errors.New("agentToken is empty. Adding labels requires a Chatwoot agent token")
	}

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/labels", client.BaseUrl, client.AccountId, conversationId)

	requestBody := AddLabelsRequest{
		Labels: []string{label},
	}

	requestBodyJSON, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyJSON))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentToken)

	_, err = http.DefaultClient.Do(request)

	return err

}

func (client *ChatwootClient) Assign(conversationId int, assignee_id int) error {

	if client.AgentToken == "" {
		return errors.New("agentToken is empty. Adding assignments requires a Chatwoot agent token")
	}

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/assignments", client.BaseUrl, client.AccountId, conversationId)

	requestBody := fmt.Sprintf(`{"assignee_id": %v}`, assignee_id)

	requestBodyAsBytes := []byte(requestBody)

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyAsBytes))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentToken)

	_, err := http.DefaultClient.Do(request)

	return err

}

func (client *ChatwootClient) AssignTeam(conversationId int, team_id int) error {

	if client.AgentToken == "" {
		return errors.New("agentToken is empty. Adding assignments requires a Chatwoot agent token")
	}

	url := fmt.Sprintf("%s/api/v1/accounts/%v/conversations/%v/assignments", client.BaseUrl, client.AccountId, conversationId)

	requestBody := fmt.Sprintf(`{"team_id": %v}`, team_id)

	requestBodyAsBytes := []byte(requestBody)

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyAsBytes))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("api_access_token", client.AgentToken)

	_, err := http.DefaultClient.Do(request)

	return err

}
