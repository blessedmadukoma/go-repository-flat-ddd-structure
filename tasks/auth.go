package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"goRepositoryPattern/integrations/mail"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type RegisterOtpInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	OTP       string `json:"otp"`
}

type PasswordResetInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Code      string `json:"code"`
}

type PasswordChangeInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

func RegisterOtpTask(i RegisterOtpInput) error {
	payload, err := json.Marshal(i)
	if err != nil {
		return err
	}

	info, err := client.Enqueue(asynq.NewTask(TypeRegisterOtpEmail, payload), asynq.Retention(24*time.Hour), asynq.Queue(TypeQueueCritical))
	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", *info)

	return nil
}

func (t Task) HandleRegisterOtpTask(ctx context.Context, a *asynq.Task) error {
	var p RegisterOtpInput
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return mail.Send(p.Email,
		"Verify your account",
		"templates/auth/register.html",
		map[string]string{"first_name": p.FirstName, "code": p.OTP},
	)
}
func PasswordResetTask(i PasswordResetInput) error {
	payload, err := json.Marshal(i)
	if err != nil {
		return err
	}

	info, err := client.Enqueue(asynq.NewTask(TypePasswordResetEmail, payload), asynq.Retention(24*time.Hour))
	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", *info)

	return nil
}

func (t Task) HandlePasswordResetTask(ctx context.Context, a *asynq.Task) error {
	var p PasswordResetInput
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return mail.Send(p.Email,
		"Forgot your password",
		"templates/auth/password_reset.html",
		map[string]string{"first_name": p.FirstName, "code": p.Code},
	)
}

func PasswordChangeTask(i PasswordChangeInput) error {
	payload, err := json.Marshal(i)
	if err != nil {
		return err
	}

	info, err := client.Enqueue(asynq.NewTask(TypePasswordChangeEmail, payload), asynq.Retention(24*time.Hour))
	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", *info)

	return nil
}

func (t Task) HandlePasswordChangeTask(ctx context.Context, a *asynq.Task) error {
	var p PasswordChangeInput
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return mail.Send(p.Email,
		"Password changed",
		"templates/auth/password_change.html",
		map[string]string{"first_name": p.FirstName},
	)
}
