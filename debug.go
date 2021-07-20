// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlchemy

import (
	"fmt"
	"strings"
	"time"

	"yunion.io/x/log"
)

var (
	DEBUG_SQLCHEMY = false
)

func sqlDebug(sqlstr string, variables []interface{}) {
	for _, v := range variables {
		switch v.(type) {
		case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			sqlstr = strings.Replace(sqlstr, "?", fmt.Sprintf(`%v`, v), 1)
		case string, time.Time:
			sqlstr = strings.Replace(sqlstr, "?", fmt.Sprintf(`"%s"`, v), 1)
		default:
			sqlstr = strings.Replace(sqlstr, "?", fmt.Sprintf(`"%v"`, v), 1)
		}
	}
	log.Debugln("SQuery ", sqlstr)
}

func (tq *SQuery) DebugQuery() {
	sqlstr := tq.String()
	vars := tq.Variables()
	sqlDebug(sqlstr, vars)
}

func (sqf *SSubQuery) DebugQuery() {
	sqlstr := sqf.Expression()
	vars := sqf.query.Variables()
	sqlDebug(sqlstr, vars)
}

func (t *STableSpec) DebugInsert(dt interface{}) error {
	return t.insert(dt, false, true)
}

func (t *STableSpec) DebugInsertOrUpdate(dt interface{}) error {
	return t.insert(dt, true, true)
}

func (ts *STableSpec) DebugUpdateFields(dt interface{}, fields map[string]interface{}) error {
	return ts.updateFields(dt, fields, true)
}
