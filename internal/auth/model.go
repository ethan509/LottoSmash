package auth

import (
	"time"
)

// TierCode 회원 등급 코드 상수
type TierCode string

const (
	TierGuest  TierCode = "GUEST"  // 게스트 (회원가입 안함)
	TierMember TierCode = "MEMBER" // 정회원 (회원가입 완료)
	TierGold   TierCode = "GOLD"   // 골드 (월정액 구독)
	TierVIP    TierCode = "VIP"    // VIP (특별 등급)
)

// MembershipTier 회원 등급 메타 정보
type MembershipTier struct {
	ID          int       `json:"id"`
	Code        TierCode  `json:"code"`
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	Description *string   `json:"description,omitempty"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID           int64           `json:"id"`
	DeviceID     *string         `json:"device_id,omitempty"`
	Email        *string         `json:"email,omitempty"`
	PasswordHash *string         `json:"-"`
	LottoTier    int             `json:"lotto_tier"`
	Tier         *MembershipTier `json:"tier,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

// IsMember 정회원 이상인지 확인 (하위 호환용)
func (u *User) IsMember() bool {
	return u.LottoTier >= 2 // MEMBER 이상
}

// HasTier 특정 등급 이상인지 확인
func (u *User) HasTier(tierLevel int) bool {
	if u.Tier != nil {
		return u.Tier.Level >= tierLevel
	}
	return u.LottoTier >= tierLevel
}

type RefreshToken struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type EmailVerification struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}

// Request/Response DTOs

type GuestLoginRequest struct {
	DeviceID string `json:"device_id"`
}

type EmailRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LinkEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type SendVerificationRequest struct {
	Email string `json:"email"`
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

type PasswordChangeRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type UserResponse struct {
	ID    int64        `json:"id"`
	Email *string      `json:"email,omitempty"`
	Tier  TierResponse `json:"tier"`
}

type TierResponse struct {
	Code  TierCode `json:"code"`
	Name  string   `json:"name"`
	Level int      `json:"level"`
}
