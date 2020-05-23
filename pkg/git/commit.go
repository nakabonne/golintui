package git

// Encloses it with `^`, to escape the double quotation in the message.
const prettyCommitFormat = "{%n  ^^^^sha^^^^: ^^^^%H^^^^,%n  ^^^^message^^^^: ^^^^%s^^^^%n},"

type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
}

func (c *Commit) ShortSha() string {
	if len(c.SHA) < 8 {
		return c.SHA
	}
	return c.SHA[:8]
}
