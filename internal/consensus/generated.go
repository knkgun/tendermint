// Code generated by tmjson. DO NOT EDIT

package consensus

import (
	"encoding/json"
	"fmt"
)

type _typeTagged struct {
	T string          `json:"type"`
	V json.RawMessage `json:"value"`
}

const _typeTag_BlockPartMessage = "tendermint/BlockPart"

// MarshalJSON implements the json.Marshaler interface for BlockPartMessage.
// It wraps the encoding in a type-tagged object.
func (v BlockPartMessage) MarshalJSON() ([]byte, error) {
	type shim BlockPartMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_BlockPartMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for BlockPartMessage.
// It expects a type-tagged object with the tag "tendermint/BlockPart".
func (v *BlockPartMessage) UnmarshalJSON(data []byte) error {
	type shim BlockPartMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_BlockPartMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_BlockPartMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_EndHeightMessage = "tendermint/wal/EndHeightMessage"

// MarshalJSON implements the json.Marshaler interface for EndHeightMessage.
// It wraps the encoding in a type-tagged object.
func (v EndHeightMessage) MarshalJSON() ([]byte, error) {
	type shim EndHeightMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_EndHeightMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for EndHeightMessage.
// It expects a type-tagged object with the tag "tendermint/wal/EndHeightMessage".
func (v *EndHeightMessage) UnmarshalJSON(data []byte) error {
	type shim EndHeightMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_EndHeightMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_EndHeightMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_HasVoteMessage = "tendermint/HasVote"

// MarshalJSON implements the json.Marshaler interface for HasVoteMessage.
// It wraps the encoding in a type-tagged object.
func (v HasVoteMessage) MarshalJSON() ([]byte, error) {
	type shim HasVoteMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_HasVoteMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for HasVoteMessage.
// It expects a type-tagged object with the tag "tendermint/HasVote".
func (v *HasVoteMessage) UnmarshalJSON(data []byte) error {
	type shim HasVoteMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_HasVoteMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_HasVoteMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_NewRoundStepMessage = "tendermint/NewRoundStepMessage"

// MarshalJSON implements the json.Marshaler interface for NewRoundStepMessage.
// It wraps the encoding in a type-tagged object.
func (v NewRoundStepMessage) MarshalJSON() ([]byte, error) {
	type shim NewRoundStepMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_NewRoundStepMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for NewRoundStepMessage.
// It expects a type-tagged object with the tag "tendermint/NewRoundStepMessage".
func (v *NewRoundStepMessage) UnmarshalJSON(data []byte) error {
	type shim NewRoundStepMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_NewRoundStepMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_NewRoundStepMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_NewValidBlockMessage = "tendermint/NewValidBlockMessage"

// MarshalJSON implements the json.Marshaler interface for NewValidBlockMessage.
// It wraps the encoding in a type-tagged object.
func (v NewValidBlockMessage) MarshalJSON() ([]byte, error) {
	type shim NewValidBlockMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_NewValidBlockMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for NewValidBlockMessage.
// It expects a type-tagged object with the tag "tendermint/NewValidBlockMessage".
func (v *NewValidBlockMessage) UnmarshalJSON(data []byte) error {
	type shim NewValidBlockMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_NewValidBlockMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_NewValidBlockMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_ProposalMessage = "tendermint/Proposal"

// MarshalJSON implements the json.Marshaler interface for ProposalMessage.
// It wraps the encoding in a type-tagged object.
func (v ProposalMessage) MarshalJSON() ([]byte, error) {
	type shim ProposalMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_ProposalMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProposalMessage.
// It expects a type-tagged object with the tag "tendermint/Proposal".
func (v *ProposalMessage) UnmarshalJSON(data []byte) error {
	type shim ProposalMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_ProposalMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_ProposalMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_ProposalPOLMessage = "tendermint/ProposalPOL"

// MarshalJSON implements the json.Marshaler interface for ProposalPOLMessage.
// It wraps the encoding in a type-tagged object.
func (v ProposalPOLMessage) MarshalJSON() ([]byte, error) {
	type shim ProposalPOLMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_ProposalPOLMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProposalPOLMessage.
// It expects a type-tagged object with the tag "tendermint/ProposalPOL".
func (v *ProposalPOLMessage) UnmarshalJSON(data []byte) error {
	type shim ProposalPOLMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_ProposalPOLMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_ProposalPOLMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_VoteMessage = "tendermint/Vote"

// MarshalJSON implements the json.Marshaler interface for VoteMessage.
// It wraps the encoding in a type-tagged object.
func (v VoteMessage) MarshalJSON() ([]byte, error) {
	type shim VoteMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_VoteMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoteMessage.
// It expects a type-tagged object with the tag "tendermint/Vote".
func (v *VoteMessage) UnmarshalJSON(data []byte) error {
	type shim VoteMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_VoteMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_VoteMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_VoteSetBitsMessage = "tendermint/VoteSetBits"

// MarshalJSON implements the json.Marshaler interface for VoteSetBitsMessage.
// It wraps the encoding in a type-tagged object.
func (v VoteSetBitsMessage) MarshalJSON() ([]byte, error) {
	type shim VoteSetBitsMessage
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_VoteSetBitsMessage, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoteSetBitsMessage.
// It expects a type-tagged object with the tag "tendermint/VoteSetBits".
func (v *VoteSetBitsMessage) UnmarshalJSON(data []byte) error {
	type shim VoteSetBitsMessage
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_VoteSetBitsMessage {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_VoteSetBitsMessage)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_VoteSetMaj23Message = "tendermint/VoteSetMaj23"

// MarshalJSON implements the json.Marshaler interface for VoteSetMaj23Message.
// It wraps the encoding in a type-tagged object.
func (v VoteSetMaj23Message) MarshalJSON() ([]byte, error) {
	type shim VoteSetMaj23Message
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_VoteSetMaj23Message, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoteSetMaj23Message.
// It expects a type-tagged object with the tag "tendermint/VoteSetMaj23".
func (v *VoteSetMaj23Message) UnmarshalJSON(data []byte) error {
	type shim VoteSetMaj23Message
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_VoteSetMaj23Message {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_VoteSetMaj23Message)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_msgInfo = "tendermint/wal/MsgInfo"

// MarshalJSON implements the json.Marshaler interface for msgInfo.
// It wraps the encoding in a type-tagged object.
func (v msgInfo) MarshalJSON() ([]byte, error) {
	type shim msgInfo
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_msgInfo, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for msgInfo.
// It expects a type-tagged object with the tag "tendermint/wal/MsgInfo".
func (v *msgInfo) UnmarshalJSON(data []byte) error {
	type shim msgInfo
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_msgInfo {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_msgInfo)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}

const _typeTag_timeoutInfo = "tendermint/wal/TimeoutInfo"

// MarshalJSON implements the json.Marshaler interface for timeoutInfo.
// It wraps the encoding in a type-tagged object.
func (v timeoutInfo) MarshalJSON() ([]byte, error) {
	type shim timeoutInfo
	value, err := json.Marshal((*shim)(&v))
	if err != nil {
		return nil, err
	}
	return json.Marshal(_typeTagged{T: _typeTag_timeoutInfo, V: value})
}

// UnmarshalJSON implements the json.Unmarshaler interface for timeoutInfo.
// It expects a type-tagged object with the tag "tendermint/wal/TimeoutInfo".
func (v *timeoutInfo) UnmarshalJSON(data []byte) error {
	type shim timeoutInfo
	var tmp _typeTagged
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	} else if tmp.T != _typeTag_timeoutInfo {
		return fmt.Errorf("wrong type tag %q for %q", tmp.T, _typeTag_timeoutInfo)
	}
	return json.Unmarshal(tmp.V, (*shim)(v))
}
