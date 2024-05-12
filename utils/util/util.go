package util

import (
	"encoding/json"
	"fmt"
	"os"
	"xyz-books/constant"
	"xyz-books/infrastructure/repository/dao"
)

func GetSeedData() (
	seedBooks *[]dao.Book,
	seedPublishers *[]dao.Publisher,
	err error) {
	rawBooks, err := os.ReadFile(constant.ProgramSeedBookDirectory)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading seed data file: %w", err)
	}
	err = json.Unmarshal(rawBooks, &seedBooks)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing seed data: %w", err)
	}

	rawPublishers, err := os.ReadFile(constant.ProgramSeedPublisherDirectory)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading seed data file: %w", err)
	}
	err = json.Unmarshal(rawPublishers, &seedPublishers)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing seed data: %w", err)
	}

	return
}

func StringToPointer(s string) *string {
	return &s
}
