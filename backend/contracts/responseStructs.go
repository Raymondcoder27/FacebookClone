package contracts

type CommentResponse struct {
	ID   uint         `json:"id"`
	Text string       `json:"text"`
	User UserResponse `json:"user"`
}

// PostResponse struct to format post data in response
type PostResponse struct {
	ID        uint              `json:"id"`
	Text      string            `json:"text"`
	Image     string            `json:"image"`
	CreatedAt string            `json:"created_at"`
	Comments  []CommentResponse `json:"comments"`
	User      UserResponse      `json:"user"`
}

// UserResponse struct to format user data in response
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
