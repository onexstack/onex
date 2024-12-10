// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.23.4
// source: nightwatch/v1/job.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Job 名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Job ID
	JobID string `protobuf:"bytes,2,opt,name=jobID,proto3" json:"jobID,omitempty"`
	// 创建人
	UserID string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	// CronJob ID
	CronJobID string `protobuf:"bytes,4,opt,name=cronJobID,proto3" json:"cronJobID,omitempty"`
	// Job 作用域
	Scope string `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	// Job 描述
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Watcher     string `protobuf:"bytes,7,opt,name=watcher,proto3" json:"watcher,omitempty"`
	// Job 参数
	Params *JobParams `protobuf:"bytes,8,opt,name=params,proto3" json:"params,omitempty"`
	// Job 执行结果
	Results *JobResults `protobuf:"bytes,9,opt,name=results,proto3" json:"results,omitempty"`
	// Job 状态：Pending、Running、Succeeded、Failed
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// Job 状态：Pending、Running、Succeeded、Failed
	Conditions []*JobCondition `protobuf:"bytes,11,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Job 开始时间
	StartedAt int64 `protobuf:"varint,12,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	// Job 结束时间
	EndedAt int64 `protobuf:"varint,13,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	// 创建时间
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// 更新时间
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *Job) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Job) GetCronJobID() string {
	if x != nil {
		return x.CronJobID
	}
	return ""
}

func (x *Job) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Job) GetWatcher() string {
	if x != nil {
		return x.Watcher
	}
	return ""
}

func (x *Job) GetParams() *JobParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Job) GetResults() *JobResults {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *Job) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Job) GetConditions() []*JobCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Job) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *Job) GetEndedAt() int64 {
	if x != nil {
		return x.EndedAt
	}
	return 0
}

func (x *Job) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Job) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type JobParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Train *TrainParams `protobuf:"bytes,1,opt,name=train,proto3" json:"train,omitempty"`
}

func (x *JobParams) Reset() {
	*x = JobParams{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobParams) ProtoMessage() {}

func (x *JobParams) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobParams.ProtoReflect.Descriptor instead.
func (*JobParams) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobParams) GetTrain() *TrainParams {
	if x != nil {
		return x.Train
	}
	return nil
}

type TrainParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdempotentExecution int64 `protobuf:"varint,1,opt,name=idempotentExecution,proto3" json:"idempotentExecution,omitempty"`
	JobTimeout          int64 `protobuf:"varint,2,opt,name=jobTimeout,proto3" json:"jobTimeout,omitempty"`
	BatchSize           int64 `protobuf:"varint,3,opt,name=batchSize,proto3" json:"batchSize,omitempty"`
}

func (x *TrainParams) Reset() {
	*x = TrainParams{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrainParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainParams) ProtoMessage() {}

func (x *TrainParams) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainParams.ProtoReflect.Descriptor instead.
func (*TrainParams) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{2}
}

func (x *TrainParams) GetIdempotentExecution() int64 {
	if x != nil {
		return x.IdempotentExecution
	}
	return 0
}

func (x *TrainParams) GetJobTimeout() int64 {
	if x != nil {
		return x.JobTimeout
	}
	return 0
}

func (x *TrainParams) GetBatchSize() int64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type JobResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Train *TrainResults `protobuf:"bytes,1,opt,name=train,proto3" json:"train,omitempty"`
}

func (x *JobResults) Reset() {
	*x = JobResults{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResults) ProtoMessage() {}

func (x *JobResults) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResults.ProtoReflect.Descriptor instead.
func (*JobResults) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobResults) GetTrain() *TrainResults {
	if x != nil {
		return x.Train
	}
	return nil
}

type TrainResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmbeddedDataPath *string `protobuf:"bytes,1,opt,name=embeddedDataPath,proto3,oneof" json:"embeddedDataPath,omitempty"`
	TaskID           *string `protobuf:"bytes,2,opt,name=taskID,proto3,oneof" json:"taskID,omitempty"`
	DataPath         *string `protobuf:"bytes,4,opt,name=dataPath,proto3,oneof" json:"dataPath,omitempty"`
	ResultPath       *string `protobuf:"bytes,3,opt,name=resultPath,proto3,oneof" json:"resultPath,omitempty"`
}

