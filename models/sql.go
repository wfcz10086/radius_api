package models
import (
	"github.com/astaxie/beego/orm"
)

type Radcheck struct {
	Id int
	Username  string 
	Attribute  string
	Op     string 
	Value    string
}

func InsertPassword(sk *Radcheck) (bool,Radcheck) {
	o := orm.NewOrm()
	var rad Radcheck
	err := o.QueryTable("radcheck").Filter("username", sk.Username).Filter("attribute","Cleartext-Password").One(&rad)
        if err == orm.ErrNoRows {
    // 没有找到记录
	o.Insert(sk)
	} 
		return err != orm.ErrNoRows, rad
	
}

func InsertExpiration(sk *Radcheck) (bool,Radcheck) {
        o := orm.NewOrm()
        var rad Radcheck
        err := o.QueryTable("radcheck").Filter("username", sk.Username).Filter("attribute","Expiration").One(&rad)
        if err == orm.ErrNoRows {
    // 没有找到记录
        o.Insert(sk)
        }
                return err != orm.ErrNoRows, rad

}

func UpdateExpiration(sk *Radcheck){
        o := orm.NewOrm()
        o.QueryTable("radcheck").Filter("username", sk.Username).Filter("attribute","Expiration").Update(orm.Params{"Value":sk.Value,})

}
func UpdatePassword(sk *Radcheck) {
        o := orm.NewOrm()

	o.QueryTable("radcheck").Filter("username", sk.Username).Filter("attribute","Cleartext-Password").Update(orm.Params{"Value":sk.Value,})


}

