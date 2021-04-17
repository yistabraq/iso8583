package spec

type Adjuster interface  {
	Get(string) string
	Set(string) string
}