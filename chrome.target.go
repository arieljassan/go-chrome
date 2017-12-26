package chrome

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/protocol"
	target "github.com/mkenney/go-chrome/target"

	log "github.com/sirupsen/logrus"
)

/*
Target - https://chromedevtools.github.io/devtools-protocol/tot/Target/
Supports additional targets discovery and allows to attach to them.
*/
type Target struct{}

/*
ActivateTarget activates (focuses) the target.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-activateTarget
*/
func (Target) ActivateTarget(
	socket *Socket,
	params *target.ActivateTargetParams,
) error {
	command := &protocol.Command{
		Method: "Target.activateTarget",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
AttachToTarget attaches to the target with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-attachToTarget
*/
func (Target) AttachToTarget(
	socket *Socket,
	params *target.AttachToTargetParams,
) error {
	command := &protocol.Command{
		Method: "Target.attachToTarget",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
CloseTarget closes the target. If the target is a page that gets closed too.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-closeTarget
*/
func (Target) CloseTarget(
	socket *Socket,
	params *target.CloseTargetParams,
) (target.CloseTargetResult, error) {
	command := &protocol.Command{
		Method: "Target.closeTarget",
	}
	result := target.CloseTargetResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CreateBrowserContext creates a new empty BrowserContext. Similar to an incognito profile but you can
have more than one. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createBrowserContext
*/
func (Target) CreateBrowserContext(
	socket *Socket,
) (target.CreateBrowserContextResult, error) {
	command := &protocol.Command{
		Method: "Target.createBrowserContext",
	}
	result := target.CreateBrowserContextResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
CreateTarget creates a new page.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-createTarget
*/
func (Target) CreateTarget(
	socket *Socket,
	params *target.CreateTargetParams,
) (target.CreateTargetResult, error) {
	command := &protocol.Command{
		Method: "Target.createTarget",
		Params: params,
	}
	result := target.CreateTargetResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
DetachFromTarget detaches session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-detachFromTarget
*/
func (Target) DetachFromTarget(
	socket *Socket,
	params *target.DetachFromTargetParams,
) error {
	command := &protocol.Command{
		Method: "Target.detachFromTarget",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
DisposeBrowserContext deletes a BrowserContext, will fail of any open page uses it. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-disposeBrowserContext
*/
func (Target) DisposeBrowserContext(
	socket *Socket,
	params *target.DisposeBrowserContextParams,
) error {
	command := &protocol.Command{
		Method: "Target.disposeBrowserContext",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetTargetInfo returns information about a target. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargetInfo
*/
func (Target) GetTargetInfo(
	socket *Socket,
	params *target.GetTargetInfoParams,
) (target.GetTargetInfoResult, error) {
	command := &protocol.Command{
		Method: "Target.getTargetInfo",
		Params: params,
	}
	result := target.GetTargetInfoResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetTargets retrieves a list of available targets.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-getTargets
*/
func (Target) GetTargets(
	socket *Socket,
	params *target.GetTargetsParams,
) error {
	command := &protocol.Command{
		Method: "Target.getTargets",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SendMessageToTarget sends protocol message over session with given id.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-sendMessageToTarget
*/
func (Target) SendMessageToTarget(
	socket *Socket,
	params *target.SendMessageToTargetParams,
) error {
	command := &protocol.Command{
		Method: "Target.sendMessageToTarget",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAttachToFrames EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAttachToFrames
*/
func (Target) SetAttachToFrames(
	socket *Socket,
	params *target.SetAttachToFramesParams,
) error {
	command := &protocol.Command{
		Method: "Target.setAttachToFrames",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetAutoAttach controls whether to automatically attach to new targets which are considered to be
related to this one. When turned on, attaches to all existing related targets as well. When turned
off, automatically detaches from all currently attached targets. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setAutoAttach
*/
func (Target) SetAutoAttach(
	socket *Socket,
	params *target.SetAutoAttachParams,
) error {
	command := &protocol.Command{
		Method: "Target.setAutoAttach",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetDiscoverTargets controls whether to discover available targets and notify via
`targetCreated/targetInfoChanged/targetDestroyed` events.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setDiscoverTargets
*/
func (Target) SetDiscoverTargets(
	socket *Socket,
	params *target.SetDiscoverTargetsParams,
) error {
	command := &protocol.Command{
		Method: "Target.setDiscoverTargets",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
SetRemoteLocations enables target discovery for the specified locations, when `setDiscoverTargets`
was set to `true`. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#method-setRemoteLocations
*/
func (Target) SetRemoteLocations(
	socket *Socket,
	params *target.SetRemoteLocationsParams,
) error {
	command := &protocol.Command{
		Method: "Target.setRemoteLocations",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAttachedToTarget adds a handler to the Target.attachedToTarget event.
Target.attachedToTarget fires when attached to target because of auto-attach or `attachToTarget`
command. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-attachedToTarget
*/
func (Target) OnAttachedToTarget(
	socket *Socket,
	callback func(event *target.AttachedToTargetEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.attachedToTarget",
		func(name string, params []byte) {
			event := &target.AttachedToTargetEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnDetachedFromTarget adds a handler to the Target.detachedFromTarget event.
Target.detachedFromTarget fires when detached from target for any reason (including
`detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been
attached to it. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-detachedFromTarget
*/
func (Target) OnDetachedFromTarget(
	socket *Socket,
	callback func(event *target.DetachedFromTargetEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.detachedFromTarget",
		func(name string, params []byte) {
			event := &target.DetachedFromTargetEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnReceivedMessageFromTarget adds a handler to the Target.receivedMessageFromTarget event.
Target.receivedMessageFromTarget fires when a new protocol message received from the session (as
reported in `attachedToTarget` event).

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-receivedMessageFromTarget
*/
func (Target) OnReceivedMessageFromTarget(
	socket *Socket,
	callback func(event *target.ReceivedMessageFromTargetEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.receivedMessageFromTarget",
		func(name string, params []byte) {
			event := &target.ReceivedMessageFromTargetEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnTargetCreated adds a handler to the Target.Created event. Target.Created fires when a possible
inspection target is created.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetCreated
*/
func (Target) OnTargetCreated(
	socket *Socket,
	callback func(event *target.CreatedEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.targetCreated",
		func(name string, params []byte) {
			event := &target.CreatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnTargetDestroyed adds a handler to the Target.Destroyed event. Target.Destroyed fires when a target
is destroyed.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetDestroyed
*/
func (Target) OnTargetDestroyed(
	socket *Socket,
	callback func(event *target.DestroyedEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.targetDestroyed",
		func(name string, params []byte) {
			event := &target.DestroyedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnTargetInfoChanged adds a handler to the Target.InfoChanged event. Target.InfoChanged fires when
some information about a target has changed. This only happens between `targetCreated` and
`targetDestroyed`.

https://chromedevtools.github.io/devtools-protocol/tot/Target/#event-targetInfoChanged
*/
func (Target) OnTargetInfoChanged(
	socket *Socket,
	callback func(event *target.InfoChangedEvent),
) {
	handler := protocol.NewEventHandler(
		"Target.targetInfoChanged",
		func(name string, params []byte) {
			event := &target.InfoChangedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
