package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendRequest(id, course string) bool {
	url := "http://vu.sbu.ac.ir/class/course.list.php"

	payload := strings.NewReader("username=" + id)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return strings.Contains(string(body), "("+course+")")
}

func main() {
	course := "4011430107601" //Software Testing course
	var list []string
	for i := 1; i < 115; i++ {
		fmt.Println(i)
		id := fmt.Sprintf("98243%03d", i)
		if sendRequest(id, course) {
			list = append(list, id)
		}
	}

	fmt.Println(list)
}
