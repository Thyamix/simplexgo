package simplexgo

// Event is the discriminated union of all API events.
type Event interface {
	eventType()
}

// ─── Contact connection events ───────────────────────────────────────────────

// ContactConnected is sent after a user connects via bot SimpleX address
// (not a business address).
type ContactConnected struct {
	User              User     `json:"user"`
	Contact           Contact  `json:"contact"`
	UserCustomProfile *Profile `json:"userCustomProfile,omitempty"`
}

// ContactUpdated is sent when the contact profile of another user is updated.
type ContactUpdated struct {
	User        User    `json:"user"`
	FromContact Contact `json:"fromContact"`
	ToContact   Contact `json:"toContact"`
}

// ContactDeletedByContact is sent when bot user's connection with another
// contact is deleted (conversation is kept).
type ContactDeletedByContact struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// ReceivedContactRequest is sent when a contact request is received.
// Only sent when auto-accept is disabled.
type ReceivedContactRequest struct {
	User           User               `json:"user"`
	ContactRequest UserContactRequest `json:"contactRequest"`
	Chat_          *AChat             `json:"chat_,omitempty"`
}

// NewMemberContactReceivedInv is sent when an invitation to connect directly
// with a group member is received.
type NewMemberContactReceivedInv struct {
	User      User        `json:"user"`
	Contact   Contact     `json:"contact"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Member    GroupMember `json:"member"`
}

// ContactSndReady is sent when connecting via 1-time invitation or after
// accepting a contact request. After this event the bot can send messages.
type ContactSndReady struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// ─── Message events ──────────────────────────────────────────────────────────

// NewChatItems is sent when one or more messages are received.
type NewChatItems struct {
	User      User        `json:"user"`
	ChatItems []AChatItem `json:"chatItems"`
}

// ChatItemReaction is sent when a message reaction is received.
type ChatItemReaction struct {
	User     User        `json:"user"`
	Added    bool        `json:"added"`
	Reaction ACIReaction `json:"reaction"`
}

// ChatItemsDeleted is sent when a message is deleted by another user.
type ChatItemsDeleted struct {
	User              User               `json:"user"`
	ChatItemDeletions []ChatItemDeletion `json:"chatItemDeletions"`
	ByUser            bool               `json:"byUser"`
	Timed             bool               `json:"timed"`
}

// ChatItemUpdated is sent when a message is updated by another user.
type ChatItemUpdated struct {
	User     User      `json:"user"`
	ChatItem AChatItem `json:"chatItem"`
}

// GroupChatItemsDeleted is sent when group messages are deleted or moderated.
type GroupChatItemsDeleted struct {
	User        User         `json:"user"`
	GroupInfo   GroupInfo    `json:"groupInfo"`
	ChatItemIDs []int64      `json:"chatItemIDs"`
	ByUser      bool         `json:"byUser"`
	Member_     *GroupMember `json:"member_,omitempty"`
}

// ChatItemsStatusesUpdated is sent for message delivery status updates.
type ChatItemsStatusesUpdated struct {
	User      User        `json:"user"`
	ChatItems []AChatItem `json:"chatItems"`
}

// ─── Group events ────────────────────────────────────────────────────────────

// ReceivedGroupInvitation is sent when a group invitation is received.
type ReceivedGroupInvitation struct {
	User           User            `json:"user"`
	GroupInfo      GroupInfo       `json:"groupInfo"`
	Contact        Contact         `json:"contact"`
	FromMemberRole GroupMemberRole `json:"fromMemberRole"`
	MemberRole     GroupMemberRole `json:"memberRole"`
}

// UserJoinedGroup is sent when the bot user joins a group (connection via
// group link completes).
type UserJoinedGroup struct {
	User       User        `json:"user"`
	GroupInfo  GroupInfo   `json:"groupInfo"`
	HostMember GroupMember `json:"hostMember"`
}

// GroupUpdated is sent when group profile or preferences are updated.
type GroupUpdated struct {
	User      User         `json:"user"`
	FromGroup GroupInfo    `json:"fromGroup"`
	ToGroup   GroupInfo    `json:"toGroup"`
	Member_   *GroupMember `json:"member_,omitempty"`
}

// JoinedGroupMember is sent when another member joins the group.
type JoinedGroupMember struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Member    GroupMember `json:"member"`
}

// MemberRole is sent when a member's (or bot user's) group role changes.
type MemberRole struct {
	User      User            `json:"user"`
	GroupInfo GroupInfo       `json:"groupInfo"`
	ByMember  GroupMember     `json:"byMember"`
	Member    GroupMember     `json:"member"`
	FromRole  GroupMemberRole `json:"fromRole"`
	ToRole    GroupMemberRole `json:"toRole"`
}

// DeletedMember is sent when another member is removed from the group.
type DeletedMember struct {
	User          User        `json:"user"`
	GroupInfo     GroupInfo   `json:"groupInfo"`
	ByMember      GroupMember `json:"byMember"`
	DeletedMember GroupMember `json:"deletedMember"`
	WithMessages  bool        `json:"withMessages"`
}

// LeftMember is sent when another member leaves the group.
type LeftMember struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Member    GroupMember `json:"member"`
}

// DeletedMemberUser is sent when the bot user is removed from the group.
type DeletedMemberUser struct {
	User         User        `json:"user"`
	GroupInfo    GroupInfo   `json:"groupInfo"`
	Member       GroupMember `json:"member"`
	WithMessages bool        `json:"withMessages"`
}

// GroupDeleted is sent when a group is deleted by its owner (not the bot user).
type GroupDeleted struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Member    GroupMember `json:"member"`
}

// ConnectedToGroupMember is sent when connected to another group member.
type ConnectedToGroupMember struct {
	User          User        `json:"user"`
	GroupInfo     GroupInfo   `json:"groupInfo"`
	Member        GroupMember `json:"member"`
	MemberContact *Contact    `json:"memberContact,omitempty"`
}

// MemberAcceptedByOther is sent when another group owner, admin or moderator
// accepts a member after review ("knocking").
type MemberAcceptedByOther struct {
	User            User        `json:"user"`
	GroupInfo       GroupInfo   `json:"groupInfo"`
	AcceptingMember GroupMember `json:"acceptingMember"`
	Member          GroupMember `json:"member"`
}

// MemberBlockedForAll is sent when another member is blocked for all members.
type MemberBlockedForAll struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	ByMember  GroupMember `json:"byMember"`
	Member    GroupMember `json:"member"`
	Blocked   bool        `json:"blocked"`
}

// GroupMemberUpdated is sent when another group member's profile is updated.
type GroupMemberUpdated struct {
	User       User        `json:"user"`
	GroupInfo  GroupInfo   `json:"groupInfo"`
	FromMember GroupMember `json:"fromMember"`
	ToMember   GroupMember `json:"toMember"`
}

// ─── File events ─────────────────────────────────────────────────────────────

// RcvFileDescrReady is sent when a file is ready to be received.
type RcvFileDescrReady struct {
	User            User            `json:"user"`
	ChatItem        AChatItem       `json:"chatItem"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
	RcvFileDescr    RcvFileDescr    `json:"rcvFileDescr"`
}

