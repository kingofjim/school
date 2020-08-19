package school

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Mail(api string, subject string, title string, content string, targetSlice... string)  {
	target := strings.Join(targetSlice, ",")
	body := fmt.Sprintf("{\"mailer_target\": \"%v\",\"mailer_subject\": \"%v\",\"mailer_title\": \"%v\",\"mailer_content\": \"%v\"}", target, subject, title, content)
	resp, err := http.Post(api, "application/json", strings.NewReader(body))
	checkErr(err)
	defer resp.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body1))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
