package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/stori-challenge/pkg/logger"
)

var ()

// func getFilesStore() filestores.FileStore {
// 	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
// 	if err != nil {
// 		logger.GetLogger().WithField("message", "aws config loading error").Panic(err)
// 	}
// 	return filestores.NewAwsBucket(config.Instance.SecretsBucket, cfg)
// }

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, evente events.CloudWatchEvent) error {
	logger.GetLogger().Info("Initialize")
	// if err := secretsScubber.ScrubMaxStorageReached(); err != nil {
	// 	logger.GetLogger().Error(err)
	// 	return err
	// }
	logger.GetLogger().Info("Done")
	return nil
}
