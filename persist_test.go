/*
 * Copyright GoIIoT (https://github.com/goiiot)
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

package libmqtt

import "testing"

func TestMemPersist(t *testing.T) {
	p := NewMemPersist(&PersistStrategy{
		MaxCount:         1,
		DropOnExceed:     true,
		DuplicateReplace: false,
	})

	p.Store("foo", nil)
	p.Store("foo", &SubscribePacket{})
	p.Store("bar", nil)

	if p.count != 1 {
		t.Log("count =", p.count)
		t.Fail()
	}

	if v, ok := p.Load("foo"); !ok || v != nil {
		t.Log("pkt =", v)
		t.Fail()
	}
}

func TestFilePersist(t *testing.T) {

}