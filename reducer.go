package mapreduce

type reducer interface {
    Reduce(key string, values []string) string
}
