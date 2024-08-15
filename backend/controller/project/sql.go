package project

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/rqlite/sql"
)

type Query struct {
	Name    string
	Content string
}

func splitQueries(input string) []Query {
	var queries []Query
	var currentQuery Query

	// Regular expression to match query name lines
	re := regexp.MustCompile(`(?i)^\s*--\s*query_name\s*:\s*(.+)\s*$`)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if match := re.FindStringSubmatch(line); len(match) > 1 {
			if currentQuery.Name != "" {
				currentQuery.Content = strings.TrimSpace(currentQuery.Content)
				queries = append(queries, currentQuery)
			}
			currentQuery = Query{
				Name: strings.TrimSpace(match[1]),
			}
		} else if currentQuery.Name != "" {
			currentQuery.Content += line + "\n"
		}
	}

	if currentQuery.Name != "" {
		currentQuery.Content = strings.TrimSpace(currentQuery.Content)
		queries = append(queries, currentQuery)
	}

	return queries
}

func prevenSQLEscape(input string) error {

	parser := sql.NewParser(strings.NewReader(input))
	stmt, err := parser.ParseStatement()
	if err != nil {
		return err
	}

	switch stmt.(type) {
	case *sql.SelectStatement:
	default:
		return fmt.Errorf("invalid statement type")
	}

	err = sql.Walk(sql.VisitEndFunc(func(n sql.Node) error {

		switch snode := n.(type) {

		case *sql.QualifiedTableName:
			tbl := snode.TableName()
			if !strings.HasPrefix(tbl, "__project__") {
				return fmt.Errorf("you sneaky sneaky")
			}

		case *sql.QualifiedRef:

			if !strings.HasPrefix(snode.Table.Name, "__project__") {
				return fmt.Errorf("you sneaky sneaky")
			}

		}

		return nil
	}), stmt)

	if err != nil {
		return err
	}

	return nil

}

func (a *ProjectController) RunQuerySQL(userId int64, pid int64, input, name string, data []any) ([]map[string]any, error) {

	queries := splitQueries(input)

	queryContent := ""
	for _, q := range queries {
		if q.Name == name {
			queryContent = q.Content
		}
	}

	pp.Println("@qq", queries)

	if queryContent == "" {
		return nil, fmt.Errorf("cannot extract query by name: %s", name)
	}

	err := prevenSQLEscape(queryContent)
	if err != nil {
		return nil, err
	}

	return a.db.RunProjectSQLQuery(pid, queryContent, data)
}
