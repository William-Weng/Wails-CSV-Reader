package utility

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func ReadCSV(path string) ([][]string, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func ReadResources(path string) (ResourceList, error) {
	f, err := os.Open(path)
	if err != nil {
		return ResourceList{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	header, err := r.Read()
	if err != nil {
		return ResourceList{}, err
	}

	index := make(map[string]int)

	for idx, name := range header {
		name = cleanBOMAndSpace(name)
		index[name] = idx
	}

	nameIdx, ok := index["Name"]
	if !ok {
		return ResourceList{}, fmt.Errorf("找不到欄位: Name")
	}

	notesIdx, hasNotes := index["Notes"]
	urlIdx, hasURL := index["URL"]
	exampleIdx, hasExample := index["Example"]
	typeIdx, hasType := index["Type"]

	var resources []Resource
	var types []string
	seen := make(map[string]struct{})

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ResourceList{}, err
		}

		resource := Resource{
			Title: row[nameIdx],
		}

		if hasNotes && len(row) > notesIdx {
			resource.Description = row[notesIdx]
		}

		if hasURL && len(row) > urlIdx {
			resource.URL = row[urlIdx]
		}

		if hasExample && len(row) > exampleIdx {
			resource.Example = row[exampleIdx]
		}

		if hasType && len(row) > typeIdx {
			resource.Types = splitTypes(row[typeIdx])

			for _, resourceType := range resource.Types {
				if _, exists := seen[resourceType]; !exists {
					seen[resourceType] = struct{}{}
					types = append(types, resourceType)
				}
			}
		}

		resources = append(resources, resource)
	}

	sort.Strings(types)

	return ResourceList{
		Types:     types,
		Resources: resources,
	}, nil
}

func splitTypes(value string) []string {
	parts := strings.Split(value, ",")

	types := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part != "" {
			types = append(types, part)
		}
	}

	return types
}
