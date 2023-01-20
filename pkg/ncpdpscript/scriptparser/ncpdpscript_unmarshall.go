package scriptparser

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Unmarshal parses the data parameter and populates the target parameter v to
// call this function, you must past a pointer to structure in the second
// parameter.
func Unmarshal(data []byte, v any) error {

	msg, err := NewNCPDPMessage(data)
	if err != nil {
		log.Panic(err)
	}

	indirectVal := reflect.Indirect(reflect.ValueOf(v))

	if !indirectVal.CanSet() {
		return fmt.Errorf("input interface is not addressable (can't Set the memory address): %#v", v)
	}
	if indirectVal.Kind() != reflect.Struct {
		return fmt.Errorf("unmarshalling only accepts pointers to structure in the second param. %v does not meet this requirement", indirectVal.Kind())
	}

	//populate group 0 then check for addition groups
	groupIndex := 0
	err2 := populateGroup(indirectVal, groupIndex, msg)
	if err2 != nil {
		return err2
	}
	//find the slice if it exists and populate it with the additional groups
	additionalGroupsPresent := false
	var additionalGroupsTagId string
	additionalGroupFieldIndex := -1
	for i := 0; i < indirectVal.NumField(); i++ {
		field := indirectVal.Field(i)
		if field.Kind() == reflect.Slice {
			ok := false
			additionalGroupsTagId, ok = indirectVal.Type().Field(i).Tag.Lookup("ncpdp")
			if ok {
				additionalGroupsPresent = true
				additionalGroupFieldIndex = i
				break
			}
		}
	}

	if additionalGroupsPresent {
		groupIndex = 1
		for msg.isGroupPresent(groupIndex) {
			log.Printf("%v %v %v", additionalGroupsTagId, additionalGroupsTagId, additionalGroupFieldIndex)
			var newGroup reflect.Value
			nElem := indirectVal.Field(additionalGroupFieldIndex).Type().Elem()
			strucIsPointer := false
			if nElem.Kind() == reflect.Pointer {
				strucIsPointer = true
				nElem = nElem.Elem()
			}
			log.Printf("%v", strucIsPointer)
			newGroup = reflect.New(nElem).Elem()
			newIndirectVal := reflect.Indirect(newGroup)
			errPop := populateGroup(newIndirectVal, groupIndex, msg)
			if errPop != nil {
				return errPop
			}
			if strucIsPointer {
				indirectVal.Field(additionalGroupFieldIndex).Set(reflect.Append(indirectVal.Field(additionalGroupFieldIndex), newIndirectVal.Addr()))
			} else {
				indirectVal.Field(additionalGroupFieldIndex).Set(reflect.Append(indirectVal.Field(additionalGroupFieldIndex), newIndirectVal))
			}
			groupIndex++
		}
	}

	return nil
	//json.Unmarshal()

}

func populateGroup(indirectVal reflect.Value, groupIndex int, msg *ncpdpMessage) error {

	for i := 0; i < indirectVal.NumField(); i++ {
		field := indirectVal.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {

			segIdentifier, ok := indirectVal.Type().Field(i).Tag.Lookup("ncpdp")
			if !ok {
				return errors.New("error all message fields which are structure must point to segment identifiers through tags")
			}

			tagParts := strings.Split(segIdentifier, ":")
			if len(tagParts) < 2 {
				return errors.New("invalid tag %s.  Tags must contain a groupId:segment id, withing a group, the groupid is not present so it would look like this ':5' where 5 is the segment id")
			}
			//segId := tagParts[1]
			if field.CanSet() && msg.isSegmentPresent(groupIndex, tagParts[1]) {

				val := reflect.New(field.Type().Elem())
				field.Set(val)
				indirectFieldVal := reflect.Indirect(field)
				for i2 := 0; i2 < indirectFieldVal.Type().NumField(); i2++ {
					log.Printf("%s - %v", indirectFieldVal.Type().Field(i2).Name, indirectFieldVal.Type().Field(i2).Tag)
					tag, ok := indirectFieldVal.Type().Field(i2).Tag.Lookup("ncpdp")
					if ok {
						log.Printf("tag to lookup:%s", tag)
						err := attachFieldValueToStruct(msg, tag, indirectFieldVal, i2, groupIndex)
						if err != nil {
							return err
						}
					}
				}
				//reflect.Indirect(field).Type().Field(0).Tag
				log.Printf("here")
			}
		}
	}

	return nil
}

