// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ActionPayload struct {
	RequestType  string  `json:"request_type"`
	K8sManifest  string  `json:"k8s_manifest"`
	Namespace    string  `json:"namespace"`
	ExternalData *string `json:"external_data"`
}

type AgentStat struct {
	Ns      int `json:"Ns"`
	Cluster int `json:"Cluster"`
	Total   int `json:"Total"`
	Active  int `json:"Active"`
}

type Annotation struct {
	Categories       string `json:"Categories"`
	Vendor           string `json:"Vendor"`
	CreatedAt        string `json:"CreatedAt"`
	Repository       string `json:"Repository"`
	Support          string `json:"Support"`
	ChartDescription string `json:"ChartDescription"`
}

type ChaosWorkFlowInput struct {
	WorkflowID          *string            `json:"workflow_id"`
	WorkflowManifest    string             `json:"workflow_manifest"`
	CronSyntax          string             `json:"cronSyntax"`
	WorkflowName        string             `json:"workflow_name"`
	WorkflowDescription string             `json:"workflow_description"`
	Weightages          []*WeightagesInput `json:"weightages"`
	IsCustomWorkflow    bool               `json:"isCustomWorkflow"`
	ProjectID           string             `json:"project_id"`
	ClusterID           string             `json:"cluster_id"`
}

type ChaosWorkFlowResponse struct {
	WorkflowID          string `json:"workflow_id"`
	CronSyntax          string `json:"cronSyntax"`
	WorkflowName        string `json:"workflow_name"`
	WorkflowDescription string `json:"workflow_description"`
	IsCustomWorkflow    bool   `json:"isCustomWorkflow"`
}

type Chart struct {
	APIVersion  string              `json:"ApiVersion"`
	Kind        string              `json:"Kind"`
	Metadata    *Metadata           `json:"Metadata"`
	Spec        *Spec               `json:"Spec"`
	PackageInfo *PackageInformation `json:"PackageInfo"`
}

type Charts struct {
	Charts []*Chart `json:"Charts"`
}

type CloningInput struct {
	HubName       string   `json:"HubName"`
	ProjectID     string   `json:"ProjectID"`
	RepoBranch    string   `json:"RepoBranch"`
	RepoURL       string   `json:"RepoURL"`
	IsPrivate     bool     `json:"IsPrivate"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
}

type Cluster struct {
	ClusterID             string  `json:"cluster_id"`
	ProjectID             string  `json:"project_id"`
	ClusterName           string  `json:"cluster_name"`
	Description           *string `json:"description"`
	PlatformName          string  `json:"platform_name"`
	AccessKey             string  `json:"access_key"`
	IsRegistered          bool    `json:"is_registered"`
	IsClusterConfirmed    bool    `json:"is_cluster_confirmed"`
	IsActive              bool    `json:"is_active"`
	UpdatedAt             string  `json:"updated_at"`
	CreatedAt             string  `json:"created_at"`
	ClusterType           string  `json:"cluster_type"`
	NoOfSchedules         *int    `json:"no_of_schedules"`
	NoOfWorkflows         *int    `json:"no_of_workflows"`
	Token                 string  `json:"token"`
	AgentNamespace        *string `json:"agent_namespace"`
	Serviceaccount        *string `json:"serviceaccount"`
	AgentScope            string  `json:"agent_scope"`
	AgentNsExists         *bool   `json:"agent_ns_exists"`
	AgentSaExists         *bool   `json:"agent_sa_exists"`
	LastWorkflowTimestamp string  `json:"last_workflow_timestamp"`
}

type ClusterAction struct {
	ProjectID string         `json:"project_id"`
	Action    *ActionPayload `json:"action"`
}

type ClusterActionInput struct {
	ClusterID string `json:"cluster_id"`
	Action    string `json:"action"`
}

type ClusterConfirmResponse struct {
	IsClusterConfirmed bool    `json:"isClusterConfirmed"`
	NewAccessKey       *string `json:"newAccessKey"`
	ClusterID          *string `json:"cluster_id"`
}

type ClusterEvent struct {
	EventID     string   `json:"event_id"`
	EventType   string   `json:"event_type"`
	EventName   string   `json:"event_name"`
	Description string   `json:"description"`
	Cluster     *Cluster `json:"cluster"`
}

type ClusterEventInput struct {
	EventName   string `json:"event_name"`
	Description string `json:"description"`
	ClusterID   string `json:"cluster_id"`
	AccessKey   string `json:"access_key"`
}

type ClusterIdentity struct {
	ClusterID string `json:"cluster_id"`
	AccessKey string `json:"access_key"`
}

type ClusterInput struct {
	ClusterName    string  `json:"cluster_name"`
	Description    *string `json:"description"`
	PlatformName   string  `json:"platform_name"`
	ProjectID      string  `json:"project_id"`
	ClusterType    string  `json:"cluster_type"`
	AgentNamespace *string `json:"agent_namespace"`
	Serviceaccount *string `json:"serviceaccount"`
	AgentScope     string  `json:"agent_scope"`
	AgentNsExists  *bool   `json:"agent_ns_exists"`
	AgentSaExists  *bool   `json:"agent_sa_exists"`
	NodeSelector   *string `json:"node_selector"`
}

type CreateMyHub struct {
	HubName       string   `json:"HubName"`
	RepoURL       string   `json:"RepoURL"`
	RepoBranch    string   `json:"RepoBranch"`
	IsPrivate     bool     `json:"IsPrivate"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
	SSHPublicKey  *string  `json:"SSHPublicKey"`
}

