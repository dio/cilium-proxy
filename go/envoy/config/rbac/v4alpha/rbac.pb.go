// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v4alpha/rbac.proto

package envoy_config_rbac_v4alpha

import (
	fmt "fmt"
	v4alpha1 "github.com/cilium/proxy/go/envoy/config/core/v4alpha"
	v4alpha "github.com/cilium/proxy/go/envoy/config/route/v4alpha"
	v3 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Should we do safe-list or block-list style access control?
type RBAC_Action int32

const (
	// The policies grant access to principals. The rest is denied. This is safe-list style
	// access control. This is the default type.
	RBAC_ALLOW RBAC_Action = 0
	// The policies deny access to principals. The rest is allowed. This is block-list style
	// access control.
	RBAC_DENY RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{0, 0}
}

// Role Based Access Control (RBAC) provides service-level and method-level access control for a
// service. RBAC policies are additive. The policies are examined in order. A request is allowed
// once a matching policy is found (suppose the `action` is ALLOW).
//
// Here is an example of RBAC configuration. It has two policies:
//
// * Service account "cluster.local/ns/default/sa/admin" has full access to the service, and so
//   does "cluster.local/ns/default/sa/superuser".
//
// * Any user can read ("GET") the service at paths with prefix "/products", so long as the
//   destination port is either 80 or 443.
//
//  .. code-block:: yaml
//
//   action: ALLOW
//   policies:
//     "service-admin":
//       permissions:
//         - any: true
//       principals:
//         - authenticated:
//             principal_name:
//               exact: "cluster.local/ns/default/sa/admin"
//         - authenticated:
//             principal_name:
//               exact: "cluster.local/ns/default/sa/superuser"
//     "product-viewer":
//       permissions:
//           - and_rules:
//               rules:
//                 - header: { name: ":method", exact_match: "GET" }
//                 - url_path:
//                     path: { prefix: "/products" }
//                 - or_rules:
//                     rules:
//                       - destination_port: 80
//                       - destination_port: 443
//       principals:
//         - any: true
//
type RBAC struct {
	// The action to take if a policy matches. The request is allowed if and only if:
	//
	//   * `action` is "ALLOWED" and at least one policy matches
	//   * `action` is "DENY" and none of the policies match
	Action RBAC_Action `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v4alpha.RBAC_Action" json:"action,omitempty"`
	// Maps from policy name to policy. A match occurs when at least one policy matches the request.
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy specifies a role and the principals that are assigned/denied the role. A policy matches if
// and only if at least one of its permissions match the action taking place AND at least one of its
// principals match the downstream AND the condition is true if specified.
type Policy struct {
	// Required. The set of permissions that define a role. Each permission is matched with OR
	// semantics. To match all actions for this policy, a single Permission with the `any` field set
	// to true should be used.
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Required. The set of principals that are assigned/denied the role based on “action”. Each
	// principal is matched with OR semantics. To match all downstreams for this policy, a single
	// Principal with the `any` field set to true should be used.
	Principals []*Principal `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	// An optional symbolic expression specifying an access control
	// :ref:`condition <arch_overview_condition>`. The condition is combined
	// with the permissions and the principals as a clause with AND semantics.
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

// Permission defines an action (or actions) that a principal can take.
// [#next-free-field: 11]
type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_UrlPath
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *v4alpha.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_UrlPath struct {
	UrlPath *v3.PathMatcher `protobuf:"bytes,10,opt,name=url_path,json=urlPath,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *v4alpha1.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *v3.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *v3.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_UrlPath) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *v4alpha.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetUrlPath() *v3.PathMatcher {
	if x, ok := m.GetRule().(*Permission_UrlPath); ok {
		return x.UrlPath
	}
	return nil
}

func (m *Permission) GetDestinationIp() *v4alpha1.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *v3.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *v3.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_UrlPath)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

// Used in the `and_rules` and `or_rules` fields in the `rule` oneof. Depending on the context,
// each are applied with the associated behavior.
type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Principal defines an identity or a group of identities for a downstream subject.
// [#next-free-field: 12]
type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_DirectRemoteIp
	//	*Principal_RemoteIp
	//	*Principal_Header
	//	*Principal_UrlPath
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_DirectRemoteIp struct {
	DirectRemoteIp *v4alpha1.CidrRange `protobuf:"bytes,10,opt,name=direct_remote_ip,json=directRemoteIp,proto3,oneof"`
}

type Principal_RemoteIp struct {
	RemoteIp *v4alpha1.CidrRange `protobuf:"bytes,11,opt,name=remote_ip,json=remoteIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *v4alpha.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_UrlPath struct {
	UrlPath *v3.PathMatcher `protobuf:"bytes,9,opt,name=url_path,json=urlPath,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *v3.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_DirectRemoteIp) isPrincipal_Identifier() {}

func (*Principal_RemoteIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_UrlPath) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetDirectRemoteIp() *v4alpha1.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_DirectRemoteIp); ok {
		return x.DirectRemoteIp
	}
	return nil
}

