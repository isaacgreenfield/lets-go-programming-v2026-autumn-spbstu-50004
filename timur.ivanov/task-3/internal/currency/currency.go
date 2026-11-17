package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading XML-file: %w", err)
	}

	var curCurs CurCurs
	if err := xml.Unmarshal(xmlData, &curCurs); err != nil {
		return fmt.Errorf("error decrypting XML-file: %w", err)
	}
	var output []OutputCurrency
	for _, cur := range curCurs.Currencies {
		numCode, err := strconv.Atoi(cur.NumCode)
		if err != nil {
			return fmt.Errorf("wrong NumCode %q: %w", cur.NumCode, err)
		}
		valueStr := strings.ReplaceAll(cur.Value, ",", ".")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return fmt.Errorf("wrong Value %q: %w", cur.Value, err)
		}
		output = append(output, OutputCurrency{
			NumCode:  numCode,
			CharCode: cur.CharCode,
			Value:    value,
		})
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})
	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("error dir creation %q: %w", dir, err)
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return fmt.Errorf("error forming JSON: %w", err)
	}
	if err := os.WriteFile(outputFile, jsonData, 0o644); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