type CreateUserInput struct {
	Username    string  `json:"username"`
	Email       *string `json:"email"`
	CompanyName *string `json:"company_name"`
	Name        *string `json:"name"`
	UserID      string  `json:"userID"`
	Role        string  `json:"role"`
}

type DSInput struct {
	DsID              *string `json:"ds_id"`
	DsName            string  `json:"ds_name"`
	DsType            string  `json:"ds_type"`
	DsURL             string  `json:"ds_url"`
	AccessType        string  `json:"access_type"`
	AuthType          string  `json:"auth_type"`
	BasicAuthUsername *string `json:"basic_auth_username"`
	BasicAuthPassword *string `json:"basic_auth_password"`
	ScrapeInterval    int     `json:"scrape_interval"`
	QueryTimeout      int     `json:"query_timeout"`
	HTTPMethod        string  `json:"http_method"`
	ProjectID         *string `json:"project_id"`
}

type DSResponse struct {
	DsID              *string `json:"ds_id"`
	DsName            *string `json:"ds_name"`
	DsType            *string `json:"ds_type"`
	DsURL             *string `json:"ds_url"`
	AccessType        *string `json:"access_type"`
	AuthType          *string `json:"auth_type"`
	BasicAuthUsername *string `json:"basic_auth_username"`
	BasicAuthPassword *string `json:"basic_auth_password"`
	ScrapeInterval    *int    `json:"scrape_interval"`
	QueryTimeout      *int    `json:"query_timeout"`
	HTTPMethod        *string `json:"http_method"`
	ProjectID         string  `json:"project_id"`
	HealthStatus      string  `json:"health_status"`
	CreatedAt         *string `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"`
}

type DateRange struct {
	StartDate string  `json:"start_date"`
	EndDate   *string `json:"end_date"`
}

type ExperimentInput struct {
	ProjectID      string  `json:"ProjectID"`
	ChartName      string  `json:"ChartName"`
	ExperimentName string  `json:"ExperimentName"`
	HubName        string  `json:"HubName"`
	FileType       *string `json:"FileType"`
}

type Experiments struct {
	Name string `json:"Name"`
	Csv  string `json:"CSV"`
	Desc string `json:"Desc"`
}

type GetWorkflowRunsInput struct {
	ProjectID      string                  `json:"project_id"`
	WorkflowRunIds []*string               `json:"workflow_run_ids"`
	WorkflowIds    []*string               `json:"workflow_ids"`
	Pagination     *Pagination             `json:"pagination"`
	Sort           *WorkflowRunSortInput   `json:"sort"`
	Filter         *WorkflowRunFilterInput `json:"filter"`
}

type GetWorkflowsOutput struct {
	TotalNoOfWorkflowRuns int            `json:"total_no_of_workflow_runs"`
	WorkflowRuns          []*WorkflowRun `json:"workflow_runs"`
}

