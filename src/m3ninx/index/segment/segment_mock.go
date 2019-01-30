// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/index/segment (interfaces: Segment,MutableSegment,Builder)

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package segment is a generated GoMock package.
package segment

import (
	"reflect"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/index"
	"github.com/m3db/m3/src/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockSegment is a mock of Segment interface
type MockSegment struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentMockRecorder
}

// MockSegmentMockRecorder is the mock recorder for MockSegment
type MockSegmentMockRecorder struct {
	mock *MockSegment
}

// NewMockSegment creates a new mock instance
func NewMockSegment(ctrl *gomock.Controller) *MockSegment {
	mock := &MockSegment{ctrl: ctrl}
	mock.recorder = &MockSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSegment) EXPECT() *MockSegmentMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSegment) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSegmentMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSegment)(nil).Close))
}

// ContainsID mocks base method
func (m *MockSegment) ContainsID(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsID indicates an expected call of ContainsID
func (mr *MockSegmentMockRecorder) ContainsID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsID", reflect.TypeOf((*MockSegment)(nil).ContainsID), arg0)
}

// FieldsIterable mocks base method
func (m *MockSegment) FieldsIterable() FieldsIterable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldsIterable")
	ret0, _ := ret[0].(FieldsIterable)
	return ret0
}

// FieldsIterable indicates an expected call of FieldsIterable
func (mr *MockSegmentMockRecorder) FieldsIterable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldsIterable", reflect.TypeOf((*MockSegment)(nil).FieldsIterable))
}

// Reader mocks base method
func (m *MockSegment) Reader() (index.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(index.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader
func (mr *MockSegmentMockRecorder) Reader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockSegment)(nil).Reader))
}

// Size mocks base method
func (m *MockSegment) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockSegmentMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockSegment)(nil).Size))
}

// TermsIterable mocks base method
func (m *MockSegment) TermsIterable() TermsIterable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TermsIterable")
	ret0, _ := ret[0].(TermsIterable)
	return ret0
}

// TermsIterable indicates an expected call of TermsIterable
func (mr *MockSegmentMockRecorder) TermsIterable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TermsIterable", reflect.TypeOf((*MockSegment)(nil).TermsIterable))
}

// MockMutableSegment is a mock of MutableSegment interface
type MockMutableSegment struct {
	ctrl     *gomock.Controller
	recorder *MockMutableSegmentMockRecorder
}

// MockMutableSegmentMockRecorder is the mock recorder for MockMutableSegment
type MockMutableSegmentMockRecorder struct {
	mock *MockMutableSegment
}

// NewMockMutableSegment creates a new mock instance
func NewMockMutableSegment(ctrl *gomock.Controller) *MockMutableSegment {
	mock := &MockMutableSegment{ctrl: ctrl}
	mock.recorder = &MockMutableSegmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutableSegment) EXPECT() *MockMutableSegmentMockRecorder {
	return m.recorder
}

// AllDocs mocks base method
func (m *MockMutableSegment) AllDocs() (index.IDDocIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllDocs")
	ret0, _ := ret[0].(index.IDDocIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDocs indicates an expected call of AllDocs
func (mr *MockMutableSegmentMockRecorder) AllDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDocs", reflect.TypeOf((*MockMutableSegment)(nil).AllDocs))
}

// Close mocks base method
func (m *MockMutableSegment) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMutableSegmentMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMutableSegment)(nil).Close))
}

// ContainsID mocks base method
func (m *MockMutableSegment) ContainsID(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsID", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsID indicates an expected call of ContainsID
func (mr *MockMutableSegmentMockRecorder) ContainsID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsID", reflect.TypeOf((*MockMutableSegment)(nil).ContainsID), arg0)
}

// Docs mocks base method
func (m *MockMutableSegment) Docs() []doc.Document {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Docs")
	ret0, _ := ret[0].([]doc.Document)
	return ret0
}

// Docs indicates an expected call of Docs
func (mr *MockMutableSegmentMockRecorder) Docs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Docs", reflect.TypeOf((*MockMutableSegment)(nil).Docs))
}

// Fields mocks base method
func (m *MockMutableSegment) Fields() (FieldsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].(FieldsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockMutableSegmentMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockMutableSegment)(nil).Fields))
}

