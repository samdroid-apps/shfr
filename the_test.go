package main

import (
	"time"
	"testing"
)

func TestRecordsAtomic(t *testing.T) {
	records = Records{}

	for i := 0; i < 10; i++ {
		go RecordMissingBundle("com.example.activity")
	}

	// Let the counter update
	time.Sleep(10 * time.Millisecond)

	r := GetRecords()
	n, ok := r["com.example.activity"]
	if !ok || *n != 10 {
		t.Errorf("10 calls to RecordMissingBundle did not give 10 in the " +
				"record; it gave %d", n)
	}
}

func TestReocrdsSaveLoad(t *testing.T) {
	ten := int64(10)
	twenty := int64(20)
	data := Records{"com.example.activity": &ten,
                    "net.example.activity": &twenty}
	records = data

    SaveRecords()
	LoadRecords()

	for k, v := range GetRecords() {
		v2, ok := data[k]
		if !ok || *v != *v2 {
			t.Errorf("Saving and loading %v resulted in %v on load",
					data, records)
		}
	}
}
