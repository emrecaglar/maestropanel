package maestropanel

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (m *MaestroPanel) getURL(a action) string {
	return m.Url + "/api/" + m.Version + a.URL
}

func (m MaestroPanel) writeData(method string, postURL string, obj interface{}) (response []byte, err error) {
	postURL += "?key=" + m.Key + "&format=json&suppress_response_codes=true"

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	convertObj2Form(writer, obj)

	err = writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, postURL, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if err != nil {
		response = nil
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		response = nil
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		response = nil
		return
	}

	response = bodyBytes
	err = nil
	return
}

// func (m MaestroPanel) writeData(method string, postURL string, obj interface{}) (response []byte, err error) {
// 	postURL += "?key=" + m.Key + "&format=json&suppress_response_codes=true"

// 	// if obj != nil {
// 	// 	postURL += convertToQueryString(obj)
// 	// }

// 	req, err := http.NewRequest(method, postURL, bytes.NewBufferString(convertToQueryString(obj)))

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	if err != nil {
// 		response = nil
// 		return
// 	}

// 	client := &http.Client{}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		response = nil
// 		return
// 	}

// 	defer resp.Body.Close()

// 	bodyBytes, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		response = nil
// 		return
// 	}

// 	response = bodyBytes
// 	err = nil
// 	return
// }
