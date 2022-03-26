// @author zan8in
// @date 2022-03-26 22:55
package allpasswords

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zan8in/allpasswords/pkg/utils"
)

type (
	Allpasswords struct {
		Options Options
		Result  Result
	}

	Result struct {
		FullName         []string
		ShortName        []string
		Phone            []string
		CompanyFullName  []string
		CompanyShortName []string
		JobNumber        []string
	}
)

func New(options Options) *Allpasswords {
	ap := &Allpasswords{Options: options}

	return ap
}

func (aps *Allpasswords) Password() (string, error) {
	aps.GetFullName()
	aps.GetShortName()
	aps.GetPhone()
	aps.GetCompanyFullName()
	aps.GetCompanyShortName()

	filename := "password" + utils.GetDate() + ".txt"
	fileHandle, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()
	buf := bufio.NewWriter(fileHandle)
	defer buf.Flush()

	for _, v := range aps.Result.FullName {
		buf.WriteString(v + "\n")
	}

	for _, v := range aps.Result.ShortName {
		buf.WriteString(v + "\n")
	}

	for _, v := range aps.Result.Phone {
		buf.WriteString(v + "\n")
	}
	for _, v := range aps.Result.CompanyFullName {
		buf.WriteString(v + "\n")
	}
	for _, v := range aps.Result.CompanyShortName {
		buf.WriteString(v + "\n")
	}
	return filename, nil
}

func (aps *Allpasswords) Account() (string, error) {
	aps.GetJobNumber()

	filename := "account" + utils.GetDate() + ".txt"
	fileHandle, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()
	buf := bufio.NewWriter(fileHandle)
	defer buf.Flush()

	for _, v := range aps.Result.JobNumber {
		buf.WriteString(v + "\n")
	}

	return filename, nil
}

func (aps *Allpasswords) GetFullName() {

	aps.Result.FullName = []string{}
	newfullname := []string{}

	if len(aps.Options.FullName) > 0 {

		for _, fullname := range aps.Options.FullName {
			if strings.Contains(fullname, " ") {
				newfullname = append(newfullname, strings.ReplaceAll(fullname, " ", "")) // "zhang san" to "zhangsan"
				f2slice := strings.Split(fullname, " ")
				if len(f2slice) > 0 {
					for _, v := range f2slice {
						newfullname = append(newfullname, strings.TrimSpace(v)) // zhang  or  san
					}
				}
			} else {
				newfullname = append(newfullname, fullname)
			}
		}

		aps.Result.FullName = append(aps.Result.FullName, newfullname...)
		for _, fullname := range newfullname {
			aps.Result.FullName = append(aps.Result.FullName, addNumber(fullname)...)
			aps.Result.FullName = append(aps.Result.FullName, addSymbolNumber(fullname)...)
			aps.Result.FullName = append(aps.Result.FullName, addNumberSymbol(fullname)...)
		}

		if len(aps.Options.BrithDay) > 0 {
			aps.Result.FullName = append(aps.Result.FullName, aps.Options.BrithDay...)
			for _, fullname := range newfullname {
				aps.Result.FullName = append(aps.Result.FullName, addBirthDay(fullname, aps.Options.BrithDay)...)
				aps.Result.FullName = append(aps.Result.FullName, addSymbolBirthDay(fullname, aps.Options.BrithDay)...)
				aps.Result.FullName = append(aps.Result.FullName, addBirthDaySymbol(fullname, aps.Options.BrithDay)...)
			}
		}

		if len(aps.Options.JobNumber) > 0 {
			aps.Result.FullName = append(aps.Result.FullName, aps.Options.JobNumber...)
			for _, fullname := range newfullname {
				aps.Result.FullName = append(aps.Result.FullName, addJobNumber(fullname, aps.Options.JobNumber)...)
				aps.Result.FullName = append(aps.Result.FullName, addSymbolJobNumber(fullname, aps.Options.JobNumber)...)
				aps.Result.FullName = append(aps.Result.FullName, addJobNumberSymbol(fullname, aps.Options.JobNumber)...)
			}
		}
		if len(aps.Options.Phone) > 0 {
			aps.Result.FullName = append(aps.Result.FullName, aps.Options.Phone...)
			for _, fullname := range newfullname {
				aps.Result.FullName = append(aps.Result.FullName, addPhone(fullname, aps.Options.Phone)...)
				aps.Result.FullName = append(aps.Result.FullName, addSymbolPhone(fullname, aps.Options.Phone)...)
				aps.Result.FullName = append(aps.Result.FullName, addPhoneSymbol(fullname, aps.Options.Phone)...)
			}
		}
	}
}

