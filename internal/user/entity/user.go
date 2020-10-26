/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/20 17:49
 */
package entity

type UserEntity struct {
	ID        int64
	CompanyID int64
	Username  string
	Password  string
	Nickname  string
	Role      string
}
