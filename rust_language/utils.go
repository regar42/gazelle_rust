package rust_language

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type logLevel int

const (
	logFatal logLevel = iota
	logErr
	logWarn
	logInfo
)

func (l *rustLang) Log(c *config.Config, level logLevel, from interface{}, msg string, args ...interface{}) {
	fmtMsg := fmt.Sprintf(msg, args...)

	var fromStr string
	switch f := from.(type) {
	case label.Label:
		fromStr = f.String()
	case string:
		fromStr = f
	case *rule.File:
		if f != nil {
			fromStr = f.Path
		} else {
			fromStr = ""
		}
	default:
		log.Panicf("unsupported from type: %v", from)
	}

	if fromStr != "" {
		fromStr = fmt.Sprintf("%s: ", fromStr)
	}

	if level == logFatal || (level != logInfo && c.Strict) {
		log.Fatalf("%s%s", fromStr, fmtMsg)
	} else {
		log.Printf("%s%s", fromStr, fmtMsg)
	}
}

func SliceContains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// ResolveIncludePaths resolves include_str!/include_bytes! paths relative to the source file.
// The srcFile is the path to the source file (e.g., "src/blocks/mod.rs"),
// and includePaths are the paths from include_str!() macros (e.g., "examples/doc.json").
// Returns the resolved paths relative to the package root (e.g., "src/blocks/examples/doc.json").
func ResolveIncludePaths(srcFile string, includePaths []string) []string {
	if len(includePaths) == 0 {
		return nil
	}

	srcDir := filepath.Dir(srcFile)
	resolved := make([]string, 0, len(includePaths))

	for _, includePath := range includePaths {
		// Resolve the path relative to the source file's directory
		resolvedPath := filepath.Join(srcDir, includePath)
		// Clean the path to handle ".." and "." properly
		resolvedPath = filepath.Clean(resolvedPath)
		resolved = append(resolved, resolvedPath)
	}

	return resolved
}

// DedupeAndSortStrings returns a deduplicated and sorted slice of strings.
func DedupeAndSortStrings(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))

	for _, s := range slice {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}

	sort.Strings(result)
	return result
}
