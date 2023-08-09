package status

type Message struct {
	Message string `json:"message"`
}

func Success() Message {
	return Message{
		Message: "success",
	}
}
