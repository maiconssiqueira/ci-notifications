package github

type Notification struct {
	Command []string
}

func NewNotification() *Notification {
	return &Notification{}
}

type status struct {
	Context     string `json:"context"`
	State       string `json:"state"`
	Description string `json:"description"`
	TargetUrl   string `json:"target_url"`
	Return      Return `json:"return_status"`
}

type labels struct {
	Labels []string `json:"labels"`
}
type markup struct {
	Issue_number int            `json:"issue_number"`
	Markups      labels         `json:"labels"`
	Return       []ReturnMarkup `json:"return_status"`
}
type releases struct {
	TagName              string `json:"tag_name"`
	TargetCommitish      string `json:"target_commitish"`
	Name                 string `json:"name"`
	Body                 string `json:"body"`
	Draft                bool   `json:"draft"`
	Prerelease           bool   `json:"prerelease"`
	GenerateReleaseNotes bool   `json:"generate_release_notes"`
	Return               Return `json:"return_status"`
}
type comments struct {
	PrNumber int    `json:"prNumber"`
	Body     string `json:"body"`
	Return   Return `json:"return_status"`
}
type Github struct {
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
	Url          string `json:"url"`
	Token        string
	Sha          string   `json:"sha"`
	Statuses     status   `json:"status"`
	Releases     releases `json:"releases"`
	Comments     comments `json:"comments"`
	Markup       markup   `json:"markup"`
}
