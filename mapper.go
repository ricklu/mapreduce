package mapreduce

type mapper interface {
    Map(key, value string) string
}
