package segment

import (
	"fmt"

	"github.com/mitch000001/go-hbci/domain"
	"github.com/mitch000001/go-hbci/element"
)

type MessageAcknowledgement struct {
	Segment
	acknowledgements []*element.AcknowledgementDataElement
}

func (m *MessageAcknowledgement) version() int         { return 2 }
func (m *MessageAcknowledgement) id() string           { return "HIRMG" }
func (m *MessageAcknowledgement) referencedId() string { return "" }
func (m *MessageAcknowledgement) sender() string       { return senderBank }

func (m *MessageAcknowledgement) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], m)
	if err != nil {
		return err
	}
	m.Segment = seg
	elements = elements[1:]
	acknowledgements := make([]*element.AcknowledgementDataElement, len(elements))
	for i, elem := range elements {
		ack := new(element.AcknowledgementDataElement)
		err := ack.UnmarshalHBCI(elem)
		if err != nil {
			return err
		}
		acknowledgements[i] = ack
	}
	m.acknowledgements = acknowledgements
	return nil
}

func (m *MessageAcknowledgement) Acknowledgements() []domain.Acknowledgement {
	acknowledgements := make([]domain.Acknowledgement, len(m.acknowledgements))
	for i, ackDe := range m.acknowledgements {
		ack := ackDe.Val()
		ack.Type = domain.MessageAcknowledgement
		acknowledgements[i] = ack
	}
	return acknowledgements
}

func (m *MessageAcknowledgement) elements() []element.DataElement {
	dataElements := make([]element.DataElement, len(m.acknowledgements))
	for i, ack := range m.acknowledgements {
		dataElements[i] = ack
	}
	return dataElements
}

func (m *MessageAcknowledgement) Value() []domain.Acknowledgement {
	acks := make([]domain.Acknowledgement, len(m.acknowledgements))
	for i, de := range m.acknowledgements {
		acks[i] = de.Val()
	}
	return acks
}

type SegmentAcknowledgement struct {
	Segment
	acknowledgements []*element.AcknowledgementDataElement
}

func (s *SegmentAcknowledgement) version() int         { return 2 }
func (s *SegmentAcknowledgement) id() string           { return "HIRMS" }
func (s *SegmentAcknowledgement) referencedId() string { return "" }
func (s *SegmentAcknowledgement) sender() string       { return senderBank }

func (s *SegmentAcknowledgement) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], s)
	if err != nil {
		return err
	}
	s.Segment = seg
	elements = elements[1:]
	acknowledgements := make([]*element.AcknowledgementDataElement, len(elements))
	for i, elem := range elements {
		ack := new(element.AcknowledgementDataElement)
		err := ack.UnmarshalHBCI(elem)
		if err != nil {
			return err
		}
		acknowledgements[i] = ack
	}
	s.acknowledgements = acknowledgements
	return nil
}

func (s *SegmentAcknowledgement) Acknowledgements() []domain.Acknowledgement {
	acknowledgements := make([]domain.Acknowledgement, len(s.acknowledgements))
	for i, ackDe := range s.acknowledgements {
		ack := ackDe.Val()
		ack.Type = domain.SegmentAcknowledgement
		acknowledgements[i] = ack
	}
	return acknowledgements
}

func (s *SegmentAcknowledgement) elements() []element.DataElement {
	dataElements := make([]element.DataElement, len(s.acknowledgements))
	for i, ack := range s.acknowledgements {
		dataElements[i] = ack
	}
	return dataElements
}
