package main

func main() {
	tv := &tv{channels: make(map[int]interface{}),maxVolume: 100}
	radio := &radio{channels: make(map[int]interface{}),maxVolume: 30}

	tvRemote := &remote{buttons: make(map[string]*button)}
	tvRemote.buttons["on"] = &button{name: "on",command:&OnCommand{device: tv} }
	tvRemote.buttons["off"] = &button{name: "off",command:&offCommand{device: tv} }
	tvRemote.buttons["goto"] = &button{name: "goto",command:&tuneCommand{device: tv,channelId: 1} }
	tvRemote.buttons["scan"] = &button{name: "scan",command:&scanCommand{device: tv} }
	tvRemote.buttons["volumeUP"] = &button{name: "volumeUP",command:&volumeUpCommand{device: tv} }
	tvRemote.buttons["volumeDown"] = &button{name: "volumeDown",command:&volumeDownCommand{device: tv} }
	tvRemote.buttons["channelNext"] = &button{name: "channelNext",command:&channelNextCommand{device: tv} }
	tvRemote.buttons["channelPrev"] = &button{name: "channelPrev",command:&channelPrevCommand{device: tv} }


	radioRemote := &remote{buttons: make(map[string]*button)}
	radioRemote.buttons["on"] = &button{name: "on",command:&OnCommand{device: radio} }
	radioRemote.buttons["off"] = &button{name: "off",command:&offCommand{device: radio} }
	radioRemote.buttons["goto"] = &button{name: "goto",command:&tuneCommand{device: radio,channelId: 1} }
	radioRemote.buttons["scan"] = &button{name: "scan",command:&scanCommand{device: radio} }
	radioRemote.buttons["volumeUP"] = &button{name: "volumeUP",command:&volumeUpCommand{device: radio} }
	radioRemote.buttons["volumeDown"] = &button{name: "volumeDown",command:&volumeDownCommand{device: radio} }
	radioRemote.buttons["channelNext"] = &button{name: "channelNext",command:&channelNextCommand{device: radio} }
	radioRemote.buttons["channelPrev"] = &button{name: "channelPrev",command:&channelPrevCommand{device: radio} }


	tvRemote.buttons["on"].press()
	tvRemote.buttons["scan"].press()


	radioRemote.buttons["on"].press()
	radioRemote.buttons["scan"].press()
}
