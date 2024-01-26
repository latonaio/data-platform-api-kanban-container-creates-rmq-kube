package requests

type Header struct {
	KanbanContainer     int      `json:"KanbanContainer"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	StorageLocation     string   `json:"StorageLocation"`
	StorageBin          string   `json:"StorageBin"`
	KanbanControlCycle  string   `json:"KanbanControlCycle"`
	Product             string   `json:"Product"`
	Capacity            *float32 `json:"Capacity"`
	CapacityUnit        *string  `json:"CapacityUnit"`
	CapacityWeight      *float32 `json:"CapacityWeight"`
	CapacityWeightUnit  *string  `json:"CapacityWeightUnit"`
	CreationDate        string   `json:"CreationDate"`
	LastChangeDate      string   `json:"LastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