func (x *TrainResults) Reset() {
	*x = TrainResults{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrainResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainResults) ProtoMessage() {}

func (x *TrainResults) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainResults.ProtoReflect.Descriptor instead.
func (*TrainResults) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{4}
}

func (x *TrainResults) GetEmbeddedDataPath() string {
	if x != nil && x.EmbeddedDataPath != nil {
		return *x.EmbeddedDataPath
	}
	return ""
}

func (x *TrainResults) GetTaskID() string {
	if x != nil && x.TaskID != nil {
		return *x.TaskID
	}
	return ""
}

func (x *TrainResults) GetDataPath() string {
	if x != nil && x.DataPath != nil {
		return *x.DataPath
	}
	return ""
}

func (x *TrainResults) GetResultPath() string {
	if x != nil && x.ResultPath != nil {
		return *x.ResultPath
	}
	return ""
}

type JobCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the condition (e.g., "Downloaed", "Embedded", "Evaluated")
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status of the condition (True, False, Unknown)
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Message provides additional information when the condition is false
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `protobuf:"bytes,4,opt,name=lastTransitionTime,proto3" json:"lastTransitionTime,omitempty"`
}

func (x *JobCondition) Reset() {
	*x = JobCondition{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCondition) ProtoMessage() {}

func (x *JobCondition) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCondition.ProtoReflect.Descriptor instead.
func (*JobCondition) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *JobCondition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *JobCondition) GetLastTransitionTime() string {
	if x != nil {
		return x.LastTransitionTime
	}
	return ""
}

type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{6}
}

func (x *CreateJobRequest) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type CreateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobResponse.ProtoReflect.Descriptor instead.
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{7}
}

func (x *CreateJobResponse) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID       string      `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Name        *string     `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string     `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Params      *JobParams  `protobuf:"bytes,4,opt,name=params,proto3,oneof" json:"params,omitempty"`
	Results     *JobResults `protobuf:"bytes,5,opt,name=results,proto3,oneof" json:"results,omitempty"`
	Status      *string     `protobuf:"bytes,6,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateJobRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *UpdateJobRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateJobRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateJobRequest) GetParams() *JobParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *UpdateJobRequest) GetResults() *JobResults {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *UpdateJobRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type UpdateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobResponse) Reset() {
	*x = UpdateJobResponse{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobResponse) ProtoMessage() {}

func (x *UpdateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobResponse) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{9}
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{10}
}

