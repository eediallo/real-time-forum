package db

// PageData represent the all the data in the system
type PageData struct {
	HeaderCSS            string
	CommentJS            string
	DashboardCSS         string
	Posts                []Post
	CssLoginPath         string
	CssRegisterPath      string
	HomePath             string
	CSSHomePage          string
	Logo                 string
	LogoCSS              string
	GoogleIcons          string
	JsImage              string
	GoIamge              string
	RustImage            string
	GolangOfficialPage   string
	RustOfficialPage     string
	JSOfficialPage       string
	LikeDislike          string
	IsAuthenticated      bool
	ErrorMessage         string
	LoginJS              string
	LikeDislikeCommentJS string
	FilterJS             string
	WS                   string
	PrivateMessageJS     string
	OnlineUsers          []User
	Username             string
	SignUpCoverImage     string
	ProfileCSS           string
	MainJS               string
}
