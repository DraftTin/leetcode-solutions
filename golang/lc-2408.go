package main

import (
	"strconv"
	"strings"
)

type SQL struct {
	tables   map[string][][]string
	idMarks  map[string]int
	colMarks map[string]int
}

func Constructor(names []string, columns []int) SQL {
	sql := SQL{}
	sql.tables = map[string][][]string{}
	sql.idMarks = map[string]int{}
	sql.colMarks = map[string]int{}
	for i := 0; i < len(names); i++ {
		sql.tables[names[i]] = [][]string{}
		sql.idMarks[names[i]] = 1
		sql.colMarks[names[i]] = columns[i]
	}
	return sql
}

func (this *SQL) Ins(name string, row []string) bool {
	if _, ok := this.tables[name]; ok == true {
		if this.colMarks[name] != len(row) {
			return false
		}
		id := this.idMarks[name]
		this.idMarks[name]++
		this.tables[name] = append(this.tables[name], append([]string{strconv.Itoa(id)}, row...))
		return true
	}
	return false

}

func (this *SQL) Rmv(name string, rowId int) {
	if table, ok := this.tables[name]; ok == true {
		for i, row := range table {
			id, _ := strconv.Atoi(row[0])
			if id == rowId {
				this.tables[name] = append(table[:i], table[i+1:]...)
			}
		}
	}
}

func (this *SQL) Sel(name string, rowId int, columnId int) string {
	if table, ok := this.tables[name]; ok == true {
		if columnId > this.colMarks[name] {
			return "<null>"
		}
		for _, row := range table {
			id, _ := strconv.Atoi(row[0])
			if id == rowId {
				return row[columnId]
			}
		}
	}
	return "<null>"
}

func (this *SQL) Exp(name string) []string {
	res := []string{}
	if table, ok := this.tables[name]; ok == true {
		for _, row := range table {
			res = append(res, strings.Join(row, ","))
		}
	}
	return res
}

/**
 * Your SQL object will be instantiated and called as such:
 * obj := Constructor(names, columns);
 * param_1 := obj.Ins(name,row);
 * obj.Rmv(name,rowId);
 * param_3 := obj.Sel(name,rowId,columnId);
 * param_4 := obj.Exp(name);
 */
