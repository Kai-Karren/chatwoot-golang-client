package chatwootclient

type ChatwootMockClient struct {
}

func NewChatwootMockClient() ChatwootMockClient {
	return ChatwootMockClient{}
}

func (c *ChatwootMockClient) AddLabel(label string) error {
	return nil
}

func (c *ChatwootMockClient) AddLabels(labels []string) error {
	return nil
}

func (c *ChatwootMockClient) Assign(conversationId int, assignee_id int) error {
	return nil
}

func (c *ChatwootMockClient) AssignTeam(conversationId int, team_id int) error {
	return nil
}

func (c *ChatwootMockClient) CreateContact(createContactRequest CreateContactRequest) (CreateContactResponse, error) {
	return CreateContactResponse{}, nil
}

func (c *ChatwootMockClient) CreateIncomingMessage(conversationId int, content string) (CreateNewMessageResponse, error) {
	return CreateNewMessageResponse{}, nil
}

func (c *ChatwootMockClient) CreateIncomingPrivateMessage(conversationId int, content string) (CreateNewMessageResponse, error) {
	return CreateNewMessageResponse{}, nil
}

func (c *ChatwootMockClient) CreateNewConversation(createNewConversationRequest CreateNewConversationRequest) (CreateNewConversationResponse, error) {
	return CreateNewConversationResponse{}, nil
}

func (c *ChatwootMockClient) CreateNewMessage(conversationId int, createMessageRequest CreateNewMessageRequest) (CreateNewMessageResponse, error) {
	return CreateNewMessageResponse{}, nil
}

func (c *ChatwootMockClient) CreateOutgoingMessage(conversationId int, content string) (CreateNewMessageResponse, error) {
	return CreateNewMessageResponse{}, nil
}

func (c *ChatwootMockClient) CreateOutgoingPrivateMessage(conversationId int, content string) (CreateNewMessageResponse, error) {
	return CreateNewMessageResponse{}, nil
}

func (c *ChatwootMockClient) GetMessages(conversationId string) (ChatwootMessages, error) {
	return ChatwootMessages{}, nil
}
