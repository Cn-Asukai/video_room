// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"video_room/app/video_room/internal/dao/internal"
)

// internalChatMessagesDao is internal type for wrapping internal DAO implements.
type internalChatMessagesDao = *internal.ChatMessagesDao

// chatMessagesDao is the data access object for table chat_messages.
// You can define custom methods on it to extend its functionality as you wish.
type chatMessagesDao struct {
	internalChatMessagesDao
}

var (
	// ChatMessages is globally public accessible object for table chat_messages operations.
	ChatMessages = chatMessagesDao{
		internal.NewChatMessagesDao(),
	}
)

// Fill with you ideas below.