type GitConfig struct {
	ProjectID     string   `json:"ProjectID"`
	Branch        string   `json:"Branch"`
	RepoURL       string   `json:"RepoURL"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
}

type GitConfigResponse struct {
	Enabled       bool      `json:"Enabled"`
	ProjectID     string    `json:"ProjectID"`
	Branch        *string   `json:"Branch"`
	RepoURL       *string   `json:"RepoURL"`
	AuthType      *AuthType `json:"AuthType"`
	Token         *string   `json:"Token"`
	UserName      *string   `json:"UserName"`
	Password      *string   `json:"Password"`
	SSHPrivateKey *string   `json:"SSHPrivateKey"`
}

type HeatmapData struct {
	Bins []*WorkflowRunsData `json:"bins"`
}

type ImageRegistryResponse struct {
	IsDefault         bool           `json:"is_default"`
	ImageRegistryInfo *ImageRegistry `json:"image_registry_info"`
	ImageRegistryID   string         `json:"image_registry_id"`
	ProjectID         string         `json:"project_id"`
	UpdatedAt         *string        `json:"updated_at"`
	CreatedAt         *string        `json:"created_at"`
	IsRemoved         *bool          `json:"is_removed"`
}

type KubeGVRRequest struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}

type KubeObjectData struct {
	RequestID string           `json:"request_id"`
	ClusterID *ClusterIdentity `json:"cluster_id"`
	KubeObj   string           `json:"kube_obj"`
}

type KubeObjectRequest struct {
	ClusterID      string          `json:"cluster_id"`
	ObjectType     string          `json:"object_type"`
	KubeObjRequest *KubeGVRRequest `json:"kube_obj_request"`
}

type KubeObjectResponse struct {
	ClusterID string `json:"cluster_id"`
	KubeObj   string `json:"kube_obj"`
}

type Link struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}

type ListWorkflowsInput struct {
	ProjectID   string               `json:"project_id"`
	WorkflowIds []*string            `json:"workflow_ids"`
	Pagination  *Pagination          `json:"pagination"`
	Sort        *WorkflowSortInput   `json:"sort"`
	Filter      *WorkflowFilterInput `json:"filter"`
}

type ListWorkflowsOutput struct {
	TotalNoOfWorkflows int         `json:"total_no_of_workflows"`
	Workflows          []*Workflow `json:"workflows"`
}

type Maintainer struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type ManifestTemplate struct {
	TemplateID          string `json:"template_id"`
	Manifest            string `json:"manifest"`
	TemplateName        string `json:"template_name"`
	TemplateDescription string `json:"template_description"`
	ProjectID           string `json:"project_id"`
	ProjectName         string `json:"project_name"`
	CreatedAt           string `json:"created_at"`
	IsRemoved           bool   `json:"is_removed"`
	IsCustomWorkflow    bool   `json:"isCustomWorkflow"`
}

type Member struct {
	UserID        string     `json:"user_id"`
	UserName      string     `json:"user_name"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Role          MemberRole `json:"role"`
	Invitation    string     `json:"invitation"`
	JoinedAt      string     `json:"joined_at"`
	DeactivatedAt string     `json:"deactivated_at"`
}

type MemberInput struct {
	ProjectID string      `json:"project_id"`
	UserID    string      `json:"user_id"`
	Role      *MemberRole `json:"role"`
}

type MemberStat struct {
	Owner *Owner `json:"Owner"`
	Total int    `json:"Total"`
}

type Metadata struct {
	Name        string      `json:"Name"`
	Version     string      `json:"Version"`
	Annotations *Annotation `json:"Annotations"`
}

