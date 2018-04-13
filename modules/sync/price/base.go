package price

type ISyncPrice interface {
	run()
}

type SyncPrice struct {
	Response string
}

func (syncPrice SyncPrice) GetResponse() string {
	return syncPrice.Response
}
