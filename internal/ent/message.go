// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"websocket-webchat-demo/internal/ent/message"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// SentAt holds the value of the "sent_at" field.
	SentAt time.Time `json:"sent_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldUsername, message.FieldMessage:
			values[i] = new(sql.NullString)
		case message.FieldSentAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				m.Username = value.String
			}
		case message.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				m.Message = value.String
			}
		case message.FieldSentAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sent_at", values[i])
			} else if value.Valid {
				m.SentAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("username=")
	builder.WriteString(m.Username)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(m.Message)
	builder.WriteString(", ")
	builder.WriteString("sent_at=")
	builder.WriteString(m.SentAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}