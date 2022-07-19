// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Defines values for ComposeStatusValue.
const (
	ComposeStatusValueFailure ComposeStatusValue = "failure"

	ComposeStatusValuePending ComposeStatusValue = "pending"

	ComposeStatusValueSuccess ComposeStatusValue = "success"
)

// Defines values for ImageStatusValue.
const (
	ImageStatusValueBuilding ImageStatusValue = "building"

	ImageStatusValueFailure ImageStatusValue = "failure"

	ImageStatusValuePending ImageStatusValue = "pending"

	ImageStatusValueRegistering ImageStatusValue = "registering"

	ImageStatusValueSuccess ImageStatusValue = "success"

	ImageStatusValueUploading ImageStatusValue = "uploading"
)

// Defines values for ImageTypes.
const (
	ImageTypesAws ImageTypes = "aws"

	ImageTypesAwsHaRhui ImageTypes = "aws-ha-rhui"

	ImageTypesAwsRhui ImageTypes = "aws-rhui"

	ImageTypesAwsSapRhui ImageTypes = "aws-sap-rhui"

	ImageTypesAzure ImageTypes = "azure"

	ImageTypesAzureRhui ImageTypes = "azure-rhui"

	ImageTypesEdgeCommit ImageTypes = "edge-commit"

	ImageTypesEdgeContainer ImageTypes = "edge-container"

	ImageTypesEdgeInstaller ImageTypes = "edge-installer"

	ImageTypesGcp ImageTypes = "gcp"

	ImageTypesGcpRhui ImageTypes = "gcp-rhui"

	ImageTypesGuestImage ImageTypes = "guest-image"

	ImageTypesImageInstaller ImageTypes = "image-installer"

	ImageTypesVsphere ImageTypes = "vsphere"
)

// Defines values for UploadStatusValue.
const (
	UploadStatusValueFailure UploadStatusValue = "failure"

	UploadStatusValuePending UploadStatusValue = "pending"

	UploadStatusValueRunning UploadStatusValue = "running"

	UploadStatusValueSuccess UploadStatusValue = "success"
)

// Defines values for UploadTypes.
const (
	UploadTypesAws UploadTypes = "aws"

	UploadTypesAwsS3 UploadTypes = "aws.s3"

	UploadTypesAzure UploadTypes = "azure"

	UploadTypesGcp UploadTypes = "gcp"
)

// AWSEC2UploadOptions defines model for AWSEC2UploadOptions.
type AWSEC2UploadOptions struct {
	Region            string   `json:"region"`
	ShareWithAccounts []string `json:"share_with_accounts"`
	SnapshotName      *string  `json:"snapshot_name,omitempty"`
}

// AWSEC2UploadStatus defines model for AWSEC2UploadStatus.
type AWSEC2UploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// AWSS3UploadOptions defines model for AWSS3UploadOptions.
type AWSS3UploadOptions struct {
	Region string `json:"region"`
}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AzureUploadOptions defines model for AzureUploadOptions.
type AzureUploadOptions struct {
	// Name of the uploaded image. It must be unique in the given resource group.
	// If name is omitted from the request, a random one based on a UUID is
	// generated.
	ImageName *string `json:"image_name,omitempty"`

	// Location where the image should be uploaded and registered.
	// How to list all locations:
	// https://docs.microsoft.com/en-us/cli/azure/account?view=azure-cli-latest#az_account_list_locations'
	Location string `json:"location"`

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded.
	// How to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeId defines model for ComposeId.
type ComposeId struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id string `json:"id"`
}

// ComposeLogs defines model for ComposeLogs.
type ComposeLogs struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	ImageBuilds []interface{} `json:"image_builds"`
	Koji        *KojiLogs     `json:"koji,omitempty"`
}

// ComposeManifests defines model for ComposeManifests.
type ComposeManifests struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Manifests []interface{} `json:"manifests"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequest   *ImageRequest   `json:"image_request,omitempty"`
	ImageRequests  *[]ImageRequest `json:"image_requests,omitempty"`
	Koji           *Koji           `json:"koji,omitempty"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	ImageStatus   ImageStatus        `json:"image_status"`
	ImageStatuses *[]ImageStatus     `json:"image_statuses,omitempty"`
	KojiStatus    *KojiStatus        `json:"koji_status,omitempty"`
	Status        ComposeStatusValue `json:"status"`
}

