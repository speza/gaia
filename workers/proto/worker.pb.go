// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

// WorkerInstance represents the identity of
// a worker instance.
type WorkerInstance struct {
	UniqueId             string   `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	WorkerSlots          int32    `protobuf:"varint,2,opt,name=worker_slots,json=workerSlots,proto3" json:"worker_slots,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerInstance) Reset()         { *m = WorkerInstance{} }
func (m *WorkerInstance) String() string { return proto.CompactTextString(m) }
func (*WorkerInstance) ProtoMessage()    {}
func (*WorkerInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{0}
}

func (m *WorkerInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerInstance.Unmarshal(m, b)
}
func (m *WorkerInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerInstance.Marshal(b, m, deterministic)
}
func (m *WorkerInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerInstance.Merge(m, src)
}
func (m *WorkerInstance) XXX_Size() int {
	return xxx_messageInfo_WorkerInstance.Size(m)
}
func (m *WorkerInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerInstance.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerInstance proto.InternalMessageInfo

func (m *WorkerInstance) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *WorkerInstance) GetWorkerSlots() int32 {
	if m != nil {
		return m.WorkerSlots
	}
	return 0
}

func (m *WorkerInstance) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// PipelineRun represents one pipeline run.
type PipelineRun struct {
	UniqueId             string   `protobuf:"bytes,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	StartDate            int64    `protobuf:"varint,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	FinishDate           int64    `protobuf:"varint,5,opt,name=finish_date,json=finishDate,proto3" json:"finish_date,omitempty"`
	ScheduleDate         int64    `protobuf:"varint,6,opt,name=schedule_date,json=scheduleDate,proto3" json:"schedule_date,omitempty"`
	PipelineId           int64    `protobuf:"varint,7,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	PipelineName         string   `protobuf:"bytes,8,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	PipelineType         string   `protobuf:"bytes,9,opt,name=pipeline_type,json=pipelineType,proto3" json:"pipeline_type,omitempty"`
	ShaSum               []byte   `protobuf:"bytes,10,opt,name=sha_sum,json=shaSum,proto3" json:"sha_sum,omitempty"`
	Jobs                 []*Job   `protobuf:"bytes,11,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Docker               bool     `protobuf:"varint,12,opt,name=docker,proto3" json:"docker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineRun) Reset()         { *m = PipelineRun{} }
func (m *PipelineRun) String() string { return proto.CompactTextString(m) }
func (*PipelineRun) ProtoMessage()    {}
func (*PipelineRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{1}
}

func (m *PipelineRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineRun.Unmarshal(m, b)
}
func (m *PipelineRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineRun.Marshal(b, m, deterministic)
}
func (m *PipelineRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineRun.Merge(m, src)
}
func (m *PipelineRun) XXX_Size() int {
	return xxx_messageInfo_PipelineRun.Size(m)
}
func (m *PipelineRun) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineRun.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineRun proto.InternalMessageInfo

func (m *PipelineRun) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *PipelineRun) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PipelineRun) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PipelineRun) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *PipelineRun) GetFinishDate() int64 {
	if m != nil {
		return m.FinishDate
	}
	return 0
}

func (m *PipelineRun) GetScheduleDate() int64 {
	if m != nil {
		return m.ScheduleDate
	}
	return 0
}

func (m *PipelineRun) GetPipelineId() int64 {
	if m != nil {
		return m.PipelineId
	}
	return 0
}

func (m *PipelineRun) GetPipelineName() string {
	if m != nil {
		return m.PipelineName
	}
	return ""
}

func (m *PipelineRun) GetPipelineType() string {
	if m != nil {
		return m.PipelineType
	}
	return ""
}

func (m *PipelineRun) GetShaSum() []byte {
	if m != nil {
		return m.ShaSum
	}
	return nil
}

func (m *PipelineRun) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *PipelineRun) GetDocker() bool {
	if m != nil {
		return m.Docker
	}
	return false
}

