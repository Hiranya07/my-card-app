package validations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ExtractPayload(r *http.Request, inputStruct interface{}) error {

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(input, inputStruct)
	if err != nil {
		return err
	}

	return nil

}
