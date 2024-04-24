package tasks

import (
	"context"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/services"
	"goRepositoryPattern/util"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeRegisterOtpEmail                   = "email:register_otp"
	TypePasswordResetEmail                 = "email:password_reset"
	TypePasswordChangeEmail                = "email:password_change"
	TypeAccountVerificationSuccessfulEmail = "email:account_verification_success"
	TypeAccountVerificationFailedEmail     = "email:account_verification_failed"
)

const (
	TypeQueueCritical = "critical"
	TypeQueueDefault  = "default"
	TypeQueueLow      = "low"
)

type Task struct {
	repo    *repository.Repository
	service *services.Service
}

func NewTask(repo *repository.Repository, service *services.Service) Task {
	return Task{repo, service}
}

var client *asynq.Client

func initializeClient(opt asynq.RedisClientOpt) *asynq.Client {
	return asynq.NewClient(opt)
}

func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		log.Printf("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}

func StartWorker(t Task) error {

	config, err := util.GetRedisConn()
	if err != nil {
		return err
	}

	client = initializeClient(*config)

	srv := asynq.NewServer(
		*config,
		asynq.Config{Concurrency: 10, Queues: map[string]int{
			TypeQueueCritical: 6,
			TypeQueueDefault:  3,
			TypeQueueLow:      1,
		}},
	)

	mux := asynq.NewServeMux()
	mux.Use(loggingMiddleware)

	t.registerTasks(mux)
	// t.registerCrons(mux)

	go func() {
		if err := srv.Run(mux); err != nil {
			log.Fatal(err)
		}
	}()

	// provider := &FileBasedConfigProvider{filename: "./cron.yml"}

	// mgr, err := asynq.NewPeriodicTaskManager(
	// 	asynq.PeriodicTaskManagerOpts{
	// 		RedisConnOpt:               *config,
	// 		PeriodicTaskConfigProvider: provider,
	// 		SyncInterval:               10 * time.Second,
	// 	})
	// if err != nil {
	// 	return err
	// }

	// if err := mgr.Run(); err != nil {
	// 	return err
	// }

	return nil
}

func (t Task) registerTasks(m *asynq.ServeMux) {
	m.HandleFunc(TypeRegisterOtpEmail, t.HandleRegisterOtpTask)
	m.HandleFunc(TypePasswordResetEmail, t.HandlePasswordResetTask)
	m.HandleFunc(TypePasswordChangeEmail, t.HandlePasswordChangeTask)
	m.HandleFunc(TypeAccountVerificationSuccessfulEmail, t.HandleAccountVerificationSuccessfulTask)
	m.HandleFunc(TypeAccountVerificationFailedEmail, t.HandleAccountVerificationFailedTask)
}
