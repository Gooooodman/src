package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"regexp"
	"strings"
)

type Tbles struct {
	RecordSid   string `json:"record_sid"`
	Server      int    `json:"server"`
	SAccountID  string `json:"s_account_id"`
	DAccountID  string `json:"d_account_id"`
	DCount      int    `json:"d_count"`
	DIP         string `json:"d_ip"`
	DItemID     int    `json:"d_item_id"`
	DItemName   string `json:"d_item_name"`
	DItemType   int    `json:"d_item_type"`
	DMac        string `json:"d_mac"`
	DMoney      int    `json:"d_money"`
	DRoleID     int    `json:"d_role_id"`
	DRoleName   string `json:"d_role_name"`
	DUUID       string `json:"d_uuid"`
	EventID     string `json:"event_id"`
	SCount      int    `json:"s_count"`
	SIP         string `json:"s_ip"`
	SItemID     int    `json:"s_item_id"`
	SItemName   string `json:"s_item_name"`
	SItemType   int    `json:"s_item_type"`
	SMac        string `json:"s_mac"`
	SMoney      int    `json:"s_money"`
	SRoleID     int    `json:"s_role_id"`
	SRoleName   string `json:"s_role_name"`
	TradeTime   int    `json:"trade_time"`
	Type        string `json:"type"`
	Incorrect   int    `json:"incorrect"`
	Content     string `json:"content"`
	EditTime    int    `json:"edit_time"`
	OpId        int    `json:"op_id"`
	ZoneId      int    `json:"zone_id"`
	GoodsDetail string `json:"goods_detail"`
}

func ReadLine(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bufio.NewReader(f)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", User, Passswd, Hostname, Port, Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("mysql conn fail :", err)
		return
	}

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			//读完了之后 break
			if err == io.EOF {
				break
			}
			return
		}
		line = strings.TrimSpace(line)
		r, _ := regexp.Compile("\\[.*\\],")
		line = r.ReplaceAllString(line, "")
		T := &Tbles{}
		json.Unmarshal([]byte(line), &T)
		stmt, err := db.Prepare(`INSERT tbl_trade (record_sid,server,s_account_id,s_role_name,s_role_id,s_item_type,s_item_id,s_item_name,s_count,s_money,s_ip,s_mac,event_id,d_account_id,d_role_name,d_role_id,d_item_type,d_item_id,d_item_name,d_count,d_money,d_ip,d_mac,d_uuid,type,trade_time,incorrect) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
		if err != nil {
			fmt.Println("Prepare fail .. ", err)
			return
		}
		_, err = stmt.Exec(T.RecordSid, T.Server, T.SAccountID, T.SRoleName, T.SRoleID, T.SItemType, T.SItemID, T.SItemName, T.SCount, T.SMoney, T.SIP, T.SMac, T.EventID, T.DAccountID, T.DRoleName, T.DRoleID, T.DItemType, T.DItemID, T.DItemName, T.DCount, T.DMoney, T.DIP, T.DMac, T.DUUID, T.Type, T.TradeTime, T.Incorrect)
		if err != nil {
			fmt.Println("Exec fail .. ", err)
			return
		}
		stmt.Close()

	}

	defer db.Close()
}

var LogFile string
var User string
var Passswd string
var Hostname string
var Port int

const (
	Database = "ssdbManager"
)

func main() {
	flag.StringVar(&LogFile, "logfile", "tbl_trade.log", "--logfile : 日志路径")
	flag.StringVar(&User, "u", "root", "-u :  mysql user")
	flag.StringVar(&Passswd, "p", "123456", "-p : mysql passwd")
	flag.StringVar(&Hostname, "h", "127.0.0.1", "-h : mysql hostname")
	flag.IntVar(&Port, "P", 3306, "-P : mysql port")
	flag.Parse()
	_, err := os.Stat(LogFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	ReadLine(LogFile)

}
