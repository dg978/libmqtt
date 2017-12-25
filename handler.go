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

// ConHandler handler the bad connect result
type ConHandler func(server string, code ConAckCode)

// PubHandler handler bad topic pub
type PubHandler func(topic string, code PubAckCode)

// SubHandler handler bad topic sub
type SubHandler func(topic string, qos QosLevel, msg []byte)

// UnSubHandler handler bad topic unSub
type UnSubHandler func(topic string, code SubAckCode)
