package mapper

import (
	"kutamukti-api/api/schedule/dto"
	"kutamukti-api/models"
	"kutamukti-api/pkg/logger"

	"github.com/go-viper/mapstructure/v2"
)

func MapScheduleInputToModel(input dto.Schedule) models.Schedules {
	var output models.Schedules

	mapstructure.Decode(input, &output)
	output.StartTime = input.StartTime
	output.EndTime = input.EndTime

	logger.Info("input: %v", input.StartTime)
	logger.Info("output: %v", output.StartTime)
	return output
}

func MapScheduleModelToOutput(input models.Schedules) dto.ScheduleResponse {
	var output dto.ScheduleResponse

	mapstructure.Decode(input, &output)
	output.StartTime = input.StartTime
	output.EndTime = input.EndTime
	output.CreatedAt = input.CreatedAt
	output.UpdatedAt = input.UpdatedAt

	return output
}