// PrivateKey represents a key.
type PrivateKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateKey) Reset()         { *m = PrivateKey{} }
func (m *PrivateKey) String() string { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()    {}
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{2}
}

func (m *PrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKey.Unmarshal(m, b)
}
func (m *PrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKey.Marshal(b, m, deterministic)
}
func (m *PrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKey.Merge(m, src)
}
func (m *PrivateKey) XXX_Size() int {
	return xxx_messageInfo_PrivateKey.Size(m)
}
func (m *PrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKey proto.InternalMessageInfo

func (m *PrivateKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PrivateKey) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PrivateKey) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// GitRepo is a git repo.
type GitRepo struct {
	PrivateKey           *PrivateKey `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Username             string      `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string      `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Url                  string      `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	SelectedBranch       string      `protobuf:"bytes,5,opt,name=selected_branch,json=selectedBranch,proto3" json:"selected_branch,omitempty"`
	Branches             []string    `protobuf:"bytes,6,rep,name=branches,proto3" json:"branches,omitempty"`
	Localdest            string      `protobuf:"bytes,7,opt,name=localdest,proto3" json:"localdest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GitRepo) Reset()         { *m = GitRepo{} }
func (m *GitRepo) String() string { return proto.CompactTextString(m) }
func (*GitRepo) ProtoMessage()    {}
func (*GitRepo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{3}
}

func (m *GitRepo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitRepo.Unmarshal(m, b)
}
func (m *GitRepo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitRepo.Marshal(b, m, deterministic)
}
func (m *GitRepo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitRepo.Merge(m, src)
}
func (m *GitRepo) XXX_Size() int {
	return xxx_messageInfo_GitRepo.Size(m)
}
func (m *GitRepo) XXX_DiscardUnknown() {
	xxx_messageInfo_GitRepo.DiscardUnknown(m)
}

var xxx_messageInfo_GitRepo proto.InternalMessageInfo

func (m *GitRepo) GetPrivateKey() *PrivateKey {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *GitRepo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GitRepo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *GitRepo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GitRepo) GetSelectedBranch() string {
	if m != nil {
		return m.SelectedBranch
	}
	return ""
}

func (m *GitRepo) GetBranches() []string {
	if m != nil {
		return m.Branches
	}
	return nil
}

func (m *GitRepo) GetLocaldest() string {
	if m != nil {
		return m.Localdest
	}
	return ""
}

type PipelineID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineID) Reset()         { *m = PipelineID{} }
func (m *PipelineID) String() string { return proto.CompactTextString(m) }
func (*PipelineID) ProtoMessage()    {}
func (*PipelineID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{4}
}

func (m *PipelineID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineID.Unmarshal(m, b)
}
func (m *PipelineID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineID.Marshal(b, m, deterministic)
}
func (m *PipelineID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineID.Merge(m, src)
}
func (m *PipelineID) XXX_Size() int {
	return xxx_messageInfo_PipelineID.Size(m)
}
func (m *PipelineID) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineID.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineID proto.InternalMessageInfo

func (m *PipelineID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Job represents one job from a pipeline run.
type Job struct {
	UniqueId             uint32      `protobuf:"varint,1,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	Title                string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DependsOn            []*Job      `protobuf:"bytes,4,rep,name=depends_on,json=dependsOn,proto3" json:"depends_on,omitempty"`
	Status               string      `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Args                 []*Argument `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{5}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetUniqueId() uint32 {
	if m != nil {
		return m.UniqueId
	}
	return 0
}

func (m *Job) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Job) GetDependsOn() []*Job {
	if m != nil {
		return m.DependsOn
	}
	return nil
}

func (m *Job) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Job) GetArgs() []*Argument {
	if m != nil {
		return m.Args
	}
	return nil
}

// Argument represents one argument from a job.
type Argument struct {
	Description          string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Argument) Reset()         { *m = Argument{} }
