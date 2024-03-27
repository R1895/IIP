package serializer

import "iip/model"

type Workshop struct {
	ID           uint          `json:"id"`
	WorkshopName string        `json:"workshop_name"`
	Worker       Worker        `json:"worker"`
	Workers      []Worker      `json:"workers"`
	Machines     []Machine     `json:"machines"`
	Tasks        []Task        `json:"tasks"`
	Environments []Environment `json:"enviroments"`
	Noises       []Noise       `json:"noises"`
	Smells       []Smell       `json:"smells"`
	Safes        []Safe        `json:"safes"`
}



func BuildWorkshop(workshop model.Workshop) Workshop {
	return Workshop{
		ID: workshop.ID,

		WorkshopName: workshop.WorkshopName,
		Worker:       BuildWorker(workshop.Worker),
		Workers:      BuildWorkers(workshop.Workers),
		Machines:     BuildMachines(workshop.Machines),
		Tasks:        BuildTasks(workshop.Tasks),
		Environments: BuildEnvironments(workshop.Environments),
		Noises:       BuildNoises(workshop.Noises),
		Smells:       BuildSmells(workshop.Smells),
		Safes:        BuildSafes(workshop.Safes),
	}
}

func BuildWorkshops(items []model.Workshop) (workshops []Workshop) {
	for _, item := range items {
		workshop := BuildWorkshop(item)
		workshops = append(workshops, workshop)
	}
	return workshops
}
