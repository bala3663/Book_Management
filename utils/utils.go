package utils

import (
	"encoding/json" //marshaling and un-marshalling
	"io/ioutil"     //function documentation details
	"net/http"      //provides http client and servers
)

// recived body in json then we have to parse
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil { //unmarshelling
			return
		}
	}
}