func (m *Argument) String() string { return proto.CompactTextString(m) }
func (*Argument) ProtoMessage()    {}
func (*Argument) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{6}
}

func (m *Argument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Argument.Unmarshal(m, b)
}
func (m *Argument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Argument.Marshal(b, m, deterministic)
}
func (m *Argument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Argument.Merge(m, src)
}
func (m *Argument) XXX_Size() int {
	return xxx_messageInfo_Argument.Size(m)
}
func (m *Argument) XXX_DiscardUnknown() {
	xxx_messageInfo_Argument.DiscardUnknown(m)
}

var xxx_messageInfo_Argument proto.InternalMessageInfo

func (m *Argument) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Argument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Argument) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Argument) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// LogChunk represents one chunk of a log file.
type LogChunk struct {
	RunId                int64    `protobuf:"varint,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	PipelineId           int64    `protobuf:"varint,2,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	Chunk                []byte   `protobuf:"bytes,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogChunk) Reset()         { *m = LogChunk{} }
func (m *LogChunk) String() string { return proto.CompactTextString(m) }
func (*LogChunk) ProtoMessage()    {}
func (*LogChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{7}
}

func (m *LogChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogChunk.Unmarshal(m, b)
}
func (m *LogChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogChunk.Marshal(b, m, deterministic)
}
func (m *LogChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogChunk.Merge(m, src)
}
func (m *LogChunk) XXX_Size() int {
	return xxx_messageInfo_LogChunk.Size(m)
}
func (m *LogChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_LogChunk.DiscardUnknown(m)
}

var xxx_messageInfo_LogChunk proto.InternalMessageInfo

func (m *LogChunk) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *LogChunk) GetPipelineId() int64 {
	if m != nil {
		return m.PipelineId
	}
	return 0
}

func (m *LogChunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

// FileChunk represents one chunk of a file.
type FileChunk struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileChunk) Reset()         { *m = FileChunk{} }
func (m *FileChunk) String() string { return proto.CompactTextString(m) }
func (*FileChunk) ProtoMessage()    {}
func (*FileChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4ff6184b07e587a, []int{8}
}

func (m *FileChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileChunk.Unmarshal(m, b)
}
func (m *FileChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileChunk.Marshal(b, m, deterministic)
}
func (m *FileChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileChunk.Merge(m, src)
}
func (m *FileChunk) XXX_Size() int {
	return xxx_messageInfo_FileChunk.Size(m)
}
func (m *FileChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_FileChunk.DiscardUnknown(m)
}

var xxx_messageInfo_FileChunk proto.InternalMessageInfo

func (m *FileChunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkerInstance)(nil), "protobuf.WorkerInstance")
	proto.RegisterType((*PipelineRun)(nil), "protobuf.PipelineRun")
	proto.RegisterType((*PrivateKey)(nil), "protobuf.PrivateKey")
	proto.RegisterType((*GitRepo)(nil), "protobuf.GitRepo")
	proto.RegisterType((*PipelineID)(nil), "protobuf.PipelineID")
	proto.RegisterType((*Job)(nil), "protobuf.Job")
	proto.RegisterType((*Argument)(nil), "protobuf.Argument")
	proto.RegisterType((*LogChunk)(nil), "protobuf.LogChunk")
	proto.RegisterType((*FileChunk)(nil), "protobuf.FileChunk")
}

func init() { proto.RegisterFile("worker.proto", fileDescriptor_e4ff6184b07e587a) }

