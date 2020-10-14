package logger

import "github.com/GSabadini/go-challenge/adapter/logger"

type Dummy struct{}

func (l Dummy) Infof(_ string, _ ...interface{})         {}
func (l Dummy) Warnf(_ string, _ ...interface{})         {}
func (l Dummy) Errorf(_ string, _ ...interface{})        {}
func (l Dummy) Fatalln(_ ...interface{})                 {}
func (l Dummy) WithFields(_ logger.Fields) logger.Logger { return EntryDummy{} }
func (l Dummy) WithError(_ error) logger.Logger          { return EntryDummy{} }

type EntryDummy struct{}

func (l EntryDummy) Infof(_ string, _ ...interface{})         {}
func (l EntryDummy) Warnf(_ string, _ ...interface{})         {}
func (l EntryDummy) Errorf(_ string, _ ...interface{})        {}
func (l EntryDummy) Fatalln(_ ...interface{})                 {}
func (l EntryDummy) WithFields(_ logger.Fields) logger.Logger { return EntryDummy{} }
func (l EntryDummy) WithError(_ error) logger.Logger          { return EntryDummy{} }
