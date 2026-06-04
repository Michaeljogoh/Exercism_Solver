package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // YYYY-MM-DD
	Description string
	Change      int // amount in cents
}

// FormatLedger formats a ledger according to locale and currency.
func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	// Validate currency 
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("invalid currency")
	}

	// Validate locale
	header, err := ledgerHeader(locale)
	if err != nil {
		return "", err
	}

	// Copy entries (do not mutate input)
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	// Sort entries
	sortEntries(entriesCopy)

	var builder strings.Builder
	builder.WriteString(header)

	for _, entry := range entriesCopy {
		line, err := formatEntry(entry, locale, currency)
		if err != nil {
			return "", err
		}
		builder.WriteString(line)
	}

	return builder.String(), nil
}

// sortEntries sorts entries by date, description, change.
func sortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Date != entries[j].Date {
			return entries[i].Date < entries[j].Date
		}
		if entries[i].Description != entries[j].Description {
			return entries[i].Description < entries[j].Description
		}
		return entries[i].Change < entries[j].Change
	})
}

// ledgerHeader returns locale-specific header.
func ledgerHeader(locale string) (string, error) {
	switch locale {
	case "en-US":
		return "Date       | Description               | Change       \n", nil
	case "nl-NL":
		return "Datum      | Omschrijving              | Verandering  \n", nil
	default:
		return "", errors.New("invalid locale")
	}
}

// formatEntry formats one ledger entry.
func formatEntry(entry Entry, locale, currency string) (string, error) {
	date, err := formatDate(entry.Date, locale)
	if err != nil {
		return "", err
	}

	description := formatDescription(entry.Description)

	amount, err := formatAmount(entry.Change, currency, locale)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%-10s | %-25s | %13s\n",
		date,
		description,
		amount,
	), nil
}

// formatDate converts YYYY-MM-DD into locale format.
func formatDate(date, locale string) (string, error) {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return "", errors.New("invalid date")
	}

	year := date[0:4]
	month := date[5:7]
	day := date[8:10]

	switch locale {
	case "en-US":
		return month + "/" + day + "/" + year, nil
	case "nl-NL":
		return day + "-" + month + "-" + year, nil
	default:
		return "", errors.New("invalid locale")
	}
}

// formatDescription truncates long descriptions.
func formatDescription(desc string) string {
	if len(desc) > 25 {
		return desc[:22] + "..."
	}
	return desc
}

// formatAmount formats currency per locale.
func formatAmount(cents int, currency, locale string) (string, error) {
	negative := cents < 0
	if negative {
		cents = -cents
	}

	symbol := currencySymbol(currency)
	if symbol == "" {
		return "", errors.New("invalid currency")
	}

	switch locale {
	case "en-US":
		return formatUSAmount(cents, symbol, negative), nil
	case "nl-NL":
		return formatDutchAmount(cents, symbol, negative), nil
	default:
		return "", errors.New("invalid locale")
	}
}

// currencySymbol returns symbol.
func currencySymbol(currency string) string {
	switch currency {
	case "USD":
		return "$"
	case "EUR":
		return "€"
	default:   
		return ""
	}
}

// US formatting
func formatUSAmount(cents int, symbol string, negative bool) string {
	dollars := cents / 100
	remainder := cents % 100

	value := addThousands(dollars, ",")

	result := fmt.Sprintf("%s%s.%02d", symbol, value, remainder)

	if negative {
		return "(" + result + ")"
	}

	return result + " "
}

// Dutch formatting
func formatDutchAmount(cents int, symbol string, negative bool) string {
	euros := cents / 100
	remainder := cents % 100

	value := addThousands(euros, ".")

	result := fmt.Sprintf("%s %s,%02d", symbol, value, remainder)

	if negative {
		result = symbol + " -" + value + "," + fmt.Sprintf("%02d", remainder)
	}

	return result + " "
}

// thousands separator helper
func addThousands(number int, separator string) string {
	s := strconv.Itoa(number)

	if len(s) <= 3 {
		return s
	}

	var parts []string

	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}

	parts = append([]string{s}, parts...)

	return strings.Join(parts, separator)  
}