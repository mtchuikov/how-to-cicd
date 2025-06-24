package main

import (
	"log/slog"
	"net/http"

	fake "github.com/brianvoe/gofakeit"
)

func generateUserV1Handler(rw http.ResponseWriter, _ *http.Request) {
	const op = "generateUserV1Handler"

	person := fake.Person()

	user := UserV1{
		Username:  fake.Username(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}

	result, err := user.ToJson()
	if err != nil {
		slog.Error(
			"failed to encode generated user",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		http.Error(
			rw,
			"failed to encode generated user",
			http.StatusInternalServerError,
		)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	_, err = rw.Write(result)
	if err != nil {
		http.Error(
			rw,
			"failed to write response body",
			http.StatusInternalServerError,
		)
		return
	}
}
