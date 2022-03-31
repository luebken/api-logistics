package model

type Vessels struct {
	Vessels []Vessel `json:"vessels"`
}

type Vessel struct {
	ID          uint64  `json:"id"`
	Description string  `json:"description"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

type Container struct {
	ID       uint64 `json:"id"`
	VesselID uint64 `json:"vessel_id"`
	Size     string `json:"size"`
}

// type Cargo struct {
// 	ID          uint64 `json:"id"`
// 	ContainerID uint64 `json:"container_id"`
// 	Weight      string `json:"weight"`
// 	Length      string `json:"length"`
// 	Width       string `json:"width"`
// }

var vessels []Vessel
var containers []Container

func init() {
	vessels = []Vessel{
		{
			ID:          1001,
			Description: "Petunia Seaways",
			Lat:         55.703461,
			Lng:         12.5937371,
		},
		{
			ID:          1002,
			Description: "Primula Seaways",
			Lat:         55.703461,
			Lng:         12.5937371,
		},
		{
			ID:          1003,
			Description: "Humber Viking",
			Lat:         55.703461,
			Lng:         12.5937371,
		},
	}

	containers = []Container{
		{
			ID:       2001,
			VesselID: 1001,
			Size:     "20ft",
		},
		{
			ID:       2002,
			VesselID: 1001,
			Size:     "20ft",
		},
		{
			ID:       2003,
			VesselID: 1002,
			Size:     "20ft",
		},
		{
			ID:       2004,
			VesselID: 1002,
			Size:     "20ft",
		},
		{
			ID:       2005,
			VesselID: 1002,
			Size:     "20ft",
		},
		{
			ID:       2006,
			VesselID: 1002,
			Size:     "20ft",
		},
	}

}

func GetAllVessels() []Vessel {
	return vessels
}

func GetAllContainers() []Container {
	return containers
}
func GetContainer(id uint64) Container {
	for _, c := range containers {
		if c.ID == id {
			return c
		}
	}
	return Container{}
}
func GetVessel(id uint64) Vessel {
	for _, v := range vessels {
		if v.ID == id {
			return v
		}
	}
	return Vessel{}
}
