// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

// This file contains e2e tests for the consumer + stats generator based on events
// gatherered with the event collector. They are placed in a separate file as they don't match
// with any specific struct but are rather integration tests for the consumer and stats generator.

//go:build linux_bpf && nvml

package gpu

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DataDog/datadog-agent/pkg/gpu/config"
	"github.com/DataDog/datadog-agent/pkg/gpu/ebpf"
	ddnvml "github.com/DataDog/datadog-agent/pkg/gpu/safenvml"
	nvmltestutil "github.com/DataDog/datadog-agent/pkg/gpu/safenvml/testutil"
	"github.com/DataDog/datadog-agent/pkg/gpu/testutil"
	"github.com/DataDog/datadog-agent/pkg/util/kernel"
)

\
	var stream *StreamHandler
	for _, s := range handlers.allStreams() {
		require.Nil(t, stream) // There should be only one stream, so it should be set to nil at this point
		stream = s
	}
return package null
