package simplexgo

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ─── Response ────────────────────────────────────────────────────────────────

// Response is the discriminated union of all API command responses.
type Response interface {
	responseType()
}

// ChatCmdError is returned by every command on failure.
type ChatCmdError struct {
	ChatError ChatError `json:"chatError"`
}

func (ChatCmdError) responseType() {}

// ─── Shared responses (used by multiple commands) ────────────────────────────

type ActiveUser struct {
	User User `json:"user"`
}

type UserProfileUpdated struct {
	User          User                     `json:"user"`
	FromProfile   Profile                  `json:"fromProfile"`
	ToProfile     Profile                  `json:"toProfile"`
	UpdateSummary UserProfileUpdateSummary `json:"updateSummary"`
}

type UserProfileNoChange struct {
	User User `json:"user"`
}

type CmdOk struct {
	User_ *User `json:"user_,omitempty"`
}

// NewChatItemsResponse re-uses the NewChatItems event struct; both share the
// same JSON shape. It is type-aliased here so callers can use it as a Response.
type NewChatItemsResponse = NewChatItems

type ChatItemUpdatedResponse = ChatItemUpdated

type ChatItemsDeletedResponse = ChatItemsDeleted

type ChatItemReactionResponse = ChatItemReaction

type GroupUpdatedResponse = GroupUpdated

type SentGroupInvitationResponse = SentGroupInvitation

type AcceptingContactRequestResponse = AcceptingContactRequest

func (ActiveUser) responseType()          {}
func (UserProfileUpdated) responseType()  {}
func (UserProfileNoChange) responseType() {}
func (CmdOk) responseType()               {}

// ─── Address command responses ────────────────────────────────────────────────

// UserContactLinkCreated is returned by APICreateMyAddress.
type UserContactLinkCreated struct {
	User            User            `json:"user"`
	ConnLinkContact CreatedConnLink `json:"connLinkContact"`
}

// UserContactLinkDeleted is returned by APIDeleteMyAddress.
type UserContactLinkDeleted struct {
	User User `json:"user"`
}

// UserContactLinkResponse is returned by APIShowMyAddress.
type UserContactLinkResponse struct {
	User        User            `json:"user"`
	ContactLink UserContactLink `json:"contactLink"`
}

// UserContactLinkUpdated is returned by APISetAddressSettings.
type UserContactLinkUpdated struct {
	User        User            `json:"user"`
	ContactLink UserContactLink `json:"contactLink"`
}

func (UserContactLinkCreated) responseType()  {}
func (UserContactLinkDeleted) responseType()  {}
func (UserContactLinkResponse) responseType() {}
func (UserContactLinkUpdated) responseType()  {}

// ─── Message command responses ────────────────────────────────────────────────

// ChatItemNotChanged is returned by APIUpdateChatItem when nothing changed.
type ChatItemNotChanged struct {
	User     User      `json:"user"`
	ChatItem AChatItem `json:"chatItem"`
}

func (ChatItemNotChanged) responseType() {}

// ─── File command responses ───────────────────────────────────────────────────

// RcvFileAcceptedResponse is returned by ReceiveFile on success.
// (Aliases the event struct; same JSON shape.)
type RcvFileAcceptedResponse = RcvFileAccepted