type MyHub struct {
	ID            string   `json:"id"`
	RepoURL       string   `json:"RepoURL"`
	RepoBranch    string   `json:"RepoBranch"`
	ProjectID     string   `json:"ProjectID"`
	HubName       string   `json:"HubName"`
	IsPrivate     bool     `json:"IsPrivate"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
	IsRemoved     bool     `json:"IsRemoved"`
	CreatedAt     string   `json:"CreatedAt"`
	UpdatedAt     string   `json:"UpdatedAt"`
	LastSyncedAt  string   `json:"LastSyncedAt"`
}

type MyHubStatus struct {
	ID            string   `json:"id"`
	RepoURL       string   `json:"RepoURL"`
	RepoBranch    string   `json:"RepoBranch"`
	IsAvailable   bool     `json:"IsAvailable"`
	TotalExp      string   `json:"TotalExp"`
	HubName       string   `json:"HubName"`
	IsPrivate     bool     `json:"IsPrivate"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	IsRemoved     bool     `json:"IsRemoved"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
	SSHPublicKey  *string  `json:"SSHPublicKey"`
	LastSyncedAt  string   `json:"LastSyncedAt"`
}

type Owner struct {
	UserID   string `json:"UserId"`
	Username string `json:"Username"`
	Name     string `json:"Name"`
}

type PackageInformation struct {
	PackageName string         `json:"PackageName"`
	Experiments []*Experiments `json:"Experiments"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PodLog struct {
	ClusterID     *ClusterIdentity `json:"cluster_id"`
	RequestID     string           `json:"request_id"`
	WorkflowRunID string           `json:"workflow_run_id"`
	PodName       string           `json:"pod_name"`
	PodType       string           `json:"pod_type"`
	Log           string           `json:"log"`
}

type PodLogRequest struct {
	ClusterID      string  `json:"cluster_id"`
	WorkflowRunID  string  `json:"workflow_run_id"`
	PodName        string  `json:"pod_name"`
	PodNamespace   string  `json:"pod_namespace"`
	PodType        string  `json:"pod_type"`
	ExpPod         *string `json:"exp_pod"`
	RunnerPod      *string `json:"runner_pod"`
	ChaosNamespace *string `json:"chaos_namespace"`
}

type PodLogResponse struct {
	WorkflowRunID string `json:"workflow_run_id"`
	PodName       string `json:"pod_name"`
	PodType       string `json:"pod_type"`
	Log           string `json:"log"`
}

type PortalDashboardData struct {
	Name          string `json:"name"`
	DashboardData string `json:"dashboard_data"`
}

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Members   []*Member `json:"members"`
	State     *string   `json:"state"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	RemovedAt string    `json:"removed_at"`
}

type ProjectData struct {
	Name      string        `json:"Name"`
	Workflows *WorkflowStat `json:"Workflows"`
	Agents    *AgentStat    `json:"Agents"`
	ProjectID string        `json:"ProjectId"`
	Members   *MemberStat   `json:"Members"`
}

type Provider struct {
	Name string `json:"Name"`
}

type SSHKey struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type Spec struct {
	DisplayName         string        `json:"DisplayName"`
	CategoryDescription string        `json:"CategoryDescription"`
	Keywords            []string      `json:"Keywords"`
	Maturity            string        `json:"Maturity"`
	Maintainers         []*Maintainer `json:"Maintainers"`
	MinKubeVersion      string        `json:"MinKubeVersion"`
	Provider            string        `json:"Provider"`
	Links               []*Link       `json:"Links"`
	Experiments         []string      `json:"Experiments"`
	ChaosExpCRDLink     string        `json:"ChaosExpCRDLink"`
	Platforms           []string      `json:"Platforms"`
	ChaosType           *string       `json:"ChaosType"`
}

type TemplateInput struct {
	Manifest            string `json:"manifest"`
	TemplateName        string `json:"template_name"`
	TemplateDescription string `json:"template_description"`
	ProjectID           string `json:"project_id"`
	IsCustomWorkflow    bool   `json:"isCustomWorkflow"`
}

type TotalCount struct {
	Projects  int           `json:"Projects"`
	Users     int           `json:"Users"`
	Agents    *AgentStat    `json:"Agents"`
	Workflows *WorkflowStat `json:"Workflows"`
}

type UpdateMyHub struct {
	ID            string   `json:"id"`
	HubName       string   `json:"HubName"`
	RepoURL       string   `json:"RepoURL"`
	RepoBranch    string   `json:"RepoBranch"`
	IsPrivate     bool     `json:"IsPrivate"`
	AuthType      AuthType `json:"AuthType"`
	Token         *string  `json:"Token"`
	UserName      *string  `json:"UserName"`
	Password      *string  `json:"Password"`
	SSHPrivateKey *string  `json:"SSHPrivateKey"`
	SSHPublicKey  *string  `json:"SSHPublicKey"`
}

type UpdateUserInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	CompanyName *string `json:"company_name"`
}

type UsageData struct {
	Projects     []*ProjectData `json:"Projects"`
	TotalEntries int            `json:"TotalEntries"`
	TotalCount   *TotalCount    `json:"TotalCount"`
}

type UsageQuery struct {
	Pagination    *Pagination     `json:"Pagination"`
	DateRange     *DateRange      `json:"DateRange"`
	Sort          *UsageSortInput `json:"Sort"`
	SearchProject *string         `json:"SearchProject"`
}

type UsageSortInput struct {
	Field      UsageSort `json:"Field"`
	Descending bool      `json:"Descending"`
}

type User struct {
	ID              string     `json:"id"`
	Username        string     `json:"username"`
	Email           *string    `json:"email"`
	IsEmailVerified *bool      `json:"is_email_verified"`
	CompanyName     *string    `json:"company_name"`
	Name            *string    `json:"name"`
	Projects        []*Project `json:"projects"`
	Role            *string    `json:"role"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	DeactivatedAt   string     `json:"deactivated_at"`
}

type WeightagesInput struct {
	ExperimentName string `json:"experiment_name"`
	Weightage      int    `json:"weightage"`
}

type Workflow struct {
	WorkflowID          string        `json:"workflow_id"`
	WorkflowManifest    string        `json:"workflow_manifest"`
	CronSyntax          string        `json:"cronSyntax"`
	ClusterName         string        `json:"cluster_name"`
	WorkflowName        string        `json:"workflow_name"`
	WorkflowDescription string        `json:"workflow_description"`
	Weightages          []*Weightages `json:"weightages"`
	IsCustomWorkflow    bool          `json:"isCustomWorkflow"`
	UpdatedAt           string        `json:"updated_at"`
	CreatedAt           string        `json:"created_at"`
	ProjectID           string        `json:"project_id"`
	ClusterID           string        `json:"cluster_id"`
	ClusterType         string        `json:"cluster_type"`
	IsRemoved           bool          `json:"isRemoved"`
}

type WorkflowFilterInput struct {
	WorkflowName *string `json:"workflow_name"`
	ClusterName  *string `json:"cluster_name"`
}

type WorkflowRun struct {
	WorkflowRunID      string        `json:"workflow_run_id"`
	WorkflowID         string        `json:"workflow_id"`
	ClusterName        string        `json:"cluster_name"`
	Weightages         []*Weightages `json:"weightages"`
	LastUpdated        string        `json:"last_updated"`
	ProjectID          string        `json:"project_id"`
	ClusterID          string        `json:"cluster_id"`
	WorkflowName       string        `json:"workflow_name"`
	ClusterType        *string       `json:"cluster_type"`
	Phase              string        `json:"phase"`
	ResiliencyScore    *float64      `json:"resiliency_score"`
	ExperimentsPassed  *int          `json:"experiments_passed"`
	ExperimentsFailed  *int          `json:"experiments_failed"`
	ExperimentsAwaited *int          `json:"experiments_awaited"`
	ExperimentsStopped *int          `json:"experiments_stopped"`
	ExperimentsNa      *int          `json:"experiments_na"`
	TotalExperiments   *int          `json:"total_experiments"`
	ExecutionData      string        `json:"execution_data"`
	IsRemoved          *bool         `json:"isRemoved"`
}

