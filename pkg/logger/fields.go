package logger

var DEBUG = true

func Error(err error) Field {
	return Field{
		Key: "error",
		Val: err,
	}
}

func Bool(key string, value bool) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func String(key, value string) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func Uint(key string, value uint) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func Int64(key string, value int64) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func Int32(key string, value int32) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func Int(key string, value int) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func Any(key string, value any) Field {
	return Field{
		Key: key,
		Val: value,
	}
}

func SafeString(key string, value string) Field {
	if DEBUG {
		return Field{Key: key, Val: value}
	} else {
		return Field{Key: key, Val: "*****"}
	}
}