func (aps *Allpasswords) GetShortName() {
	aps.Result.ShortName = []string{}
	newshortname := []string{}
	if len(aps.Options.ShortName) > 0 {

		for _, shortname := range aps.Options.ShortName {
			if strings.Contains(shortname, " ") {
				newshortname = append(newshortname, strings.ReplaceAll(shortname, " ", ""))
				f2slice := strings.Split(shortname, " ")
				if len(f2slice) > 0 {
					for _, v := range f2slice {
						newshortname = append(newshortname, strings.TrimSpace(v))
					}
				}
			} else {
				newshortname = append(newshortname, shortname)
			}
		}

		aps.Result.ShortName = append(aps.Result.ShortName, newshortname...)
		for _, shortname := range newshortname {
			aps.Result.ShortName = append(aps.Result.ShortName, addNumber(shortname)...)
			aps.Result.ShortName = append(aps.Result.ShortName, addSymbolNumber(shortname)...)
			aps.Result.ShortName = append(aps.Result.ShortName, addNumberSymbol(shortname)...)
		}

		if len(aps.Options.BrithDay) > 0 {
			aps.Result.ShortName = append(aps.Result.ShortName, aps.Options.BrithDay...)
			for _, shortname := range newshortname {
				aps.Result.ShortName = append(aps.Result.ShortName, addBirthDay(shortname, aps.Options.BrithDay)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addSymbolBirthDay(shortname, aps.Options.BrithDay)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addBirthDaySymbol(shortname, aps.Options.BrithDay)...)
			}
		}

		if len(aps.Options.JobNumber) > 0 {
			aps.Result.ShortName = append(aps.Result.ShortName, aps.Options.JobNumber...)
			for _, shortname := range newshortname {
				aps.Result.ShortName = append(aps.Result.ShortName, addJobNumber(shortname, aps.Options.JobNumber)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addSymbolJobNumber(shortname, aps.Options.JobNumber)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addJobNumberSymbol(shortname, aps.Options.JobNumber)...)
			}
		}
		if len(aps.Options.Phone) > 0 {
			aps.Result.ShortName = append(aps.Result.ShortName, aps.Options.Phone...)
			for _, shortname := range newshortname {
				aps.Result.ShortName = append(aps.Result.ShortName, addPhone(shortname, aps.Options.Phone)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addSymbolPhone(shortname, aps.Options.Phone)...)
				aps.Result.ShortName = append(aps.Result.ShortName, addPhoneSymbol(shortname, aps.Options.Phone)...)
			}
		}
	}
}

func (aps *Allpasswords) GetPhone() {
	if len(aps.Options.Phone) > 0 {
		aps.Result.Phone = []string{}
		newfullname := []string{}
		if len(aps.Options.FullName) > 0 {

			for _, fullname := range aps.Options.FullName {
				if strings.Contains(fullname, " ") {
					newfullname = append(newfullname, strings.ReplaceAll(fullname, " ", "")) // "zhang san" to "zhangsan"
					f2slice := strings.Split(fullname, " ")
					if len(f2slice) > 0 {
						for _, v := range f2slice {
							newfullname = append(newfullname, strings.TrimSpace(v)) // zhang  or  san
						}
					}
				} else {
					newfullname = append(newfullname, fullname)
				}
			}
		}
		newshortname := []string{}
		if len(aps.Options.ShortName) > 0 {

			for _, shortname := range aps.Options.ShortName {
				if strings.Contains(shortname, " ") {
					newshortname = append(newshortname, strings.ReplaceAll(shortname, " ", ""))
					f2slice := strings.Split(shortname, " ")
					if len(f2slice) > 0 {
						for _, v := range f2slice {
							newshortname = append(newshortname, strings.TrimSpace(v))
						}
					}
				} else {
					newshortname = append(newshortname, shortname)
				}
			}
		}
		aps.Result.Phone = append(aps.Result.Phone, aps.Options.Phone...)
		for _, phone := range aps.Options.Phone {
			if len(newfullname) > 0 {
				aps.Result.Phone = append(aps.Result.Phone, phoneFullName(phone, newfullname)...)
				aps.Result.Phone = append(aps.Result.Phone, phoneSymbolFullName(phone, newfullname)...)
			}
			if len(newshortname) > 0 {
				aps.Result.Phone = append(aps.Result.Phone, phoneShortName(phone, newshortname)...)
				aps.Result.Phone = append(aps.Result.Phone, phoneSymbolShortName(phone, newshortname)...)
			}
			aps.Result.Phone = append(aps.Result.Phone, phoneSymbol(phone)...)
		}
	}
}

