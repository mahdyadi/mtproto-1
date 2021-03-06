package mtproto

func (m *MTProto) MessagesGetHistory(peer TL, offsetId, offsetDate, addOffset, limit, maxId, minId int32) (*TL, error) {
	resp := make(chan response, 1)
	m.queueSend <- packetToSend{
		msg: TL_messages_getHistory{
			Peer:        peer,
			Offset_id:   offsetId,
			Offset_date: offsetDate,
			Add_offset:  addOffset,
			Limit:       limit,
			Max_id:      maxId,
			Min_id:      minId,
		},
		resp: resp,
	}
	x := <-resp

	if x.err != nil {
		return nil, x.err
	}

	return &x.data, nil
}

func (m *MTProto) MessagesGetDialogs(excludePinned bool, offsetDate, offsetId int32, offsetPeer TL, limit int32) (*TL, error) {
	resp := make(chan response, 1)
	m.queueSend <- packetToSend{
		msg: TL_messages_getDialogs{
			Exclude_pinned: excludePinned,
			Offset_date:    offsetDate,
			Offset_id:      offsetId,
			Offset_peer:    offsetPeer,
			Limit:          limit,
		},
		resp: resp,
	}
	x := <-resp

	if x.err != nil {
		return nil, x.err
	}

	return &x.data, nil
}

func (m *MTProto) MessagesSendMessage(no_webpage, silent, background, clear_draft bool, peer TL, reply_to_msg_id int32, message string, random_id int64, reply_markup TL, entities []TL) (*TL, error) {
	resp := make(chan response, 1)
	m.queueSend <- packetToSend{
		msg: TL_messages_sendMessage{
			No_webpage:      no_webpage,
			Silent:          silent,
			Background:      background,
			Clear_draft:     clear_draft,
			Peer:            peer,
			Reply_to_msg_id: reply_to_msg_id,
			Message:         message,
			Random_id:       random_id,
			Reply_markup:    reply_markup,
			Entities:        entities,
		},
		resp: resp,
	}
	x := <-resp
	if x.err != nil {
		return nil, x.err
	}

	return &x.data, nil
}
