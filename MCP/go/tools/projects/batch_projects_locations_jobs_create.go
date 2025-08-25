package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/batch-api/mcp-server/config"
	"github.com/batch-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Batch_projects_locations_jobs_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["jobId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("jobId=%v", val))
		}
		if val, ok := args["requestId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("requestId=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Job
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/jobs%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Job
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateBatch_projects_locations_jobs_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_jobs",
		mcp.WithDescription("Create a Job."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The parent resource name where the Job will be created. Pattern: \"projects/{project}/locations/{location}\"")),
		mcp.WithString("jobId", mcp.Description("ID used to uniquely identify the Job within its parent scope. This field should contain at most 63 characters and must start with lowercase characters. Only lowercase characters, numbers and '-' are accepted. The '-' character cannot be the first or the last one. A system generated ID will be used if the field is not set. The job.name field in the request will be ignored and the created resource name of the Job will be \"{parent}/jobs/{job_id}\".")),
		mcp.WithString("requestId", mcp.Description("Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).")),
		mcp.WithArray("taskGroups", mcp.Description("Input parameter: Required. TaskGroups in the Job. Only one TaskGroup is supported now.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Labels for the Job. Labels could be user provided or system generated. For example, \"labels\": { \"department\": \"finance\", \"environment\": \"test\" } You can assign up to 64 labels. [Google Compute Engine label restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) apply. Label names that start with \"goog-\" or \"google-\" are reserved.")),
		mcp.WithArray("notifications", mcp.Description("Input parameter: Notification configurations.")),
		mcp.WithString("uid", mcp.Description("Input parameter: Output only. A system generated unique ID for the Job.")),
		mcp.WithObject("status", mcp.Description("Input parameter: Job status.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The last time the Job was updated.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. When the Job was created.")),
		mcp.WithString("name", mcp.Description("Input parameter: Output only. Job name. For example: \"projects/123456/locations/us-central1/jobs/job01\".")),
		mcp.WithObject("allocationPolicy", mcp.Description("Input parameter: A Job's resource allocation policy describes when, where, and how compute resources should be allocated for the Job.")),
		mcp.WithObject("logsPolicy", mcp.Description("Input parameter: LogsPolicy describes how outputs from a Job's Tasks (stdout/stderr) will be preserved.")),
		mcp.WithString("priority", mcp.Description("Input parameter: Priority of the Job. The valid value range is [0, 100). Default value is 0. Higher value indicates higher priority. A job with higher priority value is more likely to run earlier if all other requirements are satisfied.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Batch_projects_locations_jobs_createHandler(cfg),
	}
}
