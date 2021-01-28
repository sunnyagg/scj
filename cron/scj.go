package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CronJob struct {
	Id          int64
	Group       int64
	Expression  string
	Timezone    string
	Url         string
	HttpMethod  string
	HttpHeaders string
	PostData    string
	Fail        int
	Status      int
	Name        string
	Notify      int
	Points      int
}

type Response struct {
	Status  string
	Code    int
	Data    interface{}
	Info    []string
	Message string
}

const (
	BaseUrl     = "https://www.setcronjob.com/api/"
	CronList    = "cron.list"
	CronEnable  = "cron.enable"
	CronDisable = "cron.disable"
	CronRun     = "cron.run"
)

func List(token string) ([]CronJob, error) {
	req := map[string]interface{}{}
	req["token"] = token

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := scjPost(CronList, reqBytes)

	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(res.Data)
	if err != nil {
		return nil, err
	}

	var crons []CronJob
	err = json.Unmarshal(data, &crons)
	if err != nil {
		return nil, err
	}

	return crons, nil
}

func EnableCron(token string, id int64) error {
	req := map[string]interface{}{}
	req["token"] = token
	req["id"] = id

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := scjPost(CronEnable, reqBytes)
	if res.Status != "success" {
		return fmt.Errorf("error:%s code: %d", res.Message, res.Code)
	}

	return nil
}

func DisableCron(token string, id int64) error {
	req := map[string]interface{}{}
	req["token"] = token
	req["id"] = id

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := scjPost(CronDisable, reqBytes)
	if res.Status != "success" {
		return fmt.Errorf("error:%s code: %d", res.Message, res.Code)
	}

	return nil

}

func RunCron(token string, id int64) error {
	req := map[string]interface{}{}
	req["token"] = token
	req["id"] = id

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := scjPost(CronRun, reqBytes)
	if res.Status != "success" {
		return fmt.Errorf("error:%s code: %d", res.Message, res.Code)
	}

	return nil

}

func scjPost(action string, body []byte) (*Response, error) {
	url := BaseUrl + action
	buf := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json", buf)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var res Response
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