func attachFieldValueToStruct(msg *ncpdpMessage, tag string, indirectFieldVal reflect.Value, fieldIndex int, groupId int) error {
	ncpdpField, err := msg.groups[groupId].GetFieldById(tag)
	if err == nil && ncpdpField != nil {
		log.Printf("%+v", ncpdpField)
		fieldIsPointer := false
		var fieldKind = indirectFieldVal.Type().Field(fieldIndex).Type.Kind()
		if fieldKind == reflect.Slice {
			//this is a repeating field
			var fieldTypeName string
			if indirectFieldVal.Type().Field(fieldIndex).Type.Elem().Kind() == reflect.Pointer {
				fieldIsPointer = true
				fieldKind = indirectFieldVal.Type().Field(fieldIndex).Type.Elem().Elem().Kind()
				fieldTypeName = indirectFieldVal.Type().Field(fieldIndex).Type.Elem().Elem().Name()
			} else {
				fieldKind = indirectFieldVal.Type().Field(fieldIndex).Type.Elem().Kind()
				fieldTypeName = indirectFieldVal.Type().Field(fieldIndex).Type.Elem().Name()
				fieldIsPointer = false
			}
			fieldArray, err := msg.groups[groupId].GetFieldArrayById(tag)
			if errors.Unwrap(err) != NCPDPSegmentNotFound {
				return err
			}
			err = attachValueToFieldArray(fieldKind, fieldTypeName, fieldIsPointer, indirectFieldVal, fieldIndex, fieldArray)
			return err

		} else {

			fieldTypeName := indirectFieldVal.Field(fieldIndex).Type().Name()
			if fieldKind == reflect.Pointer {
				fieldKind = indirectFieldVal.Field(fieldIndex).Type().Elem().Kind()
				fieldTypeName = indirectFieldVal.Field(fieldIndex).Type().Elem().Name()
				fieldIsPointer = true
			}
			err := attachValueToField(fieldKind, fieldTypeName, fieldIsPointer, indirectFieldVal, fieldIndex, ncpdpField)
			return err
		}

	}
	return nil
}

func attachValueToField(fieldKind reflect.Kind, fieldTypeName string, fieldIsPointer bool, indirectFieldVal reflect.Value, fieldIndex int, ncpdpField *field) error {
	switch fieldKind {
	case reflect.String:
		if fieldIsPointer {
			indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(&ncpdpField.fieldData))

		} else {
			indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(ncpdpField.fieldData))
		}
	case reflect.Int:
		intVal, intErr := ncpdpField.int()
		if intErr == nil {
			if fieldIsPointer {
				indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(intVal))
			} else {
				indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(*intVal))
			}
		}
		return intErr
	case reflect.Float64:
		if fieldTypeName == "Currency64" {
			fl64Val, flErr := ncpdpField.Currency64()
			if flErr == nil {
				if fieldIsPointer {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(fl64Val))
				} else {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(*fl64Val))
				}
			}
			return flErr
		} else {
			fl64Val, flErr := ncpdpField.Float64()
			if flErr == nil {
				if fieldIsPointer {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(fl64Val))
				} else {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(*fl64Val))
				}
			}
			return flErr
		}
	case reflect.Float32:
		if fieldTypeName == "Currency32" {
			fl32Val, flErr := ncpdpField.Currency32()
			if flErr == nil {
				if fieldIsPointer {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(fl32Val))
				} else {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(*fl32Val))
				}
			}
			return flErr
		} else {
			fl32Val, flErr := ncpdpField.Float32()
			if flErr == nil {
				if fieldIsPointer {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(fl32Val))
				} else {
					indirectFieldVal.Field(fieldIndex).Set(reflect.ValueOf(*fl32Val))
				}
			}
			return flErr
		}
	}

	return nil
}
func attachValueToFieldArray(fieldKind reflect.Kind, fieldTypeName string, fieldIsPointer bool, indirectFieldVal reflect.Value, fieldIndex int, ncpdpField []*field) error {
	switch fieldKind {
	case reflect.String:
		if fieldIsPointer {
			for _, f := range ncpdpField {
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(&f.fieldData))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}

		} else {
			for _, f := range ncpdpField {
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(f.fieldData))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}

		}
	case reflect.Int:

		for _, f := range ncpdpField {
			val, err := f.int()
			if err != nil {
				return err
			}
			b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(val))
			indirectFieldVal.Field(fieldIndex).Set(b)
		}
	case reflect.Float64:
		if fieldTypeName == "Currency64" {
			for _, f := range ncpdpField {
				val, err := f.Currency64()
				if err != nil {
					return err
				}
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(val))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}
		} else {
			for _, f := range ncpdpField {
				val, err := f.Float64()
				if err != nil {
					return err
				}
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(val))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}
		}

	case reflect.Float32:
		if fieldTypeName == "Currency64" {
			for _, f := range ncpdpField {
				val, err := f.Currency32()
				if err != nil {
					return err
				}
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(val))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}
		} else {
			for _, f := range ncpdpField {
				val, err := f.Float32()
				if err != nil {
					return err
				}
				b := reflect.Append(indirectFieldVal.Field(fieldIndex), reflect.ValueOf(val))
				indirectFieldVal.Field(fieldIndex).Set(b)
			}
		}

	}

	return nil
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "ncpdp: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "ncpdp: Unmarshal(nil " + e.Type.String() + ")"
}