func (aps *Allpasswords) GetCompanyFullName() {
	if len(aps.Options.CompanyFullName) > 0 {
		aps.Result.CompanyFullName = []string{}

		aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, aps.Options.CompanyFullName...)

		if len(aps.Options.CompanyTelephone) > 0 {
			for _, companyfullname := range aps.Options.CompanyFullName {
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNamePhone(companyfullname, aps.Options.CompanyTelephone)...)
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameSymbolPhone(companyfullname, aps.Options.CompanyTelephone)...)
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNamePhoneSymbol(companyfullname, aps.Options.CompanyTelephone)...)
			}
		}

		if len(aps.Options.JobNumber) > 0 {
			for _, companyfullname := range aps.Options.CompanyFullName {
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameJobNumber(companyfullname, aps.Options.JobNumber)...)
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameSymbolJobNumber(companyfullname, aps.Options.JobNumber)...)
				aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameJobNumberSymbol(companyfullname, aps.Options.JobNumber)...)
			}
		}

		for _, companyfullname := range aps.Options.CompanyFullName {
			aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameYears(companyfullname)...)
			aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameSymbolYears(companyfullname)...)
			aps.Result.CompanyFullName = append(aps.Result.CompanyFullName, companyFullNameYearsSymbol(companyfullname)...)

		}
	}
}

func (aps *Allpasswords) GetCompanyShortName() {
	if len(aps.Options.CompanyShortName) > 0 {
		aps.Result.CompanyShortName = []string{}

		aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, aps.Options.CompanyShortName...)

		if len(aps.Options.CompanyTelephone) > 0 {
			for _, companyshortname := range aps.Options.CompanyShortName {
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNamePhone(companyshortname, aps.Options.CompanyTelephone)...)
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameSymbolPhone(companyshortname, aps.Options.CompanyTelephone)...)
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNamePhoneSymbol(companyshortname, aps.Options.CompanyTelephone)...)
			}
		}

		if len(aps.Options.JobNumber) > 0 {
			for _, companyshortname := range aps.Options.CompanyShortName {
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameJobNumber(companyshortname, aps.Options.JobNumber)...)
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameSymbolJobNumber(companyshortname, aps.Options.JobNumber)...)
				aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameJobNumberSymbol(companyshortname, aps.Options.JobNumber)...)
			}
		}

		for _, companyshortname := range aps.Options.CompanyShortName {
			aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameYears(companyshortname)...)
			aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameSymbolYears(companyshortname)...)
			aps.Result.CompanyShortName = append(aps.Result.CompanyShortName, companyShortNameYearsSymbol(companyshortname)...)

		}
	}
}

func (aps *Allpasswords) GetJobNumber() {
	result := []string{}
	for i := 0; i <= 10; i++ {
		result = append(result, strconv.Itoa(i))
		result = append(result, "0"+strconv.Itoa(i))
		result = append(result, "00"+strconv.Itoa(i))
		result = append(result, "000"+strconv.Itoa(i))
		result = append(result, "0000"+strconv.Itoa(i))

		result = append(result, "10"+strconv.Itoa(i))
		result = append(result, "100"+strconv.Itoa(i))
		result = append(result, "1000"+strconv.Itoa(i))
	}
	result = append(result, RuleJobNumber...)

	if len(aps.Options.JobNumber) > 0 {
		for _, jn := range aps.Options.JobNumber {
			flag := false
			for _, v := range result {
				if jn == v {
					flag = true
				}
			}
			if !flag {
				result = append(result, jn)
			}
		}
	}
	aps.Result.JobNumber = append(aps.Result.JobNumber, result...)
}

func companyShortNameYearsSymbol(str string) []string {
	result := []string{}
	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		result = append(result, addSymbol(str+strconv.Itoa(i))...)
	}
	return result
}
func companyShortNameSymbolYears(str string) []string {
	result := []string{}
	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+strconv.Itoa(i))
		}
	}
	return result
}
func companyShortNameYears(str string) []string {
	result := []string{}

	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		result = append(result, str+strconv.Itoa(i))
		result = append(result, firstUpper(str)+strconv.Itoa(i))
	}

	return result
}

func companyShortNameJobNumber(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		result = append(result, str+jobnumber)
		result = append(result, firstUpper(str)+jobnumber)
	}
	return result
}
func companyShortNameSymbolJobNumber(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+jobnumber)
		}
	}
	return result
}
func companyShortNameJobNumberSymbol(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		result = append(result, addSymbol(str+jobnumber)...)
	}
	return result
}

func companyShortNamePhone(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		result = append(result, str+telephone)
		result = append(result, firstUpper(str)+telephone)
	}
	return result
}
func companyShortNameSymbolPhone(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+telephone)
		}
	}
	return result
}
func companyShortNamePhoneSymbol(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		result = append(result, addSymbol(str+telephone)...)
	}
	return result
}

