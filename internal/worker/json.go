package worker

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/osbuild/osbuild-composer/internal/common"
	"github.com/osbuild/osbuild-composer/internal/distro"
	"github.com/osbuild/osbuild-composer/internal/osbuild"
	"github.com/osbuild/osbuild-composer/internal/target"
)

//
// JSON-serializable types for the jobqueue
//

type OSBuildJob struct {
	Manifest distro.Manifest  `json:"manifest"`
	Targets  []*target.Target `json:"targets,omitempty"`
}

type OSBuildJobResult struct {
	OSBuildOutput *osbuild.Result `json:"osbuild_output,omitempty"`
}

//
// JSON-serializable types for the HTTP API
//

type statusResponse struct {
	Status string `json:"status"`
}

type requestJobResponse struct {
	Id               uuid.UUID       `json:"id"`
	Location         string          `json:"location"`
	ArtifactLocation string          `json:"artifact_location"`
	Type             string          `json:"type"`
	Args             json.RawMessage `json:"args,omitempty"`
}

type getJobResponse struct {
	Canceled bool `json:"canceled"`
}

type updateJobRequest struct {
	Status common.ImageBuildState `json:"status"`
	Result *osbuild.Result        `json:"result"`
}

type updateJobResponse struct {
}