// ComposeStatusError defines model for ComposeStatusError.
type ComposeStatusError struct {
	Details *interface{} `json:"details,omitempty"`
	Id      int          `json:"id"`
	Reason  string       `json:"reason"`
}

// ComposeStatusValue defines model for ComposeStatusValue.
type ComposeStatusValue string

// Customizations defines model for Customizations.
type Customizations struct {
	Filesystem *[]Filesystem `json:"filesystem,omitempty"`
	Packages   *[]string     `json:"packages,omitempty"`

	// Extra repositories for packages specified in customizations. These
	// repositories will only be used to depsolve and retrieve packages
	// for the OS itself (they will not be available for the build root or
	// any other part of the build process). The package_sets field for these
	// repositories is ignored.
	PayloadRepositories *[]Repository `json:"payload_repositories,omitempty"`
	Services            *struct {
		// List of services to disable by default
		Disabled *[]string `json:"disabled,omitempty"`

		// List of services to enable by default
		Enabled *[]string `json:"enabled,omitempty"`
	} `json:"services,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	Users        *[]User       `json:"users,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Code        string       `json:"code"`
	Details     *interface{} `json:"details,omitempty"`
	OperationId string       `json:"operation_id"`
	Reason      string       `json:"reason"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	// Embedded struct due to allOf(#/components/schemas/List)
	List `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Items []Error `json:"items"`
}

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    uint64 `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadOptions defines model for GCPUploadOptions.
type GCPUploadOptions struct {
	// Name of an existing STANDARD Storage class Bucket.
	Bucket string `json:"bucket"`

	// The name to use for the imported and shared Compute Engine image.
	// The image name must be unique within the GCP project, which is used
	// for the OS image upload and import. If not specified a random
	// 'composer-api-<uuid>' string is used as the image name.
	ImageName *string `json:"image_name,omitempty"`

	// The GCP region where the OS image will be imported to and shared from.
	// The value must be a valid GCP location. See https://cloud.google.com/storage/docs/locations.
	// If not specified, the multi-region location closest to the source
	// (source Storage Bucket location) is chosen automatically.
	Region string `json:"region"`

	// List of valid Google accounts to share the imported Compute Engine image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	// If not specified, the imported Compute Engine image is not shared with any
	// account.
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture string       `json:"architecture"`
	ImageType    ImageTypes   `json:"image_type"`
	Ostree       *OSTree      `json:"ostree,omitempty"`
	Repositories []Repository `json:"repositories"`

	// This should really be oneOf but AWSS3UploadOptions is a subset of
	// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
	// are also valid AWSS3UploadOptionas objects which violates the oneOf
	// rules. Therefore, we have to use anyOf here but be aware that it isn't
	// possible to mix and match more schemas together.
	UploadOptions *UploadOptions `json:"upload_options,omitempty"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Error        *ComposeStatusError `json:"error,omitempty"`
	Status       ImageStatusValue    `json:"status"`
	UploadStatus *UploadStatus       `json:"upload_status,omitempty"`
}

// ImageStatusValue defines model for ImageStatusValue.
type ImageStatusValue string

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// Koji defines model for Koji.
type Koji struct {
	Name    string `json:"name"`
	Release string `json:"release"`
	Server  string `json:"server"`
	TaskId  int    `json:"task_id"`
	Version string `json:"version"`
}

// KojiLogs defines model for KojiLogs.
type KojiLogs struct {
	Import interface{} `json:"import"`
	Init   interface{} `json:"init"`
}

// KojiStatus defines model for KojiStatus.
type KojiStatus struct {
	BuildId *int `json:"build_id,omitempty"`
}

