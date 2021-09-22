package to_entity

import (
	"errors"
	"fmt"
	"go-db2entity/string_utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ToEntity(filePath, sql string) {
	tableName, err := GetTableName(sql)
	if err != nil {
		log.Fatal(err)
	}
	fields, err := GetFieldList(sql)
	if err != nil {
		log.Fatal(err)
	}

	resText := GetResText(tableName, fields)

	ret, err := string_utils.WriteToFile(filePath, resText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}

func GetResText(tableName string, fields []*Field) string {
	primaryField := &Field{}
	maxNameLen, maxTypeLen := GetMaxFieldLen(fields)
	tableStructName := string_utils.Ucfirst(string_utils.Case2Camel(tableName))
	resImport := "import (\n\t\"encoding/json\"\n"
	hasTime := false
	resText := fmt.Sprintf("const TableName%s = \"%s\"\n\n", tableStructName, tableName)
	resText += fmt.Sprintf("type %s struct{\n", tableStructName)
	for _, field := range fields {
		if field.FieldType == "time.Time" {
			hasTime = true
		}
		space1, space2 := " ", " "
		fieldStructName := string_utils.Ucfirst(string_utils.Case2Camel(field.Name))
		for {
			end1, end2 := true, true
			if len(fieldStructName)+len(space1) < maxNameLen+1 {
				space1 += " "
				end1 = false
			}
			if len(field.FieldType)+len(space2) < maxTypeLen+1 {
				space2 += " "
				end2 = false
			}
			if end1 && end2 {
				break
			}
		}

		xorm := ""
		if field.IsPrimary {
			xorm = fmt.Sprintf("%s pk autoincr", fieldStructName)
		} else {
			if field.AllowNull {
				xorm += "null "
			} else {
				xorm += "not null "
			}
			if len(field.Default) > 0 {
				xorm += "default " + field.Default
			}
			xorm += " " + field.FieldSqlType
			if len(field.Remark) > 0 {
				xorm += " COMMENT(" + field.Remark + ")"
			}
			xorm += " '" + field.Name + "'"
		}

		resText += fmt.Sprintf("\t%s%s%s%s`json:\"%s\" xorm:\"%s\"`\n",
			fieldStructName,
			space1,
			field.FieldType,
			space2,
			field.Name,
			xorm)
		if field.IsPrimary {
			primaryField = field
		}
	}
	resText += fmt.Sprintf("}\n\n\n")
	if hasTime {
		resImport += "\t\"time\"\n"
	}
	resImport += ")\n\n"
	resText = "package entity\n\n" + resImport + resText

	resText += fmt.Sprintf("func (t *%s)TableName() string {\n\treturn TableName%s\n}\n\n",
		tableStructName, tableStructName)
	resText += fmt.Sprintf("func (t *%s) String() string {\n\tdata, _ := json.Marshal(t)\n\treturn string(data)\n}\n\n",
		tableStructName)
	resText += fmt.Sprintf("type %sList []*%s\n\n",
		tableStructName, tableStructName)

	primaryFieldName := string_utils.Ucfirst(string_utils.Case2Camel(primaryField.Name))
	resText += fmt.Sprintf("type %sMapBy%s map[%s]*%s\n\n",
		tableStructName,
		primaryFieldName,
		primaryField.FieldType,
		tableStructName)
	resText += fmt.Sprintf("func (list %sList) Get%ss() []%s {\n\tres := make([]%s, 0)\n\tfor _, l := range list {\n\t\tres = append(res, l.%s)\n\t}\n\treturn res\n}\n\n",
		tableStructName,
		primaryFieldName,
		primaryField.FieldType,
		primaryField.FieldType,
		primaryFieldName)
	resText += fmt.Sprintf("func (list %sList) GetMapBy%s() %sMapBy%s {\n\tres := make(%sMapBy%s)\n\tfor _, l := range list {\n\t\tres[l.%s] = l\n\t}\n\treturn res\n}\n\n",
		tableStructName,
		primaryFieldName,
		tableStructName,
		primaryFieldName,
		tableStructName,
		primaryFieldName,
		primaryFieldName)
	resText += fmt.Sprintf("func (list %sList) string() string {\n\tres, _ := json.Marshal(list)\n\treturn string(res)\n}\n\n",
		tableStructName)
	resText += fmt.Sprintf("// 无论何种类型的值，都可转化为string类型来构建map\nfunc (list %sList) GroupBy(visitor func(item *%s) string) map[string]%sList {\n\tres := make(map[string]%sList)\n\tfor _, l := range list {\n\t\tkey := visitor(l)\n\t\tvalue := res[key]\n\t\tif value == nil {\n\t\t\tvalue = make(%sList, 0)\n\t\t}\n\t\tres[key] = append(value, l)\n\t}\n\treturn res\n}\n\n",
		tableStructName,
		tableStructName,
		tableStructName,
		tableStructName,
		tableStructName)
	return resText
}

func GetMaxFieldLen(fields []*Field) (int, int) {
	maxNameLen := 0
	maxTypeLen := 0
	for _, field := range fields {
		nameLen := len(field.Name)
		typeLen := len(field.FieldType)
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
		if typeLen > maxTypeLen {
			maxTypeLen = typeLen
		}
	}
	return maxNameLen, maxTypeLen
}

func GetTableName(sql string) (string, error) {
	re, err := regexp.Compile("^create table [\\d\\w\\S]*")
	if err != nil {
		fmt.Println(err)
	}
	res := re.FindAll([]byte(sql), 1)
	if len(res) != 1 {
		return "", errors.New("表名匹配0个或多个，请检查数据")
	}
	ret := strings.ReplaceAll(strings.ToLower(string(res[0])), "create table ", "")
	retList := string_utils.ExplodeStr(ret, ".")
	if len(retList) == 1 {
		return retList[0], nil
	}
	if len(retList) == 2 {
		return retList[1], nil
	}
	return "", errors.New("tableName error, " + ret)
}

type Field struct {
	Name         string
	FieldType    string
	FieldSqlType string
	Default      string
	Remark       string
	IsPrimary    bool
	AllowNull    bool
}

func GetFieldList(sql string) ([]*Field, error) {
	sql, err := string_utils.GetContentInFirstBrackets(sql)
	if err != nil {
		return nil, err
	}
	sql = strings.TrimFunc(sql, string_utils.IsNeedClear)
	strListMap := make(map[string]string, 0)
	sql, strListMap = string_utils.ReplaceAllQuotationMarks(sql)

	re, err := regexp.Compile("\\([\\d]*[\\s,\\d]*\\)")
	if err != nil {
		return nil, err
	}
	fieldSqlTypeMap := make(map[string]string)
	collectNum := 0
	sql = string(re.ReplaceAllFunc([]byte(sql), func(r []byte) []byte {
		key := "$FST" + strconv.Itoa(collectNum)
		fieldSqlTypeMap[key] = string(r)
		collectNum++
		return []byte(key)
	}))
	re, err = regexp.Compile("constraint [\\S]*[\\s\\n]*[\\S]* \\([\\S]*,[\\s]*[\\S]*\\)")
	if err != nil {
		return nil, err
	}
	sql = string(re.ReplaceAll([]byte( sql), []byte("")))

	filedSqlList := string_utils.ExplodeStr(sql, ",")
	fieldsList := make([]*Field, 0)
	for _, fieldSql := range filedSqlList {
		fieldSql = strings.TrimFunc(fieldSql, string_utils.IsNeedClear)
		re, err = regexp.Compile("^[\\S]*")
		if err != nil {
			return nil, err
		}
		field := &Field{}
		r := re.Find([]byte(fieldSql))
		field.Name = string(r)

		fieldSql = strings.Replace(fieldSql, field.Name, "", 1)
		fieldSql = strings.Trim(fieldSql, " ")
		fieldSql = strings.Trim(fieldSql, "	")
		re, err = regexp.Compile("^[a-zA-Z]*(\\$FST)*[\\d]*")
		if err != nil {
			return nil, err
		}
		r1 := re.Find([]byte(fieldSql))
		fieldSqlTypeList := string_utils.ExplodeStr(string(r1), "$FST")
		if len(fieldSqlTypeList) == 1 {
			field.FieldSqlType = string(r1)
			field.FieldType = ConvertFieldType(string(r1))
		} else if len(fieldSqlTypeList) == 2 {
			if fst, ok := fieldSqlTypeMap["$FST" + fieldSqlTypeList[1]]; ok {
				field.FieldSqlType = fieldSqlTypeList[0] + fst
			}
			field.FieldType = ConvertFieldType(fieldSqlTypeList[0])
		}

		if !strings.Contains(fieldSql, "not null") {
			field.AllowNull = true
		}

		re, err = regexp.Compile("default [\\d\\w|\\$QM]*")
		if err != nil {
			return nil, err
		}
		r2 := re.FindAll([]byte(fieldSql), 1)
		if len(r2) > 0 {
			field.Default = strings.ReplaceAll(string(r2[0]), "default ", "")
			isVariable, err := regexp.Match("\\$", []byte(field.Default))
			if err != nil {
				return nil, err
			}
			if isVariable {
				if s, ok := strListMap[field.Default]; ok {
					field.Default = s
				}
			}
		}

		re, err = regexp.Compile("comment \\$QM[\\d]*")
		if err != nil {
			return nil, err
		}
		r3 := re.FindAll([]byte(fieldSql), 1)
		if len(r3) > 0 {
			variable := strings.ReplaceAll(string(r3[0]), "comment ", "")
			if s, ok := strListMap[variable]; ok {
				field.Remark = strings.Trim(strings.ReplaceAll(s, "\n", " "), "'")
			}
		}

		if strings.Contains(fieldSql, "primary key") {
			field.IsPrimary = true
			field.AllowNull = false
		}
		if len(field.Name) == 0 || len(field.FieldType) == 0 {
			continue
		}
		fieldsList = append(fieldsList, field)
	}
	return fieldsList, nil
}

func ConvertFieldType(fieldType string) string {
	fieldType = strings.ToLower(fieldType)
	switch fieldType {
	case "varchar", "date", "text", "char", "tinytext", "blob", "mediumtext", "mediumblob", "longtext",
		"longblob", "datetime", "time", "nchar", "nvarchar", "ntext":
		return "string"
	case "tinyint", "smallint", "mediumint", "int", "bit":
		return "int32"
	case "bigint", "numeric":
		return "int64"
	case "float", "decimal", "double", "money", "smallmoney":
		return "float64"
	case "timestamp":
		return "time.Time"
	default:
		return "string"
	}
}
