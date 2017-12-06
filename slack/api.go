package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"ssf/config"
	"ssf/model"
	"strconv"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

const (
	listUrl   = "https://slack.com/api/files.list"
	deleteUrl = "https://slack.com/api/files.delete"
)

func ListAllFiles() ([]model.File, error) {
	res, err := listFile(1)
	if err != nil {
		return nil, err
	}

	totalPages := res.PageInfo.Pages
	curPage := res.PageInfo.Page
	if curPage == totalPages {
		return res.Files, nil
	}

	var result = make([]model.File, 0)
	result = append(result, res.Files...)

	var ch = make(chan []model.File)
	var wg sync.WaitGroup

	for p := curPage + 1; p <= totalPages; p += 1 {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			res, err = listFile(index)
			if err != nil {
				log.Error().Msgf("list file with page:%d error:%+v\n", index, err)
				return
			}
			ch <- res.Files
		}(p)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for f := range ch {
		result = append(result, f...)
	}

	return result, nil
}

func DeleteAllFiles() ([]string, []string, error) {
	files, err := ListAllFiles()
	if err != nil {
		return nil, nil, err
	}

	success, fail := DeleteFiles(files)
	return success, fail, nil
}

func DeleteFiles(files []model.File) (success []string, fail []string) {
	var wg sync.WaitGroup
	var ch = make(chan map[string]string)

	for _, file := range files {
		wg.Add(1)

		go func(id string) {
			defer wg.Done()

			if err := deleteFile(id); err != nil {
				ch <- map[string]string{"fail": id}
			} else {
				ch <- map[string]string{"success": id}
			}
		}(file.Id)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		if file, ok := s["success"]; ok {
			success = append(success, file)
		} else if file, ok = s["fail"]; ok {
			fail = append(fail, file)
		}
	}

	return
}

func deleteFile(id string) error {
	form := url.Values{}
	form.Add("file", id)
	form.Add("token", config.Slack.Token)

	req, err := http.NewRequest("POST", deleteUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}

	type deleteRes struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}

	var res deleteRes
	if err = json.Unmarshal(b, &res); err != nil {
		return err
	}

	if !res.Ok {
		return errors.New(res.Error)
	}
	return nil
}

func listFile(page int) (*model.ListResponse, error) {

	values := url.Values{
		"token": []string{config.Slack.Token},
		"page":  []string{strconv.Itoa(page)},
	}

	if config.Slack.User != "" {
		values.Add("user", config.Slack.User)
	}

	res, err := http.PostForm(listUrl, values)
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var list model.ListResponse
	if err = json.Unmarshal(b, &list); err != nil {
		return nil, err
	}

	if !list.Ok {
		reason := list.Error
		if reason == "" {
			reason = "unknow error"
		}
		return nil, errors.New(reason)
	}
	return &list, nil
}
