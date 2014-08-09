package main

import (
	"os"
	"io/ioutil"
	"sync/atomic"
	"encoding/json"
)

const (
	RECORDS_FILE string = "records.json"
)

type Records map[string]*int64

var (
	records Records
)

func RecordMissingBundle(id string) {
	if _, ok := records[id]; !ok {
		x := int64(0)
		records[id] = &x
	}

	atomic.AddInt64(records[id], 1)
}

func GetRecords() Records {
	r := Records{}

	for k, v := range records {
		n := atomic.LoadInt64(v)
		r[k] = &n
	}

	return r
}

func SaveRecords() error {
	data, err := json.Marshal(GetRecords())
	if err != nil {
		return err
	}

	return ioutil.WriteFile(RECORDS_FILE, data, 0644)
}

func LoadRecords() {
	records = Records{}

	if _, err := os.Stat(RECORDS_FILE); os.IsNotExist(err) {
		// No records file
		return
	}

	data, err := ioutil.ReadFile(RECORDS_FILE)
	if err != nil {
		panic(err)
	}

	tmp := map[string]int64{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		panic(err)
	}

	for k, v := range tmp {
		v := v
		records[k] = &v
	}
}