func (m *Principal) GetRemoteIp() *v4alpha1.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_RemoteIp); ok {
		return x.RemoteIp
	}
	return nil
}

func (m *Principal) GetHeader() *v4alpha.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetUrlPath() *v3.PathMatcher {
	if x, ok := m.GetIdentifier().(*Principal_UrlPath); ok {
		return x.UrlPath
	}
	return nil
}

func (m *Principal) GetMetadata() *v3.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_DirectRemoteIp)(nil),
		(*Principal_RemoteIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_UrlPath)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

// Used in the `and_ids` and `or_ids` fields in the `identifier` oneof. Depending on the context,
// each are applied with the associated behavior.
type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Authentication attributes for a downstream.
type Principal_Authenticated struct {
	// The name of the principal. If set, The URI SAN or DNS SAN in that order is used from the
	// certificate, otherwise the subject field is used. If unset, it applies to any user that is
	// authenticated.
	PrincipalName        *v3.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd4c25c664ecf48, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *v3.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v4alpha.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v4alpha.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v4alpha.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v4alpha.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v4alpha.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v4alpha.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v4alpha.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v4alpha.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v4alpha.Principal.Authenticated")
}

func init() {
	proto.RegisterFile("envoy/config/rbac/v4alpha/rbac.proto", fileDescriptor_6cd4c25c664ecf48)
}

