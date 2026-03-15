package simplexgo

import "time"

type ACIReaction struct {
	ChatInfo     ChatInfo   `json:"chatInfo"`
	ChatReaction CIReaction `json:"chatReaction"`
}

type AChat struct {
	ChatInfo  ChatInfo  `json:"chatInfo"`
	ChatItem  ChatItem  `json:"chatItem"`
	ChatStats ChatStats `json:"chatStats"`
}

type AChatItem struct {
	ChatInfo ChatInfo `json:"chatInfo"`
	ChatItem ChatItem `json:"chatItem"`
}

type AddressSettings struct {
	BusinessAddress bool        `json:"businessAddress"`
	AutoAccept      *AutoAccept `json:"autoAccept,omitempty"`
	AutoReply       *MsgContent `json:"autoReply,omitempty"`
}

type AgentCryptoError interface {
	agentCryptoError()
}

type AgentCryptoErrorDecryptAES struct{}
type AgentCryptoErrorDecryptCB struct{}
type AgentCryptoErrorRatchetHeader struct{}
type AgentCryptoErrorRatchetSync struct{}

func (AgentCryptoErrorDecryptAES) agentCryptoError()    {}
func (AgentCryptoErrorDecryptCB) agentCryptoError()     {}
func (AgentCryptoErrorRatchetHeader) agentCryptoError() {}
func (AgentCryptoErrorRatchetSync) agentCryptoError()   {}

type AgentErrorType interface {
	agentErrorType()
}

type AgentErrorTypeCMD struct {
	CmdErr     CommandErrorType `json:"cmdErr"`
	ErrContext string           `json:"errContext"`
}

type AgentErrorTypeCONN struct {
	ConnErr    ConnectionErrorType `json:"connErr"`
	ErrContext string              `json:"errContext"`
}

type AgentErrorTypeNO_USER struct{}

type AgentErrorTypeSMP struct {
	ServerAddress string    `json:"serverAddress"`
	SmpErr        ErrorType `json:"smpErr"`
}

type AgentErrorTypeNTF struct {
	ServerAddress string    `json:"serverAddress"`
	NtfErr        ErrorType `json:"ntfErr"`
}

type AgentErrorTypeXFTP struct {
	ServerAddress string        `json:"serverAddress"`
	XftpErr       XFTPErrorType `json:"xftpErr"`
}

type AgentErrorTypeFILE struct {
	FileErr FileErrorType `json:"fileErr"`
}

type AgentErrorTypePROXY struct {
	ProxyServer string           `json:"proxyServer"`
	RelayServer string           `json:"relayServer"`
	ProxyErr    ProxyClientError `json:"proxyErr"`
}

type AgentErrorTypeRCP struct {
	RcpErr RCErrorType `json:"rcpErr"`
}

type AgentErrorTypeBROKER struct {
	BrokerAddress string          `json:"brokerAddress"`
	BrokerErr     BrokerErrorType `json:"brokerErr"`
}

type AgentErrorTypeAGENT struct {
	AgentErr SMPAgentError `json:"agentErr"`
}

type AgentErrorTypeINTERNAL struct {
	InternalErr string `json:"internalErr"`
}

type AgentErrorTypeCRITICAL struct {
	OfferRestart bool   `json:"offerRestart"`
	CriticalErr  string `json:"criticalErr"`
}

type AgentErrorTypeINACTIVE struct{}

func (AgentErrorTypeCMD) agentErrorType()      {}
func (AgentErrorTypeCONN) agentErrorType()     {}
func (AgentErrorTypeNO_USER) agentErrorType()  {}
func (AgentErrorTypeSMP) agentErrorType()      {}
func (AgentErrorTypeNTF) agentErrorType()      {}
func (AgentErrorTypeXFTP) agentErrorType()     {}
func (AgentErrorTypeFILE) agentErrorType()     {}
func (AgentErrorTypePROXY) agentErrorType()    {}
func (AgentErrorTypeRCP) agentErrorType()      {}
func (AgentErrorTypeBROKER) agentErrorType()   {}
func (AgentErrorTypeAGENT) agentErrorType()    {}
func (AgentErrorTypeINTERNAL) agentErrorType() {}
func (AgentErrorTypeCRITICAL) agentErrorType() {}
func (AgentErrorTypeINACTIVE) agentErrorType() {}

type AutoAccept struct {
	AcceptIncognito bool `json:"acceptIncognito"`
}

type BlockingInfo struct {
	Reason BlockingReason `json:"reason"`
}

type BlockingReason string

const (
	BlockingReasonSpam    BlockingReason = "spam"
	BlockingReasonContent BlockingReason = "content"
)

type BrokerErrorType interface {
	brokerErrorType()
}

type BrokerErrorTypeRESPONSE struct {
	RespErr string `json:"respErr"`
}

type BrokerErrorTypeUNEXPECTED struct {
	RespErr string `json:"respErr"`
}

type BrokerErrorTypeNETWORK struct {
	NetworkError NetworkError `json:"networkError"`
}

type BrokerErrorTypeHOST struct{}

type BrokerErrorTypeNO_SERVICE struct{}

type BrokerErrorTypeTRANSPORT struct {
	TransportErr TransportError `json:"transportErr"`
}

type BrokerErrorTypeTIMEOUT struct{}

func (BrokerErrorTypeRESPONSE) brokerErrorType()   {}
func (BrokerErrorTypeUNEXPECTED) brokerErrorType() {}
func (BrokerErrorTypeNETWORK) brokerErrorType()    {}
func (BrokerErrorTypeHOST) brokerErrorType()       {}
func (BrokerErrorTypeNO_SERVICE) brokerErrorType() {}
func (BrokerErrorTypeTRANSPORT) brokerErrorType()  {}
func (BrokerErrorTypeTIMEOUT) brokerErrorType()    {}

type BusinessChatInfo struct {
	ChatType   BusinessChatType `json:"chatType"`
	BusinessId string           `json:"businessId"`
	CustomerId string           `json:"customerId"`
}

type BusinessChatType string

const (
	BusinessChatTypeBusiness BusinessChatType = "business"
	BusinessChatTypeCustomer BusinessChatType = "customer"
)

type CICallStatus string

const (
	CICallStatusPending    CICallStatus = "pending"
	CICallStatusMissed     CICallStatus = "missed"
	CICallStatusRejected   CICallStatus = "rejected"
	CICallStatusAccepted   CICallStatus = "accepted"
	CICallStatusNegotiated CICallStatus = "negotiated"
	CICallStatusProgress   CICallStatus = "progress"
	CICallStatusEnded      CICallStatus = "ended"
	CICallStatusError      CICallStatus = "error"
)

type CIContent interface {
	ciContentType()
}

type CIContentSndMsgContent struct {
	MsgContent MsgContent `json:"msgContent"`
}

type CIContentRcvMsgContent struct {
	MsgContent MsgContent `json:"msgContent"`
}

type CIContentSndDeleted struct {
	DeleteMode CIDeleteMode `json:"deleteMode"`
}

type CIContentRcvDeleted struct {
	DeleteMode CIDeleteMode `json:"deleteMode"`
}

type CIContentSndCall struct {
	Status   CICallStatus `json:"status"`
	Duration int          `json:"duration"`
}

type CIContentRcvCall struct {
	Status   CICallStatus `json:"status"`
	Duration int          `json:"duration"`
}

type CIContentRcvIntegrityError struct {
	MsgError MsgErrorType `json:"msgError"`
}

type CIContentRcvDecryptionError struct {
	MsgDecryptError MsgDecryptError `json:"msgDecryptError"`
	MsgCount        uint32          `json:"msgCount"`
}

type CIContentRcvGroupInvitation struct {
	GroupInvitation CIGroupInvitation `json:"groupInvitation"`
	MemberRole      GroupMemberRole   `json:"memberRole"`
}

type CIContentSndGroupInvitation struct {
	GroupInvitation CIGroupInvitation `json:"groupInvitation"`
	MemberRole      GroupMemberRole   `json:"memberRole"`
}

type CIContentRcvDirectEvent struct {
	RcvDirectEvent RcvDirectEvent `json:"rcvDirectEvent"`
}

type CIContentRcvGroupEvent struct {
	RcvGroupEvent RcvGroupEvent `json:"rcvGroupEvent"`
}

type CIContentSndGroupEvent struct {
	SndGroupEvent SndGroupEvent `json:"sndGroupEvent"`
}

type CIContentRcvConnEvent struct {
	RcvConnEvent RcvConnEvent `json:"rcvConnEvent"`
}

type CIContentSndConnEvent struct {
	SndConnEvent SndConnEvent `json:"sndConnEvent"`
}

type CIContentRcvChatFeature struct {
	Feature ChatFeature `json:"feature"`
	Enabled PrefEnabled `json:"enabled"`
	Param   *int        `json:"param,omitempty"`
}

type CIContentSndChatFeature struct {
	Feature ChatFeature `json:"feature"`
	Enabled PrefEnabled `json:"enabled"`
	Param   *int        `json:"param,omitempty"`
}

type CIContentRcvChatPreference struct {
	Feature ChatFeature    `json:"feature"`
	Allowed FeatureAllowed `json:"allowed"`
	Param   *int           `json:"param,omitempty"`
}

type CIContentSndChatPreference struct {
	Feature ChatFeature    `json:"feature"`
	Allowed FeatureAllowed `json:"allowed"`
	Param   *int           `json:"param,omitempty"`
}

type CIContentRcvGroupFeature struct {
	GroupFeature GroupFeature     `json:"groupFeature"`
	Preference   GroupPreference  `json:"preference"`
	Param        *int             `json:"param,omitempty"`
	MemberRole   *GroupMemberRole `json:"memberRole_,omitempty"`
}

type CIContentSndGroupFeature struct {
	GroupFeature GroupFeature     `json:"groupFeature"`
	Preference   GroupPreference  `json:"preference"`
	Param        *int             `json:"param,omitempty"`
	MemberRole   *GroupMemberRole `json:"memberRole_,omitempty"`
}

type CIContentRcvChatFeatureRejected struct {
	Feature ChatFeature `json:"feature"`
}

type CIContentRcvGroupFeatureRejected struct {
	GroupFeature GroupFeature `json:"groupFeature"`
}

type CIContentSndModerated struct{}
type CIContentRcvModerated struct{}
type CIContentRcvBlocked struct{}

type CIContentSndDirectE2EEInfo struct {
	E2EEInfo E2EInfo `json:"e2eeInfo"`
}

type CIContentRcvDirectE2EEInfo struct {
	E2EEInfo E2EInfo `json:"e2eeInfo"`
}

type CIContentSndGroupE2EEInfo struct {
	E2EEInfo E2EInfo `json:"e2eeInfo"`
}

type CIContentRcvGroupE2EEInfo struct {
	E2EEInfo E2EInfo `json:"e2eeInfo"`
}

type CIContentChatBanner struct{}

func (CIContentSndMsgContent) ciContentType()           {}
func (CIContentRcvMsgContent) ciContentType()           {}
func (CIContentSndDeleted) ciContentType()              {}
func (CIContentRcvDeleted) ciContentType()              {}
func (CIContentSndCall) ciContentType()                 {}
func (CIContentRcvCall) ciContentType()                 {}
func (CIContentRcvIntegrityError) ciContentType()       {}
func (CIContentRcvDecryptionError) ciContentType()      {}
func (CIContentRcvGroupInvitation) ciContentType()      {}
func (CIContentSndGroupInvitation) ciContentType()      {}
func (CIContentRcvDirectEvent) ciContentType()          {}
func (CIContentRcvGroupEvent) ciContentType()           {}
func (CIContentSndGroupEvent) ciContentType()           {}
func (CIContentRcvConnEvent) ciContentType()            {}
func (CIContentSndConnEvent) ciContentType()            {}
func (CIContentRcvChatFeature) ciContentType()          {}
func (CIContentSndChatFeature) ciContentType()          {}
func (CIContentRcvChatPreference) ciContentType()       {}
func (CIContentSndChatPreference) ciContentType()       {}
func (CIContentRcvGroupFeature) ciContentType()         {}
func (CIContentSndGroupFeature) ciContentType()         {}
func (CIContentRcvChatFeatureRejected) ciContentType()  {}
func (CIContentRcvGroupFeatureRejected) ciContentType() {}
func (CIContentSndModerated) ciContentType()            {}
func (CIContentRcvModerated) ciContentType()            {}
func (CIContentRcvBlocked) ciContentType()              {}
func (CIContentSndDirectE2EEInfo) ciContentType()       {}
func (CIContentRcvDirectE2EEInfo) ciContentType()       {}
func (CIContentSndGroupE2EEInfo) ciContentType()        {}
func (CIContentRcvGroupE2EEInfo) ciContentType()        {}
func (CIContentChatBanner) ciContentType()              {}

type CIDeleteMode string

const (
	CIDeleteModeBroadcast    CIDeleteMode = "broadcast"
	CIDeleteModeInternal     CIDeleteMode = "internal"
	CIDeleteModeInternalMark CIDeleteMode = "internalMark"
)

type CIDeleted interface {
	ciDeleted()
}

type CIDeletedDeleted struct {
	DeletedTs *time.Time `json:"deletedTs,omitempty"`
	ChatType  ChatType   `json:"chatType"`
}

type CIDeletedBlocked struct {
	DeletedTs *time.Time `json:"deletedTs,omitempty"`
}

type CIDeletedBlockedByAdmin struct {
	DeletedTs *time.Time `json:"deletedTs,omitempty"`
}

type CIDeletedModerated struct {
	DeletedTs     *time.Time  `json:"deletedTs,omitempty"`
	ByGroupMember GroupMember `json:"byGroupMember"`
}

func (CIDeletedDeleted) ciDeleted()        {}
func (CIDeletedBlocked) ciDeleted()        {}
func (CIDeletedBlockedByAdmin) ciDeleted() {}
func (CIDeletedModerated) ciDeleted()      {}

type CIDirection interface {
	ciDirection()
}

type CIDirectionDirectSnd struct{}
type CIDirectionDirectRcv struct{}
type CIDirectionGroupSnd struct{}

type CIDirectionGroupRcv struct {
	GroupMember GroupMember `json:"groupMember"`
}

type CIDirectionLocalSnd struct{}
type CIDirectionLocalRcv struct{}

func (CIDirectionDirectSnd) ciDirection() {}
func (CIDirectionDirectRcv) ciDirection() {}
func (CIDirectionGroupSnd) ciDirection()  {}
func (CIDirectionGroupRcv) ciDirection()  {}
func (CIDirectionLocalSnd) ciDirection()  {}
func (CIDirectionLocalRcv) ciDirection()  {}

type CIFile struct {
	FileId       int64        `json:"fileId"`
	FileName     string       `json:"fileName"`
	FileSize     int64        `json:"fileSize"`
	FileSource   *CryptoFile  `json:"fileSource,omitempty"`
	FileStatus   CIFileStatus `json:"fileStatus"`
	FileProtocol FileProtocol `json:"fileProtocol"`
}

