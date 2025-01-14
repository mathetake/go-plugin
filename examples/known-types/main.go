package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/knqyf263/go-plugin/examples/known-types/known"
	"github.com/knqyf263/go-plugin/types/known/structpb"
	"github.com/knqyf263/go-plugin/types/known/timestamppb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	p, err := known.NewWellKnownPlugin(ctx, known.WellKnownPluginOption{Stdout: os.Stdout, Stderr: os.Stderr})
	if err != nil {
		return err
	}

	plugin, err := p.Load(ctx, "plugin/plugin.wasm")
	if err != nil {
		return err
	}

	value, err := structpb.NewValue(map[string]interface{}{
		"A": "Sushi",
		"B": "Tempura",
	})
	if err != nil {
		return err
	}

	start := timestamppb.Now()
	end := timestamppb.New(start.AsTime().Add(1 * time.Hour))
	reply, err := plugin.Diff(ctx, known.DiffRequest{
		Value: value,
		Start: start,
		End:   end,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Duration: %s\n", reply.GetDuration().AsDuration())

	return nil
}
