package handler

import (
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if !NotFoundHandler(w, r) {
		return
	}

	data := db.PageData{
		CSSHomePage:        indexCSS,
		Logo:               logPath,
		HomePath:           homePagePath,
		LogoCSS:            cssLogoPath,
		GoIamge:            goImagePath,
		RustImage:          rustImagePath,
		JsImage:            jsImagPath,
		GolangOfficialPage: goOfficialPagePath,
		RustOfficialPage:   rustOfficialPagePath,
		JSOfficialPage:     jsOfficialPagePath,
	}
	RenderTemplate(w, "index", data)
}
