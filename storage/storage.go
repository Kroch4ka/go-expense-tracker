package storage

import (
	"errors"
	"os"
	"time"

	"github.com/Kroch4ka/go-expense-tracker/domain"
	"github.com/gocarina/gocsv"
)

type Storage interface {
	Load() domain.ExpenseList
	Unload(domain.ExpenseList)
}

type CSVStorage struct{}

type ExpenseCSV struct {
	Id          int             `csv:"-"`
	Date        *time.Time      `csv:"date"`
	Description string          `csv:"description"`
	Amount      int             `csv:"amount"`
	Currency    domain.Currency `csv:"currency"`
}

func (c CSVStorage) Load() domain.ExpenseList {
	expensesFile, err := os.OpenFile("expenses.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer expensesFile.Close()
	var expenses []ExpenseCSV
	var list domain.ExpenseList

	if err := gocsv.UnmarshalFile(expensesFile, &expenses); err != nil && !errors.Is(err, gocsv.ErrEmptyCSVFile) {
		panic(err)
	}

	for _, v := range expenses {
		list.Add(domain.Expense{
			Id:          v.Id,
			Date:        v.Date,
			Description: v.Description,
			Amount:      v.Amount,
			Currency:    v.Currency,
		})
	}

	return list
}

func (c CSVStorage) Unload(list domain.ExpenseList) {
	expensesFile, err := os.OpenFile("expenses.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer expensesFile.Close()

	var expenses []ExpenseCSV

	list.ForEach(func(e domain.Expense) {
		expenses = append(expenses, ExpenseCSV{
			Id:          e.Id,
			Date:        e.Date,
			Description: e.Description,
			Amount:      e.Amount,
			Currency:    e.Currency,
		})
	})

	if list.IsEmpty() {
		err = expensesFile.Truncate(0)
		if err != nil {
			panic(err)
		}
		return
	}

	err = gocsv.MarshalFile(&expenses, expensesFile)

	if err != nil {
		panic(err)
	}
}
