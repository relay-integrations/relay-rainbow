package main

import (
	"flag"
	"encoding/json"
	"fmt"
	"os"
	"send2rainbowbubble/rainbow_rest"
	"github.com/puppetlabs/relay-sdk-go/pkg/log"
	"github.com/puppetlabs/relay-sdk-go/pkg/taskutil"
)

// Default configuration values for Rainbow
const (
	platform    = "openrainbow.com" // for sandbox, use "sandbox.openrainbow.com"
	callbackUrl = "http://myendpoint.com" // we don't care for receiving messages/events here..
)

type ConnectionSpec struct {
	LoginEmail string
	Password string
	AppID string
	AppSecret string
}

type Spec struct {
	Connection *ConnectionSpec
	BubbleName string
	Message string
}

func FindBubbleByName(name string, rooms []rainbow_rest.Room) (*rainbow_rest.Room, bool) {
	for i := range rooms {
		if rooms[i].Name == name {
			return &rooms[i], true
		}
	}
	return nil, false
}

func FindConversationWithJid(jid string, conversations []rainbow_rest.Conversation) (*rainbow_rest.Conversation, bool) {
	for i := range conversations {
		if conversations[i].JidIm == jid {
			return &conversations[i], true
		}
	}
	return nil, false
}

func main() {
	// Retrieve spec from Relay metadata
	u, err := taskutil.MetadataSpecURL()
	if err != nil {
		log.FatalE(err)
	}
	specURL := flag.String("spec-url", u, "url to fetch the spec from")

	flag.Parse()

	planOpts := taskutil.DefaultPlanOptions{SpecURL: *specURL}

	var spec Spec
	if err := taskutil.PopulateSpecFromDefaultPlan(&spec, planOpts); err != nil {
		log.FatalE(err)
	}

	// Validate the user provided inputs
	if spec.Connection.LoginEmail == "" {
		log.Fatal("specify the email login to use")
	} else if spec.Connection.Password == "" {
		log.Fatal("specify the password to use")
	} else if spec.Connection.AppID == "" {
		log.Fatal("specify the Rainbow App ID")
	} else if spec.Connection.AppSecret == "" {
		log.Fatal("specify the Rainbow App Secret")
	} else if spec.BubbleName == "" {
		log.Fatal("specify the Rainbow Bubble to send the message to")
	} else if spec.Message == "" {
		log.Fatal("specify the message to send the message to")
	}

	// Create Rainbow client
	r := rainbow_rest.NewRest(platform)

	// Authenticate to rainbow
	loginResponse, err := r.GetLogin(spec.Connection.LoginEmail, spec.Connection.Password, spec.Connection.AppID, spec.Connection.AppSecret)
	_ = loginResponse

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to authenticate to rainbow: %s \n", err.Error())
		return
	}

	s2sCnx, err := r.LoginS2S(callbackUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to do S2SLogin : %s \n", err.Error())
		return
	}

	defer r.DeleteS2SConnection(s2sCnx.ID) // Destroy the connection afterwards

	fmt.Printf("*********************************************\n")
	b, err := json.Marshal(s2sCnx)
	fmt.Printf("LOGIN S2S : %s\n", string(b))
	fmt.Printf("*********************************************\n")

	// Retrieve all bubbles:
	rooms, err := r.GetAllBubbles(loginResponse.LoggedInUser.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to do get bubbles : %s \n", err.Error())
		return
	}

	b, err = json.Marshal(rooms)
	fmt.Printf("res:%s\n", string(b))

	// Search for bubble by name:
	bubble, found := FindBubbleByName(spec.BubbleName, rooms)
	if !found {
		fmt.Fprintf(os.Stderr, "Bubble %s doesn't exist.\n", spec.BubbleName)
		return
	}

	// Retrieve Conversations:
	_, conversations, err := r.GetConversations(loginResponse.LoggedInUser.ID, nil)

	// Find the conversation we're looking for (// its jid matches the one of the bubble)
	conv, found := FindConversationWithJid(bubble.Jid, conversations)
	if !found {
		fmt.Fprintf(os.Stderr, "Conversation with bubble %s doesn't exist.\n", spec.BubbleName)
		return
	}

	fmt.Printf("*********************************************\n")
	b, err = json.Marshal(conv)
	fmt.Printf("conversation :%s\n", string(b))
	fmt.Printf("*********************************************\n")

	// Ok bubble found and conversation found, we're good to go:
	// Send presence
	if _, err := r.SendPresence(s2sCnx.ID); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send presence: %s.\n", err.Error())
		return
	}

	// Join the bubble
	if _, err := r.JoinBubble(s2sCnx.ID, bubble.ID); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to join bubble: %s.\n", err.Error())
		return
	}

	// Send the message
	if _, err := r.SendMessageToConversation(s2sCnx.ID, conv.ID, spec.Message); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message: %s.\n", err.Error())
		return
	}

	fmt.Printf("Done.\n")

}
