package git

type Commit struct {
	SHA           string
	Message       string
	Author        string
	UnixTimestamp int64
}

func (c *Commit) ShortSha() string {
	if len(c.SHA) < 8 {
		return c.SHA
	}
	return c.SHA[:8]
}
