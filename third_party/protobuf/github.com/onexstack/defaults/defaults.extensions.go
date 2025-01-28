// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package defaults

import (
	"errors"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	"google.golang.org/protobuf/proto"
	reflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Apply(m proto.Message) {
	if m == nil {
		return
	}
	mref := m.ProtoReflect()
	typd := mref.Type().Descriptor()
	opts := typd.Options()
	disabled := proto.GetExtension(opts, E_Disabled)
	if disabled.(bool) {
		return
	}
	ignored := proto.GetExtension(opts, E_Ignored)
	if ignored.(bool) {
		return
	}
	fields := typd.Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		if f.IsList() || f.IsMap() {
			continue
		}
		if mref.Has(f) {
			continue
		}
		v := proto.GetExtension(f.Options(), E_Value)
		if v == nil {
			continue
		}
		fd, ok := v.(*FieldDefaults)
		if !ok {
			// wtf ???
			continue
		}
		name := f.Name()
		if oo := f.ContainingOneof(); oo != nil && !oo.IsSynthetic() {
			v := proto.GetExtension(oo.Options(), E_Oneof)
			oon, ok := v.(string)
			if !ok {
				// wtf ??
				continue
			}
			if oon != string(name) {
				continue
			}
		}
		switch f.Kind() {
		case reflect.BoolKind:
			if _, ok := fd.GetType().(*FieldDefaults_Bool); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetBool()))
		case reflect.EnumKind:
			if _, ok := fd.GetType().(*FieldDefaults_Enum); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(reflect.EnumNumber(fd.GetEnum())))
		case reflect.Int32Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Int32); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetInt32()))
		case reflect.Sint32Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Sint32); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetSint32()))
		case reflect.Uint32Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Uint32); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetUint32()))
		case reflect.Int64Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Int64); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetInt64()))
		case reflect.Sint64Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Sint64); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetSint64()))
		case reflect.Uint64Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Uint64); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetUint64()))
		case reflect.Sfixed32Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Sfixed32); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetSfixed32()))
		case reflect.Fixed32Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Fixed32); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetFixed32()))
		case reflect.FloatKind:
			if _, ok := fd.GetType().(*FieldDefaults_Float); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetFloat()))
		case reflect.Sfixed64Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Sfixed64); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetSfixed64()))
		case reflect.Fixed64Kind:
			if _, ok := fd.GetType().(*FieldDefaults_Fixed64); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetFixed64()))
		case reflect.DoubleKind:
			if _, ok := fd.GetType().(*FieldDefaults_Double); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetDouble()))
		case reflect.StringKind:
			if _, ok := fd.GetType().(*FieldDefaults_String_); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetString_()))
		case reflect.BytesKind:
			if _, ok := fd.GetType().(*FieldDefaults_Bytes); !ok {
				continue
			}
			mref.Set(f, reflect.ValueOf(fd.GetBytes()))
		case reflect.MessageKind:
			m := fd.GetMessage()
			switch mref.Get(f).Message().Interface().(type) {
			case *durationpb.Duration:
				if _, ok := fd.GetType().(*FieldDefaults_Duration); !ok {
					continue
				}
				if d, err := model.ParseDuration(fd.GetDuration()); err == nil {
					mref.Set(f, reflect.ValueOf(durationpb.New(time.Duration(d)).ProtoReflect()))
				}
			case *timestamppb.Timestamp:
				if _, ok := fd.GetType().(*FieldDefaults_Timestamp); !ok {
					continue
				}
				ts := fd.GetTimestamp()
				if strings.ToLower(ts) == "now" {
					mref.Set(f, reflect.ValueOf(timestamppb.Now().ProtoReflect()))
					continue
				}
				if t, err := parseTime(ts); err == nil {
					mref.Set(f, reflect.ValueOf(timestamppb.New(t).ProtoReflect()))
				}
			case *wrapperspb.DoubleValue:
				if _, ok := fd.GetType().(*FieldDefaults_Double); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Double(fd.GetDouble()).ProtoReflect()))
			case *wrapperspb.FloatValue:
				if _, ok := fd.GetType().(*FieldDefaults_Float); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Float(fd.GetFloat()).ProtoReflect()))
			case *wrapperspb.Int64Value:
				if _, ok := fd.GetType().(*FieldDefaults_Int64); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Int64(fd.GetInt64()).ProtoReflect()))
			case *wrapperspb.UInt64Value:
				if _, ok := fd.GetType().(*FieldDefaults_Uint64); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.UInt64(fd.GetUint64()).ProtoReflect()))
			case *wrapperspb.Int32Value:
				if _, ok := fd.GetType().(*FieldDefaults_Int32); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Int32(fd.GetInt32()).ProtoReflect()))
			case *wrapperspb.UInt32Value:
				if _, ok := fd.GetType().(*FieldDefaults_Uint32); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.UInt32(fd.GetUint32()).ProtoReflect()))
			case *wrapperspb.BoolValue:
				if _, ok := fd.GetType().(*FieldDefaults_Bool); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Bool(fd.GetBool()).ProtoReflect()))
			case *wrapperspb.StringValue:
				if _, ok := fd.GetType().(*FieldDefaults_String_); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.String(fd.GetString_()).ProtoReflect()))
			case *wrapperspb.BytesValue:
				if _, ok := fd.GetType().(*FieldDefaults_Bytes); !ok {
					continue
				}
				mref.Set(f, reflect.ValueOf(wrapperspb.Bytes(fd.GetBytes()).ProtoReflect()))
			default:
				if _, ok := fd.GetType().(*FieldDefaults_Message); !ok {
					continue
				}
				if !mref.Get(f).Message().IsValid() {
					if !m.GetInitialize() {
						continue
					}
					mref.Set(f, reflect.ValueOf(mref.Get(f).Message().New()))
				}
				if !m.GetDefaults() {
					continue
				}
				Apply(mref.Get(f).Message().Interface())
			}
		case reflect.GroupKind:
		}
	}
}

func parseTime(s string) (time.Time, error) {
	for _, format := range []string{
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
	} {
		t, err := time.Parse(format, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("cannot parse timestamp, timestamp supported format: RFC822 / RFC822Z / RFC850 / RFC1123 / RFC1123Z / RFC3339")
}
