// <auto-generated>
//   This file was generated by a tool; you should avoid making direct changes.
//   Consider using 'partial classes' to extend these types
//   Input: Net_Inbox.proto
// </auto-generated>

#region Designer generated code
#pragma warning disable CS0612, CS0618, CS1591, CS3021, IDE0079, IDE1006, RCS1036, RCS1057, RCS1085, RCS1192
namespace Pb
{

    [global::ProtoBuf.ProtoContract()]
    public partial class Base_Inbox_InboxInfo : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public global::System.Collections.Generic.List<Base_Inbox_InboxItemInfo> ItemList { get; set; } = new global::System.Collections.Generic.List<Base_Inbox_InboxItemInfo>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Base_Inbox_InboxItemInfo : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Id { get; set; }

        [global::ProtoBuf.ProtoMember(2)]
        public int Type { get; set; }

        [global::ProtoBuf.ProtoMember(3)]
        [global::System.ComponentModel.DefaultValue("")]
        public string ReceiverId { get; set; } = "";

        [global::ProtoBuf.ProtoMember(4)]
        [global::System.ComponentModel.DefaultValue("")]
        public string SenderId { get; set; } = "";

        [global::ProtoBuf.ProtoMember(5)]
        [global::System.ComponentModel.DefaultValue("")]
        public string SenderName { get; set; } = "";

        [global::ProtoBuf.ProtoMember(6)]
        [global::System.ComponentModel.DefaultValue("")]
        public string Title { get; set; } = "";

        [global::ProtoBuf.ProtoMember(7)]
        [global::System.ComponentModel.DefaultValue("")]
        public string Content { get; set; } = "";

        [global::ProtoBuf.ProtoMember(8)]
        public bool IsRead { get; set; }

        [global::ProtoBuf.ProtoMember(9)]
        public bool IsClaimReward { get; set; }

        [global::ProtoBuf.ProtoMember(10)]
        public int ItemId1 { get; set; }

        [global::ProtoBuf.ProtoMember(11)]
        public int ItemNum1 { get; set; }

        [global::ProtoBuf.ProtoMember(12)]
        public int ItemId2 { get; set; }

        [global::ProtoBuf.ProtoMember(13)]
        public int ItemNum2 { get; set; }

        [global::ProtoBuf.ProtoMember(14)]
        public int ItemId3 { get; set; }

        [global::ProtoBuf.ProtoMember(15)]
        public int ItemNum3 { get; set; }

        [global::ProtoBuf.ProtoMember(16)]
        public long Ts { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_UpdateItemsInfoRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        [global::ProtoBuf.ProtoMap]
        public global::System.Collections.Generic.Dictionary<int, Base_Inbox_InboxItemInfo> UpdateItemList { get; set; } = new global::System.Collections.Generic.Dictionary<int, Base_Inbox_InboxItemInfo>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_SendEmailReq : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public Base_Inbox_InboxItemInfo InboxItemInfo { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_SendEmailRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Ret { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_DeleteEmailReq : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1, IsPacked = true)]
        public int[] EmailIds { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_DeleteEmailRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Ret { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_ClaimRewardGiftsReq : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int EmailId { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_ClaimRewardGiftsRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Ret { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_GetInboxListRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        [global::ProtoBuf.ProtoMap]
        public global::System.Collections.Generic.Dictionary<int, Base_Inbox_InboxItemInfo> InboxItemInfoList { get; set; } = new global::System.Collections.Generic.Dictionary<int, Base_Inbox_InboxItemInfo>();

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_ReceiveGiftsReq : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Index { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public partial class Net_Inbox_ReceiveGiftsRet : global::ProtoBuf.IExtensible
    {
        private global::ProtoBuf.IExtension __pbn__extensionData;
        global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
            => global::ProtoBuf.Extensible.GetExtensionObject(ref __pbn__extensionData, createIfMissing);

        [global::ProtoBuf.ProtoMember(1)]
        public int Ret { get; set; }

    }

    [global::ProtoBuf.ProtoContract()]
    public enum InboxType
    {
        InboxIdle = 0,
        InboxTypeSys = 1,
    }

}

#pragma warning restore CS0612, CS0618, CS1591, CS3021, IDE0079, IDE1006, RCS1036, RCS1057, RCS1085, RCS1192
#endregion