// List defines model for List.
type List struct {
	Kind  string `json:"kind"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
}

// OSTree defines model for OSTree.
type OSTree struct {
	// Can be either a commit (example: 02604b2da6e954bd34b8b82a835e5a77d2b60ffa), or a branch-like reference (example: rhel/8/x86_64/edge)
	Parent *string `json:"parent,omitempty"`
	Ref    *string `json:"ref,omitempty"`
	Url    *string `json:"url,omitempty"`
}

// ObjectReference defines model for ObjectReference.
type ObjectReference struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Kind string `json:"kind"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl  *string `json:"baseurl,omitempty"`
	CheckGpg *bool   `json:"check_gpg,omitempty"`

	// GPG key used to sign packages in this repository.
	Gpgkey     *string `json:"gpgkey,omitempty"`
	IgnoreSsl  *bool   `json:"ignore_ssl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`

	// Naming package sets for a repository assigns it to a specific part
	// (pipeline) of the build process.
	PackageSets *[]string `json:"package_sets,omitempty"`

	// Determines whether a valid subscription is required to access this repository.
	Rhsm *bool `json:"rhsm,omitempty"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation_key"`
	BaseUrl       string `json:"base_url"`
	Insights      bool   `json:"insights"`
	Organization  string `json:"organization"`
	ServerUrl     string `json:"server_url"`
}

// This should really be oneOf but AWSS3UploadOptions is a subset of
// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
// are also valid AWSS3UploadOptionas objects which violates the oneOf
// rules. Therefore, we have to use anyOf here but be aware that it isn't
// possible to mix and match more schemas together.
type UploadOptions interface{}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{}       `json:"options"`
	Status  UploadStatusValue `json:"status"`
	Type    UploadTypes       `json:"type"`
}

// UploadStatusValue defines model for UploadStatusValue.
type UploadStatusValue string

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// User defines model for User.
type User struct {
	Groups *[]string `json:"groups,omitempty"`
	Key    *string   `json:"key,omitempty"`
	Name   string    `json:"name"`
}

// Page defines model for page.
type Page string

// Size defines model for size.
type Size string

// PostComposeJSONBody defines parameters for PostCompose.
type PostComposeJSONBody ComposeRequest

// GetErrorListParams defines parameters for GetErrorList.
type GetErrorListParams struct {
	// Page index
	Page *Page `json:"page,omitempty"`

	// Number of items in each page
	Size *Size `json:"size,omitempty"`
}

// PostComposeJSONRequestBody defines body for PostCompose for application/json ContentType.
type PostComposeJSONRequestBody PostComposeJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create compose
	// (POST /compose)
	PostCompose(ctx echo.Context) error
	// The status of a compose
	// (GET /composes/{id})
	GetComposeStatus(ctx echo.Context, id string) error
	// Get logs for a compose.
	// (GET /composes/{id}/logs)
	GetComposeLogs(ctx echo.Context, id string) error
	// Get the manifests for a compose.
	// (GET /composes/{id}/manifests)
	GetComposeManifests(ctx echo.Context, id string) error
	// Get the metadata for a compose.
	// (GET /composes/{id}/metadata)
	GetComposeMetadata(ctx echo.Context, id string) error
	// Get a list of all possible errors
	// (GET /errors)
	GetErrorList(ctx echo.Context, params GetErrorListParams) error
	// Get error description
	// (GET /errors/{id})
	GetError(ctx echo.Context, id string) error
	// Get the openapi spec in json format
	// (GET /openapi)
	GetOpenapi(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostCompose converts echo context to params.
func (w *ServerInterfaceWrapper) PostCompose(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCompose(ctx)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, id)
	return err
}

// GetComposeLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeLogs(ctx, id)
	return err
}

// GetComposeManifests converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeManifests(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeManifests(ctx, id)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, id)
	return err
}

// GetErrorList converts echo context to params.
func (w *ServerInterfaceWrapper) GetErrorList(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetErrorListParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetErrorList(ctx, params)
	return err
}

// GetError converts echo context to params.
func (w *ServerInterfaceWrapper) GetError(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetError(ctx, id)
	return err
}

// GetOpenapi converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapi(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapi(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/compose", wrapper.PostCompose)
	router.GET(baseURL+"/composes/:id", wrapper.GetComposeStatus)
	router.GET(baseURL+"/composes/:id/logs", wrapper.GetComposeLogs)
	router.GET(baseURL+"/composes/:id/manifests", wrapper.GetComposeManifests)
	router.GET(baseURL+"/composes/:id/metadata", wrapper.GetComposeMetadata)
	router.GET(baseURL+"/errors", wrapper.GetErrorList)
	router.GET(baseURL+"/errors/:id", wrapper.GetError)
	router.GET(baseURL+"/openapi", wrapper.GetOpenapi)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