// RcvFileAcceptedSndCancelled is returned by ReceiveFile when sender already
// cancelled the file.
type RcvFileAcceptedSndCancelled struct {
	User            User            `json:"user"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
}

// SndFileCancelled is returned by CancelFile when sending was cancelled.
type SndFileCancelled struct {
	User             User              `json:"user"`
	ChatItem_        *AChatItem        `json:"chatItem_,omitempty"`
	FileTransferMeta FileTransferMeta  `json:"fileTransferMeta"`
	SndFileTransfers []SndFileTransfer `json:"sndFileTransfers"`
}

// RcvFileCancelled is returned by CancelFile when receiving was cancelled.
type RcvFileCancelled struct {
	User            User            `json:"user"`
	ChatItem_       *AChatItem      `json:"chatItem_,omitempty"`
	RcvFileTransfer RcvFileTransfer `json:"rcvFileTransfer"`
}

func (RcvFileAcceptedSndCancelled) responseType() {}
func (SndFileCancelled) responseType()            {}
func (RcvFileCancelled) responseType()            {}

// ─── Group command responses ──────────────────────────────────────────────────

// UserAcceptedGroupSent is returned by APIJoinGroup.
type UserAcceptedGroupSent struct {
	User        User      `json:"user"`
	GroupInfo   GroupInfo `json:"groupInfo"`
	HostContact *Contact  `json:"hostContact,omitempty"`
}

// MemberAccepted is returned by APIAcceptMember.
type MemberAccepted struct {
	User      User        `json:"user"`
	GroupInfo GroupInfo   `json:"groupInfo"`
	Member    GroupMember `json:"member"`
}

// MembersRoleUser is returned by APIMembersRole.
type MembersRoleUser struct {
	User      User            `json:"user"`
	GroupInfo GroupInfo       `json:"groupInfo"`
	Members   []GroupMember   `json:"members"`
	ToRole    GroupMemberRole `json:"toRole"`
}

// MembersBlockedForAllUser is returned by APIBlockMembersForAll.
type MembersBlockedForAllUser struct {
	User      User          `json:"user"`
	GroupInfo GroupInfo     `json:"groupInfo"`
	Members   []GroupMember `json:"members"`
	Blocked   bool          `json:"blocked"`
}

// UserDeletedMembers is returned by APIRemoveMembers.
type UserDeletedMembers struct {
	User         User          `json:"user"`
	GroupInfo    GroupInfo     `json:"groupInfo"`
	Members      []GroupMember `json:"members"`
	WithMessages bool          `json:"withMessages"`
}

// LeftMemberUser is returned by APILeaveGroup.
type LeftMemberUser struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

// GroupMembers is returned by APIListMembers.
type GroupMembers struct {
	User  User  `json:"user"`
	Group Group `json:"group"`
}

// GroupCreated is returned by APINewGroup.
type GroupCreated struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

func (UserAcceptedGroupSent) responseType()    {}
func (MemberAccepted) responseType()           {}
func (MembersRoleUser) responseType()          {}
func (MembersBlockedForAllUser) responseType() {}
func (UserDeletedMembers) responseType()       {}
func (LeftMemberUser) responseType()           {}
func (GroupMembers) responseType()             {}
func (GroupCreated) responseType()             {}

// ─── Group link command responses ────────────────────────────────────────────

// GroupLinkCreated is returned by APICreateGroupLink.
type GroupLinkCreated struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
	GroupLink GroupLink `json:"groupLink"`
}

// GroupLinkResponse is returned by APIGroupLinkMemberRole and APIGetGroupLink.
type GroupLinkResponse struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
	GroupLink GroupLink `json:"groupLink"`
}

// GroupLinkDeleted is returned by APIDeleteGroupLink.
type GroupLinkDeleted struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

func (GroupLinkCreated) responseType()  {}
func (GroupLinkResponse) responseType() {}
func (GroupLinkDeleted) responseType()  {}

// ─── Connection command responses ────────────────────────────────────────────

// Invitation is returned by APIAddContact.
type Invitation struct {
	User               User                     `json:"user"`
	ConnLinkInvitation CreatedConnLink          `json:"connLinkInvitation"`
	Connection         PendingContactConnection `json:"connection"`
}

// ConnectionPlanResponse is returned by APIConnectPlan.
type ConnectionPlanResponse struct {
	User           User            `json:"user"`
	ConnLink       CreatedConnLink `json:"connLink"`
	ConnectionPlan ConnectionPlan  `json:"connectionPlan"`
}

// SentConfirmation is returned by APIConnect / Connect for 1-time invitation links.
type SentConfirmation struct {
	User              User                     `json:"user"`
	Connection        PendingContactConnection `json:"connection"`
	CustomUserProfile *Profile                 `json:"customUserProfile,omitempty"`
}

// ContactAlreadyExists is returned by APIConnect / Connect when already connected.
type ContactAlreadyExists struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// SentInvitation is returned by APIConnect / Connect for contact address links.
type SentInvitation struct {
	User              User                     `json:"user"`
	Connection        PendingContactConnection `json:"connection"`
	CustomUserProfile *Profile                 `json:"customUserProfile,omitempty"`
}

// ContactRequestRejected is returned by APIRejectContact.
type ContactRequestRejected struct {
	User           User               `json:"user"`
	ContactRequest UserContactRequest `json:"contactRequest"`
	Contact_       *Contact           `json:"contact_,omitempty"`
}

func (Invitation) responseType()             {}
func (ConnectionPlanResponse) responseType() {}
func (SentConfirmation) responseType()       {}
func (ContactAlreadyExists) responseType()   {}
func (SentInvitation) responseType()         {}
func (ContactRequestRejected) responseType() {}

// ─── Chat command responses ───────────────────────────────────────────────────

// ContactsList is returned by APIListContacts.
type ContactsList struct {
	User     User      `json:"user"`
	Contacts []Contact `json:"contacts"`
}

// GroupsList is returned by APIListGroups.
type GroupsList struct {
	User   User               `json:"user"`
	Groups []GroupInfoSummary `json:"groups"`
}

// ContactDeleted is returned by APIDeleteChat for a direct chat.
type ContactDeleted struct {
	User    User    `json:"user"`
	Contact Contact `json:"contact"`
}

// ContactConnectionDeleted is returned by APIDeleteChat for a pending connection.
type ContactConnectionDeleted struct {
	User       User                     `json:"user"`
	Connection PendingContactConnection `json:"connection"`
}

// GroupDeletedUser is returned by APIDeleteChat for a group.
type GroupDeletedUser struct {
	User      User      `json:"user"`
	GroupInfo GroupInfo `json:"groupInfo"`
}

func (ContactsList) responseType()             {}
func (GroupsList) responseType()               {}
func (ContactDeleted) responseType()           {}
func (ContactConnectionDeleted) responseType() {}
func (GroupDeletedUser) responseType()         {}

// ─── User profile command responses ──────────────────────────────────────────

// UsersList is returned by ListUsers.
type UsersList struct {
	Users []UserInfo `json:"users"`
}

// ContactPrefsUpdated is returned by APISetContactPrefs.
type ContactPrefsUpdated struct {
	User        User    `json:"user"`
	FromContact Contact `json:"fromContact"`
	ToContact   Contact `json:"toContact"`
}

func (UsersList) responseType()           {}
func (ContactPrefsUpdated) responseType() {}

// ─── Commands ────────────────────────────────────────────────────────────────

// Command is the discriminated union of all API commands.
type Command interface {
	// Syntax returns the wire string to send to the SimpleX chat controller.
	Syntax() string
}

// ── Address commands ──────────────────────────────────────────────────────────

type APICreateMyAddress struct {
	UserId int64 `json:"userId"`
}

type APIDeleteMyAddress struct {
	UserId int64 `json:"userId"`
}

type APIShowMyAddress struct {
	UserId int64 `json:"userId"`
}

type APISetProfileAddress struct {
	UserId int64 `json:"userId"`
	Enable bool  `json:"enable"`
}

type APISetAddressSettings struct {
	UserId   int64           `json:"userId"`
	Settings AddressSettings `json:"settings"`
}

func (c APICreateMyAddress) Syntax() string {
	return "/_address " + itoa(c.UserId)
}

func (c APIDeleteMyAddress) Syntax() string {
	return "/_delete_address " + itoa(c.UserId)
}

func (c APIShowMyAddress) Syntax() string {
	return "/_show_address " + itoa(c.UserId)
}

func (c APISetProfileAddress) Syntax() string {
	return "/_profile_address " + itoa(c.UserId) + " " + onOff(c.Enable)
}

func (c APISetAddressSettings) Syntax() string {
	return "/_address_settings " + itoa(c.UserId) + " " + mustJSON(c.Settings)
}

// ── Message commands ──────────────────────────────────────────────────────────

type APISendMessages struct {
	SendRef          ChatRef           `json:"sendRef"`
	LiveMessage      bool              `json:"liveMessage"`
	Ttl              *int              `json:"ttl,omitempty"`
	ComposedMessages []ComposedMessage `json:"composedMessages"`
}

type APIUpdateChatItem struct {
	ChatRef        ChatRef        `json:"chatRef"`
	ChatItemId     int64          `json:"chatItemId"`
	LiveMessage    bool           `json:"liveMessage"`
	UpdatedMessage UpdatedMessage `json:"updatedMessage"`
}

type APIDeleteChatItem struct {
	ChatRef     ChatRef      `json:"chatRef"`
	ChatItemIds []int64      `json:"chatItemIds"`
	DeleteMode  CIDeleteMode `json:"deleteMode"`
}

type APIDeleteMemberChatItem struct {
	GroupId     int64   `json:"groupId"`
	ChatItemIds []int64 `json:"chatItemIds"`
}

type APIChatItemReaction struct {
	ChatRef    ChatRef     `json:"chatRef"`
	ChatItemId int64       `json:"chatItemId"`
	Add        bool        `json:"add"`
	Reaction   MsgReaction `json:"reaction"`
}

func (c APISendMessages) Syntax() string {
	s := "/_send " + chatRefStr(c.SendRef)
	if c.LiveMessage {
		s += " live=on"
	}
	if c.Ttl != nil {
		s += " ttl=" + itoa(int64(*c.Ttl))
	}
	return s + " json " + mustJSON(c.ComposedMessages)
}

func (c APIUpdateChatItem) Syntax() string {
	s := "/_update item " + chatRefStr(c.ChatRef) + " " + itoa(c.ChatItemId)
	if c.LiveMessage {
		s += " live=on"
	}
	return s + " json " + mustJSON(c.UpdatedMessage)
}

func (c APIDeleteChatItem) Syntax() string {
	return "/_delete item " + chatRefStr(c.ChatRef) + " " + joinInt64s(c.ChatItemIds) + " " + string(c.DeleteMode)
}

func (c APIDeleteMemberChatItem) Syntax() string {
	return "/_delete member item #" + itoa(c.GroupId) + " " + joinInt64s(c.ChatItemIds)
}

func (c APIChatItemReaction) Syntax() string {
	return "/_reaction " + chatRefStr(c.ChatRef) + " " + itoa(c.ChatItemId) + " " + onOff(c.Add) + " " + mustJSON(c.Reaction)
}

// ── File commands ─────────────────────────────────────────────────────────────

type ReceiveFile struct {
	FileId             int64   `json:"fileId"`
	UserApprovedRelays bool    `json:"userApprovedRelays"`
	StoreEncrypted     *bool   `json:"storeEncrypted,omitempty"`
	FileInline         *bool   `json:"fileInline,omitempty"`
	FilePath           *string `json:"filePath,omitempty"`
}

type CancelFile struct {
	FileId int64 `json:"fileId"`
}

func (c ReceiveFile) Syntax() string {
	s := "/freceive " + itoa(c.FileId)
	if c.UserApprovedRelays {
		s += " approved_relays=on"
	}
	if c.StoreEncrypted != nil {
		s += " encrypt=" + onOff(*c.StoreEncrypted)
	}
	if c.FileInline != nil {
		s += " inline=" + onOff(*c.FileInline)
	}
	if c.FilePath != nil {
		s += " " + *c.FilePath
	}
	return s
}

func (c CancelFile) Syntax() string {
	return "/fcancel " + itoa(c.FileId)
}

// ── Group commands ────────────────────────────────────────────────────────────

type APIAddMember struct {
	GroupId    int64           `json:"groupId"`
	ContactId  int64           `json:"contactId"`
	MemberRole GroupMemberRole `json:"memberRole"`
}

type APIJoinGroup struct {
	GroupId int64 `json:"groupId"`
}

type APIAcceptMember struct {
	GroupId       int64           `json:"groupId"`
	GroupMemberId int64           `json:"groupMemberId"`
	MemberRole    GroupMemberRole `json:"memberRole"`
}

type APIMembersRole struct {
	GroupId        int64           `json:"groupId"`
	GroupMemberIds []int64         `json:"groupMemberIds"`
	MemberRole     GroupMemberRole `json:"memberRole"`
}

type APIBlockMembersForAll struct {
	GroupId        int64   `json:"groupId"`
	GroupMemberIds []int64 `json:"groupMemberIds"`
	Blocked        bool    `json:"blocked"`
}

type APIRemoveMembers struct {
	GroupId        int64   `json:"groupId"`
	GroupMemberIds []int64 `json:"groupMemberIds"`
	WithMessages   bool    `json:"withMessages"`
}

type APILeaveGroup struct {
	GroupId int64 `json:"groupId"`
}

type APIListMembers struct {
	GroupId int64 `json:"groupId"`
}

type APINewGroup struct {
	UserId       int64        `json:"userId"`
	Incognito    bool         `json:"incognito"`
	GroupProfile GroupProfile `json:"groupProfile"`
}

type APIUpdateGroupProfile struct {
	GroupId      int64        `json:"groupId"`
	GroupProfile GroupProfile `json:"groupProfile"`
}

func (c APIAddMember) Syntax() string {
	return "/_add #" + itoa(c.GroupId) + " " + itoa(c.ContactId) + " " + string(c.MemberRole)
}

func (c APIJoinGroup) Syntax() string {
	return "/_join #" + itoa(c.GroupId)
}

func (c APIAcceptMember) Syntax() string {
	return "/_accept member #" + itoa(c.GroupId) + " " + itoa(c.GroupMemberId) + " " + string(c.MemberRole)
}

func (c APIMembersRole) Syntax() string {
	return "/_member role #" + itoa(c.GroupId) + " " + joinInt64s(c.GroupMemberIds) + " " + string(c.MemberRole)
}

func (c APIBlockMembersForAll) Syntax() string {
	return "/_block #" + itoa(c.GroupId) + " " + joinInt64s(c.GroupMemberIds) + " blocked=" + onOff(c.Blocked)
}

func (c APIRemoveMembers) Syntax() string {
	s := "/_remove #" + itoa(c.GroupId) + " " + joinInt64s(c.GroupMemberIds)
	if c.WithMessages {
		s += " messages=on"
	}
	return s
}

func (c APILeaveGroup) Syntax() string {
	return "/_leave #" + itoa(c.GroupId)
}

func (c APIListMembers) Syntax() string {
	return "/_members #" + itoa(c.GroupId)
}

func (c APINewGroup) Syntax() string {
	s := "/_group " + itoa(c.UserId)
	if c.Incognito {
		s += " incognito=on"
	}
	return s + " " + mustJSON(c.GroupProfile)
}

func (c APIUpdateGroupProfile) Syntax() string {
	return "/_group_profile #" + itoa(c.GroupId) + " " + mustJSON(c.GroupProfile)
}

// ── Group link commands ───────────────────────────────────────────────────────

type APICreateGroupLink struct {
	GroupId    int64           `json:"groupId"`
	MemberRole GroupMemberRole `json:"memberRole"`
}

type APIGroupLinkMemberRole struct {
	GroupId    int64           `json:"groupId"`
	MemberRole GroupMemberRole `json:"memberRole"`
}

type APIDeleteGroupLink struct {
	GroupId int64 `json:"groupId"`
}

type APIGetGroupLink struct {
	GroupId int64 `json:"groupId"`
}

func (c APICreateGroupLink) Syntax() string {
	return "/_create link #" + itoa(c.GroupId) + " " + string(c.MemberRole)
}

func (c APIGroupLinkMemberRole) Syntax() string {
	return "/_set link role #" + itoa(c.GroupId) + " " + string(c.MemberRole)
}

func (c APIDeleteGroupLink) Syntax() string {
	return "/_delete link #" + itoa(c.GroupId)
}

func (c APIGetGroupLink) Syntax() string {
	return "/_get link #" + itoa(c.GroupId)
}

// ── Connection commands ───────────────────────────────────────────────────────

type APIAddContact struct {
	UserId    int64 `json:"userId"`
	Incognito bool  `json:"incognito"`
}

type APIConnectPlan struct {
	UserId         int64   `json:"userId"`
	ConnectionLink *string `json:"connectionLink,omitempty"`
}

type APIConnect struct {
	UserId       int64            `json:"userId"`
	Incognito    bool             `json:"incognito"`
	PreparedLink *CreatedConnLink `json:"preparedLink_,omitempty"`
}

type Connect struct {
	Incognito bool    `json:"incognito"`
	ConnLink_ *string `json:"connLink_,omitempty"`
}

type APIAcceptContact struct {
	ContactReqId int64 `json:"contactReqId"`
}

type APIRejectContact struct {
	ContactReqId int64 `json:"contactReqId"`
}

func (c APIAddContact) Syntax() string {
	s := "/_connect " + itoa(c.UserId)
	if c.Incognito {
		s += " incognito=on"
	}
	return s
}

func (c APIConnectPlan) Syntax() string {
	s := "/_connect plan " + itoa(c.UserId)
	if c.ConnectionLink != nil {
		s += " " + *c.ConnectionLink
	}
	return s
}

func (c APIConnect) Syntax() string {
	s := "/_connect " + itoa(c.UserId)
	if c.PreparedLink != nil {
		s += " " + connLinkStr(*c.PreparedLink)
	}
	return s
}

func (c Connect) Syntax() string {
	s := "/connect"
	if c.ConnLink_ != nil {
		s += " " + *c.ConnLink_
	}
	return s
}

func (c APIAcceptContact) Syntax() string {
	return "/_accept " + itoa(c.ContactReqId)
}

func (c APIRejectContact) Syntax() string {
	return "/_reject " + itoa(c.ContactReqId)
}

// ── Chat commands ─────────────────────────────────────────────────────────────

type APIListContacts struct {
	UserId int64 `json:"userId"`
}

type APIListGroups struct {
	UserId     int64   `json:"userId"`
	ContactId_ *int64  `json:"contactId_,omitempty"`
	Search     *string `json:"search,omitempty"`
}

type APIDeleteChat struct {
	ChatRef        ChatRef        `json:"chatRef"`
	ChatDeleteMode ChatDeleteMode `json:"chatDeleteMode"`
}

func (c APIListContacts) Syntax() string {
	return "/_contacts " + itoa(c.UserId)
}

func (c APIListGroups) Syntax() string {
	s := "/_groups " + itoa(c.UserId)
	if c.ContactId_ != nil {
		s += " @" + itoa(*c.ContactId_)
	}
	if c.Search != nil {
		s += " " + *c.Search
	}
	return s
}

func (c APIDeleteChat) Syntax() string {
	return "/_delete " + chatRefStr(c.ChatRef) + " " + chatDeleteModeStr(c.ChatDeleteMode)
}

// ── User profile commands ─────────────────────────────────────────────────────

type ShowActiveUser struct{}

type CreateActiveUser struct {
	NewUser NewUser `json:"newUser"`
}

type ListUsers struct{}

type APISetActiveUser struct {
	UserId  int64   `json:"userId"`
	ViewPwd *string `json:"viewPwd,omitempty"`
}

type APIDeleteUser struct {
	UserId       int64   `json:"userId"`
	DelSMPQueues bool    `json:"delSMPQueues"`
	ViewPwd      *string `json:"viewPwd,omitempty"`
}

type APIUpdateProfile struct {
	UserId  int64   `json:"userId"`
	Profile Profile `json:"profile"`
}

type APISetContactPrefs struct {
	ContactId   int64       `json:"contactId"`
	Preferences Preferences `json:"preferences"`
}

func (ShowActiveUser) Syntax() string { return "/user" }

func (c CreateActiveUser) Syntax() string {
	return "/_create user " + mustJSON(c.NewUser)
}

func (ListUsers) Syntax() string { return "/users" }

func (c APISetActiveUser) Syntax() string {
	s := "/_user " + itoa(c.UserId)
	if c.ViewPwd != nil {
		s += " " + mustJSON(*c.ViewPwd)
	}
	return s
}

func (c APIDeleteUser) Syntax() string {
	s := "/_delete user " + itoa(c.UserId) + " del_smp=" + onOff(c.DelSMPQueues)
	if c.ViewPwd != nil {
		s += " " + mustJSON(*c.ViewPwd)
	}
	return s
}

func (c APIUpdateProfile) Syntax() string {
	return "/_profile " + itoa(c.UserId) + " " + mustJSON(c.Profile)
}

func (c APISetContactPrefs) Syntax() string {
	return "/_set prefs @" + itoa(c.ContactId) + " " + mustJSON(c.Preferences)
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func itoa(n int64) string {
	return strconv.FormatInt(n, 10)
}

func onOff(b bool) string {
	if b {
		return "on"
	}
	return "off"
}

func mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic("simplexgo: json.Marshal: " + err.Error())
	}
	return string(b)
}

func joinInt64s(ids []int64) string {
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = itoa(id)
	}
	return strings.Join(parts, ",")
}

// chatRefStr serialises a ChatRef to its wire form: <@|#|*><id>[<scope>]
func chatRefStr(r ChatRef) string {
	var prefix string
	switch r.ChatType {
	case ChatTypeDirect:
		prefix = "@"
	case ChatTypeGroup:
		prefix = "#"
	case ChatTypeLocal:
		prefix = "*"
	}
	s := prefix + itoa(r.ChatId)
	if r.ChatScope != nil {
		s += groupChatScopeStr(*r.ChatScope)
	}
	return s
}

func groupChatScopeStr(scope GroupChatScope) string {
	switch v := scope.(type) {
	case GroupChatScopeMemberSupport:
		if v.GroupMemberId_ != nil {
			return "(_support:" + itoa(*v.GroupMemberId_) + ")"
		}
		return "(_support)"
	}
	return ""
}

// connLinkStr serialises a CreatedConnLink to its wire form.
func connLinkStr(l CreatedConnLink) string {
	if l.ConnShortLink != nil {
		return l.ConnFullLink + " " + *l.ConnShortLink
	}
	return l.ConnFullLink
}

// chatDeleteModeStr serialises a ChatDeleteMode to its wire form.
func chatDeleteModeStr(m ChatDeleteMode) string {
	switch v := m.(type) {
	case ChatDeleteModeFull:
		if !v.Notify {
			return "full notify=off"
		}
		return "full"
	case ChatDeleteModeEntity:
		if !v.Notify {
			return "entity notify=off"
		}
		return "entity"
	case ChatDeleteModeMessages:
		return "messages"
	}
	return ""
}