// RcvFileComplete is sent when file reception is completed.
type RcvFileComplete struct {
	User     User      `json:"user"`
	ChatItem AChatItem `json:"chatItem"`
}

// SndFileCompleteXFTP is sent when file upload is completed.
type SndFileCompleteXFTP struct {
	User             User             `json:"user"`
	ChatItem         AChatItem        `json:"chatItem"`
	FileTransferMeta FileTransferMeta `json:"fileTransferMeta"`
}

// RcvFileStart is sent when file reception starts (after RcvFileDescrReady).
type RcvFileStart struct {
	User     User      `json:"user"`
	ChatItem AChatItem `json:"chatItem"`
}

// RcvFileSndCancelled is sent when a file is cancelled by the sender.
type RcvFileSndCancelled struct {
	User            User            `json:"user"`
	ChatItem        AChatItem       `json:"chatItem"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
}

// RcvFileAccepted is sent when a file is automatically accepted via CLI option.
type RcvFileAccepted struct {
	User     User      `json:"user"`
	ChatItem AChatItem `json:"chatItem"`
}

// RcvFileError is sent on error receiving a file.
type RcvFileError struct {
	User            User            `json:"user"`
	ChatItem_       *AChatItem      `json:"chatItem_,omitempty"`
	AgentError      AgentErrorType  `json:"agentError"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
}