type WorkflowRunDetails struct {
	NoOfRuns  int     `json:"no_of_runs"`
	DateStamp float64 `json:"date_stamp"`
}

type WorkflowRunFilterInput struct {
	WorkflowName   *string            `json:"workflow_name"`
	ClusterName    *string            `json:"cluster_name"`
	WorkflowStatus *WorkflowRunStatus `json:"workflow_status"`
	DateRange      *DateRange         `json:"date_range"`
}

type WorkflowRunInput struct {
	WorkflowID    string           `json:"workflow_id"`
	WorkflowRunID string           `json:"workflow_run_id"`
	WorkflowName  string           `json:"workflow_name"`
	ExecutionData string           `json:"execution_data"`
	ClusterID     *ClusterIdentity `json:"cluster_id"`
	Completed     bool             `json:"completed"`
	IsRemoved     *bool            `json:"isRemoved"`
}

type WorkflowRunSortInput struct {
	Field      WorkflowSortingField `json:"field"`
	Descending *bool                `json:"descending"`
}

type WorkflowRunStatsRequest struct {
	ProjectID   string    `json:"project_id"`
	WorkflowIds []*string `json:"workflow_ids"`
}

type WorkflowRunStatsResponse struct {
	TotalWorkflowRuns              int     `json:"total_workflow_runs"`
	SucceededWorkflowRuns          int     `json:"succeeded_workflow_runs"`
	FailedWorkflowRuns             int     `json:"failed_workflow_runs"`
	RunningWorkflowRuns            int     `json:"running_workflow_runs"`
	AverageResiliencyScore         float64 `json:"average_resiliency_score"`
	TotalExperiments               int     `json:"total_experiments"`
	ExperimentsPassed              int     `json:"experiments_passed"`
	ExperimentsFailed              int     `json:"experiments_failed"`
	ExperimentsAwaited             int     `json:"experiments_awaited"`
	ExperimentsStopped             int     `json:"experiments_stopped"`
	ExperimentsNa                  int     `json:"experiments_na"`
	PassedPercentage               float64 `json:"passed_percentage"`
	FailedPercentage               float64 `json:"failed_percentage"`
	WorkflowRunSucceededPercentage float64 `json:"workflow_run_succeeded_percentage"`
	WorkflowRunFailedPercentage    float64 `json:"workflow_run_failed_percentage"`
}

type WorkflowRunsData struct {
	Value             *float64            `json:"value"`
	WorkflowRunDetail *WorkflowRunDetails `json:"workflowRunDetail"`
}

type WorkflowSortInput struct {
	Field      WorkflowSortingField `json:"field"`
	Descending *bool                `json:"descending"`
}

type WorkflowStat struct {
	Schedules int `json:"Schedules"`
	Runs      int `json:"Runs"`
	ExpRuns   int `json:"ExpRuns"`
}

type WorkflowStats struct {
	Date  float64 `json:"date"`
	Value int     `json:"value"`
}

type AnnotationsPromResponse struct {
	Queryid      string                         `json:"queryid"`
	Legends      []*string                      `json:"legends"`
	Tsvs         [][]*AnnotationsTimeStampValue `json:"tsvs"`
	SubDataArray [][]*SubData                   `json:"subDataArray"`
}

type AnnotationsTimeStampValue struct {
	Date  *float64 `json:"date"`
	Value *int     `json:"value"`
}

type ApplicationMetadata struct {
	Namespace    string      `json:"namespace"`
	Applications []*Resource `json:"applications"`
}

type ApplicationMetadataResponse struct {
	Namespace    string              `json:"namespace"`
	Applications []*ResourceResponse `json:"applications"`
}

type ClusterRegResponse struct {
	Token       string `json:"token"`
	ClusterID   string `json:"cluster_id"`
	ClusterName string `json:"cluster_name"`
}

type CreateDBInput struct {
	DsID                      string                 `json:"ds_id"`
	DbName                    string                 `json:"db_name"`
	DbTypeName                string                 `json:"db_type_name"`
	DbTypeID                  string                 `json:"db_type_id"`
	DbInformation             *string                `json:"db_information"`
	ChaosEventQueryTemplate   string                 `json:"chaos_event_query_template"`
	ChaosVerdictQueryTemplate string                 `json:"chaos_verdict_query_template"`
	ApplicationMetadataMap    []*ApplicationMetadata `json:"application_metadata_map"`
	PanelGroups               []*PanelGroup          `json:"panel_groups"`
	EndTime                   string                 `json:"end_time"`
	StartTime                 string                 `json:"start_time"`
	ProjectID                 string                 `json:"project_id"`
	ClusterID                 string                 `json:"cluster_id"`
	RefreshRate               string                 `json:"refresh_rate"`
}

