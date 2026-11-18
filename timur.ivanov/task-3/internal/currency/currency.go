package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

type CurCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type OutputCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Process(inputFile, outputFile string) error {
	curCurs, err := decodeXML(inputFile)
	if err != nil {
		return err
	}

	output, err := convertCurrencies(curCurs)
	if err != nil {
		return err
	}

	if err := saveToJSON(outputFile, output); err != nil {
		return err
	}

	return nil
}

func decodeXML(inputFile string) (*CurCurs, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("error reading XML-file: %w", err)
	}

	var curCurs CurCurs

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&curCurs); err != nil {
		return nil, fmt.Errorf("error decrypting XML-file: %w", err)
	}

	return &curCurs, nil
}

func convertCurrencies(curCurs *CurCurs) ([]OutputCurrency, error) {
	output := make([]OutputCurrency, 0)

	for _, cur := range curCurs.Currencies {
		numCodeStr := strings.TrimSpace(cur.NumCode)

		var numCode int

		if numCodeStr != "" {
			parsedNumCode, err := strconv.Atoi(numCodeStr)
			if err != nil {
				return nil, fmt.Errorf("wrong NumCode %q: %w", cur.NumCode, err)
			}

			numCode = parsedNumCode
		}

		valueStr := strings.ReplaceAll(strings.TrimSpace(cur.Value), ",", ".")

		var value float64

		if valueStr != "" {
			parsedValue, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return nil, fmt.Errorf("wrong Value %q: %w", cur.Value, err)
			}

			value = parsedValue
		}

		output = append(output, OutputCurrency{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(cur.CharCode),
			Value:    value,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})

	return output, nil
}

func saveToJSON(outputFile string, output []OutputCurrency) error {
	dir := filepath.Dir(outputFile)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("error dir creation %q: %w", dir, err)
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("error forming JSON: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
