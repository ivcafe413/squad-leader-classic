package game

type GameMessageProcessor struct {
}

func (proc *GameMessageProcessor) ProcessInput(msg interface{}) []byte {
	//Switch on the different potential structs that the message can coerce into
	switch msg.(type) {
	default:
		return nil
	}
}
