package model
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

var (
	MyUserDao *UserDao
)

type UserDao struct{
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao){

	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){

	res,err := redis.String(conn.Do("HGET","users",id))
	if err != nil{
		if err == redis.ErrNil {//表示reids.hash中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return

}

func (this *UserDao) Login(userId int,userPwd string) (user *User,err error){
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil{
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}