package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func sendRequest(id, course string) bool {
	url := "http://vu.sbu.ac.ir/class/course.list.php"

	payload := strings.NewReader("username=" + id)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return sendRequest(id, course)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return strings.Contains(string(body), "("+course+")")
}

type result struct {
	lock sync.Mutex
	list []string
	wg   sync.WaitGroup
}

func recieveResponse(id, course string, res *result, guard chan struct{}) {
	if sendRequest(id, course) {
		res.lock.Lock()
		res.list = append(res.list, id)
		res.lock.Unlock()
	}
	<-guard
	res.wg.Done()
}

func main() {
	course := "4011430107601" //Software Testing course
	var res result
	maxParallelReqs := 10
	guard := make(chan struct{}, maxParallelReqs)

	entries := []int{96, 97, 98, 99, 400, 401}
	for _, year := range entries {
		for i := 1; i < 115; i++ {
			id := fmt.Sprintf("%d243%03d", year, i)
			res.wg.Add(1)
			guard <- struct{}{}
			go recieveResponse(id, course, &res, guard)
		}
	}

	fmt.Print(strings.Join(res.list, " "))
}
