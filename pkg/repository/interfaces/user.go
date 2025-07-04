package interfaces

import (
	"ecommerce_clean_arch/pkg/domain"
	"ecommerce_clean_arch/pkg/utils/models"
	"time"
)

type UserRepository interface {
	IsEmailExists(email string) bool
	IsPhoneExists(phone string) bool
	SaveTempUser(user models.User) error
	GetTempUserByEmail(email string) (models.TempUser, error)
	DeleteTempUser(email string) error
	SaveOrUpdateOTP(email string, otp string, otpExpiry time.Time) error
	GetOTP(email string) (string, time.Time, error)
	GetOTPByEmail(email string) (models.OTP, error)
	VerifyOTPAndMoveUser(email string, otp string) error
	SaveOTP(email, otp string, expiry time.Time) error
	DeleteOTP(email string) error
	UpdateOTP(otp models.OTP) error
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetEmailByOTP(otp string) (string, error)
	UnblockUser(email string) error
	GetProducts() ([]models.ProductResponse, error)
	ListCategory() ([]domain.Category, error)
	UserProfile(userID int) (*models.User, error)
	UpdateProfile(editProfile models.User) (*models.User, error)
	GetUserByID(userID int) (models.User, error)
	ForgotPassword(email string, newPassword string) error
	UpdatePassword(userID int, newPassword string) error
	GetPassword(userID int) (models.User, error)
	AddAddress(userID int, address models.AddAddress) (models.AddAddress, error)
	UpdateAddress(userID int, address models.AddAddress) (models.AddAddress, error)
	DeleteAddress(userID int) error
	GetAllAddresses(userID int) ([]domain.Address, error)
}
