package main

import (
	"fmt"
	"strings"
)

/*
	In this section we will learn about how to pass a builder
	as a parameter and example of how to use it
*/

// First, we create a email struct here
// That have from, to, subject, body attribute as a string
type Email struct {
	from, to, subject, body string
}

// Then we create the builder here
type EmailBuilder struct {
	email Email
}

// Create some function that simply build up the attributes
func (builder *EmailBuilder) From(from string) *EmailBuilder {
	// Here we can add bunch of logic like validation etc
	if !strings.Contains(from, "@") {
		panic("Format email not valid!, @ char not found !\n")
	}
	builder.email.from = from
	return builder
}

func (builder *EmailBuilder) To(to string) *EmailBuilder {
	// Here we can add bunch of logic like validation etc
	if !strings.Contains(to, "@") {
		panic("Format email not valid!, @ char not found !\n")
	}
	builder.email.to = to
	return builder
}

func (builder *EmailBuilder) Subject(subject string) *EmailBuilder {
	if len(subject) > 255 {
		panic("Subject is too long! max 255 char!")
	}
	builder.email.subject = subject
	return builder
}

func (builder *EmailBuilder) Body(body string) *EmailBuilder {
	builder.email.body = body
	return builder
}

// Then we create some private func sendEmail that takes email
func sendEmail(email *Email) *Email {
	return email
}

// Then we create public method SendEmail
// This method must contain a builder parameter in order to make object from builder class
type emailBuilder func(builder *EmailBuilder)

func SendEmail(action emailBuilder) *Email {
	builder := EmailBuilder{}
	action(&builder)
	return sendEmail(&builder.email)
}

func main() {
	// Here as you can see we force the client to make builder as a parameter
	// In order to build the email
	received := SendEmail(func(builder *EmailBuilder) {
		builder.From("joshuaryandafres@gmail.com").
			To("chesya@yahoo.co.id").
			Subject("Blind Spot Means Nothin!").
			Body("Please watch this movie!")
	})

	fmt.Printf("Received an email from	" + received.from + "!\n\n")
	fmt.Println("SUBJECT\t:", received.subject)
	fmt.Println(received.body)
}
