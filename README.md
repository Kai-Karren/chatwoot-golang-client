# Chatwoot golang Client
Unoffical implementation of a golang client that allows to interact with the [Chatwoot API](https://www.chatwoot.com/developers/api/). 

This is a private project in early development with the primary purpose being the creation of a golang client that can be used by agent bots to interact with Chatwoot. It does only cover a small subset of the functionality of the Chatwoot API. A full coverage of the Chatwoot API is also not planned. I am implementing the functionality that I have a concreate use for.

Currently, there exist a single client, the ChatwootClient that expects an AgentBotToken and can optionally use an AgentToken
to be able to perform operations like to assign labels or assign agents or teams to a conversation which is not allowed
when using an AgentBotToken.
