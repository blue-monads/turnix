package project

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/blue-monads/turnix/backend/services/database/autoquery"
	"github.com/blue-monads/turnix/backend/xtypes/models"
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

func prevenSQLEscape(prefix string, input string) error {

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
			if !strings.HasPrefix(tbl, prefix) {
				return fmt.Errorf("you sneaky sneaky")
			}

		case *sql.QualifiedRef:

			if !strings.HasPrefix(snode.Table.Name, prefix) {
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

func (a *ProjectController) RunQuerySQL2(userId int64, pid int64, qstr string, data []any) ([]map[string]any, error) {
	err := prevenSQLEscape("__project__", qstr)
	if err != nil {
		return nil, err
	}

	return a.db.RunProjectSQLQuery(pid, qstr, data)
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

	err := prevenSQLEscape("__project__", queryContent)
	if err != nil {
		return nil, err
	}

	return a.db.RunProjectSQLQuery(pid, queryContent, data)
}

func (a *ProjectController) ListProjectTables(userId int64, pid int64) ([]string, error) {
	return a.db.ListProjectTables(pid)
}

func (a *ProjectController) ListProjectTableColumns(userId int64, pid int64, table string) ([]models.TableColumn, error) {
	return a.db.ListProjectTableColumns(pid, table)
}

func (a *ProjectController) AutoQueryProjectTable(userId int64, pid int64, table string, opts autoquery.LoaderParams) ([]map[string]any, error) {

	prefix := fmt.Sprintf("z_%d_", pid)

	fullTableName := fmt.Sprintf("z_%d_%s", pid, table)

	builder := autoquery.NewBuilder(fullTableName, opts)
	strquery, args := builder.Build()

	err := prevenSQLEscape(prefix, strquery)
	if err != nil {
		return nil, err
	}

	pp.Println("@strquery", strquery)

	data, err := a.db.RunProjectSQLQuery(pid, strquery, args)
	if err != nil {
		return nil, err
	}

	pp.Println("@data", data)

	return data, nil
}
