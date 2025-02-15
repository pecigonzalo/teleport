/*
Copyright 2022 Gravitational, Inc.

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

package protocol

import (
	"bytes"
	"io"
	"net"
	"testing"
	"testing/iotest"

	"github.com/gravitational/trace"

	"github.com/stretchr/testify/require"
)

var (
	sampleOKPacket = &OK{
		packet: packet{
			bytes: []byte{
				0x03, 0x00, 0x00, 0x00, // header
				0x00, // type
				0x00, 0x00,
			},
		},
	}

	sampleQueryPacket = &Query{
		packet: packet{
			bytes: []byte{
				0x09, 0x00, 0x00, 0x00, // header
				0x03,                                           // type
				0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x20, 0x31, // query
			},
		},
		query: "select 1",
	}

	sampleErrorPacket = &Error{
		packet: packet{
			bytes: []byte{
				0x09, 0x00, 0x00, 0x00, // header
				0xff,       // type
				0x51, 0x04, // error code
				0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, // message
			},
		},
		message: "denied",
	}

	sampleErrorWithSQLStatePacket = &Error{
		packet: packet{
			bytes: []byte{
				0x0f, 0x00, 0x00, 0x00, // header
				0xff,       // type
				0x51, 0x04, // error code
				0x23,                         // marker #
				0x48, 0x59, 0x30, 0x30, 0x30, // state - HY000
				0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, // message
			},
		},
		message: "denied",
	}

	sampleQuitPacket = &Quit{
		packet: packet{
			bytes: []byte{
				0x01, 0x00, 0x00, 0x00, // header
				0x01, //type
			},
		},
	}

	sampleChangeUserPacket = &ChangeUser{
		packet: packet{
			bytes: []byte{
				0x05, 0x00, 0x00, 0x04, // header
				0x11,                   // type
				0x62, 0x6f, 0x62, 0x00, // null terminated "bob"
			},
		},
		user: "bob",
	}

	sampleStatementPreparePacket = &StatementPreparePacket{
		packet: packet{
			bytes: []byte{
				0x09, 0x00, 0x00, 0x00, // header
				0x16,                                           // type
				0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x20, 0x31, // query
			},
		},
		query: "select 1",
	}

	sampleStatementSendLongDataPacket = &StatementSendLongDataPacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x0a, 0x00, 0x00, 0x00, // header
					0x18,                   // type
					0x05, 0x00, 0x00, 0x00, // statement ID
					0x02, 0x00, // parameter ID
					0x62, 0x6f, 0x62, //data
				},
			},
			statementID: 5,
		},
		parameterID: 2,
		data:        []byte("bob"),
	}

	sampleStatementExecutePacket = &StatementExecutePacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x1e, 0x00, 0x00, 0x00, // header
					0x17,                   // type
					0x02, 0x00, 0x00, 0x00, // statement ID
					0x00,                   // cursor flag
					0x01, 0x00, 0x00, 0x00, // iteration count
					0x00, // nullbit map
					0x01, // new-params-bound flag

					// https://dev.mysql.com/doc/internals/en/com-query-response.html#column-type
					0xfe, 0x00, // param 1 type - MYSQL_TYPE_STRING
					0x08, 0x00, // param 2 type - MYSQL_TYPE_LONGLONG

					0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, // param 1 value - "hello"
					0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // param 2 value - 200
				},
			},
			statementID: 2,
		},
		cursorFlag: 0x00,
		iterations: 1,
		nullBitmapAndParameters: []byte{
			0x00,       // null bitmap
			0x01,       // new-params-bound flag
			0xfe, 0x00, // param 1 type - MYSQL_TYPE_STRING
			0x08, 0x00, // param 2 type - MYSQL_TYPE_LONGLONG
			0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, // param 1 value - "hello"
			0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // param 2 value - 200
		},
	}

	sampleStatementClosePacket = &StatementClosePacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x05, 0x00, 0x00, 0x00, // header
					0x19,                   // type
					0x01, 0x00, 0x00, 0x00, // statement ID
				},
			},
			statementID: 1,
		},
	}

	sampleStatementResetPacket = &StatementResetPacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x05, 0x00, 0x00, 0x00, // header
					0x1a,                   // type
					0x01, 0x00, 0x00, 0x00, // statement ID
				},
			},
			statementID: 1,
		},
	}

	sampleStatementFetchPacket = &StatementFetchPacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x09, 0x00, 0x00, 0x00, // header
					0x1c,                   // type
					0x01, 0x00, 0x00, 0x00, // statement ID
					0x0a, 0x00, 0x00, 0x00, // num rows
				},
			},
			statementID: 1,
		},
		rowsCount: 10,
	}

	sampleStatementBulkExecutePacket = &StatementBulkExecutePacket{
		statementIDPacket: statementIDPacket{
			packet: packet{
				bytes: []byte{
					0x15, 0x00, 0x00, 0x00, // header
					0xfa,                   // type
					0x01, 0x00, 0x00, 0x00, // statement ID
					0x80, 0x00, // bulkFlag
					0xfe, 0x00, // param 1 type - MYSQL_TYPE_STRING
					0x08, 0x00, // param 2 type - MYSQL_TYPE_LONGLONG
					0x01,                                                 // param 1 - null
					0x00, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // param 2 value - 200
				},
			},
			statementID: 1,
		},
		bulkFlag: 128,
		parameters: []byte{
			0xfe, 0x00, // param 1 type - MYSQL_TYPE_STRING
			0x08, 0x00, // param 2 type - MYSQL_TYPE_LONGLONG
			0x01,                                                 // param 1 - null
			0x00, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // param 2 value - 200
		},
	}
)

func TestParsePacket(t *testing.T) {
	tests := []struct {
		name           string
		input          io.Reader
		expectedPacket Packet
		expectErrorIs  func(error) bool
	}{
		{
			name:          "network error",
			input:         iotest.ErrReader(&net.OpError{}),
			expectErrorIs: trace.IsConnectionProblem,
		},
		{
			name:          "not enough data for header",
			input:         bytes.NewBuffer([]byte{0x00}),
			expectErrorIs: isUnexpectedEOFError,
		},
		{
			name:          "not enough data for payload",
			input:         bytes.NewBuffer([]byte{0xff, 0xff, 0xff, 0x00, 0x01}),
			expectErrorIs: isUnexpectedEOFError,
		},
		{
			name: "unrecognized type",
			input: bytes.NewBuffer([]byte{
				0x01, 0x00, 0x00, 0x00, // header
				0x44, // type
			}),
			expectedPacket: &Generic{
				packet: packet{
					bytes: []byte{0x01, 0x00, 0x00, 0x00, 0x44},
				},
			},
		},
		{
			name:           "OK_HEADER",
			input:          bytes.NewBuffer(sampleOKPacket.Bytes()),
			expectedPacket: sampleOKPacket,
		},
		{
			name:           "ERR_HEADER",
			input:          bytes.NewBuffer(sampleErrorPacket.Bytes()),
			expectedPacket: sampleErrorPacket,
		},
		{
			name:           "ERR_HEADER protocol 4.1",
			input:          bytes.NewBuffer(sampleErrorWithSQLStatePacket.Bytes()),
			expectedPacket: sampleErrorWithSQLStatePacket,
		},
		{
			name:           "COM_QUERY",
			input:          bytes.NewBuffer(sampleQueryPacket.Bytes()),
			expectedPacket: sampleQueryPacket,
		},
		{
			name:           "COM_QUIT",
			input:          bytes.NewBuffer(sampleQuitPacket.Bytes()),
			expectedPacket: sampleQuitPacket,
		},
		{
			name:           "COM_CHANGE_USER",
			input:          bytes.NewBuffer(sampleChangeUserPacket.Bytes()),
			expectedPacket: sampleChangeUserPacket,
		},
		{
			name:           "COM_STMT_PREPARE",
			input:          bytes.NewBuffer(sampleStatementPreparePacket.Bytes()),
			expectedPacket: sampleStatementPreparePacket,
		},
		{
			name:           "COM_STMT_SEND_LONG_DATA",
			input:          bytes.NewBuffer(sampleStatementSendLongDataPacket.Bytes()),
			expectedPacket: sampleStatementSendLongDataPacket,
		},
		{
			name:           "COM_STMT_EXECUTE",
			input:          bytes.NewBuffer(sampleStatementExecutePacket.Bytes()),
			expectedPacket: sampleStatementExecutePacket,
		},
		{
			name:           "COM_STMT_CLOSE",
			input:          bytes.NewBuffer(sampleStatementClosePacket.Bytes()),
			expectedPacket: sampleStatementClosePacket,
		},
		{
			name:           "COM_STMT_RESET",
			input:          bytes.NewBuffer(sampleStatementResetPacket.Bytes()),
			expectedPacket: sampleStatementResetPacket,
		},
		{
			name:           "COM_STMT_FETCH",
			input:          bytes.NewBuffer(sampleStatementFetchPacket.Bytes()),
			expectedPacket: sampleStatementFetchPacket,
		},
		{
			name:           "COM_STMT_BULK_EXECUTE",
			input:          bytes.NewBuffer(sampleStatementBulkExecutePacket.Bytes()),
			expectedPacket: sampleStatementBulkExecutePacket,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actualPacket, err := ParsePacket(test.input)
			if test.expectErrorIs != nil {
				require.Error(t, err)
				require.True(t, test.expectErrorIs(err))
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expectedPacket, actualPacket)
			}
		})
	}
}

func isUnexpectedEOFError(err error) bool {
	return trace.Unwrap(err) == io.ErrUnexpectedEOF
}
