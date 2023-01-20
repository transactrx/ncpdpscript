package scriptparser

import (
	"bytes"
	"fmt"
	"strings"
)

type group struct {
	groupPosition int
	segments      []*segment
	segmentsMap   map[string]*segment
}

func newGroup(data []byte, groupPosition int) (*group, error) {
	grp := group{groupPosition: groupPosition}
	grp.segmentsMap = make(map[string]*segment)

	grpBytes := bytes.Split(data, []byte(SegmentSeparator))

	for i, segBytes := range grpBytes {
		if groupPosition == 0 && i == 0 {
			seg, err := newSegmentHeader(segBytes)
			if err != nil {
				return nil, err
			}
			grp.addSegment(seg)
			continue
		}
		if i == 0 {
			//first seg separator empty seg
			continue
		}
		seg, err := newSegment(segBytes)
		if err != nil {
			return nil, err
		}
		grp.addSegment(seg)
	}
	return &grp, nil
}

func (g *group) String() (string, error) {
	result := ""

	for _, seg := range g.segments {

		segStr, err := seg.String()
		if err != nil {
			return "", err
		}
		result += segStr
	}

	if g.groupPosition != 0 {
		result = GroupSeparator + result
	}
	return result, nil

}
func (g *group) addSegment(seg *segment) {
	g.segmentsMap[seg.id] = seg
	g.segments = append(g.segments, seg)
}
func (g *group) GetFieldById(id string) (*field, error) {
	if strings.HasPrefix(id, "header:") {
		//header
		if !strings.Contains(id, ":") {
			return nil, fmt.Errorf("header field keys must be formated %s, key %s is invalid. %w", "header:bin", id, RequestStructInvalid)
		}
		segId := "header"
		fieldId := strings.Split(id, ":")[1]
		return g.getField(segId, fieldId)
	}

	if len(id) != 4 {
		return nil, fmt.Errorf("field ID %s is not valid for looking up field, it must be 4 charcters long. %w", id, RequestStructInvalid)
	}
	segId := id[:2]
	fieldId := id[2:]
	return g.getField(segId, fieldId)

}
func (g *group) GetFieldArrayById(id string) ([]*field, error) {

	if strings.HasPrefix(id, "header") {
		f, err := g.GetFieldById(id)
		if err != nil {
			return nil, err
		} else {
			res := make([]*field, 1)
			res[0] = f
			return res, nil
		}
		return nil, fmt.Errorf("header is not allowed to have repeating segment: %w", RequestStructInvalid)
	}

	segId := id[:2]
	fieldId := id[2:]
	return g.getFieldArrayById(segId, fieldId)
}

func (g *group) getFieldArrayById(segId, fieldId string) ([]*field, error) {
	seg := g.segmentsMap[segId]
	if seg == nil {
		return nil, fmt.Errorf("segment Id %s not found in group. %w", segId, NCPDPSegmentNotFound)
	}
	f := seg.GetFieldArrayById(fieldId)
	if f == nil {
		return nil, fmt.Errorf("segment %s found, but field with id %s not found within it. %w", segId, fieldId, NCPDPFieldNotFound)
	}
	return f, nil
}

func (g *group) getField(segId string, fieldId string) (*field, error) {
	seg := g.segmentsMap[segId]
	if seg == nil {
		return nil, fmt.Errorf("segment Id %s not found in group. %w", segId, NCPDPSegmentNotFound)
	}
	f := seg.GetFieldById(fieldId)
	if f == nil {
		return nil, fmt.Errorf("segment %s found, but field with id %s not found within it. %w", segId, fieldId, NCPDPFieldNotFound)
	}
	return f, nil
}

func (g *group) IsSegmentPresent(id string) bool {

	seg := g.segmentsMap[id]
	return seg != nil

}
