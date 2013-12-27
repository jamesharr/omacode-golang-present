package lockmanager

type LockManager struct {
	requestCh chan LockRequest
	releaseCh chan ReleaseRequest
}

type LockRequest struct {
	Name      string
	GrantChan chan LockGrant
}

type LockGrant struct {
	Name      string
	ReleaseCh chan ReleaseRequest
}

type ReleaseRequest struct {
	Name string
}

func Create() LockManager {
	var lm LockManager
	lm.requestCh = make(chan LockRequest)
	lm.releaseCh = make(chan ReleaseRequest)
	go lm.managerRoutine()
	return lm
}

func (lm LockManager) managerRoutine() {
	locks := make(map[string]bool)
	waiting := make(map[string][]LockRequest)

	for {
		select {
		case req := <- lm.requestCh:
			if ! locks[req.Name] {
				// Hand out the request
				locks[req.Name] = true
				req.GrantChan <- LockGrant{req.Name, lm.releaseCh}
			} else {
				// Add on to waiting queue
				waiting[req.Name] = append(waiting[req.Name], req)
			}
		case req := <- lm.releaseCh:
			// Received a release request

			locks[req.Name] = false

			q := waiting[req.Name]
			if len(q) > 0 {
				// Grant lock
				locks[req.Name] = true
				q[0].GrantChan <- LockGrant{req.Name, lm.releaseCh}

				// Manage waiting queue
				q = q[1:]
				waiting[req.Name] = q
			}
		}
	}
}

func (lm LockManager) Lock(name string) LockGrant {
	var req LockRequest
	req.Name = name
	req.GrantChan = make(chan LockGrant)
	lm.requestCh <- req
	return <-req.GrantChan
}

func (grant LockGrant) Release() {
	grant.ReleaseCh <- ReleaseRequest{grant.Name}
}
