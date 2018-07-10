package stub

import (
	"context"
	"fmt"

	"github.com/tkashem/echo-operator/pkg/apis/echo/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1.EchoApp:
		if event.Deleted {
			return nil
		}

		echo := o
		deployment := deployment(echo)

		err := sdk.Create(deployment)
		if err != nil && !apierrors.IsAlreadyExists(err) {
			return fmt.Errorf("failed to create deployment: %v", err)
		}

		if err := ensure(echo, deployment); err != nil {
			return fmt.Errorf("failed to ensure deployment size mtches: %v", err)
		}

		if err := update(echo); err != nil {
			return fmt.Errorf("filed to update status echo status: %v", err)
		}
	}

	return nil
}
