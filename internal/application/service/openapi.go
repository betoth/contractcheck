package service

import (
	"context"

	"github.com/betoth/contractcheck/internal/application/customerrors"
	"github.com/betoth/contractcheck/internal/application/ports/input"
	"github.com/betoth/contractcheck/internal/application/ports/output"
	"github.com/betoth/contractcheck/internal/application/ports/output/openapi"
)

// OpenAPILoaderParams declares the hard dependencies required to build the service.
type OpenAPILoaderParams struct {
	Loader        openapi.Loader
	Logger        output.Logger
	VersionPolicy input.VersionPolicy
}

// validate performs defensive checks on constructor params.
func (p OpenAPILoaderParams) validate() error {
	if p.Loader == nil {
		return customerrors.NewDependencyError("loader")
	}
	if p.Logger == nil {
		return customerrors.NewDependencyError("logger")
	}
	if p.VersionPolicy == nil {
		return customerrors.NewDependencyError("versionPolicy")
	}
	return nil
}

// OpenAPILoaderService is the application service (input port implementation)
type OpenAPILoaderService struct {
	loader        openapi.Loader
	logger        output.Logger
	versionPolicy input.VersionPolicy
}

// NewOpenAPILoaderService constructs the service after validating dependencies.
func NewOpenAPILoaderService(params OpenAPILoaderParams) (*OpenAPILoaderService, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}
	return &OpenAPILoaderService{
		loader:        params.Loader,
		logger:        params.Logger,
		versionPolicy: params.VersionPolicy,
	}, nil
}

// Import loads an OpenAPI spec from disk through the output adapter, then enforces
// the configured VersionPolicy. It returns a canonical OpenAPIDoc or a typed error.
func (s *OpenAPILoaderService) Import(ctx context.Context, filePath string) (openapi.OpenAPIDoc, error) {
	log := s.logger.With("local", "service.OpenAPILoaderService.Import")

	log.Info("starting to load OpenAPI spec", "file", filePath)
	doc, err := s.loader.Load(ctx, filePath)
	if err != nil {
		log.Error("failed to load OpenAPI spec", "file", filePath)
		return openapi.OpenAPIDoc{}, err
	}

	if !s.versionPolicy.IsSupported(doc.Version.Major()) {
		log.Error(
			"invalid OpenAPI version",
			"file", filePath,
			"version", doc.Version.String(),
			"accepted", s.versionPolicy.SupportedVersions(),
		)
		return openapi.OpenAPIDoc{}, customerrors.NewUnsupportedVersionError(
			filePath,
			doc.Version.String(),
			s.versionPolicy.SupportedVersions(),
		)
	}

	log.Debug("successfully loaded OpenAPI spec")
	return doc, nil
}

// compile-time check
var _ input.ImportOpenAPISpec = (*OpenAPILoaderService)(nil)