type CIFileStatus interface {
	ciFileStatus()
}

type CIFileStatusSndStored struct{}

type CIFileStatusSndTransfer struct {
	SndProgress int64 `json:"sndProgress"`
	SndTotal    int64 `json:"sndTotal"`
}

type CIFileStatusSndCancelled struct{}
type CIFileStatusSndComplete struct{}

type CIFileStatusSndError struct {
	SndFileError FileError `json:"sndFileError"`
}

type CIFileStatusSndWarning struct {
	SndFileError FileError `json:"sndFileError"`
}

type CIFileStatusRcvInvitation struct{}
type CIFileStatusRcvAccepted struct{}

type CIFileStatusRcvTransfer struct {
	RcvProgress int64 `json:"rcvProgress"`
	RcvTotal    int64 `json:"rcvTotal"`
}

type CIFileStatusRcvAborted struct{}
type CIFileStatusRcvComplete struct{}
type CIFileStatusRcvCancelled struct{}

type CIFileStatusRcvError struct {
	RcvFileError FileError `json:"rcvFileError"`
}

type CIFileStatusRcvWarning struct {
	RcvFileError FileError `json:"rcvFileError"`
}

type CIFileStatusInvalid struct {
	Text string `json:"text"`
}

func (CIFileStatusSndStored) ciFileStatus()     {}
func (CIFileStatusSndTransfer) ciFileStatus()   {}
func (CIFileStatusSndCancelled) ciFileStatus()  {}
func (CIFileStatusSndComplete) ciFileStatus()   {}
func (CIFileStatusSndError) ciFileStatus()      {}
func (CIFileStatusSndWarning) ciFileStatus()    {}
func (CIFileStatusRcvInvitation) ciFileStatus() {}
func (CIFileStatusRcvAccepted) ciFileStatus()   {}
func (CIFileStatusRcvTransfer) ciFileStatus()   {}
func (CIFileStatusRcvAborted) ciFileStatus()    {}
func (CIFileStatusRcvComplete) ciFileStatus()   {}
func (CIFileStatusRcvCancelled) ciFileStatus()  {}
func (CIFileStatusRcvError) ciFileStatus()      {}
func (CIFileStatusRcvWarning) ciFileStatus()    {}
func (CIFileStatusInvalid) ciFileStatus()       {}

type CIForwardedFrom interface {
	ciForwardedFrom()
}

type CIForwardedFromUnknown struct{}

type CIForwardedFromContact struct {
	ChatName   string       `json:"chatName"`
	MsgDir     MsgDirection `json:"msgDir"`
	ContactId  *int64       `json:"contactId,omitempty"`
	ChatItemId *int64       `json:"chatItemId,omitempty"`
}

type CIForwardedFromGroup struct {
	ChatName   string       `json:"chatName"`
	MsgDir     MsgDirection `json:"msgDir"`
	GroupId    *int64       `json:"groupId,omitempty"`
	ChatItemId *int64       `json:"chatItemId,omitempty"`
}

func (CIForwardedFromUnknown) ciForwardedFrom() {}
func (CIForwardedFromContact) ciForwardedFrom() {}
func (CIForwardedFromGroup) ciForwardedFrom()   {}

type CIGroupInvitation struct {
	GroupId          int64                   `json:"groupId"`
	GroupMemberId    int64                   `json:"groupMemberId"`
	LocalDisplayName string                  `json:"localDisplayName"`
	GroupProfile     GroupProfile            `json:"groupProfile"`
	Status           CIGroupInvitationStatus `json:"status"`
}

type CIGroupInvitationStatus string

const (
	CIGroupInvitationStatusPending  CIGroupInvitationStatus = "pending"
	CIGroupInvitationStatusAccepted CIGroupInvitationStatus = "accepted"
	CIGroupInvitationStatusRejected CIGroupInvitationStatus = "rejected"
	CIGroupInvitationStatusExpired  CIGroupInvitationStatus = "expired"
)

type CIMention struct {
	MemberId  string           `json:"memberId"`
	MemberRef *CIMentionMember `json:"memberRef,omitempty"`
}

type CIMentionMember struct {
	GroupMemberId int64           `json:"groupMemberId"`
	DisplayName   string          `json:"displayName"`
	LocalAlias    *string         `json:"localAlias,omitempty"`
	MemberRole    GroupMemberRole `json:"memberRole"`
}

