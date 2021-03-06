package driver

import (
	"fmt"
	"testing"

	"github.com/edgexfoundry/edgex-go/pkg/clients/logging"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

func init() {
	driver = new(Driver)
	driver.Logger = logger.NewClient("test", false, "", "DEBUG")
}

func TestNewResult_bool(t *testing.T) {
	var reading interface{} = true
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "bool"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.BoolValue()
	if val != true || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_uint8(t *testing.T) {
	var reading interface{} = uint8(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint8Value()
	if val != uint8(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_int8(t *testing.T) {
	var reading interface{} = int8(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int8Value()
	if val != int8(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResultFailed_int8(t *testing.T) {
	var reading interface{} = int8(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int16"},
		},
	}
	ro := models.ResourceOperation{}

	_, err := newResult(deviceObject, ro, reading)
	if err == nil {
		t.Errorf("Convert new result should be failed")
	}
}

func TestNewResult_uint16(t *testing.T) {
	var reading interface{} = uint16(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint16"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint16Value()
	if val != uint16(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_int16(t *testing.T) {
	var reading interface{} = int16(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int16"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int16Value()
	if val != int16(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_uint32(t *testing.T) {
	var reading interface{} = uint32(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint32Value()
	if val != uint32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_int32(t *testing.T) {
	var reading interface{} = int32(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int32Value()
	if val != int32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_uint64(t *testing.T) {
	var reading interface{} = uint64(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint64Value()
	if val != uint64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_int64(t *testing.T) {
	var reading interface{} = int64(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int64Value()
	if val != int64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float32(t *testing.T) {
	var reading interface{} = float32(123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "Float32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float32Value()
	if val != float32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64(t *testing.T) {
	var reading interface{} = float64(0.123)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "Float64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float64Value()
	if val != float64(0.123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToInt8(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int8Value()
	if val != int8(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToInt16(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int16"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int16Value()
	if val != int16(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToInt32(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int32Value()
	if val != int32(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToInt64(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int32Value()
	if val != int32(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToUint8(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint8Value()
	if val != uint8(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToUint16(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint16"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint16Value()
	if val != uint16(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToUint32(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint32Value()
	if val != uint32(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToUint64(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint64Value()
	if val != uint64(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToFloat32(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "float32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float32Value()
	if val != float32(reading.(float64)) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_float64ToString(t *testing.T) {
	var reading interface{} = float64(123.11)
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "string"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.StringValue()
	if val != fmt.Sprintf("%v", reading) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_string(t *testing.T) {
	var reading interface{} = "test string"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "string"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.StringValue()
	if val != "test string" || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToFloat32(t *testing.T) {
	var reading interface{} = "123.0"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "float32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float32Value()
	if val != float32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToFloat64(t *testing.T) {
	var reading interface{} = "123.0"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "float64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float64Value()
	if val != float64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToInt64(t *testing.T) {
	var reading interface{} = "123"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int64Value()
	if val != int64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToInt8(t *testing.T) {
	var reading interface{} = "123"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "int8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Int8Value()
	if val != int8(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToUint8(t *testing.T) {
	var reading interface{} = "123"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint8"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint8Value()
	if val != uint8(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToUint32(t *testing.T) {
	var reading interface{} = "123"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint32Value()
	if val != uint32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToUint64(t *testing.T) {
	var reading interface{} = "123"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint64Value()
	if val != uint64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_stringToBool(t *testing.T) {
	var reading interface{} = "true"
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "bool"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.BoolValue()
	if val != true || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_numberToUint64(t *testing.T) {
	var reading interface{} = 123
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "uint64"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Uint64Value()
	if val != uint64(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_floatNumberToFloat32(t *testing.T) {
	var reading interface{} = 123.0
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "float32"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.Float32Value()
	if val != float32(123) || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}

func TestNewResult_numberToString(t *testing.T) {
	var reading interface{} = 123
	deviceObject := models.DeviceObject{
		Properties: models.ProfileProperty{
			Value: models.PropertyValue{Type: "string"},
		},
	}
	ro := models.ResourceOperation{}

	cmdVal, _ := newResult(deviceObject, ro, reading)
	val, err := cmdVal.StringValue()
	if val != "123" || err != nil {
		t.Errorf("Convert new result(%v) failed, error: %v", val, err)
	}
}