// RcvFileWarning is sent when a warning occurs while receiving a file.
type RcvFileWarning struct {
	User            User            `json:"user"`
	ChatItem_       *AChatItem      `json:"chatItem_,omitempty"`
	AgentError      AgentErrorType  `json:"agentError"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
}

// SndFileError is sent on error sending a file.
type SndFileError struct {
	User             User             `json:"user"`
	ChatItem_        *AChatItem       `json:"chatItem_,omitempty"`
	FileTransferMeta FileTransferMeta `json:"fileTransferMeta"`
	ErrorMessage     string           `json:"errorMessage"`
}

// SndFileWarning is sent when a warning occurs while sending a file.
type SndFileWarning struct {
	User             User             `json:"user"`
	ChatItem_        *AChatItem       `json:"chatItem_,omitempty"`
	FileTransferMeta FileTransferMeta `json:"fileTransferMeta"`
	ErrorMessage     string           `json:"errorMessage"`
}

// ─── Connection progress events ──────────────────────────────────────────────

// AcceptingContactRequest is sent when automatically accepting a contact
// request via bot's SimpleX address with auto-accept enabled.
type AcceptingContactRequest struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// AcceptingBusinessRequest is sent when automatically accepting a contact
// request via bot's business address.
type AcceptingBusinessRequest struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

// ContactConnecting is sent when a contact confirms a connection (via 1-time
// invitation link or when bot connects to another SimpleX address).
type ContactConnecting struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// BusinessLinkConnecting is sent when the bot connects to another business
// address and the contact confirms the connection.
type BusinessLinkConnecting struct {
	User        User        `json:"user"`
	GroupInfo   GroupInfo   `json:"groupInfo"`
	HostMember  GroupMember `json:"hostMember"`
	FromContact Contact     `json:"fromContact"`
}

// JoinedGroupMemberConnecting is sent when a group member is announced to the
// group and will be connecting to the bot.
type JoinedGroupMemberConnecting struct {
	User       User        `json:"user"`
	GroupInfo  GroupInfo   `json:"groupInfo"`
	HostMember GroupMember `json:"hostMember"`
	Member     GroupMember `json:"member"`
}

// SentGroupInvitation is sent when another user joins a group via the bot's
// link.
type SentGroupInvitation struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Contact   Contact     `json:"contact"`
	Member    GroupMember `json:"member"`
}

// GroupLinkConnecting is sent when the bot joins a group via another user's
// link.
type GroupLinkConnecting struct {
	User       User        `json:"user"`
	GroupInfo  GroupInfo   `json:"groupInfo"`
	HostMember GroupMember `json:"hostMember"`
}

// ─── Error events ────────────────────────────────────────────────────────────

// MessageError is sent on a message error.
type MessageError struct {
	User         User   `json:"user"`
	Severity     string `json:"severity"`
	ErrorMessage string `json:"errorMessage"`
}

// ChatErrorEvent is sent on a chat error.
// Named ChatErrorEvent to avoid collision with the ChatError type.
type ChatErrorEvent struct {
	ChatError ChatError `json:"chatError"`
}

// ChatErrors is sent when multiple chat errors occur.
type ChatErrors struct {
	ChatErrors []ChatError `json:"chatErrors"`
}

// ─── Interface implementation ────────────────────────────────────────────────

func (ContactConnected) eventType()            {}
func (ContactUpdated) eventType()              {}
func (ContactDeletedByContact) eventType()     {}
func (ReceivedContactRequest) eventType()      {}
func (NewMemberContactReceivedInv) eventType() {}
func (ContactSndReady) eventType()             {}

func (NewChatItems) eventType()             {}
func (ChatItemReaction) eventType()         {}
func (ChatItemsDeleted) eventType()         {}
func (ChatItemUpdated) eventType()          {}
func (GroupChatItemsDeleted) eventType()    {}
func (ChatItemsStatusesUpdated) eventType() {}

func (ReceivedGroupInvitation) eventType() {}
func (UserJoinedGroup) eventType()         {}
func (GroupUpdated) eventType()            {}
func (JoinedGroupMember) eventType()       {}
func (MemberRole) eventType()              {}
func (DeletedMember) eventType()           {}
func (LeftMember) eventType()              {}
func (DeletedMemberUser) eventType()       {}
func (GroupDeleted) eventType()            {}
func (ConnectedToGroupMember) eventType()  {}
func (MemberAcceptedByOther) eventType()   {}
func (MemberBlockedForAll) eventType()     {}
func (GroupMemberUpdated) eventType()      {}

func (RcvFileDescrReady) eventType()   {}
func (RcvFileComplete) eventType()     {}
func (SndFileCompleteXFTP) eventType() {}
func (RcvFileStart) eventType()        {}
func (RcvFileSndCancelled) eventType() {}
func (RcvFileAccepted) eventType()     {}
func (RcvFileError) eventType()        {}
func (RcvFileWarning) eventType()      {}
func (SndFileError) eventType()        {}
func (SndFileWarning) eventType()      {}

func (AcceptingContactRequest) eventType()     {}
func (AcceptingBusinessRequest) eventType()    {}
func (ContactConnecting) eventType()           {}
func (BusinessLinkConnecting) eventType()      {}
func (JoinedGroupMemberConnecting) eventType() {}
func (SentGroupInvitation) eventType()         {}
func (GroupLinkConnecting) eventType()         {}

func (MessageError) eventType()   {}
func (ChatErrorEvent) eventType() {}
func (ChatErrors) eventType()     {}
