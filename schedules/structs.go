package schedules

import "time"

type ListCompletedJobs struct {
	Result []struct {
		Scheduler         string
		Name              string
		CompletedSystems  int
		FailedSystems     int
		InProgressSystems int
		Id                int
		Type              string
		Earliest          time.Time
	}
}

type ListFailedJobs struct {
	Result []struct {
		Scheduler         string
		Name              string
		CompletedSystems  int
		FailedSystems     int
		InProgressSystems int
		Id                int
		Type              string
		Earliest          time.Time
	}
}

type ListPendingJobs struct {
	Result []struct {
		Scheduler         string
		Name              string
		CompletedSystems  int
		FailedSystems     int
		InProgressSystems int
		Id                int
		Type              string
		Earliest          time.Time
	}
}

type ListJobs struct {
	Completed ListCompletedJobs
	Failed    ListFailedJobs
	Pending   ListPendingJobs
}
