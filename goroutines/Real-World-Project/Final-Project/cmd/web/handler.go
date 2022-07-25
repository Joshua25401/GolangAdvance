package main

import "net/http"

func (app *Config) Dashboard(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	// Parse form post
	if err := r.ParseForm(); err != nil {
		app.ErrorLog.Println(err)
	}

	// Get email and password from form post
	// We've user credentials here
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)

	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// If NOT validPassword
	if !validPassword {
		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// All validation pass
	// Put the userID to redis cache
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)
	app.Session.Put(r.Context(), "flash", "Succesfull Login!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