func (x *GetJobRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

type GetJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *GetJobResponse) Reset() {
	*x = GetJobResponse{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobResponse) ProtoMessage() {}

func (x *GetJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobResponse.ProtoReflect.Descriptor instead.
func (*GetJobResponse) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{11}
}

func (x *GetJobResponse) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type ListJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListJobRequest) Reset() {
	*x = ListJobRequest{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobRequest) ProtoMessage() {}

func (x *ListJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobRequest.ProtoReflect.Descriptor instead.
func (*ListJobRequest) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{12}
}

func (x *ListJobRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListJobRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64  `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Jobs       []*Job `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ListJobResponse) Reset() {
	*x = ListJobResponse{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobResponse) ProtoMessage() {}

func (x *ListJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobResponse.ProtoReflect.Descriptor instead.
func (*ListJobResponse) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{13}
}

func (x *ListJobResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListJobResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobIDs []string `protobuf:"bytes,1,rep,name=jobIDs,proto3" json:"jobIDs,omitempty"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteJobRequest) GetJobIDs() []string {
	if x != nil {
		return x.JobIDs
	}
	return nil
}

type DeleteJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobResponse) Reset() {
	*x = DeleteJobResponse{}
	mi := &file_nightwatch_v1_job_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobResponse) ProtoMessage() {}

func (x *DeleteJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nightwatch_v1_job_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobResponse) Descriptor() ([]byte, []int) {
	return file_nightwatch_v1_job_proto_rawDescGZIP(), []int{15}
}

var File_nightwatch_v1_job_proto protoreflect.FileDescriptor

var file_nightwatch_v1_job_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x04, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a,
	0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x33, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x7d, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x64, 0x65, 0x6d, 0x70,
	0x6f, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6a,
	0x6f, 0x62, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x35, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0xde,
	0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x2f, 0x0a, 0x10, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x23,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x84, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f,
	0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x44, 0x22, 0x9d, 0x02, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x02, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x48, 0x03, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x22, 0x2c, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x3e, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4f, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x2a, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nightwatch_v1_job_proto_rawDescOnce sync.Once
	file_nightwatch_v1_job_proto_rawDescData = file_nightwatch_v1_job_proto_rawDesc
)

func file_nightwatch_v1_job_proto_rawDescGZIP() []byte {
	file_nightwatch_v1_job_proto_rawDescOnce.Do(func() {
		file_nightwatch_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_nightwatch_v1_job_proto_rawDescData)
	})
	return file_nightwatch_v1_job_proto_rawDescData
}

var file_nightwatch_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_nightwatch_v1_job_proto_goTypes = []any{
	(*Job)(nil),                   // 0: job.Job
	(*JobParams)(nil),             // 1: job.JobParams
	(*TrainParams)(nil),           // 2: job.TrainParams
	(*JobResults)(nil),            // 3: job.JobResults
	(*TrainResults)(nil),          // 4: job.TrainResults
	(*JobCondition)(nil),          // 5: job.JobCondition
	(*CreateJobRequest)(nil),      // 6: job.CreateJobRequest
	(*CreateJobResponse)(nil),     // 7: job.CreateJobResponse
	(*UpdateJobRequest)(nil),      // 8: job.UpdateJobRequest
	(*UpdateJobResponse)(nil),     // 9: job.UpdateJobResponse
	(*GetJobRequest)(nil),         // 10: job.GetJobRequest
	(*GetJobResponse)(nil),        // 11: job.GetJobResponse
	(*ListJobRequest)(nil),        // 12: job.ListJobRequest
	(*ListJobResponse)(nil),       // 13: job.ListJobResponse
	(*DeleteJobRequest)(nil),      // 14: job.DeleteJobRequest
	(*DeleteJobResponse)(nil),     // 15: job.DeleteJobResponse
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_nightwatch_v1_job_proto_depIdxs = []int32{
	1,  // 0: job.Job.params:type_name -> job.JobParams
	3,  // 1: job.Job.results:type_name -> job.JobResults
	5,  // 2: job.Job.conditions:type_name -> job.JobCondition
	16, // 3: job.Job.createdAt:type_name -> google.protobuf.Timestamp
	16, // 4: job.Job.updatedAt:type_name -> google.protobuf.Timestamp
	2,  // 5: job.JobParams.train:type_name -> job.TrainParams
	4,  // 6: job.JobResults.train:type_name -> job.TrainResults
	0,  // 7: job.CreateJobRequest.job:type_name -> job.Job
	1,  // 8: job.UpdateJobRequest.params:type_name -> job.JobParams
	3,  // 9: job.UpdateJobRequest.results:type_name -> job.JobResults
	0,  // 10: job.GetJobResponse.job:type_name -> job.Job
	0,  // 11: job.ListJobResponse.jobs:type_name -> job.Job
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_nightwatch_v1_job_proto_init() }
func file_nightwatch_v1_job_proto_init() {
	if File_nightwatch_v1_job_proto != nil {
		return
	}
	file_nightwatch_v1_job_proto_msgTypes[4].OneofWrappers = []any{}
	file_nightwatch_v1_job_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nightwatch_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nightwatch_v1_job_proto_goTypes,
		DependencyIndexes: file_nightwatch_v1_job_proto_depIdxs,
		MessageInfos:      file_nightwatch_v1_job_proto_msgTypes,
	}.Build()
	File_nightwatch_v1_job_proto = out.File
	file_nightwatch_v1_job_proto_rawDesc = nil
	file_nightwatch_v1_job_proto_goTypes = nil
	file_nightwatch_v1_job_proto_depIdxs = nil
}
