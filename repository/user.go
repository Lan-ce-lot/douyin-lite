package repository

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              int64  `json:"id" gorm:"id,omitempty"`
	Name            string `json:"name" gorm:"comment:用户名"`
	Password        string `json:"password" gorm:"comment:用户密码"`
	Avatar          string `json:"avatar" gorm:"comment:用户头像"`
	BackgroundImage string `json:"background_image" gorm:"comment:用户背景主图"`
	Signature       string `json:"signature" gorm:"comment:用户签名"`
	FollowingCount  int64  `json:"follow_count" gorm:"comment: 关注总数"`
	FollowerCount   int64  `json:"follower_count" gorm:"comment:粉丝总数"`
	TotalFavorited  int64  `json:"total_favorited" gorm:"comment:获赞总数"`
	WorkCount       int64  `json:"work_count" gorm:"comment:作品总数"`
	FavoriteCount   int64  `json:"favorite_count" gorm:"comment:点赞总数"`
}

type UserDao struct {
}

var userDao *UserDao
var UserOnce sync.Once

func NewUserDaoInstance() *UserDao {
	UserOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Avatar = "https://douyin-lite.oss-cn-hangzhou.aliyuncs.com/avater/default.jpg"
	user.BackgroundImage = "https://douyin-lite.oss-cn-hangzhou.aliyuncs.com/background_image/default.jpg"
	user.Signature = "你所热爱的，就是你的生活"
	return nil
}

func (*UserDao) CreateUser(name string, followingCnt int64, followerCnt int64) error {
	newUser := User{
		Name:           name,
		FollowingCount: followingCnt,
		FollowerCount:  followerCnt,
	}
	err := db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (*UserDao) CreateRegisterUser(name string, password string) (*User, error) {
	newUser := User{
		Name:     name,
		Password: password,
	}
	err := db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (*UserDao) QueryIsUserExist(name string) (bool, error) {
	err := db.Where("name = ?", name).First(&User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, errors.New("数据库异常")
	}
	return true, nil
}

func (*UserDao) QueryLoginUser(name string, password string) (*User, error) {
	qUser := User{}
	err := db.Where("name = ? and password = ?", name, password).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账号或者密码不对")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}

func (*UserDao) QueryUserById(userId uint) (*User, error) {
	qUser := User{}
	err := db.Where("id = ?", userId).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("该用户不存在")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}
