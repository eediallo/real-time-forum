package handler

var (
	homePagePath = "/"
	// dashboardPagePath = "/dashboard"
	styleDir        = "/static/styles/"
	headerCSS       = styleDir + "header.css"
	indexCSS        = styleDir + "index.css"
	cssLoginPath    = styleDir + "login.css"
	dashboardCSS    = styleDir + "dashboard.css"
	cssRegisterPath = styleDir + "register.css"
	cssLogoPath     = styleDir + "logo.css"

	imagesDir = "/static/images/"
	logPath   = imagesDir + "forum_logo.png"

	jsDir                    = "/static/scripts/"
	commentJS                = jsDir + "comments.js"
	likeDislike              = jsDir + "likeDislike.js"
	loginJSPath              = jsDir + "login.js"
	likeDislikeCommentJsPath = jsDir + "likeDislikeComment.js"
	filterJsPath             = jsDir + "filter.js"
	wsPath                   = jsDir + "ws.js"
	privateMessageJS         = jsDir + "privateMessage.js"

	googleIcons = "https://fonts.googleapis.com/icon?family=Material+Icons"

	jsImagPath    = imagesDir + "JS.png"
	goImagePath   = imagesDir + "golang.png"
	rustImagePath = imagesDir + "Rust.png"

	goOfficialPagePath   = "https://go.dev/"
	jsOfficialPagePath   = "https://www.javascript.com/"
	rustOfficialPagePath = "https://www.rust-lang.org/"

	// status code

	onlineUsers = []string{}
)

type StatT = uint8

const (
	Null StatT = iota
	UserNotFound
	InvalidPasswordForUser
	DeleteExistingSession
	InsertingSession
)