func companyFullNameYearsSymbol(str string) []string {
	result := []string{}
	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		result = append(result, addSymbol(str+strconv.Itoa(i))...)
	}
	return result
}
func companyFullNameSymbolYears(str string) []string {
	result := []string{}
	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+strconv.Itoa(i))
		}
	}
	return result
}
func companyFullNameYears(str string) []string {
	result := []string{}

	for i := RuleCompanyYearsFrom; i <= time.Now().Year(); i++ {
		result = append(result, str+strconv.Itoa(i))
		result = append(result, firstUpper(str)+strconv.Itoa(i))
	}

	return result
}

func companyFullNameJobNumber(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		result = append(result, str+jobnumber)
		result = append(result, firstUpper(str)+jobnumber)
	}
	return result
}
func companyFullNameSymbolJobNumber(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+jobnumber)
		}
	}
	return result
}
func companyFullNameJobNumberSymbol(str string, jobNumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobNumbers {
		result = append(result, addSymbol(str+jobnumber)...)
	}
	return result
}

func companyFullNamePhone(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		result = append(result, str+telephone)
		result = append(result, firstUpper(str)+telephone)
	}
	return result
}
func companyFullNameSymbolPhone(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+telephone)
		}
	}
	return result
}
func companyFullNamePhoneSymbol(str string, companyTelephones []string) []string {
	result := []string{}
	for _, telephone := range companyTelephones {
		result = append(result, addSymbol(str+telephone)...)
	}
	return result
}

func phoneFullName(str string, fullnames []string) []string {
	result := []string{}
	for _, fullname := range fullnames {
		result = append(result, str+fullname)
		result = append(result, str+firstUpper(fullname))
	}
	return result
}
func phoneSymbolFullName(str string, fullnames []string) []string {
	result := []string{}
	for _, fullname := range fullnames {
		for _, symstr := range addSymbol(fullname) {
			result = append(result, str+symstr)
		}
	}
	return result
}
func phoneShortName(str string, shortnames []string) []string {
	result := []string{}
	for _, shortname := range shortnames {
		result = append(result, str+shortname)
		result = append(result, str+firstUpper(shortname))
	}
	return result
}
func phoneSymbolShortName(str string, shortnames []string) []string {
	result := []string{}
	for _, shortname := range shortnames {
		for _, symstr := range addSymbol(shortname) {
			result = append(result, str+symstr)
		}
	}
	return result
}
func phoneSymbol(str string) []string {
	result := []string{}
	for _, sym := range RuleSymbol {
		result = append(result, str+sym)
	}
	return result
}

func addPhone(str string, phones []string) []string {
	result := []string{}
	for _, phone := range phones {
		result = append(result, str+phone)
		result = append(result, firstUpper(str)+phone)
	}
	return result
}
func addSymbolPhone(str string, phones []string) []string {
	result := []string{}
	for _, phone := range phones {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+phone)
		}
	}
	return result
}
func addPhoneSymbol(str string, phones []string) []string {
	result := []string{}
	for _, phone := range phones {
		result = append(result, addSymbol(str+phone)...)
	}
	return result
}

func addJobNumber(str string, jobnumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobnumbers {
		result = append(result, str+jobnumber)
		result = append(result, firstUpper(str)+jobnumber)
	}
	return result
}
func addSymbolJobNumber(str string, jobnumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobnumbers {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+jobnumber)
		}
	}
	return result
}
func addJobNumberSymbol(str string, jobnumbers []string) []string {
	result := []string{}
	for _, jobnumber := range jobnumbers {
		result = append(result, addSymbol(str+jobnumber)...)
	}
	return result
}

func addBirthDay(str string, births []string) []string {
	result := []string{}
	for _, birth := range births {
		result = append(result, str+birth)
		result = append(result, firstUpper(str)+birth)
	}
	return result
}
func addSymbolBirthDay(str string, births []string) []string {
	result := []string{}
	for _, birth := range births {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+birth)
		}
	}
	return result
}
func addBirthDaySymbol(str string, births []string) []string {
	result := []string{}
	for _, birth := range births {
		result = append(result, addSymbol(str+birth)...)
	}
	return result
}

func addNumber(str string) []string {
	result := []string{}
	for _, number := range RuleNumber {
		result = append(result, str+number)
		result = append(result, firstUpper(str)+number)
	}
	return result
}

func addSymbolNumber(str string) []string {
	result := []string{}
	for _, number := range RuleNumber {
		for _, symstr := range addSymbol(str) {
			result = append(result, symstr+number)
		}
	}
	return result
}

func addNumberSymbol(str string) []string {
	result := []string{}
	for _, number := range RuleNumber {
		result = append(result, addSymbol(str+number)...)
	}
	return result
}

func addSymbol(str string) []string {
	result := []string{}
	for _, sym := range RuleSymbol {
		result = append(result, str+sym)
		result = append(result, firstUpper(str)+sym)
	}
	return result
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
