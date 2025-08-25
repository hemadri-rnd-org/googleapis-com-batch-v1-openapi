package main

import (
	"github.com/batch-api/mcp-server/config"
	"github.com/batch-api/mcp-server/models"
	tools_projects "github.com/batch-api/mcp-server/tools/projects"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateBatch_projects_locations_jobs_listTool(cfg),
		tools_projects.CreateBatch_projects_locations_jobs_createTool(cfg),
		tools_projects.CreateBatch_projects_locations_state_reportTool(cfg),
		tools_projects.CreateBatch_projects_locations_jobs_taskgroups_tasks_listTool(cfg),
		tools_projects.CreateBatch_projects_locations_operations_deleteTool(cfg),
		tools_projects.CreateBatch_projects_locations_operations_getTool(cfg),
		tools_projects.CreateBatch_projects_locations_listTool(cfg),
		tools_projects.CreateBatch_projects_locations_operations_listTool(cfg),
		tools_projects.CreateBatch_projects_locations_operations_cancelTool(cfg),
	}
}
