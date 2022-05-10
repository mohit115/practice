package main

type button struct{
	name string
	command command
}

func (b *button) press()  {
	b.command.execute()
}


type remote struct{
	buttons map[string]*button
}