package main

import (
	"testing"
	"time"
)

const dateFormat = "2006-01-02"

func TestGetLastUpdate_NoFile(t *testing.T) {
	expected := time.Now().UTC().AddDate(0, 0, -30).Format(dateFormat)
	actual := getLastUpdate(isNotExistMockTrue, readFileMockEmptyFile, statMock).Format(dateFormat)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetLastUpdate_File(t *testing.T) {
	expected, err := time.Parse(isoDateTimeFormat, "2026-02-26T12:23:11.723Z")
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}
	actual := getLastUpdate(isNotExistMockFalse, readFileMockFakeData, statMock)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}
