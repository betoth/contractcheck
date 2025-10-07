package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/betoth/contractcheck/internal/application/ports/input"
)

// OpenAPIVersionPolicy enforces acceptance based on *major* version only
type OpenAPIVersionPolicy struct {
	supportedVersions map[int]struct{}
}

// NewOpenAPIVersionPolicy builds a policy from a list of majors.
// Duplicates and non-positive values are ignored; result is canonicalized.
func NewOpenAPIVersionPolicy(versions []int) *OpenAPIVersionPolicy {
	norm := normalizeMajors(versions)
	supported := make(map[int]struct{}, len(norm))
	for _, v := range norm {
		supported[v] = struct{}{}
	}
	return &OpenAPIVersionPolicy{supportedVersions: supported}
}

// IsSupported reports whether a major version is allowed.
func (p *OpenAPIVersionPolicy) IsSupported(major int) bool {
	_, ok := p.supportedVersions[major]
	return ok
}

// SupportedVersions returns the allowed majors in ascending order.
// Suitable for logs, telemetry and error messages.
func (p *OpenAPIVersionPolicy) SupportedVersions() []int {
	versions := make([]int, 0, len(p.supportedVersions))
	for v := range p.supportedVersions {
		versions = append(versions, v)
	}
	sort.Ints(versions)
	return versions
}

// FormatVersions renders majors as a human-friendly list, e.g. "3.x, 4.x".
func (p *OpenAPIVersionPolicy) FormatVersions() string {
	versions := p.SupportedVersions()
	if len(versions) == 0 {
		return ""
	}
	parts := make([]string, len(versions))
	for i, m := range versions {
		parts[i] = fmt.Sprintf("%d.x", m)
	}
	return strings.Join(parts, ", ")
}

// normalizeMajors canonicalizes input majors:
// - removes duplicates
// - filters non-positive values
// - returns ascending order
func normalizeMajors(in []int) []int {
	seen := make(map[int]struct{}, len(in))
	out := make([]int, 0, len(in))
	for _, v := range in {
		if v <= 0 {
			continue
		}
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}

// Add this to ensure APIVersionPolicy implements VersionPolicy interface
var _ input.VersionPolicy = (*OpenAPIVersionPolicy)(nil)
