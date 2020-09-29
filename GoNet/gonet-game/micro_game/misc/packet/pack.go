package packet

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

type FastPack interface {
	Pack(w *Packet)
}

func Pack(tos int16, tbl interface{}, w *Packet) []byte {
	if w == nil {
		w = Writer()
	}
	w.WriteS16(tos)
	if tbl == nil {
		return w.Data()
	}
	if fastPack, ok := tbl.(FastPack); ok {
		fastPack.Pack(w)
		return w.Data()
	}
	_pack(reflect.ValueOf(tbl), w)
	return w.Data()
}

func _pack(v reflect.Value, w *Packet) {
	switch v.Kind() {
	case reflect.Bool:
		w.WriteBool(v.Bool())
	case reflect.Uint8:
		w.WriteByte(byte(v.Uint()))
	case reflect.Uint16:
		w.WriteU16(uint16(v.Uint()))
	case reflect.Uint32:
		w.WriteU32(uint32(v.Uint()))
	case reflect.Uint64:
		w.WriteU64(v.Uint())
	case reflect.Int16:
		w.WriteS16(int16(v.Int()))
	case reflect.Int32:
		w.WriteS32(int32(v.Int()))
	case reflect.Int64:
		w.WriteS64(v.Int())
	case reflect.Float32:
		w.WriteFloat32(float32(v.Float()))
	case reflect.Float64:
		w.WriteFloat64(v.Float())
	case reflect.String:
		w.WriteString(v.String())
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			_pack(v.Elem(), w)
		}
	case reflect.Slice:
		if bs, ok := v.Interface().([]byte); ok {
			w.WriteBytes(bs)
		} else {
			count := v.Len()
			w.WriteU16(uint16(count))
			for i := 0; i < count; i++ {
				_pack(v.Index(i), w)
			}
		}
	case reflect.Struct:
		fields := v.NumField()
		for i := 0; i < fields; i++ {
			_pack(v.Field(i), w)
		}
	default:
		log.Error("cannot pack type:", v)
	}
}