type CIMeta struct {
	ItemId            int64            `json:"itemId"`
	ItemTs            time.Time        `json:"itemTs"`
	ItemText          string           `json:"itemText"`
	ItemStatus        CIStatus         `json:"itemStatus"`
	SentViaProxy      *bool            `json:"sentViaProxy,omitempty"`
	ItemSharedMsgId   *string          `json:"itemSharedMsgId,omitempty"`
	ItemForwarded     *CIForwardedFrom `json:"itemForwarded,omitempty"`
	ItemDeleted       *CIDeleted       `json:"itemDeleted,omitempty"`
	ItemEdited        bool             `json:"itemEdited"`
	ItemTimed         *CITimed         `json:"itemTimed,omitempty"`
	ItemLive          *bool            `json:"itemLive,omitempty"`
	UserMention       bool             `json:"userMention"`
	Deletable         bool             `json:"deletable"`
	Editable          bool             `json:"editable"`
	ForwardedByMember *int64           `json:"forwardedByMember,omitempty"`
	ShowGroupAsSender bool             `json:"showGroupAsSender"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
}

type CIQuote struct {
	ChatDir       *CIDirection     `json:"chatDir,omitempty"`
	ItemId        *int64           `json:"itemId,omitempty"`
	SharedMsgId   *string          `json:"sharedMsgId,omitempty"`
	SentAt        time.Time        `json:"sentAt"`
	Content       MsgContent       `json:"content"`
	FormattedText *[]FormattedText `json:"formattedText,omitempty"`
}

type CIReaction struct {
	ChatDir  CIDirection `json:"chatDir"`
	ChatItem ChatItem    `json:"chatItem"`
	SentAt   time.Time   `json:"sentAt"`
	Reaction MsgReaction `json:"reaction"`
}

type CIReactionCount struct {
	Reaction     MsgReaction `json:"reaction"`
	UserReacted  bool        `json:"userReacted"`
	TotalReacted int         `json:"totalReacted"`
}

type CIStatus interface {
	ciStatus()
}

type CIStatusSndNew struct{}

type CIStatusSndSent struct {
	SndProgress SndCIStatusProgress `json:"sndProgress"`
}

type CIStatusSndRcvd struct {
	MsgRcptStatus MsgReceiptStatus    `json:"msgRcptStatus"`
	SndProgress   SndCIStatusProgress `json:"sndProgress"`
}

type CIStatusSndErrorAuth struct{}

type CIStatusSndError struct {
	AgentError SndError `json:"agentError"`
}

type CIStatusSndWarning struct {
	AgentError SndError `json:"agentError"`
}

type CIStatusRcvNew struct{}
type CIStatusRcvRead struct{}

type CIStatusInvalid struct {
	Text string `json:"text"`
}

func (CIStatusSndNew) ciStatus()       {}
func (CIStatusSndSent) ciStatus()      {}
func (CIStatusSndRcvd) ciStatus()      {}
func (CIStatusSndErrorAuth) ciStatus() {}
func (CIStatusSndError) ciStatus()     {}
func (CIStatusSndWarning) ciStatus()   {}
func (CIStatusRcvNew) ciStatus()       {}
func (CIStatusRcvRead) ciStatus()      {}
func (CIStatusInvalid) ciStatus()      {}

type CITimed struct {
	Ttl      int        `json:"ttl"`
	DeleteAt *time.Time `json:"deleteAt,omitempty"`
}

type ChatBotCommand interface {
	chatBotCommand()
}

type ChatBotCommandCommand struct {
	Keyword string  `json:"keyword"`
	Label   string  `json:"label"`
	Params  *string `json:"params,omitempty"`
}

type ChatBotCommandMenu struct {
	Label    string           `json:"label"`
	Commands []ChatBotCommand `json:"commands"`
}

func (ChatBotCommandCommand) chatBotCommand() {}
func (ChatBotCommandMenu) chatBotCommand()    {}

type ChatDeleteMode interface {
	chatDeleteMode()
}

type ChatDeleteModeFull struct {
	Notify bool `json:"notify"`
}

type ChatDeleteModeEntity struct {
	Notify bool `json:"notify"`
}

type ChatDeleteModeMessages struct{}

func (ChatDeleteModeFull) chatDeleteMode()     {}
func (ChatDeleteModeEntity) chatDeleteMode()   {}
func (ChatDeleteModeMessages) chatDeleteMode() {}

type ChatError interface {
	chatError()
}

type ChatErrorError struct {
	ErrorType ChatErrorType `json:"errorType"`
}

type ChatErrorErrorAgent struct {
	AgentError        AgentErrorType    `json:"agentError"`
	ConnectionEntity_ *ConnectionEntity `json:"connectionEntity_,omitempty"`
}

type ChatErrorErrorStore struct {
	StoreError StoreError `json:"storeError"`
}

func (ChatErrorError) chatError()      {}
func (ChatErrorErrorAgent) chatError() {}
func (ChatErrorErrorStore) chatError() {}

type ChatErrorType interface {
	chatErrorType()
}

type ChatErrorTypeNoActiveUser struct{}
type ChatErrorTypeNoConnectionUser struct {
	AgentConnId string `json:"agentConnId"`
}
type ChatErrorTypeNoSndFileUser struct {
	AgentSndFileId string `json:"agentSndFileId"`
}
type ChatErrorTypeNoRcvFileUser struct {
	AgentRcvFileId string `json:"agentRcvFileId"`
}
type ChatErrorTypeUserUnknown struct{}
type ChatErrorTypeActiveUserExists struct{}
type ChatErrorTypeUserExists struct {
	ContactName string `json:"contactName"`
}
type ChatErrorTypeDifferentActiveUser struct {
	CommandUserId int64 `json:"commandUserId"`
	ActiveUserId  int64 `json:"activeUserId"`
}
type ChatErrorTypeCantDeleteActiveUser struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeCantDeleteLastUser struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeCantHideLastUser struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeHiddenUserAlwaysMuted struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeEmptyUserPassword struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeUserAlreadyHidden struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeUserNotHidden struct {
	UserId int64 `json:"userId"`
}
type ChatErrorTypeInvalidDisplayName struct {
	DisplayName string `json:"displayName"`
	ValidName   string `json:"validName"`
}
type ChatErrorTypeChatNotStarted struct{}
type ChatErrorTypeChatNotStopped struct{}
type ChatErrorTypeChatStoreChanged struct{}
type ChatErrorTypeInvalidConnReq struct{}
type ChatErrorTypeUnsupportedConnReq struct{}
type ChatErrorTypeConnReqMessageProhibited struct{}
type ChatErrorTypeContactNotReady struct {
	Contact Contact `json:"contact"`
}
type ChatErrorTypeContactNotActive struct {
	Contact Contact `json:"contact"`
}
type ChatErrorTypeContactDisabled struct {
	Contact Contact `json:"contact"`
}
type ChatErrorTypeConnectionDisabled struct {
	Connection Connection `json:"connection"`
}
type ChatErrorTypeGroupUserRole struct {
	GroupInfo    GroupInfo       `json:"groupInfo"`
	RequiredRole GroupMemberRole `json:"requiredRole"`
}
type ChatErrorTypeGroupMemberInitialRole struct {
	GroupInfo   GroupInfo       `json:"groupInfo"`
	InitialRole GroupMemberRole `json:"initialRole"`
}
type ChatErrorTypeContactIncognitoCantInvite struct{}
type ChatErrorTypeGroupIncognitoCantInvite struct{}
type ChatErrorTypeGroupContactRole struct {
	ContactName string `json:"contactName"`
}
type ChatErrorTypeGroupDuplicateMember struct {
	ContactName string `json:"contactName"`
}
type ChatErrorTypeGroupDuplicateMemberId struct{}
type ChatErrorTypeGroupNotJoined struct {
	GroupInfo GroupInfo `json:"groupInfo"`
}
type ChatErrorTypeGroupMemberNotActive struct{}
type ChatErrorTypeCantBlockMemberForSelf struct {
	GroupInfo       GroupInfo   `json:"groupInfo"`
	Member          GroupMember `json:"member"`
	SetShowMessages bool        `json:"setShowMessages"`
}
type ChatErrorTypeGroupMemberUserRemoved struct{}
type ChatErrorTypeGroupMemberNotFound struct{}
type ChatErrorTypeGroupCantResendInvitation struct {
	GroupInfo   GroupInfo `json:"groupInfo"`
	ContactName string    `json:"contactName"`
}
type ChatErrorTypeGroupInternal struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileNotFound struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileSize struct {
	FilePath string `json:"filePath"`
}
type ChatErrorTypeFileAlreadyReceiving struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileCancelled struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileCancel struct {
	FileId  int64  `json:"fileId"`
	Message string `json:"message"`
}
type ChatErrorTypeFileAlreadyExists struct {
	FilePath string `json:"filePath"`
}
type ChatErrorTypeFileRead struct {
	FilePath string `json:"filePath"`
	Message  string `json:"message"`
}
type ChatErrorTypeFileWrite struct {
	FilePath string `json:"filePath"`
	Message  string `json:"message"`
}
type ChatErrorTypeFileSend struct {
	FileId     int64          `json:"fileId"`
	AgentError AgentErrorType `json:"agentError"`
}
type ChatErrorTypeFileRcvChunk struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileInternal struct {
	Message string `json:"message"`
}
type ChatErrorTypeFileImageType struct {
	FilePath string `json:"filePath"`
}
type ChatErrorTypeFileImageSize struct {
	FilePath string `json:"filePath"`
}
type ChatErrorTypeFileNotReceived struct {
	FileId int64 `json:"fileId"`
}
type ChatErrorTypeFileNotApproved struct {
	FileId         int64    `json:"fileId"`
	UnknownServers []string `json:"unknownServers"`
}
type ChatErrorTypeFallbackToSMPProhibited struct {
	FileId int64 `json:"fileId"`
}
type ChatErrorTypeInlineFileProhibited struct {
	FileId int64 `json:"fileId"`
}
type ChatErrorTypeInvalidForward struct{}
type ChatErrorTypeInvalidChatItemUpdate struct{}
type ChatErrorTypeInvalidChatItemDelete struct{}
type ChatErrorTypeHasCurrentCall struct{}
type ChatErrorTypeNoCurrentCall struct{}
type ChatErrorTypeCallContact struct {
	ContactId int64 `json:"contactId"`
}
type ChatErrorTypeDirectMessagesProhibited struct {
	Direction MsgDirection `json:"direction"`
	Contact   Contact      `json:"contact"`
}
type ChatErrorTypeAgentVersion struct{}
type ChatErrorTypeAgentNoSubResult struct {
	AgentConnId string `json:"agentConnId"`
}
type ChatErrorTypeCommandError struct {
	Message string `json:"message"`
}
type ChatErrorTypeAgentCommandError struct {
	Message string `json:"message"`
}
type ChatErrorTypeInvalidFileDescription struct {
	Message string `json:"message"`
}
type ChatErrorTypeConnectionIncognitoChangeProhibited struct{}
type ChatErrorTypeConnectionUserChangeProhibited struct{}
type ChatErrorTypePeerChatVRangeIncompatible struct{}
type ChatErrorTypeInternalError struct {
	Message string `json:"message"`
}
type ChatErrorTypeException struct {
	Message string `json:"message"`
}

func (ChatErrorTypeNoActiveUser) chatErrorType()                        {}
func (ChatErrorTypeNoConnectionUser) chatErrorType()                    {}
func (ChatErrorTypeNoSndFileUser) chatErrorType()                       {}
func (ChatErrorTypeNoRcvFileUser) chatErrorType()                       {}
func (ChatErrorTypeUserUnknown) chatErrorType()                         {}
func (ChatErrorTypeActiveUserExists) chatErrorType()                    {}
func (ChatErrorTypeUserExists) chatErrorType()                          {}
func (ChatErrorTypeDifferentActiveUser) chatErrorType()                 {}
func (ChatErrorTypeCantDeleteActiveUser) chatErrorType()                {}
func (ChatErrorTypeCantDeleteLastUser) chatErrorType()                  {}
func (ChatErrorTypeCantHideLastUser) chatErrorType()                    {}
func (ChatErrorTypeHiddenUserAlwaysMuted) chatErrorType()               {}
func (ChatErrorTypeEmptyUserPassword) chatErrorType()                   {}
func (ChatErrorTypeUserAlreadyHidden) chatErrorType()                   {}
func (ChatErrorTypeUserNotHidden) chatErrorType()                       {}
func (ChatErrorTypeInvalidDisplayName) chatErrorType()                  {}
func (ChatErrorTypeChatNotStarted) chatErrorType()                      {}
func (ChatErrorTypeChatNotStopped) chatErrorType()                      {}
func (ChatErrorTypeChatStoreChanged) chatErrorType()                    {}
func (ChatErrorTypeInvalidConnReq) chatErrorType()                      {}
func (ChatErrorTypeUnsupportedConnReq) chatErrorType()                  {}
func (ChatErrorTypeConnReqMessageProhibited) chatErrorType()            {}
func (ChatErrorTypeContactNotReady) chatErrorType()                     {}
func (ChatErrorTypeContactNotActive) chatErrorType()                    {}
func (ChatErrorTypeContactDisabled) chatErrorType()                     {}
func (ChatErrorTypeConnectionDisabled) chatErrorType()                  {}
func (ChatErrorTypeGroupUserRole) chatErrorType()                       {}
func (ChatErrorTypeGroupMemberInitialRole) chatErrorType()              {}
func (ChatErrorTypeContactIncognitoCantInvite) chatErrorType()          {}
func (ChatErrorTypeGroupIncognitoCantInvite) chatErrorType()            {}
func (ChatErrorTypeGroupContactRole) chatErrorType()                    {}
func (ChatErrorTypeGroupDuplicateMember) chatErrorType()                {}
func (ChatErrorTypeGroupDuplicateMemberId) chatErrorType()              {}
func (ChatErrorTypeGroupNotJoined) chatErrorType()                      {}
func (ChatErrorTypeGroupMemberNotActive) chatErrorType()                {}
func (ChatErrorTypeCantBlockMemberForSelf) chatErrorType()              {}
func (ChatErrorTypeGroupMemberUserRemoved) chatErrorType()              {}
func (ChatErrorTypeGroupMemberNotFound) chatErrorType()                 {}
func (ChatErrorTypeGroupCantResendInvitation) chatErrorType()           {}
func (ChatErrorTypeGroupInternal) chatErrorType()                       {}
func (ChatErrorTypeFileNotFound) chatErrorType()                        {}
func (ChatErrorTypeFileSize) chatErrorType()                            {}
func (ChatErrorTypeFileAlreadyReceiving) chatErrorType()                {}
func (ChatErrorTypeFileCancelled) chatErrorType()                       {}
func (ChatErrorTypeFileCancel) chatErrorType()                          {}
func (ChatErrorTypeFileAlreadyExists) chatErrorType()                   {}
func (ChatErrorTypeFileRead) chatErrorType()                            {}
func (ChatErrorTypeFileWrite) chatErrorType()                           {}
func (ChatErrorTypeFileSend) chatErrorType()                            {}
func (ChatErrorTypeFileRcvChunk) chatErrorType()                        {}
func (ChatErrorTypeFileInternal) chatErrorType()                        {}
func (ChatErrorTypeFileImageType) chatErrorType()                       {}
func (ChatErrorTypeFileImageSize) chatErrorType()                       {}
func (ChatErrorTypeFileNotReceived) chatErrorType()                     {}
func (ChatErrorTypeFileNotApproved) chatErrorType()                     {}
func (ChatErrorTypeFallbackToSMPProhibited) chatErrorType()             {}
func (ChatErrorTypeInlineFileProhibited) chatErrorType()                {}
func (ChatErrorTypeInvalidForward) chatErrorType()                      {}
func (ChatErrorTypeInvalidChatItemUpdate) chatErrorType()               {}
func (ChatErrorTypeInvalidChatItemDelete) chatErrorType()               {}
func (ChatErrorTypeHasCurrentCall) chatErrorType()                      {}
func (ChatErrorTypeNoCurrentCall) chatErrorType()                       {}
func (ChatErrorTypeCallContact) chatErrorType()                         {}
func (ChatErrorTypeDirectMessagesProhibited) chatErrorType()            {}
func (ChatErrorTypeAgentVersion) chatErrorType()                        {}
func (ChatErrorTypeAgentNoSubResult) chatErrorType()                    {}
func (ChatErrorTypeCommandError) chatErrorType()                        {}
func (ChatErrorTypeAgentCommandError) chatErrorType()                   {}
func (ChatErrorTypeInvalidFileDescription) chatErrorType()              {}
func (ChatErrorTypeConnectionIncognitoChangeProhibited) chatErrorType() {}
func (ChatErrorTypeConnectionUserChangeProhibited) chatErrorType()      {}
func (ChatErrorTypePeerChatVRangeIncompatible) chatErrorType()          {}
func (ChatErrorTypeInternalError) chatErrorType()                       {}
func (ChatErrorTypeException) chatErrorType()                           {}

type ChatFeature string

const (
	ChatFeatureTimedMessages ChatFeature = "timedMessages"
	ChatFeatureFullDelete    ChatFeature = "fullDelete"
	ChatFeatureReactions     ChatFeature = "reactions"
	ChatFeatureVoice         ChatFeature = "voice"
	ChatFeatureFiles         ChatFeature = "files"
	ChatFeatureCalls         ChatFeature = "calls"
	ChatFeatureSessions      ChatFeature = "sessions"
)

type ChatInfo interface {
	chatInfo()
}

type ChatInfoDirect struct {
	Contact Contact `json:"contact"`
}

type ChatInfoGroup struct {
	GroupInfo      GroupInfo           `json:"groupInfo"`
	GroupChatScope *GroupChatScopeInfo `json:"groupChatScope,omitempty"`
}

type ChatInfoLocal struct {
	NoteFolder NoteFolder `json:"noteFolder"`
}

type ChatInfoContactRequest struct {
	ContactRequest UserContactRequest `json:"contactRequest"`
}

type ChatInfoContactConnection struct {
	ContactConnection PendingContactConnection `json:"contactConnection"`
}

func (ChatInfoDirect) chatInfo()            {}
func (ChatInfoGroup) chatInfo()             {}
func (ChatInfoLocal) chatInfo()             {}
func (ChatInfoContactRequest) chatInfo()    {}
func (ChatInfoContactConnection) chatInfo() {}

type ChatItem struct {
	ChatDir       CIDirection          `json:"chatDir"`
	Meta          CIMeta               `json:"meta"`
	Content       CIContent            `json:"content"`
	Mentions      map[string]CIMention `json:"mentions"`
	FormattedText *[]FormattedText     `json:"formattedText,omitempty"`
	QuotedItem    *CIQuote             `json:"quotedItem,omitempty"`
	Reactions     []CIReactionCount    `json:"reactions"`
	File          *CIFile              `json:"file,omitempty"`
}

type ChatItemDeletion struct {
	DeletedChatItem AChatItem  `json:"deletedChatItem"`
	ToChatItem      *AChatItem `json:"toChatItem,omitempty"`
}

type ChatPeerType string

const (
	ChatPeerTypeHuman ChatPeerType = "human"
	ChatPeerTypeBot   ChatPeerType = "bot"
)

type ChatRef struct {
	ChatType  ChatType        `json:"chatType"`
	ChatId    int64           `json:"chatId"`
	ChatScope *GroupChatScope `json:"chatScope,omitempty"`
}

type ChatSettings struct {
	EnableNtfs MsgFilter `json:"enableNtfs"`
	SendRcpts  *bool     `json:"sendRcpts,omitempty"`
	Favorite   bool      `json:"favorite"`
}

type ChatStats struct {
	UnreadCount     int   `json:"unreadCount"`
	UnreadMentions  int   `json:"unreadMentions"`
	ReportsCount    int   `json:"reportsCount"`
	MinUnreadItemId int64 `json:"minUnreadItemId"`
	UnreadChat      bool  `json:"unreadChat"`
}

type ChatType string

const (
	ChatTypeDirect ChatType = "direct"
	ChatTypeGroup  ChatType = "group"
	ChatTypeLocal  ChatType = "local"
)

type ChatWallpaper struct {
	Preset     *string             `json:"preset,omitempty"`
	ImageFile  *string             `json:"imageFile,omitempty"`
	Background *string             `json:"background,omitempty"`
	Tint       *string             `json:"tint,omitempty"`
	ScaleType  *ChatWallpaperScale `json:"scaleType,omitempty"`
	Scale      *float64            `json:"scale,omitempty"`
}

type ChatWallpaperScale string

const (
	ChatWallpaperScaleFill   ChatWallpaperScale = "fill"
	ChatWallpaperScaleFit    ChatWallpaperScale = "fit"
	ChatWallpaperScaleRepeat ChatWallpaperScale = "repeat"
)

type Color string

const (
	ColorBlack   Color = "black"
	ColorRed     Color = "red"
	ColorGreen   Color = "green"
	ColorYellow  Color = "yellow"
	ColorBlue    Color = "blue"
	ColorMagenta Color = "magenta"
	ColorCyan    Color = "cyan"
	ColorWhite   Color = "white"
)

type CommandError interface {
	commandError()
}

type CommandErrorUNKNOWN struct{}
type CommandErrorSYNTAX struct{}
type CommandErrorPROHIBITED struct{}
type CommandErrorNO_AUTH struct{}
type CommandErrorHAS_AUTH struct{}
type CommandErrorNO_ENTITY struct{}

func (CommandErrorUNKNOWN) commandError()    {}
func (CommandErrorSYNTAX) commandError()     {}
func (CommandErrorPROHIBITED) commandError() {}
func (CommandErrorNO_AUTH) commandError()    {}
func (CommandErrorHAS_AUTH) commandError()   {}
func (CommandErrorNO_ENTITY) commandError()  {}

type CommandErrorType interface {
	commandErrorType()
}

type CommandErrorTypePROHIBITED struct{}
type CommandErrorTypeSYNTAX struct{}
type CommandErrorTypeNO_CONN struct{}
type CommandErrorTypeSIZE struct{}
type CommandErrorTypeLARGE struct{}

func (CommandErrorTypePROHIBITED) commandErrorType() {}
func (CommandErrorTypeSYNTAX) commandErrorType()     {}
func (CommandErrorTypeNO_CONN) commandErrorType()    {}
func (CommandErrorTypeSIZE) commandErrorType()       {}
func (CommandErrorTypeLARGE) commandErrorType()      {}

type ComposedMessage struct {
	FileSource   *CryptoFile      `json:"fileSource,omitempty"`
	QuotedItemId *int64           `json:"quotedItemId,omitempty"`
	MsgContent   MsgContent       `json:"msgContent"`
	Mentions     map[string]int64 `json:"mentions"`
}

type ConnStatus string

const (
	ConnStatusNew       ConnStatus = "new"
	ConnStatusPrepared  ConnStatus = "prepared"
	ConnStatusJoined    ConnStatus = "joined"
	ConnStatusRequested ConnStatus = "requested"
	ConnStatusAccepted  ConnStatus = "accepted"
	ConnStatusSndReady  ConnStatus = "snd-ready"
	ConnStatusReady     ConnStatus = "ready"
	ConnStatusDeleted   ConnStatus = "deleted"
)

type ConnType string

const (
	ConnTypeContact     ConnType = "contact"
	ConnTypeMember      ConnType = "member"
	ConnTypeUserContact ConnType = "user_contact"
)

type Connection struct {
	ConnId               int64         `json:"connId"`
	AgentConnId          string        `json:"agentConnId"`
	ConnChatVersion      int           `json:"connChatVersion"`
	PeerChatVRange       VersionRange  `json:"peerChatVRange"`
	ConnLevel            int           `json:"connLevel"`
	ViaContact           *int64        `json:"viaContact,omitempty"`
	ViaUserContactLink   *int64        `json:"viaUserContactLink,omitempty"`
	ViaGroupLink         bool          `json:"viaGroupLink"`
	GroupLinkId          *string       `json:"groupLinkId,omitempty"`
	XContactId           *string       `json:"xContactId,omitempty"`
	CustomUserProfileId  *int64        `json:"customUserProfileId,omitempty"`
	ConnType             ConnType      `json:"connType"`
	ConnStatus           ConnStatus    `json:"connStatus"`
	ContactConnInitiated bool          `json:"contactConnInitiated"`
	LocalAlias           string        `json:"localAlias"`
	EntityId             *int64        `json:"entityId,omitempty"`
	ConnectionCode       *SecurityCode `json:"connectionCode,omitempty"`
	PqSupport            bool          `json:"pqSupport"`
	PqEncryption         bool          `json:"pqEncryption"`
	PqSndEnabled         *bool         `json:"pqSndEnabled,omitempty"`
	PqRcvEnabled         *bool         `json:"pqRcvEnabled,omitempty"`
	AuthErrCounter       int           `json:"authErrCounter"`
	QuotaErrCounter      int           `json:"quotaErrCounter"`
	CreatedAt            time.Time     `json:"createdAt"`
}

type ConnectionEntity interface {
	connectionEntity()
}

type ConnectionEntityRcvDirectMsgConnection struct {
	EntityConnection Connection `json:"entityConnection"`
	Contact          *Contact   `json:"contact,omitempty"`
}

type ConnectionEntityRcvGroupMsgConnection struct {
	EntityConnection Connection  `json:"entityConnection"`
	GroupInfo        GroupInfo   `json:"groupInfo"`
	GroupMember      GroupMember `json:"groupMember"`
}

type ConnectionEntitySndFileConnection struct {
	EntityConnection Connection      `json:"entityConnection"`
	SndFileTransfer  SndFileTransfer `json:"sndFileTransfer"`
}

type ConnectionEntityRcvFileConnection struct {
	EntityConnection Connection      `json:"entityConnection"`
	RcvFileTransfer  RcvFileTransfer `json:"rcvFileTransfer"`
}

type ConnectionEntityUserContactConnection struct {
	EntityConnection Connection  `json:"entityConnection"`
	UserContact      UserContact `json:"userContact"`
}

func (ConnectionEntityRcvDirectMsgConnection) connectionEntity() {}
func (ConnectionEntityRcvGroupMsgConnection) connectionEntity()  {}
func (ConnectionEntitySndFileConnection) connectionEntity()      {}
func (ConnectionEntityRcvFileConnection) connectionEntity()      {}
func (ConnectionEntityUserContactConnection) connectionEntity()  {}

type ConnectionErrorType interface {
	connectionErrorType()
}

type ConnectionErrorTypeNOT_FOUND struct{}
type ConnectionErrorTypeDUPLICATE struct{}
type ConnectionErrorTypeSIMPLEX struct{}
type ConnectionErrorTypeNOT_ACCEPTED struct{}
type ConnectionErrorTypeNOT_AVAILABLE struct{}

func (ConnectionErrorTypeNOT_FOUND) connectionErrorType()     {}
func (ConnectionErrorTypeDUPLICATE) connectionErrorType()     {}
func (ConnectionErrorTypeSIMPLEX) connectionErrorType()       {}
func (ConnectionErrorTypeNOT_ACCEPTED) connectionErrorType()  {}
func (ConnectionErrorTypeNOT_AVAILABLE) connectionErrorType() {}

type ConnectionMode string

const (
	ConnectionModeInv ConnectionMode = "inv"
	ConnectionModeCon ConnectionMode = "con"
)

type ConnectionPlan interface {
	connectionPlan()
}

type ConnectionPlanInvitationLink struct {
	InvitationLinkPlan InvitationLinkPlan `json:"invitationLinkPlan"`
}

type ConnectionPlanContactAddress struct {
	ContactAddressPlan ContactAddressPlan `json:"contactAddressPlan"`
}

type ConnectionPlanGroupLink struct {
	GroupLinkPlan GroupLinkPlan `json:"groupLinkPlan"`
}

type ConnectionPlanError struct {
	ChatError ChatError `json:"chatError"`
}

func (ConnectionPlanInvitationLink) connectionPlan() {}
func (ConnectionPlanContactAddress) connectionPlan() {}
func (ConnectionPlanGroupLink) connectionPlan()      {}
func (ConnectionPlanError) connectionPlan()          {}

type Contact struct {
	ContactId            int64                   `json:"contactId"`
	LocalDisplayName     string                  `json:"localDisplayName"`
	Profile              LocalProfile            `json:"profile"`
	ActiveConn           *Connection             `json:"activeConn,omitempty"`
	ViaGroup             *int64                  `json:"viaGroup,omitempty"`
	ContactUsed          bool                    `json:"contactUsed"`
	ContactStatus        ContactStatus           `json:"contactStatus"`
	ChatSettings         ChatSettings            `json:"chatSettings"`
	UserPreferences      Preferences             `json:"userPreferences"`
	MergedPreferences    ContactUserPreferences  `json:"mergedPreferences"`
	CreatedAt            time.Time               `json:"createdAt"`
	UpdatedAt            time.Time               `json:"updatedAt"`
	ChatTs               *time.Time              `json:"chatTs,omitempty"`
	PreparedContact      *PreparedContact        `json:"preparedContact,omitempty"`
	ContactRequestId     *int64                  `json:"contactRequestId,omitempty"`
	ContactGroupMemberId *int64                  `json:"contactGroupMemberId,omitempty"`
	ContactGrpInvSent    bool                    `json:"contactGrpInvSent"`
	GroupDirectInv       *GroupDirectInvitation  `json:"groupDirectInv,omitempty"`
	ChatTags             []int64                 `json:"chatTags"`
	ChatItemTTL          *int64                  `json:"chatItemTTL,omitempty"`
	UiThemes             *UIThemeEntityOverrides `json:"uiThemes,omitempty"`
	ChatDeleted          bool                    `json:"chatDeleted"`
	CustomData           map[string]interface{}  `json:"customData,omitempty"`
}

type ContactAddressPlan interface {
	contactAddressPlan()
}

type ContactAddressPlanOk struct {
	ContactSLinkData_ *ContactShortLinkData `json:"contactSLinkData_,omitempty"`
}

type ContactAddressPlanOwnLink struct{}

type ContactAddressPlanConnectingConfirmReconnect struct{}

type ContactAddressPlanConnectingProhibit struct {
	Contact Contact `json:"contact"`
}

type ContactAddressPlanKnown struct {
	Contact Contact `json:"contact"`
}

type ContactAddressPlanContactViaAddress struct {
	Contact Contact `json:"contact"`
}

func (ContactAddressPlanOk) contactAddressPlan()                         {}
func (ContactAddressPlanOwnLink) contactAddressPlan()                    {}
func (ContactAddressPlanConnectingConfirmReconnect) contactAddressPlan() {}
func (ContactAddressPlanConnectingProhibit) contactAddressPlan()         {}
func (ContactAddressPlanKnown) contactAddressPlan()                      {}
func (ContactAddressPlanContactViaAddress) contactAddressPlan()          {}

type ContactShortLinkData struct {
	Profile  Profile     `json:"profile"`
	Message  *MsgContent `json:"message,omitempty"`
	Business bool        `json:"business"`
}

type ContactStatus string

const (
	ContactStatusActive        ContactStatus = "active"
	ContactStatusDeleted       ContactStatus = "deleted"
	ContactStatusDeletedByUser ContactStatus = "deletedByUser"
)

type ContactUserPref interface {
	contactUserPref()
}

type ContactUserPrefContact struct {
	Preference SimplePreference `json:"preference"`
}

type ContactUserPrefUser struct {
	Preference SimplePreference `json:"preference"`
}

func (ContactUserPrefContact) contactUserPref() {}
func (ContactUserPrefUser) contactUserPref()    {}

type ContactUserPreference struct {
	Enabled           PrefEnabled      `json:"enabled"`
	UserPreference    ContactUserPref  `json:"userPreference"`
	ContactPreference SimplePreference `json:"contactPreference"`
}

type ContactUserPreferences struct {
	TimedMessages ContactUserPreference `json:"timedMessages"`
	FullDelete    ContactUserPreference `json:"fullDelete"`
	Reactions     ContactUserPreference `json:"reactions"`
	Voice         ContactUserPreference `json:"voice"`
	Files         ContactUserPreference `json:"files"`
	Calls         ContactUserPreference `json:"calls"`
	Sessions      ContactUserPreference `json:"sessions"`
	Commands      *[]ChatBotCommand     `json:"commands,omitempty"`
}

type CreatedConnLink struct {
	ConnFullLink  string  `json:"connFullLink"`
	ConnShortLink *string `json:"connShortLink,omitempty"`
}

type CryptoFile struct {
	FilePath   string          `json:"filePath"`
	CryptoArgs *CryptoFileArgs `json:"cryptoArgs,omitempty"`
}

type CryptoFileArgs struct {
	FileKey   string `json:"fileKey"`
	FileNonce string `json:"fileNonce"`
}

type E2EInfo struct {
	PqEnabled *bool `json:"pqEnabled,omitempty"`
}

type ErrorType interface {
	errorType()
}

type ErrorTypeBLOCK struct{}
type ErrorTypeSESSION struct{}
type ErrorTypeCMD struct {
	CmdErr CommandError `json:"cmdErr"`
}
type ErrorTypePROXY struct {
	ProxyErr ProxyError `json:"proxyErr"`
}
type ErrorTypeAUTH struct{}
type ErrorTypeBLOCKED struct {
	BlockInfo BlockingInfo `json:"blockInfo"`
}
type ErrorTypeSERVICE struct{}
type ErrorTypeCRYPTO struct{}
type ErrorTypeQUOTA struct{}
type ErrorTypeSTORE struct {
	StoreErr string `json:"storeErr"`
}
type ErrorTypeNO_MSG struct{}
type ErrorTypeLARGE_MSG struct{}
type ErrorTypeEXPIRED struct{}
type ErrorTypeINTERNAL struct{}
type ErrorTypeDUPLICATE_ struct{}

func (ErrorTypeBLOCK) errorType()      {}
func (ErrorTypeSESSION) errorType()    {}
func (ErrorTypeCMD) errorType()        {}
func (ErrorTypePROXY) errorType()      {}
func (ErrorTypeAUTH) errorType()       {}
func (ErrorTypeBLOCKED) errorType()    {}
func (ErrorTypeSERVICE) errorType()    {}
func (ErrorTypeCRYPTO) errorType()     {}
func (ErrorTypeQUOTA) errorType()      {}
func (ErrorTypeSTORE) errorType()      {}
func (ErrorTypeNO_MSG) errorType()     {}
func (ErrorTypeLARGE_MSG) errorType()  {}
func (ErrorTypeEXPIRED) errorType()    {}
func (ErrorTypeINTERNAL) errorType()   {}
func (ErrorTypeDUPLICATE_) errorType() {}

type FeatureAllowed string

const (
	FeatureAllowedAlways FeatureAllowed = "always"
	FeatureAllowedYes    FeatureAllowed = "yes"
	FeatureAllowedNo     FeatureAllowed = "no"
)

type FileDescr struct {
	FileDescrText     string `json:"fileDescrText"`
	FileDescrPartNo   int    `json:"fileDescrPartNo"`
	FileDescrComplete bool   `json:"fileDescrComplete"`
}

type FileError interface {
	fileError()
}

type FileErrorAuth struct{}

type FileErrorBlocked struct {
	Server    string       `json:"server"`
	BlockInfo BlockingInfo `json:"blockInfo"`
}

type FileErrorNoFile struct{}

type FileErrorRelay struct {
	SrvError SrvError `json:"srvError"`
}

type FileErrorOther struct {
	FileError string `json:"fileError"`
}

func (FileErrorAuth) fileError()    {}
func (FileErrorBlocked) fileError() {}
func (FileErrorNoFile) fileError()  {}
func (FileErrorRelay) fileError()   {}
func (FileErrorOther) fileError()   {}

type FileErrorType interface {
	fileErrorType()
}

type FileErrorTypeNOT_APPROVED struct{}
type FileErrorTypeSIZE struct{}
type FileErrorTypeREDIRECT struct {
	RedirectError string `json:"redirectError"`
}
type FileErrorTypeFILE_IO struct {
	FileIOError string `json:"fileIOError"`
}
type FileErrorTypeNO_FILE struct{}

func (FileErrorTypeNOT_APPROVED) fileErrorType() {}
func (FileErrorTypeSIZE) fileErrorType()         {}
func (FileErrorTypeREDIRECT) fileErrorType()     {}
func (FileErrorTypeFILE_IO) fileErrorType()      {}
func (FileErrorTypeNO_FILE) fileErrorType()      {}

type FileInvitation struct {
	FileName    string          `json:"fileName"`
	FileSize    int64           `json:"fileSize"`
	FileDigest  *string         `json:"fileDigest,omitempty"`
	FileConnReq *string         `json:"fileConnReq,omitempty"`
	FileInline  *InlineFileMode `json:"fileInline,omitempty"`
	FileDescr   *FileDescr      `json:"fileDescr,omitempty"`
}

type FileProtocol string

const (
	FileProtocolSmp   FileProtocol = "smp"
	FileProtocolXftp  FileProtocol = "xftp"
	FileProtocolLocal FileProtocol = "local"
)

type FileStatus string

const (
	FileStatusNew       FileStatus = "new"
	FileStatusAccepted  FileStatus = "accepted"
	FileStatusConnected FileStatus = "connected"
	FileStatusComplete  FileStatus = "complete"
	FileStatusCancelled FileStatus = "cancelled"
)

type FileTransferMeta struct {
	FileId          int64           `json:"fileId"`
	XftpSndFile     *XFTPSndFile    `json:"xftpSndFile,omitempty"`
	XftpRedirectFor *int64          `json:"xftpRedirectFor,omitempty"`
	FileName        string          `json:"fileName"`
	FilePath        string          `json:"filePath"`
	FileSize        int64           `json:"fileSize"`
	FileInline      *InlineFileMode `json:"fileInline,omitempty"`
	ChunkSize       int64           `json:"chunkSize"`
	Cancelled       bool            `json:"cancelled"`
}

type Format interface {
	format()
}

type FormatBold struct{}
type FormatItalic struct{}
type FormatStrikeThrough struct{}
type FormatSnippet struct{}
type FormatSecret struct{}

type FormatColored struct {
	Color Color `json:"color"`
}

type FormatUri struct{}

type FormatHyperLink struct {
	ShowText *string `json:"showText,omitempty"`
	LinkUri  string  `json:"linkUri"`
}

type FormatSimplexLink struct {
	ShowText   *string         `json:"showText,omitempty"`
	LinkType   SimplexLinkType `json:"linkType"`
	SimplexUri string          `json:"simplexUri"`
	SmpHosts   []string        `json:"smpHosts"`
}

type FormatCommand struct {
	CommandStr string `json:"commandStr"`
}

type FormatMention struct {
	MemberName string `json:"memberName"`
}

type FormatEmail struct{}
type FormatPhone struct{}

func (FormatBold) format()          {}
func (FormatItalic) format()        {}
func (FormatStrikeThrough) format() {}
func (FormatSnippet) format()       {}
func (FormatSecret) format()        {}
func (FormatColored) format()       {}
func (FormatUri) format()           {}
func (FormatHyperLink) format()     {}
func (FormatSimplexLink) format()   {}
func (FormatCommand) format()       {}
func (FormatMention) format()       {}
func (FormatEmail) format()         {}
func (FormatPhone) format()         {}

type FormattedText struct {
	Format *Format `json:"format,omitempty"`
	Text   string  `json:"text"`
}

type FullGroupPreferences struct {
	TimedMessages  TimedMessagesGroupPreference `json:"timedMessages"`
	DirectMessages RoleGroupPreference          `json:"directMessages"`
	FullDelete     GroupPreference              `json:"fullDelete"`
	Reactions      GroupPreference              `json:"reactions"`
	Voice          RoleGroupPreference          `json:"voice"`
	Files          RoleGroupPreference          `json:"files"`
	SimplexLinks   RoleGroupPreference          `json:"simplexLinks"`
	Reports        GroupPreference              `json:"reports"`
	History        GroupPreference              `json:"history"`
	Sessions       RoleGroupPreference          `json:"sessions"`
	Commands       []ChatBotCommand             `json:"commands"`
}

type FullPreferences struct {
	TimedMessages TimedMessagesPreference `json:"timedMessages"`
	FullDelete    SimplePreference        `json:"fullDelete"`
	Reactions     SimplePreference        `json:"reactions"`
	Voice         SimplePreference        `json:"voice"`
	Files         SimplePreference        `json:"files"`
	Calls         SimplePreference        `json:"calls"`
	Sessions      SimplePreference        `json:"sessions"`
	Commands      *[]ChatBotCommand       `json:"commands,omitempty"`
}

type Group struct {
	GroupInfo GroupInfo     `json:"groupInfo"`
	Members   []GroupMember `json:"members"`
}

type GroupChatScope interface {
	groupChatScope()
}

type GroupChatScopeMemberSupport struct {
	GroupMemberId_ *int64 `json:"groupMemberId_,omitempty"`
}

func (GroupChatScopeMemberSupport) groupChatScope() {}

type GroupChatScopeInfo interface {
	groupChatScopeInfo()
}

type GroupChatScopeInfoMemberSupport struct {
	GroupMember_ *GroupMember `json:"groupMember_,omitempty"`
}

func (GroupChatScopeInfoMemberSupport) groupChatScopeInfo() {}

type GroupDirectInvitation struct {
	GroupDirectInvLink              string `json:"groupDirectInvLink"`
	FromGroupId_                    *int64 `json:"fromGroupId_,omitempty"`
	FromGroupMemberId_              *int64 `json:"fromGroupMemberId_,omitempty"`
	FromGroupMemberConnId_          *int64 `json:"fromGroupMemberConnId_,omitempty"`
	GroupDirectInvStartedConnection bool   `json:"groupDirectInvStartedConnection"`
}

type GroupFeature string

const (
	GroupFeatureTimedMessages  GroupFeature = "timedMessages"
	GroupFeatureDirectMessages GroupFeature = "directMessages"
	GroupFeatureFullDelete     GroupFeature = "fullDelete"
	GroupFeatureReactions      GroupFeature = "reactions"
	GroupFeatureVoice          GroupFeature = "voice"
	GroupFeatureFiles          GroupFeature = "files"
	GroupFeatureSimplexLinks   GroupFeature = "simplexLinks"
	GroupFeatureReports        GroupFeature = "reports"
	GroupFeatureHistory        GroupFeature = "history"
	GroupFeatureSessions       GroupFeature = "sessions"
)

type GroupFeatureEnabled string

const (
	GroupFeatureEnabledOn  GroupFeatureEnabled = "on"
	GroupFeatureEnabledOff GroupFeatureEnabled = "off"
)

type GroupInfo struct {
	GroupId                 int64                   `json:"groupId"`
	LocalDisplayName        string                  `json:"localDisplayName"`
	GroupProfile            GroupProfile            `json:"groupProfile"`
	LocalAlias              string                  `json:"localAlias"`
	BusinessChat            *BusinessChatInfo       `json:"businessChat,omitempty"`
	FullGroupPreferences    FullGroupPreferences    `json:"fullGroupPreferences"`
	Membership              GroupMember             `json:"membership"`
	ChatSettings            ChatSettings            `json:"chatSettings"`
	CreatedAt               time.Time               `json:"createdAt"`
	UpdatedAt               time.Time               `json:"updatedAt"`
	ChatTs                  *time.Time              `json:"chatTs,omitempty"`
	UserMemberProfileSentAt *time.Time              `json:"userMemberProfileSentAt,omitempty"`
	PreparedGroup           *PreparedGroup          `json:"preparedGroup,omitempty"`
	ChatTags                []int64                 `json:"chatTags"`
	ChatItemTTL             *int64                  `json:"chatItemTTL,omitempty"`
	UiThemes                *UIThemeEntityOverrides `json:"uiThemes,omitempty"`
	CustomData              map[string]interface{}  `json:"customData,omitempty"`
	MembersRequireAttention int                     `json:"membersRequireAttention"`
	ViaGroupLinkUri         *string                 `json:"viaGroupLinkUri,omitempty"`
}

type GroupInfoSummary struct {
	GroupInfo    GroupInfo    `json:"groupInfo"`
	GroupSummary GroupSummary `json:"groupSummary"`
}

type GroupLink struct {
	UserContactLinkId     int64           `json:"userContactLinkId"`
	ConnLinkContact       CreatedConnLink `json:"connLinkContact"`
	ShortLinkDataSet      bool            `json:"shortLinkDataSet"`
	ShortLinkLargeDataSet bool            `json:"shortLinkLargeDataSet"`
	GroupLinkId           string          `json:"groupLinkId"`
	AcceptMemberRole      GroupMemberRole `json:"acceptMemberRole"`
}

type GroupLinkPlan interface {
	groupLinkPlan()
}

type GroupLinkPlanOk struct {
	GroupSLinkData_ *GroupShortLinkData `json:"groupSLinkData_,omitempty"`
}

type GroupLinkPlanOwnLink struct {
	GroupInfo GroupInfo `json:"groupInfo"`
}

type GroupLinkPlanConnectingConfirmReconnect struct{}

type GroupLinkPlanConnectingProhibit struct {
	GroupInfo_ *GroupInfo `json:"groupInfo_,omitempty"`
}

type GroupLinkPlanKnown struct {
	GroupInfo GroupInfo `json:"groupInfo"`
}

func (GroupLinkPlanOk) groupLinkPlan()                         {}
func (GroupLinkPlanOwnLink) groupLinkPlan()                    {}
func (GroupLinkPlanConnectingConfirmReconnect) groupLinkPlan() {}
func (GroupLinkPlanConnectingProhibit) groupLinkPlan()         {}
func (GroupLinkPlanKnown) groupLinkPlan()                      {}

type GroupMember struct {
	GroupMemberId          int64               `json:"groupMemberId"`
	GroupId                int64               `json:"groupId"`
	MemberId               string              `json:"memberId"`
	MemberRole             GroupMemberRole     `json:"memberRole"`
	MemberCategory         GroupMemberCategory `json:"memberCategory"`
	MemberStatus           GroupMemberStatus   `json:"memberStatus"`
	MemberSettings         GroupMemberSettings `json:"memberSettings"`
	BlockedByAdmin         bool                `json:"blockedByAdmin"`
	InvitedBy              InvitedBy           `json:"invitedBy"`
	InvitedByGroupMemberId *int64              `json:"invitedByGroupMemberId,omitempty"`
	LocalDisplayName       string              `json:"localDisplayName"`
	MemberProfile          LocalProfile        `json:"memberProfile"`
	MemberContactId        *int64              `json:"memberContactId,omitempty"`
	MemberContactProfileId int64               `json:"memberContactProfileId"`
	ActiveConn             *Connection         `json:"activeConn,omitempty"`
	MemberChatVRange       VersionRange        `json:"memberChatVRange"`
	CreatedAt              time.Time           `json:"createdAt"`
	UpdatedAt              time.Time           `json:"updatedAt"`
	SupportChat            *GroupSupportChat   `json:"supportChat,omitempty"`
}

type GroupMemberAdmission struct {
	Review *MemberCriteria `json:"review,omitempty"`
}

type GroupMemberCategory string

const (
	GroupMemberCategoryUser    GroupMemberCategory = "user"
	GroupMemberCategoryInvitee GroupMemberCategory = "invitee"
	GroupMemberCategoryHost    GroupMemberCategory = "host"
	GroupMemberCategoryPre     GroupMemberCategory = "pre"
	GroupMemberCategoryPost    GroupMemberCategory = "post"
)

type GroupMemberRef struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type GroupMemberRole string

const (
	GroupMemberRoleObserver  GroupMemberRole = "observer"
	GroupMemberRoleAuthor    GroupMemberRole = "author"
	GroupMemberRoleMember    GroupMemberRole = "member"
	GroupMemberRoleModerator GroupMemberRole = "moderator"
	GroupMemberRoleAdmin     GroupMemberRole = "admin"
	GroupMemberRoleOwner     GroupMemberRole = "owner"
)

type GroupMemberSettings struct {
	ShowMessages bool `json:"showMessages"`
}

type GroupMemberStatus string

const (
	GroupMemberStatusRejected        GroupMemberStatus = "rejected"
	GroupMemberStatusRemoved         GroupMemberStatus = "removed"
	GroupMemberStatusLeft            GroupMemberStatus = "left"
	GroupMemberStatusDeleted         GroupMemberStatus = "deleted"
	GroupMemberStatusUnknown         GroupMemberStatus = "unknown"
	GroupMemberStatusInvited         GroupMemberStatus = "invited"
	GroupMemberStatusPendingApproval GroupMemberStatus = "pending_approval"
	GroupMemberStatusPendingReview   GroupMemberStatus = "pending_review"
	GroupMemberStatusIntroduced      GroupMemberStatus = "introduced"
	GroupMemberStatusIntroInv        GroupMemberStatus = "intro-inv"
	GroupMemberStatusAccepted        GroupMemberStatus = "accepted"
	GroupMemberStatusAnnounced       GroupMemberStatus = "announced"
	GroupMemberStatusConnected       GroupMemberStatus = "connected"
	GroupMemberStatusComplete        GroupMemberStatus = "complete"
	GroupMemberStatusCreator         GroupMemberStatus = "creator"
)

type GroupPreference struct {
	Enable GroupFeatureEnabled `json:"enable"`
}

type GroupPreferences struct {
	TimedMessages  *TimedMessagesGroupPreference `json:"timedMessages,omitempty"`
	DirectMessages *RoleGroupPreference          `json:"directMessages,omitempty"`
	FullDelete     *GroupPreference              `json:"fullDelete,omitempty"`
	Reactions      *GroupPreference              `json:"reactions,omitempty"`
	Voice          *RoleGroupPreference          `json:"voice,omitempty"`
	Files          *RoleGroupPreference          `json:"files,omitempty"`
	SimplexLinks   *RoleGroupPreference          `json:"simplexLinks,omitempty"`
	Reports        *GroupPreference              `json:"reports,omitempty"`
	History        *GroupPreference              `json:"history,omitempty"`
	Sessions       *RoleGroupPreference          `json:"sessions,omitempty"`
	Commands       *[]ChatBotCommand             `json:"commands,omitempty"`
}

type GroupProfile struct {
	DisplayName      string                `json:"displayName"`
	FullName         string                `json:"fullName"`
	ShortDescr       *string               `json:"shortDescr,omitempty"`
	Description      *string               `json:"description,omitempty"`
	Image            *string               `json:"image,omitempty"`
	GroupPreferences *GroupPreferences     `json:"groupPreferences,omitempty"`
	MemberAdmission  *GroupMemberAdmission `json:"memberAdmission,omitempty"`
}

type GroupShortLinkData struct {
	GroupProfile GroupProfile `json:"groupProfile"`
}

type GroupSummary struct {
	CurrentMembers int `json:"currentMembers"`
}

type GroupSupportChat struct {
	ChatTs              time.Time  `json:"chatTs"`
	Unread              int64      `json:"unread"`
	MemberAttention     int64      `json:"memberAttention"`
	Mentions            int64      `json:"mentions"`
	LastMsgFromMemberTs *time.Time `json:"lastMsgFromMemberTs,omitempty"`
}

type HandshakeError string

const (
	HandshakeErrorPARSE       HandshakeError = "PARSE"
	HandshakeErrorIDENTITY    HandshakeError = "IDENTITY"
	HandshakeErrorBAD_AUTH    HandshakeError = "BAD_AUTH"
	HandshakeErrorBAD_SERVICE HandshakeError = "BAD_SERVICE"
)

type InlineFileMode string

const (
	InlineFileModeOffer InlineFileMode = "offer"
	InlineFileModeSent  InlineFileMode = "sent"
)

type InvitationLinkPlan interface {
	invitationLinkPlan()
}

type InvitationLinkPlanOk struct {
	ContactSLinkData_ *ContactShortLinkData `json:"contactSLinkData_,omitempty"`
}

type InvitationLinkPlanOwnLink struct{}

type InvitationLinkPlanConnecting struct {
	Contact_ *Contact `json:"contact_,omitempty"`
}

type InvitationLinkPlanKnown struct {
	Contact Contact `json:"contact"`
}

func (InvitationLinkPlanOk) invitationLinkPlan()         {}
func (InvitationLinkPlanOwnLink) invitationLinkPlan()    {}
func (InvitationLinkPlanConnecting) invitationLinkPlan() {}
func (InvitationLinkPlanKnown) invitationLinkPlan()      {}

type InvitedBy interface {
	invitedBy()
}

type InvitedByContact struct {
	ByContactId int64 `json:"byContactId"`
}

type InvitedByUser struct{}
type InvitedByUnknown struct{}

func (InvitedByContact) invitedBy() {}
func (InvitedByUser) invitedBy()    {}
func (InvitedByUnknown) invitedBy() {}

type LinkContent interface {
	linkContent()
}

type LinkContentPage struct{}

type LinkContentImage struct{}

type LinkContentVideo struct {
	Duration *int `json:"duration,omitempty"`
}

type LinkContentUnknown struct {
	Tag  string                 `json:"tag"`
	Json map[string]interface{} `json:"json"`
}

func (LinkContentPage) linkContent()    {}
func (LinkContentImage) linkContent()   {}
func (LinkContentVideo) linkContent()   {}
func (LinkContentUnknown) linkContent() {}

type LinkPreview struct {
	Uri         string       `json:"uri"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Content     *LinkContent `json:"content,omitempty"`
}

