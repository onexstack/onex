package version

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Call the Get function to retrieve version info
	info := Get()

	// Assert that the returned Info struct contains expected values
	assert.Equal(t, gitVersion, info.GitVersion)
	assert.Equal(t, gitCommit, info.GitCommit)
	assert.Equal(t, gitTreeState, info.GitTreeState)
	assert.Equal(t, buildDate, info.BuildDate)
	assert.Equal(t, runtime.Version(), info.GoVersion)
	assert.Equal(t, runtime.Compiler, info.Compiler)
	assert.Equal(t, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), info.Platform)
}

func TestInfo_String(t *testing.T) {
	info := Info{GitVersion: "v1.0.0"}

	// Test the String method
	assert.Equal(t, "v1.0.0", info.String())
}

func TestInfo_ToJSON(t *testing.T) {
	info := Info{
		GitVersion:   "v1.0.0",
		GitCommit:    "abc123",
		GitTreeState: "clean",
		BuildDate:    "2024-01-01T00:00:00Z",
		GoVersion:    "go1.20",
		Compiler:     "gc",
		Platform:     "linux/amd64",
	}

	// Expected JSON output
	expectedJSON := `{"gitVersion":"v1.0.0","gitCommit":"abc123","gitTreeState":"clean","buildDate":"2024-01-01T00:00:00Z","goVersion":"go1.20","compiler":"gc","platform":"linux/amd64"}`

	// Test the ToJSON method
	assert.JSONEq(t, expectedJSON, info.ToJSON())
}

func TestInfo_Text(t *testing.T) {
	info := Info{
		GitVersion:   "v1.0.0",
		GitCommit:    "abc123",
		GitTreeState: "clean",
		BuildDate:    "2024-01-01T00:00:00Z",
		GoVersion:    "go1.20",
		Compiler:     "gc",
		Platform:     "linux/amd64",
	}

	// Test the Text method
	stringOutput := info.Text()
	assert.Contains(t, stringOutput, "gitVersion: "+info.GitVersion)
	assert.Contains(t, stringOutput, "gitCommit: "+info.GitCommit)
	assert.Contains(t, stringOutput, "gitTreeState: "+info.GitTreeState)
	assert.Contains(t, stringOutput, "buildDate: "+info.BuildDate)
	assert.Contains(t, stringOutput, "goVersion: "+info.GoVersion)
	assert.Contains(t, stringOutput, "compiler: "+info.Compiler)
	assert.Contains(t, stringOutput, "platform: "+info.Platform)
}
