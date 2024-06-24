package cmd

// TODO: use it to initialize according service object: https://chatgpt.com/c/8a16604b-c8c3-488f-aa1f-26cb18cbf1c9

type Purge struct {
	Config AlicloudConfig
}

func NewPurge(config AlicloudConfig) *Purge {
	p := Purge{
		Config: config,
	}
	return &p
}

func (p *Purge) Run() error {

	return nil
}

func (p *Purge) Scan() error {
	return nil
}