type LocalProfile struct {
	ProfileId   int64         `json:"profileId"`
	DisplayName string        `json:"displayName"`
	FullName    string        `json:"fullName"`
	ShortDescr  *string       `json:"shortDescr,omitempty"`
	Image       *string       `json:"image,omitempty"`
	ContactLink *string       `json:"contactLink,omitempty"`
	Preferences *Preferences  `json:"preferences,omitempty"`
	PeerType    *ChatPeerType `json:"peerType,omitempty"`
	LocalAlias  string        `json:"localAlias"`
}

type MemberCriteria string

const (
	MemberCriteriaAll MemberCriteria = "all"
)

type MsgChatLink interface {
	msgChatLink()
}

type MsgChatLinkContact struct {
	ConnLink string  `json:"connLink"`
	Profile  Profile `json:"profile"`
	Business bool    `json:"business"`
}

type MsgChatLinkInvitation struct {
	InvLink string  `json:"invLink"`
	Profile Profile `json:"profile"`
}

type MsgChatLinkGroup struct {
	ConnLink     string       `json:"connLink"`
	GroupProfile GroupProfile `json:"groupProfile"`
}

func (MsgChatLinkContact) msgChatLink()    {}
func (MsgChatLinkInvitation) msgChatLink() {}
func (MsgChatLinkGroup) msgChatLink()      {}

