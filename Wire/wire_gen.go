// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func initializeEvent(msg string) (event, error) {
	mainMessage := newMessage(msg)
	mainGreeter := newGreeter(mainMessage)
	mainEvent, err := newEvent(mainGreeter)
	if err != nil {
		return event{}, err
	}
	return mainEvent, nil
}
