// Package ohdear is a Golang SDK to interact with oh-dear REST api.
//
// The Oh Dear API lets you configure everything about our application
// through a simple, structured Application Programming Interface (API).
// Everything you see in your dashboard can be controlled with the API.
// And as a bonus, all changes you make with the API will be visible
// in realtime on your dashboard.
//
// The full api documentation can be found in
// https://ohdear.app/docs/general/welcome
//
// When instantiating a new client, you can provide
// and API token using OHDEAR_API_TOKEN which is the
// default environment variable and it will be looked
// up by the library automatically. This is strongly
// recommended as it is the most secure way to deal
// with your key.
//
// If the library cannot resolve your API token from
// the environment you can provide it when instantiating
// a new ohdear client using the NewClient constructor.
package ohdear
