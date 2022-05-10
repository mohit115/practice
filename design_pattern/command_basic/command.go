package main

type command interface{
	execute()
}

type OnCommand struct{
	device device
}

func (c *OnCommand)execute()  {
	c.device.on()
}


type offCommand struct{
	device device
}

func (c *offCommand)execute()  {
	c.device.off()
}

type tuneCommand struct{
	device device
	channelId int
}

func (c *tuneCommand)execute()  {
	c.device.tune(c.channelId)
}

type scanCommand struct{
	device device
}

func (c *scanCommand)execute()  {
	c.device.scan()
}

type volumeUpCommand struct{
	device device
}

func (c *volumeUpCommand)execute()  {
	c.device.volumeUp()
}


type volumeDownCommand struct{
	device device
}

func (c *volumeDownCommand)execute()  {
	c.device.volumeDown()
}


type channelNextCommand struct{
	device device
}

func (c *channelNextCommand)execute()  {
	c.device.channelNext()
}

type channelPrevCommand struct{
	device device
}

func (c *channelPrevCommand)execute()  {
	c.device.channelPrev()
}