type MsgContent interface {
	msgContent()
}

type MsgContentText struct {
	Text string `json:"text"`
}

type MsgContentLink struct {
	Text    string      `json:"text"`
	Preview LinkPreview `json:"preview"`
}

type MsgContentImage struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

type MsgContentVideo struct {
	Text     string `json:"text"`
	Image    string `json:"image"`
	Duration int    `json:"duration"`
}

type MsgContentVoice struct {
	Text     string `json:"text"`
	Duration int    `json:"duration"`
}

type MsgContentFile struct {
	Text string `json:"text"`
}

type MsgContentReport struct {
	Text   string       `json:"text"`
	Reason ReportReason `json:"reason"`
}

type MsgContentChat struct {
	Text     string      `json:"text"`
	ChatLink MsgChatLink `json:"chatLink"`
}

type MsgContentUnknown struct {
	Tag  string                 `json:"tag"`
	Text string                 `json:"text"`
	Json map[string]interface{} `json:"json"`
}

func (MsgContentText) msgContent()    {}
func (MsgContentLink) msgContent()    {}
func (MsgContentImage) msgContent()   {}
func (MsgContentVideo) msgContent()   {}
func (MsgContentVoice) msgContent()   {}
func (MsgContentFile) msgContent()    {}
func (MsgContentReport) msgContent()  {}
func (MsgContentChat) msgContent()    {}
func (MsgContentUnknown) msgContent() {}

