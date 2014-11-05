/**
 * Copyright 2014 Acquia, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package statsgod_test

import (
	. "github.com/acquia/statsgod/statsgod"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Logger", func() {

	Describe("Creating a silent logger", func() {
		It("should be a legit Logger struct", func() {
			logger := *CreateLogger(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
			Expect(logger.Trace).ShouldNot(Equal(nil))
			Expect(logger.Info).ShouldNot(Equal(nil))
			Expect(logger.Warning).ShouldNot(Equal(nil))
			Expect(logger.Error).ShouldNot(Equal(nil))
		})
	})

})