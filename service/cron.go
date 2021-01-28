package service

import (
	"SetCronJob/cron"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ListAllCrons(token string) error {
	crons, err := cron.List(token)
	if err != nil {
		return err
	}

	ids := ""
	for _, cron := range crons {
		ids += fmt.Sprintf("%d,", cron.Id)
	}
	ids = strings.TrimSuffix(ids, ",")
	fmt.Println(ids)
	return nil
}

func ListEnabledCrons(token string) error {
	crons, err := cron.List(token)
	if err != nil {
		return err
	}

	ids := ""
	for _, cron := range crons {
		if cron.Status == 0 {
			ids += fmt.Sprintf("%d,", cron.Id)
		}
	}
	ids = strings.TrimSuffix(ids, ",")
	fmt.Println(ids)
	return nil

}

func EnableCronList(token string, ids string) error {
	if ids == "" {
		return errors.New("cron ids missing")
	}

	idList := strings.Split(ids, ",")

	for _, idStr := range idList {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("invalid cron id: %s", idStr)
		}

		err = cron.EnableCron(token, int64(id))
		if err != nil {
			fmt.Printf("failed to enable cron with id: %d", id)
		}

		// to avoid throttling
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func DisableCronList(token string, ids string) error {
	if ids == "" {
		return errors.New("cron ids missing")
	}

	idList := strings.Split(ids, ",")

	for _, idStr := range idList {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("invalid cron id: %s", idStr)
		}

		err = cron.DisableCron(token, int64(id))
		if err != nil {
			fmt.Printf("failed to enable cron with id: %d", id)
		}

		// to avoid throttling
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func RunCronList(token string, ids string) error {
	if ids == "" {
		return errors.New("cron ids missing")
	}
	idList := strings.Split(ids, ",")

	for _, idStr := range idList {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("invalid cron id: %s", idStr)
		}

		err = cron.RunCron(token, int64(id))
		if err != nil {
			fmt.Printf("failed to enable cron with id: %d", id)
		}

		// to avoid throttling
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