type MsgDecryptError string

const (
	MsgDecryptErrorRatchetHeader  MsgDecryptError = "ratchetHeader"
	MsgDecryptErrorTooManySkipped MsgDecryptError = "tooManySkipped"
	MsgDecryptErrorRatchetEarlier MsgDecryptError = "ratchetEarlier"
	MsgDecryptErrorOther          MsgDecryptError = "other"
	MsgDecryptErrorRatchetSync    MsgDecryptError = "ratchetSync"
)

type MsgDirection string

const (
	MsgDirectionRcv MsgDirection = "rcv"
	MsgDirectionSnd MsgDirection = "snd"
)

type MsgErrorType interface {
	msgErrorType()
}

type MsgErrorTypeMsgSkipped struct {
	FromMsgId int64 `json:"fromMsgId"`
	ToMsgId   int64 `json:"toMsgId"`
}

type MsgErrorTypeMsgBadId struct {
	MsgId int64 `json:"msgId"`
}

type MsgErrorTypeMsgBadHash struct{}
type MsgErrorTypeMsgDuplicate struct{}

func (MsgErrorTypeMsgSkipped) msgErrorType()   {}
func (MsgErrorTypeMsgBadId) msgErrorType()     {}
func (MsgErrorTypeMsgBadHash) msgErrorType()   {}
func (MsgErrorTypeMsgDuplicate) msgErrorType() {}

type MsgFilter string

const (
	MsgFilterNone     MsgFilter = "none"
	MsgFilterAll      MsgFilter = "all"
	MsgFilterMentions MsgFilter = "mentions"
)

type MsgReaction interface {
	msgReaction()
}

type MsgReactionEmoji struct {
	Emoji string `json:"emoji"`
}

type MsgReactionUnknown struct {
	Tag  string                 `json:"tag"`
	Json map[string]interface{} `json:"json"`
}

func (MsgReactionEmoji) msgReaction()   {}
func (MsgReactionUnknown) msgReaction() {}

type MsgReceiptStatus string

const (
	MsgReceiptStatusOk         MsgReceiptStatus = "ok"
	MsgReceiptStatusBadMsgHash MsgReceiptStatus = "badMsgHash"
)

type NetworkError interface {
	networkError()
}

type NetworkErrorConnectError struct {
	ConnectError string `json:"connectError"`
}

type NetworkErrorTLSError struct {
	TlsError string `json:"tlsError"`
}

type NetworkErrorUnknownCAError struct{}
type NetworkErrorFailedError struct{}
type NetworkErrorTimeoutError struct{}

type NetworkErrorSubscribeError struct {
	SubscribeError string `json:"subscribeError"`
}

func (NetworkErrorConnectError) networkError()   {}
func (NetworkErrorTLSError) networkError()       {}
func (NetworkErrorUnknownCAError) networkError() {}
func (NetworkErrorFailedError) networkError()    {}
func (NetworkErrorTimeoutError) networkError()   {}
func (NetworkErrorSubscribeError) networkError() {}

type NewUser struct {
	Profile       *Profile `json:"profile,omitempty"`
	PastTimestamp bool     `json:"pastTimestamp"`
}

