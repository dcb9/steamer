package youku

type YoukuStream struct {
	Format       string
	Container    string
	VideoProfile string
	Size         uint64
}

func (s YoukuStream) DlWith() string {
	return "--format=" + s.Format
}