var fileDescriptor_e4ff6184b07e587a = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x8f, 0x1b, 0x35,
	0x14, 0xd6, 0xe4, 0xb6, 0x99, 0x33, 0xd9, 0x05, 0x4c, 0x5a, 0x46, 0x69, 0x11, 0xd9, 0x41, 0x82,
	0x3c, 0xa0, 0x6d, 0xb5, 0xa8, 0x2f, 0x50, 0x21, 0x51, 0x16, 0x56, 0x29, 0x15, 0x54, 0xb3, 0xdc,
	0xde, 0x22, 0x27, 0x3e, 0x9d, 0x98, 0x9d, 0xd8, 0x83, 0xed, 0x69, 0x95, 0x7f, 0xc8, 0x9f, 0xe1,
	0x0d, 0xf1, 0x8c, 0x6c, 0x8f, 0x33, 0xd9, 0xa6, 0xdb, 0x87, 0x3e, 0xcd, 0x9c, 0xef, 0x7c, 0x3e,
	0x3e, 0x97, 0xcf, 0x07, 0x46, 0xaf, 0xa4, 0xba, 0x46, 0x75, 0x56, 0x29, 0x69, 0x24, 0x19, 0xba,
	0xcf, 0xb2, 0x7e, 0x31, 0xb9, 0x57, 0x48, 0x59, 0x94, 0xf8, 0x20, 0x00, 0x0f, 0x70, 0x53, 0x99,
	0xad, 0xa7, 0x65, 0x0c, 0x4e, 0x7e, 0x77, 0xc7, 0xe6, 0x42, 0x1b, 0x2a, 0x56, 0x48, 0xee, 0x41,
	0x5c, 0x0b, 0xfe, 0x57, 0x8d, 0x0b, 0xce, 0xd2, 0x68, 0x1a, 0xcd, 0xe2, 0x7c, 0xe8, 0x81, 0x39,
	0x23, 0xa7, 0xe1, 0x96, 0x85, 0x2e, 0xa5, 0xd1, 0x69, 0x67, 0x1a, 0xcd, 0xfa, 0x79, 0xe2, 0xb1,
	0x2b, 0x0b, 0x11, 0x02, 0x3d, 0x43, 0x0b, 0x9d, 0x76, 0xa7, 0xdd, 0x59, 0x9c, 0xbb, 0xff, 0xec,
	0xbf, 0x0e, 0x24, 0xcf, 0x79, 0x85, 0x25, 0x17, 0x98, 0xd7, 0xe2, 0xed, 0x77, 0x9c, 0x40, 0x87,
	0x33, 0x17, 0xb9, 0x9b, 0x77, 0x38, 0x23, 0x77, 0x61, 0xa0, 0x0d, 0x35, 0xb5, 0x0d, 0x69, 0x99,
	0x8d, 0x45, 0x3e, 0x06, 0xd0, 0x86, 0x2a, 0xb3, 0x60, 0xd4, 0x60, 0xda, 0x73, 0xfc, 0xd8, 0x21,
	0x17, 0xd4, 0x20, 0xf9, 0x04, 0x92, 0x17, 0x5c, 0x70, 0xbd, 0xf6, 0xfe, 0xbe, 0xf3, 0x83, 0x87,
	0x1c, 0xe1, 0x53, 0x38, 0xd6, 0xab, 0x35, 0xb2, 0xba, 0x44, 0x4f, 0x19, 0x38, 0xca, 0x28, 0x80,
	0x21, 0x4a, 0xd5, 0x24, 0x6e, 0x73, 0x3d, 0xf2, 0x51, 0x02, 0x34, 0x67, 0x36, 0xca, 0x8e, 0x20,
	0xe8, 0x06, 0xd3, 0xa1, 0x4b, 0x72, 0x14, 0xc0, 0x9f, 0xe8, 0x06, 0x6f, 0x90, 0xcc, 0xb6, 0xc2,
	0x34, 0xbe, 0x49, 0xfa, 0x65, 0x5b, 0x21, 0xf9, 0x08, 0x8e, 0xf4, 0x9a, 0x2e, 0x74, 0xbd, 0x49,
	0x61, 0x1a, 0xcd, 0x46, 0xf9, 0x40, 0xaf, 0xe9, 0x55, 0xbd, 0x21, 0xa7, 0xd0, 0xfb, 0x53, 0x2e,
	0x75, 0x9a, 0x4c, 0xbb, 0xb3, 0xe4, 0xfc, 0xf8, 0x2c, 0x0c, 0xf2, 0xec, 0xa9, 0x5c, 0xe6, 0xce,
	0x65, 0x7b, 0xc4, 0xe4, 0xea, 0x1a, 0x55, 0x3a, 0x9a, 0x46, 0xb3, 0x61, 0xde, 0x58, 0xd9, 0x6f,
	0x00, 0xcf, 0x15, 0x7f, 0x49, 0x0d, 0xfe, 0x88, 0x5b, 0xf2, 0x3e, 0x74, 0xaf, 0x71, 0xdb, 0x34,
	0xdc, 0xfe, 0x92, 0x09, 0x0c, 0x6b, 0x8d, 0xca, 0x25, 0xde, 0x69, 0xe6, 0xd0, 0xd8, 0xd6, 0x57,
	0x51, 0xad, 0x5f, 0x49, 0xc5, 0x9a, 0xce, 0xef, 0xec, 0xec, 0x9f, 0x08, 0x8e, 0x2e, 0xb9, 0xc9,
	0xb1, 0x92, 0xe4, 0x11, 0x24, 0x95, 0xbf, 0x63, 0x11, 0xa2, 0x27, 0xe7, 0xe3, 0x36, 0xcb, 0x36,
	0x81, 0x1c, 0xaa, 0x36, 0x99, 0x77, 0xbc, 0xda, 0x16, 0x51, 0xab, 0xd2, 0xcd, 0x3b, 0xce, 0xed,
	0x2f, 0xf9, 0x1c, 0xde, 0xd3, 0x58, 0xe2, 0xca, 0x20, 0x5b, 0x2c, 0x15, 0x15, 0xab, 0xb5, 0x9b,
	0x76, 0x9c, 0x9f, 0x04, 0xf8, 0x89, 0x43, 0x6d, 0x58, 0xef, 0x47, 0x9d, 0x0e, 0x9c, 0x3c, 0x77,
	0x36, 0xb9, 0x0f, 0x71, 0x29, 0x57, 0xb4, 0x64, 0xa8, 0x8d, 0x1b, 0x73, 0x9c, 0xb7, 0x40, 0x76,
	0x1f, 0x20, 0xe8, 0x77, 0x7e, 0xd1, 0x28, 0x34, 0x0a, 0x0a, 0xcd, 0xfe, 0x8e, 0xa0, 0xfb, 0x54,
	0x2e, 0x0f, 0x65, 0x7d, 0xbc, 0x27, 0xeb, 0x31, 0xf4, 0x0d, 0x37, 0x65, 0x28, 0xd6, 0x1b, 0x64,
	0x0a, 0x09, 0x43, 0xbd, 0x52, 0xbc, 0x32, 0x5c, 0x8a, 0xa6, 0xd8, 0x7d, 0x88, 0x7c, 0x01, 0xc0,
	0xb0, 0x42, 0xc1, 0xf4, 0x42, 0x8a, 0xb4, 0xf7, 0x26, 0x0d, 0xc4, 0x0d, 0xe1, 0x67, 0xb1, 0xf7,
	0x58, 0xfa, 0x37, 0x1e, 0xcb, 0x67, 0xd0, 0xa3, 0xaa, 0xf0, 0x65, 0x27, 0xe7, 0xa4, 0x3d, 0xff,
	0xad, 0x2a, 0xea, 0x0d, 0x0a, 0x93, 0x3b, 0x7f, 0xb6, 0x86, 0x61, 0x40, 0x5e, 0xcf, 0x2d, 0x3a,
	0xcc, 0xcd, 0xbe, 0x75, 0x2b, 0x67, 0x5f, 0x92, 0xfb, 0x0f, 0x22, 0xeb, 0xb6, 0x22, 0x1b, 0x43,
	0xff, 0x25, 0x2d, 0x6b, 0x6c, 0x66, 0xe6, 0x8d, 0xec, 0x0f, 0x18, 0x3e, 0x93, 0xc5, 0x77, 0xeb,
	0x5a, 0x5c, 0x93, 0x3b, 0x30, 0x50, 0xb5, 0x58, 0xec, 0x9a, 0xda, 0x57, 0xb5, 0x98, 0xb3, 0xd7,
	0x1f, 0x5f, 0xe7, 0xe0, 0xf1, 0x8d, 0xa1, 0xbf, 0xb2, 0x01, 0xdc, 0x6d, 0xa3, 0xdc, 0x1b, 0xd9,
	0x29, 0xc4, 0x3f, 0xf0, 0x12, 0x7d, 0xe8, 0x1d, 0x25, 0xda, 0xa3, 0x9c, 0xff, 0xdb, 0x81, 0x81,
	0xdf, 0x7b, 0xe4, 0x31, 0x1c, 0x5d, 0xa2, 0xb1, 0x06, 0x49, 0xdb, 0xb6, 0xdc, 0x5c, 0x8a, 0x93,
	0x3b, 0x7b, 0x72, 0x6e, 0xf7, 0xd8, 0xc3, 0x88, 0x7c, 0x0d, 0xf0, 0x6b, 0x65, 0xb7, 0x87, 0x0b,
	0xf0, 0x66, 0xda, 0xe4, 0xee, 0x99, 0x5f, 0xc1, 0xad, 0xf7, 0x7b, 0xbb, 0x82, 0xc9, 0x63, 0x18,
	0x5d, 0x19, 0x85, 0x74, 0xf3, 0x84, 0x0b, 0xaa, 0xb6, 0xb7, 0x1d, 0xff, 0xb0, 0x85, 0x77, 0x75,
	0x3d, 0x8c, 0xc8, 0x57, 0x00, 0xfe, 0xf4, 0x33, 0x59, 0x68, 0xb2, 0x37, 0xd2, 0xd0, 0xd6, 0xdb,
	0xee, 0x9d, 0x45, 0xe4, 0x1b, 0x80, 0x0b, 0x54, 0x58, 0x70, 0x6d, 0x50, 0xbd, 0xa5, 0xee, 0xdb,
	0x32, 0x7f, 0x04, 0x70, 0x89, 0x26, 0x6c, 0x80, 0xf1, 0x61, 0xde, 0xf3, 0x8b, 0xc9, 0x07, 0x2d,
	0xda, 0x10, 0x97, 0x03, 0x87, 0x7c, 0xf9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x36, 0xeb,
	0x40, 0xab, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	// GetWork pulls work from the primary instance.
	GetWork(ctx context.Context, in *WorkerInstance, opts ...grpc.CallOption) (Worker_GetWorkClient, error)
	// UpdateWork updates work information at the primary instance.
	UpdateWork(ctx context.Context, in *PipelineRun, opts ...grpc.CallOption) (*empty.Empty, error)
	// StreamBinary streams a pipeline binary back to a worker instance.
	StreamBinary(ctx context.Context, in *PipelineRun, opts ...grpc.CallOption) (Worker_StreamBinaryClient, error)
	// StreamLogs streams pipeline run logs to the primary instance.
	StreamLogs(ctx context.Context, opts ...grpc.CallOption) (Worker_StreamLogsClient, error)
	// Deregister deregister a registered worker from the primary instance.
	Deregister(ctx context.Context, in *WorkerInstance, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetGitRepo returns git repo information to the worker based on a pipeline name.
	GetGitRepo(ctx context.Context, in *PipelineID, opts ...grpc.CallOption) (*GitRepo, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) GetWork(ctx context.Context, in *WorkerInstance, opts ...grpc.CallOption) (Worker_GetWorkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[0], "/protobuf.Worker/GetWork", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerGetWorkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Worker_GetWorkClient interface {
	Recv() (*PipelineRun, error)
	grpc.ClientStream
}

type workerGetWorkClient struct {
	grpc.ClientStream
}

func (x *workerGetWorkClient) Recv() (*PipelineRun, error) {
	m := new(PipelineRun)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) UpdateWork(ctx context.Context, in *PipelineRun, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Worker/UpdateWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) StreamBinary(ctx context.Context, in *PipelineRun, opts ...grpc.CallOption) (Worker_StreamBinaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[1], "/protobuf.Worker/StreamBinary", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerStreamBinaryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Worker_StreamBinaryClient interface {
	Recv() (*FileChunk, error)
	grpc.ClientStream
}

type workerStreamBinaryClient struct {
	grpc.ClientStream
}

func (x *workerStreamBinaryClient) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) StreamLogs(ctx context.Context, opts ...grpc.CallOption) (Worker_StreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[2], "/protobuf.Worker/StreamLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerStreamLogsClient{stream}
	return x, nil
}

type Worker_StreamLogsClient interface {
	Send(*LogChunk) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type workerStreamLogsClient struct {
	grpc.ClientStream
}

func (x *workerStreamLogsClient) Send(m *LogChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerStreamLogsClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) Deregister(ctx context.Context, in *WorkerInstance, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Worker/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetGitRepo(ctx context.Context, in *PipelineID, opts ...grpc.CallOption) (*GitRepo, error) {
	out := new(GitRepo)
	err := c.cc.Invoke(ctx, "/protobuf.Worker/GetGitRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	// GetWork pulls work from the primary instance.
	GetWork(*WorkerInstance, Worker_GetWorkServer) error
	// UpdateWork updates work information at the primary instance.
	UpdateWork(context.Context, *PipelineRun) (*empty.Empty, error)
	// StreamBinary streams a pipeline binary back to a worker instance.
	StreamBinary(*PipelineRun, Worker_StreamBinaryServer) error
	// StreamLogs streams pipeline run logs to the primary instance.
	StreamLogs(Worker_StreamLogsServer) error
	// Deregister deregister a registered worker from the primary instance.
	Deregister(context.Context, *WorkerInstance) (*empty.Empty, error)
	// GetGitRepo returns git repo information to the worker based on a pipeline name.
	GetGitRepo(context.Context, *PipelineID) (*GitRepo, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_GetWork_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WorkerInstance)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerServer).GetWork(m, &workerGetWorkServer{stream})
}

type Worker_GetWorkServer interface {
	Send(*PipelineRun) error
	grpc.ServerStream
}

type workerGetWorkServer struct {
	grpc.ServerStream
}

func (x *workerGetWorkServer) Send(m *PipelineRun) error {
	return x.ServerStream.SendMsg(m)
}

func _Worker_UpdateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineRun)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).UpdateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Worker/UpdateWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).UpdateWork(ctx, req.(*PipelineRun))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_StreamBinary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PipelineRun)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerServer).StreamBinary(m, &workerStreamBinaryServer{stream})
}

type Worker_StreamBinaryServer interface {
	Send(*FileChunk) error
	grpc.ServerStream
}

type workerStreamBinaryServer struct {
	grpc.ServerStream
}

func (x *workerStreamBinaryServer) Send(m *FileChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Worker_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).StreamLogs(&workerStreamLogsServer{stream})
}

type Worker_StreamLogsServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*LogChunk, error)
	grpc.ServerStream
}

type workerStreamLogsServer struct {
	grpc.ServerStream
}

func (x *workerStreamLogsServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerStreamLogsServer) Recv() (*LogChunk, error) {
	m := new(LogChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Worker/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Deregister(ctx, req.(*WorkerInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetGitRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetGitRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Worker/GetGitRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetGitRepo(ctx, req.(*PipelineID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateWork",
			Handler:    _Worker_UpdateWork_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Worker_Deregister_Handler,
		},
		{
			MethodName: "GetGitRepo",
			Handler:    _Worker_GetGitRepo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWork",
			Handler:       _Worker_GetWork_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamBinary",
			Handler:       _Worker_StreamBinary_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamLogs",
			Handler:       _Worker_StreamLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "worker.proto",
}
