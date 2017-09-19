package models

import (
	"time"
	"strings"
)

func Expiration() string {
	now := time.Now()
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	newtime:=hh1.Format("2006-01-02 15:04:05")
	Xtime:=strings.Split(newtime, "-")
	var Xtime_english string
	if Xtime[1] == "01" {
		Xtime_english="Jan"
	}	
	if Xtime[1] == "02" {
                Xtime_english="Feb"
	}

	if Xtime[1] == "03" {
                Xtime_english="Mar"
        }

	if Xtime[1] == "04" {
                Xtime_english="Apr"
        }

	if Xtime[1] == "05" {
                Xtime_english="May"
        }

	if Xtime[1] == "06" {
                Xtime_english="Jun"
        }

	if Xtime[1] == "07" {
                Xtime_english="Jul"
        }

	if Xtime[1] == "08" {
                Xtime_english="Aug"
        }

	if Xtime[1] == "09" {
                Xtime_english="Sep"
        }

	if Xtime[1] == "10" {
                Xtime_english="Oct"
        }
	
	if Xtime[1] == "11" {
                Xtime_english="Nov"
        }
	if Xtime[1] == "12" {
                Xtime_english="Dec"
        }
	temp1 :=strings.Split(Xtime[2], " ")
	rt_time := temp1[0]+" "+Xtime_english+" "+Xtime[0] +" "+temp1[1]
	return string(rt_time) 

}
