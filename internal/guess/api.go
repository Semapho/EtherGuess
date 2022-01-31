package guess

func (c *Client) Run(parallel int) {
	finishChan := make(chan bool)
	for i := 0; i < parallel; i++ {
		go c.guessEther(finishChan)
	}
	<-finishChan
}
