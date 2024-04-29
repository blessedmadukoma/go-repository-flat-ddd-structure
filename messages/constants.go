package messages

var AccountTokenTypeEmailConfirmationKey = 1
var AccountTokenTypeEmailConfirmation = "email_confirmation"
var AccountTokenTypePasswordReset = "password_reset"
var AccountTokenTypePasswordResetKey = 2
var AccountTokenTypeSmsConfirmationKey = 3
var AccountTokenTypeSmsConfirmation = "sms_confirmation"

func AccountTokenTypes() map[string]int {
	m := map[string]int{
		AccountTokenTypeEmailConfirmation: AccountTokenTypeEmailConfirmationKey,
		AccountTokenTypePasswordReset:     AccountTokenTypePasswordResetKey,
		AccountTokenTypeSmsConfirmation:   AccountTokenTypeSmsConfirmationKey,
	}
	return m
}

func GetAccountTokenTypeKey(t string) uint {
	if v, found := AccountTokenTypes()[t]; found {
		return uint(v)
	}
	return uint(AccountTokenTypeEmailConfirmationKey)
}

func GetAccountTokenTypeName(val int) string {
	for k, v := range AccountTokenTypes() {
		if v == val {
			return k
		}
	}
	return AccountTokenTypeEmailConfirmation
}
