/*
Copyright © 2021 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	v1mock "github.com/rancher/elemental-toolkit/pkg/mocks"
	v1 "github.com/rancher/elemental-toolkit/pkg/types/v1"
)

var _ = Describe("Runner", Label("types", "runner"), func() {
	It("Runs commands on the real Runner", func() {
		r := v1.RealRunner{}
		_, err := r.Run("pwd")
		Expect(err).To(BeNil())
	})
	It("Runs commands on the fake runner", func() {
		r := v1mock.NewFakeRunner()
		_, err := r.Run("pwd")
		Expect(err).To(BeNil())
	})
	It("Sets and gets the logger on the fake runner", func() {
		r := v1mock.NewFakeRunner()
		Expect(r.GetLogger()).To(BeNil())
		logger := v1.NewNullLogger()
		r.SetLogger(logger)
		Expect(r.GetLogger()).To(Equal(logger))
	})
	It("Sets and gets the logger on the real runner", func() {
		r := v1.RealRunner{}
		Expect(r.GetLogger()).To(BeNil())
		logger := v1.NewNullLogger()
		r.SetLogger(logger)
		Expect(r.GetLogger()).To(Equal(logger))
	})

	It("logs the command when on debug", func() {
		memLog := &bytes.Buffer{}
		logger := v1.NewBufferLogger(memLog)
		logger.SetLevel(logrus.DebugLevel)
		r := v1.RealRunner{Logger: logger}
		_, err := r.Run("echo", "-n", "Some message")
		Expect(err).To(BeNil())
		Expect(memLog.String()).To(ContainSubstring("echo -n Some message"))
	})
	It("logs when command is not found", func() {
		memLog := &bytes.Buffer{}
		logger := v1.NewBufferLogger(memLog)
		r := v1.RealRunner{Logger: logger}
		_, err := r.Run("IAmMissing")
		Expect(err).NotTo(BeNil())
		Expect(memLog.String()).To(ContainSubstring("not found"))
	})
	It("returns false if command does not exists", func() {
		r := v1.RealRunner{}
		exists := r.CommandExists("THISCOMMANDSHOULDNOTBETHERECOMEON")
		Expect(exists).To(BeFalse())
	})
	It("returns true if command exists", func() {
		r := v1.RealRunner{}
		exists := r.CommandExists("true")
		Expect(exists).To(BeTrue())
	})
})
