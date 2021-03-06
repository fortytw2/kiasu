// +build integration

package discollect

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func QueueTests(t *testing.T, q Queue, resetFunc func()) func(t *testing.T) {
	return func(t *testing.T) {

		exID := uuid.New()
		cases := []struct {
			name  string
			run   func(q Queue) error
			check func(q Queue) error
		}{
			{
				"one-in-out",
				func(q Queue) error {
					err := q.Push(context.TODO(), []*QueuedTask{
						{
							ScrapeID: exID,
						},
					})

					return err
				},
				func(q Queue) error {
					qt, err := q.Pop(context.TODO())
					if err != nil {
						return err
					}

					if qt == nil {
						return errors.New("got a nil queued task when one should exist")
					}

					if qt.ScrapeID != exID {
						return errors.New("did not get the same thing back")
					}

					return nil
				},
			},
			{
				"none-ever-queued",
				func(q Queue) error {
					return nil
				},
				func(q Queue) error {
					qt, err := q.Pop(context.TODO())
					if err != nil {
						return err
					}

					if qt != nil {
						return errors.New("got a task where none exists")
					}

					return nil
				},
			},
			{
				"queued-but-none-exist",
				func(q Queue) error {
					err := q.Push(context.TODO(), []*QueuedTask{
						{
							ScrapeID: exID,
						},
					})

					return err
				},
				func(q Queue) error {
					qt1, err := q.Pop(context.TODO())
					if err != nil {
						return err
					}

					if qt1 == nil {
						return errors.New("got nil for task 1")
					}

					ss, err := q.Status(context.TODO(), exID)
					if err != nil {
						return err
					}

					if ss.TotalTasks != 1 {
						return fmt.Errorf("wrong number of total tasks after popping one, got %d, want 1", ss.TotalTasks)
					}

					qt, err := q.Pop(context.TODO())
					if err != nil || qt != nil {
						return errors.New("got a task when non currently exist")
					}

					ss, err = q.Status(context.TODO(), exID)
					if err != nil {
						return err
					}

					if ss.TotalTasks != 1 {
						return fmt.Errorf("wrong number of total tasks, got %d, want 1", ss.TotalTasks)
					}

					err = q.Finish(context.TODO(), qt1)
					if err != nil {
						return err
					}

					ss, err = q.Status(context.TODO(), exID)
					if err != nil {
						return err
					}

					if ss.TotalTasks != 1 || ss.CompletedTasks != 1 {
						return errors.New("wrong scrape status after finishing task")
					}

					return nil
				},
			},
			{
				"correct-inflight",
				func(q Queue) error {
					err := q.Push(context.TODO(), []*QueuedTask{
						{
							ScrapeID: exID,
						},
					})

					return err
				},
				func(q Queue) error {
					qt1, err := q.Pop(context.TODO())
					if err != nil {
						return err
					}

					qt, err := q.Pop(context.TODO())
					if err != nil {
						return err
					}

					if qt != nil {
						return errors.New("got a task where none exists")
					}

					ss, err := q.Status(context.TODO(), exID)
					if err != nil {
						return err
					}

					if ss.InFlightTasks != 1 {
						return errors.New("multiple inflight tasks when none should exist")
					}

					err = q.Finish(context.TODO(), qt1)
					if err != nil {
						return err
					}

					ss, err = q.Status(context.TODO(), exID)
					if err != nil {
						return err
					}

					if ss.InFlightTasks != 0 {
						return errors.New("no tasks are inflight but one exists")
					}

					return nil
				},
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				resetFunc()

				err := c.run(q)
				if err != nil {
					t.Fatal(err)
				}

				err = c.check(q)
				if err != nil {
					t.Fatal(err)
				}
			})
		}
	}
}
