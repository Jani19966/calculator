package frontend

import (
	"calculator/backend"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func HandleRequests() {
	http.HandleFunc("/add_numbers", AddNumbers)

	err := http.ListenAndServe(":3333", nil)

	logrus.Info("Listening on port 3333")
	if errors.Is(err, http.ErrServerClosed) {
		logrus.Errorf("The server has been closed!")
	} else if err != nil {
		logrus.Errorf("An error had happened: %v", err)
	}
}

func AddNumbers(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Couldnt read request body for 'add numbers': %v", err)
	}

	jsonInputData := AddNumbersRequest
	json.Unmarshal(reqBody, &jsonInputData)

	io.WriteString(w, strconv.Itoa(backend.AddNumbers(jsonInputData.Num1, jsonInputData.Num2)))
}