type NoteFolder struct {
	NoteFolderId int64     `json:"noteFolderId"`
	UserId       int64     `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ChatTs       time.Time `json:"chatTs"`
	Favorite     bool      `json:"favorite"`
	Unread       bool      `json:"unread"`
}

type PendingContactConnection struct {
	PccConnId           int64            `json:"pccConnId"`
	PccAgentConnId      string           `json:"pccAgentConnId"`
	PccConnStatus       ConnStatus       `json:"pccConnStatus"`
	ViaContactUri       bool             `json:"viaContactUri"`
	ViaUserContactLink  *int64           `json:"viaUserContactLink,omitempty"`
	GroupLinkId         *string          `json:"groupLinkId,omitempty"`
	CustomUserProfileId *int64           `json:"customUserProfileId,omitempty"`
	ConnLinkInv         *CreatedConnLink `json:"connLinkInv,omitempty"`
	LocalAlias          string           `json:"localAlias"`
	CreatedAt           time.Time        `json:"createdAt"`
	UpdatedAt           time.Time        `json:"updatedAt"`
}

type PrefEnabled struct {
	ForUser    bool `json:"forUser"`
	ForContact bool `json:"forContact"`
}

type Preferences struct {
	TimedMessages *TimedMessagesPreference `json:"timedMessages,omitempty"`
	FullDelete    *SimplePreference        `json:"fullDelete,omitempty"`
	Reactions     *SimplePreference        `json:"reactions,omitempty"`
	Voice         *SimplePreference        `json:"voice,omitempty"`
	Files         *SimplePreference        `json:"files,omitempty"`
	Calls         *SimplePreference        `json:"calls,omitempty"`
	Sessions      *SimplePreference        `json:"sessions,omitempty"`
	Commands      *[]ChatBotCommand        `json:"commands,omitempty"`
}

type PreparedContact struct {
	ConnLinkToConnect  CreatedConnLink `json:"connLinkToConnect"`
	UiConnLinkType     ConnectionMode  `json:"uiConnLinkType"`
	WelcomeSharedMsgId *string         `json:"welcomeSharedMsgId,omitempty"`
	RequestSharedMsgId *string         `json:"requestSharedMsgId,omitempty"`
}

type PreparedGroup struct {
	ConnLinkToConnect          CreatedConnLink `json:"connLinkToConnect"`
	ConnLinkPreparedConnection bool            `json:"connLinkPreparedConnection"`
	ConnLinkStartedConnection  bool            `json:"connLinkStartedConnection"`
	WelcomeSharedMsgId         *string         `json:"welcomeSharedMsgId,omitempty"`
	RequestSharedMsgId         *string         `json:"requestSharedMsgId,omitempty"`
}

type Profile struct {
	DisplayName string        `json:"displayName"`
	FullName    string        `json:"fullName"`
	ShortDescr  *string       `json:"shortDescr,omitempty"`
	Image       *string       `json:"image,omitempty"`
	ContactLink *string       `json:"contactLink,omitempty"`
	Preferences *Preferences  `json:"preferences,omitempty"`
	PeerType    *ChatPeerType `json:"peerType,omitempty"`
}

type ProxyClientError interface {
	proxyClientError()
}

type ProxyClientErrorProtocolError struct {
	ProtocolErr ErrorType `json:"protocolErr"`
}

type ProxyClientErrorUnexpectedResponse struct {
	ResponseStr string `json:"responseStr"`
}

type ProxyClientErrorResponseError struct {
	ResponseErr ErrorType `json:"responseErr"`
}

func (ProxyClientErrorProtocolError) proxyClientError()      {}
func (ProxyClientErrorUnexpectedResponse) proxyClientError() {}
func (ProxyClientErrorResponseError) proxyClientError()      {}

type ProxyError interface {
	proxyError()
}

type ProxyErrorPROTOCOL struct {
	ProtocolErr ErrorType `json:"protocolErr"`
}

type ProxyErrorBROKER struct {
	BrokerErr BrokerErrorType `json:"brokerErr"`
}

type ProxyErrorBASIC_AUTH struct{}
type ProxyErrorNO_SESSION struct{}

func (ProxyErrorPROTOCOL) proxyError()   {}
func (ProxyErrorBROKER) proxyError()     {}
func (ProxyErrorBASIC_AUTH) proxyError() {}
func (ProxyErrorNO_SESSION) proxyError() {}

type RCErrorType interface {
	rcErrorType()
}

type RCErrorTypeInternal struct {
	InternalErr string `json:"internalErr"`
}

type RCErrorTypeIdentity struct{}
type RCErrorTypeNoLocalAddress struct{}
type RCErrorTypeNewController struct{}
type RCErrorTypeNotDiscovered struct{}
type RCErrorTypeTLSStartFailed struct{}

type RCErrorTypeException struct {
	Exception string `json:"exception"`
}

type RCErrorTypeCtrlAuth struct{}
type RCErrorTypeCtrlNotFound struct{}

type RCErrorTypeCtrlError struct {
	CtrlErr string `json:"ctrlErr"`
}

type RCErrorTypeInvitation struct{}
type RCErrorTypeVersion struct{}
type RCErrorTypeEncrypt struct{}
type RCErrorTypeDecrypt struct{}
type RCErrorTypeBlockSize struct{}

type RCErrorTypeSyntax struct {
	SyntaxErr string `json:"syntaxErr"`
}

func (RCErrorTypeInternal) rcErrorType()       {}
func (RCErrorTypeIdentity) rcErrorType()       {}
func (RCErrorTypeNoLocalAddress) rcErrorType() {}
func (RCErrorTypeNewController) rcErrorType()  {}
func (RCErrorTypeNotDiscovered) rcErrorType()  {}
func (RCErrorTypeTLSStartFailed) rcErrorType() {}
func (RCErrorTypeException) rcErrorType()      {}
func (RCErrorTypeCtrlAuth) rcErrorType()       {}
func (RCErrorTypeCtrlNotFound) rcErrorType()   {}
func (RCErrorTypeCtrlError) rcErrorType()      {}
func (RCErrorTypeInvitation) rcErrorType()     {}
func (RCErrorTypeVersion) rcErrorType()        {}
func (RCErrorTypeEncrypt) rcErrorType()        {}
func (RCErrorTypeDecrypt) rcErrorType()        {}
func (RCErrorTypeBlockSize) rcErrorType()      {}
func (RCErrorTypeSyntax) rcErrorType()         {}

type RatchetSyncState string

const (
	RatchetSyncStateOk       RatchetSyncState = "ok"
	RatchetSyncStateAllowed  RatchetSyncState = "allowed"
	RatchetSyncStateRequired RatchetSyncState = "required"
	RatchetSyncStateStarted  RatchetSyncState = "started"
	RatchetSyncStateAgreed   RatchetSyncState = "agreed"
)

type RcvConnEvent interface {
	rcvConnEvent()
}

type RcvConnEventSwitchQueue struct {
	Phase SwitchPhase `json:"phase"`
}

type RcvConnEventRatchetSync struct {
	SyncStatus RatchetSyncState `json:"syncStatus"`
}

type RcvConnEventVerificationCodeReset struct{}

type RcvConnEventPqEnabled struct {
	Enabled bool `json:"enabled"`
}

func (RcvConnEventSwitchQueue) rcvConnEvent()           {}
func (RcvConnEventRatchetSync) rcvConnEvent()           {}
func (RcvConnEventVerificationCodeReset) rcvConnEvent() {}
func (RcvConnEventPqEnabled) rcvConnEvent()             {}

type RcvDirectEvent interface {
	rcvDirectEvent()
}

type RcvDirectEventContactDeleted struct{}

type RcvDirectEventProfileUpdated struct {
	FromProfile Profile `json:"fromProfile"`
	ToProfile   Profile `json:"toProfile"`
}

type RcvDirectEventGroupInvLinkReceived struct {
	GroupProfile GroupProfile `json:"groupProfile"`
}

func (RcvDirectEventContactDeleted) rcvDirectEvent()       {}
func (RcvDirectEventProfileUpdated) rcvDirectEvent()       {}
func (RcvDirectEventGroupInvLinkReceived) rcvDirectEvent() {}

type RcvFileDescr struct {
	FileDescrId       int64  `json:"fileDescrId"`
	FileDescrText     string `json:"fileDescrText"`
	FileDescrPartNo   int    `json:"fileDescrPartNo"`
	FileDescrComplete bool   `json:"fileDescrComplete"`
}

type RcvFileInfo struct {
	FilePath    string  `json:"filePath"`
	ConnId      *int64  `json:"connId,omitempty"`
	AgentConnId *string `json:"agentConnId,omitempty"`
}

type RcvFileStatus interface {
	rcvFileStatus()
}

type RcvFileStatusNew struct{}

type RcvFileStatusAccepted struct {
	FileInfo RcvFileInfo `json:"fileInfo"`
}

type RcvFileStatusConnected struct {
	FileInfo RcvFileInfo `json:"fileInfo"`
}

type RcvFileStatusComplete struct {
	FileInfo RcvFileInfo `json:"fileInfo"`
}

type RcvFileStatusCancelled struct {
	FileInfo_ *RcvFileInfo `json:"fileInfo_,omitempty"`
}

func (RcvFileStatusNew) rcvFileStatus()       {}
func (RcvFileStatusAccepted) rcvFileStatus()  {}
func (RcvFileStatusConnected) rcvFileStatus() {}
func (RcvFileStatusComplete) rcvFileStatus()  {}
func (RcvFileStatusCancelled) rcvFileStatus() {}

type RcvFileTransfer struct {
	FileId            int64           `json:"fileId"`
	XftpRcvFile       *XFTPRcvFile    `json:"xftpRcvFile,omitempty"`
	FileInvitation    FileInvitation  `json:"fileInvitation"`
	FileStatus        RcvFileStatus   `json:"fileStatus"`
	RcvFileInline     *InlineFileMode `json:"rcvFileInline,omitempty"`
	SenderDisplayName string          `json:"senderDisplayName"`
	ChunkSize         int64           `json:"chunkSize"`
	Cancelled         bool            `json:"cancelled"`
	GrpMemberId       *int64          `json:"grpMemberId,omitempty"`
	CryptoArgs        *CryptoFileArgs `json:"cryptoArgs,omitempty"`
}

type RcvGroupEvent interface {
	rcvGroupEvent()
}

type RcvGroupEventMemberAdded struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type RcvGroupEventMemberConnected struct{}

type RcvGroupEventMemberAccepted struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type RcvGroupEventUserAccepted struct{}
type RcvGroupEventMemberLeft struct{}

type RcvGroupEventMemberRole struct {
	GroupMemberId int64           `json:"groupMemberId"`
	Profile       Profile         `json:"profile"`
	Role          GroupMemberRole `json:"role"`
}

type RcvGroupEventMemberBlocked struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
	Blocked       bool    `json:"blocked"`
}

type RcvGroupEventUserRole struct {
	Role GroupMemberRole `json:"role"`
}

type RcvGroupEventMemberDeleted struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type RcvGroupEventUserDeleted struct{}
type RcvGroupEventGroupDeleted struct{}

type RcvGroupEventGroupUpdated struct {
	GroupProfile GroupProfile `json:"groupProfile"`
}

type RcvGroupEventInvitedViaGroupLink struct{}
type RcvGroupEventMemberCreatedContact struct{}

type RcvGroupEventMemberProfileUpdated struct {
	FromProfile Profile `json:"fromProfile"`
	ToProfile   Profile `json:"toProfile"`
}

type RcvGroupEventNewMemberPendingReview struct{}

func (RcvGroupEventMemberAdded) rcvGroupEvent()            {}
func (RcvGroupEventMemberConnected) rcvGroupEvent()        {}
func (RcvGroupEventMemberAccepted) rcvGroupEvent()         {}
func (RcvGroupEventUserAccepted) rcvGroupEvent()           {}
func (RcvGroupEventMemberLeft) rcvGroupEvent()             {}
func (RcvGroupEventMemberRole) rcvGroupEvent()             {}
func (RcvGroupEventMemberBlocked) rcvGroupEvent()          {}
func (RcvGroupEventUserRole) rcvGroupEvent()               {}
func (RcvGroupEventMemberDeleted) rcvGroupEvent()          {}
func (RcvGroupEventUserDeleted) rcvGroupEvent()            {}
func (RcvGroupEventGroupDeleted) rcvGroupEvent()           {}
func (RcvGroupEventGroupUpdated) rcvGroupEvent()           {}
func (RcvGroupEventInvitedViaGroupLink) rcvGroupEvent()    {}
func (RcvGroupEventMemberCreatedContact) rcvGroupEvent()   {}
func (RcvGroupEventMemberProfileUpdated) rcvGroupEvent()   {}
func (RcvGroupEventNewMemberPendingReview) rcvGroupEvent() {}

type ReportReason string

const (
	ReportReasonSpam      ReportReason = "spam"
	ReportReasonContent   ReportReason = "content"
	ReportReasonCommunity ReportReason = "community"
	ReportReasonProfile   ReportReason = "profile"
	ReportReasonOther     ReportReason = "other"
)

type RoleGroupPreference struct {
	Enable GroupFeatureEnabled `json:"enable"`
	Role   *GroupMemberRole    `json:"role,omitempty"`
}

type SMPAgentError interface {
	smpAgentError()
}

type SMPAgentErrorA_MESSAGE struct{}

type SMPAgentErrorA_PROHIBITED struct {
	ProhibitedErr string `json:"prohibitedErr"`
}

type SMPAgentErrorA_VERSION struct{}

type SMPAgentErrorA_LINK struct {
	LinkErr string `json:"linkErr"`
}

type SMPAgentErrorA_CRYPTO struct {
	CryptoErr AgentCryptoError `json:"cryptoErr"`
}

type SMPAgentErrorA_DUPLICATE struct{}

type SMPAgentErrorA_QUEUE struct {
	QueueErr string `json:"queueErr"`
}

func (SMPAgentErrorA_MESSAGE) smpAgentError()    {}
func (SMPAgentErrorA_PROHIBITED) smpAgentError() {}
func (SMPAgentErrorA_VERSION) smpAgentError()    {}
func (SMPAgentErrorA_LINK) smpAgentError()       {}
func (SMPAgentErrorA_CRYPTO) smpAgentError()     {}
func (SMPAgentErrorA_DUPLICATE) smpAgentError()  {}
func (SMPAgentErrorA_QUEUE) smpAgentError()      {}

type SecurityCode struct {
	SecurityCode string    `json:"securityCode"`
	VerifiedAt   time.Time `json:"verifiedAt"`
}

type SimplePreference struct {
	Allow FeatureAllowed `json:"allow"`
}

type SimplexLinkType string

const (
	SimplexLinkTypeContact    SimplexLinkType = "contact"
	SimplexLinkTypeInvitation SimplexLinkType = "invitation"
	SimplexLinkTypeGroup      SimplexLinkType = "group"
	SimplexLinkTypeChannel    SimplexLinkType = "channel"
	SimplexLinkTypeRelay      SimplexLinkType = "relay"
)

type SndCIStatusProgress string

const (
	SndCIStatusProgressPartial  SndCIStatusProgress = "partial"
	SndCIStatusProgressComplete SndCIStatusProgress = "complete"
)

type SndConnEvent interface {
	sndConnEvent()
}

type SndConnEventSwitchQueue struct {
	Phase  SwitchPhase     `json:"phase"`
	Member *GroupMemberRef `json:"member,omitempty"`
}

type SndConnEventRatchetSync struct {
	SyncStatus RatchetSyncState `json:"syncStatus"`
	Member     *GroupMemberRef  `json:"member,omitempty"`
}

type SndConnEventPqEnabled struct {
	Enabled bool `json:"enabled"`
}

func (SndConnEventSwitchQueue) sndConnEvent() {}
func (SndConnEventRatchetSync) sndConnEvent() {}
func (SndConnEventPqEnabled) sndConnEvent()   {}

type SndError interface {
	sndError()
}

type SndErrorAuth struct{}
type SndErrorQuota struct{}
type SndErrorExpired struct{}

type SndErrorRelay struct {
	SrvError SrvError `json:"srvError"`
}

type SndErrorProxy struct {
	ProxyServer string   `json:"proxyServer"`
	SrvError    SrvError `json:"srvError"`
}

type SndErrorProxyRelay struct {
	ProxyServer string   `json:"proxyServer"`
	SrvError    SrvError `json:"srvError"`
}

type SndErrorOther struct {
	SndError string `json:"sndError"`
}

func (SndErrorAuth) sndError()       {}
func (SndErrorQuota) sndError()      {}
func (SndErrorExpired) sndError()    {}
func (SndErrorRelay) sndError()      {}
func (SndErrorProxy) sndError()      {}
func (SndErrorProxyRelay) sndError() {}
func (SndErrorOther) sndError()      {}

type SndFileTransfer struct {
	FileId               int64           `json:"fileId"`
	FileName             string          `json:"fileName"`
	FilePath             string          `json:"filePath"`
	FileSize             int64           `json:"fileSize"`
	ChunkSize            int64           `json:"chunkSize"`
	RecipientDisplayName string          `json:"recipientDisplayName"`
	ConnId               int64           `json:"connId"`
	AgentConnId          string          `json:"agentConnId"`
	GroupMemberId        *int64          `json:"groupMemberId,omitempty"`
	FileStatus           FileStatus      `json:"fileStatus"`
	FileDescrId          *int64          `json:"fileDescrId,omitempty"`
	FileInline           *InlineFileMode `json:"fileInline,omitempty"`
}

type SndGroupEvent interface {
	sndGroupEvent()
}

type SndGroupEventMemberRole struct {
	GroupMemberId int64           `json:"groupMemberId"`
	Profile       Profile         `json:"profile"`
	Role          GroupMemberRole `json:"role"`
}

type SndGroupEventMemberBlocked struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
	Blocked       bool    `json:"blocked"`
}

type SndGroupEventUserRole struct {
	Role GroupMemberRole `json:"role"`
}

type SndGroupEventMemberDeleted struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type SndGroupEventUserLeft struct{}

type SndGroupEventGroupUpdated struct {
	GroupProfile GroupProfile `json:"groupProfile"`
}

type SndGroupEventMemberAccepted struct {
	GroupMemberId int64   `json:"groupMemberId"`
	Profile       Profile `json:"profile"`
}

type SndGroupEventUserPendingReview struct{}

func (SndGroupEventMemberRole) sndGroupEvent()        {}
func (SndGroupEventMemberBlocked) sndGroupEvent()     {}
func (SndGroupEventUserRole) sndGroupEvent()          {}
func (SndGroupEventMemberDeleted) sndGroupEvent()     {}
func (SndGroupEventUserLeft) sndGroupEvent()          {}
func (SndGroupEventGroupUpdated) sndGroupEvent()      {}
func (SndGroupEventMemberAccepted) sndGroupEvent()    {}
func (SndGroupEventUserPendingReview) sndGroupEvent() {}

type SrvError interface {
	srvError()
}

type SrvErrorHost struct{}
type SrvErrorVersion struct{}

type SrvErrorOther struct {
	SrvError string `json:"srvError"`
}

func (SrvErrorHost) srvError()    {}
func (SrvErrorVersion) srvError() {}
func (SrvErrorOther) srvError()   {}

type StoreError interface {
	storeError()
}

type StoreErrorDuplicateName struct{}
type StoreErrorUserNotFound struct {
	UserId int64 `json:"userId"`
}
type StoreErrorUserNotFoundByName struct {
	ContactName string `json:"contactName"`
}
type StoreErrorUserNotFoundByContactId struct {
	ContactId int64 `json:"contactId"`
}
type StoreErrorUserNotFoundByGroupId struct {
	GroupId int64 `json:"groupId"`
}
type StoreErrorUserNotFoundByFileId struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorUserNotFoundByContactRequestId struct {
	ContactRequestId int64 `json:"contactRequestId"`
}
type StoreErrorContactNotFound struct {
	ContactId int64 `json:"contactId"`
}
type StoreErrorContactNotFoundByName struct {
	ContactName string `json:"contactName"`
}
type StoreErrorContactNotFoundByMemberId struct {
	GroupMemberId int64 `json:"groupMemberId"`
}
type StoreErrorContactNotReady struct {
	ContactName string `json:"contactName"`
}
type StoreErrorDuplicateContactLink struct{}
type StoreErrorUserContactLinkNotFound struct{}
type StoreErrorContactRequestNotFound struct {
	ContactRequestId int64 `json:"contactRequestId"`
}
type StoreErrorContactRequestNotFoundByName struct {
	ContactName string `json:"contactName"`
}
type StoreErrorInvalidContactRequestEntity struct {
	ContactRequestId int64 `json:"contactRequestId"`
}
type StoreErrorInvalidBusinessChatContactRequest struct{}
type StoreErrorGroupNotFound struct {
	GroupId int64 `json:"groupId"`
}
type StoreErrorGroupNotFoundByName struct {
	GroupName string `json:"groupName"`
}
type StoreErrorGroupMemberNameNotFound struct {
	GroupId         int64  `json:"groupId"`
	GroupMemberName string `json:"groupMemberName"`
}
type StoreErrorGroupMemberNotFound struct {
	GroupMemberId int64 `json:"groupMemberId"`
}
type StoreErrorGroupHostMemberNotFound struct {
	GroupId int64 `json:"groupId"`
}
type StoreErrorGroupMemberNotFoundByMemberId struct {
	MemberId string `json:"memberId"`
}
type StoreErrorMemberContactGroupMemberNotFound struct {
	ContactId int64 `json:"contactId"`
}
type StoreErrorGroupWithoutUser struct{}
type StoreErrorDuplicateGroupMember struct{}
type StoreErrorGroupAlreadyJoined struct{}
type StoreErrorGroupInvitationNotFound struct{}
type StoreErrorNoteFolderAlreadyExists struct {
	NoteFolderId int64 `json:"noteFolderId"`
}
type StoreErrorNoteFolderNotFound struct {
	NoteFolderId int64 `json:"noteFolderId"`
}
type StoreErrorUserNoteFolderNotFound struct{}
type StoreErrorSndFileNotFound struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorSndFileInvalid struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorRcvFileNotFound struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorRcvFileDescrNotFound struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorFileNotFound struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorRcvFileInvalid struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorRcvFileInvalidDescrPart struct{}
type StoreErrorLocalFileNoTransfer struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorSharedMsgIdNotFoundByFileId struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorFileIdNotFoundBySharedMsgId struct {
	SharedMsgId string `json:"sharedMsgId"`
}
type StoreErrorSndFileNotFoundXFTP struct {
	AgentSndFileId string `json:"agentSndFileId"`
}
type StoreErrorRcvFileNotFoundXFTP struct {
	AgentRcvFileId string `json:"agentRcvFileId"`
}
type StoreErrorConnectionNotFound struct {
	AgentConnId string `json:"agentConnId"`
}
type StoreErrorConnectionNotFoundById struct {
	ConnId int64 `json:"connId"`
}
type StoreErrorConnectionNotFoundByMemberId struct {
	GroupMemberId int64 `json:"groupMemberId"`
}
type StoreErrorPendingConnectionNotFound struct {
	ConnId int64 `json:"connId"`
}
type StoreErrorIntroNotFound struct{}
type StoreErrorUniqueID struct{}
type StoreErrorLargeMsg struct{}
type StoreErrorInternalError struct {
	Message string `json:"message"`
}
type StoreErrorDBException struct {
	Message string `json:"message"`
}
type StoreErrorDBBusyError struct {
	Message string `json:"message"`
}
type StoreErrorBadChatItem struct {
	ItemId int64      `json:"itemId"`
	ItemTs *time.Time `json:"itemTs,omitempty"`
}
type StoreErrorChatItemNotFound struct {
	ItemId int64 `json:"itemId"`
}
type StoreErrorChatItemNotFoundByText struct {
	Text string `json:"text"`
}
type StoreErrorChatItemSharedMsgIdNotFound struct {
	SharedMsgId string `json:"sharedMsgId"`
}
type StoreErrorChatItemNotFoundByFileId struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorChatItemNotFoundByContactId struct {
	ContactId int64 `json:"contactId"`
}
type StoreErrorChatItemNotFoundByGroupId struct {
	GroupId int64 `json:"groupId"`
}
type StoreErrorProfileNotFound struct {
	ProfileId int64 `json:"profileId"`
}
type StoreErrorDuplicateGroupLink struct {
	GroupInfo GroupInfo `json:"groupInfo"`
}
type StoreErrorGroupLinkNotFound struct {
	GroupInfo GroupInfo `json:"groupInfo"`
}
type StoreErrorHostMemberIdNotFound struct {
	GroupId int64 `json:"groupId"`
}
type StoreErrorContactNotFoundByFileId struct {
	FileId int64 `json:"fileId"`
}
type StoreErrorNoGroupSndStatus struct {
	ItemId        int64 `json:"itemId"`
	GroupMemberId int64 `json:"groupMemberId"`
}
type StoreErrorDuplicateGroupMessage struct {
	GroupId                  int64  `json:"groupId"`
	SharedMsgId              string `json:"sharedMsgId"`
	AuthorGroupMemberId      *int64 `json:"authorGroupMemberId,omitempty"`
	ForwardedByGroupMemberId *int64 `json:"forwardedByGroupMemberId,omitempty"`
}
type StoreErrorRemoteHostNotFound struct {
	RemoteHostId int64 `json:"remoteHostId"`
}
type StoreErrorRemoteHostUnknown struct{}
type StoreErrorRemoteHostDuplicateCA struct{}
type StoreErrorRemoteCtrlNotFound struct {
	RemoteCtrlId int64 `json:"remoteCtrlId"`
}
type StoreErrorRemoteCtrlDuplicateCA struct{}
type StoreErrorProhibitedDeleteUser struct {
	UserId    int64 `json:"userId"`
	ContactId int64 `json:"contactId"`
}
type StoreErrorOperatorNotFound struct {
	ServerOperatorId int64 `json:"serverOperatorId"`
}
type StoreErrorUsageConditionsNotFound struct{}
type StoreErrorInvalidQuote struct{}
type StoreErrorInvalidMention struct{}

func (StoreErrorDuplicateName) storeError()                     {}
func (StoreErrorUserNotFound) storeError()                      {}
func (StoreErrorUserNotFoundByName) storeError()                {}
func (StoreErrorUserNotFoundByContactId) storeError()           {}
func (StoreErrorUserNotFoundByGroupId) storeError()             {}
func (StoreErrorUserNotFoundByFileId) storeError()              {}
func (StoreErrorUserNotFoundByContactRequestId) storeError()    {}
func (StoreErrorContactNotFound) storeError()                   {}
func (StoreErrorContactNotFoundByName) storeError()             {}
func (StoreErrorContactNotFoundByMemberId) storeError()         {}
func (StoreErrorContactNotReady) storeError()                   {}
func (StoreErrorDuplicateContactLink) storeError()              {}
func (StoreErrorUserContactLinkNotFound) storeError()           {}
func (StoreErrorContactRequestNotFound) storeError()            {}
func (StoreErrorContactRequestNotFoundByName) storeError()      {}
func (StoreErrorInvalidContactRequestEntity) storeError()       {}
func (StoreErrorInvalidBusinessChatContactRequest) storeError() {}
func (StoreErrorGroupNotFound) storeError()                     {}
func (StoreErrorGroupNotFoundByName) storeError()               {}
func (StoreErrorGroupMemberNameNotFound) storeError()           {}
func (StoreErrorGroupMemberNotFound) storeError()               {}
func (StoreErrorGroupHostMemberNotFound) storeError()           {}
func (StoreErrorGroupMemberNotFoundByMemberId) storeError()     {}
func (StoreErrorMemberContactGroupMemberNotFound) storeError()  {}
func (StoreErrorGroupWithoutUser) storeError()                  {}
func (StoreErrorDuplicateGroupMember) storeError()              {}
func (StoreErrorGroupAlreadyJoined) storeError()                {}
func (StoreErrorGroupInvitationNotFound) storeError()           {}
func (StoreErrorNoteFolderAlreadyExists) storeError()           {}
func (StoreErrorNoteFolderNotFound) storeError()                {}
func (StoreErrorUserNoteFolderNotFound) storeError()            {}
func (StoreErrorSndFileNotFound) storeError()                   {}
func (StoreErrorSndFileInvalid) storeError()                    {}
func (StoreErrorRcvFileNotFound) storeError()                   {}
func (StoreErrorRcvFileDescrNotFound) storeError()              {}
func (StoreErrorFileNotFound) storeError()                      {}
func (StoreErrorRcvFileInvalid) storeError()                    {}
func (StoreErrorRcvFileInvalidDescrPart) storeError()           {}
func (StoreErrorLocalFileNoTransfer) storeError()               {}
func (StoreErrorSharedMsgIdNotFoundByFileId) storeError()       {}
func (StoreErrorFileIdNotFoundBySharedMsgId) storeError()       {}
func (StoreErrorSndFileNotFoundXFTP) storeError()               {}
func (StoreErrorRcvFileNotFoundXFTP) storeError()               {}
func (StoreErrorConnectionNotFound) storeError()                {}
func (StoreErrorConnectionNotFoundById) storeError()            {}
func (StoreErrorConnectionNotFoundByMemberId) storeError()      {}
func (StoreErrorPendingConnectionNotFound) storeError()         {}
func (StoreErrorIntroNotFound) storeError()                     {}
func (StoreErrorUniqueID) storeError()                          {}
func (StoreErrorLargeMsg) storeError()                          {}
func (StoreErrorInternalError) storeError()                     {}
func (StoreErrorDBException) storeError()                       {}
func (StoreErrorDBBusyError) storeError()                       {}
func (StoreErrorBadChatItem) storeError()                       {}
func (StoreErrorChatItemNotFound) storeError()                  {}
func (StoreErrorChatItemNotFoundByText) storeError()            {}
func (StoreErrorChatItemSharedMsgIdNotFound) storeError()       {}
func (StoreErrorChatItemNotFoundByFileId) storeError()          {}
func (StoreErrorChatItemNotFoundByContactId) storeError()       {}
func (StoreErrorChatItemNotFoundByGroupId) storeError()         {}
func (StoreErrorProfileNotFound) storeError()                   {}
func (StoreErrorDuplicateGroupLink) storeError()                {}
func (StoreErrorGroupLinkNotFound) storeError()                 {}
func (StoreErrorHostMemberIdNotFound) storeError()              {}
func (StoreErrorContactNotFoundByFileId) storeError()           {}
func (StoreErrorNoGroupSndStatus) storeError()                  {}
func (StoreErrorDuplicateGroupMessage) storeError()             {}
func (StoreErrorRemoteHostNotFound) storeError()                {}
func (StoreErrorRemoteHostUnknown) storeError()                 {}
func (StoreErrorRemoteHostDuplicateCA) storeError()             {}
func (StoreErrorRemoteCtrlNotFound) storeError()                {}
func (StoreErrorRemoteCtrlDuplicateCA) storeError()             {}
func (StoreErrorProhibitedDeleteUser) storeError()              {}
func (StoreErrorOperatorNotFound) storeError()                  {}
func (StoreErrorUsageConditionsNotFound) storeError()           {}
func (StoreErrorInvalidQuote) storeError()                      {}
func (StoreErrorInvalidMention) storeError()                    {}

type SwitchPhase string

const (
	SwitchPhaseStarted   SwitchPhase = "started"
	SwitchPhaseConfirmed SwitchPhase = "confirmed"
	SwitchPhaseSecured   SwitchPhase = "secured"
	SwitchPhaseCompleted SwitchPhase = "completed"
)

type TimedMessagesGroupPreference struct {
	Enable GroupFeatureEnabled `json:"enable"`
	Ttl    *int                `json:"ttl,omitempty"`
}

type TimedMessagesPreference struct {
	Allow FeatureAllowed `json:"allow"`
	Ttl   *int           `json:"ttl,omitempty"`
}

type TransportError interface {
	transportError()
}

type TransportErrorBadBlock struct{}
type TransportErrorVersion struct{}
type TransportErrorLargeMsg struct{}
type TransportErrorBadSession struct{}
type TransportErrorNoServerAuth struct{}

type TransportErrorHandshake struct {
	HandshakeErr HandshakeError `json:"handshakeErr"`
}

func (TransportErrorBadBlock) transportError()     {}
func (TransportErrorVersion) transportError()      {}
func (TransportErrorLargeMsg) transportError()     {}
func (TransportErrorBadSession) transportError()   {}
func (TransportErrorNoServerAuth) transportError() {}
func (TransportErrorHandshake) transportError()    {}

type UIColorMode string

const (
	UIColorModeLight UIColorMode = "light"
	UIColorModeDark  UIColorMode = "dark"
)

type UIColors struct {
	Accent           *string `json:"accent,omitempty"`
	AccentVariant    *string `json:"accentVariant,omitempty"`
	Secondary        *string `json:"secondary,omitempty"`
	SecondaryVariant *string `json:"secondaryVariant,omitempty"`
	Background       *string `json:"background,omitempty"`
	Menus            *string `json:"menus,omitempty"`
	Title            *string `json:"title,omitempty"`
	AccentVariant2   *string `json:"accentVariant2,omitempty"`
	SentMessage      *string `json:"sentMessage,omitempty"`
	SentReply        *string `json:"sentReply,omitempty"`
	ReceivedMessage  *string `json:"receivedMessage,omitempty"`
	ReceivedReply    *string `json:"receivedReply,omitempty"`
}

type UIThemeEntityOverride struct {
	Mode      UIColorMode    `json:"mode"`
	Wallpaper *ChatWallpaper `json:"wallpaper,omitempty"`
	Colors    UIColors       `json:"colors"`
}

type UIThemeEntityOverrides struct {
	Light *UIThemeEntityOverride `json:"light,omitempty"`
	Dark  *UIThemeEntityOverride `json:"dark,omitempty"`
}

type UpdatedMessage struct {
	MsgContent MsgContent       `json:"msgContent"`
	Mentions   map[string]int64 `json:"mentions"`
}

type User struct {
	UserId                     int64                   `json:"userId"`
	AgentUserId                int64                   `json:"agentUserId"`
	UserContactId              int64                   `json:"userContactId"`
	LocalDisplayName           string                  `json:"localDisplayName"`
	Profile                    LocalProfile            `json:"profile"`
	FullPreferences            FullPreferences         `json:"fullPreferences"`
	ActiveUser                 bool                    `json:"activeUser"`
	ActiveOrder                int64                   `json:"activeOrder"`
	ViewPwdHash                *UserPwdHash            `json:"viewPwdHash,omitempty"`
	ShowNtfs                   bool                    `json:"showNtfs"`
	SendRcptsContacts          bool                    `json:"sendRcptsContacts"`
	SendRcptsSmallGroups       bool                    `json:"sendRcptsSmallGroups"`
	AutoAcceptMemberContacts   bool                    `json:"autoAcceptMemberContacts"`
	UserMemberProfileUpdatedAt *time.Time              `json:"userMemberProfileUpdatedAt,omitempty"`
	UiThemes                   *UIThemeEntityOverrides `json:"uiThemes,omitempty"`
}

type UserContact struct {
	UserContactLinkId int64  `json:"userContactLinkId"`
	ConnReqContact    string `json:"connReqContact"`
	GroupId           *int64 `json:"groupId,omitempty"`
}

type UserContactLink struct {
	UserContactLinkId     int64           `json:"userContactLinkId"`
	ConnLinkContact       CreatedConnLink `json:"connLinkContact"`
	ShortLinkDataSet      bool            `json:"shortLinkDataSet"`
	ShortLinkLargeDataSet bool            `json:"shortLinkLargeDataSet"`
	AddressSettings       AddressSettings `json:"addressSettings"`
}

type UserContactRequest struct {
	ContactRequestId   int64        `json:"contactRequestId"`
	AgentInvitationId  string       `json:"agentInvitationId"`
	ContactId_         *int64       `json:"contactId_,omitempty"`
	BusinessGroupId_   *int64       `json:"businessGroupId_,omitempty"`
	UserContactLinkId_ *int64       `json:"userContactLinkId_,omitempty"`
	CReqChatVRange     VersionRange `json:"cReqChatVRange"`
	LocalDisplayName   string       `json:"localDisplayName"`
	ProfileId          int64        `json:"profileId"`
	Profile            Profile      `json:"profile"`
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	XContactId         *string      `json:"xContactId,omitempty"`
	PqSupport          bool         `json:"pqSupport"`
	WelcomeSharedMsgId *string      `json:"welcomeSharedMsgId,omitempty"`
	RequestSharedMsgId *string      `json:"requestSharedMsgId,omitempty"`
}

type UserInfo struct {
	User        User `json:"user"`
	UnreadCount int  `json:"unreadCount"`
}

type UserProfileUpdateSummary struct {
	UpdateSuccesses int       `json:"updateSuccesses"`
	UpdateFailures  int       `json:"updateFailures"`
	ChangedContacts []Contact `json:"changedContacts"`
}

type UserPwdHash struct {
	Hash string `json:"hash"`
	Salt string `json:"salt"`
}

type VersionRange struct {
	MinVersion int `json:"minVersion"`
	MaxVersion int `json:"maxVersion"`
}

type XFTPErrorType interface {
	xftpErrorType()
}

type XFTPErrorTypeBLOCK struct{}
type XFTPErrorTypeSESSION struct{}
type XFTPErrorTypeHANDSHAKE struct{}
type XFTPErrorTypeCMD struct {
	CmdErr CommandError `json:"cmdErr"`
}
type XFTPErrorTypeAUTH struct{}
type XFTPErrorTypeBLOCKED struct {
	BlockInfo BlockingInfo `json:"blockInfo"`
}
type XFTPErrorTypeSIZE struct{}
type XFTPErrorTypeQUOTA struct{}
type XFTPErrorTypeDIGEST struct{}
type XFTPErrorTypeCRYPTO struct{}
type XFTPErrorTypeNO_FILE struct{}
type XFTPErrorTypeHAS_FILE struct{}
type XFTPErrorTypeFILE_IO struct{}
type XFTPErrorTypeTIMEOUT struct{}
type XFTPErrorTypeINTERNAL struct{}
type XFTPErrorTypeDUPLICATE_ struct{}

func (XFTPErrorTypeBLOCK) xftpErrorType()      {}
func (XFTPErrorTypeSESSION) xftpErrorType()    {}
func (XFTPErrorTypeHANDSHAKE) xftpErrorType()  {}
func (XFTPErrorTypeCMD) xftpErrorType()        {}
func (XFTPErrorTypeAUTH) xftpErrorType()       {}
func (XFTPErrorTypeBLOCKED) xftpErrorType()    {}
func (XFTPErrorTypeSIZE) xftpErrorType()       {}
func (XFTPErrorTypeQUOTA) xftpErrorType()      {}
func (XFTPErrorTypeDIGEST) xftpErrorType()     {}
func (XFTPErrorTypeCRYPTO) xftpErrorType()     {}
func (XFTPErrorTypeNO_FILE) xftpErrorType()    {}
func (XFTPErrorTypeHAS_FILE) xftpErrorType()   {}
func (XFTPErrorTypeFILE_IO) xftpErrorType()    {}
func (XFTPErrorTypeTIMEOUT) xftpErrorType()    {}
func (XFTPErrorTypeINTERNAL) xftpErrorType()   {}
func (XFTPErrorTypeDUPLICATE_) xftpErrorType() {}

type XFTPRcvFile struct {
	RcvFileDescription  RcvFileDescr `json:"rcvFileDescription"`
	AgentRcvFileId      *string      `json:"agentRcvFileId,omitempty"`
	AgentRcvFileDeleted bool         `json:"agentRcvFileDeleted"`
	UserApprovedRelays  bool         `json:"userApprovedRelays"`
}

type XFTPSndFile struct {
	AgentSndFileId      string          `json:"agentSndFileId"`
	PrivateSndFileDescr *string         `json:"privateSndFileDescr,omitempty"`
	AgentSndFileDeleted bool            `json:"agentSndFileDeleted"`
	CryptoArgs          *CryptoFileArgs `json:"cryptoArgs,omitempty"`
}
