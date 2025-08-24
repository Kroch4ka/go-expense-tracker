package domain

import (
	"errors"
	"time"
)

type Expense struct {
	Id          int
	Date        *time.Time
	Description string
	Amount      int
	Currency    Currency
}

type ExpenseList struct {
	data    []Expense
	counter int
}

func (e *ExpenseList) ForEach(f func(Expense)) {
	for _, v := range e.data {
		f(v)
	}
}

func (e *ExpenseList) IsEmpty() bool {
	return len(e.data) == 0
}

func (e *ExpenseList) CollectSummary(filters ...ExpenseSummaryFilter) ExpenseSummary {
	total := 0

	for _, expense := range e.data {
		for _, f := range filters {
			if f.Filter(expense) {
				total += expense.Amount
			}
		}
	}

	return ExpenseSummary{
		total: total,
	}
}

func (e *ExpenseList) Add(expense Expense) {
	e.counter++
	expense.Id = e.counter
	if expense.Date == nil {
		date := time.Now()
		expense.Date = &date
	}
	e.data = append(e.data, expense)
}

func (e *ExpenseList) Delete(id int) error {
	if e.counter == 0 {
		return errors.New("not found")
	}

	e.counter -= 1
	result := make([]Expense, e.counter)
	var err error
	var isFound bool

	for _, val := range e.data {
		if val.Id == id {
			isFound = true
			continue
		}
		result = append(result, val)
	}

	e.data = result

	if !isFound {
		err = errors.New("not found")
	}

	return err
}

type ExpenseSummary struct {
	total int
}

type ExpenseSummaryFilter interface {
	Filter(Expense) bool
}

type Currency string

func (c Currency) Format() string {
	switch c {
	case USD:
		return "$"
	default:
		return ""
	}
}

type ExpenseSummaryFilterFunc func(Expense) bool

func (e ExpenseSummaryFilterFunc) Filter(expense Expense) bool {
	return e(expense)
}

func ExpenseSummaryFilterByMonth(month int) ExpenseSummaryFilterFunc {
	return func(expense Expense) bool {
		return int(expense.Date.Month()) == month
	}
}

const (
	USD Currency = "USD"
)