var fileDescriptor_6cd4c25c664ecf48 = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x18, 0x8d, 0x9d, 0xc4, 0x75, 0xbe, 0x28, 0x25, 0x0c, 0x42, 0x98, 0xac, 0xb6, 0xa4, 0xa1, 0xed,
	0x76, 0x57, 0x60, 0xab, 0x2d, 0x02, 0x14, 0xc1, 0x42, 0xdd, 0x2d, 0x4a, 0x97, 0x6d, 0x29, 0xee,
	0x01, 0xb1, 0x07, 0xa2, 0xa9, 0x3d, 0xdb, 0x0c, 0x24, 0x33, 0x66, 0x3c, 0x89, 0x9a, 0x1b, 0x37,
	0xb8, 0x70, 0xe1, 0x82, 0xc4, 0x9f, 0xc1, 0x91, 0x3b, 0x12, 0x57, 0xfe, 0x00, 0xfe, 0x0f, 0xd4,
	0xcb, 0xa2, 0x19, 0x3b, 0xbf, 0x96, 0xb4, 0x4d, 0xab, 0xbd, 0x25, 0x9e, 0xf7, 0xde, 0xbc, 0xf9,
	0xe6, 0x7b, 0x9f, 0x0d, 0x6b, 0x84, 0x0d, 0xf8, 0xd0, 0x0b, 0x39, 0x7b, 0x46, 0xcf, 0x3c, 0x71,
	0x8a, 0x43, 0x6f, 0xf0, 0x1e, 0xee, 0xc6, 0x1d, 0xac, 0xff, 0xb8, 0xb1, 0xe0, 0x92, 0xa3, 0x37,
	0x35, 0xca, 0x4d, 0x51, 0xae, 0x5e, 0xc8, 0x50, 0xb5, 0x7b, 0x33, 0x02, 0x21, 0x17, 0x64, 0x2c,
	0x80, 0xa3, 0x48, 0x90, 0x24, 0x49, 0x35, 0x6a, 0x5b, 0xb3, 0x3b, 0xf1, 0xbe, 0x9c, 0x20, 0xf5,
	0xbf, 0x76, 0xc8, 0x7b, 0x31, 0x67, 0x84, 0xc9, 0x11, 0x25, 0x33, 0x27, 0x87, 0x31, 0xf1, 0x7a,
	0x58, 0x86, 0x1d, 0x22, 0xbc, 0xc1, 0x8e, 0xd7, 0x23, 0x12, 0x47, 0x58, 0xe2, 0x0c, 0x55, 0x9f,
	0x8f, 0x8a, 0xb1, 0xec, 0x64, 0x88, 0xc6, 0x7c, 0x44, 0x22, 0x05, 0x65, 0x67, 0x19, 0x66, 0xfd,
	0x8c, 0xf3, 0xb3, 0x2e, 0xf1, 0x70, 0x4c, 0x3d, 0x72, 0x1e, 0x0b, 0x6f, 0xb0, 0xa5, 0xbd, 0x6d,
	0x79, 0xc9, 0x90, 0x49, 0x7c, 0x9e, 0xc1, 0xee, 0xf6, 0xa3, 0x18, 0x7b, 0x98, 0x31, 0x2e, 0xb1,
	0xa4, 0x9c, 0x25, 0x5e, 0x22, 0xb1, 0xec, 0x8f, 0x1c, 0xaf, 0xfe, 0x6f, 0x79, 0x40, 0x44, 0x42,
	0x39, 0x9b, 0x6c, 0xf4, 0xc6, 0x00, 0x77, 0x69, 0x84, 0xd5, 0xe9, 0xb3, 0x1f, 0xe9, 0x42, 0xe3,
	0x77, 0x13, 0x0a, 0x81, 0xbf, 0xbb, 0x87, 0x1e, 0x82, 0x85, 0x43, 0xc5, 0x76, 0x8c, 0xba, 0xb1,
	0xb9, 0xbc, 0xbd, 0xe1, 0x5e, 0x5a, 0x7e, 0x57, 0x11, 0xdc, 0x5d, 0x8d, 0x0e, 0x32, 0x16, 0x3a,
	0x00, 0x3b, 0xe6, 0x5d, 0x1a, 0x52, 0x92, 0x38, 0x66, 0x3d, 0xbf, 0x59, 0xde, 0x7e, 0xf7, 0x3a,
	0x85, 0xe3, 0x0c, 0xbf, 0xcf, 0xa4, 0x18, 0x06, 0x63, 0x7a, 0xed, 0x1b, 0xa8, 0xcc, 0x2c, 0xa1,
	0x2a, 0xe4, 0xbf, 0x23, 0x43, 0x6d, 0xac, 0x14, 0xa8, 0x9f, 0xe8, 0x03, 0x28, 0x0e, 0x70, 0xb7,
	0x4f, 0x1c, 0xb3, 0x6e, 0x6c, 0x96, 0xb7, 0x57, 0xaf, 0xd8, 0x4a, 0x4b, 0x0d, 0x83, 0x14, 0xdf,
	0x34, 0x3f, 0x34, 0x1a, 0x77, 0xc1, 0x4a, 0xcd, 0xa3, 0x12, 0x14, 0x77, 0x9f, 0x3c, 0xf9, 0xe2,
	0xab, 0x6a, 0x0e, 0xd9, 0x50, 0x78, 0xb4, 0x7f, 0xf4, 0x75, 0xd5, 0x68, 0xd6, 0x7f, 0xfb, 0xf3,
	0xa7, 0x95, 0x3b, 0x30, 0xaf, 0xfd, 0x76, 0xb4, 0xf1, 0xc6, 0xcf, 0x26, 0x58, 0xa9, 0x2c, 0xfa,
	0x12, 0xca, 0x31, 0x11, 0x3d, 0x9a, 0xa8, 0x7a, 0x27, 0x8e, 0xa1, 0x4f, 0xbe, 0x7e, 0x95, 0x9d,
	0x31, 0xda, 0xb7, 0x2f, 0xfc, 0xe2, 0x2f, 0x86, 0x69, 0x1b, 0xc1, 0xb4, 0x06, 0x3a, 0x02, 0x88,
	0x05, 0x65, 0x21, 0x8d, 0x71, 0x77, 0x54, 0xcb, 0xb5, 0xab, 0x14, 0x47, 0xe0, 0x29, 0xc1, 0x29,
	0x05, 0xf4, 0x11, 0x94, 0x42, 0xce, 0x22, 0xaa, 0x2f, 0x37, 0xaf, 0xeb, 0xb5, 0xe2, 0xa6, 0x8d,
	0xe7, 0xe2, 0x98, 0xba, 0xaa, 0xf1, 0xdc, 0x51, 0xe3, 0xb9, 0xfb, 0xe7, 0xb1, 0x08, 0x26, 0x84,
	0x66, 0x43, 0x55, 0xe3, 0x2e, 0xdc, 0x99, 0x5b, 0x8d, 0xb4, 0x08, 0x8d, 0x7f, 0x2c, 0x80, 0xc9,
	0xb9, 0x50, 0x0b, 0x4a, 0x98, 0x45, 0x6d, 0xd1, 0xef, 0x92, 0x44, 0x5f, 0x5a, 0x79, 0xfb, 0xfe,
	0x42, 0x15, 0x71, 0x4f, 0x88, 0x6c, 0xe5, 0x02, 0x1b, 0xb3, 0x28, 0x50, 0x64, 0xf4, 0x19, 0xd8,
	0x5c, 0x64, 0x42, 0xe6, 0xcd, 0x85, 0x96, 0xb8, 0x48, 0x75, 0xee, 0x40, 0x1e, 0xb3, 0xa1, 0x3e,
	0xbc, 0xed, 0x2f, 0x5d, 0xf8, 0x85, 0x6f, 0x4d, 0xdb, 0x68, 0xe5, 0x02, 0xf5, 0x14, 0xed, 0x81,
	0xd5, 0x21, 0x38, 0x22, 0xc2, 0x29, 0xcc, 0xdd, 0x42, 0x8d, 0x89, 0xf1, 0x1e, 0x2d, 0x8d, 0x3c,
	0x4c, 0xf3, 0xdc, 0xca, 0x05, 0x19, 0x15, 0x7d, 0x02, 0x76, 0x5f, 0x74, 0xdb, 0x2a, 0xff, 0x0e,
	0x68, 0x99, 0x46, 0x26, 0xa3, 0x06, 0x80, 0x9b, 0x0d, 0x00, 0x5d, 0x33, 0x2c, 0x3b, 0x13, 0xfe,
	0x52, 0x5f, 0x74, 0xd5, 0x13, 0x74, 0x08, 0xcb, 0x11, 0x49, 0x24, 0x65, 0x3a, 0xc2, 0x6d, 0x1a,
	0x3b, 0x45, 0x2d, 0xf3, 0xc2, 0xcd, 0xab, 0x59, 0x37, 0x36, 0xb3, 0x47, 0x23, 0x11, 0x60, 0x76,
	0x46, 0x5a, 0xb9, 0xa0, 0x32, 0xc5, 0x3e, 0x88, 0xd1, 0xfb, 0x50, 0x9d, 0x96, 0x8b, 0xb9, 0x90,
	0x8e, 0x55, 0x37, 0x36, 0x2b, 0x7e, 0xe9, 0xc2, 0xb7, 0x1e, 0x14, 0x9c, 0xe7, 0xcf, 0xf3, 0xad,
	0x5c, 0xf0, 0xca, 0x14, 0xe8, 0x98, 0x0b, 0x89, 0x1e, 0x81, 0x3d, 0x9a, 0x74, 0xce, 0x92, 0x36,
	0xb0, 0x71, 0xc9, 0x39, 0x0e, 0x33, 0xd8, 0xe4, 0x2c, 0x63, 0x26, 0xf2, 0xc1, 0x66, 0x5c, 0xea,
	0x8b, 0x73, 0x6c, 0xad, 0xb2, 0x58, 0x24, 0x54, 0x41, 0x18, 0x97, 0xea, 0xd2, 0xd0, 0x53, 0x78,
	0x5d, 0x90, 0xef, 0xfb, 0x24, 0x91, 0x24, 0x6a, 0x27, 0x44, 0x0c, 0x88, 0x68, 0x33, 0xdc, 0x23,
	0x4e, 0x69, 0xa6, 0x2e, 0x2f, 0xda, 0x3a, 0xd1, 0xf3, 0x75, 0x62, 0xea, 0xb5, 0xb1, 0xc8, 0x89,
	0xd6, 0x38, 0xc2, 0x3d, 0x52, 0x3b, 0x87, 0xfc, 0x09, 0x91, 0x68, 0x1f, 0x8a, 0xa3, 0x26, 0xbd,
	0x55, 0x6c, 0x53, 0x76, 0xf3, 0x81, 0x8a, 0xc8, 0x3a, 0xbc, 0x3d, 0x3f, 0x22, 0x33, 0x4d, 0xd9,
	0xdc, 0x50, 0xd8, 0x55, 0x78, 0xeb, 0x1a, 0xac, 0x5f, 0x86, 0x82, 0x12, 0x47, 0xf9, 0x7f, 0x7d,
	0xa3, 0xf1, 0xa3, 0x0d, 0xa5, 0x71, 0xca, 0xd1, 0x1e, 0x2c, 0xa9, 0x78, 0xd1, 0x68, 0x14, 0xae,
	0xcd, 0x45, 0x86, 0x43, 0x16, 0x09, 0x0b, 0xb3, 0xe8, 0x20, 0x4a, 0xd0, 0x2e, 0x58, 0x5c, 0x68,
	0x0d, 0xf3, 0xc6, 0x1a, 0x45, 0x2e, 0x94, 0xc4, 0x95, 0xa1, 0x7a, 0x0a, 0x15, 0xdc, 0x97, 0x1d,
	0xc2, 0x24, 0x0d, 0xb1, 0x24, 0x51, 0x96, 0xad, 0xed, 0x85, 0xb6, 0xd9, 0x9d, 0x66, 0xaa, 0xde,
	0x9e, 0x91, 0x42, 0xc7, 0x50, 0x8d, 0xa8, 0x20, 0xa1, 0x6c, 0x0b, 0xd2, 0xe3, 0x92, 0xa8, 0xb0,
	0xc0, 0x8d, 0xc2, 0xb2, 0x9c, 0xf2, 0x03, 0x4d, 0x3f, 0x88, 0xd1, 0x1e, 0x94, 0x26, 0x52, 0xe5,
	0x1b, 0x49, 0xd9, 0x62, 0x22, 0x32, 0x9a, 0x23, 0xd6, 0xcb, 0x99, 0x23, 0xa5, 0xdb, 0xcc, 0x91,
	0x97, 0x13, 0xe0, 0x8f, 0xc1, 0x52, 0x01, 0xa6, 0x51, 0x16, 0xdf, 0x85, 0xde, 0x3f, 0xaa, 0x35,
	0x18, 0x97, 0x07, 0x51, 0x4d, 0xa4, 0xf9, 0xfa, 0x14, 0xf2, 0x69, 0x97, 0xde, 0xe6, 0x15, 0xa6,
	0xa8, 0xcd, 0xfb, 0x2a, 0x2e, 0x6b, 0xd0, 0x98, 0x1f, 0x97, 0xe9, 0xbe, 0xac, 0xfd, 0x6a, 0x40,
	0x65, 0xa6, 0x71, 0xd0, 0xe7, 0xb0, 0x3c, 0x7e, 0x0d, 0xa6, 0xa3, 0xc3, 0x5c, 0x7c, 0x74, 0x04,
	0x95, 0x31, 0x57, 0x8d, 0x8c, 0xe6, 0x8e, 0x72, 0xe2, 0xc2, 0x3b, 0xd7, 0x38, 0x99, 0x71, 0xf0,
	0xb8, 0x60, 0x1b, 0x55, 0xb3, 0xb9, 0xae, 0xa8, 0x75, 0x58, 0xb9, 0x9a, 0xea, 0xbf, 0x0a, 0x40,
	0x23, 0xc5, 0x7d, 0x46, 0x89, 0xd0, 0xc1, 0x7f, 0x5c, 0xb0, 0x8b, 0x55, 0x2b, 0x28, 0x25, 0xbc,
	0x2f, 0x42, 0xd5, 0x9b, 0xfe, 0xc3, 0x3f, 0x7e, 0xf8, 0xeb, 0x6f, 0xcb, 0xac, 0xe6, 0xe1, 0x1e,
	0xe5, 0xe9, 0x31, 0x62, 0xc1, 0xcf, 0x87, 0x97, 0xd7, 0xd6, 0x2f, 0x05, 0xa7, 0x38, 0x3c, 0x56,
	0x1f, 0x7b, 0xc7, 0xc6, 0xa9, 0xa5, 0xbf, 0xfa, 0x76, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0x07, 0x86, 0x88, 0x82, 0x0b, 0x00, 0x00,
}
