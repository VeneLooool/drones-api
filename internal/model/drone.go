package model

type DroneStatus string

func (ds DroneStatus) String() string {
	return string(ds)
}

const (
	DroneStatusActive DroneStatus = "active"
	DroneStatusReady  DroneStatus = "ready"
	DroneStatusError  DroneStatus = "error"
)

type Drone struct {
	ID     uint64      `db:"id"`
	Name   string      `db:"name"`
	Status DroneStatus `db:"status"`
}

func (d *Drone) SetDefaultStatus() {
	d.Status = DroneStatusReady
}

func (d *Drone) SetStatus(status DroneStatus) {
	d.Status = status
}
