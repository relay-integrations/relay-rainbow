# rainbow-step-bubble-message-send

This [Rainbow](https://www.openrainbow.com) step sends a message to a desired bubble.

In order to connect to Rainbow you will need to **Create a new Rainbow application** at [https://hub.openrainbow.com/](https://hub.openrainbow.com/):

* Select "Create new application" 
* Give it a name like "notifications-from-relay"
* Click "Deploy to Rainbow Production Platform" and provide payment details if necessary

Once your application has been successfully deployed, configure the Relay step:
* Copy the following fields from the Identification section:
    * Application ID 
    * Application Secret
* Configure the Relay step with the following `connection` spec values:
    * `loginEmail` - email address associated with your Rainbow account
    * `password` - password associated with your Rainbow account
    * `appID` - application ID you copied before
    * `appSecret` - application secret you copied before

An example configuration is provided below:
```
name: send-message
image: relaysh/slack-step-message-send
spec:
  connection:
    loginEmail: !Secret loginEmail
    password: !Secret password
    appID: !Secret appID
    appSecret: !Secret appSecret 
  bubbleName: !Parameter Bubble
  message: !Parameter message # Example: "hello Relay!"
```
