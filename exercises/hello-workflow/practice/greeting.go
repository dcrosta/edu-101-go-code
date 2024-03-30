package hello

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5000 * time.Millisecond,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	future := workflow.ExecuteActivity(ctx, GreetInSpanish)

	var result string
	err := future.Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func GreetInSpanish(ctx context.Context, name string) (string, error) {
	url := "https://web.archive.org/"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("got response code %d, expected 200", resp.StatusCode)
	}

	return "¡Hola " + name + "!", nil
}
