package jsoniter

import (
	"io"
)

func (cfg *frozenConfig) BorrowStream(writer io.Writer) *Stream {
	select {
	case stream := <-cfg.streamPool:
		stream.Reset(writer)
		return stream
	default:
		return NewStream(cfg, writer, 512)
	}
}

func (cfg *frozenConfig) ReturnStream(stream *Stream) {
	stream.Error = nil
	select {
	case cfg.streamPool <- stream:
		return
	default:
		return
	}
}

func (cfg *frozenConfig) BorrowIterator(data []byte) *Iterator {
	select {
	case iter := <-cfg.iteratorPool:
		iter.ResetBytes(data)
		return iter
	default:
		return ParseBytes(cfg, data)
	}
}

func (cfg *frozenConfig) ReturnIterator(iter *Iterator) {
	iter.Error = nil
	select {
	case cfg.iteratorPool <- iter:
		return
	default:
		return
	}
}