// FieldsIterable mocks base method
func (m *MockMutableSegment) FieldsIterable() FieldsIterable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldsIterable")
	ret0, _ := ret[0].(FieldsIterable)
	return ret0
}

// FieldsIterable indicates an expected call of FieldsIterable
func (mr *MockMutableSegmentMockRecorder) FieldsIterable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldsIterable", reflect.TypeOf((*MockMutableSegment)(nil).FieldsIterable))
}

// Insert mocks base method
func (m *MockMutableSegment) Insert(arg0 doc.Document) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockMutableSegmentMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMutableSegment)(nil).Insert), arg0)
}

// InsertBatch mocks base method
func (m *MockMutableSegment) InsertBatch(arg0 index.Batch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBatch indicates an expected call of InsertBatch
func (mr *MockMutableSegmentMockRecorder) InsertBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBatch", reflect.TypeOf((*MockMutableSegment)(nil).InsertBatch), arg0)
}

// IsSealed mocks base method
func (m *MockMutableSegment) IsSealed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSealed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSealed indicates an expected call of IsSealed
func (mr *MockMutableSegmentMockRecorder) IsSealed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSealed", reflect.TypeOf((*MockMutableSegment)(nil).IsSealed))
}

// Offset mocks base method
func (m *MockMutableSegment) Offset() postings.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset")
	ret0, _ := ret[0].(postings.ID)
	return ret0
}

// Offset indicates an expected call of Offset
func (mr *MockMutableSegmentMockRecorder) Offset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockMutableSegment)(nil).Offset))
}

// Reader mocks base method
func (m *MockMutableSegment) Reader() (index.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(index.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader
func (mr *MockMutableSegmentMockRecorder) Reader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockMutableSegment)(nil).Reader))
}

// Reset mocks base method
func (m *MockMutableSegment) Reset(arg0 postings.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockMutableSegmentMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockMutableSegment)(nil).Reset), arg0)
}

// Seal mocks base method
func (m *MockMutableSegment) Seal() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal")
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockMutableSegmentMockRecorder) Seal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockMutableSegment)(nil).Seal))
}

// Size mocks base method
func (m *MockMutableSegment) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockMutableSegmentMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockMutableSegment)(nil).Size))
}

// Terms mocks base method
func (m *MockMutableSegment) Terms(arg0 []byte) (TermsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].(TermsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockMutableSegmentMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockMutableSegment)(nil).Terms), arg0)
}

// TermsIterable mocks base method
func (m *MockMutableSegment) TermsIterable() TermsIterable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TermsIterable")
	ret0, _ := ret[0].(TermsIterable)
	return ret0
}

// TermsIterable indicates an expected call of TermsIterable
func (mr *MockMutableSegmentMockRecorder) TermsIterable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TermsIterable", reflect.TypeOf((*MockMutableSegment)(nil).TermsIterable))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AllDocs mocks base method
func (m *MockBuilder) AllDocs() (index.IDDocIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllDocs")
	ret0, _ := ret[0].(index.IDDocIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDocs indicates an expected call of AllDocs
func (mr *MockBuilderMockRecorder) AllDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDocs", reflect.TypeOf((*MockBuilder)(nil).AllDocs))
}

// Docs mocks base method
func (m *MockBuilder) Docs() []doc.Document {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Docs")
	ret0, _ := ret[0].([]doc.Document)
	return ret0
}

// Docs indicates an expected call of Docs
func (mr *MockBuilderMockRecorder) Docs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Docs", reflect.TypeOf((*MockBuilder)(nil).Docs))
}

// Fields mocks base method
func (m *MockBuilder) Fields() (FieldsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].(FieldsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields
func (mr *MockBuilderMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockBuilder)(nil).Fields))
}

// Reset mocks base method
func (m *MockBuilder) Reset(arg0 postings.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockBuilderMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockBuilder)(nil).Reset), arg0)
}

// Terms mocks base method
func (m *MockBuilder) Terms(arg0 []byte) (TermsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Terms", arg0)
	ret0, _ := ret[0].(TermsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Terms indicates an expected call of Terms
func (mr *MockBuilderMockRecorder) Terms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terms", reflect.TypeOf((*MockBuilder)(nil).Terms), arg0)
}
