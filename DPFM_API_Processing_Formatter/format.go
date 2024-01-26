package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-kanban-container-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		KanbanContainer: data.KanbanContainer,
	}
}
