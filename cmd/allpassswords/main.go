package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zan8in/allpasswords/pkg/allpasswords"
)

func main() {

	options := allpasswords.Options{}

	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("FullName[姓名全拼](例如：zhang san)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.FullName = append(options.FullName, strings.ToLower(input.Text()))
		fmt.Printf("FullName[姓名全拼]：")
	}

	fmt.Printf("ShortName[姓名简拼](例如：zs)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.ShortName = append(options.ShortName, strings.ToLower(input.Text()))
		fmt.Printf("ShortName[姓名简拼]：")
	}

	fmt.Printf("BrithDay[生日](例如：19961023)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.ShortName = append(options.ShortName, strings.ToLower(input.Text()))
		fmt.Printf("BrithDay[生日]：")
	}

	fmt.Printf("Phone[个人手机号]：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.Phone = append(options.Phone, strings.ToLower(input.Text()))
		fmt.Printf("Phone[个人手机号]：")
	}

	fmt.Printf("CompanyFullName[公司全拼](例如：beijingdaxue)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.CompanyFullName = append(options.CompanyFullName, strings.ToLower(input.Text()))
		fmt.Printf("CompanyFullName[公司全拼]：")
	}

	fmt.Printf("CompanyShortName[公司简拼](例如：bjdx)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.CompanyShortName = append(options.CompanyShortName, strings.ToLower(input.Text()))
		fmt.Printf("CompanyShortName[公司简拼]：")
	}

	fmt.Printf("CompanyTelephone[公司电话]：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.CompanyTelephone = append(options.CompanyTelephone, strings.ToLower(input.Text()))
		fmt.Printf("CompanyTelephone[公司电话]：")
	}

	fmt.Printf("JobNumber[员工号码](例如：9521)：")
	for input.Scan() {
		if len(input.Text()) == 0 {
			break
		}
		options.JobNumber = append(options.JobNumber, strings.ToLower(input.Text()))
		fmt.Printf("JobNumber[员工号码]：")
	}

	aps := allpasswords.New(options)

	filename, err := aps.Password()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("生成密码文件:", filename)

	filename, err = aps.Account()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("生成账号文件:", filename)
}
