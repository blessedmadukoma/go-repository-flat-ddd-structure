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

type AccountVerificationSuccessfulInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

type AccountVerificationFailedInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Reason    string `json:"reason"`
}

func AccountVerificationSuccessfulTask(i AccountVerificationSuccessfulInput) error {
	payload, err := json.Marshal(i)
	if err != nil {
		return err
	}

	info, err := client.Enqueue(asynq.NewTask(TypeAccountVerificationSuccessfulEmail, payload), asynq.Retention(24*time.Hour), asynq.Queue(TypeQueueDefault))
	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", *info)

	return nil
}

func (t Task) HandleAccountVerificationSuccessfulTask(ctx context.Context, a *asynq.Task) error {
	var p AccountVerificationSuccessfulInput
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return mail.Send(p.Email,
		"Account verified",
		"templates/accounts/successful_verification.html",
		map[string]string{"first_name": p.FirstName},
	)
}

func AccountVerificationFailedTask(i AccountVerificationFailedInput) error {
	payload, err := json.Marshal(i)
	if err != nil {
		return err
	}

	info, err := client.Enqueue(asynq.NewTask(TypeAccountVerificationFailedEmail, payload), asynq.Retention(24*time.Hour), asynq.Queue(TypeQueueDefault))
	if err != nil {
		return err
	}

	log.Printf(" [*] Successfully enqueued task: %+v", *info)

	return nil
}

func (t Task) HandleAccountVerificationFailedTask(ctx context.Context, a *asynq.Task) error {
	var p AccountVerificationFailedInput
	if err := json.Unmarshal(a.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	return mail.Send(p.Email,
		"Failed account verification",
		"templates/accounts/failed_verification.html",
		map[string]string{"first_name": p.FirstName, "reason": p.Reason},
	)
}