type DashboardPromResponse struct {
	DashboardMetricsResponse []*MetricDataForPanelGroup `json:"dashboardMetricsResponse"`
	AnnotationsResponse      []*AnnotationsPromResponse `json:"annotationsResponse"`
}

type DataVars struct {
	URL             string `json:"url"`
	Start           string `json:"start"`
	End             string `json:"end"`
	RelativeTime    int    `json:"relative_time"`
	RefreshInterval int    `json:"refresh_interval"`
}

type DeleteDSInput struct {
	ForceDelete bool   `json:"force_delete"`
	DsID        string `json:"ds_id"`
}

type DsDetails struct {
	URL   string `json:"url"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type ImageRegistry struct {
	IsDefault         *bool   `json:"is_default"`
	ImageRegistryName string  `json:"image_registry_name"`
	ImageRepoName     string  `json:"image_repo_name"`
	ImageRegistryType string  `json:"image_registry_type"`
	SecretName        *string `json:"secret_name"`
	SecretNamespace   *string `json:"secret_namespace"`
	EnableRegistry    *bool   `json:"enable_registry"`
}

type ImageRegistryInput struct {
	IsDefault         bool    `json:"is_default"`
	ImageRegistryName string  `json:"image_registry_name"`
	ImageRepoName     string  `json:"image_repo_name"`
	ImageRegistryType string  `json:"image_registry_type"`
	SecretName        *string `json:"secret_name"`
	SecretNamespace   *string `json:"secret_namespace"`
	EnableRegistry    *bool   `json:"enable_registry"`
}

type LabelValue struct {
	Label  string    `json:"label"`
	Values []*Option `json:"values"`
}

type ListDashboardResponse struct {
	DsID                      string                         `json:"ds_id"`
	DbID                      string                         `json:"db_id"`
	DbName                    string                         `json:"db_name"`
	DbTypeID                  string                         `json:"db_type_id"`
	DbTypeName                string                         `json:"db_type_name"`
	DbInformation             *string                        `json:"db_information"`
	ChaosEventQueryTemplate   string                         `json:"chaos_event_query_template"`
	ChaosVerdictQueryTemplate string                         `json:"chaos_verdict_query_template"`
	ApplicationMetadataMap    []*ApplicationMetadataResponse `json:"application_metadata_map"`
	ClusterName               *string                        `json:"cluster_name"`
	DsName                    *string                        `json:"ds_name"`
	DsType                    *string                        `json:"ds_type"`
	DsURL                     *string                        `json:"ds_url"`
	DsHealthStatus            *string                        `json:"ds_health_status"`
	PanelGroups               []*PanelGroupResponse          `json:"panel_groups"`
	EndTime                   string                         `json:"end_time"`
	StartTime                 string                         `json:"start_time"`
	RefreshRate               string                         `json:"refresh_rate"`
	ProjectID                 string                         `json:"project_id"`
	ClusterID                 string                         `json:"cluster_id"`
	CreatedAt                 *string                        `json:"created_at"`
	UpdatedAt                 *string                        `json:"updated_at"`
	ViewedAt                  *string                        `json:"viewed_at"`
}

type MetricDataForPanel struct {
	PanelID              string                 `json:"panelID"`
	PanelMetricsResponse []*MetricsPromResponse `json:"PanelMetricsResponse"`
}

type MetricDataForPanelGroup struct {
	PanelGroupID              string                `json:"panelGroupID"`
	PanelGroupMetricsResponse []*MetricDataForPanel `json:"panelGroupMetricsResponse"`
}

type MetricsPromResponse struct {
	Queryid string                     `json:"queryid"`
	Legends []*string                  `json:"legends"`
	Tsvs    [][]*MetricsTimeStampValue `json:"tsvs"`
}

type MetricsTimeStampValue struct {
	Date  *float64 `json:"date"`
	Value *float64 `json:"value"`
}

type Option struct {
	Name string `json:"name"`
}

type Panel struct {
	PanelID      *string      `json:"panel_id"`
	DbID         *string      `json:"db_id"`
	YAxisLeft    *string      `json:"y_axis_left"`
	YAxisRight   *string      `json:"y_axis_right"`
	XAxisDown    *string      `json:"x_axis_down"`
	Unit         *string      `json:"unit"`
	PanelGroupID *string      `json:"panel_group_id"`
	CreatedAt    *string      `json:"created_at"`
	PromQueries  []*PromQuery `json:"prom_queries"`
	PanelOptions *PanelOption `json:"panel_options"`
	PanelName    string       `json:"panel_name"`
}

type PanelGroup struct {
	Panels         []*Panel `json:"panels"`
	PanelGroupName string   `json:"panel_group_name"`
}

type PanelGroupResponse struct {
	Panels         []*PanelResponse `json:"panels"`
	PanelGroupName string           `json:"panel_group_name"`
	PanelGroupID   *string          `json:"panel_group_id"`
}

type PanelOption struct {
	Points   *bool `json:"points"`
	Grids    *bool `json:"grids"`
	LeftAxis *bool `json:"left_axis"`
}

type PanelOptionResponse struct {
	Points   *bool `json:"points"`
	Grids    *bool `json:"grids"`
	LeftAxis *bool `json:"left_axis"`
}

type PanelResponse struct {
	PanelID      string               `json:"panel_id"`
	YAxisLeft    *string              `json:"y_axis_left"`
	YAxisRight   *string              `json:"y_axis_right"`
	XAxisDown    *string              `json:"x_axis_down"`
	Unit         *string              `json:"unit"`
	PromQueries  []*PromQueryResponse `json:"prom_queries"`
	PanelOptions *PanelOptionResponse `json:"panel_options"`
	PanelName    *string              `json:"panel_name"`
	CreatedAt    *string              `json:"created_at"`
}

type PromInput struct {
	Queries   []*PromQueryInput `json:"queries"`
	DsDetails *DsDetails        `json:"ds_details"`
}

type PromQuery struct {
	Queryid       string  `json:"queryid"`
	PromQueryName *string `json:"prom_query_name"`
	Legend        *string `json:"legend"`
	Resolution    *string `json:"resolution"`
	Minstep       *string `json:"minstep"`
	Line          *bool   `json:"line"`
	CloseArea     *bool   `json:"close_area"`
}

type PromQueryInput struct {
	Queryid    string  `json:"queryid"`
	Query      string  `json:"query"`
	Legend     *string `json:"legend"`
	Resolution *string `json:"resolution"`
	Minstep    int     `json:"minstep"`
}

type PromQueryResponse struct {
	Queryid       string  `json:"queryid"`
	PromQueryName *string `json:"prom_query_name"`
	Legend        *string `json:"legend"`
	Resolution    *string `json:"resolution"`
	Minstep       *string `json:"minstep"`
	Line          *bool   `json:"line"`
	CloseArea     *bool   `json:"close_area"`
}

type PromResponse struct {
	MetricsResponse     []*MetricsPromResponse     `json:"metricsResponse"`
	AnnotationsResponse []*AnnotationsPromResponse `json:"annotationsResponse"`
}

type PromSeriesInput struct {
	Series    string     `json:"series"`
	DsDetails *DsDetails `json:"ds_details"`
}

type PromSeriesListResponse struct {
	SeriesList []*string `json:"seriesList"`
}

type PromSeriesResponse struct {
	Series      string        `json:"series"`
	LabelValues []*LabelValue `json:"labelValues"`
}

type QueryMapForPanel struct {
	PanelID  string   `json:"panelID"`
	QueryIDs []string `json:"queryIDs"`
}

type QueryMapForPanelGroup struct {
	PanelGroupID  string              `json:"panelGroupID"`
	PanelQueryMap []*QueryMapForPanel `json:"panelQueryMap"`
}

type Resource struct {
	Kind  string    `json:"kind"`
	Names []*string `json:"names"`
}

type ResourceResponse struct {
	Kind  string    `json:"kind"`
	Names []*string `json:"names"`
}

type SubData struct {
	Date        *float64 `json:"date"`
	Value       string   `json:"value"`
	SubDataName string   `json:"subDataName"`
}

type UpdateDBInput struct {
	DbID                      string                   `json:"db_id"`
	DsID                      *string                  `json:"ds_id"`
	DbName                    *string                  `json:"db_name"`
	DbTypeName                *string                  `json:"db_type_name"`
	DbTypeID                  *string                  `json:"db_type_id"`
	DbInformation             *string                  `json:"db_information"`
	ChaosEventQueryTemplate   *string                  `json:"chaos_event_query_template"`
	ChaosVerdictQueryTemplate *string                  `json:"chaos_verdict_query_template"`
	ApplicationMetadataMap    []*ApplicationMetadata   `json:"application_metadata_map"`
	PanelGroups               []*UpdatePanelGroupInput `json:"panel_groups"`
	EndTime                   *string                  `json:"end_time"`
	StartTime                 *string                  `json:"start_time"`
	ClusterID                 *string                  `json:"cluster_id"`
	RefreshRate               *string                  `json:"refresh_rate"`
}

type UpdatePanelGroupInput struct {
	PanelGroupName string   `json:"panel_group_name"`
	PanelGroupID   string   `json:"panel_group_id"`
	Panels         []*Panel `json:"panels"`
}

type Weightages struct {
	ExperimentName string `json:"experiment_name"`
	Weightage      int    `json:"weightage"`
}

type AuthType string

const (
	AuthTypeNone  AuthType = "none"
	AuthTypeBasic AuthType = "basic"
	AuthTypeToken AuthType = "token"
	AuthTypeSSH   AuthType = "ssh"
)

var AllAuthType = []AuthType{
	AuthTypeNone,
	AuthTypeBasic,
	AuthTypeToken,
	AuthTypeSSH,
}

func (e AuthType) IsValid() bool {
	switch e {
	case AuthTypeNone, AuthTypeBasic, AuthTypeToken, AuthTypeSSH:
		return true
	}
	return false
}

func (e AuthType) String() string {
	return string(e)
}

func (e *AuthType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthType", str)
	}
	return nil
}

func (e AuthType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MemberRole string

const (
	MemberRoleOwner  MemberRole = "Owner"
	MemberRoleEditor MemberRole = "Editor"
	MemberRoleViewer MemberRole = "Viewer"
)

var AllMemberRole = []MemberRole{
	MemberRoleOwner,
	MemberRoleEditor,
	MemberRoleViewer,
}

func (e MemberRole) IsValid() bool {
	switch e {
	case MemberRoleOwner, MemberRoleEditor, MemberRoleViewer:
		return true
	}
	return false
}

func (e MemberRole) String() string {
	return string(e)
}

func (e *MemberRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MemberRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MemberRole", str)
	}
	return nil
}

func (e MemberRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TimeFrequency string

const (
	TimeFrequencyMonthly TimeFrequency = "Monthly"
	TimeFrequencyDaily   TimeFrequency = "Daily"
	TimeFrequencyHourly  TimeFrequency = "Hourly"
)

var AllTimeFrequency = []TimeFrequency{
	TimeFrequencyMonthly,
	TimeFrequencyDaily,
	TimeFrequencyHourly,
}

func (e TimeFrequency) IsValid() bool {
	switch e {
	case TimeFrequencyMonthly, TimeFrequencyDaily, TimeFrequencyHourly:
		return true
	}
	return false
}

func (e TimeFrequency) String() string {
	return string(e)
}

func (e *TimeFrequency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimeFrequency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TimeFrequency", str)
	}
	return nil
}

func (e TimeFrequency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UsageSort string

const (
	UsageSortProject        UsageSort = "Project"
	UsageSortOwner          UsageSort = "Owner"
	UsageSortAgents         UsageSort = "Agents"
	UsageSortSchedules      UsageSort = "Schedules"
	UsageSortWorkflowRuns   UsageSort = "WorkflowRuns"
	UsageSortExperimentRuns UsageSort = "ExperimentRuns"
	UsageSortTeamMembers    UsageSort = "TeamMembers"
)

var AllUsageSort = []UsageSort{
	UsageSortProject,
	UsageSortOwner,
	UsageSortAgents,
	UsageSortSchedules,
	UsageSortWorkflowRuns,
	UsageSortExperimentRuns,
	UsageSortTeamMembers,
}

func (e UsageSort) IsValid() bool {
	switch e {
	case UsageSortProject, UsageSortOwner, UsageSortAgents, UsageSortSchedules, UsageSortWorkflowRuns, UsageSortExperimentRuns, UsageSortTeamMembers:
		return true
	}
	return false
}

func (e UsageSort) String() string {
	return string(e)
}

func (e *UsageSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UsageSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UsageSort", str)
	}
	return nil
}

func (e UsageSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WorkflowRunStatus string

const (
	WorkflowRunStatusAll       WorkflowRunStatus = "All"
	WorkflowRunStatusFailed    WorkflowRunStatus = "Failed"
	WorkflowRunStatusRunning   WorkflowRunStatus = "Running"
	WorkflowRunStatusSucceeded WorkflowRunStatus = "Succeeded"
)

var AllWorkflowRunStatus = []WorkflowRunStatus{
	WorkflowRunStatusAll,
	WorkflowRunStatusFailed,
	WorkflowRunStatusRunning,
	WorkflowRunStatusSucceeded,
}

func (e WorkflowRunStatus) IsValid() bool {
	switch e {
	case WorkflowRunStatusAll, WorkflowRunStatusFailed, WorkflowRunStatusRunning, WorkflowRunStatusSucceeded:
		return true
	}
	return false
}

func (e WorkflowRunStatus) String() string {
	return string(e)
}

func (e *WorkflowRunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkflowRunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkflowRunStatus", str)
	}
	return nil
}

func (e WorkflowRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WorkflowSortingField string

const (
	WorkflowSortingFieldName WorkflowSortingField = "Name"
	WorkflowSortingFieldTime WorkflowSortingField = "Time"
)

var AllWorkflowSortingField = []WorkflowSortingField{
	WorkflowSortingFieldName,
	WorkflowSortingFieldTime,
}

func (e WorkflowSortingField) IsValid() bool {
	switch e {
	case WorkflowSortingFieldName, WorkflowSortingFieldTime:
		return true
	}
	return false
}

func (e WorkflowSortingField) String() string {
	return string(e)
}

func (e *WorkflowSortingField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkflowSortingField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkflowSortingField", str)
	}
	return nil
}

func (e WorkflowSortingField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
